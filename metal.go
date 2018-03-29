package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://www.metal-archives.com/browse/ajax-letter/l/A/json/1?sEcho=1&iColumns=4&sColumns=&iDisplayStart=1000&iDisplayLength=2000&mDataProp_0=0&mDataProp_1=1&mDataProp_2=2&mDataProp_3=3&iSortCol_0=0&sSortDir_0=asc&iSortingCols=1&bSortable_0=true&bSortable_1=true&bSortable_2=true&bSortable_3=false&_=1513584461617"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "eb9fefe5-c072-68c4-8805-f9a6a50908fb")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}