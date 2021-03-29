package tools

import (
	"log"
	"strconv"
)

//IsExistSlice ...
func IsExistSlice(value string, slice []string) bool {
	for k := range slice {
		if slice[k] == value {
			return true
		}
	}
	return false
}

//StringToInt ...
func StringToInt(value string) int {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("invalid string value `%s` err: %s", value, err)
		return 0
	}
	return intValue
}

//ReverseSign ...
func ReverseSign(value float64) float64 { return value * -1 }
