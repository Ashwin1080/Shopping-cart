To install node modules : npm install in the "shopping-ui" directory 
To run the frontend run the following command : npm run dev
Frontend - http://localhost:5173/ 

To start the server, in the "shopping-cart" directory run the following command : go run main.go
port - 8080

API's included and sample data for testing purpose in postman:

1. Create a New User
---------------------
URL: http://localhost:8080/users
Method: POST
Body (JSON):
{
  "username": "user1",
  "password": "password123"
}

2. Login a User
----------------
URL: http://localhost:8080/users/login
Method: POST
Body (JSON):
{
  "username": "user1",
  "password": "password123"
}

3. Create an Item
------------------
URL: http://localhost:8080/items
Method: POST
Body (JSON):
{
  "name": "Item1",
  "status": "available"
}

4. List All Items
------------------
URL: http://localhost:8080/items
Method: GET

5. Create a Cart
-----------------
URL: http://localhost:8080/carts
Method: POST
Body (JSON):
{
  "user_id": 1,
  "status": "active"
}

6. List All Carts
------------------
URL: http://localhost:8080/carts
Method: GET

7. Create an Order
-------------------
URL: http://localhost:8080/orders
Method: POST
Body (JSON):
{
  "cart_id": 1,
  "user_id": 1
}

8. List All Orders
-------------------
URL: http://localhost:8080/orders
Method: GET
(https://github.com/user-attachments/assets/1935ffcd-4e85-4124-9051-ad43b7c5098c)
