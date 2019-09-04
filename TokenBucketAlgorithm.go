package SystemdesignAlgorithms

import (
	"sync"
	"time"
)

type TokenBucket struct {
	Capacity       int
	Available      int
	ReFillRate     int64
	LastReFillTime time.Time
	lock           sync.Mutex
}

func NewTokenBucket(capacity int, refillRate int64) {
	t := new(TokenBucket)

	t.Capacity = capacity
	t.Available = capacity
	t.ReFillRate = refillRate
	t.LastReFillTime = time.Now()
}

func (b *TokenBucket) AllowRequest(n int) {

}

func (b *TokenBucket) refill() {
}
