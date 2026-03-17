package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No input provided")
		return
	}

	name := os.Args[1]

	// Read API key from environment
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Println("Missing API_KEY")
		return
	}

	url := fmt.Sprintf("https://api.agify.io?name=%s&apikey=%s", name, apiKey)

	fmt.Println("Calling API:", url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error calling API:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Println("API Response:")
	fmt.Println(string(body))
}
