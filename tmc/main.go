package tmc

import (
	"github.com/amken3d/goTrinamic"
	"machine"
	"time"
)

// Global variable to hold the latest register information.
var latestRegDisplay string

// boolToStr converts a bool to "true" or "false"
func boolToStr(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

// monitorAllShadowRegisters polls all drivers periodically and updates latestRegDisplay.
func monitorAllShadowRegisters(drivers []*Driver) {
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
				"  GCONF: Recalibrate=" + boolToStr(gconfReg.Recalibrate) +
				" Faststandstill=" + boolToStr(gconfReg.Faststandstill) +
				" EnPwmMode=" + boolToStr(gconfReg.EnPwmMode) + "\n" +
				"  GSTAT: Reset=" + boolToStr(gstatReg.Reset) +
				" DrvErr=" + boolToStr(gstatReg.DrvErr) +
				" UvCp=" + boolToStr(gstatReg.UvCp) + "\n" +
				"  IOIN: Version=" + goTrinamic.ToHex8(ioinReg.Version) +
				" DrvEnn=" + boolToStr(ioinReg.DrvEnn) + "\n\n"
		}
		latestRegDisplay = displayStr
		// Update every 5 seconds.
		time.Sleep(5 * time.Second)
	}
}

// refreshDisplay clears the screen and redraws header and dynamic content.
func refreshDisplay(header, content string) {
	// ANSI escape codes: "\033[H" moves the cursor home and "\033[2J" clears the screen.
	print("\033[H\033[2J")
	print(header)
	print(content)
}

func main() {
	println("Starting TMC Drivers")
	time.Sleep(4 * time.Second)

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// Setup SPI using compile-time config.
	spi := machine.SPI0
	if err := spi.Configure(BUS); err != nil {
		println("Failed to configure SPI", err.Error())
		return
	}
	println("Configured SPI:", spi)

	// Build a map of chip-select pins.
	csPins := make(map[uint8]machine.Pin)
	for _, driver := range Drivers {
		csPins[driver.Address] = driver.CSPin
	}
	println("CS Pins configured:", len(csPins))

	// Create SPIComm instance.
	comm := NewSPIComm(*spi, csPins)
	if err := comm.Setup(); err != nil {
		println("Error setting up SPI:", err.Error())
		return
	}
	println("Starting TMC Drivers")

	// Create drivers from the config.
	var drivers []*Driver
	for _, drvConf := range Drivers {
		driver := NewDriver(
			comm,
			drvConf.Address,
			drvConf.EnablePin,
			NewStepper(
				drvConf.StepAngle,
				drvConf.GearRatio,
				drvConf.VSupply,
				drvConf.RCoil,
				drvConf.LCoil,
				drvConf.IPeak,
				drvConf.RSense,
				drvConf.MSteps,
				drvConf.Fclk,
			),
		)
		drivers = append(drivers, driver)
		println("Created driver at address:", driver.Address)
	}

	// Start a single goroutine that polls all drivers and updates latestRegDisplay.
	go monitorAllShadowRegisters(drivers)

	// (Optional) Run LED blinking in its own goroutine.
	go func() {
		for {
			led.High()
			time.Sleep(400 * time.Millisecond)
			led.Low()
			time.Sleep(600 * time.Millisecond)
		}
	}()

	header := "=== REPL Interface ===\nCommon Info: Driver status monitor\n----------------------\n"

	// Main display refresh loop.
	for {
		content := "Current time: " + time.Now().Format("15:04:05") + "\n\n" + latestRegDisplay
		refreshDisplay(header, content)
		time.Sleep(500 * time.Millisecond)
	}
}
