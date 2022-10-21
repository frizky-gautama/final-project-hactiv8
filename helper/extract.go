package helper

import (
	"log"
	"strconv"
)

func StrToInt(val string) int {
	var num, err = strconv.Atoi(val)

	if err != nil {
		log.Fatal("error format")
	}

	return num
}

func StrToInt32(val string) int32 {
	// var num, err = strconv.Atoi(val)
	i64, err := strconv.ParseInt(val, 10, 32)
	i := int32(i64)
	if err != nil {
		log.Fatal("error format")
	}

	return i
}
