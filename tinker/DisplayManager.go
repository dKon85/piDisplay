package tinker

import (
	"fmt"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
	"github.com/Tinkerforge/go-api-bindings/lcd_20x4_bricklet"
)

var lcd lcd_20x4_bricklet.LCD20x4Bricklet
var ipcon ipconnection.IPConnection

//Ringbuffer to hold all lines to be displayed, followed by a pointer showing the first line to print
var lineBuffer = []string{"", "", "", ""}
var bufferIndex = 0

// Open a connection to the display and register listeners
func InitDisplay(addr string, uid string) {

	ipcon = ipconnection.New()

	lcd, _ = lcd_20x4_bricklet.New(uid, &ipcon) // Create device object.

	ipcon.Connect(addr) // Connect to brickd.

	ActivateDisplay()

	lcd.RegisterButtonPressedCallback(func(button uint8) {
		fmt.Printf("Button Pressed: %d\n", button)
	})

	lcd.RegisterButtonReleasedCallback(func(button uint8) {
		fmt.Printf("Button Released: %d\n", button)
	})
}

// Adds a new line to the lineBuffer and writes the buffer to the display.
func AppendText(text string) {
	lineBuffer[bufferIndex] = text
	bufferIndex = (bufferIndex + 1) % 4
	lcd.ClearDisplay()
	writeLines()
}

// Writes lineBuffer to the display, erasing everything written before.
func writeLines() {
	writeLine(lineBuffer[bufferIndex], 0)
	writeLine(lineBuffer[(bufferIndex+1)%4], 1)
	writeLine(lineBuffer[(bufferIndex+2)%4], 2)
	writeLine(lineBuffer[(bufferIndex+3)%4], 3)
}

// Writes a single line to the display while cutting it after 20 chars
func writeLine(text string, line int) {
	if len(text) > 20 {
		runes := []rune(text)
		text = string(runes[0:20])
	}
	lcd.WriteLine(uint8(line), 0, text)
}

func ActivateDisplay() {
	lcd.ClearDisplay()
	lcd.BacklightOn()
}

func DeactivateDisplay() {
	lcd.ClearDisplay()
	lcd.BacklightOff()
}

func DisconnectDisplayManager() {
	ipcon.Disconnect()
	ipcon.Close()
}
