package logging

import (
	"log"
	"os"
)

var (
    LogInfo  *log.Logger
    LogErr *log.Logger
)

// create 2 loggers, one for outputting important information, the other for errors 
func InitLogger() {
	flags := log.Ldate | log.Ltime | log.Lshortfile
	LogInfo = log.New(os.Stdout, "LOG --> [INFO]", flags)
	LogErr = log.New(os.Stderr, "LOG --> [ERROR]", flags)
}