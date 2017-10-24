package tempconv

//Convert temperature from Celsium to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

//Convert temperature from Fahrenheit to Celsium
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
