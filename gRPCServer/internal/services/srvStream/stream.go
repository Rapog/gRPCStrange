package srvStream

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	"math/rand"
	"server/internal/domain/models"
	"time"
)

type stream struct {
}

func (s stream) Stream(ctx context.Context) (*models.SrvStream, error) {
	//fmt.Printf("here \n")
	connectUuid, err := uuid.NewUUID()
	if err != nil {
		log.Printf("Uuid err: %s\n", err)
		return nil, err
	}
	//sesId := connectUuid.String()
	mean := -10.0 + rand.Float64()*20
	fmt.Printf("mean: %v\n", mean)
	std := 0.3 + rand.Float64()*1.2
	fmt.Printf("std: %v\n", std)
	utcTime := time.Now().UTC()
	freq := mean + std*rand.NormFloat64()
	fmt.Printf("freq: %v\n", freq)

	return &models.SrvStream{
		SessionId: connectUuid.String(),
		Frequency: freq,
		Time:      utcTime,
	}, nil
}

type StreamProvider interface {
	Stream(ctx context.Context) (*models.SrvStream, error)
}

func NewStrmPrvder() StreamProvider {
	return &stream{}
}
