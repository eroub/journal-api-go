package model

import (
    "time"
)

type Journal struct {
    ID              uint       `gorm:"primaryKey;autoIncrement"`
    DatetimeIn      *time.Time // Nullable datetime
    DatetimeOut     *time.Time // Nullable datetime
    TotalHrs        *float64   // Nullable float
    Ticker          *string    // Nullable string
    Direction       *string    `gorm:"type:enum('Long','Short')"` // Nullable enum
    Equity          *float64   // Nullable float
    Entry           *float64   // Nullable float
    StopLoss        *float64   // Nullable float
    Target          *float64   // Nullable float
    Size            *int       // Nullable int
    Risk            *float64   // Nullable float
    EstGain         *float64   // Nullable float
    EstRR           *float64   // Nullable float
    ExitPrice       *float64   // Nullable float
    ProjPL          *float64   // Nullable float
    RealPL          *float64   // Nullable float
    Commission      *float64   // Nullable float
    PercentChange   *float64   // Nullable float
    RealRR          *float64   // Nullable float
    Pips            *float64   // Nullable float
    MFE             *float64   // Nullable float
    MAE             *float64   // Nullable float
    MFERatio        *float64   // Nullable float
    MAERatio        *float64   // Nullable float
    Type            *string    // Nullable string (varchar(255))
    Screenshot      *string    // Nullable string (text)
    RiskUSD         *float64   // Nullable float
    EstGainUSD      *float64   // Nullable float
    ProjPLUSD       *float64   // Nullable float
    Comment         *string    // Nullable string (text)
    Status          string     `gorm:"type:enum('Open','Closed');default:'Open'"` // Non-nullable enum with default
    AccountID       *uint      // Nullable foreign key (int)
}

func (Journal) TableName() string {
    return "trade_journal"
}
