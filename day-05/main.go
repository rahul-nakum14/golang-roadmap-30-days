package main

import (
	"net/http"
	"fmt"
	"time"
	"sync"
)
var wg sync.WaitGroup

func main() {
	// go greet("Hello")
	// greet("World")

	// websites := []string{
	// 	"https://www.google.com",
	// 	"https://www.facebook.com",
	// 	"https://www.twitter.com",
	// 	"https://www.linkedin.com",
	// }
	// for _, website := range websites {
	// 	wg.Add(1)
	// 	go getStatusCode(website)
	// }
	// wg.Wait()

	wg.Add(2)
	go printNumbers(&wg)
	go printLetters(&wg)
	wg.Wait()
}

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}

func printLetters(wg *sync.WaitGroup) {
	defer wg.Done()
	for ch := 'A'; ch <= 'E'; ch++ {
		fmt.Printf("%c\n", ch)
	}
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res,err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Error:", err)
	}else{
		fmt.Printf("Response Status Code: %d and the endpoint is %s\n", res.StatusCode, endpoint)
	}
}

func greet(name string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		println(name)
	}
}