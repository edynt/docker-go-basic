package main

import (
	"fmt"
	"sync"
)

type Singleton struct{}

var (
	instance *Singleton
	once     sync.Once
)

func GetInstance() *Singleton {
	once.Do(func() {
		fmt.Printf("... get instance \n")
		instance = &Singleton{}
	})

	return instance
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			GetInstance()
			fmt.Printf("Singleton instance address: %p \n", instance)
		}()
	}
	wg.Wait()

	// s := GetInstance()

	// fmt.Printf("Singleton instance address: %p \n", s)

	// time.Sleep(2 * time.Second)

	// f := GetInstance()
	// fmt.Printf("Singleton instance address: %p \n", f)
}

// var once sync.Once

// func initialize() {
// 	fmt.Println("...Initializing...")
// }
// func main() {
// 	for i := 0; i < 5; i++ {
// 		once.Do(initialize)
// 		fmt.Println("-->", i+1, "--->")
// 	}
// }
