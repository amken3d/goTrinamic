// Partially Generated Go code
// Amken LLC

package tmc2240

type RegisterField struct {
	Name     string
	Mask     uint32
	Shift    uint32
	Register uint32
	IsSigned bool
}

// Constants
const (
	TMC2240_WRITE_BIT                           = 0x80
	TMC2240_ADDRESS_MASK                        = 0x7F
	TMC2240_MAX_VELOCITY                        = 8388096
	TMC2240_MAX_ACCELERATION             uint16 = 65535
	TMC2240_MODE_POSITION                       = 0
	TMC2240_MODE_VELPOS                         = 1
	TMC2240_MODE_VELNEG                         = 2
	TMC2240_MODE_HOLD                           = 3
	TMC2240_SW_STOPL_ENABLE                     = 0x0001
	TMC2240_SW_STOPR_ENABLE                     = 0x0002
	TMC2240_SW_STOPL_POLARITY                   = 0x0004
	TMC2240_SW_STOPR_POLARITY                   = 0x0008
	TMC2240_SW_SWAP_LR                          = 0x0010
	TMC2240_SW_LATCH_L_ACT                      = 0x0020
	TMC2240_SW_LATCH_L_INACT                    = 0x0040
	TMC2240_SW_LATCH_R_ACT                      = 0x0080
	TMC2240_SW_LATCH_R_INACT                    = 0x0100
	TMC2240_SW_LATCH_ENC                        = 0x0200
	TMC2240_SW_SG_STOP                          = 0x0400
	TMC2240_SW_SOFTSTOP                         = 0x0800
	TMC2240_RS_STOPL                            = 0x0001
	TMC2240_RS_STOPR                            = 0x0002
	TMC2240_RS_LATCHL                           = 0x0004
	TMC2240_RS_LATCHR                           = 0x0008
	TMC2240_RS_EV_STOPL                         = 0x0010
	TMC2240_RS_EV_STOPR                         = 0x0020
	TMC2240_RS_EV_STOP_SG                       = 0x0040
	TMC2240_RS_EV_POSREACHED                    = 0x0080
	TMC2240_RS_VELREACHED                       = 0x0100
	TMC2240_RS_POSREACHED                       = 0x0200
	TMC2240_RS_VZERO                            = 0x0400
	TMC2240_RS_ZEROWAIT                         = 0x0800
	TMC2240_RS_SECONDMOVE                       = 0x1000
	TMC2240_RS_SG                               = 0x2000
	TMC2240_EM_DECIMAL                          = 0x0400
	TMC2240_EM_LATCH_XACT                       = 0x0200
	TMC2240_EM_CLR_XENC                         = 0x0100
	TMC2240_EM_NEG_EDGE                         = 0x0080
	TMC2240_EM_POS_EDGE                         = 0x0040
	TMC2240_EM_CLR_ONCE                         = 0x0020
	TMC2240_EM_CLR_CONT                         = 0x0010
	TMC2240_EM_IGNORE_AB                        = 0x0008
	TMC2240_EM_POL_N                            = 0x0004
	TMC2240_EM_POL_B                            = 0x0002
	TMC2240_EM_POL_A                            = 0x0001
	TMC2240_GCONF                               = 0x00
	TMC2240_GSTAT                               = 0x01
	TMC2240_IFCNT                               = 0x02
	TMC2240_SLAVECONF                           = 0x03
	TMC2240_IOIN                                = 0x04
	TMC2240_DRV_CONF                            = 0x0A
	TMC2240_GLOBAL_SCALER                       = 0x0B
	TMC2240_IHOLD_IRUN                          = 0x10
	TMC2240_TPOWERDOWN                          = 0x11
	TMC2240_TSTEP                               = 0x12
	TMC2240_TPWMTHRS                            = 0x13
	TMC2240_TCOOLTHRS                           = 0x14
	TMC2240_THIGH                               = 0x15
	TMC2240_DIRECT_MODE                         = 0x2D
	TMC2240_ENCMODE                             = 0x38
	TMC2240_XENC                                = 0x39
	TMC2240_ENC_CONST                           = 0x3A
	TMC2240_ENC_STATUS                          = 0x3B
	TMC2240_ENC_LATCH                           = 0x3C
	TMC2240_ADC_VSUPPLY_AIN                     = 0x50
	TMC2240_ADC_TEMP                            = 0x51
	TMC2240_OTW_OV_VTH                          = 0x52
	TMC2240_MSLUT0                              = 0x60
	TMC2240_MSLUT1                              = 0x61
	TMC2240_MSLUT2                              = 0x62
	TMC2240_MSLUT3                              = 0x63
	TMC2240_MSLUT4                              = 0x64
	TMC2240_MSLUT5                              = 0x65
	TMC2240_MSLUT6                              = 0x66
	TMC2240_MSLUT7                              = 0x67
	TMC2240_MSLUTSEL                            = 0x68
	TMC2240_MSLUTSTART                          = 0x69
	TMC2240_MSCNT                               = 0x6A
	TMC2240_MSCURACT                            = 0x6B
	TMC2240_CHOPCONF                            = 0x6C
	TMC2240_COOLCONF                            = 0x6D
	TMC2240_DRVSTATUS                           = 0x6F
	TMC2240_PWMCONF                             = 0x70
	TMC2240_PWMSCALE                            = 0x71
	TMC2240_PWM_AUTO                            = 0x72
	TMC2240_SG4_THRS                            = 0x74
	TMC2240_SG4_RESULT                          = 0x75
	TMC2240_SG4_IND                             = 0x76
	TMC2240_FAST_STANDSTILL_MASK                = 0x00000002
	TMC2240_FAST_STANDSTILL_SHIFT               = 1
	TMC2240_EN_PWM_MODE_MASK                    = 0x00000004
	TMC2240_EN_PWM_MODE_SHIFT                   = 2
	TMC2240_MULTISTEP_FILT_MASK                 = 0x00000008
	TMC2240_MULTISTEP_FILT_SHIFT                = 3
	TMC2240_SHAFT_MASK                          = 0x00000010
	TMC2240_SHAFT_SHIFT                         = 4
	TMC2240_DIAG0_ERROR_MASK                    = 0x00000020
	TMC2240_DIAG0_ERROR_SHIFT                   = 5
	TMC2240_DIAG0_OTPW_MASK                     = 0x00000040
	TMC2240_DIAG0_OTPW_SHIFT                    = 6
	TMC2240_DIAG0_STALL_MASK                    = 0x00000080
	TMC2240_DIAG0_STALL_SHIFT                   = 7
	TMC2240_DIAG1_STALL_MASK                    = 0x00000100
	TMC2240_DIAG1_STALL_SHIFT                   = 8
	TMC2240_DIAG1_INDEX_MASK                    = 0x00000200
	TMC2240_DIAG1_INDEX_SHIFT                   = 9
	TMC2240_DIAG1_ONSTATE_MASK                  = 0x00000400
	TMC2240_DIAG1_ONSTATE_SHIFT                 = 10
	TMC2240_DIAG0_PUSHPULL_MASK                 = 0x00001000
	TMC2240_DIAG0_PUSHPULL_SHIFT                = 12
	TMC2240_DIAG1_PUSHPULL_MASK                 = 0x00002000
	TMC2240_DIAG1_PUSHPULL_SHIFT                = 13
	TMC2240_SMALL_HYSTERESIS_MASK               = 0x00004000
	TMC2240_SMALL_HYSTERESIS_SHIFT              = 14
	TMC2240_STOP_ENABLE_MASK                    = 0x00008000
	TMC2240_STOP_ENABLE_SHIFT                   = 15
	TMC2240_DIRECT_MODE_MASK                    = 0x00010000
	TMC2240_DIRECT_MODE_SHIFT                   = 16
	TMC2240_RESET_MASK                          = 0x00000001
	TMC2240_RESET_SHIFT                         = 0
	TMC2240_DRV_ERR_MASK                        = 0x00000002
	TMC2240_DRV_ERR_SHIFT                       = 1
	TMC2240_UV_CP_MASK                          = 0x00000004
	TMC2240_UV_CP_SHIFT                         = 2
	TMC2240_REGISTER_RESET_MASK                 = 0x00000008
	TMC2240_REGISTER_RESET_SHIFT                = 3
	TMC2240_VM_UVLO_MASK                        = 0x00000010
	TMC2240_VM_UVLO_SHIFT                       = 4
	TMC2240_IFCNT_MASK                          = 0x000000FF
	TMC2240_IFCNT_SHIFT                         = 0
	TMC2240_SLAVEADDR_MASK                      = 0x000000FF
	TMC2240_SLAVEADDR_SHIFT                     = 0
	TMC2240_SENDDELAY_MASK                      = 0x00000F00
	TMC2240_SENDDELAY_SHIFT                     = 8
	TMC2240_STEP_MASK                           = 0x00000001
	TMC2240_STEP_SHIFT                          = 0
	TMC2240_DIR_MASK                            = 0x00000002
	TMC2240_DIR_SHIFT                           = 1
	TMC2240_ENCB_MASK                           = 0x00000004
	TMC2240_ENCB_SHIFT                          = 2
	TMC2240_ENCA_MASK                           = 0x00000008
	TMC2240_ENCA_SHIFT                          = 3
	TMC2240_DRV_ENN_MASK                        = 0x00000010
	TMC2240_DRV_ENN_SHIFT                       = 4
	TMC2240_ENCN_MASK                           = 0x00000020
	TMC2240_ENCN_SHIFT                          = 5
	TMC2240_UART_EN_MASK                        = 0x00000040
	TMC2240_UART_EN_SHIFT                       = 6
	TMC2240_RESERVED_MASK                       = 0x00000080
	TMC2240_RESERVED_SHIFT                      = 7
	TMC2240_COMP_A_MASK                         = 0x00000100
	TMC2240_COMP_A_SHIFT                        = 8
	TMC2240_COMP_B_MASK                         = 0x00000200
	TMC2240_COMP_B_SHIFT                        = 9
	TMC2240_COMP_A1_A2_MASK                     = 0x00000400
	TMC2240_COMP_A1_A2_SHIFT                    = 10
	TMC2240_COMP_B1_B2_MASK                     = 0x00000800
	TMC2240_COMP_B1_B2_SHIFT                    = 11
	TMC2240_OUTPUT_MASK                         = 0x00001000
	TMC2240_OUTPUT_SHIFT                        = 12
	TMC2240_EXT_RES_DET_MASK                    = 0x00002000
	TMC2240_EXT_RES_DET_SHIFT                   = 13
	TMC2240_EXT_CLK_MASK                        = 0x00004000
	TMC2240_EXT_CLK_SHIFT                       = 14
	TMC2240_ADC_ERR_MASK                        = 0x00008000
	TMC2240_ADC_ERR_SHIFT                       = 15
	TMC2240_SILICON_RV_MASK                     = 0x00070000
	TMC2240_SILICON_RV_SHIFT                    = 16
	TMC2240_VERSION_MASK                        = 0xFF000000
	TMC2240_VERSION_SHIFT                       = 24
	TMC2240_CURRENT_RANGE_MASK                  = 0x00000003
	TMC2240_CURRENT_RANGE_SHIFT                 = 0
	TMC2240_SLOPE_CONTROL_MASK                  = 0x00000030
	TMC2240_SLOPE_CONTROL_SHIFT                 = 4
	TMC2240_GLOBALSCALER_MASK                   = 0x000000FF
	TMC2240_GLOBALSCALER_SHIFT                  = 0
	TMC2240_IHOLD_MASK                          = 0x0000001F
	TMC2240_IHOLD_SHIFT                         = 0
	TMC2240_IRUN_MASK                           = 0x00001F00
	TMC2240_IRUN_SHIFT                          = 8
	TMC2240_IHOLDDELAY_MASK                     = 0x000F0000
	TMC2240_IHOLDDELAY_SHIFT                    = 16
	TMC2240_IRUNDELAY_MASK                      = 0x0F000000
	TMC2240_IRUNDELAY_SHIFT                     = 24
	TMC2240_TPOWERDOWN_MASK                     = 0x000000FF
	TMC2240_TPOWERDOWN_SHIFT                    = 0
	TMC2240_TSTEP_MASK                          = 0x000FFFFF
	TMC2240_TSTEP_SHIFT                         = 0
	TMC2240_TPWMTHRS_MASK                       = 0x000FFFFF
	TMC2240_TPWMTHRS_SHIFT                      = 0
	TMC2240_TCOOLTHRS_MASK                      = 0x000FFFFF
	TMC2240_TCOOLTHRS_SHIFT                     = 0
	TMC2240_THIGH_MASK                          = 0x000FFFFF
	TMC2240_THIGH_SHIFT                         = 0
	TMC2240_DIRECT_COIL_A_MASK                  = 0x000001FF
	TMC2240_DIRECT_COIL_A_SHIFT                 = 0
	TMC2240_DIRECT_COIL_B_MASK                  = 0x01FF0000
	TMC2240_DIRECT_COIL_B_SHIFT                 = 16
	TMC2240_POL_A_MASK                          = 0x00000001
	TMC2240_POL_A_SHIFT                         = 0
	TMC2240_POL_B_MASK                          = 0x00000002
	TMC2240_POL_B_SHIFT                         = 1
	TMC2240_POL_N_MASK                          = 0x00000004
	TMC2240_POL_N_SHIFT                         = 2
	TMC2240_IGNORE_AB_MASK                      = 0x00000008
	TMC2240_IGNORE_AB_SHIFT                     = 3
	TMC2240_CLR_CONT_MASK                       = 0x00000010
	TMC2240_CLR_CONT_SHIFT                      = 4
	TMC2240_CLR_ONCE_MASK                       = 0x00000020
	TMC2240_CLR_ONCE_SHIFT                      = 5
	TMC2240_POS_NEG_EDGE_MASK                   = 0x000000C0
	TMC2240_POS_NEG_EDGE_SHIFT                  = 6
	TMC2240_CLR_ENC_X_MASK                      = 0x00000100
	TMC2240_CLR_ENC_X_SHIFT                     = 8
	TMC2240_LATCH_X_ACT_MASK                    = 0x00000200
	TMC2240_LATCH_X_ACT_SHIFT                   = 9
	TMC2240_ENC_SEL_DECIMAL_MASK                = 0x00000400
	TMC2240_ENC_SEL_DECIMAL_SHIFT               = 10
	TMC2240_X_ENC_MASK                          = 0xFFFFFFFF
	TMC2240_X_ENC_SHIFT                         = 0
	TMC2240_ENC_CONST_MASK                      = 0xFFFFFFFF
	TMC2240_ENC_CONST_SHIFT                     = 0
	TMC2240_N_EVENT_MASK                        = 0x00000001
	TMC2240_N_EVENT_SHIFT                       = 0
	TMC2240_DEVIATION_WARN_MASK                 = 0x00000002
	TMC2240_DEVIATION_WARN_SHIFT                = 1
	TMC2240_ENC_LATCH_MASK                      = 0xFFFFFFFF
	TMC2240_ENC_LATCH_SHIFT                     = 0
	TMC2240_ADC_VSUPPLY_MASK                    = 0x00001FFF
	TMC2240_ADC_VSUPPLY_SHIFT                   = 0
	TMC2240_ADC_AIN_MASK                        = 0x1FFF0000
	TMC2240_ADC_AIN_SHIFT                       = 16
	TMC2240_ADC_TEMP_MASK                       = 0x00001FFF
	TMC2240_ADC_TEMP_SHIFT                      = 0
	TMC2240_OVERVOLTAGE_VTH_MASK                = 0x00001FFF
	TMC2240_OVERVOLTAGE_VTH_SHIFT               = 0
	TMC2240_OVERTEMPPREWARNING_VTH_MASK         = 0x1FFF0000
	TMC2240_OVERTEMPPREWARNING_VTH_SHIFT        = 16
	TMC2240_MSLUT___MASK                        = 0xFFFFFFFF
	TMC2240_MSLUT___SHIFT                       = 0
	TMC2240_W0_MASK                             = 0x00000003
	TMC2240_W0_SHIFT                            = 0
	TMC2240_W1_MASK                             = 0x0000000C
	TMC2240_W1_SHIFT                            = 2
	TMC2240_W2_MASK                             = 0x00000030
	TMC2240_W2_SHIFT                            = 4
	TMC2240_W3_MASK                             = 0x000000C0
	TMC2240_W3_SHIFT                            = 6
	TMC2240_X1_MASK                             = 0x0000FF00
	TMC2240_X1_SHIFT                            = 8
	TMC2240_X2_MASK                             = 0x00FF0000
	TMC2240_X2_SHIFT                            = 16
	TMC2240_X3_MASK                             = 0xFF000000
	TMC2240_X3_SHIFT                            = 24
	TMC2240_START_SIN_MASK                      = 0x000000FF
	TMC2240_START_SIN_SHIFT                     = 0
	TMC2240_START_SIN90_MASK                    = 0x00FF0000
	TMC2240_START_SIN90_SHIFT                   = 16
	TMC2240_OFFSET_SIN90_MASK                   = 0xFF000000
	TMC2240_OFFSET_SIN90_SHIFT                  = 24
	TMC2240_MSCNT_MASK                          = 0x000003FF
	TMC2240_MSCNT_SHIFT                         = 0
	TMC2240_CUR_B_MASK                          = 0x000001FF
	TMC2240_CUR_B_SHIFT                         = 0
	TMC2240_CUR_A_MASK                          = 0x01FF0000
	TMC2240_CUR_A_SHIFT                         = 16
	TMC2240_TOFF_MASK                           = 0x0000000F
	TMC2240_TOFF_SHIFT                          = 0
	TMC2240_HSTRT_TFD210_MASK                   = 0x00000070
	TMC2240_HSTRT_TFD210_SHIFT                  = 4
	TMC2240_HEND_OFFSET_MASK                    = 0x00000780
	TMC2240_HEND_OFFSET_SHIFT                   = 7
	TMC2240_FD3_MASK                            = 0x00000800
	TMC2240_FD3_SHIFT                           = 11
	TMC2240_DISFDCC_MASK                        = 0x00001000
	TMC2240_DISFDCC_SHIFT                       = 12
	TMC2240_CHM_MASK                            = 0x00004000
	TMC2240_CHM_SHIFT                           = 14
	TMC2240_TBL_MASK                            = 0x00018000
	TMC2240_TBL_SHIFT                           = 15
	TMC2240_VHIGHFS_MASK                        = 0x00040000
	TMC2240_VHIGHFS_SHIFT                       = 18
	TMC2240_VHIGHCHM_MASK                       = 0x00080000
	TMC2240_VHIGHCHM_SHIFT                      = 19
	TMC2240_TPFD_MASK                           = 0x00F00000
	TMC2240_TPFD_SHIFT                          = 20
	TMC2240_MRES_MASK                           = 0x0F000000
	TMC2240_MRES_SHIFT                          = 24
	TMC2240_INTPOL_MASK                         = 0x10000000
	TMC2240_INTPOL_SHIFT                        = 28
	TMC2240_DEDGE_MASK                          = 0x20000000
	TMC2240_DEDGE_SHIFT                         = 29
	TMC2240_DISS2G_MASK                         = 0x40000000
	TMC2240_DISS2G_SHIFT                        = 30
	TMC2240_DISS2VS_MASK                        = 0x80000000
	TMC2240_DISS2VS_SHIFT                       = 31
	TMC2240_SEMIN_MASK                          = 0x0000000F
	TMC2240_SEMIN_SHIFT                         = 0
	TMC2240_SEUP_MASK                           = 0x00000060
	TMC2240_SEUP_SHIFT                          = 5
	TMC2240_SEMAX_MASK                          = 0x00000F00
	TMC2240_SEMAX_SHIFT                         = 8
	TMC2240_SEDN_MASK                           = 0x00006000
	TMC2240_SEDN_SHIFT                          = 13
	TMC2240_SEIMIN_MASK                         = 0x00008000
	TMC2240_SEIMIN_SHIFT                        = 15
	TMC2240_SGT_MASK                            = 0x007F0000
	TMC2240_SGT_SHIFT                           = 16
	TMC2240_SFILT_MASK                          = 0x01000000
	TMC2240_SFILT_SHIFT                         = 24
	TMC2240_SG_RESULT_MASK                      = 0x000003FF
	TMC2240_SG_RESULT_SHIFT                     = 0
	TMC2240_S2VSA_MASK                          = 0x00001000
	TMC2240_S2VSA_SHIFT                         = 12
	TMC2240_S2VSB_MASK                          = 0x00002000
	TMC2240_S2VSB_SHIFT                         = 13
	TMC2240_STEALTH_MASK                        = 0x00004000
	TMC2240_STEALTH_SHIFT                       = 14
	TMC2240_FSACTIVE_MASK                       = 0x00008000
	TMC2240_FSACTIVE_SHIFT                      = 15
	TMC2240_CS_ACTUAL_MASK                      = 0x001F0000
	TMC2240_CS_ACTUAL_SHIFT                     = 16
	TMC2240_STALLGUARD_MASK                     = 0x01000000
	TMC2240_STALLGUARD_SHIFT                    = 24
	TMC2240_OT_MASK                             = 0x02000000
	TMC2240_OT_SHIFT                            = 25
	TMC2240_OTPW_MASK                           = 0x04000000
	TMC2240_OTPW_SHIFT                          = 26
	TMC2240_S2GA_MASK                           = 0x08000000
	TMC2240_S2GA_SHIFT                          = 27
	TMC2240_S2GB_MASK                           = 0x10000000
	TMC2240_S2GB_SHIFT                          = 28
	TMC2240_OLA_MASK                            = 0x20000000
	TMC2240_OLA_SHIFT                           = 29
	TMC2240_OLB_MASK                            = 0x40000000
	TMC2240_OLB_SHIFT                           = 30
	TMC2240_STST_MASK                           = 0x80000000
	TMC2240_STST_SHIFT                          = 31
	TMC2240_PWM_OFS_MASK                        = 0x000000FF
	TMC2240_PWM_OFS_SHIFT                       = 0
	TMC2240_PWM_GRAD_MASK                       = 0x0000FF00
	TMC2240_PWM_GRAD_SHIFT                      = 8
	TMC2240_PWM_FREQ_MASK                       = 0x00030000
	TMC2240_PWM_FREQ_SHIFT                      = 16
	TMC2240_PWM_AUTOSCALE_MASK                  = 0x00040000
	TMC2240_PWM_AUTOSCALE_SHIFT                 = 18
	TMC2240_PWM_AUTOGRAD_MASK                   = 0x00080000
	TMC2240_PWM_AUTOGRAD_SHIFT                  = 19
	TMC2240_FREEWHEEL_MASK                      = 0x00300000
	TMC2240_FREEWHEEL_SHIFT                     = 20
	TMC2240_PWM_MEAS_SD_ENABLE_MASK             = 0x00400000
	TMC2240_PWM_MEAS_SD_ENABLE_SHIFT            = 22
	TMC2240_PWM_DIS_REG_STST_MASK               = 0x00800000
	TMC2240_PWM_DIS_REG_STST_SHIFT              = 23
	TMC2240_PWM_REG_MASK                        = 0x0F000000
	TMC2240_PWM_REG_SHIFT                       = 24
	TMC2240_PWM_LIM_MASK                        = 0xF0000000
	TMC2240_PWM_LIM_SHIFT                       = 28
	TMC2240_PWM_SCALE_SUM_MASK                  = 0x000003FF
	TMC2240_PWM_SCALE_SUM_SHIFT                 = 0
	TMC2240_PWM_SCALE_AUTO_MASK                 = 0x01FF0000
	TMC2240_PWM_SCALE_AUTO_SHIFT                = 16
	TMC2240_PWM_OFS_AUTO_MASK                   = 0x000000FF
	TMC2240_PWM_OFS_AUTO_SHIFT                  = 0
	TMC2240_PWM_GRAD_AUTO_MASK                  = 0x00FF0000
	TMC2240_PWM_GRAD_AUTO_SHIFT                 = 16
	TMC2240_SG4_THRS_MASK                       = 0x000000FF
	TMC2240_SG4_THRS_SHIFT                      = 0
	TMC2240_SG4_FILT_EN_MASK                    = 0x00000100
	TMC2240_SG4_FILT_EN_SHIFT                   = 8
	TMC2240_SG_ANGLE_OFFSET_MASK                = 0x00000200
	TMC2240_SG_ANGLE_OFFSET_SHIFT               = 9
	TMC2240_SG4_RESULT_MASK                     = 0x000003FF
	TMC2240_SG4_RESULT_SHIFT                    = 0
	TMC2240_SG4_IND_0_MASK                      = 0x000000FF
	TMC2240_SG4_IND_0_SHIFT                     = 0
	TMC2240_SG4_IND_1_MASK                      = 0x0000FF00
	TMC2240_SG4_IND_1_SHIFT                     = 8
	TMC2240_SG4_IND_2_MASK                      = 0x00FF0000
	TMC2240_SG4_IND_2_SHIFT                     = 16
	TMC2240_SG4_IND_3_MASK                      = 0xFF000000
	TMC2240_SG4_IND_3_SHIFT                     = 24
)

var TMC2240_MSLUT = []uint32{
	TMC2240_MSLUT0,
	TMC2240_MSLUT1,
	TMC2240_MSLUT2,
	TMC2240_MSLUT3,
	TMC2240_MSLUT4,
	TMC2240_MSLUT5,
	TMC2240_MSLUT6,
	TMC2240_MSLUT7,
}

// Register Fields
var (
	TMC2240_FAST_STANDSTILL_FIELD = RegisterField{
		Mask:     TMC2240_FAST_STANDSTILL_MASK,
		Shift:    TMC2240_FAST_STANDSTILL_SHIFT,
		Register: TMC2240_GCONF,
		IsSigned: false,
	}
	TMC2240_EN_PWM_MODE_FIELD = RegisterField{
		Mask:     TMC2240_EN_PWM_MODE_MASK,
		Shift:    TMC2240_EN_PWM_MODE_SHIFT,
		Register: TMC2240_GCONF,
		IsSigned: false,
	}
	TMC2240_MULTISTEP_FILT_FIELD = RegisterField{
		Mask:     TMC2240_MULTISTEP_FILT_MASK,
		Shift:    TMC2240_MULTISTEP_FILT_SHIFT,
		Register: TMC2240_GCONF,
		IsSigned: false,
	}
	TMC2240_SHAFT_FIELD = RegisterField{
		Mask:     TMC2240_SHAFT_MASK,
		Shift:    TMC2240_SHAFT_SHIFT,
		Register: TMC2240_GCONF,
		IsSigned: false,
	}
	TMC2240_DIAG0_ERROR_FIELD = RegisterField{
		Mask:     TMC2240_DIAG0_ERROR_MASK,
		Shift:    TMC2240_DIAG0_ERROR_SHIFT,
		Register: TMC2240_GCONF,
		IsSigned: false,
	}
	TMC2240_DIAG0_OTPW_FIELD = RegisterField{
		Mask:     TMC2240_DIAG0_OTPW_MASK,
		Shift:    TMC2240_DIAG0_OTPW_SHIFT,
		Register: TMC2240_GCONF,
		IsSigned: false,
	}
	TMC2240_DIAG0_STALL_FIELD = RegisterField{
		Mask:     TMC2240_DIAG0_STALL_MASK,
		Shift:    TMC2240_DIAG0_STALL_SHIFT,
		Register: TMC2240_GCONF,
		IsSigned: false,
	}
	TMC2240_DIAG1_STALL_FIELD = RegisterField{
		Mask:     TMC2240_DIAG1_STALL_MASK,
		Shift:    TMC2240_DIAG1_STALL_SHIFT,
		Register: TMC2240_GCONF,
		IsSigned: false,
	}
	TMC2240_DIAG1_INDEX_FIELD = RegisterField{
		Mask:     TMC2240_DIAG1_INDEX_MASK,
		Shift:    TMC2240_DIAG1_INDEX_SHIFT,
		Register: TMC2240_GCONF,
		IsSigned: false,
	}
	TMC2240_DIAG1_ONSTATE_FIELD = RegisterField{
		Mask:     TMC2240_DIAG1_ONSTATE_MASK,
		Shift:    TMC2240_DIAG1_ONSTATE_SHIFT,
		Register: TMC2240_GCONF,
		IsSigned: false,
	}
	TMC2240_DIAG0_PUSHPULL_FIELD = RegisterField{
		Mask:     TMC2240_DIAG0_PUSHPULL_MASK,
		Shift:    TMC2240_DIAG0_PUSHPULL_SHIFT,
		Register: TMC2240_GCONF,
		IsSigned: false,
	}
	TMC2240_DIAG1_PUSHPULL_FIELD = RegisterField{
		Mask:     TMC2240_DIAG1_PUSHPULL_MASK,
		Shift:    TMC2240_DIAG1_PUSHPULL_SHIFT,
		Register: TMC2240_GCONF,
		IsSigned: false,
	}
	TMC2240_SMALL_HYSTERESIS_FIELD = RegisterField{
		Mask:     TMC2240_SMALL_HYSTERESIS_MASK,
		Shift:    TMC2240_SMALL_HYSTERESIS_SHIFT,
		Register: TMC2240_GCONF,
		IsSigned: false,
	}
	TMC2240_STOP_ENABLE_FIELD = RegisterField{
		Mask:     TMC2240_STOP_ENABLE_MASK,
		Shift:    TMC2240_STOP_ENABLE_SHIFT,
		Register: TMC2240_GCONF,
		IsSigned: false,
	}
	TMC2240_DIRECT_MODE_FIELD = RegisterField{
		Mask:     TMC2240_DIRECT_MODE_MASK,
		Shift:    TMC2240_DIRECT_MODE_SHIFT,
		Register: TMC2240_GCONF,
		IsSigned: false,
	}
	TMC2240_RESET_FIELD = RegisterField{
		Mask:     TMC2240_RESET_MASK,
		Shift:    TMC2240_RESET_SHIFT,
		Register: TMC2240_GSTAT,
		IsSigned: false,
	}
	TMC2240_DRV_ERR_FIELD = RegisterField{
		Mask:     TMC2240_DRV_ERR_MASK,
		Shift:    TMC2240_DRV_ERR_SHIFT,
		Register: TMC2240_GSTAT,
		IsSigned: false,
	}
	TMC2240_UV_CP_FIELD = RegisterField{
		Mask:     TMC2240_UV_CP_MASK,
		Shift:    TMC2240_UV_CP_SHIFT,
		Register: TMC2240_GSTAT,
		IsSigned: false,
	}
	TMC2240_REGISTER_RESET_FIELD = RegisterField{
		Mask:     TMC2240_REGISTER_RESET_MASK,
		Shift:    TMC2240_REGISTER_RESET_SHIFT,
		Register: TMC2240_GSTAT,
		IsSigned: false,
	}
	TMC2240_VM_UVLO_FIELD = RegisterField{
		Mask:     TMC2240_VM_UVLO_MASK,
		Shift:    TMC2240_VM_UVLO_SHIFT,
		Register: TMC2240_GSTAT,
		IsSigned: false,
	}
	TMC2240_IFCNT_FIELD = RegisterField{
		Mask:     TMC2240_IFCNT_MASK,
		Shift:    TMC2240_IFCNT_SHIFT,
		Register: TMC2240_IFCNT,
		IsSigned: false,
	}
	TMC2240_SLAVEADDR_FIELD = RegisterField{
		Mask:     TMC2240_SLAVEADDR_MASK,
		Shift:    TMC2240_SLAVEADDR_SHIFT,
		Register: TMC2240_SLAVECONF,
		IsSigned: false,
	}
	TMC2240_SENDDELAY_FIELD = RegisterField{
		Mask:     TMC2240_SENDDELAY_MASK,
		Shift:    TMC2240_SENDDELAY_SHIFT,
		Register: TMC2240_SLAVECONF,
		IsSigned: false,
	}
	TMC2240_STEP_FIELD = RegisterField{
		Mask:     TMC2240_STEP_MASK,
		Shift:    TMC2240_STEP_SHIFT,
		Register: TMC2240_IOIN,
		IsSigned: false,
	}
	TMC2240_DIR_FIELD = RegisterField{
		Mask:     TMC2240_DIR_MASK,
		Shift:    TMC2240_DIR_SHIFT,
		Register: TMC2240_IOIN,
		IsSigned: false,
	}
	TMC2240_ENCB_FIELD = RegisterField{
		Mask:     TMC2240_ENCB_MASK,
		Shift:    TMC2240_ENCB_SHIFT,
		Register: TMC2240_IOIN,
		IsSigned: false,
	}
	TMC2240_ENCA_FIELD = RegisterField{
		Mask:     TMC2240_ENCA_MASK,
		Shift:    TMC2240_ENCA_SHIFT,
		Register: TMC2240_IOIN,
		IsSigned: false,
	}
	TMC2240_DRV_ENN_FIELD = RegisterField{
		Mask:     TMC2240_DRV_ENN_MASK,
		Shift:    TMC2240_DRV_ENN_SHIFT,
		Register: TMC2240_IOIN,
		IsSigned: false,
	}
	TMC2240_ENCN_FIELD = RegisterField{
		Mask:     TMC2240_ENCN_MASK,
		Shift:    TMC2240_ENCN_SHIFT,
		Register: TMC2240_IOIN,
		IsSigned: false,
	}
	TMC2240_UART_EN_FIELD = RegisterField{
		Mask:     TMC2240_UART_EN_MASK,
		Shift:    TMC2240_UART_EN_SHIFT,
		Register: TMC2240_IOIN,
		IsSigned: false,
	}
	TMC2240_RESERVED_FIELD = RegisterField{
		Mask:     TMC2240_RESERVED_MASK,
		Shift:    TMC2240_RESERVED_SHIFT,
		Register: TMC2240_IOIN,
		IsSigned: false,
	}
	TMC2240_COMP_A_FIELD = RegisterField{
		Mask:     TMC2240_COMP_A_MASK,
		Shift:    TMC2240_COMP_A_SHIFT,
		Register: TMC2240_IOIN,
		IsSigned: false,
	}
	TMC2240_COMP_B_FIELD = RegisterField{
		Mask:     TMC2240_COMP_B_MASK,
		Shift:    TMC2240_COMP_B_SHIFT,
		Register: TMC2240_IOIN,
		IsSigned: false,
	}
	TMC2240_COMP_A1_A2_FIELD = RegisterField{
		Mask:     TMC2240_COMP_A1_A2_MASK,
		Shift:    TMC2240_COMP_A1_A2_SHIFT,
		Register: TMC2240_IOIN,
		IsSigned: false,
	}
	TMC2240_COMP_B1_B2_FIELD = RegisterField{
		Mask:     TMC2240_COMP_B1_B2_MASK,
		Shift:    TMC2240_COMP_B1_B2_SHIFT,
		Register: TMC2240_IOIN,
		IsSigned: false,
	}
	TMC2240_OUTPUT_FIELD = RegisterField{
		Mask:     TMC2240_OUTPUT_MASK,
		Shift:    TMC2240_OUTPUT_SHIFT,
		Register: TMC2240_IOIN,
		IsSigned: false,
	}
	TMC2240_EXT_RES_DET_FIELD = RegisterField{
		Mask:     TMC2240_EXT_RES_DET_MASK,
		Shift:    TMC2240_EXT_RES_DET_SHIFT,
		Register: TMC2240_IOIN,
		IsSigned: false,
	}
	TMC2240_EXT_CLK_FIELD = RegisterField{
		Mask:     TMC2240_EXT_CLK_MASK,
		Shift:    TMC2240_EXT_CLK_SHIFT,
		Register: TMC2240_IOIN,
		IsSigned: false,
	}
	TMC2240_ADC_ERR_FIELD = RegisterField{
		Mask:     TMC2240_ADC_ERR_MASK,
		Shift:    TMC2240_ADC_ERR_SHIFT,
		Register: TMC2240_IOIN,
		IsSigned: false,
	}
	TMC2240_SILICON_RV_FIELD = RegisterField{
		Mask:     TMC2240_SILICON_RV_MASK,
		Shift:    TMC2240_SILICON_RV_SHIFT,
		Register: TMC2240_IOIN,
		IsSigned: false,
	}
	TMC2240_VERSION_FIELD = RegisterField{
		Mask:     TMC2240_VERSION_MASK,
		Shift:    TMC2240_VERSION_SHIFT,
		Register: TMC2240_IOIN,
		IsSigned: false,
	}
	TMC2240_CURRENT_RANGE_FIELD = RegisterField{
		Mask:     TMC2240_CURRENT_RANGE_MASK,
		Shift:    TMC2240_CURRENT_RANGE_SHIFT,
		Register: TMC2240_DRV_CONF,
		IsSigned: false,
	}
	TMC2240_SLOPE_CONTROL_FIELD = RegisterField{
		Mask:     TMC2240_SLOPE_CONTROL_MASK,
		Shift:    TMC2240_SLOPE_CONTROL_SHIFT,
		Register: TMC2240_DRV_CONF,
		IsSigned: false,
	}
	TMC2240_GLOBALSCALER_FIELD = RegisterField{
		Mask:     TMC2240_GLOBALSCALER_MASK,
		Shift:    TMC2240_GLOBALSCALER_SHIFT,
		Register: TMC2240_GLOBAL_SCALER,
		IsSigned: false,
	}
	TMC2240_IHOLD_FIELD = RegisterField{
		Mask:     TMC2240_IHOLD_MASK,
		Shift:    TMC2240_IHOLD_SHIFT,
		Register: TMC2240_IHOLD_IRUN,
		IsSigned: false,
	}
	TMC2240_IRUN_FIELD = RegisterField{
		Mask:     TMC2240_IRUN_MASK,
		Shift:    TMC2240_IRUN_SHIFT,
		Register: TMC2240_IHOLD_IRUN,
		IsSigned: false,
	}
	TMC2240_IHOLDDELAY_FIELD = RegisterField{
		Mask:     TMC2240_IHOLDDELAY_MASK,
		Shift:    TMC2240_IHOLDDELAY_SHIFT,
		Register: TMC2240_IHOLD_IRUN,
		IsSigned: false,
	}
	TMC2240_IRUNDELAY_FIELD = RegisterField{
		Mask:     TMC2240_IRUNDELAY_MASK,
		Shift:    TMC2240_IRUNDELAY_SHIFT,
		Register: TMC2240_IHOLD_IRUN,
		IsSigned: false,
	}
	TMC2240_TPOWERDOWN_FIELD = RegisterField{
		Mask:     TMC2240_TPOWERDOWN_MASK,
		Shift:    TMC2240_TPOWERDOWN_SHIFT,
		Register: TMC2240_TPOWERDOWN,
		IsSigned: false,
	}
	TMC2240_TSTEP_FIELD = RegisterField{
		Mask:     TMC2240_TSTEP_MASK,
		Shift:    TMC2240_TSTEP_SHIFT,
		Register: TMC2240_TSTEP,
		IsSigned: false,
	}
	TMC2240_TPWMTHRS_FIELD = RegisterField{
		Mask:     TMC2240_TPWMTHRS_MASK,
		Shift:    TMC2240_TPWMTHRS_SHIFT,
		Register: TMC2240_TPWMTHRS,
		IsSigned: false,
	}
	TMC2240_TCOOLTHRS_FIELD = RegisterField{
		Mask:     TMC2240_TCOOLTHRS_MASK,
		Shift:    TMC2240_TCOOLTHRS_SHIFT,
		Register: TMC2240_TCOOLTHRS,
		IsSigned: false,
	}
	TMC2240_THIGH_FIELD = RegisterField{
		Mask:     TMC2240_THIGH_MASK,
		Shift:    TMC2240_THIGH_SHIFT,
		Register: TMC2240_THIGH,
		IsSigned: false,
	}
	TMC2240_DIRECT_COIL_A_FIELD = RegisterField{
		Mask:     TMC2240_DIRECT_COIL_A_MASK,
		Shift:    TMC2240_DIRECT_COIL_A_SHIFT,
		Register: TMC2240_DIRECT_MODE,
		IsSigned: true,
	}
	TMC2240_DIRECT_COIL_B_FIELD = RegisterField{
		Mask:     TMC2240_DIRECT_COIL_B_MASK,
		Shift:    TMC2240_DIRECT_COIL_B_SHIFT,
		Register: TMC2240_DIRECT_MODE,
		IsSigned: true,
	}
	TMC2240_POL_A_FIELD = RegisterField{
		Mask:     TMC2240_POL_A_MASK,
		Shift:    TMC2240_POL_A_SHIFT,
		Register: TMC2240_ENCMODE,
		IsSigned: false,
	}
	TMC2240_POL_B_FIELD = RegisterField{
		Mask:     TMC2240_POL_B_MASK,
		Shift:    TMC2240_POL_B_SHIFT,
		Register: TMC2240_ENCMODE,
		IsSigned: false,
	}
	TMC2240_POL_N_FIELD = RegisterField{
		Mask:     TMC2240_POL_N_MASK,
		Shift:    TMC2240_POL_N_SHIFT,
		Register: TMC2240_ENCMODE,
		IsSigned: false,
	}
	TMC2240_IGNORE_AB_FIELD = RegisterField{
		Mask:     TMC2240_IGNORE_AB_MASK,
		Shift:    TMC2240_IGNORE_AB_SHIFT,
		Register: TMC2240_ENCMODE,
		IsSigned: false,
	}
	TMC2240_CLR_CONT_FIELD = RegisterField{
		Mask:     TMC2240_CLR_CONT_MASK,
		Shift:    TMC2240_CLR_CONT_SHIFT,
		Register: TMC2240_ENCMODE,
		IsSigned: false,
	}
	TMC2240_CLR_ONCE_FIELD = RegisterField{
		Mask:     TMC2240_CLR_ONCE_MASK,
		Shift:    TMC2240_CLR_ONCE_SHIFT,
		Register: TMC2240_ENCMODE,
		IsSigned: false,
	}
	TMC2240_POS_NEG_EDGE_FIELD = RegisterField{
		Mask:     TMC2240_POS_NEG_EDGE_MASK,
		Shift:    TMC2240_POS_NEG_EDGE_SHIFT,
		Register: TMC2240_ENCMODE,
		IsSigned: false,
	}
	TMC2240_CLR_ENC_X_FIELD = RegisterField{
		Mask:     TMC2240_CLR_ENC_X_MASK,
		Shift:    TMC2240_CLR_ENC_X_SHIFT,
		Register: TMC2240_ENCMODE,
		IsSigned: false,
	}
	TMC2240_LATCH_X_ACT_FIELD = RegisterField{
		Mask:     TMC2240_LATCH_X_ACT_MASK,
		Shift:    TMC2240_LATCH_X_ACT_SHIFT,
		Register: TMC2240_ENCMODE,
		IsSigned: false,
	}
	TMC2240_ENC_SEL_DECIMAL_FIELD = RegisterField{
		Mask:     TMC2240_ENC_SEL_DECIMAL_MASK,
		Shift:    TMC2240_ENC_SEL_DECIMAL_SHIFT,
		Register: TMC2240_ENCMODE,
		IsSigned: false,
	}
	TMC2240_X_ENC_FIELD = RegisterField{
		Mask:     TMC2240_X_ENC_MASK,
		Shift:    TMC2240_X_ENC_SHIFT,
		Register: TMC2240_XENC,
		IsSigned: true,
	}
	TMC2240_ENC_CONST_FIELD = RegisterField{
		Mask:     TMC2240_ENC_CONST_MASK,
		Shift:    TMC2240_ENC_CONST_SHIFT,
		Register: TMC2240_ENC_CONST,
		IsSigned: true,
	}
	TMC2240_N_EVENT_FIELD = RegisterField{
		Mask:     TMC2240_N_EVENT_MASK,
		Shift:    TMC2240_N_EVENT_SHIFT,
		Register: TMC2240_ENC_STATUS,
		IsSigned: false,
	}
	TMC2240_DEVIATION_WARN_FIELD = RegisterField{
		Mask:     TMC2240_DEVIATION_WARN_MASK,
		Shift:    TMC2240_DEVIATION_WARN_SHIFT,
		Register: TMC2240_ENC_STATUS,
		IsSigned: false,
	}
	TMC2240_ENC_LATCH_FIELD = RegisterField{
		Mask:     TMC2240_ENC_LATCH_MASK,
		Shift:    TMC2240_ENC_LATCH_SHIFT,
		Register: TMC2240_ENC_LATCH,
		IsSigned: false,
	}
	TMC2240_ADC_VSUPPLY_FIELD = RegisterField{
		Mask:     TMC2240_ADC_VSUPPLY_MASK,
		Shift:    TMC2240_ADC_VSUPPLY_SHIFT,
		Register: TMC2240_ADC_VSUPPLY_AIN,
		IsSigned: true,
	}
	TMC2240_ADC_AIN_FIELD = RegisterField{
		Mask:     TMC2240_ADC_AIN_MASK,
		Shift:    TMC2240_ADC_AIN_SHIFT,
		Register: TMC2240_ADC_VSUPPLY_AIN,
		IsSigned: true,
	}
	TMC2240_ADC_TEMP_FIELD = RegisterField{
		Mask:     TMC2240_ADC_TEMP_MASK,
		Shift:    TMC2240_ADC_TEMP_SHIFT,
		Register: TMC2240_ADC_TEMP,
		IsSigned: true,
	}
	TMC2240_OVERVOLTAGE_VTH_FIELD = RegisterField{
		Mask:     TMC2240_OVERVOLTAGE_VTH_MASK,
		Shift:    TMC2240_OVERVOLTAGE_VTH_SHIFT,
		Register: TMC2240_OTW_OV_VTH,
		IsSigned: false,
	}
	TMC2240_OVERTEMPPREWARNING_VTH_FIELD = RegisterField{
		Mask:     TMC2240_OVERTEMPPREWARNING_VTH_MASK,
		Shift:    TMC2240_OVERTEMPPREWARNING_VTH_SHIFT,
		Register: TMC2240_OTW_OV_VTH,
		IsSigned: false,
	}
	TMC2240_W0_FIELD = RegisterField{
		Mask:     TMC2240_W0_MASK,
		Shift:    TMC2240_W0_SHIFT,
		Register: TMC2240_MSLUTSEL,
		IsSigned: false,
	}
	TMC2240_W1_FIELD = RegisterField{
		Mask:     TMC2240_W1_MASK,
		Shift:    TMC2240_W1_SHIFT,
		Register: TMC2240_MSLUTSEL,
		IsSigned: false,
	}
	TMC2240_W2_FIELD = RegisterField{
		Mask:     TMC2240_W2_MASK,
		Shift:    TMC2240_W2_SHIFT,
		Register: TMC2240_MSLUTSEL,
		IsSigned: false,
	}
	TMC2240_W3_FIELD = RegisterField{
		Mask:     TMC2240_W3_MASK,
		Shift:    TMC2240_W3_SHIFT,
		Register: TMC2240_MSLUTSEL,
		IsSigned: false,
	}
	TMC2240_X1_FIELD = RegisterField{
		Mask:     TMC2240_X1_MASK,
		Shift:    TMC2240_X1_SHIFT,
		Register: TMC2240_MSLUTSEL,
		IsSigned: false,
	}
	TMC2240_X2_FIELD = RegisterField{
		Mask:     TMC2240_X2_MASK,
		Shift:    TMC2240_X2_SHIFT,
		Register: TMC2240_MSLUTSEL,
		IsSigned: false,
	}
	TMC2240_X3_FIELD = RegisterField{
		Mask:     TMC2240_X3_MASK,
		Shift:    TMC2240_X3_SHIFT,
		Register: TMC2240_MSLUTSEL,
		IsSigned: false,
	}
	TMC2240_START_SIN_FIELD = RegisterField{
		Mask:     TMC2240_START_SIN_MASK,
		Shift:    TMC2240_START_SIN_SHIFT,
		Register: TMC2240_MSLUTSTART,
		IsSigned: false,
	}
	TMC2240_START_SIN90_FIELD = RegisterField{
		Mask:     TMC2240_START_SIN90_MASK,
		Shift:    TMC2240_START_SIN90_SHIFT,
		Register: TMC2240_MSLUTSTART,
		IsSigned: false,
	}
	TMC2240_OFFSET_SIN90_FIELD = RegisterField{
		Mask:     TMC2240_OFFSET_SIN90_MASK,
		Shift:    TMC2240_OFFSET_SIN90_SHIFT,
		Register: TMC2240_MSLUTSTART,
		IsSigned: false,
	}
	TMC2240_MSCNT_FIELD = RegisterField{
		Mask:     TMC2240_MSCNT_MASK,
		Shift:    TMC2240_MSCNT_SHIFT,
		Register: TMC2240_MSCNT,
		IsSigned: false,
	}
	TMC2240_CUR_B_FIELD = RegisterField{
		Mask:     TMC2240_CUR_B_MASK,
		Shift:    TMC2240_CUR_B_SHIFT,
		Register: TMC2240_MSCURACT,
		IsSigned: true,
	}
	TMC2240_CUR_A_FIELD = RegisterField{
		Mask:     TMC2240_CUR_A_MASK,
		Shift:    TMC2240_CUR_A_SHIFT,
		Register: TMC2240_MSCURACT,
		IsSigned: true,
	}
	TMC2240_TOFF_FIELD = RegisterField{
		Mask:     TMC2240_TOFF_MASK,
		Shift:    TMC2240_TOFF_SHIFT,
		Register: TMC2240_CHOPCONF,
		IsSigned: false,
	}
	TMC2240_HSTRT_TFD210_FIELD = RegisterField{
		Mask:     TMC2240_HSTRT_TFD210_MASK,
		Shift:    TMC2240_HSTRT_TFD210_SHIFT,
		Register: TMC2240_CHOPCONF,
		IsSigned: false,
	}
	TMC2240_HEND_OFFSET_FIELD = RegisterField{
		Mask:     TMC2240_HEND_OFFSET_MASK,
		Shift:    TMC2240_HEND_OFFSET_SHIFT,
		Register: TMC2240_CHOPCONF,
		IsSigned: false,
	}
	TMC2240_FD3_FIELD = RegisterField{
		Mask:     TMC2240_FD3_MASK,
		Shift:    TMC2240_FD3_SHIFT,
		Register: TMC2240_CHOPCONF,
		IsSigned: false,
	}
	TMC2240_DISFDCC_FIELD = RegisterField{
		Mask:     TMC2240_DISFDCC_MASK,
		Shift:    TMC2240_DISFDCC_SHIFT,
		Register: TMC2240_CHOPCONF,
		IsSigned: false,
	}
	TMC2240_CHM_FIELD = RegisterField{
		Mask:     TMC2240_CHM_MASK,
		Shift:    TMC2240_CHM_SHIFT,
		Register: TMC2240_CHOPCONF,
		IsSigned: false,
	}
	TMC2240_TBL_FIELD = RegisterField{
		Mask:     TMC2240_TBL_MASK,
		Shift:    TMC2240_TBL_SHIFT,
		Register: TMC2240_CHOPCONF,
		IsSigned: false,
	}
	TMC2240_VHIGHFS_FIELD = RegisterField{
		Mask:     TMC2240_VHIGHFS_MASK,
		Shift:    TMC2240_VHIGHFS_SHIFT,
		Register: TMC2240_CHOPCONF,
		IsSigned: false,
	}
	TMC2240_VHIGHCHM_FIELD = RegisterField{
		Mask:     TMC2240_VHIGHCHM_MASK,
		Shift:    TMC2240_VHIGHCHM_SHIFT,
		Register: TMC2240_CHOPCONF,
		IsSigned: false,
	}
	TMC2240_TPFD_FIELD = RegisterField{
		Mask:     TMC2240_TPFD_MASK,
		Shift:    TMC2240_TPFD_SHIFT,
		Register: TMC2240_CHOPCONF,
		IsSigned: false,
	}
	TMC2240_MRES_FIELD = RegisterField{
		Mask:     TMC2240_MRES_MASK,
		Shift:    TMC2240_MRES_SHIFT,
		Register: TMC2240_CHOPCONF,
		IsSigned: false,
	}
	TMC2240_INTPOL_FIELD = RegisterField{
		Mask:     TMC2240_INTPOL_MASK,
		Shift:    TMC2240_INTPOL_SHIFT,
		Register: TMC2240_CHOPCONF,
		IsSigned: false,
	}
	TMC2240_DEDGE_FIELD = RegisterField{
		Mask:     TMC2240_DEDGE_MASK,
		Shift:    TMC2240_DEDGE_SHIFT,
		Register: TMC2240_CHOPCONF,
		IsSigned: false,
	}
	TMC2240_DISS2G_FIELD = RegisterField{
		Mask:     TMC2240_DISS2G_MASK,
		Shift:    TMC2240_DISS2G_SHIFT,
		Register: TMC2240_CHOPCONF,
		IsSigned: false,
	}
	TMC2240_DISS2VS_FIELD = RegisterField{
		Mask:     TMC2240_DISS2VS_MASK,
		Shift:    TMC2240_DISS2VS_SHIFT,
		Register: TMC2240_CHOPCONF,
		IsSigned: false,
	}
	TMC2240_SEMIN_FIELD = RegisterField{
		Mask:     TMC2240_SEMIN_MASK,
		Shift:    TMC2240_SEMIN_SHIFT,
		Register: TMC2240_COOLCONF,
		IsSigned: false,
	}
	TMC2240_SEUP_FIELD = RegisterField{
		Mask:     TMC2240_SEUP_MASK,
		Shift:    TMC2240_SEUP_SHIFT,
		Register: TMC2240_COOLCONF,
		IsSigned: false,
	}
	TMC2240_SEMAX_FIELD = RegisterField{
		Mask:     TMC2240_SEMAX_MASK,
		Shift:    TMC2240_SEMAX_SHIFT,
		Register: TMC2240_COOLCONF,
		IsSigned: false,
	}
	TMC2240_SEDN_FIELD = RegisterField{
		Mask:     TMC2240_SEDN_MASK,
		Shift:    TMC2240_SEDN_SHIFT,
		Register: TMC2240_COOLCONF,
		IsSigned: false,
	}
	TMC2240_SEIMIN_FIELD = RegisterField{
		Mask:     TMC2240_SEIMIN_MASK,
		Shift:    TMC2240_SEIMIN_SHIFT,
		Register: TMC2240_COOLCONF,
		IsSigned: false,
	}
	TMC2240_SGT_FIELD = RegisterField{
		Mask:     TMC2240_SGT_MASK,
		Shift:    TMC2240_SGT_SHIFT,
		Register: TMC2240_COOLCONF,
		IsSigned: false,
	}
	TMC2240_SFILT_FIELD = RegisterField{
		Mask:     TMC2240_SFILT_MASK,
		Shift:    TMC2240_SFILT_SHIFT,
		Register: TMC2240_COOLCONF,
		IsSigned: false,
	}
	TMC2240_SG_RESULT_FIELD = RegisterField{
		Mask:     TMC2240_SG_RESULT_MASK,
		Shift:    TMC2240_SG_RESULT_SHIFT,
		Register: TMC2240_DRVSTATUS,
		IsSigned: false,
	}
	TMC2240_S2VSA_FIELD = RegisterField{
		Mask:     TMC2240_S2VSA_MASK,
		Shift:    TMC2240_S2VSA_SHIFT,
		Register: TMC2240_DRVSTATUS,
		IsSigned: false,
	}
	TMC2240_S2VSB_FIELD = RegisterField{
		Mask:     TMC2240_S2VSB_MASK,
		Shift:    TMC2240_S2VSB_SHIFT,
		Register: TMC2240_DRVSTATUS,
		IsSigned: false,
	}
	TMC2240_STEALTH_FIELD = RegisterField{
		Mask:     TMC2240_STEALTH_MASK,
		Shift:    TMC2240_STEALTH_SHIFT,
		Register: TMC2240_DRVSTATUS,
		IsSigned: false,
	}
	TMC2240_FSACTIVE_FIELD = RegisterField{
		Mask:     TMC2240_FSACTIVE_MASK,
		Shift:    TMC2240_FSACTIVE_SHIFT,
		Register: TMC2240_DRVSTATUS,
		IsSigned: false,
	}
	TMC2240_CS_ACTUAL_FIELD = RegisterField{
		Mask:     TMC2240_CS_ACTUAL_MASK,
		Shift:    TMC2240_CS_ACTUAL_SHIFT,
		Register: TMC2240_DRVSTATUS,
		IsSigned: false,
	}
	TMC2240_STALLGUARD_FIELD = RegisterField{
		Mask:     TMC2240_STALLGUARD_MASK,
		Shift:    TMC2240_STALLGUARD_SHIFT,
		Register: TMC2240_DRVSTATUS,
		IsSigned: false,
	}
	TMC2240_OT_FIELD = RegisterField{
		Mask:     TMC2240_OT_MASK,
		Shift:    TMC2240_OT_SHIFT,
		Register: TMC2240_DRVSTATUS,
		IsSigned: false,
	}
	TMC2240_OTPW_FIELD = RegisterField{
		Mask:     TMC2240_OTPW_MASK,
		Shift:    TMC2240_OTPW_SHIFT,
		Register: TMC2240_DRVSTATUS,
		IsSigned: false,
	}
	TMC2240_S2GA_FIELD = RegisterField{
		Mask:     TMC2240_S2GA_MASK,
		Shift:    TMC2240_S2GA_SHIFT,
		Register: TMC2240_DRVSTATUS,
		IsSigned: false,
	}
	TMC2240_S2GB_FIELD = RegisterField{
		Mask:     TMC2240_S2GB_MASK,
		Shift:    TMC2240_S2GB_SHIFT,
		Register: TMC2240_DRVSTATUS,
		IsSigned: false,
	}
	TMC2240_OLA_FIELD = RegisterField{
		Mask:     TMC2240_OLA_MASK,
		Shift:    TMC2240_OLA_SHIFT,
		Register: TMC2240_DRVSTATUS,
		IsSigned: false,
	}
	TMC2240_OLB_FIELD = RegisterField{
		Mask:     TMC2240_OLB_MASK,
		Shift:    TMC2240_OLB_SHIFT,
		Register: TMC2240_DRVSTATUS,
		IsSigned: false,
	}
	TMC2240_STST_FIELD = RegisterField{
		Mask:     TMC2240_STST_MASK,
		Shift:    TMC2240_STST_SHIFT,
		Register: TMC2240_DRVSTATUS,
		IsSigned: false,
	}
	TMC2240_PWM_OFS_FIELD = RegisterField{
		Mask:     TMC2240_PWM_OFS_MASK,
		Shift:    TMC2240_PWM_OFS_SHIFT,
		Register: TMC2240_PWMCONF,
		IsSigned: false,
	}
	TMC2240_PWM_GRAD_FIELD = RegisterField{
		Mask:     TMC2240_PWM_GRAD_MASK,
		Shift:    TMC2240_PWM_GRAD_SHIFT,
		Register: TMC2240_PWMCONF,
		IsSigned: false,
	}
	TMC2240_PWM_FREQ_FIELD = RegisterField{
		Mask:     TMC2240_PWM_FREQ_MASK,
		Shift:    TMC2240_PWM_FREQ_SHIFT,
		Register: TMC2240_PWMCONF,
		IsSigned: false,
	}
	TMC2240_PWM_AUTOSCALE_FIELD = RegisterField{
		Mask:     TMC2240_PWM_AUTOSCALE_MASK,
		Shift:    TMC2240_PWM_AUTOSCALE_SHIFT,
		Register: TMC2240_PWMCONF,
		IsSigned: false,
	}
	TMC2240_PWM_AUTOGRAD_FIELD = RegisterField{
		Mask:     TMC2240_PWM_AUTOGRAD_MASK,
		Shift:    TMC2240_PWM_AUTOGRAD_SHIFT,
		Register: TMC2240_PWMCONF,
		IsSigned: false,
	}
	TMC2240_FREEWHEEL_FIELD = RegisterField{
		Mask:     TMC2240_FREEWHEEL_MASK,
		Shift:    TMC2240_FREEWHEEL_SHIFT,
		Register: TMC2240_PWMCONF,
		IsSigned: false,
	}
	TMC2240_PWM_MEAS_SD_ENABLE_FIELD = RegisterField{
		Mask:     TMC2240_PWM_MEAS_SD_ENABLE_MASK,
		Shift:    TMC2240_PWM_MEAS_SD_ENABLE_SHIFT,
		Register: TMC2240_PWMCONF,
		IsSigned: false,
	}
	TMC2240_PWM_DIS_REG_STST_FIELD = RegisterField{
		Mask:     TMC2240_PWM_DIS_REG_STST_MASK,
		Shift:    TMC2240_PWM_DIS_REG_STST_SHIFT,
		Register: TMC2240_PWMCONF,
		IsSigned: false,
	}
	TMC2240_PWM_REG_FIELD = RegisterField{
		Mask:     TMC2240_PWM_REG_MASK,
		Shift:    TMC2240_PWM_REG_SHIFT,
		Register: TMC2240_PWMCONF,
		IsSigned: false,
	}
	TMC2240_PWM_LIM_FIELD = RegisterField{
		Mask:     TMC2240_PWM_LIM_MASK,
		Shift:    TMC2240_PWM_LIM_SHIFT,
		Register: TMC2240_PWMCONF,
		IsSigned: false,
	}
	TMC2240_PWM_SCALE_SUM_FIELD = RegisterField{
		Mask:     TMC2240_PWM_SCALE_SUM_MASK,
		Shift:    TMC2240_PWM_SCALE_SUM_SHIFT,
		Register: TMC2240_PWMSCALE,
		IsSigned: false,
	}
	TMC2240_PWM_SCALE_AUTO_FIELD = RegisterField{
		Mask:     TMC2240_PWM_SCALE_AUTO_MASK,
		Shift:    TMC2240_PWM_SCALE_AUTO_SHIFT,
		Register: TMC2240_PWMSCALE,
		IsSigned: false,
	}
	TMC2240_PWM_OFS_AUTO_FIELD = RegisterField{
		Mask:     TMC2240_PWM_OFS_AUTO_MASK,
		Shift:    TMC2240_PWM_OFS_AUTO_SHIFT,
		Register: TMC2240_PWM_AUTO,
		IsSigned: false,
	}
	TMC2240_PWM_GRAD_AUTO_FIELD = RegisterField{
		Mask:     TMC2240_PWM_GRAD_AUTO_MASK,
		Shift:    TMC2240_PWM_GRAD_AUTO_SHIFT,
		Register: TMC2240_PWM_AUTO,
		IsSigned: false,
	}
	TMC2240_SG4_THRS_FIELD = RegisterField{
		Mask:     TMC2240_SG4_THRS_MASK,
		Shift:    TMC2240_SG4_THRS_SHIFT,
		Register: TMC2240_SG4_THRS,
		IsSigned: false,
	}
	TMC2240_SG4_FILT_EN_FIELD = RegisterField{
		Mask:     TMC2240_SG4_FILT_EN_MASK,
		Shift:    TMC2240_SG4_FILT_EN_SHIFT,
		Register: TMC2240_SG4_THRS,
		IsSigned: false,
	}
	TMC2240_SG_ANGLE_OFFSET_FIELD = RegisterField{
		Mask:     TMC2240_SG_ANGLE_OFFSET_MASK,
		Shift:    TMC2240_SG_ANGLE_OFFSET_SHIFT,
		Register: TMC2240_SG4_THRS,
		IsSigned: false,
	}
	TMC2240_SG4_RESULT_FIELD = RegisterField{
		Mask:     TMC2240_SG4_RESULT_MASK,
		Shift:    TMC2240_SG4_RESULT_SHIFT,
		Register: TMC2240_SG4_RESULT,
		IsSigned: false,
	}
	TMC2240_SG4_IND_0_FIELD = RegisterField{
		Mask:     TMC2240_SG4_IND_0_MASK,
		Shift:    TMC2240_SG4_IND_0_SHIFT,
		Register: TMC2240_SG4_IND,
		IsSigned: false,
	}
	TMC2240_SG4_IND_1_FIELD = RegisterField{
		Mask:     TMC2240_SG4_IND_1_MASK,
		Shift:    TMC2240_SG4_IND_1_SHIFT,
		Register: TMC2240_SG4_IND,
		IsSigned: false,
	}
	TMC2240_SG4_IND_2_FIELD = RegisterField{
		Mask:     TMC2240_SG4_IND_2_MASK,
		Shift:    TMC2240_SG4_IND_2_SHIFT,
		Register: TMC2240_SG4_IND,
		IsSigned: false,
	}
	TMC2240_SG4_IND_3_FIELD = RegisterField{
		Mask:     TMC2240_SG4_IND_3_MASK,
		Shift:    TMC2240_SG4_IND_3_SHIFT,
		Register: TMC2240_SG4_IND,
		IsSigned: false,
	}
	SineWaveTableDefault = []uint32{
		0xAAAAB554, 0x4A9554AA, 0x24492929, 0x10104222,
		0xFBFFFFFF, 0xB5BB777D, 0x49295556, 0x404222}

	TMC2240_MSLUT_FIELDS = []RegisterField{
		{
			Mask:     0xFFFFFFFF,
			Shift:    0,
			Register: TMC2240_MSLUT0,
			IsSigned: false},
		{
			Mask:     0xFFFFFFFF,
			Shift:    0,
			Register: TMC2240_MSLUT1,
			IsSigned: false},
		{
			Mask:     0xFFFFFFFF,
			Shift:    0,
			Register: TMC2240_MSLUT2,
			IsSigned: false},
		{
			Mask:     0xFFFFFFFF,
			Shift:    0,
			Register: TMC2240_MSLUT3,
			IsSigned: false},
		{
			Mask:     0xFFFFFFFF,
			Shift:    0,
			Register: TMC2240_MSLUT4,
			IsSigned: false},
		{
			Mask:     0xFFFFFFFF,
			Shift:    0,
			Register: TMC2240_MSLUT5,
			IsSigned: false},
		{
			Mask:     0xFFFFFFFF,
			Shift:    0,
			Register: TMC2240_MSLUT6,
			IsSigned: false},
		{
			Mask:     0xFFFFFFFF,
			Shift:    0,
			Register: TMC2240_MSLUT7,
			IsSigned: false},
	}
)
