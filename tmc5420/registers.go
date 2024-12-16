package tmc5420

// Generated Go code

type RegisterField struct {
	Name     string
	Mask     uint32
	Shift    uint32
	Register uint32
	IsSigned bool
}

// Constants
const (
	TMC5240_WRITE_BIT                         = 0x80
	TMC5240_ADDRESS_MASK                      = 0x7F
	TMC5240_MAX_VELOCITY                      = 8388096
	TMC5240_MAX_ACCELERATION                  = 8388096
	DEFAULT_ICID                              = 0
	TMC5240_MODE_POSITION                     = 0
	TMC5240_MODE_VELPOS                       = 1
	TMC5240_MODE_VELNEG                       = 2
	TMC5240_MODE_HOLD                         = 3
	TMC5240_SW_STOPL_ENABLE                   = 0x0001
	TMC5240_SW_STOPR_ENABLE                   = 0x0002
	TMC5240_SW_STOPL_POLARITY                 = 0x0004
	TMC5240_SW_STOPR_POLARITY                 = 0x0008
	TMC5240_SW_SWAP_LR                        = 0x0010
	TMC5240_SW_LATCH_L_ACT                    = 0x0020
	TMC5240_SW_LATCH_L_INACT                  = 0x0040
	TMC5240_SW_LATCH_R_ACT                    = 0x0080
	TMC5240_SW_LATCH_R_INACT                  = 0x0100
	TMC5240_SW_LATCH_ENC                      = 0x0200
	TMC5240_SW_SG_STOP                        = 0x0400
	TMC5240_SW_SOFTSTOP                       = 0x0800
	TMC5240_RS_STOPL                          = 0x0001
	TMC5240_RS_STOPR                          = 0x0002
	TMC5240_RS_LATCHL                         = 0x0004
	TMC5240_RS_LATCHR                         = 0x0008
	TMC5240_RS_EV_STOPL                       = 0x0010
	TMC5240_RS_EV_STOPR                       = 0x0020
	TMC5240_RS_EV_STOP_SG                     = 0x0040
	TMC5240_RS_EV_POSREACHED                  = 0x0080
	TMC5240_RS_VELREACHED                     = 0x0100
	TMC5240_RS_POSREACHED                     = 0x0200
	TMC5240_RS_VZERO                          = 0x0400
	TMC5240_RS_ZEROWAIT                       = 0x0800
	TMC5240_RS_SECONDMOVE                     = 0x1000
	TMC5240_RS_SG                             = 0x2000
	TMC5240_EM_DECIMAL                        = 0x0400
	TMC5240_EM_LATCH_XACT                     = 0x0200
	TMC5240_EM_CLR_XENC                       = 0x0100
	TMC5240_EM_NEG_EDGE                       = 0x0080
	TMC5240_EM_POS_EDGE                       = 0x0040
	TMC5240_EM_CLR_ONCE                       = 0x0020
	TMC5240_EM_CLR_CONT                       = 0x0010
	TMC5240_EM_IGNORE_AB                      = 0x0008
	TMC5240_EM_POL_N                          = 0x0004
	TMC5240_EM_POL_B                          = 0x0002
	TMC5240_EM_POL_A                          = 0x0001
	TMC5240_GCONF                             = 0x00
	TMC5240_GSTAT                             = 0x01
	TMC5240_IFCNT                             = 0x02
	TMC5240_SLAVECONF                         = 0x03
	TMC5240_INP_OUT                           = 0x04
	TMC5240_X_COMPARE                         = 0x05
	TMC5240_OTP_PROG                          = 0x06
	TMC5240_DRV_CONF                          = 0x0A
	TMC5240_GLOBAL_SCALER                     = 0x0B
	TMC5240_IHOLD_IRUN                        = 0x10
	TMC5240_TPOWERDOWN                        = 0x11
	TMC5240_TSTEP                             = 0x12
	TMC5240_TPWMTHRS                          = 0x13
	TMC5240_TCOOLTHRS                         = 0x14
	TMC5240_THIGH                             = 0x15
	TMC5240_RAMPMODE                          = 0x20
	TMC5240_XACTUAL                           = 0x21
	TMC5240_VACTUAL                           = 0x22
	TMC5240_VSTART                            = 0x23
	TMC5240_A1                                = 0x24
	TMC5240_V1                                = 0x25
	TMC5240_AMAX                              = 0x26
	TMC5240_VMAX                              = 0x27
	TMC5240_DMAX                              = 0x28
	TMC5240_TVMAX                             = 0x29
	TMC5240_D1                                = 0x2A
	TMC5240_VSTOP                             = 0x2B
	TMC5240_TZEROWAIT                         = 0x2C
	TMC5240_XTARGET                           = 0x2D
	TMC5240_V2                                = 0x2E
	TMC5240_A2                                = 0x2F
	TMC5240_D2                                = 0x30
	TMC5240_AACTUAL                           = 0x31
	TMC5240_VDCMIN                            = 0x33
	TMC5240_SWMODE                            = 0x34
	TMC5240_RAMPSTAT                          = 0x35
	TMC5240_XLATCH                            = 0x36
	TMC5240_ENCMODE                           = 0x38
	TMC5240_XENC                              = 0x39
	TMC5240_ENC_CONST                         = 0x3A
	TMC5240_ENC_STATUS                        = 0x3B
	TMC5240_ENC_LATCH                         = 0x3C
	TMC5240_ENC_DEVIATION                     = 0x3D
	TMC5240_VIRTUAL_STOP_L                    = 0x3E
	TMC5240_VIRTUAL_STOP_R                    = 0x3F
	TMC5240_ADC_VSUPPLY_AIN                   = 0x50
	TMC5240_ADC_TEMP                          = 0x51
	TMC5240_OTW_OV_VTH                        = 0x52
	TMC5240_MSLUT0                            = 0x60
	TMC5240_MSLUT1                            = 0x61
	TMC5240_MSLUT2                            = 0x62
	TMC5240_MSLUT3                            = 0x63
	TMC5240_MSLUT4                            = 0x64
	TMC5240_MSLUT5                            = 0x65
	TMC5240_MSLUT6                            = 0x66
	TMC5240_MSLUT7                            = 0x67
	TMC5240_MSLUTSEL                          = 0x68
	TMC5240_MSLUTSTART                        = 0x69
	TMC5240_MSCNT                             = 0x6A
	TMC5240_MSCURACT                          = 0x6B
	TMC5240_CHOPCONF                          = 0x6C
	TMC5240_COOLCONF                          = 0x6D
	TMC5240_DCCTRL                            = 0x6E
	TMC5240_DRVSTATUS                         = 0x6F
	TMC5240_PWMCONF                           = 0x70
	TMC5240_PWMSCALE                          = 0x71
	TMC5240_PWM_AUTO                          = 0x72
	TMC5240_SG4_THRS                          = 0x74
	TMC5240_SG4_RESULT                        = 0x75
	TMC5240_SG4_IND                           = 0x76
	TMC5240_SPI_STATUS_RESET_FLAG_MASK        = 0x01 /* GSTAT[0] - 1: Signals, that a reset has occurred (clear by reading GSTAT) */
	TMC5240_SPI_STATUS_RESET_FLAG_SHIFT       = 0
	TMC5240_SPI_STATUS_DRIVER_ERROR_MASK      = 0x02 /* GSTAT[1] – 1: Signals driver 1 driver error (clear by reading GSTAT) */
	TMC5240_SPI_STATUS_DRIVER_ERROR_SHIFT     = 1
	TMC5240_SPI_STATUS_SG2_MASK               = 0x04 /* DRV_STATUS[24] – 1: Signals StallGuard flag active */
	TMC5240_SPI_STATUS_SG2_SHIFT              = 2
	TMC5240_SPI_STATUS_STANDSTILL_MASK        = 0x08 /* DRV_STATUS[31] – 1: Signals motor stand still */
	TMC5240_SPI_STATUS_STANDSTILL_SHIFT       = 3
	TMC5240_SPI_STATUS_VELOCITY_REACHED_MASK  = 0x10 /* RAMP_STAT[8] – 1: Signals target velocity reached (motion controller only) */
	TMC5240_SPI_STATUS_VELOCITY_REACHED_SHIFT = 4
	TMC5240_SPI_STATUS_POSITION_REACHED_MASK  = 0x20 /* RAMP_STAT[9] – 1: Signals target position reached (motion controller only) */
	TMC5240_SPI_STATUS_POSITION_REACHED_SHIFT = 5
	TMC5240_SPI_STATUS_STATUS_STOP_L_MASK     = 0x40 /* RAMP_STAT[0] – 1: Signals stop left switch status (motion controller only) */
	TMC5240_SPI_STATUS_STATUS_STOP_L_SHIFT    = 6
	TMC5240_SPI_STATUS_STATUS_STOP_R_MASK     = 0x80 /* RAMP_STAT[1] – 1: Signals stop right switch status (motion controller only) */
	TMC5240_SPI_STATUS_STATUS_STOP_R_SHIFT    = 7
	TMC5240_FAST_STANDSTILL_MASK              = 0x02       // GCONF // Timeout for step execution until standstill detection
	TMC5240_FAST_STANDSTILL_SHIFT             = 1          // Timeout for step execution until standstill detection
	TMC5240_EN_PWM_MODE_MASK                  = 0x04       // GCONF // Enable the stealthChop(TM) mode
	TMC5240_EN_PWM_MODE_SHIFT                 = 2          // Enable the stealthChop(TM) mode
	TMC5240_MULTISTEP_FILT_MASK               = 0x08       // GCONF // Enable step input filtering for stealthChop
	TMC5240_MULTISTEP_FILT_SHIFT              = 3          // Enable step input filtering for stealthChop
	TMC5240_SHAFT_MASK                        = 0x10       // GCONF // Change motor direction / direction sign
	TMC5240_SHAFT_SHIFT                       = 4          // Change motor direction / direction sign
	TMC5240_DIAG0_ERROR_MASK                  = 0x20       // GCONF // DIAG0 output configuration (only with SD_MODE=1).; DIAG0 always shows the reset-status, i.e. is active low during reset condition.
	TMC5240_DIAG0_ERROR_SHIFT                 = 5          // DIAG0 output configuration (only with SD_MODE=1).; DIAG0 always shows the reset-status, i.e. is active low during reset condition.
	TMC5240_DIAG0_OTPW_MASK                   = 0x40       // GCONF // DIAG0 output configuration (only with SD_MODE=1).
	TMC5240_DIAG0_OTPW_SHIFT                  = 6          // DIAG0 output configuration (only with SD_MODE=1).
	TMC5240_DIAG0_STALL_STEP_MASK             = 0x80       // GCONF // DIAG0 output configuration (see differences depending on SD_MODE setting).
	TMC5240_DIAG0_STALL_STEP_SHIFT            = 7          // DIAG0 output configuration (see differences depending on SD_MODE setting).
	TMC5240_DIAG1_STALL_DIR_MASK              = 0x100      // GCONF // DIAG1 output configuration (see differences depending on SD_MODE setting).
	TMC5240_DIAG1_STALL_DIR_SHIFT             = 8          // DIAG1 output configuration (see differences depending on SD_MODE setting).
	TMC5240_DIAG1_INDEX_MASK                  = 0x200      // GCONF // DIAG1 output configuration.
	TMC5240_DIAG1_INDEX_SHIFT                 = 9          // DIAG1 output configuration.
	TMC5240_DIAG1_ONSTATE_MASK                = 0x400      // GCONF // DIAG1 output configuration.
	TMC5240_DIAG1_ONSTATE_SHIFT               = 10         // DIAG1 output configuration.
	TMC5240_DIAG0_INT_PUSHPULL_MASK           = 0x1000     // GCONF // DIAG0 output type configuration.
	TMC5240_DIAG0_INT_PUSHPULL_SHIFT          = 12         // DIAG0 output type configuration.
	TMC5240_DIAG1_POSCOMP_PUSHPULL_MASK       = 0x2000     // GCONF // DIAG1 output type configuration.
	TMC5240_DIAG1_POSCOMP_PUSHPULL_SHIFT      = 13         // DIAG1 output type configuration.
	TMC5240_SMALL_HYSTERESIS_MASK             = 0x4000     // GCONF // SMALL_HYSTERESIS
	TMC5240_SMALL_HYSTERESIS_SHIFT            = 14         // SMALL_HYSTERESIS
	TMC5240_STOP_ENABLE_MASK                  = 0x8000     // GCONF // Motor hard stop function enable.
	TMC5240_STOP_ENABLE_SHIFT                 = 15         // Motor hard stop function enable.
	TMC5240_DIRECT_MODE_MASK                  = 0x10000    // GCONF // Enable direct motor phase current control via serial interface.
	TMC5240_DIRECT_MODE_SHIFT                 = 16         // Enable direct motor phase current control via serial interface.
	TMC5240_LENGTH_STEP_PULSE_MASK            = 0x1E0000   // GCONF // cDriver only: length_step_pulse = 0: STEP output toggles upon each step;; length_step_pulse = 1...15: STEP pin high time in number of clock cycles
	TMC5240_LENGTH_STEP_PULSE_SHIFT           = 17         // cDriver only: length_step_pulse = 0: STEP output toggles upon each step;; length_step_pulse = 1...15: STEP pin high time in number of clock cycles
	TMC5240_RESET_MASK                        = 0x01       // GSTAT // Reset flag#type=COW
	TMC5240_RESET_SHIFT                       = 0          // Reset flag#type=COW
	TMC5240_DRV_ERR_MASK                      = 0x02       // GSTAT // Driver error flag#type=COW
	TMC5240_DRV_ERR_SHIFT                     = 1          // Driver error flag#type=COW
	TMC5240_UV_CP_MASK                        = 0x04       // GSTAT // Charge pump undervoltage condition flag#type=COW
	TMC5240_UV_CP_SHIFT                       = 2          // Charge pump undervoltage condition flag#type=COW
	TMC5240_REGISTER_RESET_MASK               = 0x08       // GSTAT // REGISTER_RESET
	TMC5240_REGISTER_RESET_SHIFT              = 3          // REGISTER_RESET
	TMC5240_VM_UVLO_MASK                      = 0x10       // GSTAT // 1: VM undervoltage has occured since last reset.
	TMC5240_VM_UVLO_SHIFT                     = 4          // 1: VM undervoltage has occured since last reset.
	TMC5240_IFCNT_MASK                        = 0xFF       // IFCNT // Interface transmission counter. This register becomes incremented with each successful UART interface write access. It can be read out to check the serial transmission for lost data. Read accesses do not change the content. Disabled in SPI operation. The counter wraps around from 255 to 0.
	TMC5240_IFCNT_SHIFT                       = 0          // Interface transmission counter. This register becomes incremented with each successful UART interface write access. It can be read out to check the serial transmission for lost data. Read accesses do not change the content. Disabled in SPI operation. The counter wraps around from 255 to 0.
	TMC5240_SLAVEADDR_MASK                    = 0xFF       // SLAVECONF // SLAVEADDR:; These eight bits set the address of unit for the UART interface. The address becomes incremented by one, two or three, as defined by SDI and SCK.; SCK, SDI; 00: +0; 01: +1; 10: +2; 11: +3; Range: 0-254 (do not increment beyond 254)
	TMC5240_SLAVEADDR_SHIFT                   = 0          // SLAVEADDR:; These eight bits set the address of unit for the UART interface. The address becomes incremented by one, two or three, as defined by SDI and SCK.; SCK, SDI; 00: +0; 01: +1; 10: +2; 11: +3; Range: 0-254 (do not increment beyond 254)
	TMC5240_SENDDELAY_MASK                    = 0xF00      // SLAVECONF // SWUART Slave Configuration
	TMC5240_SENDDELAY_SHIFT                   = 8          // SWUART Slave Configuration
	TMC5240_REFL_STEP_MASK                    = 0x01       // INP_OUT // REFL_STEP
	TMC5240_REFL_STEP_SHIFT                   = 0          // REFL_STEP
	TMC5240_REFR_DIR_MASK                     = 0x02       // INP_OUT // REFR_DIR
	TMC5240_REFR_DIR_SHIFT                    = 1          // REFR_DIR
	TMC5240_ENCB_CFG4_MASK                    = 0x04       // INP_OUT // B-channel state
	TMC5240_ENCB_CFG4_SHIFT                   = 2          // B-channel state
	TMC5240_ENCA_CFG5_MASK                    = 0x08       // INP_OUT // A-channel state
	TMC5240_ENCA_CFG5_SHIFT                   = 3          // A-channel state
	TMC5240_DRV_ENN_MASK                      = 0x10       // INP_OUT // Driver disabled/enabled state.
	TMC5240_DRV_ENN_SHIFT                     = 4          // Driver disabled/enabled state.
	TMC5240_ENCN_CFG6_MASK                    = 0x20       // INP_OUT // N-channel state
	TMC5240_ENCN_CFG6_SHIFT                   = 5          // N-channel state
	TMC5240_UART_EN_MASK                      = 0x40       // INP_OUT // 1 = UART interface is enabled
	TMC5240_UART_EN_SHIFT                     = 6          // 1 = UART interface is enabled
	TMC5240_COMP_A_MASK                       = 0x100      // INP_OUT // COMP_A (chopper comparator A, for IC test)
	TMC5240_COMP_A_SHIFT                      = 8          // COMP_A (chopper comparator A, for IC test)
	TMC5240_COMP_B_MASK                       = 0x200      // INP_OUT // COMP_B (chopper comparator B, for IC test)
	TMC5240_COMP_B_SHIFT                      = 9          // COMP_B (chopper comparator B, for IC test)
	TMC5240_COMP_A1_A2_MASK                   = 0x400      // INP_OUT // COMP_A1_A2 (StallGuard4 comparator A, for IC test)
	TMC5240_COMP_A1_A2_SHIFT                  = 10         // COMP_A1_A2 (StallGuard4 comparator A, for IC test)
	TMC5240_COMP_B1_B2_MASK                   = 0x800      // INP_OUT // COMP_B1_B2 (StallGuard4 comparator B, for IC test)
	TMC5240_COMP_B1_B2_SHIFT                  = 11         // COMP_B1_B2 (StallGuard4 comparator B, for IC test)
	TMC5240_OUTPUT_MASK                       = 0x1000     // INP_OUT // Output polarity of SDO_CFG0 pin when UART is enabled via pin UART_EN_CFG7. Its main purpose it to use SDO_CFG0 as NAO next address output signal for chain addressing of multiple ICs. Attention: Reset Value is 1 for use as NAO to next IC in single wire chain
	TMC5240_OUTPUT_SHIFT                      = 12         // Output polarity of SDO_CFG0 pin when UART is enabled via pin UART_EN_CFG7. Its main purpose it to use SDO_CFG0 as NAO next address output signal for chain addressing of multiple ICs. Attention: Reset Value is 1 for use as NAO to next IC in single wire chain
	TMC5240_EXT_RES_DET_MASK                  = 0x2000     // INP_OUT // 1: External resistor between REF and GND; 0: No external resistor detected
	TMC5240_EXT_RES_DET_SHIFT                 = 13         // 1: External resistor between REF and GND; 0: No external resistor detected
	TMC5240_EXT_CLK_MASK                      = 0x4000     // INP_OUT // 0: The internal oscillator is used for generating the clock-signal (12.5 MHz).; 1: The external oscillator is used for generating the clock-signal.
	TMC5240_EXT_CLK_SHIFT                     = 14         // 0: The internal oscillator is used for generating the clock-signal (12.5 MHz).; 1: The external oscillator is used for generating the clock-signal.
	TMC5240_ADC_ERR_MASK                      = 0x8000     // INP_OUT // 1: Signals that the ADC is not working correctly. Do not utilize ADC-features.Adc is stuck in configurationmode and very likely does not receive an ACK_OUT
	TMC5240_ADC_ERR_SHIFT                     = 15         // 1: Signals that the ADC is not working correctly. Do not utilize ADC-features.Adc is stuck in configurationmode and very likely does not receive an ACK_OUT
	TMC5240_SILICON_RV_MASK                   = 0x70000    // INP_OUT // Silicon revision number
	TMC5240_SILICON_RV_SHIFT                  = 16         // Silicon revision number
	TMC5240_VERSION_MASK                      = 0xFF000000 // INP_OUT // 0x40 = first version of the IC; Identical numbers mean full digital compatibility.
	TMC5240_VERSION_SHIFT                     = 24         // 0x40 = first version of the IC; Identical numbers mean full digital compatibility.
	TMC5240_X_COMPARE_MASK                    = 0xFFFFFFFF // X_COMPARE // Position comparison register for motion controller position strobe.; X_COMPARE is an absolute position.; The Position pulse is available on output SWP_DIAG1.; XACTUAL = X_COMPARE: Output signal PP (position pulse) becomes high.; It returns to a low state, if the positions mismatch.; If X_COMPARE_REPEAT is >1, X_COMPARE is the position reference for; the periodic position strobe trigger output.
	TMC5240_X_COMPARE_SHIFT                   = 0          // Position comparison register for motion controller position strobe.; X_COMPARE is an absolute position.; The Position pulse is available on output SWP_DIAG1.; XACTUAL = X_COMPARE: Output signal PP (position pulse) becomes high.; It returns to a low state, if the positions mismatch.; If X_COMPARE_REPEAT is >1, X_COMPARE is the position reference for; the periodic position strobe trigger output.
	TMC5240_X_COMPARE_REPEAT_MASK             = 0xFFFFFF   // X_COMPARE_REPEAT // This register defines a relative distance in microsteps (based on MRES confguration).; If set to >1, the position compare pulse is raised every time a multiple of X_PDISTANCE µsteps have been made.; Thereby, the X_COMPARE register defines the base position for the modulo calculation of X_PDISTANCE steps have been made into positive or negative direction.
	TMC5240_X_COMPARE_REPEAT_SHIFT            = 0          // This register defines a relative distance in microsteps (based on MRES confguration).; If set to >1, the position compare pulse is raised every time a multiple of X_PDISTANCE µsteps have been made.; Thereby, the X_COMPARE register defines the base position for the modulo calculation of X_PDISTANCE steps have been made into positive or negative direction.
	TMC5240_CURRENT_RANGE_MASK                = 0x03       // DRV_CONF // This setting allows a basic adaptation of the drivers RDSon current sensing to the motor current range. Select the lowest fitting range for best current precision. The value is the peak current setting.
	TMC5240_CURRENT_RANGE_SHIFT               = 0          // This setting allows a basic adaptation of the drivers RDSon current sensing to the motor current range. Select the lowest fitting range for best current precision. The value is the peak current setting.
	TMC5240_SLOPE_CONTROL_MASK                = 0x30       // DRV_CONF // Slope Control Setting
	TMC5240_SLOPE_CONTROL_SHIFT               = 4          // Slope Control Setting
	TMC5240_BBM_CLKS_MASK                     = 0xF00      // DRV_CONF // BBM_CLKS
	TMC5240_BBM_CLKS_SHIFT                    = 8          // BBM_CLKS
	TMC5240_RNDTF_MASK                        = 0x2000     // CHOPCONF // random TOFF time
	TMC5240_RNDTF_SHIFT                       = 13         // min.: 0, max.: 1, default: 0
	TMC5240_GLOBAL_SCALER_MASK                = 0xFF       // GLOBAL_SCALER // Global scaling of Motor current. This value is multiplied to the current scaling in order to adapt a drive to a certain motor type. This value should be chosen before tuning other settings, because it also influences chopper hysteresis.; 0:; Full Scale (or write 256); 1 … 31:; Not allowed for operation; 32 … 255:; 32/256 … 255/256 of maximum current.; Hint: Values >128 recommended for best results
	TMC5240_GLOBAL_SCALER_SHIFT               = 0          // Global scaling of Motor current. This value is multiplied to the current scaling in order to adapt a drive to a certain motor type. This value should be chosen before tuning other settings, because it also influences chopper hysteresis.; 0:; Full Scale (or write 256); 1 … 31:; Not allowed for operation; 32 … 255:; 32/256 … 255/256 of maximum current.; Hint: Values >128 recommended for best results
	TMC5240_IHOLD_MASK                        = 0x1F       // IHOLD_IRUN // Standstill current (0=1/32…31=32/32); In combination with StealthChop mode, setting IHOLD=0 allows to choose freewheeling or coil short circuit for motor stand still.
	TMC5240_IHOLD_SHIFT                       = 0          // Standstill current (0=1/32…31=32/32); In combination with StealthChop mode, setting IHOLD=0 allows to choose freewheeling or coil short circuit for motor stand still.
	TMC5240_IRUN_MASK                         = 0x1F00     // IHOLD_IRUN // Motor run current (0=1/32…31=32/32); Hint: Choose sense resistors in a way, that normal IRUN is 16 to 31 for best microstep performance.
	TMC5240_IRUN_SHIFT                        = 8          // Motor run current (0=1/32…31=32/32); Hint: Choose sense resistors in a way, that normal IRUN is 16 to 31 for best microstep performance.
	TMC5240_IHOLDDELAY_MASK                   = 0xF0000    // IHOLD_IRUN // Controls the number of clock cycles for motor power down after a motion as soon as standstill is detected (stst=1) and TPOWERDOWN has expired. The smooth transition avoids a motor jerk upon power down.; 0:; instant power down; 1..15:; Delay per current reduction step in multiple of 2^18 clocks
	TMC5240_IHOLDDELAY_SHIFT                  = 16         // Controls the number of clock cycles for motor power down after a motion as soon as standstill is detected (stst=1) and TPOWERDOWN has expired. The smooth transition avoids a motor jerk upon power down.; 0:; instant power down; 1..15:; Delay per current reduction step in multiple of 2^18 clocks
	TMC5240_IRUNDELAY_MASK                    = 0xF000000  // IHOLD_IRUN // Controls the number of clock cycles for motor power up after start is detected. 0:; instant power up 1..15:; Delay per current increment step in multiple of IRUNDELAY * 512 clocks
	TMC5240_IRUNDELAY_SHIFT                   = 24         // Controls the number of clock cycles for motor power up after start is detected. 0:; instant power up 1..15:; Delay per current increment step in multiple of IRUNDELAY * 512 clocks
	TMC5240_TPOWERDOWN_MASK                   = 0xFF       // TPOWERDOWN // TPOWERDOWN sets the delay time after stand still (stst) of the motor to motor current power down. Time range is about 0 to 4 seconds.; Attention: A minimum setting of 2 is required to allow automatic tuning of StealthChop PWM_OFFS_AUTO.; Reset Default = 10; 0…((2^8)-1) * 2^18 tCLK
	TMC5240_TPOWERDOWN_SHIFT                  = 0          // TPOWERDOWN sets the delay time after stand still (stst) of the motor to motor current power down. Time range is about 0 to 4 seconds.; Attention: A minimum setting of 2 is required to allow automatic tuning of StealthChop PWM_OFFS_AUTO.; Reset Default = 10; 0…((2^8)-1) * 2^18 tCLK
	TMC5240_TSTEP_MASK                        = 0xFFFFF    // TSTEP // Actual measured time between two 1/256 microsteps derived from the step input frequency in units of 1/fCLK. Measured value is (2^20)-1 in case of overflow or stand still.; All TSTEP related thresholds use a hysteresis of 1/16 of the compare value to compensate for jitter in the clock or the step frequency. The flag small_hysteresis modifies the hysteresis to a smaller value of 1/32.; (Txxx*15/16)-1 or; (Txxx*31/32)-1 is used as a second compare value for each comparison value.; This means, that the lower switching velocity equals the calculated setting, but the upper switching velocity is higher as defined by the hysteresis setting.; When working with the motion controller, the measured TSTEP for a given velocity V is in the range; (224 / V) = TSTEP = 224 / V - 1.; In DcStep mode TSTEP will not show the mean velocity of the motor, but the velocities for each microstep, which may not be stable and thus does not represent the real motor velocity in case it runs slower than the target velocity.
	TMC5240_TSTEP_SHIFT                       = 0          // Actual measured time between two 1/256 microsteps derived from the step input frequency in units of 1/fCLK. Measured value is (2^20)-1 in case of overflow or stand still.; All TSTEP related thresholds use a hysteresis of 1/16 of the compare value to compensate for jitter in the clock or the step frequency. The flag small_hysteresis modifies the hysteresis to a smaller value of 1/32.; (Txxx*15/16)-1 or; (Txxx*31/32)-1 is used as a second compare value for each comparison value.; This means, that the lower switching velocity equals the calculated setting, but the upper switching velocity is higher as defined by the hysteresis setting.; When working with the motion controller, the measured TSTEP for a given velocity V is in the range; (224 / V) = TSTEP = 224 / V - 1.; In DcStep mode TSTEP will not show the mean velocity of the motor, but the velocities for each microstep, which may not be stable and thus does not represent the real motor velocity in case it runs slower than the target velocity.
	TMC5240_TPWMTHRS_MASK                     = 0xFFFFF    // TPWMTHRS // This is the upper velocity for StealthChop voltage PWM mode.; TSTEP = TPWMTHRS; StealthChop PWM mode is enabled, if configured; DcStep is disabled
	TMC5240_TPWMTHRS_SHIFT                    = 0          // This is the upper velocity for StealthChop voltage PWM mode.; TSTEP = TPWMTHRS; StealthChop PWM mode is enabled, if configured; DcStep is disabled
	TMC5240_TCOOLTHRS_MASK                    = 0xFFFFF    // TCOOLTHRS // This is the lower threshold velocity for switching on smart energy CoolStep and StallGuard feature. (unsigned); Set this parameter to disable CoolStep at low speeds, where it cannot work reliably. The stop on stall function (enable with sg_stop when using internal motion controller) and the stall output signal become enabled when exceeding this velocity. In non-DcStep mode, it becomes disabled again once the velocity falls below this threshold.; TCOOLTHRS = TSTEP = THIGH:; CoolStep is enabled, if configured; TCOOLTHRS = TSTEP; Stop on stall is enabled, if configured; Stall output signal (DIAG0/1) is enabled, if configured
	TMC5240_TCOOLTHRS_SHIFT                   = 0          // This is the lower threshold velocity for switching on smart energy CoolStep and StallGuard feature. (unsigned); Set this parameter to disable CoolStep at low speeds, where it cannot work reliably. The stop on stall function (enable with sg_stop when using internal motion controller) and the stall output signal become enabled when exceeding this velocity. In non-DcStep mode, it becomes disabled again once the velocity falls below this threshold.; TCOOLTHRS = TSTEP = THIGH:; CoolStep is enabled, if configured; TCOOLTHRS = TSTEP; Stop on stall is enabled, if configured; Stall output signal (DIAG0/1) is enabled, if configured
	TMC5240_THIGH_MASK                        = 0xFFFFF    // THIGH // This velocity setting allows velocity dependent switching into a different chopper mode and fullstepping to maximize torque. (unsigned); The stall detection feature becomes switched off for 2-3 electrical periods whenever passing THIGH threshold to compensate for the effect of switching modes.; TSTEP = THIGH:; CoolStep is disabled (motor runs with normal current scale); StealthChop voltage PWM mode is disabled; If vhighchm is set, the chopper switches to chm=1 with TFD=0 (constant off time with slow decay, only).; If vhighfs is set, the motor operates in fullstep mode and the stall detection becomes switched over to DcStep stall detection.
	TMC5240_THIGH_SHIFT                       = 0          // This velocity setting allows velocity dependent switching into a different chopper mode and fullstepping to maximize torque. (unsigned); The stall detection feature becomes switched off for 2-3 electrical periods whenever passing THIGH threshold to compensate for the effect of switching modes.; TSTEP = THIGH:; CoolStep is disabled (motor runs with normal current scale); StealthChop voltage PWM mode is disabled; If vhighchm is set, the chopper switches to chm=1 with TFD=0 (constant off time with slow decay, only).; If vhighfs is set, the motor operates in fullstep mode and the stall detection becomes switched over to DcStep stall detection.
	TMC5240_RAMPMODE_MASK                     = 0x03       // RAMPMODE // Motion Controller ramping mode
	TMC5240_RAMPMODE_SHIFT                    = 0          // Motion Controller ramping mode
	TMC5240_XACTUAL_MASK                      = 0xFFFFFFFF // XACTUAL // Actual motor position (signed); Hint: This value normally should only be modified, when homing the drive. In positioning mode, modifying the register content will start a motion.
	TMC5240_XACTUAL_SHIFT                     = 0          // Actual motor position (signed); Hint: This value normally should only be modified, when homing the drive. In positioning mode, modifying the register content will start a motion.
	TMC5240_VACTUAL_MASK                      = 0xFFFFFF   // VACTUAL // Actual motor velocity from ramp generator (signed); The sign matches the motion direction. A negative sign means motion to lower XACTUAL.; +-(2^23)-1; [µsteps / t]
	TMC5240_VACTUAL_SHIFT                     = 0          // Actual motor velocity from ramp generator (signed); The sign matches the motion direction. A negative sign means motion to lower XACTUAL.; +-(2^23)-1; [µsteps / t]
	TMC5240_VSTART_MASK                       = 0x3FFFF    // VSTART // Motor start velocity (unsigned); For universal use, set VSTOP = VSTART. This is not required if the motion distance is sufficient to ensure deceleration from VSTART to VSTOP.; 0…(2^18)-1; [µsteps / t]
	TMC5240_VSTART_SHIFT                      = 0          // Motor start velocity (unsigned); For universal use, set VSTOP = VSTART. This is not required if the motion distance is sufficient to ensure deceleration from VSTART to VSTOP.; 0…(2^18)-1; [µsteps / t]
	TMC5240_A1_MASK                           = 0x3FFFF    // A1 // First acceleration between VSTART and V1 (unsigned); 0…(2^18)-1; [µsteps / ta²]
	TMC5240_A1_SHIFT                          = 0          // First acceleration between VSTART and V1 (unsigned); 0…(2^18)-1; [µsteps / ta²]
	TMC5240_V1_MASK                           = 0xFFFFF    // V1 // First acceleration / deceleration phase threshold velocity (unsigned); 0: Disables A1 and D1 phase, use AMAX, DMAX only; 0…(2^20)-1; [µsteps / t]
	TMC5240_V1_SHIFT                          = 0          // First acceleration / deceleration phase threshold velocity (unsigned); 0: Disables A1 and D1 phase, use AMAX, DMAX only; 0…(2^20)-1; [µsteps / t]
	TMC5240_AMAX_MASK                         = 0x3FFFF    // AMAX // Second acceleration between V1 and VMAX (unsigned); This is the acceleration and deceleration value for velocity mode.; 0…(2^18)-1; [µsteps / ta²]
	TMC5240_AMAX_SHIFT                        = 0          // Second acceleration between V1 and VMAX (unsigned); This is the acceleration and deceleration value for velocity mode.; 0…(2^18)-1; [µsteps / ta²]
	TMC5240_VMAX_MASK                         = 0x7FFFFF   // VMAX // Motion ramp target velocity (for positioning ensure VMAX = VSTART) (unsigned); This is the target velocity in velocity mode. It can be changed any time during a motion.; 0…(2^23)-512; [µsteps / t]
	TMC5240_VMAX_SHIFT                        = 0          // Motion ramp target velocity (for positioning ensure VMAX = VSTART) (unsigned); This is the target velocity in velocity mode. It can be changed any time during a motion.; 0…(2^23)-512; [µsteps / t]
	TMC5240_DMAX_MASK                         = 0x3FFFF    // DMAX // Deceleration between VMAX and V1 (unsigned); 0…(2^18)-1; [µsteps / ta²]
	TMC5240_DMAX_SHIFT                        = 0          // Deceleration between VMAX and V1 (unsigned); 0…(2^18)-1; [µsteps / ta²]
	TMC5240_TVMAX_MASK                        = 0xFFFF     // TVMAX // Minimum time for constant velocity segments in multiple of 512 clocks.; 0: Disables minimum duration setting for constant velocity phase; >0: A minimum duration of constant velocity is inserted in between of any change from acceleration to deceleration or vice versa to reduce jerk; 0…(2^16)-1 * 512 tCLK; Note: Configure this register after setting VMAX when in Positionmode and standstill. Set TVMAX=0 during velocity mode to avoid triggering the TVMAX delay when switching back to ramp mode.
	TMC5240_TVMAX_SHIFT                       = 0          // Minimum time for constant velocity segments in multiple of 512 clocks.; 0: Disables minimum duration setting for constant velocity phase; >0: A minimum duration of constant velocity is inserted in between of any change from acceleration to deceleration or vice versa to reduce jerk; 0…(2^16)-1 * 512 tCLK; Note: Configure this register after setting VMAX when in Positionmode and standstill. Set TVMAX=0 during velocity mode to avoid triggering the TVMAX delay when switching back to ramp mode.
	TMC5240_D1_MASK                           = 0x3FFFF    // D1 // Deceleration between V1 and VSTOP (unsigned); Attention: Do not set 0 in positioning mode, even if V1=0!; 1…(2^16)-1; [µsteps / ta²]; Reset Default = 10
	TMC5240_D1_SHIFT                          = 0          // Deceleration between V1 and VSTOP (unsigned); Attention: Do not set 0 in positioning mode, even if V1=0!; 1…(2^16)-1; [µsteps / ta²]; Reset Default = 10
	TMC5240_VSTOP_MASK                        = 0x3FFFF    // VSTOP // Motor stop velocity (unsigned); Hint: Set VSTOP = VSTART to allow positioning for short distances; Attention: Do not set 0 in positioning mode, minimum 10 recommend!; 1…(2^18)-1; [µsteps / t]; Reset Default=10
	TMC5240_VSTOP_SHIFT                       = 0          // Motor stop velocity (unsigned); Hint: Set VSTOP = VSTART to allow positioning for short distances; Attention: Do not set 0 in positioning mode, minimum 10 recommend!; 1…(2^18)-1; [µsteps / t]; Reset Default=10
	TMC5240_TZEROWAIT_MASK                    = 0xFFFF     // TZEROWAIT // Defines the waiting time after ramping down to zero velocity before next movement or direction inversion can start. Time range is about 0 to 2 seconds.; This setting avoids excess acceleration e.g. from VSTOP to -VSTART.; 0…(2^16)-1 * 512 tCLK
	TMC5240_TZEROWAIT_SHIFT                   = 0          // Defines the waiting time after ramping down to zero velocity before next movement or direction inversion can start. Time range is about 0 to 2 seconds.; This setting avoids excess acceleration e.g. from VSTOP to -VSTART.; 0…(2^16)-1 * 512 tCLK
	TMC5240_XTARGET_MASK                      = 0xFFFFFFFF // XTARGET // Target position for ramp mode (signed). Write a new target position to this register in order to activate the ramp generator positioning in RAMPMODE=0. Initialize all velocity, acceleration and deceleration parameters before.; Hint: The position is allowed to wrap around, thus, XTARGET value optionally can be treated as an unsigned number.; Hint: The maximum possible displacement is +/-((2^31)-1).; Hint: When increasing V1, D1 or DMAX during a motion, rewrite XTARGET afterwards in order to trigger a second acceleration phase, if desired.; -2^31…; +(2^31)-1
	TMC5240_XTARGET_SHIFT                     = 0          // Target position for ramp mode (signed). Write a new target position to this register in order to activate the ramp generator positioning in RAMPMODE=0. Initialize all velocity, acceleration and deceleration parameters before.; Hint: The position is allowed to wrap around, thus, XTARGET value optionally can be treated as an unsigned number.; Hint: The maximum possible displacement is +/-((2^31)-1).; Hint: When increasing V1, D1 or DMAX during a motion, rewrite XTARGET afterwards in order to trigger a second acceleration phase, if desired.; -2^31…; +(2^31)-1
	TMC5240_V2_MASK                           = 0xFFFFF    // V2 // Velocity difference from VMAX for activation of acceleration segments with AMAX/2 and DMAX/2.; 0: Disables AMAX/2 and DMAX/2 phase, use AMAX, DMAX only; 0…(2^20)-1; [µsteps / t]
	TMC5240_V2_SHIFT                          = 0          // Velocity difference from VMAX for activation of acceleration segments with AMAX/2 and DMAX/2.; 0: Disables AMAX/2 and DMAX/2 phase, use AMAX, DMAX only; 0…(2^20)-1; [µsteps / t]
	TMC5240_A2_MASK                           = 0x3FFFF    // A2 // Acceleration between V1 and V2 (unsigned); 0…(2^18)-1 [µsteps / ta²]
	TMC5240_A2_SHIFT                          = 0          // Acceleration between V1 and V2 (unsigned); 0…(2^18)-1 [µsteps / ta²]
	TMC5240_D2_MASK                           = 0x3FFFF    // D2 // Deceleration between V2 and V1 (unsigned); Attention: Do not set 0 in positioning mode, even if V2=0!; 1…(2^18)-1; [µsteps / ta²]; Reset Default = 10
	TMC5240_D2_SHIFT                          = 0          // Deceleration between V2 and V1 (unsigned); Attention: Do not set 0 in positioning mode, even if V2=0!; 1…(2^18)-1; [µsteps / ta²]; Reset Default = 10
	TMC5240_AACTUAL_MASK                      = 0xFFFFFF   // AACTUAL // Current acceleration used by the rampgenerator
	TMC5240_AACTUAL_SHIFT                     = 0          // Current acceleration used by the rampgenerator
	TMC5240_VDCMIN_MASK                       = 0x7FFF00   // VDCMIN // Automatic commutation DcStep becomes enabled above velocity VDCMIN (unsigned) (only when using internal ramp generator, not for STEP/DIR interface modes – in STEP/DIR mode, DcStep is not available); In this mode, the actual position is determined by the sensor­less motor commutation and becomes fed back to XACTUAL. In case the motor becomes heavily loaded, VDCMIN also is used as the minimum step velocity. Activate stop on stall (sg_stop) to detect step loss.; 0: Disable, DcStep off; |VACT| = VDCMIN = 256:; Triggers the same actions as exceeding THIGH setting.; Switches on automatic commutation DcStep; Hint: Also set DCCTRL parameters in order to operate DcStep.; (Only bits 22… 8 are used for value and for comparison)
	TMC5240_VDCMIN_SHIFT                      = 8          // Automatic commutation DcStep becomes enabled above velocity VDCMIN (unsigned) (only when using internal ramp generator, not for STEP/DIR interface modes – in STEP/DIR mode, DcStep is not available); In this mode, the actual position is determined by the sensor­less motor commutation and becomes fed back to XACTUAL. In case the motor becomes heavily loaded, VDCMIN also is used as the minimum step velocity. Activate stop on stall (sg_stop) to detect step loss.; 0: Disable, DcStep off; |VACT| = VDCMIN = 256:; Triggers the same actions as exceeding THIGH setting.; Switches on automatic commutation DcStep; Hint: Also set DCCTRL parameters in order to operate DcStep.; (Only bits 22… 8 are used for value and for comparison)
	TMC5240_STOP_L_ENABLE_MASK                = 0x01       // SW_MODE // 1: Enables automatic motor stop during active left reference switch input; Hint: The motor restarts in case the stop switch becomes released.
	TMC5240_STOP_L_ENABLE_SHIFT               = 0          // 1: Enables automatic motor stop during active left reference switch input; Hint: The motor restarts in case the stop switch becomes released.
	TMC5240_STOP_R_ENABLE_MASK                = 0x02       // SW_MODE // 1: Enables automatic motor stop during active right reference switch input; Hint: The motor restarts in case the stop switch becomes released.
	TMC5240_STOP_R_ENABLE_SHIFT               = 1          // 1: Enables automatic motor stop during active right reference switch input; Hint: The motor restarts in case the stop switch becomes released.
	TMC5240_POL_STOP_L_MASK                   = 0x04       // SW_MODE // Sets the active polarity of the left reference switch input; 0=non-inverted, high active: a high level on REFL stops the motor; 1=inverted, low active: a low level on REFL stops the motor
	TMC5240_POL_STOP_L_SHIFT                  = 2          // Sets the active polarity of the left reference switch input; 0=non-inverted, high active: a high level on REFL stops the motor; 1=inverted, low active: a low level on REFL stops the motor
	TMC5240_POL_STOP_R_MASK                   = 0x08       // SW_MODE // Sets the active polarity of the right reference switch input; 0=non-inverted, high active: a high level on REFR stops the motor; 1=inverted, low active: a low level on REFR stops the motor
	TMC5240_POL_STOP_R_SHIFT                  = 3          // Sets the active polarity of the right reference switch input; 0=non-inverted, high active: a high level on REFR stops the motor; 1=inverted, low active: a low level on REFR stops the motor
	TMC5240_SWAP_LR_MASK                      = 0x10       // SW_MODE // 1: Swap the left and the right reference switch input REFL and REFR
	TMC5240_SWAP_LR_SHIFT                     = 4          // 1: Swap the left and the right reference switch input REFL and REFR
	TMC5240_LATCH_L_ACTIVE_MASK               = 0x20       // SW_MODE // 1: Activates latching of the position to XLATCH upon an active going edge on the left reference switch input REFL.; Hint: Activate latch_l_active to detect any spurious stop event by reading status_latch_l.
	TMC5240_LATCH_L_ACTIVE_SHIFT              = 5          // 1: Activates latching of the position to XLATCH upon an active going edge on the left reference switch input REFL.; Hint: Activate latch_l_active to detect any spurious stop event by reading status_latch_l.
	TMC5240_LATCH_L_INACTIVE_MASK             = 0x40       // SW_MODE // 1: Activates latching of the position to XLATCH upon an inactive going edge on the left reference switch input REFL. The active level is defined by pol_stop_l.
	TMC5240_LATCH_L_INACTIVE_SHIFT            = 6          // 1: Activates latching of the position to XLATCH upon an inactive going edge on the left reference switch input REFL. The active level is defined by pol_stop_l.
	TMC5240_LATCH_R_ACTIVE_MASK               = 0x80       // SW_MODE // 1: Activates latching of the position to XLATCH upon an active going edge on the right reference switch input REFR.; Hint: Activate latch_r_active to detect any spurious stop event by reading status_latch_r.
	TMC5240_LATCH_R_ACTIVE_SHIFT              = 7          // 1: Activates latching of the position to XLATCH upon an active going edge on the right reference switch input REFR.; Hint: Activate latch_r_active to detect any spurious stop event by reading status_latch_r.
	TMC5240_LATCH_R_INACTIVE_MASK             = 0x100      // SW_MODE // 1: Activates latching of the position to XLATCH upon an inactive going edge on the right reference switch input REFR. The active level is defined by pol_stop_r.
	TMC5240_LATCH_R_INACTIVE_SHIFT            = 8          // 1: Activates latching of the position to XLATCH upon an inactive going edge on the right reference switch input REFR. The active level is defined by pol_stop_r.
	TMC5240_EN_LATCH_ENCODER_MASK             = 0x200      // SW_MODE // 1: Latch encoder position to ENC_LATCH upon reference switch event.
	TMC5240_EN_LATCH_ENCODER_SHIFT            = 9          // 1: Latch encoder position to ENC_LATCH upon reference switch event.
	TMC5240_SG_STOP_MASK                      = 0x400      // SW_MODE // Enable stop by StallGuard2 (also available in DcStep mode). Disable to release motor after stop event. Program TCOOLTHRS for velocity threshold.; Hint: Do not enable during motor spin-up, wait until the motor velocity exceeds a certain value, where StallGuard2 delivers a stable result. This velocity threshold should be programmed using TCOOLTHRS.
	TMC5240_SG_STOP_SHIFT                     = 10         // Enable stop by StallGuard2 (also available in DcStep mode). Disable to release motor after stop event. Program TCOOLTHRS for velocity threshold.; Hint: Do not enable during motor spin-up, wait until the motor velocity exceeds a certain value, where StallGuard2 delivers a stable result. This velocity threshold should be programmed using TCOOLTHRS.
	TMC5240_EN_SOFTSTOP_MASK                  = 0x800      // SW_MODE // 0: Hard stop; 1: Soft stop; The soft stop mode always uses the deceleration ramp settings DMAX, V1, D1, V2, D2, VSTOP and TZEROWAIT for stopping the motor. A stop occurs when the velocity sign matches the reference switch position (REFL, VIRTUAL_STOP_L for negative velocities, REFR, VIRTUAL_STOP_R for positive velocities) and the respective switch stop function is enabled.; A hard stop also uses TZEROWAIT before the motor becomes released.; Attention: Do not use soft stop in combination with StallGuard2. Use soft stop for StealthChop operation at high velocity. In this case, hard stop must be avoided, as it could result in severe overcurrent.
	TMC5240_EN_SOFTSTOP_SHIFT                 = 11         // 0: Hard stop; 1: Soft stop; The soft stop mode always uses the deceleration ramp settings DMAX, V1, D1, V2, D2, VSTOP and TZEROWAIT for stopping the motor. A stop occurs when the velocity sign matches the reference switch position (REFL, VIRTUAL_STOP_L for negative velocities, REFR, VIRTUAL_STOP_R for positive velocities) and the respective switch stop function is enabled.; A hard stop also uses TZEROWAIT before the motor becomes released.; Attention: Do not use soft stop in combination with StallGuard2. Use soft stop for StealthChop operation at high velocity. In this case, hard stop must be avoided, as it could result in severe overcurrent.
	TMC5240_EN_VIRTUAL_STOP_L_MASK            = 0x1000     // SW_MODE // 1: Enables automatic motor stop during active left virtual stop condition
	TMC5240_EN_VIRTUAL_STOP_L_SHIFT           = 12         // 1: Enables automatic motor stop during active left virtual stop condition
	TMC5240_EN_VIRTUAL_STOP_R_MASK            = 0x2000     // SW_MODE // 1: Enables automatic motor stop during active right virtual stop condition
	TMC5240_EN_VIRTUAL_STOP_R_SHIFT           = 13         // 1: Enables automatic motor stop during active right virtual stop condition
	TMC5240_VIRTUAL_STOP_ENC_MASK             = 0x4000     // SW_MODE // Source for virtual stop (VIRTUAL_STOP_L and VIRTUAL_STOP_R); 0: Virtual stop relates to ramp generator position XACTUAL; 1: Virtual stop relates to encoder position X_ENC
	TMC5240_VIRTUAL_STOP_ENC_SHIFT            = 14         // Source for virtual stop (VIRTUAL_STOP_L and VIRTUAL_STOP_R); 0: Virtual stop relates to ramp generator position XACTUAL; 1: Virtual stop relates to encoder position X_ENC
	TMC5240_STATUS_STOP_L_MASK                = 0x01       // RAMP_STAT // Reference switch left status (1=active)
	TMC5240_STATUS_STOP_L_SHIFT               = 0          // Reference switch left status (1=active)
	TMC5240_STATUS_STOP_R_MASK                = 0x02       // RAMP_STAT // Reference switch right status (1=active)
	TMC5240_STATUS_STOP_R_SHIFT               = 1          // Reference switch right status (1=active)
	TMC5240_STATUS_LATCH_L_MASK               = 0x04       // RAMP_STAT // 1: Latch left ready; (enable position latching using SW_MODE settings; latch_l_active or latch_l_inactive); (Write ‘1’ to clear)#type=COW
	TMC5240_STATUS_LATCH_L_SHIFT              = 2          // 1: Latch left ready; (enable position latching using SW_MODE settings; latch_l_active or latch_l_inactive); (Write ‘1’ to clear)#type=COW
	TMC5240_STATUS_LATCH_R_MASK               = 0x08       // RAMP_STAT // 1: Latch right ready; (enable position latching using SW_MODE settings; latch_r_active or latch_r_inactive); (Write ‘1’ to clear)#type=COW
	TMC5240_STATUS_LATCH_R_SHIFT              = 3          // 1: Latch right ready; (enable position latching using SW_MODE settings; latch_r_active or latch_r_inactive); (Write ‘1’ to clear)#type=COW
	TMC5240_EVENT_STOP_L_MASK                 = 0x10       // RAMP_STAT // 1: Active stop left condition due to stop switch or virtual stop.; The stop condition and the interrupt condition can be removed by setting RAMP_MODE to hold mode or by commanding a move to the opposite direction. In soft_stop mode, the condition will remain active until the motor has stopped motion into the direction of the stop switch. Disabling the stop switch or the stop function also clears the flag, but the motor will continue motion.; This bit is ORed to the interrupt output signal.
	TMC5240_EVENT_STOP_L_SHIFT                = 4          // 1: Active stop left condition due to stop switch or virtual stop.; The stop condition and the interrupt condition can be removed by setting RAMP_MODE to hold mode or by commanding a move to the opposite direction. In soft_stop mode, the condition will remain active until the motor has stopped motion into the direction of the stop switch. Disabling the stop switch or the stop function also clears the flag, but the motor will continue motion.; This bit is ORed to the interrupt output signal.
	TMC5240_EVENT_STOP_R_MASK                 = 0x20       // RAMP_STAT // 1: Active stop right condition due to stop switch or virtual stop.; The stop condition and the interrupt condition can be removed by setting RAMP_MODE to hold mode or by commanding a move to the opposite direction. In soft_stop mode, the condition will remain active until the motor has stopped motion into the direction of the stop switch. Disabling the stop switch or the stop function also clears the flag, but the motor will continue motion.; This bit is ORed to the interrupt output signal.
	TMC5240_EVENT_STOP_R_SHIFT                = 5          // 1: Active stop right condition due to stop switch or virtual stop.; The stop condition and the interrupt condition can be removed by setting RAMP_MODE to hold mode or by commanding a move to the opposite direction. In soft_stop mode, the condition will remain active until the motor has stopped motion into the direction of the stop switch. Disabling the stop switch or the stop function also clears the flag, but the motor will continue motion.; This bit is ORed to the interrupt output signal.
	TMC5240_EVENT_STOP_SG_MASK                = 0x40       // RAMP_STAT // 1: Signals an active StallGuard2 stop event.; Resetting the register will clear the stall condition and the motor may re-start motion, unless the motion controller has been stopped.; (Write ‘1’ to clear flag and interrupt condition); This bit is ORed to the interrupt output signal.#type=COW
	TMC5240_EVENT_STOP_SG_SHIFT               = 6          // 1: Signals an active StallGuard2 stop event.; Resetting the register will clear the stall condition and the motor may re-start motion, unless the motion controller has been stopped.; (Write ‘1’ to clear flag and interrupt condition); This bit is ORed to the interrupt output signal.#type=COW
	TMC5240_EVENT_POS_REACHED_MASK            = 0x80       // RAMP_STAT // 1: Signals, that the target position has been reached (position_reached becoming active).; (Write ‘1’ to clear flag and interrupt condition); This bit is ORed to the interrupt output signal.#type=COW
	TMC5240_EVENT_POS_REACHED_SHIFT           = 7          // 1: Signals, that the target position has been reached (position_reached becoming active).; (Write ‘1’ to clear flag and interrupt condition); This bit is ORed to the interrupt output signal.#type=COW
	TMC5240_VELOCITY_REACHED_MASK             = 0x100      // RAMP_STAT // 1: Signals, that the target velocity is reached.; This flag becomes set while VACTUAL and VMAX match.
	TMC5240_VELOCITY_REACHED_SHIFT            = 8          // 1: Signals, that the target velocity is reached.; This flag becomes set while VACTUAL and VMAX match.
	TMC5240_POSITION_REACHED_MASK             = 0x200      // RAMP_STAT // 1: Signals, that the target position is reached.; This flag becomes set while XACTUAL and XTARGET match.
	TMC5240_POSITION_REACHED_SHIFT            = 9          // 1: Signals, that the target position is reached.; This flag becomes set while XACTUAL and XTARGET match.
	TMC5240_VZERO_MASK                        = 0x400      // RAMP_STAT // 1: Signals, that the actual velocity is 0.
	TMC5240_VZERO_SHIFT                       = 10         // 1: Signals, that the actual velocity is 0.
	TMC5240_T_ZEROWAIT_ACTIVE_MASK            = 0x800      // RAMP_STAT // 1: Signals, that TZEROWAIT is active after a motor stop. During this time, the motor is in standstill.
	TMC5240_T_ZEROWAIT_ACTIVE_SHIFT           = 11         // 1: Signals, that TZEROWAIT is active after a motor stop. During this time, the motor is in standstill.
	TMC5240_SECOND_MOVE_MASK                  = 0x1000     // RAMP_STAT // 1: Signals that the automatic ramp required moving back in the opposite direction, e.g. due to on-the-fly parameter change; (Write ‘1’ to clear)#type=COW
	TMC5240_SECOND_MOVE_SHIFT                 = 12         // 1: Signals that the automatic ramp required moving back in the opposite direction, e.g. due to on-the-fly parameter change; (Write ‘1’ to clear)#type=COW
	TMC5240_STATUS_SG_MASK                    = 0x2000     // RAMP_STAT // 1: Signals an active StallGuard2 input from the CoolStep driver or from the DcStep unit, if enabled.; Hint: When polling this flag, stall events may be missed – activate sg_stop to be sure not to miss the stall event.
	TMC5240_STATUS_SG_SHIFT                   = 13         // 1: Signals an active StallGuard2 input from the CoolStep driver or from the DcStep unit, if enabled.; Hint: When polling this flag, stall events may be missed – activate sg_stop to be sure not to miss the stall event.
	TMC5240_STATUS_VIRTUAL_STOP_L_MASK        = 0x4000     // RAMP_STAT // Virtual reference switch left status (1=active)
	TMC5240_STATUS_VIRTUAL_STOP_L_SHIFT       = 14         // Virtual reference switch left status (1=active)
	TMC5240_STATUS_VIRTUAL_STOP_R_MASK        = 0x8000     // RAMP_STAT // Virtual reference switch right status (1=active)
	TMC5240_STATUS_VIRTUAL_STOP_R_SHIFT       = 15         // Virtual reference switch right status (1=active)
	TMC5240_XLATCH_MASK                       = 0xFFFFFFFF // XLATCH // Ramp generator latch position, latches XACTUAL upon a programmable switch event (see SW_MODE).; Hint: The encoder position can be latched to ENC_LATCH together with XLATCH to allow consistency checks.
	TMC5240_XLATCH_SHIFT                      = 0          // Ramp generator latch position, latches XACTUAL upon a programmable switch event (see SW_MODE).; Hint: The encoder position can be latched to ENC_LATCH together with XLATCH to allow consistency checks.
	TMC5240_POL_A_MASK                        = 0x01       // ENCMODE // Required A polarity for an N channel event
	TMC5240_POL_A_SHIFT                       = 0          // Required A polarity for an N channel event
	TMC5240_POL_B_MASK                        = 0x02       // ENCMODE // Required B polarity for an N channel event
	TMC5240_POL_B_SHIFT                       = 1          // Required B polarity for an N channel event
	TMC5240_POL_N_MASK                        = 0x04       // ENCMODE // Defines active polarity of N
	TMC5240_POL_N_SHIFT                       = 2          // Defines active polarity of N
	TMC5240_IGNORE_AB_MASK                    = 0x08       // ENCMODE // N event configuration
	TMC5240_IGNORE_AB_SHIFT                   = 3          // N event configuration
	TMC5240_CLR_CONT_MASK                     = 0x10       // ENCMODE // Position latch configuration
	TMC5240_CLR_CONT_SHIFT                    = 4          // Position latch configuration
	TMC5240_CLR_ONCE_MASK                     = 0x20       // ENCMODE // Position latch configuration
	TMC5240_CLR_ONCE_SHIFT                    = 5          // Position latch configuration
	TMC5240_POS_NEG_EDGE_MASK                 = 0xC0       // ENCMODE // N channel event sensitivity
	TMC5240_POS_NEG_EDGE_SHIFT                = 6          // N channel event sensitivity
	TMC5240_CLR_ENC_X_MASK                    = 0x100      // ENCMODE // Encoder latch configuration
	TMC5240_CLR_ENC_X_SHIFT                   = 8          // Encoder latch configuration
	TMC5240_LATCH_X_ACT_MASK                  = 0x200      // ENCMODE // Position latch configuration
	TMC5240_LATCH_X_ACT_SHIFT                 = 9          // Position latch configuration
	TMC5240_ENC_SEL_DECIMAL_MASK              = 0x400      // ENCMODE // Encoder prescaler mode selection
	TMC5240_ENC_SEL_DECIMAL_SHIFT             = 10         // Encoder prescaler mode selection
	TMC5240_X_ENC_MASK                        = 0xFFFFFFFF // X_ENC // Actual encoder position (signed)
	TMC5240_X_ENC_SHIFT                       = 0          // Actual encoder position (signed)
	TMC5240_ENC_CONST_MASK                    = 0xFFFFFFFF // ENC_CONST // Accumulation constant (signed); 16 bit integer part, 16 bit fractional part; X_ENC accumulates; +/- ENC_CONST / (2^16*X_ENC) (binary); or; +/-ENC_CONST / (10^4*X_ENC) (decimal); ENCMODE bit enc_sel_decimal switches between decimal and binary setting.; Use the sign, to match rotation direction!; binary:; ± [µsteps/2^16]; ±(0 …; 32767.999847); decimal:; ±(0.0 … 32767.9999); reset default = 1.0 (=65536)
	TMC5240_ENC_CONST_SHIFT                   = 0          // Accumulation constant (signed); 16 bit integer part, 16 bit fractional part; X_ENC accumulates; +/- ENC_CONST / (2^16*X_ENC) (binary); or; +/-ENC_CONST / (10^4*X_ENC) (decimal); ENCMODE bit enc_sel_decimal switches between decimal and binary setting.; Use the sign, to match rotation direction!; binary:; ± [µsteps/2^16]; ±(0 …; 32767.999847); decimal:; ±(0.0 … 32767.9999); reset default = 1.0 (=65536)
	TMC5240_N_EVENT_MASK                      = 0x01       // ENC_STATUS // #type=COW
	TMC5240_N_EVENT_SHIFT                     = 0          // #type=COW
	TMC5240_DEVIATION_WARN_MASK               = 0x02       // ENC_STATUS // #type=COW
	TMC5240_DEVIATION_WARN_SHIFT              = 1          // #type=COW
	TMC5240_ENC_LATCH_MASK                    = 0xFFFFFFFF // ENC_LATCH // Encoder position X_ENC latched on N event
	TMC5240_ENC_LATCH_SHIFT                   = 0          // Encoder position X_ENC latched on N event
	TMC5240_ENC_DEVIATION_MASK                = 0xFFFFF    // ENC_DEVIATION // Maximum number of steps deviation between encoder counter and XACTUAL for deviation warning; Result in flag ENC_STATUS.deviation_warn; 0=Function is off.
	TMC5240_ENC_DEVIATION_SHIFT               = 0          // Maximum number of steps deviation between encoder counter and XACTUAL for deviation warning; Result in flag ENC_STATUS.deviation_warn; 0=Function is off.
	TMC5240_VIRTUAL_STOP_L_MASK               = 0xFFFFFFFF // VIRTUAL_STOP_L // Virtual stop switch based on encoder or ramp position. A stop is raised, based on the signed comparison; virtual_stop_enc = 1:; X_ENC VIRTUAL_STOP_L; virtual_stop_enc = 0:; XACTUAL VIRTUAL_STOP_L; -2^31…; +(2^31)-1
	TMC5240_VIRTUAL_STOP_L_SHIFT              = 0          // Virtual stop switch based on encoder or ramp position. A stop is raised, based on the signed comparison; virtual_stop_enc = 1:; X_ENC VIRTUAL_STOP_L; virtual_stop_enc = 0:; XACTUAL VIRTUAL_STOP_L; -2^31…; +(2^31)-1
	TMC5240_VIRTUAL_STOP_R_MASK               = 0xFFFFFFFF // VIRTUAL_STOP_R // Virtual Stop Switch based on Encoder. A stop is raised, based on the signed comparison; virtual_stop_enc = 1:; X_ENC >= VIRTUAL_STOP_R; virtual_stop_enc = 0:; XACTUAL >= VIRTUAL_STOP_R; -2^31…; +(2^31)-1
	TMC5240_VIRTUAL_STOP_R_SHIFT              = 0          // Virtual Stop Switch based on Encoder. A stop is raised, based on the signed comparison; virtual_stop_enc = 1:; X_ENC >= VIRTUAL_STOP_R; virtual_stop_enc = 0:; XACTUAL >= VIRTUAL_STOP_R; -2^31…; +(2^31)-1
	TMC5240_ADC_VSUPPLY_MASK                  = 0x1FFF     // ADC_VSUPPLY_AIN // Actual Value of voltage on VS (filtered with low pass filter), update rate: each 2048 clocks [als Plan für alle 3 Werte, den Rest schicken wir den ADC ins idle]
	TMC5240_ADC_VSUPPLY_SHIFT                 = 0          // Actual Value of voltage on VS (filtered with low pass filter), update rate: each 2048 clocks [als Plan für alle 3 Werte, den Rest schicken wir den ADC ins idle]
	TMC5240_ADC_AIN_MASK                      = 0x1FFF0000 // ADC_VSUPPLY_AIN // Value of voltage on ADC_AIN pin in integer
	TMC5240_ADC_AIN_SHIFT                     = 16         // Value of voltage on ADC_AIN pin in integer
	TMC5240_ADC_TEMP_MASK                     = 0x1FFF     // ADC_TEMP // Analog voltage on general purpose input AIN [zu klären ist der Range – ggf. 0-1.25V ist m.E. der native Bereich des ADCs]
	TMC5240_ADC_TEMP_SHIFT                    = 0          // Analog voltage on general purpose input AIN [zu klären ist der Range – ggf. 0-1.25V ist m.E. der native Bereich des ADCs]
	TMC5240_OVERVOLTAGE_VTH_MASK              = 0x1FFF     // OTW_OV_VTH // Overvoltage threshold for output OV. Default: 38V, 36 V equals 1.125 V at ADC inputs
	TMC5240_OVERVOLTAGE_VTH_SHIFT             = 0          // Overvoltage threshold for output OV. Default: 38V, 36 V equals 1.125 V at ADC inputs
	TMC5240_OVERTEMPPREWARNING_VTH_MASK       = 0x1FFF0000 // OTW_OV_VTH // Overtemperature warning threshold register:; ADC_TEMP >= OVERTEMPPREWARNING_VTH; Overtemperatureprewarning will be triggered; (Reset: 0xB92 equals 120°C)
	TMC5240_OVERTEMPPREWARNING_VTH_SHIFT      = 16         // Overtemperature warning threshold register:; ADC_TEMP >= OVERTEMPPREWARNING_VTH; Overtemperatureprewarning will be triggered; (Reset: 0xB92 equals 120°C)
	TMC5240_MSLUT_0_MASK                      = 0xFFFFFFFF // MSLUT_0 // Each bit gives the difference between entry x and entry x+1 when combined with the cor­res­ponding MSLUTSEL W bits:; 0: W=; %00: -1; %01: +0; %10: +1; %11: +2; 1: W=; %00: +0; %01: +1; %10: +2; %11: +3; This is the differential coding for the first quarter of a wave. Start values for CUR_A and CUR_B are stored for MSCNT position 0 in START_SIN and START_SIN90.; ofs31, ofs30, …, ofs01, ofs00; …; ofs255, ofs254, …, ofs225, ofs224; reset default= sine wave table
	TMC5240_MSLUT_0_SHIFT                     = 0          // Each bit gives the difference between entry x and entry x+1 when combined with the cor­res­ponding MSLUTSEL W bits:; 0: W=; %00: -1; %01: +0; %10: +1; %11: +2; 1: W=; %00: +0; %01: +1; %10: +2; %11: +3; This is the differential coding for the first quarter of a wave. Start values for CUR_A and CUR_B are stored for MSCNT position 0 in START_SIN and START_SIN90.; ofs31, ofs30, …, ofs01, ofs00; …; ofs255, ofs254, …, ofs225, ofs224; reset default= sine wave table
	TMC5240_MSLUT_1_MASK                      = 0xFFFFFFFF // MSLUT_1 // Each bit gives the difference between entry x and entry x+1 when combined with the cor­res­ponding MSLUTSEL W bits:; 0: W=; %00: -1; %01: +0; %10: +1; %11: +2; 1: W=; %00: +0; %01: +1; %10: +2; %11: +3; This is the differential coding for the first quarter of a wave. Start values for CUR_A and CUR_B are stored for MSCNT position 0 in START_SIN and START_SIN90.; ofs31, ofs30, …, ofs01, ofs00; …; ofs255, ofs254, …, ofs225, ofs224; reset default= sine wave table; ???????
	TMC5240_MSLUT_1_SHIFT                     = 0          // Each bit gives the difference between entry x and entry x+1 when combined with the cor­res­ponding MSLUTSEL W bits:; 0: W=; %00: -1; %01: +0; %10: +1; %11: +2; 1: W=; %00: +0; %01: +1; %10: +2; %11: +3; This is the differential coding for the first quarter of a wave. Start values for CUR_A and CUR_B are stored for MSCNT position 0 in START_SIN and START_SIN90.; ofs31, ofs30, …, ofs01, ofs00; …; ofs255, ofs254, …, ofs225, ofs224; reset default= sine wave table; ???????
	TMC5240_MSLUT_2_MASK                      = 0xFFFFFFFF // MSLUT_2 // Each bit gives the difference between entry x and entry x+1 when combined with the cor­res­ponding MSLUTSEL W bits:; 0: W=; %00: -1; %01: +0; %10: +1; %11: +2; 1: W=; %00: +0; %01: +1; %10: +2; %11: +3; This is the differential coding for the first quarter of a wave. Start values for CUR_A and CUR_B are stored for MSCNT position 0 in START_SIN and START_SIN90.; ofs31, ofs30, …, ofs01, ofs00; …; ofs255, ofs254, …, ofs225, ofs224; reset default= sine wave table; ???????
	TMC5240_MSLUT_2_SHIFT                     = 0          // Each bit gives the difference between entry x and entry x+1 when combined with the cor­res­ponding MSLUTSEL W bits:; 0: W=; %00: -1; %01: +0; %10: +1; %11: +2; 1: W=; %00: +0; %01: +1; %10: +2; %11: +3; This is the differential coding for the first quarter of a wave. Start values for CUR_A and CUR_B are stored for MSCNT position 0 in START_SIN and START_SIN90.; ofs31, ofs30, …, ofs01, ofs00; …; ofs255, ofs254, …, ofs225, ofs224; reset default= sine wave table; ???????
	TMC5240_MSLUT_3_MASK                      = 0xFFFFFFFF // MSLUT_3 // Each bit gives the difference between entry x and entry x+1 when combined with the cor­res­ponding MSLUTSEL W bits:; 0: W=; %00: -1; %01: +0; %10: +1; %11: +2; 1: W=; %00: +0; %01: +1; %10: +2; %11: +3; This is the differential coding for the first quarter of a wave. Start values for CUR_A and CUR_B are stored for MSCNT position 0 in START_SIN and START_SIN90.; ofs31, ofs30, …, ofs01, ofs00; …; ofs255, ofs254, …, ofs225, ofs224; reset default= sine wave table; ???????
	TMC5240_MSLUT_3_SHIFT                     = 0          // Each bit gives the difference between entry x and entry x+1 when combined with the cor­res­ponding MSLUTSEL W bits:; 0: W=; %00: -1; %01: +0; %10: +1; %11: +2; 1: W=; %00: +0; %01: +1; %10: +2; %11: +3; This is the differential coding for the first quarter of a wave. Start values for CUR_A and CUR_B are stored for MSCNT position 0 in START_SIN and START_SIN90.; ofs31, ofs30, …, ofs01, ofs00; …; ofs255, ofs254, …, ofs225, ofs224; reset default= sine wave table; ???????
	TMC5240_MSLUT_4_MASK                      = 0xFFFFFFFF // MSLUT_4 // Each bit gives the difference between entry x and entry x+1 when combined with the cor­res­ponding MSLUTSEL W bits:; 0: W=; %00: -1; %01: +0; %10: +1; %11: +2; 1: W=; %00: +0; %01: +1; %10: +2; %11: +3; This is the differential coding for the first quarter of a wave. Start values for CUR_A and CUR_B are stored for MSCNT position 0 in START_SIN and START_SIN90.; ofs31, ofs30, …, ofs01, ofs00; …; ofs255, ofs254, …, ofs225, ofs224; reset default= sine wave table; ???????
	TMC5240_MSLUT_4_SHIFT                     = 0          // Each bit gives the difference between entry x and entry x+1 when combined with the cor­res­ponding MSLUTSEL W bits:; 0: W=; %00: -1; %01: +0; %10: +1; %11: +2; 1: W=; %00: +0; %01: +1; %10: +2; %11: +3; This is the differential coding for the first quarter of a wave. Start values for CUR_A and CUR_B are stored for MSCNT position 0 in START_SIN and START_SIN90.; ofs31, ofs30, …, ofs01, ofs00; …; ofs255, ofs254, …, ofs225, ofs224; reset default= sine wave table; ???????
	TMC5240_MSLUT_5_MASK                      = 0xFFFFFFFF // MSLUT_5 // Each bit gives the difference between entry x and entry x+1 when combined with the cor­res­ponding MSLUTSEL W bits:; 0: W=; %00: -1; %01: +0; %10: +1; %11: +2; 1: W=; %00: +0; %01: +1; %10: +2; %11: +3; This is the differential coding for the first quarter of a wave. Start values for CUR_A and CUR_B are stored for MSCNT position 0 in START_SIN and START_SIN90.; ofs31, ofs30, …, ofs01, ofs00; …; ofs255, ofs254, …, ofs225, ofs224; reset default= sine wave table; ???????
	TMC5240_MSLUT_5_SHIFT                     = 0          // Each bit gives the difference between entry x and entry x+1 when combined with the cor­res­ponding MSLUTSEL W bits:; 0: W=; %00: -1; %01: +0; %10: +1; %11: +2; 1: W=; %00: +0; %01: +1; %10: +2; %11: +3; This is the differential coding for the first quarter of a wave. Start values for CUR_A and CUR_B are stored for MSCNT position 0 in START_SIN and START_SIN90.; ofs31, ofs30, …, ofs01, ofs00; …; ofs255, ofs254, …, ofs225, ofs224; reset default= sine wave table; ???????
	TMC5240_MSLUT_6_MASK                      = 0xFFFFFFFF // MSLUT_6 // Each bit gives the difference between entry x and entry x+1 when combined with the cor­res­ponding MSLUTSEL W bits:; 0: W=; %00: -1; %01: +0; %10: +1; %11: +2; 1: W=; %00: +0; %01: +1; %10: +2; %11: +3; This is the differential coding for the first quarter of a wave. Start values for CUR_A and CUR_B are stored for MSCNT position 0 in START_SIN and START_SIN90.; ofs31, ofs30, …, ofs01, ofs00; …; ofs255, ofs254, …, ofs225, ofs224; reset default= sine wave table; ???????
	TMC5240_MSLUT_6_SHIFT                     = 0          // Each bit gives the difference between entry x and entry x+1 when combined with the cor­res­ponding MSLUTSEL W bits:; 0: W=; %00: -1; %01: +0; %10: +1; %11: +2; 1: W=; %00: +0; %01: +1; %10: +2; %11: +3; This is the differential coding for the first quarter of a wave. Start values for CUR_A and CUR_B are stored for MSCNT position 0 in START_SIN and START_SIN90.; ofs31, ofs30, …, ofs01, ofs00; …; ofs255, ofs254, …, ofs225, ofs224; reset default= sine wave table; ???????
	TMC5240_MSLUT_7_MASK                      = 0xFFFFFFFF // MSLUT_7 // Each bit gives the difference between entry x and entry x+1 when combined with the cor­res­ponding MSLUTSEL W bits:; 0: W=; %00: -1; %01: +0; %10: +1; %11: +2; 1: W=; %00: +0; %01: +1; %10: +2; %11: +3; This is the differential coding for the first quarter of a wave. Start values for CUR_A and CUR_B are stored for MSCNT position 0 in START_SIN and START_SIN90.; ofs31, ofs30, …, ofs01, ofs00; …; ofs255, ofs254, …, ofs225, ofs224; reset default= sine wave table; ???????
	TMC5240_MSLUT_7_SHIFT                     = 0          // Each bit gives the difference between entry x and entry x+1 when combined with the cor­res­ponding MSLUTSEL W bits:; 0: W=; %00: -1; %01: +0; %10: +1; %11: +2; 1: W=; %00: +0; %01: +1; %10: +2; %11: +3; This is the differential coding for the first quarter of a wave. Start values for CUR_A and CUR_B are stored for MSCNT position 0 in START_SIN and START_SIN90.; ofs31, ofs30, …, ofs01, ofs00; …; ofs255, ofs254, …, ofs225, ofs224; reset default= sine wave table; ???????
	TMC5240_W0_MASK                           = 0x03       // MSLUTSEL // LUT width select from ofs00 to ofs(X1-1); Width control bit coding W0…W3:; %00:; MSLUT entry 0, 1 select: -1, +0; %01:; MSLUT entry 0, 1 select: +0, +1; %10:; MSLUT entry 0, 1 select: +1, +2; %11:; MSLUT entry 0, 1 select: +2, +3
	TMC5240_W0_SHIFT                          = 0          // LUT width select from ofs00 to ofs(X1-1); Width control bit coding W0…W3:; %00:; MSLUT entry 0, 1 select: -1, +0; %01:; MSLUT entry 0, 1 select: +0, +1; %10:; MSLUT entry 0, 1 select: +1, +2; %11:; MSLUT entry 0, 1 select: +2, +3
	TMC5240_W1_MASK                           = 0x0C       // MSLUTSEL // LUT width select from ofs(X1) to ofs(X2-1); Width control bit coding W0…W3:; %00:; MSLUT entry 0, 1 select: -1, +0; %01:; MSLUT entry 0, 1 select: +0, +1; %10:; MSLUT entry 0, 1 select: +1, +2; %11:; MSLUT entry 0, 1 select: +2, +3
	TMC5240_W1_SHIFT                          = 2          // LUT width select from ofs(X1) to ofs(X2-1); Width control bit coding W0…W3:; %00:; MSLUT entry 0, 1 select: -1, +0; %01:; MSLUT entry 0, 1 select: +0, +1; %10:; MSLUT entry 0, 1 select: +1, +2; %11:; MSLUT entry 0, 1 select: +2, +3
	TMC5240_W2_MASK                           = 0x30       // MSLUTSEL // LUT width select from ofs(X2) to ofs(X3-1); Width control bit coding W0…W3:; %00:; MSLUT entry 0, 1 select: -1, +0; %01:; MSLUT entry 0, 1 select: +0, +1; %10:; MSLUT entry 0, 1 select: +1, +2; %11:; MSLUT entry 0, 1 select: +2, +3
	TMC5240_W2_SHIFT                          = 4          // LUT width select from ofs(X2) to ofs(X3-1); Width control bit coding W0…W3:; %00:; MSLUT entry 0, 1 select: -1, +0; %01:; MSLUT entry 0, 1 select: +0, +1; %10:; MSLUT entry 0, 1 select: +1, +2; %11:; MSLUT entry 0, 1 select: +2, +3
	TMC5240_W3_MASK                           = 0xC0       // MSLUTSEL // LUT width select from ofs(X3) to ofs255; Width control bit coding W0…W3:; %00:; MSLUT entry 0, 1 select: -1, +0; %01:; MSLUT entry 0, 1 select: +0, +1; %10:; MSLUT entry 0, 1 select: +1, +2; %11:; MSLUT entry 0, 1 select: +2, +3
	TMC5240_W3_SHIFT                          = 6          // LUT width select from ofs(X3) to ofs255; Width control bit coding W0…W3:; %00:; MSLUT entry 0, 1 select: -1, +0; %01:; MSLUT entry 0, 1 select: +0, +1; %10:; MSLUT entry 0, 1 select: +1, +2; %11:; MSLUT entry 0, 1 select: +2, +3
	TMC5240_X1_MASK                           = 0xFF00     // MSLUTSEL // LUT segment 1 start; The sine wave look up table can be divided into up to four segments using an individual step width control entry Wx. The segment borders are selected by X1, X2 and X3.; Segment 0 goes from 0 to X1-1.; Segment 1 goes from X1 to X2-1.; Segment 2 goes from X2 to X3-1.; Segment 3 goes from X3 to 255.; For defined response the values shall satisfy:; 0X1X2X3
	TMC5240_X1_SHIFT                          = 8          // LUT segment 1 start; The sine wave look up table can be divided into up to four segments using an individual step width control entry Wx. The segment borders are selected by X1, X2 and X3.; Segment 0 goes from 0 to X1-1.; Segment 1 goes from X1 to X2-1.; Segment 2 goes from X2 to X3-1.; Segment 3 goes from X3 to 255.; For defined response the values shall satisfy:; 0X1X2X3
	TMC5240_X2_MASK                           = 0xFF0000   // MSLUTSEL // LUT segment 1 start; The sine wave look up table can be divided into up to four segments using an individual step width control entry Wx. The segment borders are selected by X1, X2 and X3.; Segment 0 goes from 0 to X1-1.; Segment 1 goes from X1 to X2-1.; Segment 2 goes from X2 to X3-1.; Segment 3 goes from X3 to 255.; For defined response the values shall satisfy:; 0X1X2X3
	TMC5240_X2_SHIFT                          = 16         // LUT segment 1 start; The sine wave look up table can be divided into up to four segments using an individual step width control entry Wx. The segment borders are selected by X1, X2 and X3.; Segment 0 goes from 0 to X1-1.; Segment 1 goes from X1 to X2-1.; Segment 2 goes from X2 to X3-1.; Segment 3 goes from X3 to 255.; For defined response the values shall satisfy:; 0X1X2X3
	TMC5240_X3_MASK                           = 0xFF000000 // MSLUTSEL // LUT segment 1 start; The sine wave look up table can be divided into up to four segments using an individual step width control entry Wx. The segment borders are selected by X1, X2 and X3.; Segment 0 goes from 0 to X1-1.; Segment 1 goes from X1 to X2-1.; Segment 2 goes from X2 to X3-1.; Segment 3 goes from X3 to 255.; For defined response the values shall satisfy:; 0X1X2X3
	TMC5240_X3_SHIFT                          = 24         // LUT segment 1 start; The sine wave look up table can be divided into up to four segments using an individual step width control entry Wx. The segment borders are selected by X1, X2 and X3.; Segment 0 goes from 0 to X1-1.; Segment 1 goes from X1 to X2-1.; Segment 2 goes from X2 to X3-1.; Segment 3 goes from X3 to 255.; For defined response the values shall satisfy:; 0X1X2X3
	TMC5240_START_SIN_MASK                    = 0xFF       // MSLUTSTART // START_SIN gives the absolute value at microstep table entry 0.
	TMC5240_START_SIN_SHIFT                   = 0          // START_SIN gives the absolute value at microstep table entry 0.
	TMC5240_START_SIN90_MASK                  = 0xFF0000   // MSLUTSTART // START_SIN­90 gives the absolute value for cosine wave microstep table entry at MSCNT=0 (table position 256+OFFSET_SIN90).
	TMC5240_START_SIN90_SHIFT                 = 16         // START_SIN­90 gives the absolute value for cosine wave microstep table entry at MSCNT=0 (table position 256+OFFSET_SIN90).
	TMC5240_OFFSET_SIN90_MASK                 = 0xFF000000 // MSLUTSTART // Signed offset for cosine wave +-127 microsteps. Adapt START_SIN90 to match the microstep wave table at position MSCNT=0
	TMC5240_OFFSET_SIN90_SHIFT                = 24         // Signed offset for cosine wave +-127 microsteps. Adapt START_SIN90 to match the microstep wave table at position MSCNT=0
	TMC5240_MSCNT_MASK                        = 0x3FF      // MSCNT // Microstep counter. Indicates actual position in the microstep table for CUR_A. CUR_B uses an offset of 256 (2 phase motor).; Hint: Move to a position where MSCNT is zero before re-initializing MSLUTSTART or MSLUT and MSLUTSEL.
	TMC5240_MSCNT_SHIFT                       = 0          // Microstep counter. Indicates actual position in the microstep table for CUR_A. CUR_B uses an offset of 256 (2 phase motor).; Hint: Move to a position where MSCNT is zero before re-initializing MSLUTSTART or MSLUT and MSLUTSEL.
	TMC5240_CUR_B_MASK                        = 0x1FF      // MSCURACT // Actual microstep current for motor phase B (sine wave) as read from MSLUT (not scaled by current)
	TMC5240_CUR_B_SHIFT                       = 0          // Actual microstep current for motor phase B (sine wave) as read from MSLUT (not scaled by current)
	TMC5240_CUR_A_MASK                        = 0x1FF0000  // MSCURACT // Actual microstep current for motor phase A (co-sine wave) as read from MSLUT (not scaled by current)
	TMC5240_CUR_A_SHIFT                       = 16         // Actual microstep current for motor phase A (co-sine wave) as read from MSLUT (not scaled by current)
	TMC5240_TOFF_MASK                         = 0x0F       // CHOPCONF // TOFF off time and driver enable; Off time setting controls duration of slow decay phase; NCLK= 24 + 32*TOFF; %0000: Driver disable, all bridges off; %0001: 1 – use only with TBL = 2; %0010 … %1111: 2 … 15
	TMC5240_TOFF_SHIFT                        = 0          // TOFF off time and driver enable; Off time setting controls duration of slow decay phase; NCLK= 24 + 32*TOFF; %0000: Driver disable, all bridges off; %0001: 1 – use only with TBL = 2; %0010 … %1111: 2 … 15
	TMC5240_TFD_ALL_MASK                      = 0x70       // CHOPCONF // with chm=0; HSTRT; hysteresis start value added to HEND; %000 … %111:; Add 1, 2, …, 8 to hysteresis low value HEND; (1/512 of this setting adds to current setting); Attention: Effective HEND+HSTRT = 16.; Hint: Hysteresis decrement is done each 16 clocks; with chm=1; TFD [2..0]; fast decay time setting; Fast decay time setting (MSB:; fd3):; %0000 … %1111:; Fast decay time setting TFD with; NCLK= 32*TFD (%0000: slow decay only)
	TMC5240_TFD_ALL_SHIFT                     = 4          // with chm=0; HSTRT; hysteresis start value added to HEND; %000 … %111:; Add 1, 2, …, 8 to hysteresis low value HEND; (1/512 of this setting adds to current setting); Attention: Effective HEND+HSTRT = 16.; Hint: Hysteresis decrement is done each 16 clocks; with chm=1; TFD [2..0]; fast decay time setting; Fast decay time setting (MSB:; fd3):; %0000 … %1111:; Fast decay time setting TFD with; NCLK= 32*TFD (%0000: slow decay only)
	TMC5240_HEND_OFFSET_MASK                  = 0x780      // CHOPCONF // with chm=0; HEND; hysteresis low value; %0000 … %1111:; Hysteresis is -3, -2, -1, 0, 1, …, 12; (1/512 of this setting adds to current setting); This is the hysteresis value which becomes used for the hysteresis chopper.; with chm=1; OFFSET; sine wave offset; %0000 … %1111:; Offset is -3, -2, -1, 0, 1, …, 12; This is the sine wave offset and 1/512 of the value becomes added to the absolute value of each sine wave entry.
	TMC5240_HEND_OFFSET_SHIFT                 = 7          // with chm=0; HEND; hysteresis low value; %0000 … %1111:; Hysteresis is -3, -2, -1, 0, 1, …, 12; (1/512 of this setting adds to current setting); This is the hysteresis value which becomes used for the hysteresis chopper.; with chm=1; OFFSET; sine wave offset; %0000 … %1111:; Offset is -3, -2, -1, 0, 1, …, 12; This is the sine wave offset and 1/512 of the value becomes added to the absolute value of each sine wave entry.
	TMC5240_FD3_MASK                          = 0x800      // CHOPCONF // TFD[3]; with chm=1:; MSB of fast decay time setting TFD
	TMC5240_FD3_SHIFT                         = 11         // TFD[3]; with chm=1:; MSB of fast decay time setting TFD
	TMC5240_DISFDCC_MASK                      = 0x1000     // CHOPCONF // fast decay mode; with chm=1:; disfdcc=1 disables current comparator usage for termi­nation of the fast decay cycle
	TMC5240_DISFDCC_SHIFT                     = 12         // fast decay mode; with chm=1:; disfdcc=1 disables current comparator usage for termi­nation of the fast decay cycle
	TMC5240_CHM_MASK                          = 0x4000     // CHOPCONF // chopper mode
	TMC5240_CHM_SHIFT                         = 14         // chopper mode
	TMC5240_TBL_MASK                          = 0x18000    // CHOPCONF // TBL / blank time select; %00 … %11:; Set comparator blank time to 16, 24, 36 or 54 clocks; Hint: %01 or %10 is recommended for most applications; (Reset Default: OTP %01 or %10)
	TMC5240_TBL_SHIFT                         = 15         // TBL / blank time select; %00 … %11:; Set comparator blank time to 16, 24, 36 or 54 clocks; Hint: %01 or %10 is recommended for most applications; (Reset Default: OTP %01 or %10)
	TMC5240_VHIGHFS_MASK                      = 0x40000    // CHOPCONF // high velocity fullstep selection; This bit enables switching to fullstep, when VHIGH is exceeded. Switching takes place only at 45° position. The fullstep target current uses the current value from the microstep table at the 45° position.
	TMC5240_VHIGHFS_SHIFT                     = 18         // high velocity fullstep selection; This bit enables switching to fullstep, when VHIGH is exceeded. Switching takes place only at 45° position. The fullstep target current uses the current value from the microstep table at the 45° position.
	TMC5240_VHIGHCHM_MASK                     = 0x80000    // CHOPCONF // high velocity chopper mode; This bit enables switching to chm=1 and fd=0, when VHIGH is exceeded. This way, a higher velocity can be achieved. Can be combined with vhighfs=1. If set, the TOFF setting automatically becomes doubled during high velocity operation in order to avoid doubling of the chopper frequency.
	TMC5240_VHIGHCHM_SHIFT                    = 19         // high velocity chopper mode; This bit enables switching to chm=1 and fd=0, when VHIGH is exceeded. This way, a higher velocity can be achieved. Can be combined with vhighfs=1. If set, the TOFF setting automatically becomes doubled during high velocity operation in order to avoid doubling of the chopper frequency.
	TMC5240_TPFD_MASK                         = 0xF00000   // CHOPCONF // passive fast decay time; TPFD allows dampening of motor mid-range resonances.; Passive fast decay time setting controls duration of the fast decay phase inserted after bridge polarity change; NCLK= 128*TPFD; %0000: Disable; %0001 … %1111: 1 … 15
	TMC5240_TPFD_SHIFT                        = 20         // passive fast decay time; TPFD allows dampening of motor mid-range resonances.; Passive fast decay time setting controls duration of the fast decay phase inserted after bridge polarity change; NCLK= 128*TPFD; %0000: Disable; %0001 … %1111: 1 … 15
	TMC5240_MRES_MASK                         = 0xF000000  // CHOPCONF // micro step resolution selection; %0000:; Native 256 microstep setting. Normally use this setting with the internal motion controller.; %0001 … %1000:; 128, 64, 32, 16, 8, 4, 2, FULLSTEP; Reduced microstep resolution esp. for STEP/DIR operation.; The resolution gives the number of microstep entries per sine quarter wave.; The driver automatically uses microstep positions which result in a symmetrical wave, when choosing a lower microstep resolution.; step width=2^MRES [microsteps]
	TMC5240_MRES_SHIFT                        = 24         // micro step resolution selection; %0000:; Native 256 microstep setting. Normally use this setting with the internal motion controller.; %0001 … %1000:; 128, 64, 32, 16, 8, 4, 2, FULLSTEP; Reduced microstep resolution esp. for STEP/DIR operation.; The resolution gives the number of microstep entries per sine quarter wave.; The driver automatically uses microstep positions which result in a symmetrical wave, when choosing a lower microstep resolution.; step width=2^MRES [microsteps]
	TMC5240_INTPOL_MASK                       = 0x10000000 // CHOPCONF // interpolation to 256 microsteps
	TMC5240_INTPOL_SHIFT                      = 28         // interpolation to 256 microsteps
	TMC5240_DEDGE_MASK                        = 0x20000000 // CHOPCONF // enable double edge step pulses
	TMC5240_DEDGE_SHIFT                       = 29         // enable double edge step pulses
	TMC5240_DISS2G_MASK                       = 0x40000000 // CHOPCONF // short to GND protection disable
	TMC5240_DISS2G_SHIFT                      = 30         // short to GND protection disable
	TMC5240_DISS2VS_MASK                      = 0x80000000 // CHOPCONF // short to supply protection disable
	TMC5240_DISS2VS_SHIFT                     = 31         // short to supply protection disable
	TMC5240_SEMIN_MASK                        = 0x0F       // COOLCONF // minimum StallGuard2 value for smart current control and; smart current enable; If the StallGuard2 result falls below SEMIN*32, the motor current becomes increased to reduce motor load angle.; %0000: smart current control CoolStep off; %0001 … %1111: 1 … 15
	TMC5240_SEMIN_SHIFT                       = 0          // minimum StallGuard2 value for smart current control and; smart current enable; If the StallGuard2 result falls below SEMIN*32, the motor current becomes increased to reduce motor load angle.; %0000: smart current control CoolStep off; %0001 … %1111: 1 … 15
	TMC5240_SEUP_MASK                         = 0x60       // COOLCONF // current up step width; Current increment steps per measured StallGuard2 value; %00 … %11: 1, 2, 4, 8
	TMC5240_SEUP_SHIFT                        = 5          // current up step width; Current increment steps per measured StallGuard2 value; %00 … %11: 1, 2, 4, 8
	TMC5240_SEMAX_MASK                        = 0xF00      // COOLCONF // StallGuard2 hysteresis value for smart current control; If the StallGuard2 result is equal to or above (SEMIN+SEMAX+1)*32, the motor current becomes decreased to save energy.; %0000 … %1111: 0 … 15
	TMC5240_SEMAX_SHIFT                       = 8          // StallGuard2 hysteresis value for smart current control; If the StallGuard2 result is equal to or above (SEMIN+SEMAX+1)*32, the motor current becomes decreased to save energy.; %0000 … %1111: 0 … 15
	TMC5240_SEDN_MASK                         = 0x6000     // COOLCONF // current down step speed; %00: For each 32 StallGuard2 values decrease by one; %01: For each 8 StallGuard2 values decrease by one; %10: For each 2 StallGuard2 values decrease by one; %11: For each StallGuard2 value decrease by one
	TMC5240_SEDN_SHIFT                        = 13         // current down step speed; %00: For each 32 StallGuard2 values decrease by one; %01: For each 8 StallGuard2 values decrease by one; %10: For each 2 StallGuard2 values decrease by one; %11: For each StallGuard2 value decrease by one
	TMC5240_SEIMIN_MASK                       = 0x8000     // COOLCONF // minimum current for smart current control
	TMC5240_SEIMIN_SHIFT                      = 15         // minimum current for smart current control
	TMC5240_SGT_MASK                          = 0x7F0000   // COOLCONF // StallGuard2 threshold value; This signed value controls StallGuard2 level for stall output and sets the optimum measurement range for readout. A lower value gives a higher sensitivity. Zero is the starting value working with most motors.; -64 to +63:; A higher value makes StallGuard2 less sensi­tive and requires more torque to indicate a stall.
	TMC5240_SGT_SHIFT                         = 16         // StallGuard2 threshold value; This signed value controls StallGuard2 level for stall output and sets the optimum measurement range for readout. A lower value gives a higher sensitivity. Zero is the starting value working with most motors.; -64 to +63:; A higher value makes StallGuard2 less sensi­tive and requires more torque to indicate a stall.
	TMC5240_SFILT_MASK                        = 0x1000000  // COOLCONF // StallGuard2 and StallGuard4 filter enable
	TMC5240_SFILT_SHIFT                       = 24         // StallGuard2 and StallGuard4 filter enable
	TMC5240_DC_TIME_MASK                      = 0x3FF      // DCCTRL // Upper PWM on time limit for commutation (DC_TIME * 1/fCLK). Set slightly above effective blank time TBL.
	TMC5240_DC_TIME_SHIFT                     = 0          // Upper PWM on time limit for commutation (DC_TIME * 1/fCLK). Set slightly above effective blank time TBL.
	TMC5240_DC_SG_MASK                        = 0xFF0000   // DCCTRL // Max. PWM on time for step loss detection using DcStep StallGuard2 in DcStep mode. (DC_SG * 16/fCLK); Set slightly higher than DC_TIME/16; 0=disable
	TMC5240_DC_SG_SHIFT                       = 16         // Max. PWM on time for step loss detection using DcStep StallGuard2 in DcStep mode. (DC_SG * 16/fCLK); Set slightly higher than DC_TIME/16; 0=disable
	TMC5240_SG_RESULT_MASK                    = 0x3FF      // DRV_STATUS // StallGuard2 result respectively StallGuard4 result (depending on actual chopper mode) resp. PWM on time for coil A in stand still with SpreadCycle for motor temperature detection.; Mechanical load measurement:; The StallGuard2/4 result gives a means to measure mecha­nical motor load. A higher value means lower mecha­nical load. For StallGuard2, a value of 0 signals highest load. With opti­mum SGT setting, this is an indicator for a motor stall. The stall detection compares SG_RESULT to 0 in order to detect a stall. SG_RESULT is used as a base for CoolStep operation, by comparing it to a pro­grammable upper and a lower limit. It is not applicable in StealthChop mode.; StallGuard2 works best with microstep operation or DcStep.; Temperature measurement during SpreadCycle mode:; In standstill, no StallGuard2 result can be obtained. SG_RESULT shows the chopper on-time for motor coil A instead. Move the motor to a determined micro­step position at a certain current setting to get a rough estimation of motor temperature by a reading the chopper on-time. As the motor heats up, its coil resistance rises and the chopper on-time increases. For StallGuard4 specifics, please refer SG4_RESULT.
	TMC5240_SG_RESULT_SHIFT                   = 0          // StallGuard2 result respectively StallGuard4 result (depending on actual chopper mode) resp. PWM on time for coil A in stand still with SpreadCycle for motor temperature detection.; Mechanical load measurement:; The StallGuard2/4 result gives a means to measure mecha­nical motor load. A higher value means lower mecha­nical load. For StallGuard2, a value of 0 signals highest load. With opti­mum SGT setting, this is an indicator for a motor stall. The stall detection compares SG_RESULT to 0 in order to detect a stall. SG_RESULT is used as a base for CoolStep operation, by comparing it to a pro­grammable upper and a lower limit. It is not applicable in StealthChop mode.; StallGuard2 works best with microstep operation or DcStep.; Temperature measurement during SpreadCycle mode:; In standstill, no StallGuard2 result can be obtained. SG_RESULT shows the chopper on-time for motor coil A instead. Move the motor to a determined micro­step position at a certain current setting to get a rough estimation of motor temperature by a reading the chopper on-time. As the motor heats up, its coil resistance rises and the chopper on-time increases. For StallGuard4 specifics, please refer SG4_RESULT.
	TMC5240_S2VSA_MASK                        = 0x1000     // DRV_STATUS // short to supply indicator phase A
	TMC5240_S2VSA_SHIFT                       = 12         // short to supply indicator phase A
	TMC5240_S2VSB_MASK                        = 0x2000     // DRV_STATUS // short to supply indicator phase B
	TMC5240_S2VSB_SHIFT                       = 13         // short to supply indicator phase B
	TMC5240_STEALTH_MASK                      = 0x4000     // DRV_STATUS // StealthChop indicator
	TMC5240_STEALTH_SHIFT                     = 14         // StealthChop indicator
	TMC5240_FSACTIVE_MASK                     = 0x8000     // DRV_STATUS // full step active indicator
	TMC5240_FSACTIVE_SHIFT                    = 15         // full step active indicator
	TMC5240_CS_ACTUAL_MASK                    = 0x1F0000   // DRV_STATUS // actual motor current / smart energy current; Actual current control scaling, for monitoring smart energy current scaling controlled via settings in register COOLCONF, or for monitoring the function of the automatic current scaling
	TMC5240_CS_ACTUAL_SHIFT                   = 16         // actual motor current / smart energy current; Actual current control scaling, for monitoring smart energy current scaling controlled via settings in register COOLCONF, or for monitoring the function of the automatic current scaling
	TMC5240_STALLGUARD_MASK                   = 0x1000000  // DRV_STATUS // StallGuard2/StallGuard4 status
	TMC5240_STALLGUARD_SHIFT                  = 24         // StallGuard2/StallGuard4 status
	TMC5240_OT_MASK                           = 0x2000000  // DRV_STATUS // overtemperature flag
	TMC5240_OT_SHIFT                          = 25         // overtemperature flag
	TMC5240_OTPW_MASK                         = 0x4000000  // DRV_STATUS // overtemperature pre-warning flag
	TMC5240_OTPW_SHIFT                        = 26         // overtemperature pre-warning flag
	TMC5240_S2GA_MASK                         = 0x8000000  // DRV_STATUS // short to ground indicator phase A
	TMC5240_S2GA_SHIFT                        = 27         // short to ground indicator phase A
	TMC5240_S2GB_MASK                         = 0x10000000 // DRV_STATUS // short to ground indicator phase B
	TMC5240_S2GB_SHIFT                        = 28         // short to ground indicator phase B
	TMC5240_OLA_MASK                          = 0x20000000 // DRV_STATUS // open load indicator phase A
	TMC5240_OLA_SHIFT                         = 29         // open load indicator phase A
	TMC5240_OLB_MASK                          = 0x40000000 // DRV_STATUS // open load indicator phase B
	TMC5240_OLB_SHIFT                         = 30         // open load indicator phase B
	TMC5240_STST_MASK                         = 0x80000000 // DRV_STATUS // standstill indicator; This flag indicates motor stand still in each operation mode. This occurs 2^20 clocks after the last step pulse.
	TMC5240_STST_SHIFT                        = 31         // standstill indicator; This flag indicates motor stand still in each operation mode. This occurs 2^20 clocks after the last step pulse.
	TMC5240_PWM_OFS_MASK                      = 0xFF       // PWMCONF // User defined PWM amplitude offset (0-255) related to full motor current (CS_ACTUAL=31) in stand still.; (Reset default=30); Use PWM_OFS as initial value for automatic scaling to speed up the automatic tuning process. To do this, set PWM_OFS to the determined, application specific value, with pwm_autoscale=0. Only afterwards, set pwm_autoscale=1. Enable StealthChop when finished.; PWM_OFS = 0 will disable scaling down motor current below a motor specific lower measurement threshold. This setting should only be used under certain conditions, i.e. when the power supply voltage can vary up and down by a factor of two or more. It prevents the motor going out of regulation, but it also prevents power down below the regulation limit.; PWM_OFS > 0 allows automatic scaling to low PWM duty cycles even below the lower regulation threshold. This allows low (standstill) current settings based on the actual (hold) current scale (register IHOLD_IRUN).
	TMC5240_PWM_OFS_SHIFT                     = 0          // User defined PWM amplitude offset (0-255) related to full motor current (CS_ACTUAL=31) in stand still.; (Reset default=30); Use PWM_OFS as initial value for automatic scaling to speed up the automatic tuning process. To do this, set PWM_OFS to the determined, application specific value, with pwm_autoscale=0. Only afterwards, set pwm_autoscale=1. Enable StealthChop when finished.; PWM_OFS = 0 will disable scaling down motor current below a motor specific lower measurement threshold. This setting should only be used under certain conditions, i.e. when the power supply voltage can vary up and down by a factor of two or more. It prevents the motor going out of regulation, but it also prevents power down below the regulation limit.; PWM_OFS > 0 allows automatic scaling to low PWM duty cycles even below the lower regulation threshold. This allows low (standstill) current settings based on the actual (hold) current scale (register IHOLD_IRUN).
	TMC5240_PWM_GRAD_MASK                     = 0xFF00     // PWMCONF // Velocity dependent gradient for PWM amplitude:; PWM_GRAD * 256 / TSTEP; This value is added to PWM_OFS to compensate for the velocity-dependent motor back-EMF.; Use PWM_GRAD as initial value for automatic scaling to speed up the automatic tuning process. To do this, set PWM_GRAD to the determined, application specific value, with pwm_autoscale=0. Only afterwards, set pwm_autoscale=1. Enable StealthChop when finished.; Hint:; After initial tuning, the required initial value can be read out from PWM_GRAD_AUTO.
	TMC5240_PWM_GRAD_SHIFT                    = 8          // Velocity dependent gradient for PWM amplitude:; PWM_GRAD * 256 / TSTEP; This value is added to PWM_OFS to compensate for the velocity-dependent motor back-EMF.; Use PWM_GRAD as initial value for automatic scaling to speed up the automatic tuning process. To do this, set PWM_GRAD to the determined, application specific value, with pwm_autoscale=0. Only afterwards, set pwm_autoscale=1. Enable StealthChop when finished.; Hint:; After initial tuning, the required initial value can be read out from PWM_GRAD_AUTO.
	TMC5240_PWM_FREQ_MASK                     = 0x30000    // PWMCONF // PWM frequency selection:; %00:; fPWM=2/1024 fCLK (Reset default); %01:; fPWM=2/683 fCLK; %10:; fPWM=2/512 fCLK; %11:; fPWM=2/410 fCLK
	TMC5240_PWM_FREQ_SHIFT                    = 16         // PWM frequency selection:; %00:; fPWM=2/1024 fCLK (Reset default); %01:; fPWM=2/683 fCLK; %10:; fPWM=2/512 fCLK; %11:; fPWM=2/410 fCLK
	TMC5240_PWM_AUTOSCALE_MASK                = 0x40000    // PWMCONF // PWM automatic amplitude scaling
	TMC5240_PWM_AUTOSCALE_SHIFT               = 18         // PWM automatic amplitude scaling
	TMC5240_PWM_AUTOGRAD_MASK                 = 0x80000    // PWMCONF // PWM automatic gradient adaptation
	TMC5240_PWM_AUTOGRAD_SHIFT                = 19         // PWM automatic gradient adaptation
	TMC5240_FREEWHEEL_MASK                    = 0x300000   // PWMCONF // Allows different standstill modes; Stand still option when motor current setting is zero (I_HOLD=0).; %00:; Normal operation; %01:; Freewheeling; %10:; Coil shorted using LS drivers; %11:; Coil shorted using HS drivers
	TMC5240_FREEWHEEL_SHIFT                   = 20         // Allows different standstill modes; Stand still option when motor current setting is zero (I_HOLD=0).; %00:; Normal operation; %01:; Freewheeling; %10:; Coil shorted using LS drivers; %11:; Coil shorted using HS drivers
	TMC5240_PWM_MEAS_SD_ENABLE_MASK           = 0x400000   // PWMCONF // Default=0; 1: Uses slow decay phases on low side to measure the motor current to reduce the lower current limit.
	TMC5240_PWM_MEAS_SD_ENABLE_SHIFT          = 22         // Default=0; 1: Uses slow decay phases on low side to measure the motor current to reduce the lower current limit.
	TMC5240_PWM_DIS_REG_STST_MASK             = 0x800000   // PWMCONF // 1= Disable current regulation when motor is in standstill and current is reduced (less than IRUN). This option eliminates any regulation noise during standstill.
	TMC5240_PWM_DIS_REG_STST_SHIFT            = 23         // 1= Disable current regulation when motor is in standstill and current is reduced (less than IRUN). This option eliminates any regulation noise during standstill.
	TMC5240_PWM_REG_MASK                      = 0xF000000  // PWMCONF // Regulation loop gradient; User defined maximum PWM amplitude change per half wave when using pwm_autoscale=1. (1…15):; 1: 0.5 increments (slowest regulation); 2: 1 increment; 3: 1.5 increments; 4: 2 increments (Reset default)); …; 8: 4 increments; ...; 15: 7.5 increments (fastest regulation)
	TMC5240_PWM_REG_SHIFT                     = 24         // Regulation loop gradient; User defined maximum PWM amplitude change per half wave when using pwm_autoscale=1. (1…15):; 1: 0.5 increments (slowest regulation); 2: 1 increment; 3: 1.5 increments; 4: 2 increments (Reset default)); …; 8: 4 increments; ...; 15: 7.5 increments (fastest regulation)
	TMC5240_PWM_LIM_MASK                      = 0xF0000000 // PWMCONF // PWM automatic scale amplitude limit when switching on; Limit for PWM_SCALE_AUTO when switching back from SpreadCycle to StealthChop. This value defines the upper limit for bits 7 to 4 of the automatic current control when switching back. It can be set to reduce the current jerk during mode change back to StealthChop.; It does not limit PWM_GRAD or PWM_GRAD_AUTO offset.; (Default = 12)
	TMC5240_PWM_LIM_SHIFT                     = 28         // PWM automatic scale amplitude limit when switching on; Limit for PWM_SCALE_AUTO when switching back from SpreadCycle to StealthChop. This value defines the upper limit for bits 7 to 4 of the automatic current control when switching back. It can be set to reduce the current jerk during mode change back to StealthChop.; It does not limit PWM_GRAD or PWM_GRAD_AUTO offset.; (Default = 12)
	TMC5240_PWM_SCALE_SUM_MASK                = 0x3FF      // PWM_SCALE // Bits: 9...0: [0...1023]PWM_SCALE_SUM: Actual PWM duty cycle. This value is used for scaling the values CUR_A and CUR_B read from the sine wave table. 1023: maximum duty cycle. This value is extended by two bits [1,0] for higher precision of duty cycle read out. Bits 9..2 correspond to the 8 bit values in other PWM duty cycle related registers.
	TMC5240_PWM_SCALE_SUM_SHIFT               = 0          // Bits: 9...0: [0...1023]PWM_SCALE_SUM: Actual PWM duty cycle. This value is used for scaling the values CUR_A and CUR_B read from the sine wave table. 1023: maximum duty cycle. This value is extended by two bits [1,0] for higher precision of duty cycle read out. Bits 9..2 correspond to the 8 bit values in other PWM duty cycle related registers.
	TMC5240_PWM_SCALE_AUTO_MASK               = 0x1FF0000  // PWM_SCALE // PWM_SCALE_AUTO
	TMC5240_PWM_SCALE_AUTO_SHIFT              = 16         // PWM_SCALE_AUTO
	TMC5240_PWM_OFS_AUTO_MASK                 = 0xFF       // PWM_AUTO // Automatically determined offset value
	TMC5240_PWM_OFS_AUTO_SHIFT                = 0          // Automatically determined offset value
	TMC5240_PWM_GRAD_AUTO_MASK                = 0xFF0000   // PWM_AUTO // Automatically determined gradient value
	TMC5240_PWM_GRAD_AUTO_SHIFT               = 16         // Automatically determined gradient value
	TMC5240_SG4_THRS_MASK                     = 0xFF       // SG4_THRS // Detection threshold for stall. The StallGuard4 value SG4_RESULT becomes compared to the double of this threshold.; A stall is signaled with; SG_RESULT = SG4_THRS*2
	TMC5240_SG4_THRS_SHIFT                    = 0          // Detection threshold for stall. The StallGuard4 value SG4_RESULT becomes compared to the double of this threshold.; A stall is signaled with; SG_RESULT = SG4_THRS*2
	TMC5240_SG4_FILT_EN_MASK                  = 0x100      // SG4_THRS // 1: enable SG4 filter, 0: disable SG4 filter
	TMC5240_SG4_FILT_EN_SHIFT                 = 8          // 1: enable SG4 filter, 0: disable SG4 filter
	TMC5240_SG_ANGLE_OFFSET_MASK              = 0x200      // SG4_THRS // 1: Automatic phase shift compensation based on StallGuard4, when switching from StealthChop to SpreadCycle controlled via TPWMTHRS
	TMC5240_SG_ANGLE_OFFSET_SHIFT             = 9          // 1: Automatic phase shift compensation based on StallGuard4, when switching from StealthChop to SpreadCycle controlled via TPWMTHRS
	TMC5240_SG4_RESULT_MASK                   = 0x3FF      // SG4_RESULT // StallGuard result for StallGuard4, only.; .; SG4_RESULT becomes updated with each fullstep, independent of TCOOLTHRS and SG4THRS. A higher value signals a lower motor load and more torque headroom.; Intended for StealthChop mode, only. Bits 9 and 0 will always show 0. Scaling to 10 bit is for compatibility to StallGuard2.
	TMC5240_SG4_RESULT_SHIFT                  = 0          // StallGuard result for StallGuard4, only.; .; SG4_RESULT becomes updated with each fullstep, independent of TCOOLTHRS and SG4THRS. A higher value signals a lower motor load and more torque headroom.; Intended for StealthChop mode, only. Bits 9 and 0 will always show 0. Scaling to 10 bit is for compatibility to StallGuard2.
	TMC5240_SG4_IND_0_MASK                    = 0xFF       // SG4_IND // displays SG4 measurement; When SG4_filt_en = 1:; Displays SG4 measurement 0 used as filter input
	TMC5240_SG4_IND_0_SHIFT                   = 0          // displays SG4 measurement; When SG4_filt_en = 1:; Displays SG4 measurement 0 used as filter input
	TMC5240_SG4_IND_1_MASK                    = 0xFF00     // SG4_IND // When SG4_filt_en = 1:; Displays SG4 measurement 1 used as filter input
	TMC5240_SG4_IND_1_SHIFT                   = 8          // When SG4_filt_en = 1:; Displays SG4 measurement 1 used as filter input
	TMC5240_SG4_IND_2_MASK                    = 0xFF0000   // SG4_IND // When SG4_filt_en = 1:; Displays SG4 measurement 2 used as filter input
	TMC5240_SG4_IND_2_SHIFT                   = 16         // When SG4_filt_en = 1:; Displays SG4 measurement 2 used as filter input
	TMC5240_SG4_IND_3_MASK                    = 0xFF000000 // SG4_IND // When SG4_filt_en = 1:; Displays SG4 measurement 3 used as filter input
	TMC5240_SG4_IND_3_SHIFT                   = 24         // When SG4_filt_en = 1:; Displays SG4 measurement 3 used as filter input
)

// Register Fields
var (
	TMC5240_SPI_STATUS_RESET_FLAG_FIELD = RegisterField{
		Mask:     TMC5240_SPI_STATUS_RESET_FLAG_MASK,
		Shift:    TMC5240_SPI_STATUS_RESET_FLAG_SHIFT,
		Register: TMC5240_GSTAT,
		IsSigned: false,
	}
	TMC5240_SPI_STATUS_DRIVER_ERROR_FIELD = RegisterField{
		Mask:     TMC5240_SPI_STATUS_DRIVER_ERROR_MASK,
		Shift:    TMC5240_SPI_STATUS_DRIVER_ERROR_SHIFT,
		Register: TMC5240_GSTAT,
		IsSigned: false,
	}
	TMC5240_SPI_STATUS_SG2_FIELD = RegisterField{
		Mask:     TMC5240_SPI_STATUS_SG2_MASK,
		Shift:    TMC5240_SPI_STATUS_SG2_SHIFT,
		Register: TMC5240_DRVSTATUS,
		IsSigned: false,
	}
	TMC5240_SPI_STATUS_STANDSTILL_FIELD = RegisterField{
		Mask:     TMC5240_SPI_STATUS_STANDSTILL_MASK,
		Shift:    TMC5240_SPI_STATUS_STANDSTILL_SHIFT,
		Register: TMC5240_DRVSTATUS,
		IsSigned: false,
	}
	TMC5240_SPI_STATUS_VELOCITY_REACHED_FIELD = RegisterField{
		Mask:     TMC5240_SPI_STATUS_VELOCITY_REACHED_MASK,
		Shift:    TMC5240_SPI_STATUS_VELOCITY_REACHED_SHIFT,
		Register: TMC5240_RAMPSTAT,
		IsSigned: false,
	}
	TMC5240_SPI_STATUS_POSITION_REACHED_FIELD = RegisterField{
		Mask:     TMC5240_SPI_STATUS_POSITION_REACHED_MASK,
		Shift:    TMC5240_SPI_STATUS_POSITION_REACHED_SHIFT,
		Register: TMC5240_RAMPSTAT,
		IsSigned: false,
	}
	TMC5240_SPI_STATUS_STATUS_STOP_L_FIELD = RegisterField{
		Mask:     TMC5240_SPI_STATUS_STATUS_STOP_L_MASK,
		Shift:    TMC5240_SPI_STATUS_STATUS_STOP_L_SHIFT,
		Register: TMC5240_RAMPSTAT,
		IsSigned: false,
	}
	TMC5240_SPI_STATUS_STATUS_STOP_R_FIELD = RegisterField{
		Mask:     TMC5240_SPI_STATUS_STATUS_STOP_R_MASK,
		Shift:    TMC5240_SPI_STATUS_STATUS_STOP_R_SHIFT,
		Register: TMC5240_RAMPSTAT,
		IsSigned: false,
	}
	TMC5240_FAST_STANDSTILL_FIELD = RegisterField{
		Mask:     TMC5240_FAST_STANDSTILL_MASK,
		Shift:    TMC5240_FAST_STANDSTILL_SHIFT,
		Register: TMC5240_GCONF,
		IsSigned: false,
	}
	TMC5240_EN_PWM_MODE_FIELD = RegisterField{
		Mask:     TMC5240_EN_PWM_MODE_MASK,
		Shift:    TMC5240_EN_PWM_MODE_SHIFT,
		Register: TMC5240_GCONF,
		IsSigned: false,
	}
	TMC5240_MULTISTEP_FILT_FIELD = RegisterField{
		Mask:     TMC5240_MULTISTEP_FILT_MASK,
		Shift:    TMC5240_MULTISTEP_FILT_SHIFT,
		Register: TMC5240_GCONF,
		IsSigned: false,
	}
	TMC5240_SHAFT_FIELD = RegisterField{
		Mask:     TMC5240_SHAFT_MASK,
		Shift:    TMC5240_SHAFT_SHIFT,
		Register: TMC5240_GCONF,
		IsSigned: false,
	}
	TMC5240_DIAG0_ERROR_FIELD = RegisterField{
		Mask:     TMC5240_DIAG0_ERROR_MASK,
		Shift:    TMC5240_DIAG0_ERROR_SHIFT,
		Register: TMC5240_GCONF,
		IsSigned: false,
	}
	TMC5240_DIAG0_OTPW_FIELD = RegisterField{
		Mask:     TMC5240_DIAG0_OTPW_MASK,
		Shift:    TMC5240_DIAG0_OTPW_SHIFT,
		Register: TMC5240_GCONF,
		IsSigned: false,
	}
	TMC5240_DIAG0_STALL_STEP_FIELD = RegisterField{
		Mask:     TMC5240_DIAG0_STALL_STEP_MASK,
		Shift:    TMC5240_DIAG0_STALL_STEP_SHIFT,
		Register: TMC5240_GCONF,
		IsSigned: false,
	}
	TMC5240_DIAG1_STALL_DIR_FIELD = RegisterField{
		Mask:     TMC5240_DIAG1_STALL_DIR_MASK,
		Shift:    TMC5240_DIAG1_STALL_DIR_SHIFT,
		Register: TMC5240_GCONF,
		IsSigned: false,
	}
	TMC5240_DIAG1_INDEX_FIELD = RegisterField{
		Mask:     TMC5240_DIAG1_INDEX_MASK,
		Shift:    TMC5240_DIAG1_INDEX_SHIFT,
		Register: TMC5240_GCONF,
		IsSigned: false,
	}
	TMC5240_DIAG1_ONSTATE_FIELD = RegisterField{
		Mask:     TMC5240_DIAG1_ONSTATE_MASK,
		Shift:    TMC5240_DIAG1_ONSTATE_SHIFT,
		Register: TMC5240_GCONF,
		IsSigned: false,
	}
	TMC5240_DIAG0_INT_PUSHPULL_FIELD = RegisterField{
		Mask:     TMC5240_DIAG0_INT_PUSHPULL_MASK,
		Shift:    TMC5240_DIAG0_INT_PUSHPULL_SHIFT,
		Register: TMC5240_GCONF,
		IsSigned: false,
	}
	TMC5240_DIAG1_POSCOMP_PUSHPULL_FIELD = RegisterField{
		Mask:     TMC5240_DIAG1_POSCOMP_PUSHPULL_MASK,
		Shift:    TMC5240_DIAG1_POSCOMP_PUSHPULL_SHIFT,
		Register: TMC5240_GCONF,
		IsSigned: false,
	}
	TMC5240_SMALL_HYSTERESIS_FIELD = RegisterField{
		Mask:     TMC5240_SMALL_HYSTERESIS_MASK,
		Shift:    TMC5240_SMALL_HYSTERESIS_SHIFT,
		Register: TMC5240_GCONF,
		IsSigned: false,
	}
	TMC5240_STOP_ENABLE_FIELD = RegisterField{
		Mask:     TMC5240_STOP_ENABLE_MASK,
		Shift:    TMC5240_STOP_ENABLE_SHIFT,
		Register: TMC5240_GCONF,
		IsSigned: false,
	}
	TMC5240_DIRECT_MODE_FIELD = RegisterField{
		Mask:     TMC5240_DIRECT_MODE_MASK,
		Shift:    TMC5240_DIRECT_MODE_SHIFT,
		Register: TMC5240_GCONF,
		IsSigned: false,
	}
	TMC5240_LENGTH_STEP_PULSE_FIELD = RegisterField{
		Mask:     TMC5240_LENGTH_STEP_PULSE_MASK,
		Shift:    TMC5240_LENGTH_STEP_PULSE_SHIFT,
		Register: TMC5240_GCONF,
		IsSigned: false,
	}
	TMC5240_RESET_FIELD = RegisterField{
		Mask:     TMC5240_RESET_MASK,
		Shift:    TMC5240_RESET_SHIFT,
		Register: TMC5240_GSTAT,
		IsSigned: false,
	}
	TMC5240_DRV_ERR_FIELD = RegisterField{
		Mask:     TMC5240_DRV_ERR_MASK,
		Shift:    TMC5240_DRV_ERR_SHIFT,
		Register: TMC5240_GSTAT,
		IsSigned: false,
	}
	TMC5240_UV_CP_FIELD = RegisterField{
		Mask:     TMC5240_UV_CP_MASK,
		Shift:    TMC5240_UV_CP_SHIFT,
		Register: TMC5240_GSTAT,
		IsSigned: false,
	}
	TMC5240_REGISTER_RESET_FIELD = RegisterField{
		Mask:     TMC5240_REGISTER_RESET_MASK,
		Shift:    TMC5240_REGISTER_RESET_SHIFT,
		Register: TMC5240_GSTAT,
		IsSigned: false,
	}
	TMC5240_VM_UVLO_FIELD = RegisterField{
		Mask:     TMC5240_VM_UVLO_MASK,
		Shift:    TMC5240_VM_UVLO_SHIFT,
		Register: TMC5240_GSTAT,
		IsSigned: false,
	}
	TMC5240_IFCNT_FIELD = RegisterField{
		Mask:     TMC5240_IFCNT_MASK,
		Shift:    TMC5240_IFCNT_SHIFT,
		Register: TMC5240_IFCNT,
		IsSigned: false,
	}
	TMC5240_SLAVEADDR_FIELD = RegisterField{
		Mask:     TMC5240_SLAVEADDR_MASK,
		Shift:    TMC5240_SLAVEADDR_SHIFT,
		Register: TMC5240_SLAVECONF,
		IsSigned: false,
	}
	TMC5240_SENDDELAY_FIELD = RegisterField{
		Mask:     TMC5240_SENDDELAY_MASK,
		Shift:    TMC5240_SENDDELAY_SHIFT,
		Register: TMC5240_SLAVECONF,
		IsSigned: false,
	}
	TMC5240_REFL_STEP_FIELD = RegisterField{
		Mask:     TMC5240_REFL_STEP_MASK,
		Shift:    TMC5240_REFL_STEP_SHIFT,
		Register: TMC5240_INP_OUT,
		IsSigned: false,
	}
	TMC5240_REFR_DIR_FIELD = RegisterField{
		Mask:     TMC5240_REFR_DIR_MASK,
		Shift:    TMC5240_REFR_DIR_SHIFT,
		Register: TMC5240_INP_OUT,
		IsSigned: false,
	}
	TMC5240_ENCB_CFG4_FIELD = RegisterField{
		Mask:     TMC5240_ENCB_CFG4_MASK,
		Shift:    TMC5240_ENCB_CFG4_SHIFT,
		Register: TMC5240_INP_OUT,
		IsSigned: false,
	}
	TMC5240_ENCA_CFG5_FIELD = RegisterField{
		Mask:     TMC5240_ENCA_CFG5_MASK,
		Shift:    TMC5240_ENCA_CFG5_SHIFT,
		Register: TMC5240_INP_OUT,
		IsSigned: false,
	}
	TMC5240_DRV_ENN_FIELD = RegisterField{
		Mask:     TMC5240_DRV_ENN_MASK,
		Shift:    TMC5240_DRV_ENN_SHIFT,
		Register: TMC5240_INP_OUT,
		IsSigned: false,
	}
	TMC5240_ENCN_CFG6_FIELD = RegisterField{
		Mask:     TMC5240_ENCN_CFG6_MASK,
		Shift:    TMC5240_ENCN_CFG6_SHIFT,
		Register: TMC5240_INP_OUT,
		IsSigned: false,
	}
	TMC5240_UART_EN_FIELD = RegisterField{
		Mask:     TMC5240_UART_EN_MASK,
		Shift:    TMC5240_UART_EN_SHIFT,
		Register: TMC5240_INP_OUT,
		IsSigned: false,
	}
	TMC5240_COMP_A_FIELD = RegisterField{
		Mask:     TMC5240_COMP_A_MASK,
		Shift:    TMC5240_COMP_A_SHIFT,
		Register: TMC5240_INP_OUT,
		IsSigned: false,
	}
	TMC5240_COMP_B_FIELD = RegisterField{
		Mask:     TMC5240_COMP_B_MASK,
		Shift:    TMC5240_COMP_B_SHIFT,
		Register: TMC5240_INP_OUT,
		IsSigned: false,
	}
	TMC5240_COMP_A1_A2_FIELD = RegisterField{
		Mask:     TMC5240_COMP_A1_A2_MASK,
		Shift:    TMC5240_COMP_A1_A2_SHIFT,
		Register: TMC5240_INP_OUT,
		IsSigned: false,
	}
	TMC5240_COMP_B1_B2_FIELD = RegisterField{
		Mask:     TMC5240_COMP_B1_B2_MASK,
		Shift:    TMC5240_COMP_B1_B2_SHIFT,
		Register: TMC5240_INP_OUT,
		IsSigned: false,
	}
	TMC5240_OUTPUT_FIELD = RegisterField{
		Mask:     TMC5240_OUTPUT_MASK,
		Shift:    TMC5240_OUTPUT_SHIFT,
		Register: TMC5240_INP_OUT,
		IsSigned: false,
	}
	TMC5240_EXT_RES_DET_FIELD = RegisterField{
		Mask:     TMC5240_EXT_RES_DET_MASK,
		Shift:    TMC5240_EXT_RES_DET_SHIFT,
		Register: TMC5240_INP_OUT,
		IsSigned: false,
	}
	TMC5240_EXT_CLK_FIELD = RegisterField{
		Mask:     TMC5240_EXT_CLK_MASK,
		Shift:    TMC5240_EXT_CLK_SHIFT,
		Register: TMC5240_INP_OUT,
		IsSigned: false,
	}
	TMC5240_ADC_ERR_FIELD = RegisterField{
		Mask:     TMC5240_ADC_ERR_MASK,
		Shift:    TMC5240_ADC_ERR_SHIFT,
		Register: TMC5240_INP_OUT,
		IsSigned: false,
	}
	TMC5240_SILICON_RV_FIELD = RegisterField{
		Mask:     TMC5240_SILICON_RV_MASK,
		Shift:    TMC5240_SILICON_RV_SHIFT,
		Register: TMC5240_INP_OUT,
		IsSigned: false,
	}
	TMC5240_VERSION_FIELD = RegisterField{
		Mask:     TMC5240_VERSION_MASK,
		Shift:    TMC5240_VERSION_SHIFT,
		Register: TMC5240_INP_OUT,
		IsSigned: false,
	}
	TMC5240_X_COMPARE_FIELD = RegisterField{
		Mask:     TMC5240_X_COMPARE_MASK,
		Shift:    TMC5240_X_COMPARE_SHIFT,
		Register: TMC5240_X_COMPARE,
		IsSigned: true,
	}
	TMC5240_X_COMPARE_REPEAT_FIELD = RegisterField{
		Mask:     TMC5240_X_COMPARE_REPEAT_MASK,
		Shift:    TMC5240_X_COMPARE_REPEAT_SHIFT,
		Register: TMC5240_X_COMPARE, //TODO Check this register
		IsSigned: false,
	}
	TMC5240_CURRENT_RANGE_FIELD = RegisterField{
		Mask:     TMC5240_CURRENT_RANGE_MASK,
		Shift:    TMC5240_CURRENT_RANGE_SHIFT,
		Register: TMC5240_DRV_CONF,
		IsSigned: false,
	}
	TMC5240_SLOPE_CONTROL_FIELD = RegisterField{
		Mask:     TMC5240_SLOPE_CONTROL_MASK,
		Shift:    TMC5240_SLOPE_CONTROL_SHIFT,
		Register: TMC5240_DRV_CONF,
		IsSigned: false,
	}
	TMC5240_BBM_CLKS_FIELD = RegisterField{
		Mask:     TMC5240_BBM_CLKS_MASK,
		Shift:    TMC5240_BBM_CLKS_SHIFT,
		Register: TMC5240_DRV_CONF,
		IsSigned: false,
	}
	TMC5240_RNDTF_FIELD = RegisterField{
		Mask:     TMC5240_RNDTF_MASK,
		Shift:    TMC5240_RNDTF_SHIFT,
		Register: TMC5240_CHOPCONF,
		IsSigned: false,
	}
	TMC5240_GLOBAL_SCALER_FIELD = RegisterField{
		Mask:     TMC5240_GLOBAL_SCALER_MASK,
		Shift:    TMC5240_GLOBAL_SCALER_SHIFT,
		Register: TMC5240_GLOBAL_SCALER,
		IsSigned: false,
	}
	TMC5240_IHOLD_FIELD = RegisterField{
		Mask:     TMC5240_IHOLD_MASK,
		Shift:    TMC5240_IHOLD_SHIFT,
		Register: TMC5240_IHOLD_IRUN,
		IsSigned: false,
	}
	TMC5240_IRUN_FIELD = RegisterField{
		Mask:     TMC5240_IRUN_MASK,
		Shift:    TMC5240_IRUN_SHIFT,
		Register: TMC5240_IHOLD_IRUN,
		IsSigned: false,
	}
	TMC5240_IHOLDDELAY_FIELD = RegisterField{
		Mask:     TMC5240_IHOLDDELAY_MASK,
		Shift:    TMC5240_IHOLDDELAY_SHIFT,
		Register: TMC5240_IHOLD_IRUN,
		IsSigned: false,
	}
	TMC5240_IRUNDELAY_FIELD = RegisterField{
		Mask:     TMC5240_IRUNDELAY_MASK,
		Shift:    TMC5240_IRUNDELAY_SHIFT,
		Register: TMC5240_IHOLD_IRUN,
		IsSigned: false,
	}
	TMC5240_TPOWERDOWN_FIELD = RegisterField{
		Mask:     TMC5240_TPOWERDOWN_MASK,
		Shift:    TMC5240_TPOWERDOWN_SHIFT,
		Register: TMC5240_TPOWERDOWN,
		IsSigned: false,
	}
	TMC5240_TSTEP_FIELD = RegisterField{
		Mask:     TMC5240_TSTEP_MASK,
		Shift:    TMC5240_TSTEP_SHIFT,
		Register: TMC5240_TSTEP,
		IsSigned: false,
	}
	TMC5240_TPWMTHRS_FIELD = RegisterField{
		Mask:     TMC5240_TPWMTHRS_MASK,
		Shift:    TMC5240_TPWMTHRS_SHIFT,
		Register: TMC5240_TPWMTHRS,
		IsSigned: false,
	}
	TMC5240_TCOOLTHRS_FIELD = RegisterField{
		Mask:     TMC5240_TCOOLTHRS_MASK,
		Shift:    TMC5240_TCOOLTHRS_SHIFT,
		Register: TMC5240_TCOOLTHRS,
		IsSigned: false,
	}
	TMC5240_THIGH_FIELD = RegisterField{
		Mask:     TMC5240_THIGH_MASK,
		Shift:    TMC5240_THIGH_SHIFT,
		Register: TMC5240_THIGH,
		IsSigned: false,
	}
	TMC5240_RAMPMODE_FIELD = RegisterField{
		Mask:     TMC5240_RAMPMODE_MASK,
		Shift:    TMC5240_RAMPMODE_SHIFT,
		Register: TMC5240_RAMPMODE,
		IsSigned: false,
	}
	TMC5240_XACTUAL_FIELD = RegisterField{
		Mask:     TMC5240_XACTUAL_MASK,
		Shift:    TMC5240_XACTUAL_SHIFT,
		Register: TMC5240_XACTUAL,
		IsSigned: true,
	}
	TMC5240_VACTUAL_FIELD = RegisterField{
		Mask:     TMC5240_VACTUAL_MASK,
		Shift:    TMC5240_VACTUAL_SHIFT,
		Register: TMC5240_VACTUAL,
		IsSigned: true,
	}
	TMC5240_VSTART_FIELD = RegisterField{
		Mask:     TMC5240_VSTART_MASK,
		Shift:    TMC5240_VSTART_SHIFT,
		Register: TMC5240_VSTART,
		IsSigned: false,
	}
	TMC5240_A1_FIELD = RegisterField{
		Mask:     TMC5240_A1_MASK,
		Shift:    TMC5240_A1_SHIFT,
		Register: TMC5240_A1,
		IsSigned: false,
	}
	TMC5240_V1_FIELD = RegisterField{
		Mask:     TMC5240_V1_MASK,
		Shift:    TMC5240_V1_SHIFT,
		Register: TMC5240_V1,
		IsSigned: false,
	}
	TMC5240_AMAX_FIELD = RegisterField{
		Mask:     TMC5240_AMAX_MASK,
		Shift:    TMC5240_AMAX_SHIFT,
		Register: TMC5240_AMAX,
		IsSigned: false,
	}
	TMC5240_VMAX_FIELD = RegisterField{
		Mask:     TMC5240_VMAX_MASK,
		Shift:    TMC5240_VMAX_SHIFT,
		Register: TMC5240_VMAX,
		IsSigned: false,
	}
	TMC5240_DMAX_FIELD = RegisterField{
		Mask:     TMC5240_DMAX_MASK,
		Shift:    TMC5240_DMAX_SHIFT,
		Register: TMC5240_DMAX,
		IsSigned: false,
	}
	TMC5240_TVMAX_FIELD = RegisterField{
		Mask:     TMC5240_TVMAX_MASK,
		Shift:    TMC5240_TVMAX_SHIFT,
		Register: TMC5240_TVMAX,
		IsSigned: false,
	}
	TMC5240_D1_FIELD = RegisterField{
		Mask:     TMC5240_D1_MASK,
		Shift:    TMC5240_D1_SHIFT,
		Register: TMC5240_D1,
		IsSigned: false,
	}
	TMC5240_VSTOP_FIELD = RegisterField{
		Mask:     TMC5240_VSTOP_MASK,
		Shift:    TMC5240_VSTOP_SHIFT,
		Register: TMC5240_VSTOP,
		IsSigned: false,
	}
	TMC5240_TZEROWAIT_FIELD = RegisterField{
		Mask:     TMC5240_TZEROWAIT_MASK,
		Shift:    TMC5240_TZEROWAIT_SHIFT,
		Register: TMC5240_TZEROWAIT,
		IsSigned: false,
	}
	TMC5240_XTARGET_FIELD = RegisterField{
		Mask:     TMC5240_XTARGET_MASK,
		Shift:    TMC5240_XTARGET_SHIFT,
		Register: TMC5240_XTARGET,
		IsSigned: true,
	}
	TMC5240_V2_FIELD = RegisterField{
		Mask:     TMC5240_V2_MASK,
		Shift:    TMC5240_V2_SHIFT,
		Register: TMC5240_V2,
		IsSigned: false,
	}
	TMC5240_A2_FIELD = RegisterField{
		Mask:     TMC5240_A2_MASK,
		Shift:    TMC5240_A2_SHIFT,
		Register: TMC5240_A2,
		IsSigned: false,
	}
	TMC5240_D2_FIELD = RegisterField{
		Mask:     TMC5240_D2_MASK,
		Shift:    TMC5240_D2_SHIFT,
		Register: TMC5240_D2,
		IsSigned: false,
	}
	TMC5240_AACTUAL_FIELD = RegisterField{
		Mask:     TMC5240_AACTUAL_MASK,
		Shift:    TMC5240_AACTUAL_SHIFT,
		Register: TMC5240_AACTUAL,
		IsSigned: true,
	}
	TMC5240_VDCMIN_FIELD = RegisterField{
		Mask:     TMC5240_VDCMIN_MASK,
		Shift:    TMC5240_VDCMIN_SHIFT,
		Register: TMC5240_VDCMIN,
		IsSigned: false,
	}
	TMC5240_STOP_L_ENABLE_FIELD = RegisterField{
		Mask:     TMC5240_STOP_L_ENABLE_MASK,
		Shift:    TMC5240_STOP_L_ENABLE_SHIFT,
		Register: TMC5240_SWMODE,
		IsSigned: false,
	}
	TMC5240_STOP_R_ENABLE_FIELD = RegisterField{
		Mask:     TMC5240_STOP_R_ENABLE_MASK,
		Shift:    TMC5240_STOP_R_ENABLE_SHIFT,
		Register: TMC5240_SWMODE,
		IsSigned: false,
	}
	TMC5240_POL_STOP_L_FIELD = RegisterField{
		Mask:     TMC5240_POL_STOP_L_MASK,
		Shift:    TMC5240_POL_STOP_L_SHIFT,
		Register: TMC5240_SWMODE,
		IsSigned: false,
	}
	TMC5240_POL_STOP_R_FIELD = RegisterField{
		Mask:     TMC5240_POL_STOP_R_MASK,
		Shift:    TMC5240_POL_STOP_R_SHIFT,
		Register: TMC5240_SWMODE,
		IsSigned: false,
	}
	TMC5240_SWAP_LR_FIELD = RegisterField{
		Mask:     TMC5240_SWAP_LR_MASK,
		Shift:    TMC5240_SWAP_LR_SHIFT,
		Register: TMC5240_SWMODE,
		IsSigned: false,
	}
	TMC5240_LATCH_L_ACTIVE_FIELD = RegisterField{
		Mask:     TMC5240_LATCH_L_ACTIVE_MASK,
		Shift:    TMC5240_LATCH_L_ACTIVE_SHIFT,
		Register: TMC5240_SWMODE,
		IsSigned: false,
	}
	TMC5240_LATCH_L_INACTIVE_FIELD = RegisterField{
		Mask:     TMC5240_LATCH_L_INACTIVE_MASK,
		Shift:    TMC5240_LATCH_L_INACTIVE_SHIFT,
		Register: TMC5240_SWMODE,
		IsSigned: false,
	}
	TMC5240_LATCH_R_ACTIVE_FIELD = RegisterField{
		Mask:     TMC5240_LATCH_R_ACTIVE_MASK,
		Shift:    TMC5240_LATCH_R_ACTIVE_SHIFT,
		Register: TMC5240_SWMODE,
		IsSigned: false,
	}
	TMC5240_LATCH_R_INACTIVE_FIELD = RegisterField{
		Mask:     TMC5240_LATCH_R_INACTIVE_MASK,
		Shift:    TMC5240_LATCH_R_INACTIVE_SHIFT,
		Register: TMC5240_SWMODE,
		IsSigned: false,
	}
	TMC5240_EN_LATCH_ENCODER_FIELD = RegisterField{
		Mask:     TMC5240_EN_LATCH_ENCODER_MASK,
		Shift:    TMC5240_EN_LATCH_ENCODER_SHIFT,
		Register: TMC5240_SWMODE,
		IsSigned: false,
	}
	TMC5240_SG_STOP_FIELD = RegisterField{
		Mask:     TMC5240_SG_STOP_MASK,
		Shift:    TMC5240_SG_STOP_SHIFT,
		Register: TMC5240_SWMODE,
		IsSigned: false,
	}
	TMC5240_EN_SOFTSTOP_FIELD = RegisterField{
		Mask:     TMC5240_EN_SOFTSTOP_MASK,
		Shift:    TMC5240_EN_SOFTSTOP_SHIFT,
		Register: TMC5240_SWMODE,
		IsSigned: false,
	}
	TMC5240_EN_VIRTUAL_STOP_L_FIELD = RegisterField{
		Mask:     TMC5240_EN_VIRTUAL_STOP_L_MASK,
		Shift:    TMC5240_EN_VIRTUAL_STOP_L_SHIFT,
		Register: TMC5240_SWMODE,
		IsSigned: false,
	}
	TMC5240_EN_VIRTUAL_STOP_R_FIELD = RegisterField{
		Mask:     TMC5240_EN_VIRTUAL_STOP_R_MASK,
		Shift:    TMC5240_EN_VIRTUAL_STOP_R_SHIFT,
		Register: TMC5240_SWMODE,
		IsSigned: false,
	}
	TMC5240_VIRTUAL_STOP_ENC_FIELD = RegisterField{
		Mask:     TMC5240_VIRTUAL_STOP_ENC_MASK,
		Shift:    TMC5240_VIRTUAL_STOP_ENC_SHIFT,
		Register: TMC5240_SWMODE,
		IsSigned: false,
	}
	TMC5240_STATUS_STOP_L_FIELD = RegisterField{
		Mask:     TMC5240_STATUS_STOP_L_MASK,
		Shift:    TMC5240_STATUS_STOP_L_SHIFT,
		Register: TMC5240_RAMPSTAT,
		IsSigned: false,
	}
	TMC5240_STATUS_STOP_R_FIELD = RegisterField{
		Mask:     TMC5240_STATUS_STOP_R_MASK,
		Shift:    TMC5240_STATUS_STOP_R_SHIFT,
		Register: TMC5240_RAMPSTAT,
		IsSigned: false,
	}
	TMC5240_STATUS_LATCH_L_FIELD = RegisterField{
		Mask:     TMC5240_STATUS_LATCH_L_MASK,
		Shift:    TMC5240_STATUS_LATCH_L_SHIFT,
		Register: TMC5240_RAMPSTAT,
		IsSigned: false,
	}
	TMC5240_STATUS_LATCH_R_FIELD = RegisterField{
		Mask:     TMC5240_STATUS_LATCH_R_MASK,
		Shift:    TMC5240_STATUS_LATCH_R_SHIFT,
		Register: TMC5240_RAMPSTAT,
		IsSigned: false,
	}
	TMC5240_EVENT_STOP_L_FIELD = RegisterField{
		Mask:     TMC5240_EVENT_STOP_L_MASK,
		Shift:    TMC5240_EVENT_STOP_L_SHIFT,
		Register: TMC5240_RAMPSTAT,
		IsSigned: false,
	}
	TMC5240_EVENT_STOP_R_FIELD = RegisterField{
		Mask:     TMC5240_EVENT_STOP_R_MASK,
		Shift:    TMC5240_EVENT_STOP_R_SHIFT,
		Register: TMC5240_RAMPSTAT,
		IsSigned: false,
	}
	TMC5240_EVENT_STOP_SG_FIELD = RegisterField{
		Mask:     TMC5240_EVENT_STOP_SG_MASK,
		Shift:    TMC5240_EVENT_STOP_SG_SHIFT,
		Register: TMC5240_RAMPSTAT,
		IsSigned: false,
	}
	TMC5240_EVENT_POS_REACHED_FIELD = RegisterField{
		Mask:     TMC5240_EVENT_POS_REACHED_MASK,
		Shift:    TMC5240_EVENT_POS_REACHED_SHIFT,
		Register: TMC5240_RAMPSTAT,
		IsSigned: false,
	}
	TMC5240_VELOCITY_REACHED_FIELD = RegisterField{
		Mask:     TMC5240_VELOCITY_REACHED_MASK,
		Shift:    TMC5240_VELOCITY_REACHED_SHIFT,
		Register: TMC5240_RAMPSTAT,
		IsSigned: false,
	}
	TMC5240_POSITION_REACHED_FIELD = RegisterField{
		Mask:     TMC5240_POSITION_REACHED_MASK,
		Shift:    TMC5240_POSITION_REACHED_SHIFT,
		Register: TMC5240_RAMPSTAT,
		IsSigned: false,
	}
	TMC5240_VZERO_FIELD = RegisterField{
		Mask:     TMC5240_VZERO_MASK,
		Shift:    TMC5240_VZERO_SHIFT,
		Register: TMC5240_RAMPSTAT,
		IsSigned: false,
	}
	TMC5240_T_ZEROWAIT_ACTIVE_FIELD = RegisterField{
		Mask:     TMC5240_T_ZEROWAIT_ACTIVE_MASK,
		Shift:    TMC5240_T_ZEROWAIT_ACTIVE_SHIFT,
		Register: TMC5240_RAMPSTAT,
		IsSigned: false,
	}
	TMC5240_SECOND_MOVE_FIELD = RegisterField{
		Mask:     TMC5240_SECOND_MOVE_MASK,
		Shift:    TMC5240_SECOND_MOVE_SHIFT,
		Register: TMC5240_RAMPSTAT,
		IsSigned: false,
	}
	TMC5240_STATUS_SG_FIELD = RegisterField{
		Mask:     TMC5240_STATUS_SG_MASK,
		Shift:    TMC5240_STATUS_SG_SHIFT,
		Register: TMC5240_RAMPSTAT,
		IsSigned: false,
	}
	TMC5240_STATUS_VIRTUAL_STOP_L_FIELD = RegisterField{
		Mask:     TMC5240_STATUS_VIRTUAL_STOP_L_MASK,
		Shift:    TMC5240_STATUS_VIRTUAL_STOP_L_SHIFT,
		Register: TMC5240_RAMPSTAT,
		IsSigned: false,
	}
	TMC5240_STATUS_VIRTUAL_STOP_R_FIELD = RegisterField{
		Mask:     TMC5240_STATUS_VIRTUAL_STOP_R_MASK,
		Shift:    TMC5240_STATUS_VIRTUAL_STOP_R_SHIFT,
		Register: TMC5240_RAMPSTAT,
		IsSigned: false,
	}
	TMC5240_XLATCH_FIELD = RegisterField{
		Mask:     TMC5240_XLATCH_MASK,
		Shift:    TMC5240_XLATCH_SHIFT,
		Register: TMC5240_XLATCH,
		IsSigned: true,
	}
	TMC5240_POL_A_FIELD = RegisterField{
		Mask:     TMC5240_POL_A_MASK,
		Shift:    TMC5240_POL_A_SHIFT,
		Register: TMC5240_ENCMODE,
		IsSigned: false,
	}
	TMC5240_POL_B_FIELD = RegisterField{
		Mask:     TMC5240_POL_B_MASK,
		Shift:    TMC5240_POL_B_SHIFT,
		Register: TMC5240_ENCMODE,
		IsSigned: false,
	}
	TMC5240_POL_N_FIELD = RegisterField{
		Mask:     TMC5240_POL_N_MASK,
		Shift:    TMC5240_POL_N_SHIFT,
		Register: TMC5240_ENCMODE,
		IsSigned: false,
	}
	TMC5240_IGNORE_AB_FIELD = RegisterField{
		Mask:     TMC5240_IGNORE_AB_MASK,
		Shift:    TMC5240_IGNORE_AB_SHIFT,
		Register: TMC5240_ENCMODE,
		IsSigned: false,
	}
	TMC5240_CLR_CONT_FIELD = RegisterField{
		Mask:     TMC5240_CLR_CONT_MASK,
		Shift:    TMC5240_CLR_CONT_SHIFT,
		Register: TMC5240_ENCMODE,
		IsSigned: false,
	}
	TMC5240_CLR_ONCE_FIELD = RegisterField{
		Mask:     TMC5240_CLR_ONCE_MASK,
		Shift:    TMC5240_CLR_ONCE_SHIFT,
		Register: TMC5240_ENCMODE,
		IsSigned: false,
	}
	TMC5240_POS_NEG_EDGE_FIELD = RegisterField{
		Mask:     TMC5240_POS_NEG_EDGE_MASK,
		Shift:    TMC5240_POS_NEG_EDGE_SHIFT,
		Register: TMC5240_ENCMODE,
		IsSigned: false,
	}
	TMC5240_CLR_ENC_X_FIELD = RegisterField{
		Mask:     TMC5240_CLR_ENC_X_MASK,
		Shift:    TMC5240_CLR_ENC_X_SHIFT,
		Register: TMC5240_ENCMODE,
		IsSigned: false,
	}
	TMC5240_LATCH_X_ACT_FIELD = RegisterField{
		Mask:     TMC5240_LATCH_X_ACT_MASK,
		Shift:    TMC5240_LATCH_X_ACT_SHIFT,
		Register: TMC5240_ENCMODE,
		IsSigned: false,
	}
	TMC5240_ENC_SEL_DECIMAL_FIELD = RegisterField{
		Mask:     TMC5240_ENC_SEL_DECIMAL_MASK,
		Shift:    TMC5240_ENC_SEL_DECIMAL_SHIFT,
		Register: TMC5240_ENCMODE,
		IsSigned: false,
	}
	TMC5240_X_ENC_FIELD = RegisterField{
		Mask:     TMC5240_X_ENC_MASK,
		Shift:    TMC5240_X_ENC_SHIFT,
		Register: TMC5240_XENC,
		IsSigned: true,
	}
	TMC5240_ENC_CONST_FIELD = RegisterField{
		Mask:     TMC5240_ENC_CONST_MASK,
		Shift:    TMC5240_ENC_CONST_SHIFT,
		Register: TMC5240_ENC_CONST,
		IsSigned: true,
	}
	TMC5240_N_EVENT_FIELD = RegisterField{
		Mask:     TMC5240_N_EVENT_MASK,
		Shift:    TMC5240_N_EVENT_SHIFT,
		Register: TMC5240_ENC_STATUS,
		IsSigned: false,
	}
	TMC5240_DEVIATION_WARN_FIELD = RegisterField{
		Mask:     TMC5240_DEVIATION_WARN_MASK,
		Shift:    TMC5240_DEVIATION_WARN_SHIFT,
		Register: TMC5240_ENC_STATUS,
		IsSigned: false,
	}
	TMC5240_ENC_LATCH_FIELD = RegisterField{
		Mask:     TMC5240_ENC_LATCH_MASK,
		Shift:    TMC5240_ENC_LATCH_SHIFT,
		Register: TMC5240_ENC_LATCH,
		IsSigned: false,
	}
	TMC5240_ENC_DEVIATION_FIELD = RegisterField{
		Mask:     TMC5240_ENC_DEVIATION_MASK,
		Shift:    TMC5240_ENC_DEVIATION_SHIFT,
		Register: TMC5240_ENC_DEVIATION,
		IsSigned: false,
	}
	TMC5240_VIRTUAL_STOP_L_FIELD = RegisterField{
		Mask:     TMC5240_VIRTUAL_STOP_L_MASK,
		Shift:    TMC5240_VIRTUAL_STOP_L_SHIFT,
		Register: TMC5240_VIRTUAL_STOP_L,
		IsSigned: true,
	}
	TMC5240_VIRTUAL_STOP_R_FIELD = RegisterField{
		Mask:     TMC5240_VIRTUAL_STOP_R_MASK,
		Shift:    TMC5240_VIRTUAL_STOP_R_SHIFT,
		Register: TMC5240_VIRTUAL_STOP_R,
		IsSigned: true,
	}
	TMC5240_ADC_VSUPPLY_FIELD = RegisterField{
		Mask:     TMC5240_ADC_VSUPPLY_MASK,
		Shift:    TMC5240_ADC_VSUPPLY_SHIFT,
		Register: TMC5240_ADC_VSUPPLY_AIN,
		IsSigned: true,
	}
	TMC5240_ADC_AIN_FIELD = RegisterField{
		Mask:     TMC5240_ADC_AIN_MASK,
		Shift:    TMC5240_ADC_AIN_SHIFT,
		Register: TMC5240_ADC_VSUPPLY_AIN,
		IsSigned: true,
	}
	TMC5240_ADC_TEMP_FIELD = RegisterField{
		Mask:     TMC5240_ADC_TEMP_MASK,
		Shift:    TMC5240_ADC_TEMP_SHIFT,
		Register: TMC5240_ADC_TEMP,
		IsSigned: true,
	}
	TMC5240_OVERVOLTAGE_VTH_FIELD = RegisterField{
		Mask:     TMC5240_OVERVOLTAGE_VTH_MASK,
		Shift:    TMC5240_OVERVOLTAGE_VTH_SHIFT,
		Register: TMC5240_OTW_OV_VTH,
		IsSigned: false,
	}
	TMC5240_OVERTEMPPREWARNING_VTH_FIELD = RegisterField{
		Mask:     TMC5240_OVERTEMPPREWARNING_VTH_MASK,
		Shift:    TMC5240_OVERTEMPPREWARNING_VTH_SHIFT,
		Register: TMC5240_OTW_OV_VTH,
		IsSigned: false,
	}
	TMC5240_MSLUT_0_FIELD = RegisterField{
		Mask:     TMC5240_MSLUT_0_MASK,
		Shift:    TMC5240_MSLUT_0_SHIFT,
		Register: TMC5240_MSLUT0,
		IsSigned: false,
	}
	TMC5240_MSLUT_1_FIELD = RegisterField{
		Mask:     TMC5240_MSLUT_1_MASK,
		Shift:    TMC5240_MSLUT_1_SHIFT,
		Register: TMC5240_MSLUT1,
		IsSigned: false,
	}
	TMC5240_MSLUT_2_FIELD = RegisterField{
		Mask:     TMC5240_MSLUT_2_MASK,
		Shift:    TMC5240_MSLUT_2_SHIFT,
		Register: TMC5240_MSLUT2,
		IsSigned: false,
	}
	TMC5240_MSLUT_3_FIELD = RegisterField{
		Mask:     TMC5240_MSLUT_3_MASK,
		Shift:    TMC5240_MSLUT_3_SHIFT,
		Register: TMC5240_MSLUT3,
		IsSigned: false,
	}
	TMC5240_MSLUT_4_FIELD = RegisterField{
		Mask:     TMC5240_MSLUT_4_MASK,
		Shift:    TMC5240_MSLUT_4_SHIFT,
		Register: TMC5240_MSLUT4,
		IsSigned: false,
	}
	TMC5240_MSLUT_5_FIELD = RegisterField{
		Mask:     TMC5240_MSLUT_5_MASK,
		Shift:    TMC5240_MSLUT_5_SHIFT,
		Register: TMC5240_MSLUT5,
		IsSigned: false,
	}
	TMC5240_MSLUT_6_FIELD = RegisterField{
		Mask:     TMC5240_MSLUT_6_MASK,
		Shift:    TMC5240_MSLUT_6_SHIFT,
		Register: TMC5240_MSLUT6,
		IsSigned: false,
	}
	TMC5240_MSLUT_7_FIELD = RegisterField{
		Mask:     TMC5240_MSLUT_7_MASK,
		Shift:    TMC5240_MSLUT_7_SHIFT,
		Register: TMC5240_MSLUT7,
		IsSigned: false,
	}
	TMC5240_W0_FIELD = RegisterField{
		Mask:     TMC5240_W0_MASK,
		Shift:    TMC5240_W0_SHIFT,
		Register: TMC5240_MSLUTSEL,
		IsSigned: false,
	}
	TMC5240_W1_FIELD = RegisterField{
		Mask:     TMC5240_W1_MASK,
		Shift:    TMC5240_W1_SHIFT,
		Register: TMC5240_MSLUTSEL,
		IsSigned: false,
	}
	TMC5240_W2_FIELD = RegisterField{
		Mask:     TMC5240_W2_MASK,
		Shift:    TMC5240_W2_SHIFT,
		Register: TMC5240_MSLUTSEL,
		IsSigned: false,
	}
	TMC5240_W3_FIELD = RegisterField{
		Mask:     TMC5240_W3_MASK,
		Shift:    TMC5240_W3_SHIFT,
		Register: TMC5240_MSLUTSEL,
		IsSigned: false,
	}
	TMC5240_X1_FIELD = RegisterField{
		Mask:     TMC5240_X1_MASK,
		Shift:    TMC5240_X1_SHIFT,
		Register: TMC5240_MSLUTSEL,
		IsSigned: false,
	}
	TMC5240_X2_FIELD = RegisterField{
		Mask:     TMC5240_X2_MASK,
		Shift:    TMC5240_X2_SHIFT,
		Register: TMC5240_MSLUTSEL,
		IsSigned: false,
	}
	TMC5240_X3_FIELD = RegisterField{
		Mask:     TMC5240_X3_MASK,
		Shift:    TMC5240_X3_SHIFT,
		Register: TMC5240_MSLUTSEL,
		IsSigned: false,
	}
	TMC5240_START_SIN_FIELD = RegisterField{
		Mask:     TMC5240_START_SIN_MASK,
		Shift:    TMC5240_START_SIN_SHIFT,
		Register: TMC5240_MSLUTSTART,
		IsSigned: false,
	}
	TMC5240_START_SIN90_FIELD = RegisterField{
		Mask:     TMC5240_START_SIN90_MASK,
		Shift:    TMC5240_START_SIN90_SHIFT,
		Register: TMC5240_MSLUTSTART,
		IsSigned: false,
	}
	TMC5240_OFFSET_SIN90_FIELD = RegisterField{
		Mask:     TMC5240_OFFSET_SIN90_MASK,
		Shift:    TMC5240_OFFSET_SIN90_SHIFT,
		Register: TMC5240_MSLUTSTART,
		IsSigned: false,
	}
	TMC5240_MSCNT_FIELD = RegisterField{
		Mask:     TMC5240_MSCNT_MASK,
		Shift:    TMC5240_MSCNT_SHIFT,
		Register: TMC5240_MSCNT,
		IsSigned: false,
	}
	TMC5240_CUR_B_FIELD = RegisterField{
		Mask:     TMC5240_CUR_B_MASK,
		Shift:    TMC5240_CUR_B_SHIFT,
		Register: TMC5240_MSCURACT,
		IsSigned: true,
	}
	TMC5240_CUR_A_FIELD = RegisterField{
		Mask:     TMC5240_CUR_A_MASK,
		Shift:    TMC5240_CUR_A_SHIFT,
		Register: TMC5240_MSCURACT,
		IsSigned: true,
	}
	TMC5240_TOFF_FIELD = RegisterField{
		Mask:     TMC5240_TOFF_MASK,
		Shift:    TMC5240_TOFF_SHIFT,
		Register: TMC5240_CHOPCONF,
		IsSigned: false,
	}
	TMC5240_TFD_ALL_FIELD = RegisterField{
		Mask:     TMC5240_TFD_ALL_MASK,
		Shift:    TMC5240_TFD_ALL_SHIFT,
		Register: TMC5240_CHOPCONF,
		IsSigned: false,
	}
	TMC5240_HEND_OFFSET_FIELD = RegisterField{
		Mask:     TMC5240_HEND_OFFSET_MASK,
		Shift:    TMC5240_HEND_OFFSET_SHIFT,
		Register: TMC5240_CHOPCONF,
		IsSigned: false,
	}
	TMC5240_FD3_FIELD = RegisterField{
		Mask:     TMC5240_FD3_MASK,
		Shift:    TMC5240_FD3_SHIFT,
		Register: TMC5240_CHOPCONF,
		IsSigned: false,
	}
	TMC5240_DISFDCC_FIELD = RegisterField{
		Mask:     TMC5240_DISFDCC_MASK,
		Shift:    TMC5240_DISFDCC_SHIFT,
		Register: TMC5240_CHOPCONF,
		IsSigned: false,
	}
	TMC5240_CHM_FIELD = RegisterField{
		Mask:     TMC5240_CHM_MASK,
		Shift:    TMC5240_CHM_SHIFT,
		Register: TMC5240_CHOPCONF,
		IsSigned: false,
	}
	TMC5240_TBL_FIELD = RegisterField{
		Mask:     TMC5240_TBL_MASK,
		Shift:    TMC5240_TBL_SHIFT,
		Register: TMC5240_CHOPCONF,
		IsSigned: false,
	}
	TMC5240_VHIGHFS_FIELD = RegisterField{
		Mask:     TMC5240_VHIGHFS_MASK,
		Shift:    TMC5240_VHIGHFS_SHIFT,
		Register: TMC5240_CHOPCONF,
		IsSigned: false,
	}
	TMC5240_VHIGHCHM_FIELD = RegisterField{
		Mask:     TMC5240_VHIGHCHM_MASK,
		Shift:    TMC5240_VHIGHCHM_SHIFT,
		Register: TMC5240_CHOPCONF,
		IsSigned: false,
	}
	TMC5240_TPFD_FIELD = RegisterField{
		Mask:     TMC5240_TPFD_MASK,
		Shift:    TMC5240_TPFD_SHIFT,
		Register: TMC5240_CHOPCONF,
		IsSigned: false,
	}
	TMC5240_MRES_FIELD = RegisterField{
		Mask:     TMC5240_MRES_MASK,
		Shift:    TMC5240_MRES_SHIFT,
		Register: TMC5240_CHOPCONF,
		IsSigned: false,
	}
	TMC5240_INTPOL_FIELD = RegisterField{
		Mask:     TMC5240_INTPOL_MASK,
		Shift:    TMC5240_INTPOL_SHIFT,
		Register: TMC5240_CHOPCONF,
		IsSigned: false,
	}
	TMC5240_DEDGE_FIELD = RegisterField{
		Mask:     TMC5240_DEDGE_MASK,
		Shift:    TMC5240_DEDGE_SHIFT,
		Register: TMC5240_CHOPCONF,
		IsSigned: false,
	}
	TMC5240_DISS2G_FIELD = RegisterField{
		Mask:     TMC5240_DISS2G_MASK,
		Shift:    TMC5240_DISS2G_SHIFT,
		Register: TMC5240_CHOPCONF,
		IsSigned: false,
	}
	TMC5240_DISS2VS_FIELD = RegisterField{
		Mask:     TMC5240_DISS2VS_MASK,
		Shift:    TMC5240_DISS2VS_SHIFT,
		Register: TMC5240_CHOPCONF,
		IsSigned: false,
	}
	TMC5240_SEMIN_FIELD = RegisterField{
		Mask:     TMC5240_SEMIN_MASK,
		Shift:    TMC5240_SEMIN_SHIFT,
		Register: TMC5240_COOLCONF,
		IsSigned: false,
	}
	TMC5240_SEUP_FIELD = RegisterField{
		Mask:     TMC5240_SEUP_MASK,
		Shift:    TMC5240_SEUP_SHIFT,
		Register: TMC5240_COOLCONF,
		IsSigned: false,
	}
	TMC5240_SEMAX_FIELD = RegisterField{
		Mask:     TMC5240_SEMAX_MASK,
		Shift:    TMC5240_SEMAX_SHIFT,
		Register: TMC5240_COOLCONF,
		IsSigned: false,
	}
	TMC5240_SEDN_FIELD = RegisterField{
		Mask:     TMC5240_SEDN_MASK,
		Shift:    TMC5240_SEDN_SHIFT,
		Register: TMC5240_COOLCONF,
		IsSigned: false,
	}
	TMC5240_SEIMIN_FIELD = RegisterField{
		Mask:     TMC5240_SEIMIN_MASK,
		Shift:    TMC5240_SEIMIN_SHIFT,
		Register: TMC5240_COOLCONF,
		IsSigned: false,
	}
	TMC5240_SGT_FIELD = RegisterField{
		Mask:     TMC5240_SGT_MASK,
		Shift:    TMC5240_SGT_SHIFT,
		Register: TMC5240_COOLCONF,
		IsSigned: false,
	}
	TMC5240_SFILT_FIELD = RegisterField{
		Mask:     TMC5240_SFILT_MASK,
		Shift:    TMC5240_SFILT_SHIFT,
		Register: TMC5240_COOLCONF,
		IsSigned: false,
	}
	TMC5240_DC_TIME_FIELD = RegisterField{
		Mask:     TMC5240_DC_TIME_MASK,
		Shift:    TMC5240_DC_TIME_SHIFT,
		Register: TMC5240_DCCTRL,
		IsSigned: false,
	}
	TMC5240_DC_SG_FIELD = RegisterField{
		Mask:     TMC5240_DC_SG_MASK,
		Shift:    TMC5240_DC_SG_SHIFT,
		Register: TMC5240_DCCTRL,
		IsSigned: false,
	}
	TMC5240_SG_RESULT_FIELD = RegisterField{
		Mask:     TMC5240_SG_RESULT_MASK,
		Shift:    TMC5240_SG_RESULT_SHIFT,
		Register: TMC5240_DRVSTATUS,
		IsSigned: false,
	}
	TMC5240_S2VSA_FIELD = RegisterField{
		Mask:     TMC5240_S2VSA_MASK,
		Shift:    TMC5240_S2VSA_SHIFT,
		Register: TMC5240_DRVSTATUS,
		IsSigned: false,
	}
	TMC5240_S2VSB_FIELD = RegisterField{
		Mask:     TMC5240_S2VSB_MASK,
		Shift:    TMC5240_S2VSB_SHIFT,
		Register: TMC5240_DRVSTATUS,
		IsSigned: false,
	}
	TMC5240_STEALTH_FIELD = RegisterField{
		Mask:     TMC5240_STEALTH_MASK,
		Shift:    TMC5240_STEALTH_SHIFT,
		Register: TMC5240_DRVSTATUS,
		IsSigned: false,
	}
	TMC5240_FSACTIVE_FIELD = RegisterField{
		Mask:     TMC5240_FSACTIVE_MASK,
		Shift:    TMC5240_FSACTIVE_SHIFT,
		Register: TMC5240_DRVSTATUS,
		IsSigned: false,
	}
	TMC5240_CS_ACTUAL_FIELD = RegisterField{
		Mask:     TMC5240_CS_ACTUAL_MASK,
		Shift:    TMC5240_CS_ACTUAL_SHIFT,
		Register: TMC5240_DRVSTATUS,
		IsSigned: false,
	}
	TMC5240_STALLGUARD_FIELD = RegisterField{
		Mask:     TMC5240_STALLGUARD_MASK,
		Shift:    TMC5240_STALLGUARD_SHIFT,
		Register: TMC5240_DRVSTATUS,
		IsSigned: false,
	}
	TMC5240_OT_FIELD = RegisterField{
		Mask:     TMC5240_OT_MASK,
		Shift:    TMC5240_OT_SHIFT,
		Register: TMC5240_DRVSTATUS,
		IsSigned: false,
	}
	TMC5240_OTPW_FIELD = RegisterField{
		Mask:     TMC5240_OTPW_MASK,
		Shift:    TMC5240_OTPW_SHIFT,
		Register: TMC5240_DRVSTATUS,
		IsSigned: false,
	}
	TMC5240_S2GA_FIELD = RegisterField{
		Mask:     TMC5240_S2GA_MASK,
		Shift:    TMC5240_S2GA_SHIFT,
		Register: TMC5240_DRVSTATUS,
		IsSigned: false,
	}
	TMC5240_S2GB_FIELD = RegisterField{
		Mask:     TMC5240_S2GB_MASK,
		Shift:    TMC5240_S2GB_SHIFT,
		Register: TMC5240_DRVSTATUS,
		IsSigned: false,
	}
	TMC5240_OLA_FIELD = RegisterField{
		Mask:     TMC5240_OLA_MASK,
		Shift:    TMC5240_OLA_SHIFT,
		Register: TMC5240_DRVSTATUS,
		IsSigned: false,
	}
	TMC5240_OLB_FIELD = RegisterField{
		Mask:     TMC5240_OLB_MASK,
		Shift:    TMC5240_OLB_SHIFT,
		Register: TMC5240_DRVSTATUS,
		IsSigned: false,
	}
	TMC5240_STST_FIELD = RegisterField{
		Mask:     TMC5240_STST_MASK,
		Shift:    TMC5240_STST_SHIFT,
		Register: TMC5240_DRVSTATUS,
		IsSigned: false,
	}
	TMC5240_PWM_OFS_FIELD = RegisterField{
		Mask:     TMC5240_PWM_OFS_MASK,
		Shift:    TMC5240_PWM_OFS_SHIFT,
		Register: TMC5240_PWMCONF,
		IsSigned: false,
	}
	TMC5240_PWM_GRAD_FIELD = RegisterField{
		Mask:     TMC5240_PWM_GRAD_MASK,
		Shift:    TMC5240_PWM_GRAD_SHIFT,
		Register: TMC5240_PWMCONF,
		IsSigned: false,
	}
	TMC5240_PWM_FREQ_FIELD = RegisterField{
		Mask:     TMC5240_PWM_FREQ_MASK,
		Shift:    TMC5240_PWM_FREQ_SHIFT,
		Register: TMC5240_PWMCONF,
		IsSigned: false,
	}
	TMC5240_PWM_AUTOSCALE_FIELD = RegisterField{
		Mask:     TMC5240_PWM_AUTOSCALE_MASK,
		Shift:    TMC5240_PWM_AUTOSCALE_SHIFT,
		Register: TMC5240_PWMCONF,
		IsSigned: false,
	}
	TMC5240_PWM_AUTOGRAD_FIELD = RegisterField{
		Mask:     TMC5240_PWM_AUTOGRAD_MASK,
		Shift:    TMC5240_PWM_AUTOGRAD_SHIFT,
		Register: TMC5240_PWMCONF,
		IsSigned: false,
	}
	TMC5240_FREEWHEEL_FIELD = RegisterField{
		Mask:     TMC5240_FREEWHEEL_MASK,
		Shift:    TMC5240_FREEWHEEL_SHIFT,
		Register: TMC5240_PWMCONF,
		IsSigned: false,
	}
	TMC5240_PWM_MEAS_SD_ENABLE_FIELD = RegisterField{
		Mask:     TMC5240_PWM_MEAS_SD_ENABLE_MASK,
		Shift:    TMC5240_PWM_MEAS_SD_ENABLE_SHIFT,
		Register: TMC5240_PWMCONF,
		IsSigned: false,
	}
	TMC5240_PWM_DIS_REG_STST_FIELD = RegisterField{
		Mask:     TMC5240_PWM_DIS_REG_STST_MASK,
		Shift:    TMC5240_PWM_DIS_REG_STST_SHIFT,
		Register: TMC5240_PWMCONF,
		IsSigned: false,
	}
	TMC5240_PWM_REG_FIELD = RegisterField{
		Mask:     TMC5240_PWM_REG_MASK,
		Shift:    TMC5240_PWM_REG_SHIFT,
		Register: TMC5240_PWMCONF,
		IsSigned: false,
	}
	TMC5240_PWM_LIM_FIELD = RegisterField{
		Mask:     TMC5240_PWM_LIM_MASK,
		Shift:    TMC5240_PWM_LIM_SHIFT,
		Register: TMC5240_PWMCONF,
		IsSigned: false,
	}
	TMC5240_PWM_SCALE_SUM_FIELD = RegisterField{
		Mask:     TMC5240_PWM_SCALE_SUM_MASK,
		Shift:    TMC5240_PWM_SCALE_SUM_SHIFT,
		Register: TMC5240_PWMSCALE,
		IsSigned: false,
	}
	TMC5240_PWM_SCALE_AUTO_FIELD = RegisterField{
		Mask:     TMC5240_PWM_SCALE_AUTO_MASK,
		Shift:    TMC5240_PWM_SCALE_AUTO_SHIFT,
		Register: TMC5240_PWMSCALE,
		IsSigned: false,
	}
	TMC5240_PWM_OFS_AUTO_FIELD = RegisterField{
		Mask:     TMC5240_PWM_OFS_AUTO_MASK,
		Shift:    TMC5240_PWM_OFS_AUTO_SHIFT,
		Register: TMC5240_PWM_AUTO,
		IsSigned: false,
	}
	TMC5240_PWM_GRAD_AUTO_FIELD = RegisterField{
		Mask:     TMC5240_PWM_GRAD_AUTO_MASK,
		Shift:    TMC5240_PWM_GRAD_AUTO_SHIFT,
		Register: TMC5240_PWM_AUTO,
		IsSigned: false,
	}
	TMC5240_SG4_THRS_FIELD = RegisterField{
		Mask:     TMC5240_SG4_THRS_MASK,
		Shift:    TMC5240_SG4_THRS_SHIFT,
		Register: TMC5240_SG4_THRS,
		IsSigned: false,
	}
	TMC5240_SG4_FILT_EN_FIELD = RegisterField{
		Mask:     TMC5240_SG4_FILT_EN_MASK,
		Shift:    TMC5240_SG4_FILT_EN_SHIFT,
		Register: TMC5240_SG4_THRS,
		IsSigned: false,
	}
	TMC5240_SG_ANGLE_OFFSET_FIELD = RegisterField{
		Mask:     TMC5240_SG_ANGLE_OFFSET_MASK,
		Shift:    TMC5240_SG_ANGLE_OFFSET_SHIFT,
		Register: TMC5240_SG4_THRS,
		IsSigned: false,
	}
	TMC5240_SG4_RESULT_FIELD = RegisterField{
		Mask:     TMC5240_SG4_RESULT_MASK,
		Shift:    TMC5240_SG4_RESULT_SHIFT,
		Register: TMC5240_SG4_RESULT,
		IsSigned: false,
	}
	TMC5240_SG4_IND_0_FIELD = RegisterField{
		Mask:     TMC5240_SG4_IND_0_MASK,
		Shift:    TMC5240_SG4_IND_0_SHIFT,
		Register: TMC5240_SG4_IND,
		IsSigned: false,
	}
	TMC5240_SG4_IND_1_FIELD = RegisterField{
		Mask:     TMC5240_SG4_IND_1_MASK,
		Shift:    TMC5240_SG4_IND_1_SHIFT,
		Register: TMC5240_SG4_IND,
		IsSigned: false,
	}
	TMC5240_SG4_IND_2_FIELD = RegisterField{
		Mask:     TMC5240_SG4_IND_2_MASK,
		Shift:    TMC5240_SG4_IND_2_SHIFT,
		Register: TMC5240_SG4_IND,
		IsSigned: false,
	}
	TMC5240_SG4_IND_3_FIELD = RegisterField{
		Mask:     TMC5240_SG4_IND_3_MASK,
		Shift:    TMC5240_SG4_IND_3_SHIFT,
		Register: TMC5240_SG4_IND,
		IsSigned: false,
	}
)
