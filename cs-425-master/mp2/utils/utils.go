package utils

import (
	"math/rand"
	"strings"
	"time"

	"../config"

	"hash/fnv"
	"math"
)

// Hash the given string to 32-bit integer
func HashString(str string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(str))
	return h.Sum32() % config.RING_SIZE
}

// Given the hash ring, return the machine ID that corresponds to the lowest entry on the hash ring
func GetLowestHashRingEntry(hashRing map[string]uint32) string {
	var lowestHash uint32 = math.MaxUint32
	var lowestMachine string = ""

	for machineID, hashNum := range hashRing {
		if hashNum < lowestHash {
			lowestMachine = machineID
		}
	}

	return lowestMachine
}

// Get a machine's IP address from its ID
func GetIPFromID(machineID string) string {
	return strings.Split(machineID, ":")[0]
}

// Shuffle the given string array
func ShuffleList(list *[]string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(*list), func(i, j int) {
		(*list)[i], (*list)[j] = (*list)[j], (*list)[i]
	})
}
