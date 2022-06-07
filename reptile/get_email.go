package reptile

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var reQQEmail = "(\\d+)@qq.com"

func GetEmail() {
	resp, err := http.Get("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	HandleError(err, "http.Get url")
	defer resp.Body.Close()

	pagesBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")

	pagesStr := string(pagesBytes)
	result := regexp.MustCompile(reQQEmail)
	results := result.FindAllStringSubmatch(pagesStr, -1)
	//fmt.Println(results, result)
	for _, re := range results {
		fmt.Println(re)
		fmt.Println("Email:", re[0])
		fmt.Println("QQ:", re[1])
	}
}

func HandleError(err error, reason string) {
	if err != nil {
		fmt.Println(reason)
	}
}
