package main

import "sync"

import "fmt" 
import "net/http"

func main()  {
	var wg sync.WaitGroup

	var urls = []string{
		"http://www.baidu.com",
		"http://www.sina.com",
		"http://www.qiniu.com",
	}

	for _, url := range urls {
		wg.Add(1)
		go func (url string)  {
			defer wg.Done()
			_, err := http.Get(url)
			fmt.Println(url, err)
		}(url)
	}
	wg.Wait()
	fmt.Println("over")
}