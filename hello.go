package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(Hello("Ted"))
	fmt.Println(Hello(""))
}

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
