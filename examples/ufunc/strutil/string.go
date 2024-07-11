package main

import (
	"log"

	"github.com/clh021/gutils/ufunc/strutil"
)

func testString() {
	strArr := []string{"a", "b", "c"}
	str := "b"
	log.Println(strArr, str)
	log.Println("Contains: ", strutil.Contains(strArr, str))
}