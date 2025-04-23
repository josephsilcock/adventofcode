package utils

import (
	"log"
	"strconv"
)

func ConvertStringToInt(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	return value

}
