package misc

type Template struct {
	MCU       string //add on run - m328p if the board is arduino uno
	FLAG      string
	FILE      string
	LINK_OPT  string
	AD_CONF   string //add on run - your avrdude.conf location
	PROGRAMER string //add on run - your programmer - arduino if the board is arduino
	PORT_NAME string //add on run - your port normally is COM3
}

func NewTemplate() Template {
	res := Template{
		FLAG:     "-c -std=gnu99 -Os -Wall -ffunction-sections -fdata-sections -mmcu=$(MCU) -DF_CPU=16000000",
		FILE:     "main",
		LINK_OPT: "-Os -mmcu=$(MCU) -ffunction-sections -fdata-sections -Wl,--gc-sections",
	}
	return res
}

const CMD_TEMPLATE = `default: comp link extract upload
comp:
	avr-gcc $(FLAG) $(FILE).c -o $(FILE).o 
link:
	avr-gcc $(LINK_OPT) $(FILE).o -o $(FILE).elf
extract:
	avr-objcopy -O ihex -R .eeprom $(FILE).elf $(FILE).ihex
upload:
	avrdude -p $(MCU) -c $(PROGRAMER) -b 19600 -P $(PORT_NAME) -U flash:w:$(FILE).ihex:i`
