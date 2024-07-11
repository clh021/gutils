package main

import (
	"log"

	"github.com/clh021/gutils/ufunc/strutil"
)

func testEncode() {
	utf8Str := "这是一段 utf8 的文字"
	gbkStr, err := strutil.Utf8ToGbk([]byte(utf8Str))
	if err != nil {
		log.Println("Utf8ToGbk: ", err)
	}
	log.Println("gbkStr: ", gbkStr)
	utf8Byte, err := strutil.GbkToUtf8(gbkStr)
	if err != nil {
		log.Println("GbkToUtf8: ", err)
	}
	log.Println("utf8Str: ", string(utf8Byte))
}
