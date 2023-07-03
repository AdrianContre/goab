package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

var (
	totalRequests int64          //counter to monotorize the total requests
	totalErrors   int64          //counter for all errors
	startTime     time.Time      // start time when starts the requests
	wg            sync.WaitGroup // variable that allows us to wait that all requests are finished

)

/*
Given a url and a bool, makes the request to the url
taking care about if KeepAlive is activated or no
*/
func httpRequest(url string, k bool) {
	defer wg.Done()

	client := &http.Client{}

	if k {
		transport := http.DefaultTransport.(*http.Transport).Clone()
		transport.DisableKeepAlives = true
		client.Transport = transport
	}

	resp, err := client.Get(url) // solicitud HTTP to URL,don't take care of the response because we only launch a request
	if err != nil {
		log.Printf("Error making the request %s", err)
		atomic.AddInt64(&totalErrors, 1) //increment atomitically the errors
		return
	}
	resp.Body.Close() //close the response, if don't closed, errors will occur

	atomic.AddInt64(&totalRequests, 1) //increment atomitically the successful requests
}

func main() {
	var (
		n int
		c int
		k bool
	)
	//Define flags for extracting the arguments
	flag.BoolVar(&k, "k", false, "-k option")
	flag.IntVar(&c, "c", 0, "-c option")
	flag.IntVar(&n, "n", 0, "-n option")

	//Analize the flags given
	flag.Parse()

	//After the values of the flags
	args := flag.Args()
	if len(args) != 1 {
		log.Fatal("URL required")
	}

	url := args[0]

	//if 0 values, we assign default values
	if n <= 0 {
		n = 10000
	}
	if c <= 0 {
		c = 100
	}

	//inicialization
	totalRequests = 0
	totalErrors = 0
	startTime = time.Now()

	//loop for initialize all the routines
	for i := 0; i < n; i++ {
		wg.Add(1)
		go httpRequest(url, k)
		//wait if there are more than c go routines at one time
		if (i+1)%c == 0 {
			wg.Wait()
		}
	}

	//waiting to all routines finish
	wg.Wait()

	//computation of time transcurred
	elapsed := time.Since(startTime)

	//Computation of tps, avgLatency and the percentage of error
	tps := float64(totalRequests) / elapsed.Seconds()
	averageLatency := elapsed.Seconds() / float64(totalRequests)
	errorPercentage := (float64(totalErrors) / float64(totalRequests)) * 100

	// Print the results
	fmt.Println(totalRequests)
	fmt.Printf("TPS: %.2f\n", tps)
	fmt.Printf("Latency: %.6f segundos\n", averageLatency)
	fmt.Printf("Requests with error: %d (%.2f%%)\n", totalErrors, errorPercentage)
}
