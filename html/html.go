package html

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

func Titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			title := r.FindStringSubmatch(string(html))

			if cap(title) == 0 {
				c <- "Cannot get title from " + url
				return
			}

			c <- title[1]
		}(url)
	}
	return c
}
