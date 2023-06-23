package util

import (
	"os"
	"os/exec"
	"reflect"
)

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func IsStructEmpty(s interface{}) bool {
	v := reflect.ValueOf(s)
	return v.Kind() == reflect.Struct
}