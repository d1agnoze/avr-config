package misc

import (
	"fmt"
	"os"
)

type config struct {
	exe string
	lib string
}

func NewConfig(exe string, lib string) config {
	return config{exe: exe, lib: lib}
}
func Gen_Makefile(tem Template) error {
	declear := fmt.Sprintf(`
	MCU = %v
	FLAG = %v
	FILE = %v
	# OTHER_FILES = abc
	LINK_OPT = %v
	PROGRAMER = %v
	PORT_NAME = %v
	`, tem.MCU, tem.FLAG, tem.FILE, tem.LINK_OPT, tem.PROGRAMER, tem.PORT_NAME)
	declear += CMD_TEMPLATE
	err := createFile("Makefile", declear)
	return err
}
func createFile(path string, content string) error {
	//make .vscode config
	if _, err := os.Stat("./.vscode"); os.IsNotExist(err) {
		// path/to/whatever does not exist
		err := os.MkdirAll(".vscode", 0777)
		if err != nil {
			return err
		}
	}
	// err = ioutil.WriteFile(path, []byte(content), 0644)
	// if err != nil {
	// 	return err
	// }
	err := os.WriteFile(path, []byte(content), 0666)
	return err
}
func Gen_CONFIG(conf config) error {
	workspace := fmt.Sprintf(`
	{
		"configurations": [
			{
				"name": "Win32",
				"includePath": [
					"${workspaceFolder}/**",
					"%v"
				],
				"defines": [
					"_DEBUG",
					"UNICODE",
					"_UNICODE"
				],
				"compilerPath": "%v",
				"cStandard": "c11",
				"cppStandard": "gnu++14",
				"intelliSenseMode": "windows-gcc-x86"
			}
		],
		"version": 4
	}
	`, conf.lib, conf.exe)
	err := createFile(".vscode/c_cpp_properties.json", workspace)
	return err
}
