package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

type ResponseSt struct {
	A int `json:"a"`
	B int `json:"b"`
}

const unInteger = 2

func main() {
	file, err := os.Open("input.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	resJs, err := ioutil.ReadAll(file)
	if err != nil { // если конец файла
		log.Fatal(err)
	}
	var resObj []ResponseSt
	err = json.Unmarshal(resJs, &resObj)
	if err != nil { // если конец файла
		log.Fatal(err)
	}

	lenRes := len(resObj)
	ch := 0
	findCh(&lenRes, &ch, resObj)

	lenResCal := len(resObj)
	indexI := 0
	var wg sync.WaitGroup
	wg.Add(ch)
	c := make(chan int, ch)

	for i := ch; i > 0; i-- { //indexI <= lenResCal {
		go calculations(resObj, indexI, lenResCal, &wg, c)
		indexI += 2
	}

	wg.Wait()

	req := 0
	for i := ch; i > 0; i-- {
		req += <-c

	}
	fmt.Println(req / ch)
}

func findCh(lenRes *int, ch *int, resObj []ResponseSt) {
	for *lenRes > 0 {
		*lenRes -= unInteger
		*ch++
		if *lenRes <= 0 {
			return
		}
		findCh(lenRes, ch, resObj)
	}
}

func calculations(resObj []ResponseSt, indexI int, lenRes int, wg *sync.WaitGroup, c chan int) {
	j := indexI
	sum := 0
	que := 0
	for i := 0; i < unInteger; i++ {
		que++
		if j >= lenRes {
			continue
		}
		sum += resObj[j].A + resObj[j].B
		j++
	}
	req := sum / que

	c <- req
	defer wg.Done()
}
