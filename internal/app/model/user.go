package model

type User struct {
    UserID   uint   `gorm:"primaryKey;autoIncrement"`
    Username string `gorm:"size:255;not null"`
    Password string `gorm:"size:255;not null"`
    Email    string `gorm:"size:255"`
}

func (User) TableName() string {
    return "Users"
}
