package radarr_test

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SkYNewZ/radarr"
)

// Instantiate a standard client
func ExampleNew_basic() {
	client, err := radarr.New("https://my.radarr-instance.fr", "radarr-api-key", nil)
	if err != nil {
		log.Fatalln(err)
	}

	movie, err := client.Movies.Get(217)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s", movie.Title)
}

func ExampleNew_verbose() {
	client, err := radarr.New("https://my.radarr-instance.fr", "radarr-api-key", nil, &radarr.ClientOptions{
		Verbose: true,
	})
	if err != nil {
		log.Fatalln(err)
	}

	movie, err := client.Movies.Get(217)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s", movie.Title)
}

// Instantiate a client with a custom HTTP client
func ExampleNew_advanced() {
	client, err := radarr.New("https://my.radarr-instance.fr", "radarr-api-key", &http.Client{
		Timeout: time.Second * 10,
	})
	if err != nil {
		log.Fatalln(err)
	}

	movie, err := client.Movies.Get(217)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s", movie.Title)
}
