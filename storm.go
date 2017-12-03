package main

import (
	"github.com/asdine/storm"
	"github.com/astaxie/beego/logs"
)

type Job struct {
	JobId  string `storm:"id"`
	Status string `storm:"index"`
}

func main() {

	db, err := storm.Open("my.db")
	if err != nil {
		logs.Info(err)
	}
	defer db.Close()

	job := Job{"111", "ok"}
	err2 := db.Save(&job)
	if err2 != nil {
		logs.Error(err2)
	}
	err3 := db.Update(&Job{"111", "ok1"})
	if err3 != nil {
		logs.Error(err3)
	}


	var jobs []Job
	err22 := db.Find("Status", "ok1", &jobs)
	if err22 != nil {
		logs.Error(err22)
	}
	logs.Info(jobs)

}
