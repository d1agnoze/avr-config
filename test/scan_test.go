package misc_test

import (
	"avr-config/cmd/avr-conf/internal/misc"
	"log"
	"os"
	"testing"
)

func TestFilter(t *testing.T) {
	res := misc.CutFilePathFromTail("C:\\Users\\Thuan\\Downloads\\avrdude-v7.2-windows-x64\\avrdude.exe", "\\")
	if res != "C:\\Users\\Thuan\\Downloads\\avrdude-v7.2-windows-x64" {
		t.Errorf("failed: %v", res)
	}
	res2 := misc.CutFilePathFromTail("/home/vdac/avrdude/avrdude.exe", "/")
	log.Printf("----------\n%s\n%s\n", res, res2)
	if res2 != "/home/vdac/avrdude" {
		t.Errorf("failed: %v", res)
	}
}
func TestScan(t *testing.T) {
	var res map[string]string
	res, _ = misc.Scan()
	log.Printf("\n%v ---> %v\n", res["avr-gcc"], res["include"])
	if len(res) != 2 {
		t.Errorf("Sum thing is wrong")
	}
	if _, err := os.Stat(res["avr-gcc"]); os.IsNotExist(err) {
		// path/to/whatever does not exist
		t.Errorf("avr gcc not found")
	}
	if _, err := os.Stat(res["include"]); os.IsNotExist(err) {
		// path/to/whatever does not exist
		t.Errorf("avr include not found")
	}

}
func TestDIR(t *testing.T) {
	if _, err := os.Stat("./makefile"); os.IsNotExist(err) {
		t.Errorf("file not found")
	}
}
