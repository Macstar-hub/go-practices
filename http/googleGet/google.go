package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type bodyPost struct {
	assetGeram int
	newCoin    int
	oldCoin    int
	semiCoin   int
}

func main() {
	// url := "https://google.com"
	// httpGet(url)
	httpPost()

}

func httpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Cannot get google.com", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Cannot get page with status code: ", http.StatusOK)
		fmt.Println(respConverter(resp))
	}
}

func httpPost() {

	// sampleAsset := bodyPost{
	// 	assetGeram: 472,
	// 	newCoin:    14,
	// 	oldCoin:    1,
	// 	semiCoin:   6,
	// }

	url := "http://185.81.97.192/api/v1/assetcalc"
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader("assetGeram=472&newCoin=14&oldCoin=1&semiCoin=6"))
	if err != nil {
		log.Println("Cannot make post request with error: ", err)
	}

	fmt.Println(respConverter(resp))
}

func respConverter(resp *http.Response) string {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Cannot make convert body: ", err)
	}

	return string(body)
}
