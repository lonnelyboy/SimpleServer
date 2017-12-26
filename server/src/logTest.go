package main

import (
	log "github.com/cihub/seelog"
)

func test1(){
	//Output log
	defer log.Flush()
	log.Info("Hello from Seelog")
}

func test2(){
	defer log.Flush()
	//Loading config
	logPath := "D:\\Project\\GGbound\\bin\\seelog.xml"
	logger, err := log.LoggerFromConfigAsFile(logPath)
	if err != nil{
		log.Error("test2.load config failed")
		return
	}
	log.ReplaceLogger(logger)
}

func exceptLog()  {
	defer log.Flush()
	log.Trace("Trace programme info")
	log.Debug("Debug programme info")
}

func test3() {
	defer log.Flush()
}

func main() {
	//test1()
	test2()
	exceptLog()
	defer log.Flush()
	log.Debug("sssssssssssss")
	log.Info("222222222222222222")
}