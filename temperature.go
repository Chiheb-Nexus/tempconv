package tempconv

import "fmt"

// Celsius °C
type Celsius float64

// Fahrenheit °F
type Fahrenheit float64

// Constants for Celsius and Fahrenheit
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%.2f°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.2f°F", f)
}

// CToF Convert Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC Convert Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius(f - 32*5/9)
}
