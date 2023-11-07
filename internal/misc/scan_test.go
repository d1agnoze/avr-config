package misc

import (
	"log"
	"testing"
)

func TestFilter(t *testing.T) {
	res := cutFilePathFromTail("C:\\Users\\Thuan\\Downloads\\avrdude-v7.2-windows-x64\\avrdude.exe", "\\")
	if res != "C:\\Users\\Thuan\\Downloads\\avrdude-v7.2-windows-x64" {
		t.Errorf("failed: %v", res)
	}
	res2 := cutFilePathFromTail("/home/vdac/avrdude/avrdude.exe", "/")
	if res2 != "/home/vdac/avrdude" {
		t.Errorf("failed: %v", res)
	}
	log.Printf("%s\n%s", res, res2)
}
