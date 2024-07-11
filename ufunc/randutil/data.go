package randutil

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type RandomGenerator struct {
	rand *rand.Rand
}

func NewRandomGenerator() *RandomGenerator {
	source := rand.NewSource(time.Now().UnixNano())
	return &RandomGenerator{rand: rand.New(source)}
}

func (rg *RandomGenerator) GetRandomInt(max int) int {
	return rg.GetRandomIntRange(0, max)
}

func (rg *RandomGenerator) GetRandomIntRange(min, max int) int {
	if max <= min {
		panic(fmt.Sprintf("invalid range: max(%d) must be greater than min(%d)", max, min))
	}
	return rg.rand.Intn(max-min) + min
}

func (rg *RandomGenerator) GetEvenlyItemFromArray(index, total int, arr []string) string {
	if len(arr) == 0 {
		panic("array cannot be empty")
	}
	itemCount := len(arr)
	indexCellCount := int(math.Ceil(float64(total) / float64(itemCount)))
	itemIndex := int(math.Ceil(float64(index) / float64(indexCellCount)))
	return arr[itemIndex%len(arr)]
}

func (rg *RandomGenerator) GetRandomBoolWithProbability(probability float32) bool {
	return rg.rand.Float32() < probability
}

func (rg *RandomGenerator) GetRandomStringFromArray(items []string) string {
	if len(items) == 0 {
		panic("items array cannot be empty")
	}
	return items[rg.GetRandomInt(len(items))]
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func (rg *RandomGenerator) GetRandomStringBytes(minlen, maxlen int) string {
	n := rg.GetRandomIntRange(minlen, maxlen+1)
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rg.GetRandomInt(len(letterBytes))]
	}
	return string(b)
}
