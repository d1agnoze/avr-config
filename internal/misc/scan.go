package misc

import (
	"log"
	"os/exec"
	"runtime"
	"strings"
	"sync"
)

func Scan(wait *sync.WaitGroup) map[string]string {
	defer wait.Done()
	res := make(map[string]string, 4)
	var cmd, cmd2, sep, assumePath string
	if runtime.GOOS == "windows" {
		cmd = "where"
	} else {
		cmd = "which"
	}
	if out, err := exec.Command(cmd, "avr-gcc").Output(); err != nil && strings.Trim(string(out), " ") != "" {
		res["avr-gcc"] = string(out)
	} else {
		log.Println("ERRRO", err)
	}
	if out, err := exec.Command(cmd, "avrdude").Output(); err != nil && strings.Trim(string(out), " ") != "" {
		res["avrdude"] = string(out)
		if runtime.GOOS == "windows" {
			cmd2 = "if"
			sep = "\\"
			assumePath = res["avrdude"] + sep + "avrdude.conf"
			if out, err := exec.Command(cmd2, "EXIST", assumePath+" echo true").Output(); err != nil && strings.Trim(string(out), " ") != "true" {
				res["conf"] = string(out)
			} else {
				log.Println("ERRRO", err)
			}
		} else {
			cmd2 = "/bin/sh"
			sep = "/"
			assumePath = res["avrdude"] + sep + "avrdude.conf"
			if out, err := exec.Command(cmd2, "-c", "[ -f "+assumePath+" ] && echo true").Output(); err != nil && strings.Trim(string(out), " ") != "true" {
				res["conf"] = string(out)
			} else {
				log.Println("ERRRO", err)
			}
		}
	} else {
		log.Println("ERRRO", err)
	}
	log.Println(res)
	return res
}
func cutFilePathFromTail(path string, sep string) string {
	index := strings.LastIndex(path, sep)
	if index == -1 {
		return path
	}
	return path[:index]
}
