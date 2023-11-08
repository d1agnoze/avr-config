package misc

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func Scan() (map[string]string, error) {
	res := make(map[string]string, 4)
	var cmd, sep string
	if runtime.GOOS == "windows" {
		cmd = "where"
		sep = "\\"
	} else {
		cmd = "which"
		sep = "/"
	}
	if out, err := exec.Command(cmd, "avr-gcc").Output(); err == nil && strings.Trim(string(out), " ") != "" {
		res["avr-gcc"] = string(out)
		assumePath := CutFilePathFromTail(CutFilePathFromTail(string(out), sep), sep)
		if _, err := os.Stat(assumePath); !os.IsNotExist(err) {
			res["include"] = assumePath + "\\avr\\include"
		} else {
			res["include"] = `program did not find AVR include directory😢, please specify it in vscode config`
			return res, nil
		}
	} else {
		return nil, err
	}
	return res, nil
}
func CutFilePathFromTail(path string, sep string) string {
	path = strings.ReplaceAll(path, "\r", "")
	path = strings.ReplaceAll(path, "\n", "")
	index := strings.LastIndex(path, sep)
	if index == -1 {
		return path
	}
	return path[:index]
}
