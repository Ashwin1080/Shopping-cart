package main
 
import (
	"log"
	"net/http"
	"shopping-cart/models"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("shopping_cart.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB.AutoMigrate(&models.User{}, &models.Cart{}, &models.Item{}, &models.CartItem{}, &models.Order{})
	log.Println("Database migration completed")
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.CreatedAt = time.Now()
	DB.Create(&user)
	c.JSON(http.StatusOK, user)
}

func ListUsers(c *gin.Context) {
	var users []models.User
	DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func LoginUser(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := DB.Where("username = ? AND password = ?", loginData.Username, loginData.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	user.Token = "newly_generated_token"
	DB.Save(&user)
	c.JSON(http.StatusOK, gin.H{"token": user.Token})
}

func CreateItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item.CreatedAt = time.Now()
	DB.Create(&item)
	c.JSON(http.StatusOK, item)
}

func ListItems(c *gin.Context) {
	var items []models.Item
	DB.Find(&items)
	c.JSON(http.StatusOK, items)
}

func CreateCart(c *gin.Context) {
	var cart models.Cart
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cart.CreatedAt = time.Now()
	DB.Create(&cart)
	c.JSON(http.StatusOK, cart)
}

func ListCarts(c *gin.Context) {
	var carts []models.Cart
	DB.Find(&carts)
	c.JSON(http.StatusOK, carts)
}

func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order.CreatedAt = time.Now()
	DB.Create(&order)
	c.JSON(http.StatusOK, order)
}

func ListOrders(c *gin.Context) {
	var orders []models.Order
	DB.Find(&orders)
	c.JSON(http.StatusOK, orders)
}

func main() {
	initDB()

	r := gin.Default()

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Allow your frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	r.POST("/users", CreateUser)
	r.GET("/users", ListUsers)
	r.POST("/users/login", LoginUser)
	r.POST("/items", CreateItem)
	r.GET("/items", ListItems)
	r.POST("/carts", CreateCart)
	r.GET("/carts", ListCarts)
	r.POST("/orders", CreateOrder)
	r.GET("/orders", ListOrders)

	r.Run(":8080")
}
