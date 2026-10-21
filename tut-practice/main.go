
package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const myurl = "https://jsonplaceholder.typicode.com/posts/1"
const postsURL = "https://jsonplaceholder.typicode.com/posts"

func main() {
	fmt.Println("\n========== GET REQUEST ==========")
	performGetRequest(myurl)

	fmt.Println("\n========== GET REQUEST WITH BUILDER ==========")
	performGetRequestWithBuilder(myurl)

	fmt.Println("\n========== POST REQUEST WITH JSON ==========")
	performPostRequest(postsURL)

	fmt.Println("\n========== POST REQUEST WITH FORM DATA ==========")
	performPostFormRequest(postsURL)

	fmt.Println("\n========== URL PARSING ==========")
	parseURL("https://example.com/posts?userId=1&name=Rahul")

	fmt.Println("\n========== URL CONSTRUCTION ==========")
	createURL()
}

// 1. GET request
func performGetRequest(endpoint string) {
	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("GET request error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Response reading error:", err)
		return
	}

	fmt.Println("Status:", resp.StatusCode)
	fmt.Println("Content Length:", resp.ContentLength)
	fmt.Println("Response:", string(body))
}

// 2. GET request using strings.Builder
func performGetRequestWithBuilder(endpoint string) {
	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("GET request error:", err)
		return
	}
	defer resp.Body.Close()

	var builder strings.Builder

	byteCount, err := io.Copy(&builder, resp.Body)
	if err != nil {
		fmt.Println("Response reading error:", err)
		return
	}

	fmt.Println("Status:", resp.StatusCode)
	fmt.Println("Bytes read:", byteCount)
	fmt.Println("Response:", builder.String())
}

// 3. POST request with JSON
func performPostRequest(endpoint string) {
	jsonData := []byte(`{
		"title": "foo",
		"body": "bar",
		"userId": 1
	}`)

	resp, err := http.Post(
		endpoint,
		"application/json",
		bytes.NewReader(jsonData),
	)
	if err != nil {
		fmt.Println("POST request error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Response reading error:", err)
		return
	}

	fmt.Println("Status:", resp.StatusCode)
	fmt.Println("Response:", string(body))
}

// 4. POST request with form data
func performPostFormRequest(endpoint string) {
	data := url.Values{}
	data.Set("title", "foo")
	data.Set("body", "bar")
	data.Set("userId", "1")

	resp, err := http.PostForm(endpoint, data)
	if err != nil {
		fmt.Println("POST form error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Response reading error:", err)
		return
	}

	fmt.Println("Status:", resp.StatusCode)
	fmt.Println("Response:", string(body))
}

// 5. Parse URL and query parameters
func parseURL(rawURL string) {
	result, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("URL parsing error:", err)
		return
	}

	fmt.Println("Scheme:", result.Scheme)
	fmt.Println("Host:", result.Host)
	fmt.Println("Path:", result.Path)
	fmt.Println("Raw Query:", result.RawQuery)

	params := result.Query()

	fmt.Println("userId:", params.Get("userId"))
	fmt.Println("name:", params.Get("name"))

	for key, values := range params {
		fmt.Println("Key:", key, "Values:", values)
	}
}

// 6. Construct a URL
func createURL() {
	createdURL := &url.URL{
		Scheme: "https",
		Host:   "jsonplaceholder.typicode.com",
		Path:   "/posts/1",
	}

	params := url.Values{}
	params.Set("userId", "1")
	params.Set("name", "Rahul")

	createdURL.RawQuery = params.Encode()

	fmt.Println("Created URL:", createdURL.String())
}
