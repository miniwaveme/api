package logger

import (
	"reflect"
	"testing"
)

func TestLogger(t *testing.T) {

	log := GetLogger()
	if "*logrus.Logger" != reflect.TypeOf(log).String() {
		t.Error("Test failed")
	}
}
