package logger

import (
	log "github.com/cihub/seelog"
	"strings"
	"fmt"
)

var file string

type LoggingConfig struct {
	logLevel string `json:"log_level"`
	logFile  string `config:"logs"`
}

func SetLogging(isFlagDebug bool, logLevel string, logFile string) {
	loggingConfig := &LoggingConfig{}

	if isFlagDebug {
		loggingConfig.logLevel = "debug"

	}
	if len(logLevel) > 0 {
		loggingConfig.logLevel = strings.ToLower(logLevel)
	}
	if len(logFile) > 0 {
		file = logFile
	}
	if file == "" {
		file = "./freeci.log"
	}
	if loggingConfig.logLevel == "" {
		loggingConfig.logLevel = "info"
	}

	//
	//consoleWriter, _ := NewConsoleWriter()
	//format := "[%Date(01-02) %Time] [%LEV] [%File:%Line] %Msg%n"
	//formatter, err := log.NewFormatter(format)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//rollingFileWriter, _ := NewRollingFileWriterSize(file, rollingArchiveNone, "", 10000000000, 5, rollingNameModePostfix)
	//bufferedWriter, _ := NewBufferedWriter(rollingFileWriter, 10000, 1000)
	////logging receivers
	//receivers := []interface{}{consoleWriter, bufferedWriter}
	//
	//root, err := log.NewSplitDispatcher(formatter, receivers)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//l, _ := log.LogLevelFromString(strings.ToLower(loggingConfig.logLevel))
	//
	//golbalConstraints, err := log.NewMinMaxConstraints(l, log.CriticalLvl)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//exceptions := []*log.LogLevelException{}
	//logger := log.NewAsyncLoopLogger(log.NewLoggerConfig(golbalConstraints, exceptions, root))

	logger, err := log.LoggerFromConfigAsFile("seelog.xml")
	if err != nil {
		panic(err)
	}
	//log.ReplaceLogger(logger)
	err = log.ReplaceLogger(logger)
	if err != nil {
		fmt.Println(err)
	}
}

func Flush() {
	log.Flush()
}
