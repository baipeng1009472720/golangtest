package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"
)

var ch = make(chan string)

func main() {

	for i := 0; i < 10; i++ {
		go get(i*10, 10)
		time.Sleep(2 * time.Second)

	}

	go func() {
		for {
			select {
			case val := <-ch:
				fmt.Println(val)
			}
		}
	}()
	//hah

	time.Sleep(20 * time.Second)

}

func get(offset int, size int) map[string]interface{} {

	url := "https://mp.weixin.qq.com/mp/profile_ext?action=getmsg&__biz=MzA3ODUwMDI5OA==&f=json&offset=%d&count=%d&is_ok=1&scene=&uin=777&key=777&pass_ticket=YcwvdYzzp%2FQQQoMVoHJPDjMfLAhg73h2fB7Qb2W9AIPX3wFKPfubYjcGi7ykowad&wxtoken=&appmsg_token=933_D9O8t91hCg49FrJhgXA0rvHIdlqLyZTDbLehtw~~&x5=0&f=json"

	url = fmt.Sprintf(url, offset, size)
	fmt.Println(url)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Host", "mp.weixin.qq.com")
	req.Header.Add("Cookie", "pass_ticket=YcwvdYzzp/QQQoMVoHJPDjMfLAhg73h2fB7Qb2W9AIPX3wFKPfubYjcGi7ykowad; wap_sid2=CJL94rAGElxXUWl6b0tBRElNdzNqdXFSYWIyVVpZNjF0NTg1cHJiYW5kUmNEYkN2RkxjUm1MXzV0V2ZYblp0VU1EZ2RvLXZrOTJVQXRQSXhtUVoyM2hDMFROdE5RNlVEQUFBfjDFyI/RBTgMQJRO; wxuin=1712897682; wxtokenkey=98709c1ba8cf8072774b08b44dc2e0864f3a8bd687ba7dfc5dba60ff465b96a7")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	var dat map[string]interface{}
	json.Unmarshal(body, &dat)
	str := string(body)
	//fmt.Println(str)
	ch <- str
	//fmt.Println(dat)
	//fmt.Println(len(dat))
	return dat
}
