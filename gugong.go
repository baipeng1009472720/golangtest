package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "http://mobileapps.dpm.org.cn/AppInterfaces/Content.aspx?from=2017-11-01&m=get_date_content&to=2017-11-29"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "6e4b3d72-843f-e0da-5dd7-1fe479de1766")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}