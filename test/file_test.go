package test_test

import (
	"avr-config/cmd/avr-conf/internal/misc"
	"fmt"
	"os"
	"testing"
)

func TestFormatter(t *testing.T) {
	path := "avr\\include"
	path = misc.WindowsFileFormat(path)
	fmt.Println(path)
	if path != "avr\\\\include" {
		t.Errorf("sum thing is wrong %v", path)
	}
}
func TestFileCheck(t *testing.T) {
	res := ""
	pwd, err := os.Getwd()
	pwd = "C:/Users/Thuan/Documents/avr-test"
	var conf, mf string
	if err != nil {
		mf = pwd + "/Makefile"
		conf = pwd + "/.vscode/c_cpp_properties.json"
		if misc.FileExists(mf) {
			res += mf + "EXISTS ✔\n"
		}
		if misc.FileExists(conf) {
			res += conf + "EXISTS ✔\n"
		}
	}
	if res == "" {
		t.Error("sum thin is wrong")
	}
}
