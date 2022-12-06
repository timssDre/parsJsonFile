package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type ResponseSt struct {
	A int `json:"a"`
	B int `json:"b"`
}

const unInteger = 1000

func main() {
	resJs, err := os.ReadFile("input.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var resObj []ResponseSt
	err = json.Unmarshal(resJs, &resObj)
	if err != nil {
		log.Fatal(err)
	}

	lenRes := len(resObj)

	workerCount := findCh(lenRes)

	ch := make(chan int, 10)

	for i := 0; i < workerCount; i++ {
		go calculations(resObj, i*unInteger, lenRes, ch)
	}

	req := 0
	for i := 0; i < workerCount; i++ {
		req += <-ch

	}

	fmt.Println(req / workerCount)
}

func findCh(lenRes int) int {
	res := lenRes / unInteger
	if lenRes%unInteger > 0 {
		res++
	}
	return res
}

func calculations(resObj []ResponseSt, indexI int, lenRes int, doneCh chan int) {
	i := indexI
	sum := 0
	for ; i < indexI+unInteger; i++ {
		if i >= lenRes {
			break
		}
		sum += resObj[i].A + resObj[i].B
	}
	req := sum / (i - indexI)

	doneCh <- req
}

//func calc(block []ResponseSt, wg *sync.WaitGroup, doneCh chan int) {
//	defer wg.Done()
//
//	result := 0
//
//	for _, item := range block {
//		result += item.A + item.B
//	}
//
//	doneCh <- result / len(block) * 2
//}
