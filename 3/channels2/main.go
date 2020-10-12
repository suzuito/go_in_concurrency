package main

import (
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"os/signal"
	"sync"
)

type ContentType string

const (
	ContentTypeTweet     ContentType = "tweet"
	ContentType2ch       ContentType = "2ch"
	ContentTypeFacebook  ContentType = "facebook"
	ContentTypeYouTube   ContentType = "youtube"
	ContentTypeTiktok    ContentType = "tiktok"
	ContentTypeYahooNews ContentType = "yahoo_news"
	ContentTypeBakusai   ContentType = "bakusai"
)

type Kind string

const (
	KindSNS     Kind = "sns"
	KindForum   Kind = "forum"
	KindVideo   Kind = "video"
	KindNews    Kind = "news"
	KindUnknown Kind = "unknown"
)

type Content struct {
	Type   ContentType
	Kind   Kind
	Title  string
	Body   string
	RawURL string
	URL    *url.URL
}

func main() {
	contents := []Content{}
	n := 100000
	for i := 0; i < n; i++ {
		c := Content{}
		i := rand.Int()
		if i%1 == 0 {
			c.RawURL = fmt.Sprintf("https://twitter.com/%d", i)
		}
		if i%2 == 0 {
			c.RawURL = fmt.Sprintf("https://2ch.com/%d", i)
		}
		if i%3 == 0 {
			c.RawURL = fmt.Sprintf("https://www.facebook.com/%d", i)
		}
		if i%4 == 0 {
			c.RawURL = fmt.Sprintf("https://www.tiktok.com/%d", i)
		}
		if i%5 == 0 {
			c.RawURL = fmt.Sprintf("https://www.yahoo.co.jp/%d", i)
		}
		if i%6 == 0 {
			c.RawURL = fmt.Sprintf("https://bakusai/%d", i)
		}
		if i%7 == 0 {
			c.RawURL = fmt.Sprintf("http://foo.com/%d", i)
		}
		contents = append(contents, c)
	}

	g := sync.WaitGroup{}

	var src chan *Content = make(chan *Content)
	g.Add(1)
	var dstErr chan error = make(chan error)
	g.Add(1)

	var s1 chan *Content = make(chan *Content)
	g.Add(1)
	var s2 chan *Content = make(chan *Content)
	g.Add(1)
	go parseURL(&g, src, s1, dstErr)
	go whichKind(&g, s1, s2, dstErr)
	go printContent(&g, s2, dstErr)
	go func(g *sync.WaitGroup) {
		for err := range dstErr {
			fmt.Printf("Error is occured: %+v\n", err)
		}
		g.Done()
	}(&g)

	for i := range contents {
		src <- &contents[i]
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		<-sig
		close(src)
		close(dstErr)
		close(s1)
		close(s2)
	}()

	g.Wait()
}

func parseURL(
	g *sync.WaitGroup,
	src <-chan *Content,
	dst chan<- *Content,
	dstErr chan<- error,
) {
	for c := range src {
		var err error
		c.URL, err = url.Parse(c.RawURL)
		if err != nil {
			dstErr <- err
			continue
		}
		dst <- c
	}
	g.Done()
}

func whichKind(
	g *sync.WaitGroup,
	src <-chan *Content,
	dst chan<- *Content,
	dstErr chan<- error,
) {
	for c := range src {
		if c.URL.Host == "twitter.com" || c.URL.Host == "www.facebook.com" {
			c.Kind = KindSNS
		} else if c.URL.Host == "2ch.com" || c.URL.Host == "bakusai" {
			c.Kind = KindForum
		} else if c.URL.Host == "www.tiktok.com" || c.URL.Host == "www.youtube.com" {
			c.Kind = KindVideo
		} else if c.URL.Host == "www.yahoo.co.jp" {
			c.Kind = KindNews
		} else {
			c.Kind = KindUnknown
		}
		dst <- c
	}
	g.Done()
}

func printContent(
	g *sync.WaitGroup,
	src <-chan *Content,
	dstErr chan<- error,
) {
	for c := range src {
		fmt.Println(c)
	}
	g.Done()
}
