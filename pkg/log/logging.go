package log

import (
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

func Trace(s ...interface{}) {
   logger.Trace(s)
}

func Debug(s ...interface{}) {
   logger.Debug(s)
}

func Info(s ...interface{}) {
   logger.Info(s)
}

func Warn(s ...interface{}) {
   logger.Warn(s)
}

func Error(s ...interface{}) {
   logger.Error(s)
}

func Fatal(s ...interface{}) {
   logger.Fatal(s)
}
