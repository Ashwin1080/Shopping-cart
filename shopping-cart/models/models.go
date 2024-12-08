package models

import "time"

type User struct {
    ID        uint      `gorm:"primaryKey"`
    Username  string    `gorm:"unique;not null"`
    Password  string    `gorm:"not null"`
    Token     string    `gorm:"type:text"`
    CartID    *uint
    CreatedAt time.Time
}

type Cart struct {
    ID        uint      `gorm:"primaryKey"`
    UserID    uint      `gorm:"not null"`
    Status    string    `gorm:"default:'active'"`
    CreatedAt time.Time
}

type Item struct {
    ID        uint      `gorm:"primaryKey"`
    Name      string    `gorm:"not null"`
    Status    string    `gorm:"default:'available'"`
    CreatedAt time.Time
}

type CartItem struct {
    CartID uint `gorm:"primaryKey"`
    ItemID uint `gorm:"primaryKey"`
}

type Order struct {
    ID        uint      `gorm:"primaryKey"`
    CartID    uint      `gorm:"not null"`
    UserID    uint      `gorm:"not null"`
    CreatedAt time.Time
}
