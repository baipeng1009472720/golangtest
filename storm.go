package main

import (
	"github.com/asdine/storm"
	"github.com/astaxie/beego/logs"
)

type Job struct {
	JobId  string `storm:"id"`
	status string `storm:"index"`
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

	var job3 Job
	err3 := db.One("JobId", "111", &job3)
	if err2 != nil {
		logs.Error(err3)
	}

	logs.Info(job3)

}
