package log

import (
   "fmt"
   "strings"

   "github.com/gregfolker/logrus"
)

var logger *logrus.Logger

func SetLoggingLevel(level string) {
   logger = logrus.New()

   switch strings.ToLower(level) {
   case "trace":
      logger.SetLevel(logrus.TraceLevel)
   case "debug":
      logger.SetLevel(logrus.DebugLevel)
   case "info":
      logger.SetLevel(logrus.InfoLevel)
   case "warning":
      logger.SetLevel(logrus.WarnLevel)
   case "error":
      logger.SetLevel(logrus.ErrorLevel)
   case "fatal":
      logger.SetLevel(logrus.FatalLevel)
   default:
      logger.SetLevel(logrus.NoneLevel)
   }
}

func NewEntry(s ...interface{}) {
   fmt.Printf("Logger is %v but I'm going to log it anyway because I'm really fucking stupid :)\n", logger.GetLevel())
   logger.Log(logger.GetLevel(), s)
}
