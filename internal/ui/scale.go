package ui

var windowDPI = 141.2

// DPIScale returns the scale factor for interface elements based
// on screen DPI.
// Scale would normally be between 1x-2x
func DPIScale() float32 {
	return float32(windowDPI) / DefaultScreenDPI
}
