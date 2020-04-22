package tempconv

import (
	"fmt"
	"testing"
)

// typeSwitch Test the type of a variable
func typeSwitch(value interface{}, name string) bool {
	switch value.(type) {
	case Celsius:
		return name == "Celsius"
	case Fahrenheit:
		return name == "Fahrenheit"
	default:
		return false
	}
}

// TestConversionTypes tests the convertion of temperature float type to Celsius and Fahrenheit
func TestConversionTypes(t *testing.T) {
	var temperature = 35
	if tempCelsius := Celsius(temperature); !typeSwitch(tempCelsius, "Celsius") {
		msg := fmt.Sprintf("Type mismatched: Celsius and %v\n", tempCelsius)
		t.Fatal(msg)
	}
	if tempFahrenheit := Fahrenheit(temperature); !typeSwitch(tempFahrenheit, "Fahrenheit") {
		msg := fmt.Sprintf("Type mismatched: Fahrenheit and %v", tempFahrenheit)
		t.Fatal(msg)
	}
}

// TestCelsiusToFahrenheit test the conversion of Celsius temperature to Fahrenheit
func TestCelsiusToFahrenheit(t *testing.T) {
	if temp := CToF(Celsius(FreezingC)); temp.String() != "32.00°F" {
		t.Fatal(temp)
	}
	if temp := CToF(Celsius(AbsoluteZeroC)); temp.String() != "-459.67°F" {
		t.Fatal(temp)
	}
	if temp := CToF(Celsius(BoilingC)); temp.String() != "212.00°F" {
		t.Fatal(temp)
	}
}

func TestFahrenheitToCelsis(t *testing.T) {
	if temp := FToC(Fahrenheit(FreezingC)); temp.String() != "-17.00°C" {
		t.Fatal(temp)
	}
	if temp := FToC(Fahrenheit(AbsoluteZeroC)); temp.String() != "-290.15°C" {
		t.Fatal(temp)
	}
	if temp := FToC(Fahrenheit(BoilingC)); temp.String() != "83.00°C" {
		t.Fatal(temp)
	}
}
