package main

import (
    "fmt"
    "os"
    "github.com/shodanhq/shodan-go"
)

func main() {
    // Get the Shodan API key from the environment variable
    apiKey := os.Getenv("SHODAN_API_KEY")

    // Create a new Shodan client
    client := shodan.NewClient(apiKey)

    // Get the search query from the user input
    fmt.Print("Enter search query: ")
    var query string
    fmt.Scanln(&query)

    // Search for the given query
    results, err := client.Search(query)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Print the results
    fmt.Println(results)
}
