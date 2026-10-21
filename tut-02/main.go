package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)
const myurl string = "https://jsonplaceholder.typicode.com/posts/1"

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
		fmt.Println("Query Parameters:",qParams["userId"])
		fmt.Println("Query Parameters:",qParams["name"])	

		for _, value := range qParams {
			fmt.Println("Value:", value)
		}

		creaeeUrl := &url.URL{
			Scheme: "https",
			Host:   "jsonplaceholder.typicode.com",
			Path:   "/posts/1",
			RawPath: "user rahul",
			RawQuery: "userId=1&name=Rahul",
		}
		fmt.Println("Created URL:", creaeeUrl.String())

		// perFormGetRequest(myurl)
		// perFormPostRequest("https://jsonplaceholder.typicode.com/posts")
		performPostformRequest("https://jsonplaceholder.typicode.com/posts")
	}

func perFormGetRequest(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer resp.Body.Close()

	var reponseString  strings.Builder
	databytes, err := io.ReadAll(resp.Body)
	bytecount, _ := reponseString.Write(databytes)

	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Status code:", resp.StatusCode)
	fmt.Println("Contenet Length:" , resp.ContentLength)
	fmt.Println("Response body:", string(databytes))
	fmt.Println("Byte count:", bytecount)

	fmt.Println(reponseString.String())
}


func perFormPostRequest(url string) {

	requestBody := strings.NewReader(`{
		"title": "foo",
		"body": "bar",
		"userId": 1
	  }`)

	  response,err:=http.Post(url,"application/json", requestBody)
	  if err != nil {
		  fmt.Println("Error making POST request:", err)
		  return
	  }
	  defer response.Body.Close()	
	  content,_:=io.ReadAll(response.Body)
	  fmt.Println(string(content))
}

func performPostformRequest(endpoint  string) {

	data := url.Values{}
	data.Set("title", "foo")
	data.Set("body", "bar")
	data.Set("userId", "1")

	response, err := http.PostForm(endpoint , data)
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}