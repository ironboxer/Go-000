package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Subset returns a subset of backends according to clientId and subsetSize
func Subset(backends []int, clientID int, subsetSize int) []int {
	subsetCount := len(backends) / subsetSize
	round := clientID / subsetCount
	rand.Seed(int64(round))
	rand.Shuffle(len(backends), func(i, j int) {
		backends[i], backends[j] = backends[j], backends[i]
	})
	subsetID := clientID % subsetCount
	start := subsetID % subsetSize
	return backends[start : start+subsetSize]
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	subsetSize := 10
	clientID := rand.Intn(100)
	backends := make([]int, 100)
	for i := range backends {
		backends[i] = rand.Intn(10000)
	}
	fmt.Printf("clientID: %d, subsetSize: %d\n\nbackends list: %v\n\n", clientID, subsetSize, backends)
	for i := 0; i < 10; i++ {
		backup := make([]int, 100)
		copy(backup, backends)
		subset := Subset(backup, clientID, subsetSize)
		fmt.Printf("subset: %v\n", subset)
	}
}
