package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var urls = []string{
	"https://golang.org",
	"https://godoc.org",
	"https://golang.org",
	"https://godoc.org",
	"https://play.golang.org",
	"http://gopl.io",
	"https://golang.org",
	"https://godoc.org",
}

func main() {

	start := time.Now()

	for _, url := range urls {
		result, err := http.Get(url)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		defer result.Body.Close()

		elapsed := time.Since(start).Seconds()

		fmt.Printf("http.Get to %s took %v seconds \n", url, elapsed)

	}

}
