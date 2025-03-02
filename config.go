package main

//
//// SPIConfig holds the SPI settings.
//var BUS = machine.SPIConfig{
//	LSBFirst:  false,
//	SDI:       machine.GPIO0,
//	SDO:       machine.GPIO3,
//	SCK:       machine.GPIO2,
//	Mode:      3,
//	Frequency: 4000000,
//}
//
//// DriverConfig holds settings for each driver.
//type DriverConfig struct {
//	Address      uint8
//	CSPin        machine.Pin
//	EnablePin    machine.Pin
//	DrvStrength  uint8
//	BbmTime      uint8
//	BbmClks      uint8
//	GlobalScaler uint16
//	Ihold        uint8
//	Irun         uint8
//	PwmGradInit  uint16
//	PwmOfsInit   uint16
//	Freewheeling uint8
//
//	StepAngle float32
//	GearRatio float32
//	VSupply   float32
//	RCoil     float32
//	LCoil     float32
//	IPeak     float32
//	RSense    float32
//	MSteps    uint8
//	Fclk      uint8
//}
//
//// Predefined configuration for your drivers.
//// You can add as many as you need.
//var Drivers = []DriverConfig{
//	{
//		Address:      0,
//		CSPin:        machine.Pin(1),
//		EnablePin:    machine.Pin(11),
//		DrvStrength:  2,
//		BbmTime:      5,
//		BbmClks:      4,
//		GlobalScaler: 128,
//		Ihold:        16,
//		Irun:         31,
//		PwmGradInit:  12,
//		PwmOfsInit:   196,
//		Freewheeling: 3,
//		StepAngle:    1.8,
//		GearRatio:    1.0,
//		VSupply:      24.0,
//		RCoil:        1.2,
//		LCoil:        0.005,
//		IPeak:        2.0,
//		RSense:       0.1,
//		MSteps:       16,
//		Fclk:         12,
//	},
//	//{
//	//	Address:      1,
//	//	CSPin:        machine.Pin(12),
//	//	EnablePin:    machine.Pin(13),
//	//	DrvStrength:  3,
//	//	BbmTime:      8,
//	//	BbmClks:      5,
//	//	GlobalScaler: 256,
//	//	Ihold:        10,
//	//	Irun:         25,
//	//	PwmGradInit:  12,
//	//	PwmOfsInit:   200,
//	//	Freewheeling: 0,
//	//	StepAngle:    0.9,
//	//	GearRatio:    2.0,
//	//	VSupply:      24.0,
//	//	RCoil:        0.6,
//	//	LCoil:        0.003,
//	//	IPeak:        3.0,
//	//	RSense:       0.05,
//	//	MSteps:       32,
//	//	Fclk:         12,
//	//},
//}
