package parser

import (
	"errors"
	"regexp"
	"strings"
)

const basicTrans = `<div class="trans-container">[\s]*<ul>[\s]*([\s\S]*?)</ul>[\s]*</div>`
const basicItems = `<li>([^>]*)</li>`

// ParseBaseTrans prase basic translation
func ParseBaseTrans(content []byte) ([]string, error) {
	re := regexp.MustCompile(basicTrans)
	matchers := re.FindAllSubmatch(content, 1)
	var trans []string
	re = regexp.MustCompile(basicItems)
	match := re.FindAllSubmatch(matchers[0][1], -1)
	for i := 0; i < len(match); i++ {
		trans = append(trans, string(match[i][1]))
	}
	return trans, nil
}

// WebTrans hold web translation
type WebTrans struct {
	Recommend   string
	Translation string
	URL         string
}

const webTrans = `<div class="title">短语</div>([\s\S]*?)<div[^>]*>`
const webItemsURL = ``
const webItemsRecommend = ``
const webItemsTranslation = ``

const webItem = `<span class="contentTitle"><a class="search-js" href="([^>]*?)">([^<]*?)</a></span>([\s\S]*?)</p>`

// ParseWebTrans parse web translation
func ParseWebTrans(content []byte) []WebTrans {
	re := regexp.MustCompile(webTrans)
	matchers := re.FindAllSubmatch(content, 1)

	re = regexp.MustCompile(webItem)
	match := re.FindAllSubmatch(matchers[0][1], -1)
	webTrans := []WebTrans{}
	for i := 0; i < len(match); i++ {
		webTrans = append(webTrans, WebTrans{
			URL:         "https://www.youdao.com" + string(match[i][1]),
			Translation: RemoveSpace(string(match[i][3])),
			Recommend:   string(match[i][2]),
		})
	}
	return webTrans
}

const webItemgray = `<span class=gray>[^<]+?</span>`

// RemoveSpace remove the space etc
func RemoveSpace(trans string) string {
	re := regexp.MustCompile(webItemgray)
	s := re.FindString(trans)
	//fmt.Println(len(s))
	var result string
	result = strings.Replace(trans, s, "", -1)
	result = strings.Replace(result, " ", "", -1)
	result = strings.Replace(result, "\n", "", -1)
	return result
}

const errorinput = `<div class="error-typo">`

// BadInput input the bad word
func BadInput(content []byte) error {
	re := regexp.MustCompile(errorinput)
	matchers := re.Find(content)
	if matchers != nil {
		//fmt.Println("Bad input")
		return errors.New("BadInput")
	}
	return nil
}

// ParseProTrans parse professional translation
// ParseEETrans parse English explanation
