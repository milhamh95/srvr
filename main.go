package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func(n int) {
			defer wg.Done()

			//time.Sleep(time.Duration(n) * 100 * time.Millisecond)

			resp, err := http.Get("http://localhost:1323/counter2")
			if err != nil {
				log.Fatalln(err)
			}

			b, _ := io.ReadAll(resp.Body)
			fmt.Printf("request #%02d: %d %q\n", n, resp.StatusCode, bytes.TrimSpace(b))
		}(i + 1)
	}
}
