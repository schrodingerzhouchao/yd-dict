package fetcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Fetch fetch the url and return the text of url body.
/*
func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		//return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
		if resp.StatusCode == http.StatusForbidden {
			client := &http.Client{}
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				log.Fatalln(err)
			}
			req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linu…) Gecko/20100101 Firefox/65.0")
			resp, err = client.Do(req)
			if err != nil {
				log.Fatalln(err)
			}
		} else {
			return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)

		}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, nil
}
*/
func Fetch(url []byte) ([]byte, error) {
	resp, err := http.Get(string(url))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		//return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
		if resp.StatusCode == http.StatusForbidden {
			client := &http.Client{}
			req, err := http.NewRequest("GET", string(url), nil)
			if err != nil {
				log.Fatalln(err)
			}
			req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linu…) Gecko/20100101 Firefox/65.0")
			resp, err = client.Do(req)
			if err != nil {
				log.Fatalln(err)
			}
		} else {
			return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)

		}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, nil
}
