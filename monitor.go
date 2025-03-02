package goTrinamic

import (
	"time"
)

// monitorAllShadowRegisters polls all drivers periodically and updates latestRegDisplay.
func MonitorShadowRegs(drivers []*Driver) string {
	for {
		var displayStr string
		for _, d := range drivers {
			// Create shadow instances for registers.
			gconfReg := NewGCONF()
			rawGCONF, err := d.ReadRegister(GCONF)
			if err != nil {
				displayStr += "Driver " + string(d.Address) + " error reading GCONF: " + err.Error() + "\n"
				continue
			}
			gconfReg.Unpack(rawGCONF)

			gstatReg := NewGSTAT()
			rawGSTAT, err := d.ReadRegister(GSTAT)
			if err != nil {
				displayStr += "Driver " + string(d.Address) + " error reading GSTAT: " + err.Error() + "\n"
				continue
			}
			gstatReg.Unpack(rawGSTAT)

			ioinReg := NewIOIN()
			rawIOIN, err := d.ReadRegister(IOIN)
			if err != nil {
				displayStr += "Driver " + string(d.Address) + " error reading IOIN: " + err.Error() + "\n"
				continue
			}
			ioinReg.Unpack(rawIOIN)

			// Build a display string for this driver.
			displayStr += "Driver " + string(d.Address) + ":\n" +
				"  GCONF: Recalibrate=" + BoolToStr(gconfReg.Recalibrate) +
				" Faststandstill=" + BoolToStr(gconfReg.Faststandstill) +
				" EnPwmMode=" + BoolToStr(gconfReg.EnPwmMode) + "\n" +
				"  GSTAT: Reset=" + BoolToStr(gstatReg.Reset) +
				" DrvErr=" + BoolToStr(gstatReg.DrvErr) +
				" UvCp=" + BoolToStr(gstatReg.UvCp) + "\n" +
				"  IOIN: Version=" + ToHex8(ioinReg.Version) +
				" DrvEnn=" + BoolToStr(ioinReg.DrvEnn) + "\n\n"
		}

		// Update every 5 seconds.
		time.Sleep(5 * time.Second)
		return displayStr
	}

}
