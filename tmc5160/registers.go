package tmc5160

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
	TMC5160_MOTORS                                  = 1
	TMC5160_WRITE_BIT                               = 0x80
	TMC5160_ADDRESS_MASK                            = 0x7F
	TMC5160_MAX_VELOCITY                            = 8388096
	TMC5160_MAX_ACCELERATION                        = 65535
	TMC5160_MODE_POSITION                           = 0
	TMC5160_MODE_VELPOS                             = 1
	TMC5160_MODE_VELNEG                             = 2
	TMC5160_MODE_HOLD                               = 3
	TMC5160_SW_STOPL_ENABLE                         = 0x0001
	TMC5160_SW_STOPR_ENABLE                         = 0x0002
	TMC5160_SW_STOPL_POLARITY                       = 0x0004
	TMC5160_SW_STOPR_POLARITY                       = 0x0008
	TMC5160_SW_SWAP_LR                              = 0x0010
	TMC5160_SW_LATCH_L_ACT                          = 0x0020
	TMC5160_SW_LATCH_L_INACT                        = 0x0040
	TMC5160_SW_LATCH_R_ACT                          = 0x0080
	TMC5160_SW_LATCH_R_INACT                        = 0x0100
	TMC5160_SW_LATCH_ENC                            = 0x0200
	TMC5160_SW_SG_STOP                              = 0x0400
	TMC5160_SW_SOFTSTOP                             = 0x0800
	TMC5160_RS_STOPL                                = 0x0001
	TMC5160_RS_STOPR                                = 0x0002
	TMC5160_RS_LATCHL                               = 0x0004
	TMC5160_RS_LATCHR                               = 0x0008
	TMC5160_RS_EV_STOPL                             = 0x0010
	TMC5160_RS_EV_STOPR                             = 0x0020
	TMC5160_RS_EV_STOP_SG                           = 0x0040
	TMC5160_RS_EV_POSREACHED                        = 0x0080
	TMC5160_RS_VELREACHED                           = 0x0100
	TMC5160_RS_POSREACHED                           = 0x0200
	TMC5160_RS_VZERO                                = 0x0400
	TMC5160_RS_ZEROWAIT                             = 0x0800
	TMC5160_RS_SECONDMOVE                           = 0x1000
	TMC5160_RS_SG                                   = 0x2000
	TMC5160_EM_DECIMAL                              = 0x0400
	TMC5160_EM_LATCH_XACT                           = 0x0200
	TMC5160_EM_CLR_XENC                             = 0x0100
	TMC5160_EM_NEG_EDGE                             = 0x0080
	TMC5160_EM_POS_EDGE                             = 0x0040
	TMC5160_EM_CLR_ONCE                             = 0x0020
	TMC5160_EM_CLR_CONT                             = 0x0010
	TMC5160_EM_IGNORE_AB                            = 0x0008
	TMC5160_EM_POL_N                                = 0x0004
	TMC5160_EM_POL_B                                = 0x0002
	TMC5160_EM_POL_A                                = 0x0001
	TMC5160_GCONF                                   = 0x00
	TMC5160_GSTAT                                   = 0x01
	TMC5160_IFCNT                                   = 0x02
	TMC5160_SLAVECONF                               = 0x03
	TMC5160_INP_OUT                                 = 0x04
	TMC5160_X_COMPARE                               = 0x05
	TMC5160_OTP_PROG                                = 0x06
	TMC5160_OTP_READ                                = 0x07
	TMC5160_FACTORY_CONF                            = 0x08
	TMC5160_SHORT_CONF                              = 0x09
	TMC5160_DRV_CONF                                = 0x0A
	TMC5160_GLOBAL_SCALER                           = 0x0B
	TMC5160_OFFSET_READ                             = 0x0C
	TMC5160_IHOLD_IRUN                              = 0x10
	TMC5160_TPOWERDOWN                              = 0x11
	TMC5160_TSTEP                                   = 0x12
	TMC5160_TPWMTHRS                                = 0x13
	TMC5160_TCOOLTHRS                               = 0x14
	TMC5160_THIGH                                   = 0x15
	TMC5160_RAMPMODE                                = 0x20
	TMC5160_XACTUAL                                 = 0x21
	TMC5160_VACTUAL                                 = 0x22
	TMC5160_VSTART                                  = 0x23
	TMC5160_A1                                      = 0x24
	TMC5160_V1                                      = 0x25
	TMC5160_AMAX                                    = 0x26
	TMC5160_VMAX                                    = 0x27
	TMC5160_DMAX                                    = 0x28
	TMC5160_D1                                      = 0x2A
	TMC5160_VSTOP                                   = 0x2B
	TMC5160_TZEROWAIT                               = 0x2C
	TMC5160_XTARGET                                 = 0x2D
	TMC5160_VDCMIN                                  = 0x33
	TMC5160_SWMODE                                  = 0x34
	TMC5160_RAMPSTAT                                = 0x35
	TMC5160_XLATCH                                  = 0x36
	TMC5160_ENCMODE                                 = 0x38
	TMC5160_XENC                                    = 0x39
	TMC5160_ENC_CONST                               = 0x3A
	TMC5160_ENC_STATUS                              = 0x3B
	TMC5160_ENC_LATCH                               = 0x3C
	TMC5160_ENC_DEVIATION                           = 0x3D
	TMC5160_MSLUT0                                  = 0x60
	TMC5160_MSLUT1                                  = 0x61
	TMC5160_MSLUT2                                  = 0x62
	TMC5160_MSLUT3                                  = 0x63
	TMC5160_MSLUT4                                  = 0x64
	TMC5160_MSLUT5                                  = 0x65
	TMC5160_MSLUT6                                  = 0x66
	TMC5160_MSLUT7                                  = 0x67
	TMC5160_MSLUTSEL                                = 0x68
	TMC5160_MSLUTSTART                              = 0x69
	TMC5160_MSCNT                                   = 0x6A
	TMC5160_MSCURACT                                = 0x6B
	TMC5160_CHOPCONF                                = 0x6C
	TMC5160_COOLCONF                                = 0x6D
	TMC5160_DCCTRL                                  = 0x6E
	TMC5160_DRV_STATUS                              = 0x6F
	TMC5160_PWMCONF                                 = 0x70
	TMC5160_PWMSCALE                                = 0x71
	TMC5160_PWM_AUTO                                = 0x72
	TMC5160_LOST_STEPS                              = 0x73
	TMC5160_SPI_STATUS_RESET_FLAG_MASK              = 0x01 /* GSTAT0 - 1: Signals, that a reset has occurred (clear by reading GSTAT) */
	TMC5160_SPI_STATUS_RESET_FLAG_SHIFT             = 0
	TMC5160_SPI_STATUS_DRIVER_ERROR_MASK            = 0x02 /* GSTAT1 – 1: Signals driver 1 driver error (clear by reading GSTAT) */
	TMC5160_SPI_STATUS_DRIVER_ERROR_SHIFT           = 1
	TMC5160_SPI_STATUS_SG2_MASK                     = 0x04 /* DRV_STATUS[24] – 1: Signals StallGuard flag active */
	TMC5160_SPI_STATUS_SG2_SHIFT                    = 2
	TMC5160_SPI_STATUS_STANDSTILL_MASK              = 0x08 /* DRV_STATUS[31] – 1: Signals motor stand still */
	TMC5160_SPI_STATUS_STANDSTILL_SHIFT             = 3
	TMC5160_SPI_STATUS_VELOCITY_REACHED_MASK        = 0x10 /* RAMPSTAT[8] – 1: Signals target velocity reached (motion controller only) */
	TMC5160_SPI_STATUS_VELOCITY_REACHED_SHIFT       = 4
	TMC5160_SPI_STATUS_POSITION_REACHED_MASK        = 0x20 /* RAMPSTAT[9] – 1: Signals target position reached (motion controller only) */
	TMC5160_SPI_STATUS_POSITION_REACHED_SHIFT       = 5
	TMC5160_SPI_STATUS_STATUS_STOP_L_MASK           = 0x40 /* RAMPSTAT0 – 1: Signals stop left switch status (motion controller only) */
	TMC5160_SPI_STATUS_STATUS_STOP_L_SHIFT          = 6
	TMC5160_SPI_STATUS_STATUS_STOP_R_MASK           = 0x80 /* RAMPSTAT1 – 1: Signals stop right switch status (motion controller only) */
	TMC5160_SPI_STATUS_STATUS_STOP_R_SHIFT          = 7
	TMC5160_RECALIBRATE_MASK                        = 0x00000001
	TMC5160_RECALIBRATE_SHIFT                       = 0
	TMC5160_FASTSTANDSTILL_MASK                     = 0x00000002
	TMC5160_FASTSTANDSTILL_SHIFT                    = 1
	TMC5160_EN_PWM_MODE_MASK                        = 0x00000004
	TMC5160_EN_PWM_MODE_SHIFT                       = 2
	TMC5160_MULTISTEP_FILT_MASK                     = 0x00000008
	TMC5160_MULTISTEP_FILT_SHIFT                    = 3
	TMC5160_SHAFT_MASK                              = 0x00000010
	TMC5160_SHAFT_SHIFT                             = 4
	TMC5160_DIAG0_ERROR__ONLY_WITH_SD_MODE_1__MASK  = 0x00000020
	TMC5160_DIAG0_ERROR__ONLY_WITH_SD_MODE_1__SHIFT = 5
	TMC5160_DIAG0_OTPW__ONLY_WITH_SD_MODE_1__MASK   = 0x00000040
	TMC5160_DIAG0_OTPW__ONLY_WITH_SD_MODE_1__SHIFT  = 6
	TMC5160_DIAG0_STALL_MASK                        = 0x00000080
	TMC5160_DIAG0_STALL_SHIFT                       = 7
	TMC5160_DIAG1_STALL_MASK                        = 0x00000100
	TMC5160_DIAG1_STALL_SHIFT                       = 8
	TMC5160_DIAG1_INDEX_MASK                        = 0x00000200
	TMC5160_DIAG1_INDEX_SHIFT                       = 9
	TMC5160_DIAG1_ONSTATE_MASK                      = 0x00000400
	TMC5160_DIAG1_ONSTATE_SHIFT                     = 10
	TMC5160_DIAG1_STEPS_SKIPPED_MASK                = 0x00000800
	TMC5160_DIAG1_STEPS_SKIPPED_SHIFT               = 11
	TMC5160_DIAG0_INT_PUSHPULL_MASK                 = 0x00001000
	TMC5160_DIAG0_INT_PUSHPULL_SHIFT                = 12
	TMC5160_DIAG1_POSCOMP_PUSHPULL_MASK             = 0x00002000
	TMC5160_DIAG1_POSCOMP_PUSHPULL_SHIFT            = 13
	TMC5160_SMALL_HYSTERESIS_MASK                   = 0x00004000
	TMC5160_SMALL_HYSTERESIS_SHIFT                  = 14
	TMC5160_STOP_ENABLE_MASK                        = 0x00008000
	TMC5160_STOP_ENABLE_SHIFT                       = 15
	TMC5160_DIRECT_MODE_MASK                        = 0x00010000
	TMC5160_DIRECT_MODE_SHIFT                       = 16
	TMC5160_TEST_MODE_MASK                          = 0x00020000
	TMC5160_TEST_MODE_SHIFT                         = 17
	TMC5160_DIAG0_STEP_MASK                         = 0x00000080
	TMC5160_DIAG0_STEP_SHIFT                        = 7
	TMC5160_DIAG1_DIR_MASK                          = 0x00000100
	TMC5160_DIAG1_DIR_SHIFT                         = 8

	TMC5160_RESET_MASK                = 0x00000001
	TMC5160_RESET_SHIFT               = 0
	TMC5160_DRV_ERR_MASK              = 0x00000002
	TMC5160_DRV_ERR_SHIFT             = 1
	TMC5160_UV_CP_MASK                = 0x00000004
	TMC5160_UV_CP_SHIFT               = 2
	TMC5160_IFCNT_MASK                = 0x000000FF
	TMC5160_IFCNT_SHIFT               = 0
	TMC5160_SLAVEADDR_MASK            = 0x000000FF
	TMC5160_SLAVEADDR_SHIFT           = 0
	TMC5160_SENDDELAY_MASK            = 0x00000F00
	TMC5160_SENDDELAY_SHIFT           = 8
	TMC5160_REFL_STEP_MASK            = 0x00000001
	TMC5160_REFL_STEP_SHIFT           = 0
	TMC5160_REFR_DIR_MASK             = 0x00000002
	TMC5160_REFR_DIR_SHIFT            = 1
	TMC5160_ENCB_DCEN_CFG4_MASK       = 0x00000004
	TMC5160_ENCB_DCEN_CFG4_SHIFT      = 2
	TMC5160_ENCA_DCIN_CFG5_MASK       = 0x00000008
	TMC5160_ENCA_DCIN_CFG5_SHIFT      = 3
	TMC5160_DRV_ENN_CFG6_MASK         = 0x00000010
	TMC5160_DRV_ENN_CFG6_SHIFT        = 4
	TMC5160_ENC_N_DCO_MASK            = 0x00000020
	TMC5160_ENC_N_DCO_SHIFT           = 5
	TMC5160_SD_MODE_MASK              = 0x00000040
	TMC5160_SD_MODE_SHIFT             = 6
	TMC5160_SWCOMP_IN_MASK            = 0x00000080
	TMC5160_SWCOMP_IN_SHIFT           = 7
	TMC5160_VERSION_MASK              = 0xFF000000
	TMC5160_VERSION_SHIFT             = 24
	TMC5160_OUTPUT_PIN_POLARITY_MASK  = 0x00000001
	TMC5160_OUTPUT_PIN_POLARITY_SHIFT = 0
	TMC5160_X_COMPARE_MASK            = 0xFFFFFFFF
	TMC5160_X_COMPARE_SHIFT           = 0
	TMC5160_OTPBIT_MASK               = 0x00000007
	TMC5160_OTPBIT_SHIFT              = 0
	TMC5160_OTPBYTE_MASK              = 0x00000030
	TMC5160_OTPBYTE_SHIFT             = 4
	TMC5160_OTPMAGIC_MASK             = 0x0000FF00
	TMC5160_OTPMAGIC_SHIFT            = 8
	TMC5160_OTP_TBL_MASK              = 0x00000080
	TMC5160_OTP_TBL_SHIFT             = 7
	TMC5160_OTP_BBM_MASK              = 0x00000040
	TMC5160_OTP_BBM_SHIFT             = 6
	TMC5160_OTP_S2_LEVEL_MASK         = 0x00000020
	TMC5160_OTP_S2_LEVEL_SHIFT        = 5
	TMC5160_OTP_FCLKTRIM_MASK         = 0x0000001F
	TMC5160_OTP_FCLKTRIM_SHIFT        = 0
	TMC5160_FCLKTRIM_MASK             = 0x0000001F
	TMC5160_FCLKTRIM_SHIFT            = 0
	TMC5160_S2VS_LEVEL_MASK           = 0x0000000F
	TMC5160_S2VS_LEVEL_SHIFT          = 0
	TMC5160_S2GND_LEVEL_MASK          = 0x00000F00
	TMC5160_S2GND_LEVEL_SHIFT         = 8
	TMC5160_SHORTFILTER_MASK          = 0x00030000
	TMC5160_SHORTFILTER_SHIFT         = 16
	TMC5160_SHORTDELAY_MASK           = 0x00040000
	TMC5160_SHORTDELAY_SHIFT          = 18
	TMC5160_BBMTIME_MASK              = 0x0000001F
	TMC5160_BBMTIME_SHIFT             = 0
	TMC5160_BBMCLKS_MASK              = 0x00000F00
	TMC5160_BBMCLKS_SHIFT             = 8
	TMC5160_OTSELECT_MASK             = 0x00030000
	TMC5160_OTSELECT_SHIFT            = 16
	TMC5160_DRVSTRENGTH_MASK          = 0x000C0000
	TMC5160_DRVSTRENGTH_SHIFT         = 18
	TMC5160_FILT_ISENSE_MASK          = 0x00300000
	TMC5160_FILT_ISENSE_SHIFT         = 20
	TMC5160_GLOBAL_SCALER_MASK        = 0x000000FF
	TMC5160_GLOBAL_SCALER_SHIFT       = 0
	TMC5160_OFFSET_READ_A_MASK        = 0x0000FF00
	TMC5160_OFFSET_READ_A_SHIFT       = 8
	TMC5160_OFFSET_READ_B_MASK        = 0x000000FF
	TMC5160_OFFSET_READ_B_SHIFT       = 0
	TMC5160_IHOLD_MASK                = 0x0000001F
	TMC5160_IHOLD_SHIFT               = 0
	TMC5160_IRUN_MASK                 = 0x00001F00
	TMC5160_IRUN_SHIFT                = 8
	TMC5160_IHOLDDELAY_MASK           = 0x000F0000
	TMC5160_IHOLDDELAY_SHIFT          = 16
	TMC5160_TPOWERDOWN_MASK           = 0x000000FF
	TMC5160_TPOWERDOWN_SHIFT          = 0
	TMC5160_TSTEP_MASK                = 0x000FFFFF
	TMC5160_TSTEP_SHIFT               = 0
	TMC5160_TPWMTHRS_MASK             = 0x000FFFFF
	TMC5160_TPWMTHRS_SHIFT            = 0
	TMC5160_TCOOLTHRS_MASK            = 0x000FFFFF
	TMC5160_TCOOLTHRS_SHIFT           = 0
	TMC5160_THIGH_MASK                = 0x000FFFFF
	TMC5160_THIGH_SHIFT               = 0
	TMC5160_RAMPMODE_MASK             = 0x00000003
	TMC5160_RAMPMODE_SHIFT            = 0
	TMC5160_XACTUAL_MASK              = 0xFFFFFFFF
	TMC5160_XACTUAL_SHIFT             = 0
	TMC5160_VACTUAL_MASK              = 0x00FFFFFF
	TMC5160_VACTUAL_SHIFT             = 0
	TMC5160_VSTART_MASK               = 0x0003FFFF
	TMC5160_VSTART_SHIFT              = 0
	TMC5160_A1_MASK                   = 0x0000FFFF
	TMC5160_A1_SHIFT                  = 0
	TMC5160_V1__MASK                  = 0x000FFFFF
	TMC5160_V1__SHIFT                 = 0
	TMC5160_AMAX_MASK                 = 0x0000FFFF
	TMC5160_AMAX_SHIFT                = 0
	TMC5160_VMAX_MASK                 = 0x007FFFFF
	TMC5160_VMAX_SHIFT                = 0
	TMC5160_DMAX_MASK                 = 0x0000FFFF
	TMC5160_DMAX_SHIFT                = 0
	TMC5160_D1_MASK                   = 0x0000FFFF
	TMC5160_D1_SHIFT                  = 0
	TMC5160_VSTOP_MASK                = 0x0003FFFF
	TMC5160_VSTOP_SHIFT               = 0
	TMC5160_TZEROWAIT_MASK            = 0x0000FFFF
	TMC5160_TZEROWAIT_SHIFT           = 0
	TMC5160_XTARGET_MASK              = 0xFFFFFFFF
	TMC5160_XTARGET_SHIFT             = 0
	TMC5160_VDCMIN_MASK               = 0x007FFFFF
	TMC5160_VDCMIN_SHIFT              = 0
	TMC5160_STOP_L_ENABLE_MASK        = 0x00000001
	TMC5160_STOP_L_ENABLE_SHIFT       = 0
	TMC5160_STOP_R_ENABLE_MASK        = 0x00000002
	TMC5160_STOP_R_ENABLE_SHIFT       = 1
	TMC5160_POL_STOP_L_MASK           = 0x00000004
	TMC5160_POL_STOP_L_SHIFT          = 2
	TMC5160_POL_STOP_R_MASK           = 0x00000008
	TMC5160_POL_STOP_R_SHIFT          = 3
	TMC5160_SWAP_LR_MASK              = 0x00000010
	TMC5160_SWAP_LR_SHIFT             = 4
	TMC5160_LATCH_L_ACTIVE_MASK       = 0x00000020
	TMC5160_LATCH_L_ACTIVE_SHIFT      = 5
	TMC5160_LATCH_L_INACTIVE_MASK     = 0x00000040
	TMC5160_LATCH_L_INACTIVE_SHIFT    = 6
	TMC5160_LATCH_R_ACTIVE_MASK       = 0x00000080
	TMC5160_LATCH_R_ACTIVE_SHIFT      = 7
	TMC5160_LATCH_R_INACTIVE_MASK     = 0x00000100
	TMC5160_LATCH_R_INACTIVE_SHIFT    = 8
	TMC5160_EN_LATCH_ENCODER_MASK     = 0x00000200
	TMC5160_EN_LATCH_ENCODER_SHIFT    = 9
	TMC5160_SG_STOP_MASK              = 0x00000400
	TMC5160_SG_STOP_SHIFT             = 10
	TMC5160_EN_SOFTSTOP_MASK          = 0x00000800
	TMC5160_EN_SOFTSTOP_SHIFT         = 11
	TMC5160_STATUS_STOP_L_MASK        = 0x00000001
	TMC5160_STATUS_STOP_L_SHIFT       = 0
	TMC5160_STATUS_STOP_R_MASK        = 0x00000002
	TMC5160_STATUS_STOP_R_SHIFT       = 1
	TMC5160_STATUS_LATCH_L_MASK       = 0x00000004
	TMC5160_STATUS_LATCH_L_SHIFT      = 2
	TMC5160_STATUS_LATCH_R_MASK       = 0x00000008
	TMC5160_STATUS_LATCH_R_SHIFT      = 3
	TMC5160_EVENT_STOP_L_MASK         = 0x00000010
	TMC5160_EVENT_STOP_L_SHIFT        = 4
	TMC5160_EVENT_STOP_R_MASK         = 0x00000020
	TMC5160_EVENT_STOP_R_SHIFT        = 5
	TMC5160_EVENT_STOP_SG_MASK        = 0x00000040
	TMC5160_EVENT_STOP_SG_SHIFT       = 6
	TMC5160_EVENT_POS_REACHED_MASK    = 0x00000080
	TMC5160_EVENT_POS_REACHED_SHIFT   = 7
	TMC5160_VELOCITY_REACHED_MASK     = 0x00000100
	TMC5160_VELOCITY_REACHED_SHIFT    = 8
	TMC5160_POSITION_REACHED_MASK     = 0x00000200
	TMC5160_POSITION_REACHED_SHIFT    = 9
	TMC5160_VZERO_MASK                = 0x00000400
	TMC5160_VZERO_SHIFT               = 10
	TMC5160_T_ZEROWAIT_ACTIVE_MASK    = 0x00000800
	TMC5160_T_ZEROWAIT_ACTIVE_SHIFT   = 11
	TMC5160_SECOND_MOVE_MASK          = 0x00001000
	TMC5160_SECOND_MOVE_SHIFT         = 12
	TMC5160_STATUS_SG_MASK            = 0x00002000
	TMC5160_STATUS_SG_SHIFT           = 13
	TMC5160_XLATCH_MASK               = 0xFFFFFFFF
	TMC5160_XLATCH_SHIFT              = 0
	TMC5160_POL_A_MASK                = 0x00000001
	TMC5160_POL_A_SHIFT               = 0
	TMC5160_POL_B_MASK                = 0x00000002
	TMC5160_POL_B_SHIFT               = 1
	TMC5160_POL_N_MASK                = 0x00000004
	TMC5160_POL_N_SHIFT               = 2
	TMC5160_IGNORE_AB_MASK            = 0x00000008
	TMC5160_IGNORE_AB_SHIFT           = 3
	TMC5160_CLR_CONT_MASK             = 0x00000010
	TMC5160_CLR_CONT_SHIFT            = 4
	TMC5160_CLR_ONCE_MASK             = 0x00000020
	TMC5160_CLR_ONCE_SHIFT            = 5
	TMC5160_POS_EDGE_NEG_EDGE_MASK    = 0x000000C0
	TMC5160_POS_EDGE_NEG_EDGE_SHIFT   = 6
	TMC5160_CLR_ENC_X_MASK            = 0x00000100
	TMC5160_CLR_ENC_X_SHIFT           = 8
	TMC5160_LATCH_X_ACT_MASK          = 0x00000200
	TMC5160_LATCH_X_ACT_SHIFT         = 9
	TMC5160_ENC_SEL_DECIMAL_MASK      = 0x00000400
	TMC5160_ENC_SEL_DECIMAL_SHIFT     = 10
	TMC5160_X_ENC_MASK                = 0xFFFFFFFF
	TMC5160_X_ENC_SHIFT               = 0
	TMC5160_INTEGER_MASK              = 0xFFFF0000
	TMC5160_INTEGER_SHIFT             = 16
	TMC5160_FRACTIONAL_MASK           = 0x0000FFFF
	TMC5160_FRACTIONAL_SHIFT          = 0
	TMC5160_N_EVENT_MASK              = 0x00000001
	TMC5160_N_EVENT_SHIFT             = 0
	TMC5160_DEVIATION_WARN_MASK       = 0x00000002
	TMC5160_DEVIATION_WARN_SHIFT      = 1
	TMC5160_ENC_LATCH_MASK            = 0xFFFFFFFF
	TMC5160_ENC_LATCH_SHIFT           = 0
	TMC5160_ENC_DEVIATION_MASK        = 0x000FFFFF
	TMC5160_ENC_DEVIATION_SHIFT       = 0
	TMC5160_OFS0_MASK                 = 0x00000001
	TMC5160_OFS0_SHIFT                = 0
	TMC5160_OFS1_MASK                 = 0x00000002
	TMC5160_OFS1_SHIFT                = 1
	TMC5160_OFS2_MASK                 = 0x00000004
	TMC5160_OFS2_SHIFT                = 2
	TMC5160_OFS3_MASK                 = 0x00000008
	TMC5160_OFS3_SHIFT                = 3
	TMC5160_OFS4_MASK                 = 0x00000010
	TMC5160_OFS4_SHIFT                = 4
	TMC5160_OFS5_MASK                 = 0x00000020
	TMC5160_OFS5_SHIFT                = 5
	TMC5160_OFS6_MASK                 = 0x00000040
	TMC5160_OFS6_SHIFT                = 6
	TMC5160_OFS7_MASK                 = 0x00000080
	TMC5160_OFS7_SHIFT                = 7
	TMC5160_OFS8_MASK                 = 0x00000100
	TMC5160_OFS8_SHIFT                = 8
	TMC5160_OFS9_MASK                 = 0x00000200
	TMC5160_OFS9_SHIFT                = 9
	TMC5160_OFS10_MASK                = 0x00000400
	TMC5160_OFS10_SHIFT               = 10
	TMC5160_OFS11_MASK                = 0x00000800
	TMC5160_OFS11_SHIFT               = 11
	TMC5160_OFS12_MASK                = 0x00001000
	TMC5160_OFS12_SHIFT               = 12
	TMC5160_OFS13_MASK                = 0x00002000
	TMC5160_OFS13_SHIFT               = 13
	TMC5160_OFS14_MASK                = 0x00004000
	TMC5160_OFS14_SHIFT               = 14
	TMC5160_OFS15_MASK                = 0x00008000
	TMC5160_OFS15_SHIFT               = 15
	TMC5160_OFS16_MASK                = 0x00010000
	TMC5160_OFS16_SHIFT               = 16
	TMC5160_OFS17_MASK                = 0x00020000
	TMC5160_OFS17_SHIFT               = 17
	TMC5160_OFS18_MASK                = 0x00040000
	TMC5160_OFS18_SHIFT               = 18
	TMC5160_OFS19_MASK                = 0x00080000
	TMC5160_OFS19_SHIFT               = 19
	TMC5160_OFS20_MASK                = 0x00100000
	TMC5160_OFS20_SHIFT               = 20
	TMC5160_OFS21_MASK                = 0x00200000
	TMC5160_OFS21_SHIFT               = 21
	TMC5160_OFS22_MASK                = 0x00400000
	TMC5160_OFS22_SHIFT               = 22
	TMC5160_OFS23_MASK                = 0x00800000
	TMC5160_OFS23_SHIFT               = 23
	TMC5160_OFS24_MASK                = 0x01000000
	TMC5160_OFS24_SHIFT               = 24
	TMC5160_OFS25_MASK                = 0x02000000
	TMC5160_OFS25_SHIFT               = 25
	TMC5160_OFS26_MASK                = 0x04000000
	TMC5160_OFS26_SHIFT               = 26
	TMC5160_OFS27_MASK                = 0x08000000
	TMC5160_OFS27_SHIFT               = 27
	TMC5160_OFS28_MASK                = 0x10000000
	TMC5160_OFS28_SHIFT               = 28
	TMC5160_OFS29_MASK                = 0x20000000
	TMC5160_OFS29_SHIFT               = 29
	TMC5160_OFS30_MASK                = 0x40000000
	TMC5160_OFS30_SHIFT               = 30
	TMC5160_OFS31_MASK                = 0x80000000
	TMC5160_OFS31_SHIFT               = 31
	TMC5160_OFS32_MASK                = 0x00000001
	TMC5160_OFS32_SHIFT               = 0
	TMC5160_OFS33_MASK                = 0x00000002
	TMC5160_OFS33_SHIFT               = 1
	TMC5160_OFS34_MASK                = 0x00000004
	TMC5160_OFS34_SHIFT               = 2
	TMC5160_OFS35_MASK                = 0x00000008
	TMC5160_OFS35_SHIFT               = 3
	TMC5160_OFS36_MASK                = 0x00000010
	TMC5160_OFS36_SHIFT               = 4
	TMC5160_OFS37_MASK                = 0x00000020
	TMC5160_OFS37_SHIFT               = 5
	TMC5160_OFS38_MASK                = 0x00000040
	TMC5160_OFS38_SHIFT               = 6
	TMC5160_OFS39_MASK                = 0x00000080
	TMC5160_OFS39_SHIFT               = 7
	TMC5160_OFS40_MASK                = 0x00000100
	TMC5160_OFS40_SHIFT               = 8
	TMC5160_OFS41_MASK                = 0x00000200
	TMC5160_OFS41_SHIFT               = 9
	TMC5160_OFS42_MASK                = 0x00000400
	TMC5160_OFS42_SHIFT               = 10
	TMC5160_OFS43_MASK                = 0x00000800
	TMC5160_OFS43_SHIFT               = 11
	TMC5160_OFS44_MASK                = 0x00001000
	TMC5160_OFS44_SHIFT               = 12
	TMC5160_OFS45_MASK                = 0x00002000
	TMC5160_OFS45_SHIFT               = 13
	TMC5160_OFS46_MASK                = 0x00004000
	TMC5160_OFS46_SHIFT               = 14
	TMC5160_OFS47_MASK                = 0x00008000
	TMC5160_OFS47_SHIFT               = 15
	TMC5160_OFS48_MASK                = 0x00010000
	TMC5160_OFS48_SHIFT               = 16
	TMC5160_OFS49_MASK                = 0x00020000
	TMC5160_OFS49_SHIFT               = 17
	TMC5160_OFS50_MASK                = 0x00040000
	TMC5160_OFS50_SHIFT               = 18
	TMC5160_OFS51_MASK                = 0x00080000
	TMC5160_OFS51_SHIFT               = 19
	TMC5160_OFS52_MASK                = 0x00100000
	TMC5160_OFS52_SHIFT               = 20
	TMC5160_OFS53_MASK                = 0x00200000
	TMC5160_OFS53_SHIFT               = 21
	TMC5160_OFS54_MASK                = 0x00400000
	TMC5160_OFS54_SHIFT               = 22
	TMC5160_OFS55_MASK                = 0x00800000
	TMC5160_OFS55_SHIFT               = 23
	TMC5160_OFS56_MASK                = 0x01000000
	TMC5160_OFS56_SHIFT               = 24
	TMC5160_OFS57_MASK                = 0x02000000
	TMC5160_OFS57_SHIFT               = 25
	TMC5160_OFS58_MASK                = 0x04000000
	TMC5160_OFS58_SHIFT               = 26
	TMC5160_OFS59_MASK                = 0x08000000
	TMC5160_OFS59_SHIFT               = 27
	TMC5160_OFS60_MASK                = 0x10000000
	TMC5160_OFS60_SHIFT               = 28
	TMC5160_OFS61_MASK                = 0x20000000
	TMC5160_OFS61_SHIFT               = 29
	TMC5160_OFS62_MASK                = 0x40000000
	TMC5160_OFS62_SHIFT               = 30
	TMC5160_OFS63_MASK                = 0x80000000
	TMC5160_OFS63_SHIFT               = 31
	TMC5160_OFS64_MASK                = 0x00000001
	TMC5160_OFS64_SHIFT               = 0
	TMC5160_OFS65_MASK                = 0x00000002
	TMC5160_OFS65_SHIFT               = 1
	TMC5160_OFS66_MASK                = 0x00000004
	TMC5160_OFS66_SHIFT               = 2
	TMC5160_OFS67_MASK                = 0x00000008
	TMC5160_OFS67_SHIFT               = 3
	TMC5160_OFS68_MASK                = 0x00000010
	TMC5160_OFS68_SHIFT               = 4
	TMC5160_OFS69_MASK                = 0x00000020
	TMC5160_OFS69_SHIFT               = 5
	TMC5160_OFS70_MASK                = 0x00000040
	TMC5160_OFS70_SHIFT               = 6
	TMC5160_OFS71_MASK                = 0x00000080
	TMC5160_OFS71_SHIFT               = 7
	TMC5160_OFS72_MASK                = 0x00000100
	TMC5160_OFS72_SHIFT               = 8
	TMC5160_OFS73_MASK                = 0x00000200
	TMC5160_OFS73_SHIFT               = 9
	TMC5160_OFS74_MASK                = 0x00000400
	TMC5160_OFS74_SHIFT               = 10
	TMC5160_OFS75_MASK                = 0x00000800
	TMC5160_OFS75_SHIFT               = 11
	TMC5160_OFS76_MASK                = 0x00001000
	TMC5160_OFS76_SHIFT               = 12
	TMC5160_OFS77_MASK                = 0x00002000
	TMC5160_OFS77_SHIFT               = 13
	TMC5160_OFS78_MASK                = 0x00004000
	TMC5160_OFS78_SHIFT               = 14
	TMC5160_OFS79_MASK                = 0x00008000
	TMC5160_OFS79_SHIFT               = 15
	TMC5160_OFS80_MASK                = 0x00010000
	TMC5160_OFS80_SHIFT               = 16
	TMC5160_OFS81_MASK                = 0x00020000
	TMC5160_OFS81_SHIFT               = 17
	TMC5160_OFS82_MASK                = 0x00040000
	TMC5160_OFS82_SHIFT               = 18
	TMC5160_OFS83_MASK                = 0x00080000
	TMC5160_OFS83_SHIFT               = 19
	TMC5160_OFS84_MASK                = 0x00100000
	TMC5160_OFS84_SHIFT               = 20
	TMC5160_OFS85_MASK                = 0x00200000
	TMC5160_OFS85_SHIFT               = 21
	TMC5160_OFS86_MASK                = 0x00400000
	TMC5160_OFS86_SHIFT               = 22
	TMC5160_OFS87_MASK                = 0x00800000
	TMC5160_OFS87_SHIFT               = 23
	TMC5160_OFS88_MASK                = 0x01000000
	TMC5160_OFS88_SHIFT               = 24
	TMC5160_OFS89_MASK                = 0x02000000
	TMC5160_OFS89_SHIFT               = 25
	TMC5160_OFS90_MASK                = 0x04000000
	TMC5160_OFS90_SHIFT               = 26
	TMC5160_OFS91_MASK                = 0x08000000
	TMC5160_OFS91_SHIFT               = 27
	TMC5160_OFS92_MASK                = 0x10000000
	TMC5160_OFS92_SHIFT               = 28
	TMC5160_OFS93_MASK                = 0x20000000
	TMC5160_OFS93_SHIFT               = 29
	TMC5160_OFS94_MASK                = 0x40000000
	TMC5160_OFS94_SHIFT               = 30
	TMC5160_OFS95_MASK                = 0x80000000
	TMC5160_OFS95_SHIFT               = 31
	TMC5160_OFS96_MASK                = 0x00000001
	TMC5160_OFS96_SHIFT               = 0
	TMC5160_OFS97_MASK                = 0x00000002
	TMC5160_OFS97_SHIFT               = 1
	TMC5160_OFS98_MASK                = 0x00000004
	TMC5160_OFS98_SHIFT               = 2
	TMC5160_OFS99_MASK                = 0x00000008
	TMC5160_OFS99_SHIFT               = 3
	TMC5160_OFS100_MASK               = 0x00000010
	TMC5160_OFS100_SHIFT              = 4
	TMC5160_OFS101_MASK               = 0x00000020
	TMC5160_OFS101_SHIFT              = 5
	TMC5160_OFS102_MASK               = 0x00000040
	TMC5160_OFS102_SHIFT              = 6
	TMC5160_OFS103_MASK               = 0x00000080
	TMC5160_OFS103_SHIFT              = 7
	TMC5160_OFS104_MASK               = 0x00000100
	TMC5160_OFS104_SHIFT              = 8
	TMC5160_OFS105_MASK               = 0x00000200
	TMC5160_OFS105_SHIFT              = 9
	TMC5160_OFS106_MASK               = 0x00000400
	TMC5160_OFS106_SHIFT              = 10
	TMC5160_OFS107_MASK               = 0x00000800
	TMC5160_OFS107_SHIFT              = 11
	TMC5160_OFS108_MASK               = 0x00001000
	TMC5160_OFS108_SHIFT              = 12
	TMC5160_OFS109_MASK               = 0x00002000
	TMC5160_OFS109_SHIFT              = 13
	TMC5160_OFS110_MASK               = 0x00004000
	TMC5160_OFS110_SHIFT              = 14
	TMC5160_OFS111_MASK               = 0x00008000
	TMC5160_OFS111_SHIFT              = 15
	TMC5160_OFS112_MASK               = 0x00010000
	TMC5160_OFS112_SHIFT              = 16
	TMC5160_OFS113_MASK               = 0x00020000
	TMC5160_OFS113_SHIFT              = 17
	TMC5160_OFS114_MASK               = 0x00040000
	TMC5160_OFS114_SHIFT              = 18
	TMC5160_OFS115_MASK               = 0x00080000
	TMC5160_OFS115_SHIFT              = 19
	TMC5160_OFS116_MASK               = 0x00100000
	TMC5160_OFS116_SHIFT              = 20
	TMC5160_OFS117_MASK               = 0x00200000
	TMC5160_OFS117_SHIFT              = 21
	TMC5160_OFS118_MASK               = 0x00400000
	TMC5160_OFS118_SHIFT              = 22
	TMC5160_OFS119_MASK               = 0x00800000
	TMC5160_OFS119_SHIFT              = 23
	TMC5160_OFS120_MASK               = 0x01000000
	TMC5160_OFS120_SHIFT              = 24
	TMC5160_OFS121_MASK               = 0x02000000
	TMC5160_OFS121_SHIFT              = 25
	TMC5160_OFS122_MASK               = 0x04000000
	TMC5160_OFS122_SHIFT              = 26
	TMC5160_OFS123_MASK               = 0x08000000
	TMC5160_OFS123_SHIFT              = 27
	TMC5160_OFS124_MASK               = 0x10000000
	TMC5160_OFS124_SHIFT              = 28
	TMC5160_OFS125_MASK               = 0x20000000
	TMC5160_OFS125_SHIFT              = 29
	TMC5160_OFS126_MASK               = 0x40000000
	TMC5160_OFS126_SHIFT              = 30
	TMC5160_OFS127_MASK               = 0x80000000
	TMC5160_OFS127_SHIFT              = 31
	TMC5160_OFS128_MASK               = 0x00000001
	TMC5160_OFS128_SHIFT              = 0
	TMC5160_OFS129_MASK               = 0x00000002
	TMC5160_OFS129_SHIFT              = 1
	TMC5160_OFS130_MASK               = 0x00000004
	TMC5160_OFS130_SHIFT              = 2
	TMC5160_OFS131_MASK               = 0x00000008
	TMC5160_OFS131_SHIFT              = 3
	TMC5160_OFS132_MASK               = 0x00000010
	TMC5160_OFS132_SHIFT              = 4
	TMC5160_OFS133_MASK               = 0x00000020
	TMC5160_OFS133_SHIFT              = 5
	TMC5160_OFS134_MASK               = 0x00000040
	TMC5160_OFS134_SHIFT              = 6
	TMC5160_OFS135_MASK               = 0x00000080
	TMC5160_OFS135_SHIFT              = 7
	TMC5160_OFS136_MASK               = 0x00000100
	TMC5160_OFS136_SHIFT              = 8
	TMC5160_OFS137_MASK               = 0x00000200
	TMC5160_OFS137_SHIFT              = 9
	TMC5160_OFS138_MASK               = 0x00000400
	TMC5160_OFS138_SHIFT              = 10
	TMC5160_OFS139_MASK               = 0x00000800
	TMC5160_OFS139_SHIFT              = 11
	TMC5160_OFS140_MASK               = 0x00001000
	TMC5160_OFS140_SHIFT              = 12
	TMC5160_OFS141_MASK               = 0x00002000
	TMC5160_OFS141_SHIFT              = 13
	TMC5160_OFS142_MASK               = 0x00004000
	TMC5160_OFS142_SHIFT              = 14
	TMC5160_OFS143_MASK               = 0x00008000
	TMC5160_OFS143_SHIFT              = 15
	TMC5160_OFS144_MASK               = 0x00010000
	TMC5160_OFS144_SHIFT              = 16
	TMC5160_OFS145_MASK               = 0x00020000
	TMC5160_OFS145_SHIFT              = 17
	TMC5160_OFS146_MASK               = 0x00040000
	TMC5160_OFS146_SHIFT              = 18
	TMC5160_OFS147_MASK               = 0x00080000
	TMC5160_OFS147_SHIFT              = 19
	TMC5160_OFS148_MASK               = 0x00100000
	TMC5160_OFS148_SHIFT              = 20
	TMC5160_OFS149_MASK               = 0x00200000
	TMC5160_OFS149_SHIFT              = 21
	TMC5160_OFS150_MASK               = 0x00400000
	TMC5160_OFS150_SHIFT              = 22
	TMC5160_OFS151_MASK               = 0x00800000
	TMC5160_OFS151_SHIFT              = 23
	TMC5160_OFS152_MASK               = 0x01000000
	TMC5160_OFS152_SHIFT              = 24
	TMC5160_OFS153_MASK               = 0x02000000
	TMC5160_OFS153_SHIFT              = 25
	TMC5160_OFS154_MASK               = 0x04000000
	TMC5160_OFS154_SHIFT              = 26
	TMC5160_OFS155_MASK               = 0x08000000
	TMC5160_OFS155_SHIFT              = 27
	TMC5160_OFS156_MASK               = 0x10000000
	TMC5160_OFS156_SHIFT              = 28
	TMC5160_OFS157_MASK               = 0x20000000
	TMC5160_OFS157_SHIFT              = 29
	TMC5160_OFS158_MASK               = 0x40000000
	TMC5160_OFS158_SHIFT              = 30
	TMC5160_OFS159_MASK               = 0x80000000
	TMC5160_OFS159_SHIFT              = 31
	TMC5160_OFS160_MASK               = 0x00000001
	TMC5160_OFS160_SHIFT              = 0
	TMC5160_OFS161_MASK               = 0x00000002
	TMC5160_OFS161_SHIFT              = 1
	TMC5160_OFS162_MASK               = 0x00000004
	TMC5160_OFS162_SHIFT              = 2
	TMC5160_OFS163_MASK               = 0x00000008
	TMC5160_OFS163_SHIFT              = 3
	TMC5160_OFS164_MASK               = 0x00000010
	TMC5160_OFS164_SHIFT              = 4
	TMC5160_OFS165_MASK               = 0x00000020
	TMC5160_OFS165_SHIFT              = 5
	TMC5160_OFS166_MASK               = 0x00000040
	TMC5160_OFS166_SHIFT              = 6
	TMC5160_OFS167_MASK               = 0x00000080
	TMC5160_OFS167_SHIFT              = 7
	TMC5160_OFS168_MASK               = 0x00000100
	TMC5160_OFS168_SHIFT              = 8
	TMC5160_OFS169_MASK               = 0x00000200
	TMC5160_OFS169_SHIFT              = 9
	TMC5160_OFS170_MASK               = 0x00000400
	TMC5160_OFS170_SHIFT              = 10
	TMC5160_OFS171_MASK               = 0x00000800
	TMC5160_OFS171_SHIFT              = 11
	TMC5160_OFS172_MASK               = 0x00001000
	TMC5160_OFS172_SHIFT              = 12
	TMC5160_OFS173_MASK               = 0x00002000
	TMC5160_OFS173_SHIFT              = 13
	TMC5160_OFS174_MASK               = 0x00004000
	TMC5160_OFS174_SHIFT              = 14
	TMC5160_OFS175_MASK               = 0x00008000
	TMC5160_OFS175_SHIFT              = 15
	TMC5160_OFS176_MASK               = 0x00010000
	TMC5160_OFS176_SHIFT              = 16
	TMC5160_OFS177_MASK               = 0x00020000
	TMC5160_OFS177_SHIFT              = 17
	TMC5160_OFS178_MASK               = 0x00040000
	TMC5160_OFS178_SHIFT              = 18
	TMC5160_OFS179_MASK               = 0x00080000
	TMC5160_OFS179_SHIFT              = 19
	TMC5160_OFS180_MASK               = 0x00100000
	TMC5160_OFS180_SHIFT              = 20
	TMC5160_OFS181_MASK               = 0x00200000
	TMC5160_OFS181_SHIFT              = 21
	TMC5160_OFS182_MASK               = 0x00400000
	TMC5160_OFS182_SHIFT              = 22
	TMC5160_OFS183_MASK               = 0x00800000
	TMC5160_OFS183_SHIFT              = 23
	TMC5160_OFS184_MASK               = 0x01000000
	TMC5160_OFS184_SHIFT              = 24
	TMC5160_OFS185_MASK               = 0x02000000
	TMC5160_OFS185_SHIFT              = 25
	TMC5160_OFS186_MASK               = 0x04000000
	TMC5160_OFS186_SHIFT              = 26
	TMC5160_OFS187_MASK               = 0x08000000
	TMC5160_OFS187_SHIFT              = 27
	TMC5160_OFS188_MASK               = 0x10000000
	TMC5160_OFS188_SHIFT              = 28
	TMC5160_OFS189_MASK               = 0x20000000
	TMC5160_OFS189_SHIFT              = 29
	TMC5160_OFS190_MASK               = 0x40000000
	TMC5160_OFS190_SHIFT              = 30
	TMC5160_OFS191_MASK               = 0x80000000
	TMC5160_OFS191_SHIFT              = 31
	TMC5160_OFS192_MASK               = 0x00000001
	TMC5160_OFS192_SHIFT              = 0
	TMC5160_OFS193_MASK               = 0x00000002
	TMC5160_OFS193_SHIFT              = 1
	TMC5160_OFS194_MASK               = 0x00000004
	TMC5160_OFS194_SHIFT              = 2
	TMC5160_OFS195_MASK               = 0x00000008
	TMC5160_OFS195_SHIFT              = 3
	TMC5160_OFS196_MASK               = 0x00000010
	TMC5160_OFS196_SHIFT              = 4
	TMC5160_OFS197_MASK               = 0x00000020
	TMC5160_OFS197_SHIFT              = 5
	TMC5160_OFS198_MASK               = 0x00000040
	TMC5160_OFS198_SHIFT              = 6
	TMC5160_OFS199_MASK               = 0x00000080
	TMC5160_OFS199_SHIFT              = 7
	TMC5160_OFS200_MASK               = 0x00000100
	TMC5160_OFS200_SHIFT              = 8
	TMC5160_OFS201_MASK               = 0x00000200
	TMC5160_OFS201_SHIFT              = 9
	TMC5160_OFS202_MASK               = 0x00000400
	TMC5160_OFS202_SHIFT              = 10
	TMC5160_OFS203_MASK               = 0x00000800
	TMC5160_OFS203_SHIFT              = 11
	TMC5160_OFS204_MASK               = 0x00001000
	TMC5160_OFS204_SHIFT              = 12
	TMC5160_OFS205_MASK               = 0x00002000
	TMC5160_OFS205_SHIFT              = 13
	TMC5160_OFS206_MASK               = 0x00004000
	TMC5160_OFS206_SHIFT              = 14
	TMC5160_OFS207_MASK               = 0x00008000
	TMC5160_OFS207_SHIFT              = 15
	TMC5160_OFS208_MASK               = 0x00010000
	TMC5160_OFS208_SHIFT              = 16
	TMC5160_OFS209_MASK               = 0x00020000
	TMC5160_OFS209_SHIFT              = 17
	TMC5160_OFS210_MASK               = 0x00040000
	TMC5160_OFS210_SHIFT              = 18
	TMC5160_OFS211_MASK               = 0x00080000
	TMC5160_OFS211_SHIFT              = 19
	TMC5160_OFS212_MASK               = 0x00100000
	TMC5160_OFS212_SHIFT              = 20
	TMC5160_OFS213_MASK               = 0x00200000
	TMC5160_OFS213_SHIFT              = 21
	TMC5160_OFS214_MASK               = 0x00400000
	TMC5160_OFS214_SHIFT              = 22
	TMC5160_OFS215_MASK               = 0x00800000
	TMC5160_OFS215_SHIFT              = 23
	TMC5160_OFS216_MASK               = 0x01000000
	TMC5160_OFS216_SHIFT              = 24
	TMC5160_OFS217_MASK               = 0x02000000
	TMC5160_OFS217_SHIFT              = 25
	TMC5160_OFS218_MASK               = 0x04000000
	TMC5160_OFS218_SHIFT              = 26
	TMC5160_OFS219_MASK               = 0x08000000
	TMC5160_OFS219_SHIFT              = 27
	TMC5160_OFS220_MASK               = 0x10000000
	TMC5160_OFS220_SHIFT              = 28
	TMC5160_OFS221_MASK               = 0x20000000
	TMC5160_OFS221_SHIFT              = 29
	TMC5160_OFS222_MASK               = 0x40000000
	TMC5160_OFS222_SHIFT              = 30
	TMC5160_OFS223_MASK               = 0x80000000
	TMC5160_OFS223_SHIFT              = 31
	TMC5160_OFS224_MASK               = 0x00000001
	TMC5160_OFS224_SHIFT              = 0
	TMC5160_OFS225_MASK               = 0x00000002
	TMC5160_OFS225_SHIFT              = 1
	TMC5160_OFS226_MASK               = 0x00000004
	TMC5160_OFS226_SHIFT              = 2
	TMC5160_OFS227_MASK               = 0x00000008
	TMC5160_OFS227_SHIFT              = 3
	TMC5160_OFS228_MASK               = 0x00000010
	TMC5160_OFS228_SHIFT              = 4
	TMC5160_OFS229_MASK               = 0x00000020
	TMC5160_OFS229_SHIFT              = 5
	TMC5160_OFS230_MASK               = 0x00000040
	TMC5160_OFS230_SHIFT              = 6
	TMC5160_OFS231_MASK               = 0x00000080
	TMC5160_OFS231_SHIFT              = 7
	TMC5160_OFS232_MASK               = 0x00000100
	TMC5160_OFS232_SHIFT              = 8
	TMC5160_OFS233_MASK               = 0x00000200
	TMC5160_OFS233_SHIFT              = 9
	TMC5160_OFS234_MASK               = 0x00000400
	TMC5160_OFS234_SHIFT              = 10
	TMC5160_OFS235_MASK               = 0x00000800
	TMC5160_OFS235_SHIFT              = 11
	TMC5160_OFS236_MASK               = 0x00001000
	TMC5160_OFS236_SHIFT              = 12
	TMC5160_OFS237_MASK               = 0x00002000
	TMC5160_OFS237_SHIFT              = 13
	TMC5160_OFS238_MASK               = 0x00004000
	TMC5160_OFS238_SHIFT              = 14
	TMC5160_OFS239_MASK               = 0x00008000
	TMC5160_OFS239_SHIFT              = 15
	TMC5160_OFS240_MASK               = 0x00010000
	TMC5160_OFS240_SHIFT              = 16
	TMC5160_OFS241_MASK               = 0x00020000
	TMC5160_OFS241_SHIFT              = 17
	TMC5160_OFS242_MASK               = 0x00040000
	TMC5160_OFS242_SHIFT              = 18
	TMC5160_OFS243_MASK               = 0x00080000
	TMC5160_OFS243_SHIFT              = 19
	TMC5160_OFS244_MASK               = 0x00100000
	TMC5160_OFS244_SHIFT              = 20
	TMC5160_OFS245_MASK               = 0x00200000
	TMC5160_OFS245_SHIFT              = 21
	TMC5160_OFS246_MASK               = 0x00400000
	TMC5160_OFS246_SHIFT              = 22
	TMC5160_OFS247_MASK               = 0x00800000
	TMC5160_OFS247_SHIFT              = 23
	TMC5160_OFS248_MASK               = 0x01000000
	TMC5160_OFS248_SHIFT              = 24
	TMC5160_OFS249_MASK               = 0x02000000
	TMC5160_OFS249_SHIFT              = 25
	TMC5160_OFS250_MASK               = 0x04000000
	TMC5160_OFS250_SHIFT              = 26
	TMC5160_OFS251_MASK               = 0x08000000
	TMC5160_OFS251_SHIFT              = 27
	TMC5160_OFS252_MASK               = 0x10000000
	TMC5160_OFS252_SHIFT              = 28
	TMC5160_OFS253_MASK               = 0x20000000
	TMC5160_OFS253_SHIFT              = 29
	TMC5160_OFS254_MASK               = 0x40000000
	TMC5160_OFS254_SHIFT              = 30
	TMC5160_OFS255_MASK               = 0x80000000
	TMC5160_OFS255_SHIFT              = 31
	TMC5160_W0_MASK                   = 0x00000003
	TMC5160_W0_SHIFT                  = 0
	TMC5160_W1_MASK                   = 0x0000000C
	TMC5160_W1_SHIFT                  = 2
	TMC5160_W2_MASK                   = 0x00000030
	TMC5160_W2_SHIFT                  = 4
	TMC5160_W3_MASK                   = 0x000000C0
	TMC5160_W3_SHIFT                  = 6
	TMC5160_X1_MASK                   = 0x0000FF00
	TMC5160_X1_SHIFT                  = 8
	TMC5160_X2_MASK                   = 0x00FF0000
	TMC5160_X2_SHIFT                  = 16
	TMC5160_X3_MASK                   = 0xFF000000
	TMC5160_X3_SHIFT                  = 24
	TMC5160_START_SIN_MASK            = 0x000000FF
	TMC5160_START_SIN_SHIFT           = 0
	TMC5160_START_SIN90_MASK          = 0x00FF0000
	TMC5160_START_SIN90_SHIFT         = 16
	TMC5160_MSCNT_MASK                = 0x000003FF
	TMC5160_MSCNT_SHIFT               = 0
	TMC5160_CUR_A_MASK                = 0x000001FF
	TMC5160_CUR_A_SHIFT               = 0
	TMC5160_CUR_B_MASK                = 0x01FF0000
	TMC5160_CUR_B_SHIFT               = 16
	TMC5160_TOFF_MASK                 = 0x0000000F
	TMC5160_TOFF_SHIFT                = 0
	TMC5160_TFD_2__0__MASK            = 0x00000070
	TMC5160_TFD_2__0__SHIFT           = 4
	TMC5160_OFFSET_MASK               = 0x00000780
	TMC5160_OFFSET_SHIFT              = 7
	TMC5160_TFD___MASK                = 0x00000800
	TMC5160_TFD___SHIFT               = 11
	TMC5160_DISFDCC_MASK              = 0x00001000
	TMC5160_DISFDCC_SHIFT             = 12
	TMC5160_CHM_MASK                  = 0x00004000
	TMC5160_CHM_SHIFT                 = 14
	TMC5160_TBL_MASK                  = 0x00018000
	TMC5160_TBL_SHIFT                 = 15
	TMC5160_VHIGHFS_MASK              = 0x00040000
	TMC5160_VHIGHFS_SHIFT             = 18
	TMC5160_VHIGHCHM_MASK             = 0x00080000
	TMC5160_VHIGHCHM_SHIFT            = 19
	TMC5160_TPFD_MASK                 = 0x00F00000
	TMC5160_TPFD_SHIFT                = 20
	TMC5160_MRES_MASK                 = 0x0F000000
	TMC5160_MRES_SHIFT                = 24
	TMC5160_INTPOL_MASK               = 0x10000000
	TMC5160_INTPOL_SHIFT              = 28
	TMC5160_DEDGE_MASK                = 0x20000000
	TMC5160_DEDGE_SHIFT               = 29
	TMC5160_DISS2G_MASK               = 0x40000000
	TMC5160_DISS2G_SHIFT              = 30
	TMC5160_DISS2VS_MASK              = 0x80000000
	TMC5160_DISS2VS_SHIFT             = 31
	TMC5160_RNDTF_MASK                = 0x00002000
	TMC5160_RNDTF_SHIFT               = 13
	TMC5160_VSENSE_MASK               = 0x00020000
	TMC5160_VSENSE_SHIFT              = 17
	TMC5160_HSTRT_MASK                = 0x00000070
	TMC5160_HSTRT_SHIFT               = 4
	TMC5160_HEND_MASK                 = 0x00000780
	TMC5160_HEND_SHIFT                = 7
	TMC5160_SEMIN_MASK                = 0x0000000F
	TMC5160_SEMIN_SHIFT               = 0
	TMC5160_SEUP_MASK                 = 0x00000060
	TMC5160_SEUP_SHIFT                = 5
	TMC5160_SEMAX_MASK                = 0x00000F00
	TMC5160_SEMAX_SHIFT               = 8
	TMC5160_SEDN_MASK                 = 0x00006000
	TMC5160_SEDN_SHIFT                = 13
	TMC5160_SEIMIN_MASK               = 0x00008000
	TMC5160_SEIMIN_SHIFT              = 15
	TMC5160_SGT_MASK                  = 0x007F0000
	TMC5160_SGT_SHIFT                 = 16
	TMC5160_SFILT_MASK                = 0x01000000
	TMC5160_SFILT_SHIFT               = 24
	TMC5160_DC_TIME_MASK              = 0x000003FF
	TMC5160_DC_TIME_SHIFT             = 0
	TMC5160_DC_SG_MASK                = 0x00FF0000
	TMC5160_DC_SG_SHIFT               = 16
	TMC5160_SG_RESULT_MASK            = 0x000003FF
	TMC5160_SG_RESULT_SHIFT           = 0
	TMC5160_S2VSA_MASK                = 0x00001000
	TMC5160_S2VSA_SHIFT               = 12
	TMC5160_S2VSB_MASK                = 0x00002000
	TMC5160_S2VSB_SHIFT               = 13
	TMC5160_STEALTH_MASK              = 0x00004000
	TMC5160_STEALTH_SHIFT             = 14
	TMC5160_FSACTIVE_MASK             = 0x00008000
	TMC5160_FSACTIVE_SHIFT            = 15
	TMC5160_CS_ACTUAL_MASK            = 0x001F0000
	TMC5160_CS_ACTUAL_SHIFT           = 16
	TMC5160_STALLGUARD_MASK           = 0x01000000
	TMC5160_STALLGUARD_SHIFT          = 24
	TMC5160_OT_MASK                   = 0x02000000
	TMC5160_OT_SHIFT                  = 25
	TMC5160_OTPW_MASK                 = 0x04000000
	TMC5160_OTPW_SHIFT                = 26
	TMC5160_S2GA_MASK                 = 0x08000000
	TMC5160_S2GA_SHIFT                = 27
	TMC5160_S2GB_MASK                 = 0x10000000
	TMC5160_S2GB_SHIFT                = 28
	TMC5160_OLA_MASK                  = 0x20000000
	TMC5160_OLA_SHIFT                 = 29
	TMC5160_OLB_MASK                  = 0x40000000
	TMC5160_OLB_SHIFT                 = 30
	TMC5160_STST_MASK                 = 0x80000000
	TMC5160_STST_SHIFT                = 31
	TMC5160_PWM_OFS_MASK              = 0x000000FF
	TMC5160_PWM_OFS_SHIFT             = 0
	TMC5160_PWM_GRAD_MASK             = 0x0000FF00
	TMC5160_PWM_GRAD_SHIFT            = 8
	TMC5160_PWM_FREQ_MASK             = 0x00030000
	TMC5160_PWM_FREQ_SHIFT            = 16
	TMC5160_PWM_AUTOSCALE_MASK        = 0x00040000
	TMC5160_PWM_AUTOSCALE_SHIFT       = 18
	TMC5160_PWM_AUTOGRAD_MASK         = 0x00080000
	TMC5160_PWM_AUTOGRAD_SHIFT        = 19
	TMC5160_FREEWHEEL_MASK            = 0x00300000
	TMC5160_FREEWHEEL_SHIFT           = 20
	TMC5160_PWM_REG_MASK              = 0x0F000000
	TMC5160_PWM_REG_SHIFT             = 24
	TMC5160_PWM_LIM_MASK              = 0xF0000000
	TMC5160_PWM_LIM_SHIFT             = 28
	TMC5160_PWM_SCALE_SUM_MASK        = 0x000000FF
	TMC5160_PWM_SCALE_SUM_SHIFT       = 0
	TMC5160_PWM_SCALE_AUTO_MASK       = 0x01FF0000
	TMC5160_PWM_SCALE_AUTO_SHIFT      = 16
	TMC5160_PWM_OFS_AUTO_MASK         = 0x000000FF
	TMC5160_PWM_OFS_AUTO_SHIFT        = 0
	TMC5160_PWM_GRAD_AUTO_MASK        = 0x00FF0000
	TMC5160_PWM_GRAD_AUTO_SHIFT       = 16
	TMC5160_LOST_STEPS_MASK           = 0x000FFFFF
	TMC5160_LOST_STEPS_SHIFT          = 0
)

// Register Fields
var (
	TMC5160_SPI_STATUS_RESET_FLAG_FIELD = RegisterField{
		Mask:     TMC5160_SPI_STATUS_RESET_FLAG_MASK,
		Shift:    TMC5160_SPI_STATUS_RESET_FLAG_SHIFT,
		Register: TMC5160_GSTAT,
		IsSigned: false,
	}
	TMC5160_SPI_STATUS_DRIVER_ERROR_FIELD = RegisterField{
		Mask:     TMC5160_SPI_STATUS_DRIVER_ERROR_MASK,
		Shift:    TMC5160_SPI_STATUS_DRIVER_ERROR_SHIFT,
		Register: TMC5160_GSTAT,
		IsSigned: false,
	}
	TMC5160_SPI_STATUS_SG2_FIELD = RegisterField{
		Mask:     TMC5160_SPI_STATUS_SG2_MASK,
		Shift:    TMC5160_SPI_STATUS_SG2_SHIFT,
		Register: TMC5160_DRV_STATUS,
		IsSigned: false,
	}
	TMC5160_SPI_STATUS_STANDSTILL_FIELD = RegisterField{
		Mask:     TMC5160_SPI_STATUS_STANDSTILL_MASK,
		Shift:    TMC5160_SPI_STATUS_STANDSTILL_SHIFT,
		Register: TMC5160_DRV_STATUS,
		IsSigned: false,
	}
	TMC5160_SPI_STATUS_VELOCITY_REACHED_FIELD = RegisterField{
		Mask:     TMC5160_SPI_STATUS_VELOCITY_REACHED_MASK,
		Shift:    TMC5160_SPI_STATUS_VELOCITY_REACHED_SHIFT,
		Register: TMC5160_RAMPSTAT,
		IsSigned: false,
	}
	TMC5160_SPI_STATUS_POSITION_REACHED_FIELD = RegisterField{
		Mask:     TMC5160_SPI_STATUS_POSITION_REACHED_MASK,
		Shift:    TMC5160_SPI_STATUS_POSITION_REACHED_SHIFT,
		Register: TMC5160_RAMPSTAT,
		IsSigned: false,
	}
	TMC5160_SPI_STATUS_STATUS_STOP_L_FIELD = RegisterField{
		Mask:     TMC5160_SPI_STATUS_STATUS_STOP_L_MASK,
		Shift:    TMC5160_SPI_STATUS_STATUS_STOP_L_SHIFT,
		Register: TMC5160_RAMPSTAT,
		IsSigned: false,
	}
	TMC5160_SPI_STATUS_STATUS_STOP_R_FIELD = RegisterField{
		Mask:     TMC5160_SPI_STATUS_STATUS_STOP_R_MASK,
		Shift:    TMC5160_SPI_STATUS_STATUS_STOP_R_SHIFT,
		Register: TMC5160_RAMPSTAT,
		IsSigned: false,
	}
	TMC5160_RECALIBRATE_FIELD = RegisterField{
		Mask:     TMC5160_RECALIBRATE_MASK,
		Shift:    TMC5160_RECALIBRATE_SHIFT,
		Register: TMC5160_GCONF,
		IsSigned: false,
	}
	TMC5160_FASTSTANDSTILL_FIELD = RegisterField{
		Mask:     TMC5160_FASTSTANDSTILL_MASK,
		Shift:    TMC5160_FASTSTANDSTILL_SHIFT,
		Register: TMC5160_GCONF,
		IsSigned: false,
	}
	TMC5160_EN_PWM_MODE_FIELD = RegisterField{
		Mask:     TMC5160_EN_PWM_MODE_MASK,
		Shift:    TMC5160_EN_PWM_MODE_SHIFT,
		Register: TMC5160_GCONF,
		IsSigned: false,
	}
	TMC5160_MULTISTEP_FILT_FIELD = RegisterField{
		Mask:     TMC5160_MULTISTEP_FILT_MASK,
		Shift:    TMC5160_MULTISTEP_FILT_SHIFT,
		Register: TMC5160_GCONF,
		IsSigned: false,
	}
	TMC5160_SHAFT_FIELD = RegisterField{
		Mask:     TMC5160_SHAFT_MASK,
		Shift:    TMC5160_SHAFT_SHIFT,
		Register: TMC5160_GCONF,
		IsSigned: false,
	}
	TMC5160_DIAG0_ERROR__ONLY_WITH_SD_MODE_1__FIELD = RegisterField{
		Mask:     TMC5160_DIAG0_ERROR__ONLY_WITH_SD_MODE_1__MASK,
		Shift:    TMC5160_DIAG0_ERROR__ONLY_WITH_SD_MODE_1__SHIFT,
		Register: TMC5160_GCONF,
		IsSigned: false,
	}
	TMC5160_DIAG0_OTPW__ONLY_WITH_SD_MODE_1__FIELD = RegisterField{
		Mask:     TMC5160_DIAG0_OTPW__ONLY_WITH_SD_MODE_1__MASK,
		Shift:    TMC5160_DIAG0_OTPW__ONLY_WITH_SD_MODE_1__SHIFT,
		Register: TMC5160_GCONF,
		IsSigned: false,
	}
	TMC5160_DIAG0_STALL_FIELD = RegisterField{
		Mask:     TMC5160_DIAG0_STALL_MASK,
		Shift:    TMC5160_DIAG0_STALL_SHIFT,
		Register: TMC5160_GCONF,
		IsSigned: false,
	}
	TMC5160_DIAG1_STALL_FIELD = RegisterField{
		Mask:     TMC5160_DIAG1_STALL_MASK,
		Shift:    TMC5160_DIAG1_STALL_SHIFT,
		Register: TMC5160_GCONF,
		IsSigned: false,
	}
	TMC5160_DIAG1_INDEX_FIELD = RegisterField{
		Mask:     TMC5160_DIAG1_INDEX_MASK,
		Shift:    TMC5160_DIAG1_INDEX_SHIFT,
		Register: TMC5160_GCONF,
		IsSigned: false,
	}
	TMC5160_DIAG1_ONSTATE_FIELD = RegisterField{
		Mask:     TMC5160_DIAG1_ONSTATE_MASK,
		Shift:    TMC5160_DIAG1_ONSTATE_SHIFT,
		Register: TMC5160_GCONF,
		IsSigned: false,
	}
	TMC5160_DIAG1_STEPS_SKIPPED_FIELD = RegisterField{
		Mask:     TMC5160_DIAG1_STEPS_SKIPPED_MASK,
		Shift:    TMC5160_DIAG1_STEPS_SKIPPED_SHIFT,
		Register: TMC5160_GCONF,
		IsSigned: false,
	}
	TMC5160_DIAG0_INT_PUSHPULL_FIELD = RegisterField{
		Mask:     TMC5160_DIAG0_INT_PUSHPULL_MASK,
		Shift:    TMC5160_DIAG0_INT_PUSHPULL_SHIFT,
		Register: TMC5160_GCONF,
		IsSigned: false,
	}
	TMC5160_DIAG1_POSCOMP_PUSHPULL_FIELD = RegisterField{
		Mask:     TMC5160_DIAG1_POSCOMP_PUSHPULL_MASK,
		Shift:    TMC5160_DIAG1_POSCOMP_PUSHPULL_SHIFT,
		Register: TMC5160_GCONF,
		IsSigned: false,
	}
	TMC5160_SMALL_HYSTERESIS_FIELD = RegisterField{
		Mask:     TMC5160_SMALL_HYSTERESIS_MASK,
		Shift:    TMC5160_SMALL_HYSTERESIS_SHIFT,
		Register: TMC5160_GCONF,
		IsSigned: false,
	}
	TMC5160_STOP_ENABLE_FIELD = RegisterField{
		Mask:     TMC5160_STOP_ENABLE_MASK,
		Shift:    TMC5160_STOP_ENABLE_SHIFT,
		Register: TMC5160_GCONF,
		IsSigned: false,
	}
	TMC5160_DIRECT_MODE_FIELD = RegisterField{
		Mask:     TMC5160_DIRECT_MODE_MASK,
		Shift:    TMC5160_DIRECT_MODE_SHIFT,
		Register: TMC5160_GCONF,
		IsSigned: false,
	}
	TMC5160_TEST_MODE_FIELD = RegisterField{
		Mask:     TMC5160_TEST_MODE_MASK,
		Shift:    TMC5160_TEST_MODE_SHIFT,
		Register: TMC5160_GCONF,
		IsSigned: false,
	}
	TMC5160_DIAG0_STEP_FIELD = RegisterField{
		Mask:     TMC5160_DIAG0_STEP_MASK,
		Shift:    TMC5160_DIAG0_STEP_SHIFT,
		Register: TMC5160_GCONF,
		IsSigned: false,
	}
	TMC5160_DIAG1_DIR_FIELD = RegisterField{
		Mask:     TMC5160_DIAG1_DIR_MASK,
		Shift:    TMC5160_DIAG1_DIR_SHIFT,
		Register: TMC5160_GCONF,
		IsSigned: false,
	}
	TMC5160_RESET_FIELD = RegisterField{
		Mask:     TMC5160_RESET_MASK,
		Shift:    TMC5160_RESET_SHIFT,
		Register: TMC5160_GSTAT,
		IsSigned: false,
	}
	TMC5160_DRV_ERR_FIELD = RegisterField{
		Mask:     TMC5160_DRV_ERR_MASK,
		Shift:    TMC5160_DRV_ERR_SHIFT,
		Register: TMC5160_GSTAT,
		IsSigned: false,
	}
	TMC5160_UV_CP_FIELD = RegisterField{
		Mask:     TMC5160_UV_CP_MASK,
		Shift:    TMC5160_UV_CP_SHIFT,
		Register: TMC5160_GSTAT,
		IsSigned: false,
	}
	TMC5160_IFCNT_FIELD = RegisterField{
		Mask:     TMC5160_IFCNT_MASK,
		Shift:    TMC5160_IFCNT_SHIFT,
		Register: TMC5160_IFCNT,
		IsSigned: false,
	}
	TMC5160_SLAVEADDR_FIELD = RegisterField{
		Mask:     TMC5160_SLAVEADDR_MASK,
		Shift:    TMC5160_SLAVEADDR_SHIFT,
		Register: TMC5160_SLAVECONF,
		IsSigned: false,
	}
	TMC5160_SENDDELAY_FIELD = RegisterField{
		Mask:     TMC5160_SENDDELAY_MASK,
		Shift:    TMC5160_SENDDELAY_SHIFT,
		Register: TMC5160_SLAVECONF,
		IsSigned: false,
	}
	TMC5160_REFL_STEP_FIELD = RegisterField{
		Mask:     TMC5160_REFL_STEP_MASK,
		Shift:    TMC5160_REFL_STEP_SHIFT,
		Register: TMC5160_INP_OUT,
		IsSigned: false,
	}
	TMC5160_REFR_DIR_FIELD = RegisterField{
		Mask:     TMC5160_REFR_DIR_MASK,
		Shift:    TMC5160_REFR_DIR_SHIFT,
		Register: TMC5160_INP_OUT,
		IsSigned: false,
	}
	TMC5160_ENCB_DCEN_CFG4_FIELD = RegisterField{
		Mask:     TMC5160_ENCB_DCEN_CFG4_MASK,
		Shift:    TMC5160_ENCB_DCEN_CFG4_SHIFT,
		Register: TMC5160_INP_OUT,
		IsSigned: false,
	}
	TMC5160_ENCA_DCIN_CFG5_FIELD = RegisterField{
		Mask:     TMC5160_ENCA_DCIN_CFG5_MASK,
		Shift:    TMC5160_ENCA_DCIN_CFG5_SHIFT,
		Register: TMC5160_INP_OUT,
		IsSigned: false,
	}
	TMC5160_DRV_ENN_CFG6_FIELD = RegisterField{
		Mask:     TMC5160_DRV_ENN_CFG6_MASK,
		Shift:    TMC5160_DRV_ENN_CFG6_SHIFT,
		Register: TMC5160_INP_OUT,
		IsSigned: false,
	}
	TMC5160_ENC_N_DCO_FIELD = RegisterField{
		Mask:     TMC5160_ENC_N_DCO_MASK,
		Shift:    TMC5160_ENC_N_DCO_SHIFT,
		Register: TMC5160_INP_OUT,
		IsSigned: false,
	}
	TMC5160_SD_MODE_FIELD = RegisterField{
		Mask:     TMC5160_SD_MODE_MASK,
		Shift:    TMC5160_SD_MODE_SHIFT,
		Register: TMC5160_INP_OUT,
		IsSigned: false,
	}
	TMC5160_SWCOMP_IN_FIELD = RegisterField{
		Mask:     TMC5160_SWCOMP_IN_MASK,
		Shift:    TMC5160_SWCOMP_IN_SHIFT,
		Register: TMC5160_INP_OUT,
		IsSigned: false,
	}
	TMC5160_VERSION_FIELD = RegisterField{
		Mask:     TMC5160_VERSION_MASK,
		Shift:    TMC5160_VERSION_SHIFT,
		Register: TMC5160_INP_OUT,
		IsSigned: false,
	}
	TMC5160_OUTPUT_PIN_POLARITY_FIELD = RegisterField{
		Mask:     TMC5160_OUTPUT_PIN_POLARITY_MASK,
		Shift:    TMC5160_OUTPUT_PIN_POLARITY_SHIFT,
		Register: TMC5160_INP_OUT,
		IsSigned: false,
	}
	TMC5160_X_COMPARE_FIELD = RegisterField{
		Mask:     TMC5160_X_COMPARE_MASK,
		Shift:    TMC5160_X_COMPARE_SHIFT,
		Register: TMC5160_X_COMPARE,
		IsSigned: false,
	}
	TMC5160_OTPBIT_FIELD = RegisterField{
		Mask:     TMC5160_OTPBIT_MASK,
		Shift:    TMC5160_OTPBIT_SHIFT,
		Register: TMC5160_OTP_PROG,
		IsSigned: false,
	}
	TMC5160_OTPBYTE_FIELD = RegisterField{
		Mask:     TMC5160_OTPBYTE_MASK,
		Shift:    TMC5160_OTPBYTE_SHIFT,
		Register: TMC5160_OTP_PROG,
		IsSigned: false,
	}
	TMC5160_OTPMAGIC_FIELD = RegisterField{
		Mask:     TMC5160_OTPMAGIC_MASK,
		Shift:    TMC5160_OTPMAGIC_SHIFT,
		Register: TMC5160_OTP_PROG,
		IsSigned: false,
	}
	TMC5160_OTP_TBL_FIELD = RegisterField{
		Mask:     TMC5160_OTP_TBL_MASK,
		Shift:    TMC5160_OTP_TBL_SHIFT,
		Register: TMC5160_OTP_READ,
		IsSigned: false,
	}
	TMC5160_OTP_BBM_FIELD = RegisterField{
		Mask:     TMC5160_OTP_BBM_MASK,
		Shift:    TMC5160_OTP_BBM_SHIFT,
		Register: TMC5160_OTP_READ,
		IsSigned: false,
	}
	TMC5160_OTP_S2_LEVEL_FIELD = RegisterField{
		Mask:     TMC5160_OTP_S2_LEVEL_MASK,
		Shift:    TMC5160_OTP_S2_LEVEL_SHIFT,
		Register: TMC5160_OTP_READ,
		IsSigned: false,
	}
	TMC5160_OTP_FCLKTRIM_FIELD = RegisterField{
		Mask:     TMC5160_OTP_FCLKTRIM_MASK,
		Shift:    TMC5160_OTP_FCLKTRIM_SHIFT,
		Register: TMC5160_OTP_READ,
		IsSigned: false,
	}
	TMC5160_FCLKTRIM_FIELD = RegisterField{
		Mask:     TMC5160_FCLKTRIM_MASK,
		Shift:    TMC5160_FCLKTRIM_SHIFT,
		Register: TMC5160_FACTORY_CONF,
		IsSigned: false,
	}
	TMC5160_S2VS_LEVEL_FIELD = RegisterField{
		Mask:     TMC5160_S2VS_LEVEL_MASK,
		Shift:    TMC5160_S2VS_LEVEL_SHIFT,
		Register: TMC5160_SHORT_CONF,
		IsSigned: false,
	}
	TMC5160_S2GND_LEVEL_FIELD = RegisterField{
		Mask:     TMC5160_S2GND_LEVEL_MASK,
		Shift:    TMC5160_S2GND_LEVEL_SHIFT,
		Register: TMC5160_SHORT_CONF,
		IsSigned: false,
	}
	TMC5160_SHORTFILTER_FIELD = RegisterField{
		Mask:     TMC5160_SHORTFILTER_MASK,
		Shift:    TMC5160_SHORTFILTER_SHIFT,
		Register: TMC5160_SHORT_CONF,
		IsSigned: false,
	}
	TMC5160_SHORTDELAY_FIELD = RegisterField{
		Mask:     TMC5160_SHORTDELAY_MASK,
		Shift:    TMC5160_SHORTDELAY_SHIFT,
		Register: TMC5160_SHORT_CONF,
		IsSigned: false,
	}
	TMC5160_BBMTIME_FIELD = RegisterField{
		Mask:     TMC5160_BBMTIME_MASK,
		Shift:    TMC5160_BBMTIME_SHIFT,
		Register: TMC5160_DRV_CONF,
		IsSigned: false,
	}
	TMC5160_BBMCLKS_FIELD = RegisterField{
		Mask:     TMC5160_BBMCLKS_MASK,
		Shift:    TMC5160_BBMCLKS_SHIFT,
		Register: TMC5160_DRV_CONF,
		IsSigned: false,
	}
	TMC5160_OTSELECT_FIELD = RegisterField{
		Mask:     TMC5160_OTSELECT_MASK,
		Shift:    TMC5160_OTSELECT_SHIFT,
		Register: TMC5160_DRV_CONF,
		IsSigned: false,
	}
	TMC5160_DRVSTRENGTH_FIELD = RegisterField{
		Mask:     TMC5160_DRVSTRENGTH_MASK,
		Shift:    TMC5160_DRVSTRENGTH_SHIFT,
		Register: TMC5160_DRV_CONF,
		IsSigned: false,
	}
	TMC5160_FILT_ISENSE_FIELD = RegisterField{
		Mask:     TMC5160_FILT_ISENSE_MASK,
		Shift:    TMC5160_FILT_ISENSE_SHIFT,
		Register: TMC5160_DRV_CONF,
		IsSigned: false,
	}
	TMC5160_GLOBAL_SCALER_FIELD = RegisterField{
		Mask:     TMC5160_GLOBAL_SCALER_MASK,
		Shift:    TMC5160_GLOBAL_SCALER_SHIFT,
		Register: TMC5160_GLOBAL_SCALER,
		IsSigned: false,
	}
	TMC5160_OFFSET_READ_A_FIELD = RegisterField{
		Mask:     TMC5160_OFFSET_READ_A_MASK,
		Shift:    TMC5160_OFFSET_READ_A_SHIFT,
		Register: TMC5160_OFFSET_READ,
		IsSigned: true,
	}
	TMC5160_OFFSET_READ_B_FIELD = RegisterField{
		Mask:     TMC5160_OFFSET_READ_B_MASK,
		Shift:    TMC5160_OFFSET_READ_B_SHIFT,
		Register: TMC5160_OFFSET_READ,
		IsSigned: true,
	}
	TMC5160_IHOLD_FIELD = RegisterField{
		Mask:     TMC5160_IHOLD_MASK,
		Shift:    TMC5160_IHOLD_SHIFT,
		Register: TMC5160_IHOLD_IRUN,
		IsSigned: false,
	}
	TMC5160_IRUN_FIELD = RegisterField{
		Mask:     TMC5160_IRUN_MASK,
		Shift:    TMC5160_IRUN_SHIFT,
		Register: TMC5160_IHOLD_IRUN,
		IsSigned: false,
	}
	TMC5160_IHOLDDELAY_FIELD = RegisterField{
		Mask:     TMC5160_IHOLDDELAY_MASK,
		Shift:    TMC5160_IHOLDDELAY_SHIFT,
		Register: TMC5160_IHOLD_IRUN,
		IsSigned: false,
	}
	TMC5160_TPOWERDOWN_FIELD = RegisterField{
		Mask:     TMC5160_TPOWERDOWN_MASK,
		Shift:    TMC5160_TPOWERDOWN_SHIFT,
		Register: TMC5160_TPOWERDOWN,
		IsSigned: false,
	}
	TMC5160_TSTEP_FIELD = RegisterField{
		Mask:     TMC5160_TSTEP_MASK,
		Shift:    TMC5160_TSTEP_SHIFT,
		Register: TMC5160_TSTEP,
		IsSigned: false,
	}
	TMC5160_TPWMTHRS_FIELD = RegisterField{
		Mask:     TMC5160_TPWMTHRS_MASK,
		Shift:    TMC5160_TPWMTHRS_SHIFT,
		Register: TMC5160_TPWMTHRS,
		IsSigned: false,
	}
	TMC5160_TCOOLTHRS_FIELD = RegisterField{
		Mask:     TMC5160_TCOOLTHRS_MASK,
		Shift:    TMC5160_TCOOLTHRS_SHIFT,
		Register: TMC5160_TCOOLTHRS,
		IsSigned: false,
	}
	TMC5160_THIGH_FIELD = RegisterField{
		Mask:     TMC5160_THIGH_MASK,
		Shift:    TMC5160_THIGH_SHIFT,
		Register: TMC5160_THIGH,
		IsSigned: false,
	}
	TMC5160_RAMPMODE_FIELD = RegisterField{
		Mask:     TMC5160_RAMPMODE_MASK,
		Shift:    TMC5160_RAMPMODE_SHIFT,
		Register: TMC5160_RAMPMODE,
		IsSigned: false,
	}
	TMC5160_XACTUAL_FIELD = RegisterField{
		Mask:     TMC5160_XACTUAL_MASK,
		Shift:    TMC5160_XACTUAL_SHIFT,
		Register: TMC5160_XACTUAL,
		IsSigned: true,
	}
	TMC5160_VACTUAL_FIELD = RegisterField{
		Mask:     TMC5160_VACTUAL_MASK,
		Shift:    TMC5160_VACTUAL_SHIFT,
		Register: TMC5160_VACTUAL,
		IsSigned: true,
	}
	TMC5160_VSTART_FIELD = RegisterField{
		Mask:     TMC5160_VSTART_MASK,
		Shift:    TMC5160_VSTART_SHIFT,
		Register: TMC5160_VSTART,
		IsSigned: false,
	}
	TMC5160_A1_FIELD = RegisterField{
		Mask:     TMC5160_A1_MASK,
		Shift:    TMC5160_A1_SHIFT,
		Register: TMC5160_A1,
		IsSigned: false,
	}
	TMC5160_V1__FIELD = RegisterField{
		Mask:     TMC5160_V1__MASK,
		Shift:    TMC5160_V1__SHIFT,
		Register: TMC5160_V1,
		IsSigned: false,
	}
	TMC5160_AMAX_FIELD = RegisterField{
		Mask:     TMC5160_AMAX_MASK,
		Shift:    TMC5160_AMAX_SHIFT,
		Register: TMC5160_AMAX,
		IsSigned: false,
	}
	TMC5160_VMAX_FIELD = RegisterField{
		Mask:     TMC5160_VMAX_MASK,
		Shift:    TMC5160_VMAX_SHIFT,
		Register: TMC5160_VMAX,
		IsSigned: false,
	}
	TMC5160_DMAX_FIELD = RegisterField{
		Mask:     TMC5160_DMAX_MASK,
		Shift:    TMC5160_DMAX_SHIFT,
		Register: TMC5160_DMAX,
		IsSigned: false,
	}
	TMC5160_D1_FIELD = RegisterField{
		Mask:     TMC5160_D1_MASK,
		Shift:    TMC5160_D1_SHIFT,
		Register: TMC5160_D1,
		IsSigned: false,
	}
	TMC5160_VSTOP_FIELD = RegisterField{
		Mask:     TMC5160_VSTOP_MASK,
		Shift:    TMC5160_VSTOP_SHIFT,
		Register: TMC5160_VSTOP,
		IsSigned: false,
	}
	TMC5160_TZEROWAIT_FIELD = RegisterField{
		Mask:     TMC5160_TZEROWAIT_MASK,
		Shift:    TMC5160_TZEROWAIT_SHIFT,
		Register: TMC5160_TZEROWAIT,
		IsSigned: false,
	}
	TMC5160_XTARGET_FIELD = RegisterField{
		Mask:     TMC5160_XTARGET_MASK,
		Shift:    TMC5160_XTARGET_SHIFT,
		Register: TMC5160_XTARGET,
		IsSigned: true,
	}
	TMC5160_VDCMIN_FIELD = RegisterField{
		Mask:     TMC5160_VDCMIN_MASK,
		Shift:    TMC5160_VDCMIN_SHIFT,
		Register: TMC5160_VDCMIN,
		IsSigned: false,
	}
	TMC5160_STOP_L_ENABLE_FIELD = RegisterField{
		Mask:     TMC5160_STOP_L_ENABLE_MASK,
		Shift:    TMC5160_STOP_L_ENABLE_SHIFT,
		Register: TMC5160_SWMODE,
		IsSigned: false,
	}
	TMC5160_STOP_R_ENABLE_FIELD = RegisterField{
		Mask:     TMC5160_STOP_R_ENABLE_MASK,
		Shift:    TMC5160_STOP_R_ENABLE_SHIFT,
		Register: TMC5160_SWMODE,
		IsSigned: false,
	}
	TMC5160_POL_STOP_L_FIELD = RegisterField{
		Mask:     TMC5160_POL_STOP_L_MASK,
		Shift:    TMC5160_POL_STOP_L_SHIFT,
		Register: TMC5160_SWMODE,
		IsSigned: false,
	}
	TMC5160_POL_STOP_R_FIELD = RegisterField{
		Mask:     TMC5160_POL_STOP_R_MASK,
		Shift:    TMC5160_POL_STOP_R_SHIFT,
		Register: TMC5160_SWMODE,
		IsSigned: false,
	}
	TMC5160_SWAP_LR_FIELD = RegisterField{
		Mask:     TMC5160_SWAP_LR_MASK,
		Shift:    TMC5160_SWAP_LR_SHIFT,
		Register: TMC5160_SWMODE,
		IsSigned: false,
	}
	TMC5160_LATCH_L_ACTIVE_FIELD = RegisterField{
		Mask:     TMC5160_LATCH_L_ACTIVE_MASK,
		Shift:    TMC5160_LATCH_L_ACTIVE_SHIFT,
		Register: TMC5160_SWMODE,
		IsSigned: false,
	}
	TMC5160_LATCH_L_INACTIVE_FIELD = RegisterField{
		Mask:     TMC5160_LATCH_L_INACTIVE_MASK,
		Shift:    TMC5160_LATCH_L_INACTIVE_SHIFT,
		Register: TMC5160_SWMODE,
		IsSigned: false,
	}
	TMC5160_LATCH_R_ACTIVE_FIELD = RegisterField{
		Mask:     TMC5160_LATCH_R_ACTIVE_MASK,
		Shift:    TMC5160_LATCH_R_ACTIVE_SHIFT,
		Register: TMC5160_SWMODE,
		IsSigned: false,
	}
	TMC5160_LATCH_R_INACTIVE_FIELD = RegisterField{
		Mask:     TMC5160_LATCH_R_INACTIVE_MASK,
		Shift:    TMC5160_LATCH_R_INACTIVE_SHIFT,
		Register: TMC5160_SWMODE,
		IsSigned: false,
	}
	TMC5160_EN_LATCH_ENCODER_FIELD = RegisterField{
		Mask:     TMC5160_EN_LATCH_ENCODER_MASK,
		Shift:    TMC5160_EN_LATCH_ENCODER_SHIFT,
		Register: TMC5160_SWMODE,
		IsSigned: false,
	}
	TMC5160_SG_STOP_FIELD = RegisterField{
		Mask:     TMC5160_SG_STOP_MASK,
		Shift:    TMC5160_SG_STOP_SHIFT,
		Register: TMC5160_SWMODE,
		IsSigned: false,
	}
	TMC5160_EN_SOFTSTOP_FIELD = RegisterField{
		Mask:     TMC5160_EN_SOFTSTOP_MASK,
		Shift:    TMC5160_EN_SOFTSTOP_SHIFT,
		Register: TMC5160_SWMODE,
		IsSigned: false,
	}
	TMC5160_STATUS_STOP_L_FIELD = RegisterField{
		Mask:     TMC5160_STATUS_STOP_L_MASK,
		Shift:    TMC5160_STATUS_STOP_L_SHIFT,
		Register: TMC5160_RAMPSTAT,
		IsSigned: false,
	}
	TMC5160_STATUS_STOP_R_FIELD = RegisterField{
		Mask:     TMC5160_STATUS_STOP_R_MASK,
		Shift:    TMC5160_STATUS_STOP_R_SHIFT,
		Register: TMC5160_RAMPSTAT,
		IsSigned: false,
	}
	TMC5160_STATUS_LATCH_L_FIELD = RegisterField{
		Mask:     TMC5160_STATUS_LATCH_L_MASK,
		Shift:    TMC5160_STATUS_LATCH_L_SHIFT,
		Register: TMC5160_RAMPSTAT,
		IsSigned: false,
	}
	TMC5160_STATUS_LATCH_R_FIELD = RegisterField{
		Mask:     TMC5160_STATUS_LATCH_R_MASK,
		Shift:    TMC5160_STATUS_LATCH_R_SHIFT,
		Register: TMC5160_RAMPSTAT,
		IsSigned: false,
	}
	TMC5160_EVENT_STOP_L_FIELD = RegisterField{
		Mask:     TMC5160_EVENT_STOP_L_MASK,
		Shift:    TMC5160_EVENT_STOP_L_SHIFT,
		Register: TMC5160_RAMPSTAT,
		IsSigned: false,
	}
	TMC5160_EVENT_STOP_R_FIELD = RegisterField{
		Mask:     TMC5160_EVENT_STOP_R_MASK,
		Shift:    TMC5160_EVENT_STOP_R_SHIFT,
		Register: TMC5160_RAMPSTAT,
		IsSigned: false,
	}
	TMC5160_EVENT_STOP_SG_FIELD = RegisterField{
		Mask:     TMC5160_EVENT_STOP_SG_MASK,
		Shift:    TMC5160_EVENT_STOP_SG_SHIFT,
		Register: TMC5160_RAMPSTAT,
		IsSigned: false,
	}
	TMC5160_EVENT_POS_REACHED_FIELD = RegisterField{
		Mask:     TMC5160_EVENT_POS_REACHED_MASK,
		Shift:    TMC5160_EVENT_POS_REACHED_SHIFT,
		Register: TMC5160_RAMPSTAT,
		IsSigned: false,
	}
	TMC5160_VELOCITY_REACHED_FIELD = RegisterField{
		Mask:     TMC5160_VELOCITY_REACHED_MASK,
		Shift:    TMC5160_VELOCITY_REACHED_SHIFT,
		Register: TMC5160_RAMPSTAT,
		IsSigned: false,
	}
	TMC5160_POSITION_REACHED_FIELD = RegisterField{
		Mask:     TMC5160_POSITION_REACHED_MASK,
		Shift:    TMC5160_POSITION_REACHED_SHIFT,
		Register: TMC5160_RAMPSTAT,
		IsSigned: false,
	}
	TMC5160_VZERO_FIELD = RegisterField{
		Mask:     TMC5160_VZERO_MASK,
		Shift:    TMC5160_VZERO_SHIFT,
		Register: TMC5160_RAMPSTAT,
		IsSigned: false,
	}
	TMC5160_T_ZEROWAIT_ACTIVE_FIELD = RegisterField{
		Mask:     TMC5160_T_ZEROWAIT_ACTIVE_MASK,
		Shift:    TMC5160_T_ZEROWAIT_ACTIVE_SHIFT,
		Register: TMC5160_RAMPSTAT,
		IsSigned: false,
	}
	TMC5160_SECOND_MOVE_FIELD = RegisterField{
		Mask:     TMC5160_SECOND_MOVE_MASK,
		Shift:    TMC5160_SECOND_MOVE_SHIFT,
		Register: TMC5160_RAMPSTAT,
		IsSigned: false,
	}
	TMC5160_STATUS_SG_FIELD = RegisterField{
		Mask:     TMC5160_STATUS_SG_MASK,
		Shift:    TMC5160_STATUS_SG_SHIFT,
		Register: TMC5160_RAMPSTAT,
		IsSigned: false,
	}
	TMC5160_XLATCH_FIELD = RegisterField{
		Mask:     TMC5160_XLATCH_MASK,
		Shift:    TMC5160_XLATCH_SHIFT,
		Register: TMC5160_XLATCH,
		IsSigned: false,
	}
	TMC5160_POL_A_FIELD = RegisterField{
		Mask:     TMC5160_POL_A_MASK,
		Shift:    TMC5160_POL_A_SHIFT,
		Register: TMC5160_ENCMODE,
		IsSigned: false,
	}
	TMC5160_POL_B_FIELD = RegisterField{
		Mask:     TMC5160_POL_B_MASK,
		Shift:    TMC5160_POL_B_SHIFT,
		Register: TMC5160_ENCMODE,
		IsSigned: false,
	}
	TMC5160_POL_N_FIELD = RegisterField{
		Mask:     TMC5160_POL_N_MASK,
		Shift:    TMC5160_POL_N_SHIFT,
		Register: TMC5160_ENCMODE,
		IsSigned: false,
	}
	TMC5160_IGNORE_AB_FIELD = RegisterField{
		Mask:     TMC5160_IGNORE_AB_MASK,
		Shift:    TMC5160_IGNORE_AB_SHIFT,
		Register: TMC5160_ENCMODE,
		IsSigned: false,
	}
	TMC5160_CLR_CONT_FIELD = RegisterField{
		Mask:     TMC5160_CLR_CONT_MASK,
		Shift:    TMC5160_CLR_CONT_SHIFT,
		Register: TMC5160_ENCMODE,
		IsSigned: false,
	}
	TMC5160_CLR_ONCE_FIELD = RegisterField{
		Mask:     TMC5160_CLR_ONCE_MASK,
		Shift:    TMC5160_CLR_ONCE_SHIFT,
		Register: TMC5160_ENCMODE,
		IsSigned: false,
	}
	TMC5160_POS_EDGE_NEG_EDGE_FIELD = RegisterField{
		Mask:     TMC5160_POS_EDGE_NEG_EDGE_MASK,
		Shift:    TMC5160_POS_EDGE_NEG_EDGE_SHIFT,
		Register: TMC5160_ENCMODE,
		IsSigned: false,
	}
	TMC5160_CLR_ENC_X_FIELD = RegisterField{
		Mask:     TMC5160_CLR_ENC_X_MASK,
		Shift:    TMC5160_CLR_ENC_X_SHIFT,
		Register: TMC5160_ENCMODE,
		IsSigned: false,
	}
	TMC5160_LATCH_X_ACT_FIELD = RegisterField{
		Mask:     TMC5160_LATCH_X_ACT_MASK,
		Shift:    TMC5160_LATCH_X_ACT_SHIFT,
		Register: TMC5160_ENCMODE,
		IsSigned: false,
	}
	TMC5160_ENC_SEL_DECIMAL_FIELD = RegisterField{
		Mask:     TMC5160_ENC_SEL_DECIMAL_MASK,
		Shift:    TMC5160_ENC_SEL_DECIMAL_SHIFT,
		Register: TMC5160_ENCMODE,
		IsSigned: false,
	}
	TMC5160_X_ENC_FIELD = RegisterField{
		Mask:     TMC5160_X_ENC_MASK,
		Shift:    TMC5160_X_ENC_SHIFT,
		Register: TMC5160_XENC,
		IsSigned: true,
	}
	TMC5160_INTEGER_FIELD = RegisterField{
		Mask:     TMC5160_INTEGER_MASK,
		Shift:    TMC5160_INTEGER_SHIFT,
		Register: TMC5160_ENC_CONST,
		IsSigned: false,
	}
	TMC5160_FRACTIONAL_FIELD = RegisterField{
		Mask:     TMC5160_FRACTIONAL_MASK,
		Shift:    TMC5160_FRACTIONAL_SHIFT,
		Register: TMC5160_ENC_CONST,
		IsSigned: false,
	}
	TMC5160_N_EVENT_FIELD = RegisterField{
		Mask:     TMC5160_N_EVENT_MASK,
		Shift:    TMC5160_N_EVENT_SHIFT,
		Register: TMC5160_ENC_STATUS,
		IsSigned: false,
	}
	TMC5160_DEVIATION_WARN_FIELD = RegisterField{
		Mask:     TMC5160_DEVIATION_WARN_MASK,
		Shift:    TMC5160_DEVIATION_WARN_SHIFT,
		Register: TMC5160_ENC_STATUS,
		IsSigned: false,
	}
	TMC5160_ENC_LATCH_FIELD = RegisterField{
		Mask:     TMC5160_ENC_LATCH_MASK,
		Shift:    TMC5160_ENC_LATCH_SHIFT,
		Register: TMC5160_ENC_LATCH,
		IsSigned: true,
	}
	TMC5160_ENC_DEVIATION_FIELD = RegisterField{
		Mask:     TMC5160_ENC_DEVIATION_MASK,
		Shift:    TMC5160_ENC_DEVIATION_SHIFT,
		Register: TMC5160_ENC_DEVIATION,
		IsSigned: false,
	}
	TMC5160_OFS0_FIELD = RegisterField{
		Mask:     TMC5160_OFS0_MASK,
		Shift:    TMC5160_OFS0_SHIFT,
		Register: TMC5160_MSLUT0,
		IsSigned: false,
	}
	TMC5160_OFS1_FIELD = RegisterField{
		Mask:     TMC5160_OFS1_MASK,
		Shift:    TMC5160_OFS1_SHIFT,
		Register: TMC5160_MSLUT0,
		IsSigned: false,
	}
	TMC5160_OFS2_FIELD = RegisterField{
		Mask:     TMC5160_OFS2_MASK,
		Shift:    TMC5160_OFS2_SHIFT,
		Register: TMC5160_MSLUT0,
		IsSigned: false,
	}
	TMC5160_OFS3_FIELD = RegisterField{
		Mask:     TMC5160_OFS3_MASK,
		Shift:    TMC5160_OFS3_SHIFT,
		Register: TMC5160_MSLUT0,
		IsSigned: false,
	}
	TMC5160_OFS4_FIELD = RegisterField{
		Mask:     TMC5160_OFS4_MASK,
		Shift:    TMC5160_OFS4_SHIFT,
		Register: TMC5160_MSLUT0,
		IsSigned: false,
	}
	TMC5160_OFS5_FIELD = RegisterField{
		Mask:     TMC5160_OFS5_MASK,
		Shift:    TMC5160_OFS5_SHIFT,
		Register: TMC5160_MSLUT0,
		IsSigned: false,
	}
	TMC5160_OFS6_FIELD = RegisterField{
		Mask:     TMC5160_OFS6_MASK,
		Shift:    TMC5160_OFS6_SHIFT,
		Register: TMC5160_MSLUT0,
		IsSigned: false,
	}
	TMC5160_OFS7_FIELD = RegisterField{
		Mask:     TMC5160_OFS7_MASK,
		Shift:    TMC5160_OFS7_SHIFT,
		Register: TMC5160_MSLUT0,
		IsSigned: false,
	}
	TMC5160_OFS8_FIELD = RegisterField{
		Mask:     TMC5160_OFS8_MASK,
		Shift:    TMC5160_OFS8_SHIFT,
		Register: TMC5160_MSLUT0,
		IsSigned: false,
	}
	TMC5160_OFS9_FIELD = RegisterField{
		Mask:     TMC5160_OFS9_MASK,
		Shift:    TMC5160_OFS9_SHIFT,
		Register: TMC5160_MSLUT0,
		IsSigned: false,
	}
	TMC5160_OFS10_FIELD = RegisterField{
		Mask:     TMC5160_OFS10_MASK,
		Shift:    TMC5160_OFS10_SHIFT,
		Register: TMC5160_MSLUT0,
		IsSigned: false,
	}
	TMC5160_OFS11_FIELD = RegisterField{
		Mask:     TMC5160_OFS11_MASK,
		Shift:    TMC5160_OFS11_SHIFT,
		Register: TMC5160_MSLUT0,
		IsSigned: false,
	}
	TMC5160_OFS12_FIELD = RegisterField{
		Mask:     TMC5160_OFS12_MASK,
		Shift:    TMC5160_OFS12_SHIFT,
		Register: TMC5160_MSLUT0,
		IsSigned: false,
	}
	TMC5160_OFS13_FIELD = RegisterField{
		Mask:     TMC5160_OFS13_MASK,
		Shift:    TMC5160_OFS13_SHIFT,
		Register: TMC5160_MSLUT0,
		IsSigned: false,
	}
	TMC5160_OFS14_FIELD = RegisterField{
		Mask:     TMC5160_OFS14_MASK,
		Shift:    TMC5160_OFS14_SHIFT,
		Register: TMC5160_MSLUT0,
		IsSigned: false,
	}
	TMC5160_OFS15_FIELD = RegisterField{
		Mask:     TMC5160_OFS15_MASK,
		Shift:    TMC5160_OFS15_SHIFT,
		Register: TMC5160_MSLUT0,
		IsSigned: false,
	}
	TMC5160_OFS16_FIELD = RegisterField{
		Mask:     TMC5160_OFS16_MASK,
		Shift:    TMC5160_OFS16_SHIFT,
		Register: TMC5160_MSLUT0,
		IsSigned: false,
	}
	TMC5160_OFS17_FIELD = RegisterField{
		Mask:     TMC5160_OFS17_MASK,
		Shift:    TMC5160_OFS17_SHIFT,
		Register: TMC5160_MSLUT0,
		IsSigned: false,
	}
	TMC5160_OFS18_FIELD = RegisterField{
		Mask:     TMC5160_OFS18_MASK,
		Shift:    TMC5160_OFS18_SHIFT,
		Register: TMC5160_MSLUT0,
		IsSigned: false,
	}
	TMC5160_OFS19_FIELD = RegisterField{
		Mask:     TMC5160_OFS19_MASK,
		Shift:    TMC5160_OFS19_SHIFT,
		Register: TMC5160_MSLUT0,
		IsSigned: false,
	}
	TMC5160_OFS20_FIELD = RegisterField{
		Mask:     TMC5160_OFS20_MASK,
		Shift:    TMC5160_OFS20_SHIFT,
		Register: TMC5160_MSLUT0,
		IsSigned: false,
	}
	TMC5160_OFS21_FIELD = RegisterField{
		Mask:     TMC5160_OFS21_MASK,
		Shift:    TMC5160_OFS21_SHIFT,
		Register: TMC5160_MSLUT0,
		IsSigned: false,
	}
	TMC5160_OFS22_FIELD = RegisterField{
		Mask:     TMC5160_OFS22_MASK,
		Shift:    TMC5160_OFS22_SHIFT,
		Register: TMC5160_MSLUT0,
		IsSigned: false,
	}
	TMC5160_OFS23_FIELD = RegisterField{
		Mask:     TMC5160_OFS23_MASK,
		Shift:    TMC5160_OFS23_SHIFT,
		Register: TMC5160_MSLUT0,
		IsSigned: false,
	}
	TMC5160_OFS24_FIELD = RegisterField{
		Mask:     TMC5160_OFS24_MASK,
		Shift:    TMC5160_OFS24_SHIFT,
		Register: TMC5160_MSLUT0,
		IsSigned: false,
	}
	TMC5160_OFS25_FIELD = RegisterField{
		Mask:     TMC5160_OFS25_MASK,
		Shift:    TMC5160_OFS25_SHIFT,
		Register: TMC5160_MSLUT0,
		IsSigned: false,
	}
	TMC5160_OFS26_FIELD = RegisterField{
		Mask:     TMC5160_OFS26_MASK,
		Shift:    TMC5160_OFS26_SHIFT,
		Register: TMC5160_MSLUT0,
		IsSigned: false,
	}
	TMC5160_OFS27_FIELD = RegisterField{
		Mask:     TMC5160_OFS27_MASK,
		Shift:    TMC5160_OFS27_SHIFT,
		Register: TMC5160_MSLUT0,
		IsSigned: false,
	}
	TMC5160_OFS28_FIELD = RegisterField{
		Mask:     TMC5160_OFS28_MASK,
		Shift:    TMC5160_OFS28_SHIFT,
		Register: TMC5160_MSLUT0,
		IsSigned: false,
	}
	TMC5160_OFS29_FIELD = RegisterField{
		Mask:     TMC5160_OFS29_MASK,
		Shift:    TMC5160_OFS29_SHIFT,
		Register: TMC5160_MSLUT0,
		IsSigned: false,
	}
	TMC5160_OFS30_FIELD = RegisterField{
		Mask:     TMC5160_OFS30_MASK,
		Shift:    TMC5160_OFS30_SHIFT,
		Register: TMC5160_MSLUT0,
		IsSigned: false,
	}
	TMC5160_OFS31_FIELD = RegisterField{
		Mask:     TMC5160_OFS31_MASK,
		Shift:    TMC5160_OFS31_SHIFT,
		Register: TMC5160_MSLUT0,
		IsSigned: false,
	}
	TMC5160_OFS32_FIELD = RegisterField{
		Mask:     TMC5160_OFS32_MASK,
		Shift:    TMC5160_OFS32_SHIFT,
		Register: TMC5160_MSLUT1,
		IsSigned: false,
	}
	TMC5160_OFS33_FIELD = RegisterField{
		Mask:     TMC5160_OFS33_MASK,
		Shift:    TMC5160_OFS33_SHIFT,
		Register: TMC5160_MSLUT1,
		IsSigned: false,
	}
	TMC5160_OFS34_FIELD = RegisterField{
		Mask:     TMC5160_OFS34_MASK,
		Shift:    TMC5160_OFS34_SHIFT,
		Register: TMC5160_MSLUT1,
		IsSigned: false,
	}
	TMC5160_OFS35_FIELD = RegisterField{
		Mask:     TMC5160_OFS35_MASK,
		Shift:    TMC5160_OFS35_SHIFT,
		Register: TMC5160_MSLUT1,
		IsSigned: false,
	}
	TMC5160_OFS36_FIELD = RegisterField{
		Mask:     TMC5160_OFS36_MASK,
		Shift:    TMC5160_OFS36_SHIFT,
		Register: TMC5160_MSLUT1,
		IsSigned: false,
	}
	TMC5160_OFS37_FIELD = RegisterField{
		Mask:     TMC5160_OFS37_MASK,
		Shift:    TMC5160_OFS37_SHIFT,
		Register: TMC5160_MSLUT1,
		IsSigned: false,
	}
	TMC5160_OFS38_FIELD = RegisterField{
		Mask:     TMC5160_OFS38_MASK,
		Shift:    TMC5160_OFS38_SHIFT,
		Register: TMC5160_MSLUT1,
		IsSigned: false,
	}
	TMC5160_OFS39_FIELD = RegisterField{
		Mask:     TMC5160_OFS39_MASK,
		Shift:    TMC5160_OFS39_SHIFT,
		Register: TMC5160_MSLUT1,
		IsSigned: false,
	}
	TMC5160_OFS40_FIELD = RegisterField{
		Mask:     TMC5160_OFS40_MASK,
		Shift:    TMC5160_OFS40_SHIFT,
		Register: TMC5160_MSLUT1,
		IsSigned: false,
	}
	TMC5160_OFS41_FIELD = RegisterField{
		Mask:     TMC5160_OFS41_MASK,
		Shift:    TMC5160_OFS41_SHIFT,
		Register: TMC5160_MSLUT1,
		IsSigned: false,
	}
	TMC5160_OFS42_FIELD = RegisterField{
		Mask:     TMC5160_OFS42_MASK,
		Shift:    TMC5160_OFS42_SHIFT,
		Register: TMC5160_MSLUT1,
		IsSigned: false,
	}
	TMC5160_OFS43_FIELD = RegisterField{
		Mask:     TMC5160_OFS43_MASK,
		Shift:    TMC5160_OFS43_SHIFT,
		Register: TMC5160_MSLUT1,
		IsSigned: false,
	}
	TMC5160_OFS44_FIELD = RegisterField{
		Mask:     TMC5160_OFS44_MASK,
		Shift:    TMC5160_OFS44_SHIFT,
		Register: TMC5160_MSLUT1,
		IsSigned: false,
	}
	TMC5160_OFS45_FIELD = RegisterField{
		Mask:     TMC5160_OFS45_MASK,
		Shift:    TMC5160_OFS45_SHIFT,
		Register: TMC5160_MSLUT1,
		IsSigned: false,
	}
	TMC5160_OFS46_FIELD = RegisterField{
		Mask:     TMC5160_OFS46_MASK,
		Shift:    TMC5160_OFS46_SHIFT,
		Register: TMC5160_MSLUT1,
		IsSigned: false,
	}
	TMC5160_OFS47_FIELD = RegisterField{
		Mask:     TMC5160_OFS47_MASK,
		Shift:    TMC5160_OFS47_SHIFT,
		Register: TMC5160_MSLUT1,
		IsSigned: false,
	}
	TMC5160_OFS48_FIELD = RegisterField{
		Mask:     TMC5160_OFS48_MASK,
		Shift:    TMC5160_OFS48_SHIFT,
		Register: TMC5160_MSLUT1,
		IsSigned: false,
	}
	TMC5160_OFS49_FIELD = RegisterField{
		Mask:     TMC5160_OFS49_MASK,
		Shift:    TMC5160_OFS49_SHIFT,
		Register: TMC5160_MSLUT1,
		IsSigned: false,
	}
	TMC5160_OFS50_FIELD = RegisterField{
		Mask:     TMC5160_OFS50_MASK,
		Shift:    TMC5160_OFS50_SHIFT,
		Register: TMC5160_MSLUT1,
		IsSigned: false,
	}
	TMC5160_OFS51_FIELD = RegisterField{
		Mask:     TMC5160_OFS51_MASK,
		Shift:    TMC5160_OFS51_SHIFT,
		Register: TMC5160_MSLUT1,
		IsSigned: false,
	}
	TMC5160_OFS52_FIELD = RegisterField{
		Mask:     TMC5160_OFS52_MASK,
		Shift:    TMC5160_OFS52_SHIFT,
		Register: TMC5160_MSLUT1,
		IsSigned: false,
	}
	TMC5160_OFS53_FIELD = RegisterField{
		Mask:     TMC5160_OFS53_MASK,
		Shift:    TMC5160_OFS53_SHIFT,
		Register: TMC5160_MSLUT1,
		IsSigned: false,
	}
	TMC5160_OFS54_FIELD = RegisterField{
		Mask:     TMC5160_OFS54_MASK,
		Shift:    TMC5160_OFS54_SHIFT,
		Register: TMC5160_MSLUT1,
		IsSigned: false,
	}
	TMC5160_OFS55_FIELD = RegisterField{
		Mask:     TMC5160_OFS55_MASK,
		Shift:    TMC5160_OFS55_SHIFT,
		Register: TMC5160_MSLUT1,
		IsSigned: false,
	}
	TMC5160_OFS56_FIELD = RegisterField{
		Mask:     TMC5160_OFS56_MASK,
		Shift:    TMC5160_OFS56_SHIFT,
		Register: TMC5160_MSLUT1,
		IsSigned: false,
	}
	TMC5160_OFS57_FIELD = RegisterField{
		Mask:     TMC5160_OFS57_MASK,
		Shift:    TMC5160_OFS57_SHIFT,
		Register: TMC5160_MSLUT1,
		IsSigned: false,
	}
	TMC5160_OFS58_FIELD = RegisterField{
		Mask:     TMC5160_OFS58_MASK,
		Shift:    TMC5160_OFS58_SHIFT,
		Register: TMC5160_MSLUT1,
		IsSigned: false,
	}
	TMC5160_OFS59_FIELD = RegisterField{
		Mask:     TMC5160_OFS59_MASK,
		Shift:    TMC5160_OFS59_SHIFT,
		Register: TMC5160_MSLUT1,
		IsSigned: false,
	}
	TMC5160_OFS60_FIELD = RegisterField{
		Mask:     TMC5160_OFS60_MASK,
		Shift:    TMC5160_OFS60_SHIFT,
		Register: TMC5160_MSLUT1,
		IsSigned: false,
	}
	TMC5160_OFS61_FIELD = RegisterField{
		Mask:     TMC5160_OFS61_MASK,
		Shift:    TMC5160_OFS61_SHIFT,
		Register: TMC5160_MSLUT1,
		IsSigned: false,
	}
	TMC5160_OFS62_FIELD = RegisterField{
		Mask:     TMC5160_OFS62_MASK,
		Shift:    TMC5160_OFS62_SHIFT,
		Register: TMC5160_MSLUT1,
		IsSigned: false,
	}
	TMC5160_OFS63_FIELD = RegisterField{
		Mask:     TMC5160_OFS63_MASK,
		Shift:    TMC5160_OFS63_SHIFT,
		Register: TMC5160_MSLUT1,
		IsSigned: false,
	}
	TMC5160_OFS64_FIELD = RegisterField{
		Mask:     TMC5160_OFS64_MASK,
		Shift:    TMC5160_OFS64_SHIFT,
		Register: TMC5160_MSLUT2,
		IsSigned: false,
	}
	TMC5160_OFS65_FIELD = RegisterField{
		Mask:     TMC5160_OFS65_MASK,
		Shift:    TMC5160_OFS65_SHIFT,
		Register: TMC5160_MSLUT2,
		IsSigned: false,
	}
	TMC5160_OFS66_FIELD = RegisterField{
		Mask:     TMC5160_OFS66_MASK,
		Shift:    TMC5160_OFS66_SHIFT,
		Register: TMC5160_MSLUT2,
		IsSigned: false,
	}
	TMC5160_OFS67_FIELD = RegisterField{
		Mask:     TMC5160_OFS67_MASK,
		Shift:    TMC5160_OFS67_SHIFT,
		Register: TMC5160_MSLUT2,
		IsSigned: false,
	}
	TMC5160_OFS68_FIELD = RegisterField{
		Mask:     TMC5160_OFS68_MASK,
		Shift:    TMC5160_OFS68_SHIFT,
		Register: TMC5160_MSLUT2,
		IsSigned: false,
	}
	TMC5160_OFS69_FIELD = RegisterField{
		Mask:     TMC5160_OFS69_MASK,
		Shift:    TMC5160_OFS69_SHIFT,
		Register: TMC5160_MSLUT2,
		IsSigned: false,
	}
	TMC5160_OFS70_FIELD = RegisterField{
		Mask:     TMC5160_OFS70_MASK,
		Shift:    TMC5160_OFS70_SHIFT,
		Register: TMC5160_MSLUT2,
		IsSigned: false,
	}
	TMC5160_OFS71_FIELD = RegisterField{
		Mask:     TMC5160_OFS71_MASK,
		Shift:    TMC5160_OFS71_SHIFT,
		Register: TMC5160_MSLUT2,
		IsSigned: false,
	}
	TMC5160_OFS72_FIELD = RegisterField{
		Mask:     TMC5160_OFS72_MASK,
		Shift:    TMC5160_OFS72_SHIFT,
		Register: TMC5160_MSLUT2,
		IsSigned: false,
	}
	TMC5160_OFS73_FIELD = RegisterField{
		Mask:     TMC5160_OFS73_MASK,
		Shift:    TMC5160_OFS73_SHIFT,
		Register: TMC5160_MSLUT2,
		IsSigned: false,
	}
	TMC5160_OFS74_FIELD = RegisterField{
		Mask:     TMC5160_OFS74_MASK,
		Shift:    TMC5160_OFS74_SHIFT,
		Register: TMC5160_MSLUT2,
		IsSigned: false,
	}
	TMC5160_OFS75_FIELD = RegisterField{
		Mask:     TMC5160_OFS75_MASK,
		Shift:    TMC5160_OFS75_SHIFT,
		Register: TMC5160_MSLUT2,
		IsSigned: false,
	}
	TMC5160_OFS76_FIELD = RegisterField{
		Mask:     TMC5160_OFS76_MASK,
		Shift:    TMC5160_OFS76_SHIFT,
		Register: TMC5160_MSLUT2,
		IsSigned: false,
	}
	TMC5160_OFS77_FIELD = RegisterField{
		Mask:     TMC5160_OFS77_MASK,
		Shift:    TMC5160_OFS77_SHIFT,
		Register: TMC5160_MSLUT2,
		IsSigned: false,
	}
	TMC5160_OFS78_FIELD = RegisterField{
		Mask:     TMC5160_OFS78_MASK,
		Shift:    TMC5160_OFS78_SHIFT,
		Register: TMC5160_MSLUT2,
		IsSigned: false,
	}
	TMC5160_OFS79_FIELD = RegisterField{
		Mask:     TMC5160_OFS79_MASK,
		Shift:    TMC5160_OFS79_SHIFT,
		Register: TMC5160_MSLUT2,
		IsSigned: false,
	}
	TMC5160_OFS80_FIELD = RegisterField{
		Mask:     TMC5160_OFS80_MASK,
		Shift:    TMC5160_OFS80_SHIFT,
		Register: TMC5160_MSLUT2,
		IsSigned: false,
	}
	TMC5160_OFS81_FIELD = RegisterField{
		Mask:     TMC5160_OFS81_MASK,
		Shift:    TMC5160_OFS81_SHIFT,
		Register: TMC5160_MSLUT2,
		IsSigned: false,
	}
	TMC5160_OFS82_FIELD = RegisterField{
		Mask:     TMC5160_OFS82_MASK,
		Shift:    TMC5160_OFS82_SHIFT,
		Register: TMC5160_MSLUT2,
		IsSigned: false,
	}
	TMC5160_OFS83_FIELD = RegisterField{
		Mask:     TMC5160_OFS83_MASK,
		Shift:    TMC5160_OFS83_SHIFT,
		Register: TMC5160_MSLUT2,
		IsSigned: false,
	}
	TMC5160_OFS84_FIELD = RegisterField{
		Mask:     TMC5160_OFS84_MASK,
		Shift:    TMC5160_OFS84_SHIFT,
		Register: TMC5160_MSLUT2,
		IsSigned: false,
	}
	TMC5160_OFS85_FIELD = RegisterField{
		Mask:     TMC5160_OFS85_MASK,
		Shift:    TMC5160_OFS85_SHIFT,
		Register: TMC5160_MSLUT2,
		IsSigned: false,
	}
	TMC5160_OFS86_FIELD = RegisterField{
		Mask:     TMC5160_OFS86_MASK,
		Shift:    TMC5160_OFS86_SHIFT,
		Register: TMC5160_MSLUT2,
		IsSigned: false,
	}
	TMC5160_OFS87_FIELD = RegisterField{
		Mask:     TMC5160_OFS87_MASK,
		Shift:    TMC5160_OFS87_SHIFT,
		Register: TMC5160_MSLUT2,
		IsSigned: false,
	}
	TMC5160_OFS88_FIELD = RegisterField{
		Mask:     TMC5160_OFS88_MASK,
		Shift:    TMC5160_OFS88_SHIFT,
		Register: TMC5160_MSLUT2,
		IsSigned: false,
	}
	TMC5160_OFS89_FIELD = RegisterField{
		Mask:     TMC5160_OFS89_MASK,
		Shift:    TMC5160_OFS89_SHIFT,
		Register: TMC5160_MSLUT2,
		IsSigned: false,
	}
	TMC5160_OFS90_FIELD = RegisterField{
		Mask:     TMC5160_OFS90_MASK,
		Shift:    TMC5160_OFS90_SHIFT,
		Register: TMC5160_MSLUT2,
		IsSigned: false,
	}
	TMC5160_OFS91_FIELD = RegisterField{
		Mask:     TMC5160_OFS91_MASK,
		Shift:    TMC5160_OFS91_SHIFT,
		Register: TMC5160_MSLUT2,
		IsSigned: false,
	}
	TMC5160_OFS92_FIELD = RegisterField{
		Mask:     TMC5160_OFS92_MASK,
		Shift:    TMC5160_OFS92_SHIFT,
		Register: TMC5160_MSLUT2,
		IsSigned: false,
	}
	TMC5160_OFS93_FIELD = RegisterField{
		Mask:     TMC5160_OFS93_MASK,
		Shift:    TMC5160_OFS93_SHIFT,
		Register: TMC5160_MSLUT2,
		IsSigned: false,
	}
	TMC5160_OFS94_FIELD = RegisterField{
		Mask:     TMC5160_OFS94_MASK,
		Shift:    TMC5160_OFS94_SHIFT,
		Register: TMC5160_MSLUT2,
		IsSigned: false,
	}
	TMC5160_OFS95_FIELD = RegisterField{
		Mask:     TMC5160_OFS95_MASK,
		Shift:    TMC5160_OFS95_SHIFT,
		Register: TMC5160_MSLUT2,
		IsSigned: false,
	}
	TMC5160_OFS96_FIELD = RegisterField{
		Mask:     TMC5160_OFS96_MASK,
		Shift:    TMC5160_OFS96_SHIFT,
		Register: TMC5160_MSLUT3,
		IsSigned: false,
	}
	TMC5160_OFS97_FIELD = RegisterField{
		Mask:     TMC5160_OFS97_MASK,
		Shift:    TMC5160_OFS97_SHIFT,
		Register: TMC5160_MSLUT3,
		IsSigned: false,
	}
	TMC5160_OFS98_FIELD = RegisterField{
		Mask:     TMC5160_OFS98_MASK,
		Shift:    TMC5160_OFS98_SHIFT,
		Register: TMC5160_MSLUT3,
		IsSigned: false,
	}
	TMC5160_OFS99_FIELD = RegisterField{
		Mask:     TMC5160_OFS99_MASK,
		Shift:    TMC5160_OFS99_SHIFT,
		Register: TMC5160_MSLUT3,
		IsSigned: false,
	}
	TMC5160_OFS100_FIELD = RegisterField{
		Mask:     TMC5160_OFS100_MASK,
		Shift:    TMC5160_OFS100_SHIFT,
		Register: TMC5160_MSLUT3,
		IsSigned: false,
	}
	TMC5160_OFS101_FIELD = RegisterField{
		Mask:     TMC5160_OFS101_MASK,
		Shift:    TMC5160_OFS101_SHIFT,
		Register: TMC5160_MSLUT3,
		IsSigned: false,
	}
	TMC5160_OFS102_FIELD = RegisterField{
		Mask:     TMC5160_OFS102_MASK,
		Shift:    TMC5160_OFS102_SHIFT,
		Register: TMC5160_MSLUT3,
		IsSigned: false,
	}
	TMC5160_OFS103_FIELD = RegisterField{
		Mask:     TMC5160_OFS103_MASK,
		Shift:    TMC5160_OFS103_SHIFT,
		Register: TMC5160_MSLUT3,
		IsSigned: false,
	}
	TMC5160_OFS104_FIELD = RegisterField{
		Mask:     TMC5160_OFS104_MASK,
		Shift:    TMC5160_OFS104_SHIFT,
		Register: TMC5160_MSLUT3,
		IsSigned: false,
	}
	TMC5160_OFS105_FIELD = RegisterField{
		Mask:     TMC5160_OFS105_MASK,
		Shift:    TMC5160_OFS105_SHIFT,
		Register: TMC5160_MSLUT3,
		IsSigned: false,
	}
	TMC5160_OFS106_FIELD = RegisterField{
		Mask:     TMC5160_OFS106_MASK,
		Shift:    TMC5160_OFS106_SHIFT,
		Register: TMC5160_MSLUT3,
		IsSigned: false,
	}
	TMC5160_OFS107_FIELD = RegisterField{
		Mask:     TMC5160_OFS107_MASK,
		Shift:    TMC5160_OFS107_SHIFT,
		Register: TMC5160_MSLUT3,
		IsSigned: false,
	}
	TMC5160_OFS108_FIELD = RegisterField{
		Mask:     TMC5160_OFS108_MASK,
		Shift:    TMC5160_OFS108_SHIFT,
		Register: TMC5160_MSLUT3,
		IsSigned: false,
	}
	TMC5160_OFS109_FIELD = RegisterField{
		Mask:     TMC5160_OFS109_MASK,
		Shift:    TMC5160_OFS109_SHIFT,
		Register: TMC5160_MSLUT3,
		IsSigned: false,
	}
	TMC5160_OFS110_FIELD = RegisterField{
		Mask:     TMC5160_OFS110_MASK,
		Shift:    TMC5160_OFS110_SHIFT,
		Register: TMC5160_MSLUT3,
		IsSigned: false,
	}
	TMC5160_OFS111_FIELD = RegisterField{
		Mask:     TMC5160_OFS111_MASK,
		Shift:    TMC5160_OFS111_SHIFT,
		Register: TMC5160_MSLUT3,
		IsSigned: false,
	}
	TMC5160_OFS112_FIELD = RegisterField{
		Mask:     TMC5160_OFS112_MASK,
		Shift:    TMC5160_OFS112_SHIFT,
		Register: TMC5160_MSLUT3,
		IsSigned: false,
	}
	TMC5160_OFS113_FIELD = RegisterField{
		Mask:     TMC5160_OFS113_MASK,
		Shift:    TMC5160_OFS113_SHIFT,
		Register: TMC5160_MSLUT3,
		IsSigned: false,
	}
	TMC5160_OFS114_FIELD = RegisterField{
		Mask:     TMC5160_OFS114_MASK,
		Shift:    TMC5160_OFS114_SHIFT,
		Register: TMC5160_MSLUT3,
		IsSigned: false,
	}
	TMC5160_OFS115_FIELD = RegisterField{
		Mask:     TMC5160_OFS115_MASK,
		Shift:    TMC5160_OFS115_SHIFT,
		Register: TMC5160_MSLUT3,
		IsSigned: false,
	}
	TMC5160_OFS116_FIELD = RegisterField{
		Mask:     TMC5160_OFS116_MASK,
		Shift:    TMC5160_OFS116_SHIFT,
		Register: TMC5160_MSLUT3,
		IsSigned: false,
	}
	TMC5160_OFS117_FIELD = RegisterField{
		Mask:     TMC5160_OFS117_MASK,
		Shift:    TMC5160_OFS117_SHIFT,
		Register: TMC5160_MSLUT3,
		IsSigned: false,
	}
	TMC5160_OFS118_FIELD = RegisterField{
		Mask:     TMC5160_OFS118_MASK,
		Shift:    TMC5160_OFS118_SHIFT,
		Register: TMC5160_MSLUT3,
		IsSigned: false,
	}
	TMC5160_OFS119_FIELD = RegisterField{
		Mask:     TMC5160_OFS119_MASK,
		Shift:    TMC5160_OFS119_SHIFT,
		Register: TMC5160_MSLUT3,
		IsSigned: false,
	}
	TMC5160_OFS120_FIELD = RegisterField{
		Mask:     TMC5160_OFS120_MASK,
		Shift:    TMC5160_OFS120_SHIFT,
		Register: TMC5160_MSLUT3,
		IsSigned: false,
	}
	TMC5160_OFS121_FIELD = RegisterField{
		Mask:     TMC5160_OFS121_MASK,
		Shift:    TMC5160_OFS121_SHIFT,
		Register: TMC5160_MSLUT3,
		IsSigned: false,
	}
	TMC5160_OFS122_FIELD = RegisterField{
		Mask:     TMC5160_OFS122_MASK,
		Shift:    TMC5160_OFS122_SHIFT,
		Register: TMC5160_MSLUT3,
		IsSigned: false,
	}
	TMC5160_OFS123_FIELD = RegisterField{
		Mask:     TMC5160_OFS123_MASK,
		Shift:    TMC5160_OFS123_SHIFT,
		Register: TMC5160_MSLUT3,
		IsSigned: false,
	}
	TMC5160_OFS124_FIELD = RegisterField{
		Mask:     TMC5160_OFS124_MASK,
		Shift:    TMC5160_OFS124_SHIFT,
		Register: TMC5160_MSLUT3,
		IsSigned: false,
	}
	TMC5160_OFS125_FIELD = RegisterField{
		Mask:     TMC5160_OFS125_MASK,
		Shift:    TMC5160_OFS125_SHIFT,
		Register: TMC5160_MSLUT3,
		IsSigned: false,
	}
	TMC5160_OFS126_FIELD = RegisterField{
		Mask:     TMC5160_OFS126_MASK,
		Shift:    TMC5160_OFS126_SHIFT,
		Register: TMC5160_MSLUT3,
		IsSigned: false,
	}
	TMC5160_OFS127_FIELD = RegisterField{
		Mask:     TMC5160_OFS127_MASK,
		Shift:    TMC5160_OFS127_SHIFT,
		Register: TMC5160_MSLUT3,
		IsSigned: false,
	}
	TMC5160_OFS128_FIELD = RegisterField{
		Mask:     TMC5160_OFS128_MASK,
		Shift:    TMC5160_OFS128_SHIFT,
		Register: TMC5160_MSLUT4,
		IsSigned: false,
	}
	TMC5160_OFS129_FIELD = RegisterField{
		Mask:     TMC5160_OFS129_MASK,
		Shift:    TMC5160_OFS129_SHIFT,
		Register: TMC5160_MSLUT4,
		IsSigned: false,
	}
	TMC5160_OFS130_FIELD = RegisterField{
		Mask:     TMC5160_OFS130_MASK,
		Shift:    TMC5160_OFS130_SHIFT,
		Register: TMC5160_MSLUT4,
		IsSigned: false,
	}
	TMC5160_OFS131_FIELD = RegisterField{
		Mask:     TMC5160_OFS131_MASK,
		Shift:    TMC5160_OFS131_SHIFT,
		Register: TMC5160_MSLUT4,
		IsSigned: false,
	}
	TMC5160_OFS132_FIELD = RegisterField{
		Mask:     TMC5160_OFS132_MASK,
		Shift:    TMC5160_OFS132_SHIFT,
		Register: TMC5160_MSLUT4,
		IsSigned: false,
	}
	TMC5160_OFS133_FIELD = RegisterField{
		Mask:     TMC5160_OFS133_MASK,
		Shift:    TMC5160_OFS133_SHIFT,
		Register: TMC5160_MSLUT4,
		IsSigned: false,
	}
	TMC5160_OFS134_FIELD = RegisterField{
		Mask:     TMC5160_OFS134_MASK,
		Shift:    TMC5160_OFS134_SHIFT,
		Register: TMC5160_MSLUT4,
		IsSigned: false,
	}
	TMC5160_OFS135_FIELD = RegisterField{
		Mask:     TMC5160_OFS135_MASK,
		Shift:    TMC5160_OFS135_SHIFT,
		Register: TMC5160_MSLUT4,
		IsSigned: false,
	}
	TMC5160_OFS136_FIELD = RegisterField{
		Mask:     TMC5160_OFS136_MASK,
		Shift:    TMC5160_OFS136_SHIFT,
		Register: TMC5160_MSLUT4,
		IsSigned: false,
	}
	TMC5160_OFS137_FIELD = RegisterField{
		Mask:     TMC5160_OFS137_MASK,
		Shift:    TMC5160_OFS137_SHIFT,
		Register: TMC5160_MSLUT4,
		IsSigned: false,
	}
	TMC5160_OFS138_FIELD = RegisterField{
		Mask:     TMC5160_OFS138_MASK,
		Shift:    TMC5160_OFS138_SHIFT,
		Register: TMC5160_MSLUT4,
		IsSigned: false,
	}
	TMC5160_OFS139_FIELD = RegisterField{
		Mask:     TMC5160_OFS139_MASK,
		Shift:    TMC5160_OFS139_SHIFT,
		Register: TMC5160_MSLUT4,
		IsSigned: false,
	}
	TMC5160_OFS140_FIELD = RegisterField{
		Mask:     TMC5160_OFS140_MASK,
		Shift:    TMC5160_OFS140_SHIFT,
		Register: TMC5160_MSLUT4,
		IsSigned: false,
	}
	TMC5160_OFS141_FIELD = RegisterField{
		Mask:     TMC5160_OFS141_MASK,
		Shift:    TMC5160_OFS141_SHIFT,
		Register: TMC5160_MSLUT4,
		IsSigned: false,
	}
	TMC5160_OFS142_FIELD = RegisterField{
		Mask:     TMC5160_OFS142_MASK,
		Shift:    TMC5160_OFS142_SHIFT,
		Register: TMC5160_MSLUT4,
		IsSigned: false,
	}
	TMC5160_OFS143_FIELD = RegisterField{
		Mask:     TMC5160_OFS143_MASK,
		Shift:    TMC5160_OFS143_SHIFT,
		Register: TMC5160_MSLUT4,
		IsSigned: false,
	}
	TMC5160_OFS144_FIELD = RegisterField{
		Mask:     TMC5160_OFS144_MASK,
		Shift:    TMC5160_OFS144_SHIFT,
		Register: TMC5160_MSLUT4,
		IsSigned: false,
	}
	TMC5160_OFS145_FIELD = RegisterField{
		Mask:     TMC5160_OFS145_MASK,
		Shift:    TMC5160_OFS145_SHIFT,
		Register: TMC5160_MSLUT4,
		IsSigned: false,
	}
	TMC5160_OFS146_FIELD = RegisterField{
		Mask:     TMC5160_OFS146_MASK,
		Shift:    TMC5160_OFS146_SHIFT,
		Register: TMC5160_MSLUT4,
		IsSigned: false,
	}
	TMC5160_OFS147_FIELD = RegisterField{
		Mask:     TMC5160_OFS147_MASK,
		Shift:    TMC5160_OFS147_SHIFT,
		Register: TMC5160_MSLUT4,
		IsSigned: false,
	}
	TMC5160_OFS148_FIELD = RegisterField{
		Mask:     TMC5160_OFS148_MASK,
		Shift:    TMC5160_OFS148_SHIFT,
		Register: TMC5160_MSLUT4,
		IsSigned: false,
	}
	TMC5160_OFS149_FIELD = RegisterField{
		Mask:     TMC5160_OFS149_MASK,
		Shift:    TMC5160_OFS149_SHIFT,
		Register: TMC5160_MSLUT4,
		IsSigned: false,
	}
	TMC5160_OFS150_FIELD = RegisterField{
		Mask:     TMC5160_OFS150_MASK,
		Shift:    TMC5160_OFS150_SHIFT,
		Register: TMC5160_MSLUT4,
		IsSigned: false,
	}
	TMC5160_OFS151_FIELD = RegisterField{
		Mask:     TMC5160_OFS151_MASK,
		Shift:    TMC5160_OFS151_SHIFT,
		Register: TMC5160_MSLUT4,
		IsSigned: false,
	}
	TMC5160_OFS152_FIELD = RegisterField{
		Mask:     TMC5160_OFS152_MASK,
		Shift:    TMC5160_OFS152_SHIFT,
		Register: TMC5160_MSLUT4,
		IsSigned: false,
	}
	TMC5160_OFS153_FIELD = RegisterField{
		Mask:     TMC5160_OFS153_MASK,
		Shift:    TMC5160_OFS153_SHIFT,
		Register: TMC5160_MSLUT4,
		IsSigned: false,
	}
	TMC5160_OFS154_FIELD = RegisterField{
		Mask:     TMC5160_OFS154_MASK,
		Shift:    TMC5160_OFS154_SHIFT,
		Register: TMC5160_MSLUT4,
		IsSigned: false,
	}
	TMC5160_OFS155_FIELD = RegisterField{
		Mask:     TMC5160_OFS155_MASK,
		Shift:    TMC5160_OFS155_SHIFT,
		Register: TMC5160_MSLUT4,
		IsSigned: false,
	}
	TMC5160_OFS156_FIELD = RegisterField{
		Mask:     TMC5160_OFS156_MASK,
		Shift:    TMC5160_OFS156_SHIFT,
		Register: TMC5160_MSLUT4,
		IsSigned: false,
	}
	TMC5160_OFS157_FIELD = RegisterField{
		Mask:     TMC5160_OFS157_MASK,
		Shift:    TMC5160_OFS157_SHIFT,
		Register: TMC5160_MSLUT4,
		IsSigned: false,
	}
	TMC5160_OFS158_FIELD = RegisterField{
		Mask:     TMC5160_OFS158_MASK,
		Shift:    TMC5160_OFS158_SHIFT,
		Register: TMC5160_MSLUT4,
		IsSigned: false,
	}
	TMC5160_OFS159_FIELD = RegisterField{
		Mask:     TMC5160_OFS159_MASK,
		Shift:    TMC5160_OFS159_SHIFT,
		Register: TMC5160_MSLUT4,
		IsSigned: false,
	}
	TMC5160_OFS160_FIELD = RegisterField{
		Mask:     TMC5160_OFS160_MASK,
		Shift:    TMC5160_OFS160_SHIFT,
		Register: TMC5160_MSLUT5,
		IsSigned: false,
	}
	TMC5160_OFS161_FIELD = RegisterField{
		Mask:     TMC5160_OFS161_MASK,
		Shift:    TMC5160_OFS161_SHIFT,
		Register: TMC5160_MSLUT5,
		IsSigned: false,
	}
	TMC5160_OFS162_FIELD = RegisterField{
		Mask:     TMC5160_OFS162_MASK,
		Shift:    TMC5160_OFS162_SHIFT,
		Register: TMC5160_MSLUT5,
		IsSigned: false,
	}
	TMC5160_OFS163_FIELD = RegisterField{
		Mask:     TMC5160_OFS163_MASK,
		Shift:    TMC5160_OFS163_SHIFT,
		Register: TMC5160_MSLUT5,
		IsSigned: false,
	}
	TMC5160_OFS164_FIELD = RegisterField{
		Mask:     TMC5160_OFS164_MASK,
		Shift:    TMC5160_OFS164_SHIFT,
		Register: TMC5160_MSLUT5,
		IsSigned: false,
	}
	TMC5160_OFS165_FIELD = RegisterField{
		Mask:     TMC5160_OFS165_MASK,
		Shift:    TMC5160_OFS165_SHIFT,
		Register: TMC5160_MSLUT5,
		IsSigned: false,
	}
	TMC5160_OFS166_FIELD = RegisterField{
		Mask:     TMC5160_OFS166_MASK,
		Shift:    TMC5160_OFS166_SHIFT,
		Register: TMC5160_MSLUT5,
		IsSigned: false,
	}
	TMC5160_OFS167_FIELD = RegisterField{
		Mask:     TMC5160_OFS167_MASK,
		Shift:    TMC5160_OFS167_SHIFT,
		Register: TMC5160_MSLUT5,
		IsSigned: false,
	}
	TMC5160_OFS168_FIELD = RegisterField{
		Mask:     TMC5160_OFS168_MASK,
		Shift:    TMC5160_OFS168_SHIFT,
		Register: TMC5160_MSLUT5,
		IsSigned: false,
	}
	TMC5160_OFS169_FIELD = RegisterField{
		Mask:     TMC5160_OFS169_MASK,
		Shift:    TMC5160_OFS169_SHIFT,
		Register: TMC5160_MSLUT5,
		IsSigned: false,
	}
	TMC5160_OFS170_FIELD = RegisterField{
		Mask:     TMC5160_OFS170_MASK,
		Shift:    TMC5160_OFS170_SHIFT,
		Register: TMC5160_MSLUT5,
		IsSigned: false,
	}
	TMC5160_OFS171_FIELD = RegisterField{
		Mask:     TMC5160_OFS171_MASK,
		Shift:    TMC5160_OFS171_SHIFT,
		Register: TMC5160_MSLUT5,
		IsSigned: false,
	}
	TMC5160_OFS172_FIELD = RegisterField{
		Mask:     TMC5160_OFS172_MASK,
		Shift:    TMC5160_OFS172_SHIFT,
		Register: TMC5160_MSLUT5,
		IsSigned: false,
	}
	TMC5160_OFS173_FIELD = RegisterField{
		Mask:     TMC5160_OFS173_MASK,
		Shift:    TMC5160_OFS173_SHIFT,
		Register: TMC5160_MSLUT5,
		IsSigned: false,
	}
	TMC5160_OFS174_FIELD = RegisterField{
		Mask:     TMC5160_OFS174_MASK,
		Shift:    TMC5160_OFS174_SHIFT,
		Register: TMC5160_MSLUT5,
		IsSigned: false,
	}
	TMC5160_OFS175_FIELD = RegisterField{
		Mask:     TMC5160_OFS175_MASK,
		Shift:    TMC5160_OFS175_SHIFT,
		Register: TMC5160_MSLUT5,
		IsSigned: false,
	}
	TMC5160_OFS176_FIELD = RegisterField{
		Mask:     TMC5160_OFS176_MASK,
		Shift:    TMC5160_OFS176_SHIFT,
		Register: TMC5160_MSLUT5,
		IsSigned: false,
	}
	TMC5160_OFS177_FIELD = RegisterField{
		Mask:     TMC5160_OFS177_MASK,
		Shift:    TMC5160_OFS177_SHIFT,
		Register: TMC5160_MSLUT5,
		IsSigned: false,
	}
	TMC5160_OFS178_FIELD = RegisterField{
		Mask:     TMC5160_OFS178_MASK,
		Shift:    TMC5160_OFS178_SHIFT,
		Register: TMC5160_MSLUT5,
		IsSigned: false,
	}
	TMC5160_OFS179_FIELD = RegisterField{
		Mask:     TMC5160_OFS179_MASK,
		Shift:    TMC5160_OFS179_SHIFT,
		Register: TMC5160_MSLUT5,
		IsSigned: false,
	}
	TMC5160_OFS180_FIELD = RegisterField{
		Mask:     TMC5160_OFS180_MASK,
		Shift:    TMC5160_OFS180_SHIFT,
		Register: TMC5160_MSLUT5,
		IsSigned: false,
	}
	TMC5160_OFS181_FIELD = RegisterField{
		Mask:     TMC5160_OFS181_MASK,
		Shift:    TMC5160_OFS181_SHIFT,
		Register: TMC5160_MSLUT5,
		IsSigned: false,
	}
	TMC5160_OFS182_FIELD = RegisterField{
		Mask:     TMC5160_OFS182_MASK,
		Shift:    TMC5160_OFS182_SHIFT,
		Register: TMC5160_MSLUT5,
		IsSigned: false,
	}
	TMC5160_OFS183_FIELD = RegisterField{
		Mask:     TMC5160_OFS183_MASK,
		Shift:    TMC5160_OFS183_SHIFT,
		Register: TMC5160_MSLUT5,
		IsSigned: false,
	}
	TMC5160_OFS184_FIELD = RegisterField{
		Mask:     TMC5160_OFS184_MASK,
		Shift:    TMC5160_OFS184_SHIFT,
		Register: TMC5160_MSLUT5,
		IsSigned: false,
	}
	TMC5160_OFS185_FIELD = RegisterField{
		Mask:     TMC5160_OFS185_MASK,
		Shift:    TMC5160_OFS185_SHIFT,
		Register: TMC5160_MSLUT5,
		IsSigned: false,
	}
	TMC5160_OFS186_FIELD = RegisterField{
		Mask:     TMC5160_OFS186_MASK,
		Shift:    TMC5160_OFS186_SHIFT,
		Register: TMC5160_MSLUT5,
		IsSigned: false,
	}
	TMC5160_OFS187_FIELD = RegisterField{
		Mask:     TMC5160_OFS187_MASK,
		Shift:    TMC5160_OFS187_SHIFT,
		Register: TMC5160_MSLUT5,
		IsSigned: false,
	}
	TMC5160_OFS188_FIELD = RegisterField{
		Mask:     TMC5160_OFS188_MASK,
		Shift:    TMC5160_OFS188_SHIFT,
		Register: TMC5160_MSLUT5,
		IsSigned: false,
	}
	TMC5160_OFS189_FIELD = RegisterField{
		Mask:     TMC5160_OFS189_MASK,
		Shift:    TMC5160_OFS189_SHIFT,
		Register: TMC5160_MSLUT5,
		IsSigned: false,
	}
	TMC5160_OFS190_FIELD = RegisterField{
		Mask:     TMC5160_OFS190_MASK,
		Shift:    TMC5160_OFS190_SHIFT,
		Register: TMC5160_MSLUT5,
		IsSigned: false,
	}
	TMC5160_OFS191_FIELD = RegisterField{
		Mask:     TMC5160_OFS191_MASK,
		Shift:    TMC5160_OFS191_SHIFT,
		Register: TMC5160_MSLUT5,
		IsSigned: false,
	}
	TMC5160_OFS192_FIELD = RegisterField{
		Mask:     TMC5160_OFS192_MASK,
		Shift:    TMC5160_OFS192_SHIFT,
		Register: TMC5160_MSLUT6,
		IsSigned: false,
	}
	TMC5160_OFS193_FIELD = RegisterField{
		Mask:     TMC5160_OFS193_MASK,
		Shift:    TMC5160_OFS193_SHIFT,
		Register: TMC5160_MSLUT6,
		IsSigned: false,
	}
	TMC5160_OFS194_FIELD = RegisterField{
		Mask:     TMC5160_OFS194_MASK,
		Shift:    TMC5160_OFS194_SHIFT,
		Register: TMC5160_MSLUT6,
		IsSigned: false,
	}
	TMC5160_OFS195_FIELD = RegisterField{
		Mask:     TMC5160_OFS195_MASK,
		Shift:    TMC5160_OFS195_SHIFT,
		Register: TMC5160_MSLUT6,
		IsSigned: false,
	}
	TMC5160_OFS196_FIELD = RegisterField{
		Mask:     TMC5160_OFS196_MASK,
		Shift:    TMC5160_OFS196_SHIFT,
		Register: TMC5160_MSLUT6,
		IsSigned: false,
	}
	TMC5160_OFS197_FIELD = RegisterField{
		Mask:     TMC5160_OFS197_MASK,
		Shift:    TMC5160_OFS197_SHIFT,
		Register: TMC5160_MSLUT6,
		IsSigned: false,
	}
	TMC5160_OFS198_FIELD = RegisterField{
		Mask:     TMC5160_OFS198_MASK,
		Shift:    TMC5160_OFS198_SHIFT,
		Register: TMC5160_MSLUT6,
		IsSigned: false,
	}
	TMC5160_OFS199_FIELD = RegisterField{
		Mask:     TMC5160_OFS199_MASK,
		Shift:    TMC5160_OFS199_SHIFT,
		Register: TMC5160_MSLUT6,
		IsSigned: false,
	}
	TMC5160_OFS200_FIELD = RegisterField{
		Mask:     TMC5160_OFS200_MASK,
		Shift:    TMC5160_OFS200_SHIFT,
		Register: TMC5160_MSLUT6,
		IsSigned: false,
	}
	TMC5160_OFS201_FIELD = RegisterField{
		Mask:     TMC5160_OFS201_MASK,
		Shift:    TMC5160_OFS201_SHIFT,
		Register: TMC5160_MSLUT6,
		IsSigned: false,
	}
	TMC5160_OFS202_FIELD = RegisterField{
		Mask:     TMC5160_OFS202_MASK,
		Shift:    TMC5160_OFS202_SHIFT,
		Register: TMC5160_MSLUT6,
		IsSigned: false,
	}
	TMC5160_OFS203_FIELD = RegisterField{
		Mask:     TMC5160_OFS203_MASK,
		Shift:    TMC5160_OFS203_SHIFT,
		Register: TMC5160_MSLUT6,
		IsSigned: false,
	}
	TMC5160_OFS204_FIELD = RegisterField{
		Mask:     TMC5160_OFS204_MASK,
		Shift:    TMC5160_OFS204_SHIFT,
		Register: TMC5160_MSLUT6,
		IsSigned: false,
	}
	TMC5160_OFS205_FIELD = RegisterField{
		Mask:     TMC5160_OFS205_MASK,
		Shift:    TMC5160_OFS205_SHIFT,
		Register: TMC5160_MSLUT6,
		IsSigned: false,
	}
	TMC5160_OFS206_FIELD = RegisterField{
		Mask:     TMC5160_OFS206_MASK,
		Shift:    TMC5160_OFS206_SHIFT,
		Register: TMC5160_MSLUT6,
		IsSigned: false,
	}
	TMC5160_OFS207_FIELD = RegisterField{
		Mask:     TMC5160_OFS207_MASK,
		Shift:    TMC5160_OFS207_SHIFT,
		Register: TMC5160_MSLUT6,
		IsSigned: false,
	}
	TMC5160_OFS208_FIELD = RegisterField{
		Mask:     TMC5160_OFS208_MASK,
		Shift:    TMC5160_OFS208_SHIFT,
		Register: TMC5160_MSLUT6,
		IsSigned: false,
	}
	TMC5160_OFS209_FIELD = RegisterField{
		Mask:     TMC5160_OFS209_MASK,
		Shift:    TMC5160_OFS209_SHIFT,
		Register: TMC5160_MSLUT6,
		IsSigned: false,
	}
	TMC5160_OFS210_FIELD = RegisterField{
		Mask:     TMC5160_OFS210_MASK,
		Shift:    TMC5160_OFS210_SHIFT,
		Register: TMC5160_MSLUT6,
		IsSigned: false,
	}
	TMC5160_OFS211_FIELD = RegisterField{
		Mask:     TMC5160_OFS211_MASK,
		Shift:    TMC5160_OFS211_SHIFT,
		Register: TMC5160_MSLUT6,
		IsSigned: false,
	}
	TMC5160_OFS212_FIELD = RegisterField{
		Mask:     TMC5160_OFS212_MASK,
		Shift:    TMC5160_OFS212_SHIFT,
		Register: TMC5160_MSLUT6,
		IsSigned: false,
	}
	TMC5160_OFS213_FIELD = RegisterField{
		Mask:     TMC5160_OFS213_MASK,
		Shift:    TMC5160_OFS213_SHIFT,
		Register: TMC5160_MSLUT6,
		IsSigned: false,
	}
	TMC5160_OFS214_FIELD = RegisterField{
		Mask:     TMC5160_OFS214_MASK,
		Shift:    TMC5160_OFS214_SHIFT,
		Register: TMC5160_MSLUT6,
		IsSigned: false,
	}
	TMC5160_OFS215_FIELD = RegisterField{
		Mask:     TMC5160_OFS215_MASK,
		Shift:    TMC5160_OFS215_SHIFT,
		Register: TMC5160_MSLUT6,
		IsSigned: false,
	}
	TMC5160_OFS216_FIELD = RegisterField{
		Mask:     TMC5160_OFS216_MASK,
		Shift:    TMC5160_OFS216_SHIFT,
		Register: TMC5160_MSLUT6,
		IsSigned: false,
	}
	TMC5160_OFS217_FIELD = RegisterField{
		Mask:     TMC5160_OFS217_MASK,
		Shift:    TMC5160_OFS217_SHIFT,
		Register: TMC5160_MSLUT6,
		IsSigned: false,
	}
	TMC5160_OFS218_FIELD = RegisterField{
		Mask:     TMC5160_OFS218_MASK,
		Shift:    TMC5160_OFS218_SHIFT,
		Register: TMC5160_MSLUT6,
		IsSigned: false,
	}
	TMC5160_OFS219_FIELD = RegisterField{
		Mask:     TMC5160_OFS219_MASK,
		Shift:    TMC5160_OFS219_SHIFT,
		Register: TMC5160_MSLUT6,
		IsSigned: false,
	}
	TMC5160_OFS220_FIELD = RegisterField{
		Mask:     TMC5160_OFS220_MASK,
		Shift:    TMC5160_OFS220_SHIFT,
		Register: TMC5160_MSLUT6,
		IsSigned: false,
	}
	TMC5160_OFS221_FIELD = RegisterField{
		Mask:     TMC5160_OFS221_MASK,
		Shift:    TMC5160_OFS221_SHIFT,
		Register: TMC5160_MSLUT6,
		IsSigned: false,
	}
	TMC5160_OFS222_FIELD = RegisterField{
		Mask:     TMC5160_OFS222_MASK,
		Shift:    TMC5160_OFS222_SHIFT,
		Register: TMC5160_MSLUT6,
		IsSigned: false,
	}
	TMC5160_OFS223_FIELD = RegisterField{
		Mask:     TMC5160_OFS223_MASK,
		Shift:    TMC5160_OFS223_SHIFT,
		Register: TMC5160_MSLUT6,
		IsSigned: false,
	}
	TMC5160_OFS224_FIELD = RegisterField{
		Mask:     TMC5160_OFS224_MASK,
		Shift:    TMC5160_OFS224_SHIFT,
		Register: TMC5160_MSLUT7,
		IsSigned: false,
	}
	TMC5160_OFS225_FIELD = RegisterField{
		Mask:     TMC5160_OFS225_MASK,
		Shift:    TMC5160_OFS225_SHIFT,
		Register: TMC5160_MSLUT7,
		IsSigned: false,
	}
	TMC5160_OFS226_FIELD = RegisterField{
		Mask:     TMC5160_OFS226_MASK,
		Shift:    TMC5160_OFS226_SHIFT,
		Register: TMC5160_MSLUT7,
		IsSigned: false,
	}
	TMC5160_OFS227_FIELD = RegisterField{
		Mask:     TMC5160_OFS227_MASK,
		Shift:    TMC5160_OFS227_SHIFT,
		Register: TMC5160_MSLUT7,
		IsSigned: false,
	}
	TMC5160_OFS228_FIELD = RegisterField{
		Mask:     TMC5160_OFS228_MASK,
		Shift:    TMC5160_OFS228_SHIFT,
		Register: TMC5160_MSLUT7,
		IsSigned: false,
	}
	TMC5160_OFS229_FIELD = RegisterField{
		Mask:     TMC5160_OFS229_MASK,
		Shift:    TMC5160_OFS229_SHIFT,
		Register: TMC5160_MSLUT7,
		IsSigned: false,
	}
	TMC5160_OFS230_FIELD = RegisterField{
		Mask:     TMC5160_OFS230_MASK,
		Shift:    TMC5160_OFS230_SHIFT,
		Register: TMC5160_MSLUT7,
		IsSigned: false,
	}
	TMC5160_OFS231_FIELD = RegisterField{
		Mask:     TMC5160_OFS231_MASK,
		Shift:    TMC5160_OFS231_SHIFT,
		Register: TMC5160_MSLUT7,
		IsSigned: false,
	}
	TMC5160_OFS232_FIELD = RegisterField{
		Mask:     TMC5160_OFS232_MASK,
		Shift:    TMC5160_OFS232_SHIFT,
		Register: TMC5160_MSLUT7,
		IsSigned: false,
	}
	TMC5160_OFS233_FIELD = RegisterField{
		Mask:     TMC5160_OFS233_MASK,
		Shift:    TMC5160_OFS233_SHIFT,
		Register: TMC5160_MSLUT7,
		IsSigned: false,
	}
	TMC5160_OFS234_FIELD = RegisterField{
		Mask:     TMC5160_OFS234_MASK,
		Shift:    TMC5160_OFS234_SHIFT,
		Register: TMC5160_MSLUT7,
		IsSigned: false,
	}
	TMC5160_OFS235_FIELD = RegisterField{
		Mask:     TMC5160_OFS235_MASK,
		Shift:    TMC5160_OFS235_SHIFT,
		Register: TMC5160_MSLUT7,
		IsSigned: false,
	}
	TMC5160_OFS236_FIELD = RegisterField{
		Mask:     TMC5160_OFS236_MASK,
		Shift:    TMC5160_OFS236_SHIFT,
		Register: TMC5160_MSLUT7,
		IsSigned: false,
	}
	TMC5160_OFS237_FIELD = RegisterField{
		Mask:     TMC5160_OFS237_MASK,
		Shift:    TMC5160_OFS237_SHIFT,
		Register: TMC5160_MSLUT7,
		IsSigned: false,
	}
	TMC5160_OFS238_FIELD = RegisterField{
		Mask:     TMC5160_OFS238_MASK,
		Shift:    TMC5160_OFS238_SHIFT,
		Register: TMC5160_MSLUT7,
		IsSigned: false,
	}
	TMC5160_OFS239_FIELD = RegisterField{
		Mask:     TMC5160_OFS239_MASK,
		Shift:    TMC5160_OFS239_SHIFT,
		Register: TMC5160_MSLUT7,
		IsSigned: false,
	}
	TMC5160_OFS240_FIELD = RegisterField{
		Mask:     TMC5160_OFS240_MASK,
		Shift:    TMC5160_OFS240_SHIFT,
		Register: TMC5160_MSLUT7,
		IsSigned: false,
	}
	TMC5160_OFS241_FIELD = RegisterField{
		Mask:     TMC5160_OFS241_MASK,
		Shift:    TMC5160_OFS241_SHIFT,
		Register: TMC5160_MSLUT7,
		IsSigned: false,
	}
	TMC5160_OFS242_FIELD = RegisterField{
		Mask:     TMC5160_OFS242_MASK,
		Shift:    TMC5160_OFS242_SHIFT,
		Register: TMC5160_MSLUT7,
		IsSigned: false,
	}
	TMC5160_OFS243_FIELD = RegisterField{
		Mask:     TMC5160_OFS243_MASK,
		Shift:    TMC5160_OFS243_SHIFT,
		Register: TMC5160_MSLUT7,
		IsSigned: false,
	}
	TMC5160_OFS244_FIELD = RegisterField{
		Mask:     TMC5160_OFS244_MASK,
		Shift:    TMC5160_OFS244_SHIFT,
		Register: TMC5160_MSLUT7,
		IsSigned: false,
	}
	TMC5160_OFS245_FIELD = RegisterField{
		Mask:     TMC5160_OFS245_MASK,
		Shift:    TMC5160_OFS245_SHIFT,
		Register: TMC5160_MSLUT7,
		IsSigned: false,
	}
	TMC5160_OFS246_FIELD = RegisterField{
		Mask:     TMC5160_OFS246_MASK,
		Shift:    TMC5160_OFS246_SHIFT,
		Register: TMC5160_MSLUT7,
		IsSigned: false,
	}
	TMC5160_OFS247_FIELD = RegisterField{
		Mask:     TMC5160_OFS247_MASK,
		Shift:    TMC5160_OFS247_SHIFT,
		Register: TMC5160_MSLUT7,
		IsSigned: false,
	}
	TMC5160_OFS248_FIELD = RegisterField{
		Mask:     TMC5160_OFS248_MASK,
		Shift:    TMC5160_OFS248_SHIFT,
		Register: TMC5160_MSLUT7,
		IsSigned: false,
	}
	TMC5160_OFS249_FIELD = RegisterField{
		Mask:     TMC5160_OFS249_MASK,
		Shift:    TMC5160_OFS249_SHIFT,
		Register: TMC5160_MSLUT7,
		IsSigned: false,
	}
	TMC5160_OFS250_FIELD = RegisterField{
		Mask:     TMC5160_OFS250_MASK,
		Shift:    TMC5160_OFS250_SHIFT,
		Register: TMC5160_MSLUT7,
		IsSigned: false,
	}
	TMC5160_OFS251_FIELD = RegisterField{
		Mask:     TMC5160_OFS251_MASK,
		Shift:    TMC5160_OFS251_SHIFT,
		Register: TMC5160_MSLUT7,
		IsSigned: false,
	}
	TMC5160_OFS252_FIELD = RegisterField{
		Mask:     TMC5160_OFS252_MASK,
		Shift:    TMC5160_OFS252_SHIFT,
		Register: TMC5160_MSLUT7,
		IsSigned: false,
	}
	TMC5160_OFS253_FIELD = RegisterField{
		Mask:     TMC5160_OFS253_MASK,
		Shift:    TMC5160_OFS253_SHIFT,
		Register: TMC5160_MSLUT7,
		IsSigned: false,
	}
	TMC5160_OFS254_FIELD = RegisterField{
		Mask:     TMC5160_OFS254_MASK,
		Shift:    TMC5160_OFS254_SHIFT,
		Register: TMC5160_MSLUT7,
		IsSigned: false,
	}
	TMC5160_OFS255_FIELD = RegisterField{
		Mask:     TMC5160_OFS255_MASK,
		Shift:    TMC5160_OFS255_SHIFT,
		Register: TMC5160_MSLUT7,
		IsSigned: false,
	}
	TMC5160_W0_FIELD = RegisterField{
		Mask:     TMC5160_W0_MASK,
		Shift:    TMC5160_W0_SHIFT,
		Register: TMC5160_MSLUTSEL,
		IsSigned: false,
	}
	TMC5160_W1_FIELD = RegisterField{
		Mask:     TMC5160_W1_MASK,
		Shift:    TMC5160_W1_SHIFT,
		Register: TMC5160_MSLUTSEL,
		IsSigned: false,
	}
	TMC5160_W2_FIELD = RegisterField{
		Mask:     TMC5160_W2_MASK,
		Shift:    TMC5160_W2_SHIFT,
		Register: TMC5160_MSLUTSEL,
		IsSigned: false,
	}
	TMC5160_W3_FIELD = RegisterField{
		Mask:     TMC5160_W3_MASK,
		Shift:    TMC5160_W3_SHIFT,
		Register: TMC5160_MSLUTSEL,
		IsSigned: false,
	}
	TMC5160_X1_FIELD = RegisterField{
		Mask:     TMC5160_X1_MASK,
		Shift:    TMC5160_X1_SHIFT,
		Register: TMC5160_MSLUTSEL,
		IsSigned: false,
	}
	TMC5160_X2_FIELD = RegisterField{
		Mask:     TMC5160_X2_MASK,
		Shift:    TMC5160_X2_SHIFT,
		Register: TMC5160_MSLUTSEL,
		IsSigned: false,
	}
	TMC5160_X3_FIELD = RegisterField{
		Mask:     TMC5160_X3_MASK,
		Shift:    TMC5160_X3_SHIFT,
		Register: TMC5160_MSLUTSEL,
		IsSigned: false,
	}
	TMC5160_START_SIN_FIELD = RegisterField{
		Mask:     TMC5160_START_SIN_MASK,
		Shift:    TMC5160_START_SIN_SHIFT,
		Register: TMC5160_MSLUTSTART,
		IsSigned: false,
	}
	TMC5160_START_SIN90_FIELD = RegisterField{
		Mask:     TMC5160_START_SIN90_MASK,
		Shift:    TMC5160_START_SIN90_SHIFT,
		Register: TMC5160_MSLUTSTART,
		IsSigned: false,
	}
	TMC5160_MSCNT_FIELD = RegisterField{
		Mask:     TMC5160_MSCNT_MASK,
		Shift:    TMC5160_MSCNT_SHIFT,
		Register: TMC5160_MSCNT,
		IsSigned: false,
	}
	TMC5160_CUR_A_FIELD = RegisterField{
		Mask:     TMC5160_CUR_A_MASK,
		Shift:    TMC5160_CUR_A_SHIFT,
		Register: TMC5160_MSCURACT,
		IsSigned: true,
	}
	TMC5160_CUR_B_FIELD = RegisterField{
		Mask:     TMC5160_CUR_B_MASK,
		Shift:    TMC5160_CUR_B_SHIFT,
		Register: TMC5160_MSCURACT,
		IsSigned: true,
	}

	TMC5160_TFD_2__0__FIELD = RegisterField{
		Mask:     TMC5160_TFD_2__0__MASK,
		Shift:    TMC5160_TFD_2__0__SHIFT,
		Register: TMC5160_CHOPCONF,
		IsSigned: false,
	}

	TMC5160_TFD___FIELD = RegisterField{
		Mask:     TMC5160_TFD___MASK,
		Shift:    TMC5160_TFD___SHIFT,
		Register: TMC5160_CHOPCONF,
		IsSigned: false,
	}

	TMC5160_VHIGHFS_FIELD = RegisterField{
		Mask:     TMC5160_VHIGHFS_MASK,
		Shift:    TMC5160_VHIGHFS_SHIFT,
		Register: TMC5160_CHOPCONF,
		IsSigned: false,
	}
	TMC5160_VHIGHCHM_FIELD = RegisterField{
		Mask:     TMC5160_VHIGHCHM_MASK,
		Shift:    TMC5160_VHIGHCHM_SHIFT,
		Register: TMC5160_CHOPCONF,
		IsSigned: false,
	}
	TMC5160_TPFD_FIELD = RegisterField{
		Mask:     TMC5160_TPFD_MASK,
		Shift:    TMC5160_TPFD_SHIFT,
		Register: TMC5160_CHOPCONF,
		IsSigned: false,
	}
	TMC5160_MRES_FIELD = RegisterField{
		Mask:     TMC5160_MRES_MASK,
		Shift:    TMC5160_MRES_SHIFT,
		Register: TMC5160_CHOPCONF,
		IsSigned: false,
	}
	TMC5160_INTPOL_FIELD = RegisterField{
		Mask:     TMC5160_INTPOL_MASK,
		Shift:    TMC5160_INTPOL_SHIFT,
		Register: TMC5160_CHOPCONF,
		IsSigned: false,
	}
	TMC5160_DEDGE_FIELD = RegisterField{
		Mask:     TMC5160_DEDGE_MASK,
		Shift:    TMC5160_DEDGE_SHIFT,
		Register: TMC5160_CHOPCONF,
		IsSigned: false,
	}
	TMC5160_DISS2G_FIELD = RegisterField{
		Mask:     TMC5160_DISS2G_MASK,
		Shift:    TMC5160_DISS2G_SHIFT,
		Register: TMC5160_CHOPCONF,
		IsSigned: false,
	}
	TMC5160_DISS2VS_FIELD = RegisterField{
		Mask:     TMC5160_DISS2VS_MASK,
		Shift:    TMC5160_DISS2VS_SHIFT,
		Register: TMC5160_CHOPCONF,
		IsSigned: false,
	}
	TMC5160_TOFF_FIELD = RegisterField{
		Mask:     TMC5160_TOFF_MASK,
		Shift:    TMC5160_TOFF_SHIFT,
		Register: TMC5160_CHOPCONF,
		IsSigned: false,
	}
	TMC5160_TFD_ALL_FIELD = RegisterField{
		Mask:     TMC5160_TFD_2__0__MASK,
		Shift:    TMC5160_TFD_2__0__SHIFT,
		Register: TMC5160_CHOPCONF,
		IsSigned: false,
	}
	TMC5160_OFFSET_FIELD = RegisterField{
		Mask:     TMC5160_OFFSET_MASK,
		Shift:    TMC5160_OFFSET_SHIFT,
		Register: TMC5160_CHOPCONF,
		IsSigned: false,
	}
	TMC5160_TFD_3_FIELD = RegisterField{
		Mask:     TMC5160_TFD___MASK,
		Shift:    TMC5160_TFD___SHIFT,
		Register: TMC5160_CHOPCONF,
		IsSigned: false,
	}
	TMC5160_DISFDCC_FIELD = RegisterField{
		Mask:     TMC5160_DISFDCC_MASK,
		Shift:    TMC5160_DISFDCC_SHIFT,
		Register: TMC5160_CHOPCONF,
		IsSigned: false,
	}
	TMC5160_RNDTF_FIELD = RegisterField{
		Mask:     TMC5160_RNDTF_MASK,
		Shift:    TMC5160_RNDTF_SHIFT,
		Register: TMC5160_CHOPCONF,
		IsSigned: false,
	}
	TMC5160_CHM_FIELD = RegisterField{
		Mask:     TMC5160_CHM_MASK,
		Shift:    TMC5160_CHM_SHIFT,
		Register: TMC5160_CHOPCONF,
		IsSigned: false,
	}
	TMC5160_TBL_FIELD = RegisterField{
		Mask:     TMC5160_TBL_MASK,
		Shift:    TMC5160_TBL_SHIFT,
		Register: TMC5160_CHOPCONF,
		IsSigned: false,
	}
	TMC5160_VSENSE_FIELD = RegisterField{
		Mask:     TMC5160_VSENSE_MASK,
		Shift:    TMC5160_VSENSE_SHIFT,
		Register: TMC5160_CHOPCONF,
		IsSigned: false,
	}

	TMC5160_HSTRT_FIELD = RegisterField{
		Mask:     TMC5160_HSTRT_MASK,
		Shift:    TMC5160_HSTRT_SHIFT,
		Register: TMC5160_CHOPCONF,
		IsSigned: false,
	}
	TMC5160_HEND_FIELD = RegisterField{
		Mask:     TMC5160_HEND_MASK,
		Shift:    TMC5160_HEND_SHIFT,
		Register: TMC5160_CHOPCONF,
		IsSigned: false,
	}

	TMC5160_SEMIN_FIELD = RegisterField{
		Mask:     TMC5160_SEMIN_MASK,
		Shift:    TMC5160_SEMIN_SHIFT,
		Register: TMC5160_COOLCONF,
		IsSigned: false,
	}
	TMC5160_SEUP_FIELD = RegisterField{
		Mask:     TMC5160_SEUP_MASK,
		Shift:    TMC5160_SEUP_SHIFT,
		Register: TMC5160_COOLCONF,
		IsSigned: false,
	}
	TMC5160_SEMAX_FIELD = RegisterField{
		Mask:     TMC5160_SEMAX_MASK,
		Shift:    TMC5160_SEMAX_SHIFT,
		Register: TMC5160_COOLCONF,
		IsSigned: false,
	}
	TMC5160_SEDN_FIELD = RegisterField{
		Mask:     TMC5160_SEDN_MASK,
		Shift:    TMC5160_SEDN_SHIFT,
		Register: TMC5160_COOLCONF,
		IsSigned: false,
	}
	TMC5160_SEIMIN_FIELD = RegisterField{
		Mask:     TMC5160_SEIMIN_MASK,
		Shift:    TMC5160_SEIMIN_SHIFT,
		Register: TMC5160_COOLCONF,
		IsSigned: false,
	}
	TMC5160_SGT_FIELD = RegisterField{
		Mask:     TMC5160_SGT_MASK,
		Shift:    TMC5160_SGT_SHIFT,
		Register: TMC5160_COOLCONF,
		IsSigned: true,
	}
	TMC5160_SFILT_FIELD = RegisterField{
		Mask:     TMC5160_SFILT_MASK,
		Shift:    TMC5160_SFILT_SHIFT,
		Register: TMC5160_COOLCONF,
		IsSigned: false,
	}
	TMC5160_DC_TIME_FIELD = RegisterField{
		Mask:     TMC5160_DC_TIME_MASK,
		Shift:    TMC5160_DC_TIME_SHIFT,
		Register: TMC5160_DCCTRL,
		IsSigned: false,
	}
	TMC5160_DC_SG_FIELD = RegisterField{
		Mask:     TMC5160_DC_SG_MASK,
		Shift:    TMC5160_DC_SG_SHIFT,
		Register: TMC5160_DCCTRL,
		IsSigned: false,
	}
	TMC5160_SG_RESULT_FIELD = RegisterField{
		Mask:     TMC5160_SG_RESULT_MASK,
		Shift:    TMC5160_SG_RESULT_SHIFT,
		Register: TMC5160_DRV_STATUS,
		IsSigned: false,
	}
	TMC5160_S2VSA_FIELD = RegisterField{
		Mask:     TMC5160_S2VSA_MASK,
		Shift:    TMC5160_S2VSA_SHIFT,
		Register: TMC5160_DRV_STATUS,
		IsSigned: false,
	}
	TMC5160_S2VSB_FIELD = RegisterField{
		Mask:     TMC5160_S2VSB_MASK,
		Shift:    TMC5160_S2VSB_SHIFT,
		Register: TMC5160_DRV_STATUS,
		IsSigned: false,
	}
	TMC5160_STEALTH_FIELD = RegisterField{
		Mask:     TMC5160_STEALTH_MASK,
		Shift:    TMC5160_STEALTH_SHIFT,
		Register: TMC5160_DRV_STATUS,
		IsSigned: false,
	}
	TMC5160_FSACTIVE_FIELD = RegisterField{
		Mask:     TMC5160_FSACTIVE_MASK,
		Shift:    TMC5160_FSACTIVE_SHIFT,
		Register: TMC5160_DRV_STATUS,
		IsSigned: false,
	}
	TMC5160_CS_ACTUAL_FIELD = RegisterField{
		Mask:     TMC5160_CS_ACTUAL_MASK,
		Shift:    TMC5160_CS_ACTUAL_SHIFT,
		Register: TMC5160_DRV_STATUS,
		IsSigned: false,
	}
	TMC5160_STALLGUARD_FIELD = RegisterField{
		Mask:     TMC5160_STALLGUARD_MASK,
		Shift:    TMC5160_STALLGUARD_SHIFT,
		Register: TMC5160_DRV_STATUS,
		IsSigned: false,
	}
	TMC5160_OT_FIELD = RegisterField{
		Mask:     TMC5160_OT_MASK,
		Shift:    TMC5160_OT_SHIFT,
		Register: TMC5160_DRV_STATUS,
		IsSigned: false,
	}
	TMC5160_OTPW_FIELD = RegisterField{
		Mask:     TMC5160_OTPW_MASK,
		Shift:    TMC5160_OTPW_SHIFT,
		Register: TMC5160_DRV_STATUS,
		IsSigned: false,
	}
	TMC5160_S2GA_FIELD = RegisterField{
		Mask:     TMC5160_S2GA_MASK,
		Shift:    TMC5160_S2GA_SHIFT,
		Register: TMC5160_DRV_STATUS,
		IsSigned: false,
	}
	TMC5160_S2GB_FIELD = RegisterField{
		Mask:     TMC5160_S2GB_MASK,
		Shift:    TMC5160_S2GB_SHIFT,
		Register: TMC5160_DRV_STATUS,
		IsSigned: false,
	}
	TMC5160_OLA_FIELD = RegisterField{
		Mask:     TMC5160_OLA_MASK,
		Shift:    TMC5160_OLA_SHIFT,
		Register: TMC5160_DRV_STATUS,
		IsSigned: false,
	}
	TMC5160_OLB_FIELD = RegisterField{
		Mask:     TMC5160_OLB_MASK,
		Shift:    TMC5160_OLB_SHIFT,
		Register: TMC5160_DRV_STATUS,
		IsSigned: false,
	}
	TMC5160_STST_FIELD = RegisterField{
		Mask:     TMC5160_STST_MASK,
		Shift:    TMC5160_STST_SHIFT,
		Register: TMC5160_DRV_STATUS,
		IsSigned: false,
	}
	TMC5160_PWM_OFS_FIELD = RegisterField{
		Mask:     TMC5160_PWM_OFS_MASK,
		Shift:    TMC5160_PWM_OFS_SHIFT,
		Register: TMC5160_PWMCONF,
		IsSigned: false,
	}
	TMC5160_PWM_GRAD_FIELD = RegisterField{
		Mask:     TMC5160_PWM_GRAD_MASK,
		Shift:    TMC5160_PWM_GRAD_SHIFT,
		Register: TMC5160_PWMCONF,
		IsSigned: false,
	}
	TMC5160_PWM_FREQ_FIELD = RegisterField{
		Mask:     TMC5160_PWM_FREQ_MASK,
		Shift:    TMC5160_PWM_FREQ_SHIFT,
		Register: TMC5160_PWMCONF,
		IsSigned: false,
	}
	TMC5160_PWM_AUTOSCALE_FIELD = RegisterField{
		Mask:     TMC5160_PWM_AUTOSCALE_MASK,
		Shift:    TMC5160_PWM_AUTOSCALE_SHIFT,
		Register: TMC5160_PWMCONF,
		IsSigned: false,
	}
	TMC5160_PWM_AUTOGRAD_FIELD = RegisterField{
		Mask:     TMC5160_PWM_AUTOGRAD_MASK,
		Shift:    TMC5160_PWM_AUTOGRAD_SHIFT,
		Register: TMC5160_PWMCONF,
		IsSigned: false,
	}
	TMC5160_FREEWHEEL_FIELD = RegisterField{
		Mask:     TMC5160_FREEWHEEL_MASK,
		Shift:    TMC5160_FREEWHEEL_SHIFT,
		Register: TMC5160_PWMCONF,
		IsSigned: false,
	}
	TMC5160_PWM_REG_FIELD = RegisterField{
		Mask:     TMC5160_PWM_REG_MASK,
		Shift:    TMC5160_PWM_REG_SHIFT,
		Register: TMC5160_PWMCONF,
		IsSigned: false,
	}
	TMC5160_PWM_LIM_FIELD = RegisterField{
		Mask:     TMC5160_PWM_LIM_MASK,
		Shift:    TMC5160_PWM_LIM_SHIFT,
		Register: TMC5160_PWMCONF,
		IsSigned: false,
	}
	TMC5160_PWM_SCALE_SUM_FIELD = RegisterField{
		Mask:     TMC5160_PWM_SCALE_SUM_MASK,
		Shift:    TMC5160_PWM_SCALE_SUM_SHIFT,
		Register: TMC5160_PWMSCALE,
		IsSigned: false,
	}
	TMC5160_PWM_SCALE_AUTO_FIELD = RegisterField{
		Mask:     TMC5160_PWM_SCALE_AUTO_MASK,
		Shift:    TMC5160_PWM_SCALE_AUTO_SHIFT,
		Register: TMC5160_PWMSCALE,
		IsSigned: true,
	}
	TMC5160_PWM_OFS_AUTO_FIELD = RegisterField{
		Mask:     TMC5160_PWM_OFS_AUTO_MASK,
		Shift:    TMC5160_PWM_OFS_AUTO_SHIFT,
		Register: TMC5160_PWM_AUTO,
		IsSigned: false,
	}
	TMC5160_PWM_GRAD_AUTO_FIELD = RegisterField{
		Mask:     TMC5160_PWM_GRAD_AUTO_MASK,
		Shift:    TMC5160_PWM_GRAD_AUTO_SHIFT,
		Register: TMC5160_PWM_AUTO,
		IsSigned: false,
	}
	TMC5160_LOST_STEPS_FIELD = RegisterField{
		Mask:     TMC5160_LOST_STEPS_MASK,
		Shift:    TMC5160_LOST_STEPS_SHIFT,
		Register: TMC5160_LOST_STEPS,
		IsSigned: false,
	}
)
