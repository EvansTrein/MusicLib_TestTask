package logging

import (
	"log"
	"os"
)

var (
    LogInfo  *log.Logger
    LogErr *log.Logger
)

func InitLogger() {
	flags := log.Ldate | log.Ltime | log.Lshortfile
	LogInfo = log.New(os.Stdout, "LOG --> [INFO]", flags)
	LogErr = log.New(os.Stdout, "LOG --> [ERROR]", flags)

}