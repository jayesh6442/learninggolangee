package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jayesh6442/onedarkage/storage/one"
)

func main() {
	fmt.Println("hi there we are in the go code base and we are here for end to end master go and make some more monkey")
	one.One()
	val := uuid.New()
	fmt.Println(val)
	go fmt.Println("hi there we are in gorotien")
	time.Sleep(1 * time.Second)
}
