package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"godemo/download"
	"time"
)

const (
	URL      = "http://www.quanjing.com/"
	savePath = "E://temp//"
)

func main() {

	buf := make(chan string)
	flg := make(chan int)
	var cache []string

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org", "www.quanjing.com"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("img", func(e *colly.HTMLElement) {
		link := e.Attr("src")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		//c.Visit(e.Request.AbsoluteURL(link))
		imgUrl := URL + link
		cache = append(cache, imgUrl)
		/*fmt.Print(imgUrl)
		down := download.Download{Url:imgUrl}
		down.GetImg2()*/

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("http://www.quanjing.com/")

	go producer(buf, cache)
	go consumer(buf, flg)
	<-flg
}

/*func getImg(url string) (n int64, err error) {
	path := strings.Split(url, "/")
	var name string
	if len(path) > 1 {
		name = savePath + path[len(path)-1]
	}
	fmt.Println(name)
	out, err := os.Create(name)
	defer out.Close()
	resp, err := http.Get(url)
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	n, err = io.Copy(out, bytes.NewReader(pix))
	return

}*/

func producer(c chan string, cache []string) {
	fmt.Print(time.Now())
	defer close(c) // 关闭channel
	for _, v := range cache {
		fmt.Print(v)
		c <- v
	}
	/*for i := 0; i < 10; i++{
		c <- i // 阻塞，直到数据被消费者取走后，才能发送下一条数据
	}*/
}

func consumer(c chan string, f chan int) {
	for {
		if v, ok := <-c; ok {
			fmt.Print(v) // 阻塞，直到生产者放入数据后继续读取数据
			download.GetImg3(v)
		} else {
			break
		}
	}
	f <- 1 //发送数据，通知main函数已接受完成
	fmt.Print(time.Now())
}
