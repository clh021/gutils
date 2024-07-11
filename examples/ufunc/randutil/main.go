package main

import (
	"log"

	"github.com/clh021/gutils/ufunc/randutil"
)

func main() {
	rg := randutil.NewRandomGenerator()
	log.Println("GetRandomInt: ", rg.GetRandomInt(100))
	log.Println("GetRandomIntRange: ", rg.GetRandomIntRange(10, 20))
	log.Println("GetEvenlyItemFromArray: ", rg.GetEvenlyItemFromArray(1, 10, []string{"a", "b", "c"}))
	log.Println("GetRandomBoolWithProbability: ", rg.GetRandomBoolWithProbability(0.5))
	log.Println("GetRandomStringFromArray: ", rg.GetRandomStringFromArray([]string{"a", "b", "c"}))
	log.Println("GetRandomStringBytes: ", rg.GetRandomStringBytes(10, 20))
}