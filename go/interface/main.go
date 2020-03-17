package main

import (
	"fmt"
	"retriever/mock"
	"retriever/real"
	"time"
)

// 使用者規定Retriever interface 有 Get方法
type test interface {
	Get(url string) string
}

func download(r test) string {
	return r.Get("http://www.google.com")
}

func main() {

	var r test
	r = &mock.Retriever{"mock retriever"}
	inspect(r)

	r = real.Retriever{"Mozilla/5.0", time.Minute}
	inspect(r)

	// 用r.(類型) 來判斷是否是該類型
	//if mockRetriever, ok := r.(mock.Retriever) ; ok {
	//	fmt.Println("this is a mockRetriever",mockRetriever)
	//} else {
	//	fmt.Println("this is not a mockretriever",mockRetriever)
	//}
	//
	//if realRetriever, ok := r.(real.Retriever) ; ok {
	//	fmt.Println("this is a realRetriever :",realRetriever)
	//} else {
	//	fmt.Println("this is not a realRetriever",realRetriever)
	//}

}

func inspect(r test) {
	//fmt.Println(download(r))
	fmt.Println("Type switch => ")
	fmt.Printf("%T %v\n", r, r)
	switch r.(type) {
	case *mock.Retriever:
		fmt.Println("this is a mockRetriever :", r)
	case mock.Retriever:
		fmt.Println("this is a mockRetriever :", r)
	case *real.Retriever:
		fmt.Println("this is a realRetriever :", r)
	case real.Retriever:
		fmt.Println("this is a realRetriever :", r)
	}
}
