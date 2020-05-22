package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://facebook.com",
		"http://google.com",
		"http://stackoverflow.com",
		"http://golang.org",
	}

	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}
	//stacking of print statements does work but we can use a for loop too
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)

	//for loop

	//for i := 0; i < len(links); i++ {
	//	fmt.Println(<-c)
	//}

	//we can get multiple get request on the same link by passing the link in the channel itself and then again calling th checkLink() fucntion with the recieved link
	//for {
	//	go checkLink(<-c, c)
	//}

	//alternative way to write the same for loop
	for l := range c {
		go checkLink(l, c)
	}

}
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "may be down")
		time.Sleep(3 * time.Second)
		c <- link
		return
	}
	fmt.Println(link, "is up and running")
	time.Sleep(3 * time.Second)
	c <- link
	return
}
