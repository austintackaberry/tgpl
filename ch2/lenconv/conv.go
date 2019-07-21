package lenconv

// CToF converts a Celsius temperature to Fahrenheit.
func MToFt(m Meters) Feet { return Feet(m / 0.3048) }

// FToC converts a Fahrenheit temperature to Celsius.
func FtToM(ft Feet) Meters { return Meters(ft * 0.3048) }
