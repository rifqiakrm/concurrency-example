package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	websites := map[string]string{
		"google":  "https://www.google.com",
		"youtube": "https://youtube.com",
		"pornHub": "https://www.pornhub.com",
	}

	for i, v := range websites {
		fmt.Println(checkWebsite(i, v))
	}

	elapsed := time.Since(start)

	fmt.Println("Execution time : ", elapsed.String())
}

func checkWebsite(t, w string) string {
	if _, err := http.Get(w); err != nil {
		return fmt.Sprintf("man, %s is down :(", t)
	}

	return fmt.Sprintf("yeay, %s is up :)", t)
}
