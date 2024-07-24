package postgresGorm

import "time"

type Anomaly struct {
	ID        uint      `gorm:"primaryKey"`
	SessionId string    `gorm:"size:255;not null"`
	Frequency float64   `gorm:"not null"`
	Timestamp time.Time `gorm:"type:timestamptz;not null"`
}
