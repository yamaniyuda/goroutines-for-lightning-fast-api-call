package main

import (
	"fmt"
	"goroutines-for-lightning-fast-api-call/app/helper"
	"goroutines-for-lightning-fast-api-call/domain/repository"
	"sync"
	"time"
)

func main() {
	defer helper.TimeCheck(time.Now(), "getting time")
	fmt.Println("Starting synchronouse calls...")

	var waitGroup sync.WaitGroup
	waitGroup.Add(10)

	go func() {
		repository.GetNatureData()
		waitGroup.Done()
	}()

	go func() {
		repository.GetGenerationData()
		waitGroup.Done()
	}()

	waitGroup.Wait()
}
