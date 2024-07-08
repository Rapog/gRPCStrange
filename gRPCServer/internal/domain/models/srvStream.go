package models

import "time"

type SrvStream struct {
	SessionId string
	Frequency float64
	Time      time.Time
}
