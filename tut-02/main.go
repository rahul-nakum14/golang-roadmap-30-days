package main 

import (
	"fmt"
	"net/http"
	// "io"
	"net/url"
)
const myurl string = "https://jsonplaceholder.typicode.com/posts/1?userId=1"

func main() {

		resp, err := http.Get(myurl)
		if err != nil {
			fmt.Println("Error making GET request:", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Println("Error: received non-OK status code:", resp.StatusCode)
			return
		}

		// databytes, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}
		// fmt.Println("Response body:", string(databytes))

		result , _ := url.Parse(myurl)
		fmt.Println("Scheme:", result.Scheme)
		fmt.Println("Host:", result.Host)
		fmt.Println("Path:", result.Path)
		fmt.Println(result.RawQuery)


		qParams := result.Query()
		fmt.Println("Query Parameters:")
	}