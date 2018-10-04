package main

import (
	"fmt"
	"time"

	"github.com/xiaozefeng/go-codes/basic/retriever/mock"
	"github.com/xiaozefeng/go-codes/basic/retriever/real"
)

func main() {
	var r Retriever
	mockRetriever := &mock.Retriever{Contents: "Hello World"}
	r = mockRetriever
	inspect(r)
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Second,
	}
	inspect(r)

	// type assertion
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock mockRetriever")
	}

	fmt.Println()
	// Try a session
	fmt.Println("try a session")
	fmt.Println(session(mockRetriever))
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

// interface compose
type RetrieverPoster interface {
	Retriever
	Poster
}

const url = "http://www.baidu.com"

func post(poster Poster) {
	poster.Post(url, map[string]string{
		"name": "test",
	})
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked baidu.com",
	})
	return s.Get(url)
}

// http request
func Download(r Retriever) string {
	return r.Get(url)
}
