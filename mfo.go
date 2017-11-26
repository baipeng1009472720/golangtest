package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://m.weibo.cn/api/container/getIndex?containerid=1076032415848337&page=1"

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("type", "2")
	req.Header.Add("adcode", "3100")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "c8fe7741-4768-6d51-8a51-039791f87b47")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
