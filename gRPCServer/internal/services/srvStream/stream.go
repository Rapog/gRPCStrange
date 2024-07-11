package srvStream

import (
	"context"
	"log"
	"math/rand"
	"server/internal/cache"
	"server/internal/domain/models"
	"time"
)

//var (
//	mean       float64
//	isInitMean = false
//	std        float64
//	isInitStd  = false
//	mu         = sync.Mutex{}
//)

type stream struct {
	cache.Cache[*models.MeanStd]
}

func (s stream) Stream(ctx context.Context, sessionId string) (*models.SrvStream, error) {
	//fmt.Printf("here \n")
	//connectUuid, err := uuid.NewUUID()
	//if err != nil {
	//	log.Printf("Uuid err: %s\n", err)
	//	return nil, err
	//}
	//mu.Lock()
	//if !isInitMean {
	//	mean = -10.0 + rand.Float64()*20
	//	isInitMean = true
	//}
	//
	//if !isInitStd {
	//	std = 0.3 + rand.Float64()*1.2
	//	isInitStd = true
	//}
	//mu.Unlock()
	////sesId := connectUuid.String()
	//
	//fmt.Printf("mean: %v\n", mean)
	//
	//fmt.Printf("std: %v\n", std)
	//utcTime := time.Now().UTC()
	//
	//freq := mean + std*rand.NormFloat64()
	//fmt.Printf("freq: %v\n", freq)

	if s.Cache.Has(sessionId) {
		res := s.Cache.Get(sessionId)
		//log.Println("Mean: ", res.Mean)
		//log.Println("STD: ", res.Std)
		return &models.SrvStream{
			SessionId: sessionId,
			Frequency: res.Mean + res.Std*rand.NormFloat64(),
			Time:      time.Now().UTC(),
		}, nil
	} else {
		newMeanStd := &models.MeanStd{
			Mean: -10 + rand.Float64()*20,
			Std:  0.3 + rand.Float64()*1.2,
		}
		s.Cache.Set(sessionId, newMeanStd)

		log.Println("MeanFirst: ", newMeanStd.Mean)
		log.Println("STDFirst: ", newMeanStd.Std)

		return &models.SrvStream{
			SessionId: sessionId,
			Frequency: newMeanStd.Mean + newMeanStd.Std*rand.NormFloat64(),
			Time:      time.Now(),
		}, nil
	}
}

type StreamProvider interface {
	Stream(ctx context.Context, sessionId string) (*models.SrvStream, error)
}

func NewStrmPrvder(cache cache.Cache[*models.MeanStd]) StreamProvider {
	return &stream{Cache: cache}
}
