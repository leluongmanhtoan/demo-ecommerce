package log

import (
	"fmt"
	"path/filepath"
	"runtime"

	log "github.com/sirupsen/logrus"
)

func getMeta() (meta string) {
	_, path, numLine, _ := runtime.Caller(1)
	srcFile := filepath.Base(path)
	meta = fmt.Sprintf("%s:%d", srcFile, numLine)
	return
}

func Info(msg ...any) {
	log.WithFields(log.Fields{
		"meta": getMeta(),
	}).Info(msg...)
}

func Warning(msg ...any) {
	log.WithFields(log.Fields{
		"meta": getMeta(),
	}).Warning(msg...)
}

func Error(err ...any) {
	log.WithFields(log.Fields{
		"meta": getMeta(),
	}).Error(err...)
}

func Debug(value ...any) {
	log.WithFields(log.Fields{
		"meta": getMeta(),
	}).Debug(value...)
}

func Fatal(value ...any) {
	log.WithFields(log.Fields{
		"meta": getMeta(),
	}).Fatal(value...)
}

func Println(value ...any) {
	log.WithFields(log.Fields{
		"meta": getMeta(),
	}).Println(value...)
}

func Infof(format string, msg ...any) {
	log.WithFields(log.Fields{
		"meta": getMeta(),
	}).Infof(format, msg...)
}

func Warningf(format string, msg ...any) {
	log.WithFields(log.Fields{
		"meta": getMeta(),
	}).Warningf(format, msg...)
}

func Errorf(format string, err ...any) {
	log.WithFields(log.Fields{
		"meta": getMeta(),
	}).Errorf(format, err...)
}

func Debugf(format string, value ...any) {
	log.WithFields(log.Fields{
		"meta": getMeta(),
	}).Debugf(format, value...)
}

func Fatalf(format string, value ...any) {
	log.WithFields(log.Fields{
		"meta": getMeta(),
	}).Fatalf(format, value...)
}
