package utils

import "fmt"

func CheckError(err error) bool {
	if err != nil {
		fmt.Println("There was an error!")
		return true
	}
	return false
}
