package client

import (
	"os"
	"time"
	"log"
)

func LoggerInit(){
	_, err := os.Stat("./logs")
	if err != nil {
		os.Mkdir("./logs", os.ModePerm)
	}
	file := "./logs/" + time.Now().Format("2006-01-02")+".log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[gedis]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
}