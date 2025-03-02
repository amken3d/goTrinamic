package goTrinamic

// CommonRegisters holds the intersection of registers that
// appear on TMC5240, TMC2240, and TMC5160.
const (
	GCONF         = 0x00
	GSTAT         = 0x01
	IFCNT         = 0x02
	SLAVECONF     = 0x03
	IOIN          = 0x04
	DRV_CONF      = 0x0A
	GLOBAL_SCALER = 0x0B
	IHOLD_IRUN    = 0x10
	TPOWERDOWN    = 0x11
	TSTEP         = 0x12
	TPWMTHRS      = 0x13
	TCOOLTHRS     = 0x14
	THIGH         = 0x15
	ENCMODE       = 0x38
	XENC          = 0x39
	ENC_CONST     = 0x3A
	ENC_STATUS    = 0x3B
	ENC_LATCH     = 0x3C
	MSLUT0        = 0x60
	MSLUT1        = 0x61
	MSLUT2        = 0x62
	MSLUT3        = 0x63
	MSLUT4        = 0x64
	MSLUT5        = 0x65
	MSLUT6        = 0x66
	MSLUT7        = 0x67
	MSLUTSEL      = 0x68
	MSLUTSTART    = 0x69
	MSCNT         = 0x6A
	MSCURACT      = 0x6B
	CHOPCONF      = 0x6C
	COOLCONF      = 0x6D
	DRVSTATUS     = 0x6F
	PWMCONF       = 0x70
	PWMSCALE      = 0x71
	PWM_AUTO      = 0x72
	SG4_RESULT    = 0x75
	SG4_IND       = 0x76
)

// Registers specific to TMC5160
const (
	OTP_READ     = 0x07
	FACTORY_CONF = 0x08
	SHORT_CONF   = 0x09
	OFFSET_READ  = 0x0C
	LOST_STEPS   = 0x73
	RAMPMODE     = 0x20
)

//Registers specific to TMC5240

const (
	TVMAX           = 0x29
	XTARGET         = 0x2D
	V2              = 0x2E
	A2              = 0x2F
	D2              = 0x30
	AACTUAL         = 0x31
	VIRTUAL_STOP_L  = 0x3E
	VIRTUAL_STOP_R  = 0x3F
	ADC_VSUPPLY_AIN = 0x50
)

// Registers specific to TMC2240
const (
	DIRECT_MODE = 0x2D
)

// Registers common to TMC5160 and TMC5240
const (
	X_COMPARE     = 0x05
	OTP_PROG      = 0x06
	XACTUAL       = 0x21
	VACTUAL       = 0x22
	VSTART        = 0x23
	A1            = 0x24
	V1            = 0x25
	AMAX          = 0x26
	VMAX          = 0x27
	DMAX          = 0x28
	D1            = 0x2A
	VSTOP         = 0x2B
	TZEROWAIT     = 0x2C
	VDCMIN        = 0x33
	SWMODE        = 0x34
	RAMPSTAT      = 0x35
	XLATCH        = 0x36
	ENC_DEVIATION = 0x3D
	DCCTRL        = 0x6E
)

//Regsiters Common to TMC5240 and TMC 2240

const (
	ADC_TEMP   = 0x51
	OTW_OV_VTH = 0x52
	SG4_THRS   = 0x74
)

const (
	DEFAULT_F_CLK = 12000000
)
