package main

import (
	log "github.com/cihub/seelog"
	ulog "./logger"
	"flag"
	"fmt"
	"strings"
)

func main() {

	//logger.SetLogging(nil, "info", "123")
	var logLevel = flag.String("log", "info", "the log level,options:trace,debug,info,warn,error")
	fmt.Println(strings.ToLower(*logLevel))

	ulog.SetLogging(false, *logLevel, "")
	defer func() {
		log.Flush()
		ulog.Flush()
	}()

	//logger, err := log.LoggerFromConfigAsFile("./seelog.xml")

	//if err != nil {
	//	log.Error(err)
	//}

	//log.ReplaceLogger(logger)
	for {
		log.Info("hello form seelog Info")
		log.Debug("hello form seelog Debug")
		log.Warn("hello form seelog Warn")
		log.Error("hello form seelog Error")
		log.Infof("")
		log.Debugf("")
		log.Warnf("")
		log.Errorf("")
	}

	log.Flush()
	ulog.Flush()
}
