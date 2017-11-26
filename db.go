package main

import (
	"github.com/boltdb/bolt"
	"github.com/astaxie/beego/logs"
	"path"
	"fmt"
)

func main() {
	db, err := bolt.Open("freeci.db", 0600, nil)
	if err != nil {
		logs.Error(err)
	}
	//读写
	db.Update(func(tx *bolt.Tx) error {
		return nil
	})
	//只读
	db.View(func(tx *bolt.Tx) error {
		return nil
	})
	defer db.Close()

	folder := path.Join("1212", "blob")
	fmt.Println(folder)

}
