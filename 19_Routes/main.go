package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func f(url string) {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(body))
}

//
func main() {
	var wg sync.WaitGroup
	urls := []string{
		"https://reactjs.org/",
		"https://www.youtube.com/user/TechGuyWeb/search?query=go",
		"http://govspy.peterbe.com/#goroutines",
	}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			f(url)
		}(url)
	}
	//
	wg.Wait()
}

// See the example in https://golang.org/pkg/sync/#WaitGroup
