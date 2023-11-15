package model

import (
    "time"
)

type Transaction struct {
    TransactionID uint `gorm:"primaryKey;autoIncrement"`
    AccountID     *uint // Pointer to uint allows NULL
    Type          string `gorm:"type:enum('Starting_Balance','Withdrawal','Deposit','Inflow','Outflow','Profit','Loss')"`
    Amount        float64
    Date          time.Time
}

func (Transaction) TableName() string {
    return "Transactions"
}
