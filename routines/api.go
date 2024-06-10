package routines

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Api() {
	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com",
		"https://graph.microsoft.com",
	}

	ch := make(chan string)

	for _, url := range apis {
		go checkApi(url, ch)
		log.Println(<-ch)
	}

	elapsed := time.Since(start)
	log.Printf("Time elapsed: %s\n", elapsed)
}

func checkApi(url string, ch chan string) {
	_, err := http.Get(url)

	if err != nil {
		log.Println(err)
	}

	ch <- fmt.Sprintf("{url: %s, status: ok}\n", url)
}
