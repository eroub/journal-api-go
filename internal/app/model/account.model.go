package model

type Account struct {
    AccountID        uint    `gorm:"primaryKey;autoIncrement"`
    AccountName      string  `gorm:"size:255;not null"`
    UserID           *uint   // Pointer to uint allows NULL
    Equity           float64 `gorm:"type:decimal(15,2)"`
    DefaultRiskPercent float64 `gorm:"type:decimal(3,2)"`
}

func (Account) TableName() string {
    return "Accounts"
}
