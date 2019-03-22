package main

import (
	"flag"
	"fmt"
	"gopro/crawlerYoudao/fetcher"
	"gopro/crawlerYoudao/parser"
	"io/ioutil"
	"log"
	"net/http"
)

// Request request the url and return the body and error
func Request(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return string(body), nil

}

const basicTrans = `<div class="trans-container">[\s]*<ul>[\s]*([\s\S]*?)</ul>[\s]*</div>`
const basicItems = `<li>([^>]*)</li>`

var lookup = flag.String("E", "a", "English words")

func main() {

	flag.Parse()
	w := flag.Lookup("E").Value.String()
	fmt.Println(w)
	var url = "https://www.youdao.com/w/" + w
	result, err := fetcher.Fetch([]byte(url))
	if err != nil {
		log.Println("request", err.Error())
	}
	err = parser.BadInput(result)
	if err != nil {
		log.Fatalln(err)
	}
	word, _ := parser.ParseBaseTrans(result)
	fmt.Println("Translation:")
	for i := 0; i < len(word); i++ {
		fmt.Println("\t", word[i])
	}

	//fmt.Println(parser.ParseWebTrans(result))
	webItems := parser.ParseWebTrans(result)
	fmt.Println("Net explanation:")
	for i, item := range webItems {
		fmt.Printf("\t %d. %s\t%s\n", i, item.Recommend, item.Translation)
		fmt.Println("\t", item.URL)
	}

}
