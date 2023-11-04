17

## source: [](https://stackoverflow.com/questions/32413959/avr-gcc-with-arduino)

I recommend the following set of command line options for compiling:

```
avr-gcc -c
        -std=gnu99
        -Os
        -Wall
        -ffunction-sections -fdata-sections
        -mmcu=m328p
        -DF_CPU=16000000

```

And for linking:

```
avr-gcc -Os
        -mmcu=m328p
        -ffunction-sections -fdata-sections
        -Wl,--gc-sections

```

Where...

-   `-c` means "compile to object file only, do not link"
-   `-std=gnu99` means "My code conforms to C99 and I use GNU extensions"
-   `-Os` means "optimize for executable size rather than code speed"
-   `-Wall` means "turn on (almost) all warnings"
-   `-ffunction-sections -fdata-sections` is necessary for the `-Wl,--gc-sections` optimization
-   `-mmcu=m328p` means "the MCU part number is ATmega 328P"
-   `-DF_CPU=16000000` means "the clock frequency is 16 MHz" (adjust for your actual clock frequency)
-   `-Wl,--gc-sections` means "tell the linker to drop unused function and data sections" (this helps reduce code size).

In order to actually compile your code, you would first issue the `avr-gcc` command with the "compile only flags", like this:

```
avr-gcc -c -std=gnu99 <etc.> MyProgram.c -o MyProgram.o

```

Then you would repeat this for all of your source files. Finally, you would link the resulting object files together by invoking AVR-GCC in link mode:

```
avr-gcc -Os <etc.> MyProgram.o SomeUtility.o -o TheExecutable.elf

```

This generates an ELF file, which isn't directly executable by your MCU. Thus, you'll need to extract the useful part (the raw machine code) from it in the Intel Hex format:

```
avr-objcopy -O ihex -R .eeprom TheExecutable.elf TheExecutable.ihex

```

Finally, you will need AVRdude to upload the contents of the hex file to the MCU:

```
avrdude -C /path/to/avrdude.conf
        -p m328p
        -c PROGRAMMER_NAME
        -b 19600
        -P PORT_NAME
        -U flash:w:TheExecutable.ihex:i

```

Where...

-   `-C /path/to/avrdude.conf` means "use this file as the configuration file"
-   `-c PROGRAMMER_NAME` means "I am using a programmer of type PROGRAMMER_NAME" (you will need to fill this in yourself depending on what kind of programmer you use).
-   `-b 19600` is the baud rate (you may need to adjust this depending on the baud rate you set or have pre-programmed into the bootloader)
-   `-P PORT_NAME` means "the programmer is connected to port PORT_NAME". On Linux, it will most often be something like `/dev/ttyusbN`, where N is some number.
-   `-U flash:w:TheExecutable.ihex:i` means "write to the Flash memory the contents of TheExecutable.ihex which is in Intel Hex format".