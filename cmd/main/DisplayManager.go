package main

import(
	"fmt"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
	"github.com/Tinkerforge/go-api-bindings/lcd_20x4_bricklet"
)

var lcd lcd_20x4_bricklet.LCD20x4Bricklet
var ipcon ipconnection.IPConnection

func InitDisplay(){

	ipcon = ipconnection.New()

	lcd, _ = lcd_20x4_bricklet.New(UID, &ipcon) // Create device object.

	ipcon.Connect(ADDR) // Connect to brickd.
	// Don't use device before ipcon is connected.


	lcd.BacklightOn()
	lcd.RegisterButtonPressedCallback(func(button uint8) {
		fmt.Printf("Button Pressed: %d\n", button)
	})

	lcd.RegisterButtonReleasedCallback(func(button uint8) {
		fmt.Printf("Button Released: %d\n", button)
	})
}

func WriteLine( text string, line int ){
	if len(text) > 20 {
		runes := []rune(text)
		text = string(runes[0:19])
	}
	lcd.WriteLine(uint8(line), 0, text)
}

func ActivateDisplay(){
	lcd.ClearDisplay()
	lcd.BacklightOn()
}

func DeactivateDisplay(){
	lcd.ClearDisplay()
	lcd.BacklightOff()
}

func DisconnectDisplayManager(){
	ipcon.Disconnect()
	ipcon.Close()
}
