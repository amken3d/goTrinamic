package tmc4671

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
	TMC4671_CHIPINFO_DATA                                                       = 0x00
	TMC4671_CHIPINFO_ADDR                                                       = 0x01
	TMC4671_ADC_RAW_DATA                                                        = 0x02
	TMC4671_ADC_RAW_ADDR                                                        = 0x03
	TMC4671_DSADC_MCFG_B_MCFG_A                                                 = 0x04
	TMC4671_DSADC_MCLK_A                                                        = 0x05
	TMC4671_DSADC_MCLK_B                                                        = 0x06
	TMC4671_DSADC_MDEC_B_MDEC_A                                                 = 0x07
	TMC4671_ADC_I1_SCALE_OFFSET                                                 = 0x08
	TMC4671_ADC_I0_SCALE_OFFSET                                                 = 0x09
	TMC4671_ADC_I_SELECT                                                        = 0x0A
	TMC4671_ADC_I1_I0_EXT                                                       = 0x0B
	TMC4671_DS_ANALOG_INPUT_STAGE_CFG                                           = 0x0C
	TMC4671_AENC_0_SCALE_OFFSET                                                 = 0x0D
	TMC4671_AENC_1_SCALE_OFFSET                                                 = 0x0E
	TMC4671_AENC_2_SCALE_OFFSET                                                 = 0x0F
	TMC4671_AENC_SELECT                                                         = 0x11
	TMC4671_ADC_IWY_IUX                                                         = 0x12
	TMC4671_ADC_IV                                                              = 0x13
	TMC4671_AENC_WY_UX                                                          = 0x15
	TMC4671_AENC_VN                                                             = 0x16
	TMC4671_PWM_POLARITIES                                                      = 0x17
	TMC4671_PWM_MAXCNT                                                          = 0x18
	TMC4671_PWM_SV_CHOP                                                         = 0x1A
	TMC4671_MOTOR_TYPE_N_POLE_PAIRS                                             = 0x1B
	TMC4671_PHI_E_EXT                                                           = 0x1C
	TMC4671_PHI_M_EXT                                                           = 0x1D
	TMC4671_POSITION_EXT                                                        = 0x1E
	TMC4671_OPENLOOP_MODE                                                       = 0x1F
	TMC4671_OPENLOOP_ACCELERATION                                               = 0x20
	TMC4671_OPENLOOP_VELOCITY_TARGET                                            = 0x21
	TMC4671_OPENLOOP_VELOCITY_ACTUAL                                            = 0x22
	TMC4671_OPENLOOP_PHI                                                        = 0x23
	TMC4671_UQ_UD_EXT                                                           = 0x24
	TMC4671_ABN_DECODER_MODE                                                    = 0x25
	TMC4671_ABN_DECODER_PPR                                                     = 0x26
	TMC4671_ABN_DECODER_COUNT                                                   = 0x27
	TMC4671_ABN_DECODER_COUNT_N                                                 = 0x28
	TMC4671_ABN_DECODER_PHI_E_PHI_M_OFFSET                                      = 0x29
	TMC4671_ABN_DECODER_PHI_E_PHI_M                                             = 0x2A
	TMC4671_ABN_2_DECODER_MODE                                                  = 0x2C
	TMC4671_ABN_2_DECODER_PPR                                                   = 0x2D
	TMC4671_ABN_2_DECODER_COUNT                                                 = 0x2E
	TMC4671_ABN_2_DECODER_COUNT_N                                               = 0x2F
	TMC4671_ABN_2_DECODER_PHI_M_OFFSET                                          = 0x30
	TMC4671_ABN_2_DECODER_PHI_M                                                 = 0x31
	TMC4671_HALL_MODE                                                           = 0x33
	TMC4671_HALL_POSITION_060_000                                               = 0x34
	TMC4671_HALL_POSITION_180_120                                               = 0x35
	TMC4671_HALL_POSITION_300_240                                               = 0x36
	TMC4671_HALL_PHI_E_PHI_M_OFFSET                                             = 0x37
	TMC4671_HALL_DPHI_MAX                                                       = 0x38
	TMC4671_HALL_PHI_E_INTERPOLATED_PHI_E                                       = 0x39
	TMC4671_HALL_PHI_M                                                          = 0x3A
	TMC4671_AENC_DECODER_MODE                                                   = 0x3B
	TMC4671_AENC_DECODER_N_THRESHOLD                                            = 0x3C
	TMC4671_AENC_DECODER_PHI_A_RAW                                              = 0x3D
	TMC4671_AENC_DECODER_PHI_A_OFFSET                                           = 0x3E
	TMC4671_AENC_DECODER_PHI_A                                                  = 0x3F
	TMC4671_AENC_DECODER_PPR                                                    = 0x40
	TMC4671_AENC_DECODER_COUNT                                                  = 0x41
	TMC4671_AENC_DECODER_COUNT_N                                                = 0x42
	TMC4671_AENC_DECODER_PHI_E_PHI_M_OFFSET                                     = 0x45
	TMC4671_AENC_DECODER_PHI_E_PHI_M                                            = 0x46
	TMC4671_AENC_DECODER_POSITION                                               = 0x47
	TMC4671_CONFIG_DATA                                                         = 0x4D
	TMC4671_CONFIG_ADDR                                                         = 0x4E
	TMC4671_VELOCITY_SELECTION                                                  = 0x50
	TMC4671_POSITION_SELECTION                                                  = 0x51
	TMC4671_PHI_E_SELECTION                                                     = 0x52
	TMC4671_PHI_E                                                               = 0x53
	TMC4671_PID_FLUX_P_FLUX_I                                                   = 0x54
	TMC4671_PID_TORQUE_P_TORQUE_I                                               = 0x56
	TMC4671_PID_VELOCITY_P_VELOCITY_I                                           = 0x58
	TMC4671_PID_POSITION_P_POSITION_I                                           = 0x5A
	TMC4671_PID_TORQUE_FLUX_TARGET_DDT_LIMITS                                   = 0x5C
	TMC4671_PIDOUT_UQ_UD_LIMITS                                                 = 0x5D
	TMC4671_PID_TORQUE_FLUX_LIMITS                                              = 0x5E
	TMC4671_PID_VELOCITY_LIMIT                                                  = 0x60
	TMC4671_PID_POSITION_LIMIT_LOW                                              = 0x61
	TMC4671_PID_POSITION_LIMIT_HIGH                                             = 0x62
	TMC4671_MODE_RAMP_MODE_MOTION                                               = 0x63
	TMC4671_PID_TORQUE_FLUX_TARGET                                              = 0x64
	TMC4671_PID_TORQUE_FLUX_OFFSET                                              = 0x65
	TMC4671_PID_VELOCITY_TARGET                                                 = 0x66
	TMC4671_PID_VELOCITY_OFFSET                                                 = 0x67
	TMC4671_PID_POSITION_TARGET                                                 = 0x68
	TMC4671_PID_TORQUE_FLUX_ACTUAL                                              = 0x69
	TMC4671_PID_VELOCITY_ACTUAL                                                 = 0x6A
	TMC4671_PID_POSITION_ACTUAL                                                 = 0x6B
	TMC4671_PID_ERROR_DATA                                                      = 0x6C
	TMC4671_PID_ERROR_ADDR                                                      = 0x6D
	TMC4671_INTERIM_DATA                                                        = 0x6E
	TMC4671_INTERIM_ADDR                                                        = 0x6F
	TMC4671_WATCHDOG_CFG                                                        = 0x74
	TMC4671_ADC_VM_LIMITS                                                       = 0x75
	TMC4671_INPUTS_RAW                                                          = 0x76
	TMC4671_OUTPUTS_RAW                                                         = 0x77
	TMC4671_STEP_WIDTH                                                          = 0x78
	TMC4671_UART_BPS                                                            = 0x79
	TMC4671_UART_ADDRS                                                          = 0x7A
	TMC4671_GPIO_DSADCI_CONFIG                                                  = 0x7B
	TMC4671_STATUS_FLAGS                                                        = 0x7C
	TMC4671_STATUS_MASK                                                         = 0x7D
	CHIPINFO_ADDR_SI_TYPE                                                       = 0
	CHIPINFO_ADDR_SI_VERSION                                                    = 1
	CHIPINFO_ADDR_SI_DATA                                                       = 2
	CHIPINFO_ADDR_SI_TIME                                                       = 3
	CHIPINFO_ADDR_SI_VARIANT                                                    = 4
	CHIPINFO_ADDR_SI_BUILD                                                      = 5
	ADC_RAW_ADDR_ADC_I1_RAW_ADC_I0_RAW                                          = 0
	ADC_RAW_ADDR_ADC_AGPI_A_RAW_ADC_VM_RAW                                      = 1
	ADC_RAW_ADDR_ADC_AENC_UX_RAW_ADC_AGPI_B_RAW                                 = 2
	ADC_RAW_ADDR_ADC_AENC_WY_RAW_ADC_AENC_VN_RAW                                = 3
	CONFIG_ADDR_BIQUAD_X_A_1                                                    = 1
	CONFIG_ADDR_BIQUAD_X_A_2                                                    = 2
	CONFIG_ADDR_BIQUAD_X_B_0                                                    = 4
	CONFIG_ADDR_BIQUAD_X_B_1                                                    = 5
	CONFIG_ADDR_BIQUAD_X_B_2                                                    = 6
	CONFIG_ADDR_BIQUAD_X_ENABLE                                                 = 7
	CONFIG_ADDR_BIQUAD_V_A_1                                                    = 9
	CONFIG_ADDR_BIQUAD_V_A_2                                                    = 10
	CONFIG_ADDR_BIQUAD_V_B_0                                                    = 12
	CONFIG_ADDR_BIQUAD_V_B_1                                                    = 13
	CONFIG_ADDR_BIQUAD_V_B_2                                                    = 14
	CONFIG_ADDR_BIQUAD_V_ENABLE                                                 = 15
	CONFIG_ADDR_BIQUAD_T_A_1                                                    = 17
	CONFIG_ADDR_BIQUAD_T_A_2                                                    = 18
	CONFIG_ADDR_BIQUAD_T_B_0                                                    = 20
	CONFIG_ADDR_BIQUAD_T_B_1                                                    = 21
	CONFIG_ADDR_BIQUAD_T_B_2                                                    = 22
	CONFIG_ADDR_BIQUAD_T_ENABLE                                                 = 23
	CONFIG_ADDR_BIQUAD_F_A_1                                                    = 25
	CONFIG_ADDR_BIQUAD_F_A_2                                                    = 26
	CONFIG_ADDR_BIQUAD_F_B_0                                                    = 28
	CONFIG_ADDR_BIQUAD_F_B_1                                                    = 29
	CONFIG_ADDR_BIQUAD_F_B_2                                                    = 30
	CONFIG_ADDR_BIQUAD_F_ENABLE                                                 = 31
	CONFIG_ADDR_PRBS_AMPLITUDE                                                  = 32
	CONFIG_ADDR_PRBS_DOWN_SAMPLING_RATIO                                        = 33
	CONFIG_ADDR_FEED_FORWARD_VELOCITY_GAIN                                      = 40
	CONFIG_ADDR_FEED_FORWARD_VELOCITY_FILTER_CONSTANT                           = 41
	CONFIG_ADDR_FEED_FORWARD_TORQUE_GAIN                                        = 42
	CONFIG_ADDR_FEED_FORWARD_TORQUE_FILTER_CONSTANT                             = 43
	CONFIG_ADDR_VELOCITY_METER_PPTM_MIN_POS_DEV                                 = 50
	CONFIG_ADDR_REF_SWITCH_CONFIG                                               = 51
	CONFIG_ADDR_ENCODER_INIT_HALL_ENABLE                                        = 52
	CONFIG_ADDR_SINGLE_PIN_IF_STATUS_CFG                                        = 60
	CONFIG_ADDR_SINGLE_PIN_IF_SCALE_OFFSET                                      = 61
	PID_ERROR_ADDR_PID_TORQUE_ERROR                                             = 0
	PID_ERROR_ADDR_PID_FLUX_ERROR                                               = 1
	PID_ERROR_ADDR_PID_VELOCITY_ERROR                                           = 2
	PID_ERROR_ADDR_PID_POSITION_ERROR                                           = 3
	PID_ERROR_ADDR_PID_TORQUE_ERROR_SUM                                         = 4
	PID_ERROR_ADDR_PID_FLUX_ERROR_SUM                                           = 5
	PID_ERROR_ADDR_PID_VELOCITY_ERROR_SUM                                       = 6
	PID_ERROR_ADDR_PID_POSITION_ERROR_SUM                                       = 7
	INTERIM_ADDR_PIDIN_TARGET_TORQUE                                            = 0
	INTERIM_ADDR_PIDIN_TARGET_FLUX                                              = 1
	INTERIM_ADDR_PIDIN_TARGET_VELOCITY                                          = 2
	INTERIM_ADDR_PIDIN_TARGET_POSITION                                          = 3
	INTERIM_ADDR_PIDOUT_TARGET_TORQUE                                           = 4
	INTERIM_ADDR_PIDOUT_TARGET_FLUX                                             = 5
	INTERIM_ADDR_PIDOUT_TARGET_VELOCITY                                         = 6
	INTERIM_ADDR_PIDOUT_TARGET_POSITION                                         = 7
	INTERIM_ADDR_FOC_IWY_IUX                                                    = 8
	INTERIM_ADDR_FOC_IV                                                         = 9
	INTERIM_ADDR_FOC_IB_IA                                                      = 10
	INTERIM_ADDR_FOC_IQ_ID                                                      = 11
	INTERIM_ADDR_FOC_UQ_UD                                                      = 12
	INTERIM_ADDR_FOC_UQ_UD_LIMITED                                              = 13
	INTERIM_ADDR_FOC_UB_UA                                                      = 14
	INTERIM_ADDR_FOC_UWY_UUX                                                    = 15
	INTERIM_ADDR_FOC_UV                                                         = 16
	INTERIM_ADDR_PWM_WY_UX                                                      = 17
	INTERIM_ADDR_PWM_UV                                                         = 18
	INTERIM_ADDR_ADC_I1_I0                                                      = 19
	INTERIM_ADDR_PID_TORQUE_TARGET_FLUX_TARGET_TORQUE_ACTUAL_FLUX_ACTUAL_DIV256 = 20
	INTERIM_ADDR_PID_TORQUE_TARGET_TORQUE_ACTUAL                                = 21
	INTERIM_ADDR_PID_FLUX_TARGET_FLUX_ACTUAL                                    = 22
	INTERIM_ADDR_PID_VELOCITY_TARGET_VELOCITY_ACTUAL_DIV256                     = 23
	INTERIM_ADDR_PID_VELOCITY_TARGET_VELOCITY_ACTUAL                            = 24
	INTERIM_ADDR_PID_POSITION_TARGET_POSITION_ACTUAL_DIV256                     = 25
	INTERIM_ADDR_PID_POSITION_TARGET_POSITION_ACTUAL                            = 26
	INTERIM_ADDR_FF_VELOCITY                                                    = 27
	INTERIM_ADDR_FF_TORQUE                                                      = 28
	INTERIM_ADDR_ACTUAL_VELOCITY_PPTM                                           = 29
	INTERIM_ADDR_REF_SWITCH_STATUS                                              = 30
	INTERIM_ADDR_HOME_POSITION                                                  = 31
	INTERIM_ADDR_LEFT_POSITION                                                  = 32
	INTERIM_ADDR_RIGHT_POSITION                                                 = 33
	INTERIM_ADDR_ENC_INIT_HALL_STATUS                                           = 34
	INTERIM_ADDR_ENC_INIT_HALL_PHI_E_ABN_OFFSET                                 = 35
	INTERIM_ADDR_ENC_INIT_HALL_PHI_E_AENC_OFFSET                                = 36
	INTERIM_ADDR_ENC_INIT_HALL_PHI_A_AENC_OFFSET                                = 37
	INTERIM_ADDR_enc_init_mini_move_u_d_status                                  = 40
	INTERIM_ADDR_enc_init_mini_move_phi_e_phi_e_offset                          = 41
	INTERIM_ADDR_SINGLE_PIN_IF_PWM_DUTY_CYCLE_TORQUE_TARGET                     = 42
	INTERIM_ADDR_SINGLE_PIN_IF_VELOCITY_TARGET                                  = 43
	INTERIM_ADDR_SINGLE_PIN_IF_POSITION_TARGET                                  = 44
	INTERIM_ADDR_DEBUG_VALUE_1_0                                                = 192
	INTERIM_ADDR_DEBUG_VALUE_3_2                                                = 193
	INTERIM_ADDR_DEBUG_VALUE_5_4                                                = 194
	INTERIM_ADDR_DEBUG_VALUE_7_6                                                = 195
	INTERIM_ADDR_DEBUG_VALUE_9_8                                                = 196
	INTERIM_ADDR_DEBUG_VALUE_11_10                                              = 197
	INTERIM_ADDR_DEBUG_VALUE_13_12                                              = 198
	INTERIM_ADDR_DEBUG_VALUE_15_14                                              = 199
	INTERIM_ADDR_DEBUG_VALUE_16                                                 = 200
	INTERIM_ADDR_DEBUG_VALUE_17                                                 = 201
	INTERIM_ADDR_DEBUG_VALUE_18                                                 = 202
	INTERIM_ADDR_DEBUG_VALUE_19                                                 = 203
	INTERIM_ADDR_CONFIG_REG_0                                                   = 208
	INTERIM_ADDR_CONFIG_REG_1                                                   = 209
	INTERIM_ADDR_CTRL_PARAM_10                                                  = 210
	INTERIM_ADDR_CTRL_PARAM_32                                                  = 211
	INTERIM_ADDR_STATUS_REG_0                                                   = 212
	INTERIM_ADDR_STATUS_REG_1                                                   = 213
	INTERIM_ADDR_STATUS_PARAM_10                                                = 214
	INTERIM_ADDR_STATUS_PARAM_32                                                = 215
	TMC4671_NO_MOTOR                                                            = 0
	TMC4671_SINGLE_PHASE_DC                                                     = 1
	TMC4671_TWO_PHASE_STEPPER                                                   = 2
	TMC4671_THREE_PHASE_BLDC                                                    = 3
	TMC4671_MOTION_MODE_STOPPED                                                 = 0
	TMC4671_MOTION_MODE_TORQUE                                                  = 1
	TMC4671_MOTION_MODE_VELOCITY                                                = 2
	TMC4671_MOTION_MODE_POSITION                                                = 3
	TMC4671_MOTION_MODE_PRBS_FLUX                                               = 4
	TMC4671_MOTION_MODE_PRBS_TORQUE                                             = 5
	TMC4671_MOTION_MODE_PRBS_VELOCITY                                           = 6
	TMC4671_MOTION_MODE_PRBS_POSITION                                           = 7
	TMC4671_MOTION_MODE_UQ_UD_EXT                                               = 8
	TMC4671_PHI_E_EXTERNAL                                                      = 1
	TMC4671_PHI_E_OPEN_LOOP                                                     = 2
	TMC4671_PHI_E_ABN                                                           = 3
	TMC4671_PHI_E_HALL                                                          = 5
	TMC4671_PHI_E_AENC                                                          = 6
	TMC4671_PHI_A_AENC                                                          = 7
	TMC4671_VELOCITY_PHI_E_SELECTION                                            = 0
	TMC4671_VELOCITY_PHI_E_EXT                                                  = 1
	TMC4671_VELOCITY_PHI_E_OPENLOOP                                             = 2
	TMC4671_VELOCITY_PHI_E_ABN                                                  = 3
	TMC4671_VELOCITY_PHI_E_HAL                                                  = 5
	TMC4671_VELOCITY_PHI_E_AENC                                                 = 6
	TMC4671_VELOCITY_PHI_A_AENC                                                 = 7
	TMC4671_VELOCITY_PHI_M_ABN                                                  = 9
	TMC4671_VELOCITY_PHI_M_ABN_2                                                = 10
	TMC4671_VELOCITY_PHI_M_AENC                                                 = 11
	TMC4671_VELOCITY_PHI_M_HAL                                                  = 12
	TMC4671_POSITION_PHI_E_SELECTION                                            = 0
	TMC4671_POSITION_PHI_E_EXT                                                  = 1
	TMC4671_POSITION_PHI_E_OPENLOOP                                             = 2
	TMC4671_POSITION_PHI_E_ABN                                                  = 3
	TMC4671_POSITION_PHI_E_HAL                                                  = 5
	TMC4671_POSITION_PHI_E_AENC                                                 = 6
	TMC4671_POSITION_PHI_A_AENC                                                 = 7
	TMC4671_POSITION_PHI_M_ABN                                                  = 9
	TMC4671_POSITION_PHI_M_ABN_2                                                = 10
	TMC4671_POSITION_PHI_M_AENC                                                 = 11
	TMC4671_POSITION_PHI_M_HAL                                                  = 12
	TMC4671_MOTORS                                                              = 1
	TMC4671_SI_TYPE_MASK                                                        = 0xFFFFFFFF
	TMC4671_SI_TYPE_SHIFT                                                       = 0
	TMC4671_SI_VERSION_MASK                                                     = 0xFFFFFFFF
	TMC4671_SI_VERSION_SHIFT                                                    = 0
	TMC4671_SI_DATE_MASK                                                        = 0xFFFFFFFF
	TMC4671_SI_DATE_SHIFT                                                       = 0
	TMC4671_SI_TIME_MASK                                                        = 0xFFFFFFFF
	TMC4671_SI_TIME_SHIFT                                                       = 0
	TMC4671_SI_VARIANT_MASK                                                     = 0xFFFFFFFF
	TMC4671_SI_VARIANT_SHIFT                                                    = 0
	TMC4671_SI_BUILD_MASK                                                       = 0xFFFFFFFF
	TMC4671_SI_BUILD_SHIFT                                                      = 0
	TMC4671_CHIP_INFO_ADDRESS_MASK                                              = 0x000000FF
	TMC4671_CHIP_INFO_ADDRESS_SHIFT                                             = 0
	TMC4671_ADC_I0_RAW_MASK                                                     = 0x0000FFFF
	TMC4671_ADC_I0_RAW_SHIFT                                                    = 0
	TMC4671_ADC_I1_RAW_MASK                                                     = 0xFFFF0000
	TMC4671_ADC_I1_RAW_SHIFT                                                    = 16
	TMC4671_ADC_VM_RAW_MASK                                                     = 0x0000FFFF
	TMC4671_ADC_VM_RAW_SHIFT                                                    = 0
	TMC4671_ADC_AGPI_A_RAW_MASK                                                 = 0xFFFF0000
	TMC4671_ADC_AGPI_A_RAW_SHIFT                                                = 16
	TMC4671_ADC_AGPI_B_RAW_MASK                                                 = 0x0000FFFF
	TMC4671_ADC_AGPI_B_RAW_SHIFT                                                = 0
	TMC4671_ADC_AENC_UX_RAW_MASK                                                = 0xFFFF0000
	TMC4671_ADC_AENC_UX_RAW_SHIFT                                               = 16
	TMC4671_ADC_AENC_VN_RAW_MASK                                                = 0x0000FFFF
	TMC4671_ADC_AENC_VN_RAW_SHIFT                                               = 0
	TMC4671_ADC_AENC_WY_RAW_MASK                                                = 0xFFFF0000
	TMC4671_ADC_AENC_WY_RAW_SHIFT                                               = 16
	TMC4671_ADC_RAW_ADDR_MASK                                                   = 0x000000FF
	TMC4671_ADC_RAW_ADDR_SHIFT                                                  = 0
	TMC4671_CFG_DSMODULATOR_A_MASK                                              = 0x00000003
	TMC4671_CFG_DSMODULATOR_A_SHIFT                                             = 0
	TMC4671_MCLK_POLARITY_A_MASK                                                = 0x00000004
	TMC4671_MCLK_POLARITY_A_SHIFT                                               = 2
	TMC4671_MDAT_POLARITY_A_MASK                                                = 0x00000008
	TMC4671_MDAT_POLARITY_A_SHIFT                                               = 3
	TMC4671_SEL_NCLK_MCLK_I_A_MASK                                              = 0x00000010
	TMC4671_SEL_NCLK_MCLK_I_A_SHIFT                                             = 4
	TMC4671_BLANKING_A_MASK                                                     = 0x0000FF00
	TMC4671_BLANKING_A_SHIFT                                                    = 8
	TMC4671_CFG_DSMODULATOR_B_MASK                                              = 0x00030000
	TMC4671_CFG_DSMODULATOR_B_SHIFT                                             = 16
	TMC4671_MCLK_POLARITY_B_MASK                                                = 0x00040000
	TMC4671_MCLK_POLARITY_B_SHIFT                                               = 18
	TMC4671_MDAT_POLARITY_B_MASK                                                = 0x00080000
	TMC4671_MDAT_POLARITY_B_SHIFT                                               = 19
	TMC4671_SEL_NCLK_MCLK_I_B_MASK                                              = 0x00100000
	TMC4671_SEL_NCLK_MCLK_I_B_SHIFT                                             = 20
	TMC4671_BLANKING_B_MASK                                                     = 0xFF000000
	TMC4671_BLANKING_B_SHIFT                                                    = 24
	TMC4671_DSADC_MCLK_A_MASK                                                   = 0xFFFFFFFF
	TMC4671_DSADC_MCLK_A_SHIFT                                                  = 0
	TMC4671_DSADC_MCLK_B_MASK                                                   = 0xFFFFFFFF
	TMC4671_DSADC_MCLK_B_SHIFT                                                  = 0
	TMC4671_DSADC_MDEC_A_MASK                                                   = 0x0000FFFF
	TMC4671_DSADC_MDEC_A_SHIFT                                                  = 0
	TMC4671_DSADC_MDEC_B_MASK                                                   = 0xFFFF0000
	TMC4671_DSADC_MDEC_B_SHIFT                                                  = 16
	TMC4671_ADC_I1_OFFSET_MASK                                                  = 0x0000FFFF
	TMC4671_ADC_I1_OFFSET_SHIFT                                                 = 0
	TMC4671_ADC_I1_SCALE_MASK                                                   = 0xFFFF0000
	TMC4671_ADC_I1_SCALE_SHIFT                                                  = 16
	TMC4671_ADC_I0_OFFSET_MASK                                                  = 0x0000FFFF
	TMC4671_ADC_I0_OFFSET_SHIFT                                                 = 0
	TMC4671_ADC_I0_SCALE_MASK                                                   = 0xFFFF0000
	TMC4671_ADC_I0_SCALE_SHIFT                                                  = 16
	TMC4671_ADC_I0_SELECT_MASK                                                  = 0x000000FF
	TMC4671_ADC_I0_SELECT_SHIFT                                                 = 0
	TMC4671_ADC_I1_SELECT_MASK                                                  = 0x0000FF00
	TMC4671_ADC_I1_SELECT_SHIFT                                                 = 8
	TMC4671_ADC_I_UX_SELECT_MASK                                                = 0x03000000
	TMC4671_ADC_I_UX_SELECT_SHIFT                                               = 24
	TMC4671_ADC_I_V_SELECT_MASK                                                 = 0x0C000000
	TMC4671_ADC_I_V_SELECT_SHIFT                                                = 26
	TMC4671_ADC_I_WY_SELECT_MASK                                                = 0x30000000
	TMC4671_ADC_I_WY_SELECT_SHIFT                                               = 28
	TMC4671_ADC_I0_EXT_MASK                                                     = 0x0000FFFF
	TMC4671_ADC_I0_EXT_SHIFT                                                    = 0
	TMC4671_ADC_I1_EXT_MASK                                                     = 0xFFFF0000
	TMC4671_ADC_I1_EXT_SHIFT                                                    = 16
	TMC4671_ADC_I0_MASK                                                         = 0x0000000F
	TMC4671_ADC_I0_SHIFT                                                        = 0
	TMC4671_ADC_I1_MASK                                                         = 0x000000F0
	TMC4671_ADC_I1_SHIFT                                                        = 4
	TMC4671_ADC_VM_MASK                                                         = 0x00000F00
	TMC4671_ADC_VM_SHIFT                                                        = 8
	TMC4671_ADC_AGPI_A_MASK                                                     = 0x0000F000
	TMC4671_ADC_AGPI_A_SHIFT                                                    = 12
	TMC4671_ADC_AGPI_B_MASK                                                     = 0x000F0000
	TMC4671_ADC_AGPI_B_SHIFT                                                    = 16
	TMC4671_ADC_AENC_UX_MASK                                                    = 0x00F00000
	TMC4671_ADC_AENC_UX_SHIFT                                                   = 20
	TMC4671_ADC_AENC_VN_MASK                                                    = 0x0F000000
	TMC4671_ADC_AENC_VN_SHIFT                                                   = 24
	TMC4671_ADC_AENC_WY_MASK                                                    = 0xF0000000
	TMC4671_ADC_AENC_WY_SHIFT                                                   = 28
	TMC4671_AENC_0_OFFSET_MASK                                                  = 0x0000FFFF
	TMC4671_AENC_0_OFFSET_SHIFT                                                 = 0
	TMC4671_AENC_0_SCALE_MASK                                                   = 0xFFFF0000
	TMC4671_AENC_0_SCALE_SHIFT                                                  = 16
	TMC4671_AENC_1_OFFSET_MASK                                                  = 0x0000FFFF
	TMC4671_AENC_1_OFFSET_SHIFT                                                 = 0
	TMC4671_AENC_1_SCALE_MASK                                                   = 0xFFFF0000
	TMC4671_AENC_1_SCALE_SHIFT                                                  = 16
	TMC4671_AENC_2_OFFSET_MASK                                                  = 0x0000FFFF
	TMC4671_AENC_2_OFFSET_SHIFT                                                 = 0
	TMC4671_AENC_2_SCALE_MASK                                                   = 0xFFFF0000
	TMC4671_AENC_2_SCALE_SHIFT                                                  = 16
	TMC4671_AENC_0_SELECT_MASK                                                  = 0x000000FF
	TMC4671_AENC_0_SELECT_SHIFT                                                 = 0
	TMC4671_AENC_1_SELECT_MASK                                                  = 0x0000FF00
	TMC4671_AENC_1_SELECT_SHIFT                                                 = 8
	TMC4671_AENC_2_SELECT_MASK                                                  = 0x00FF0000
	TMC4671_AENC_2_SELECT_SHIFT                                                 = 16
	TMC4671_ADC_IUX_MASK                                                        = 0x0000FFFF
	TMC4671_ADC_IUX_SHIFT                                                       = 0
	TMC4671_ADC_IWY_MASK                                                        = 0xFFFF0000
	TMC4671_ADC_IWY_SHIFT                                                       = 16
	TMC4671_ADC_IV_MASK                                                         = 0x0000FFFF
	TMC4671_ADC_IV_SHIFT                                                        = 0
	TMC4671_AENC_UX_MASK                                                        = 0x0000FFFF
	TMC4671_AENC_UX_SHIFT                                                       = 0
	TMC4671_AENC_WY_MASK                                                        = 0xFFFF0000
	TMC4671_AENC_WY_SHIFT                                                       = 16
	TMC4671_AENC_VN_MASK                                                        = 0x0000FFFF
	TMC4671_AENC_VN_SHIFT                                                       = 0
	TMC4671_PWM_POLARITIES_0_MASK                                               = 0x00000001
	TMC4671_PWM_POLARITIES_0_SHIFT                                              = 0
	TMC4671_PWM_POLARITIES_1_MASK                                               = 0x00000002
	TMC4671_PWM_POLARITIES_1_SHIFT                                              = 1
	TMC4671_PWM_MAXCNT_MASK                                                     = 0x0000FFFF
	TMC4671_PWM_MAXCNT_SHIFT                                                    = 0
	TMC4671_PWM_BBM_L_MASK                                                      = 0x000000FF
	TMC4671_PWM_BBM_L_SHIFT                                                     = 0
	TMC4671_PWM_CHOP_MASK                                                       = 0x000000FF
	TMC4671_PWM_CHOP_SHIFT                                                      = 0
	TMC4671_PWM_SV_MASK                                                         = 0x00000100
	TMC4671_PWM_SV_SHIFT                                                        = 8
	TMC4671_N_POLE_PAIRS_MASK                                                   = 0x0000FFFF
	TMC4671_N_POLE_PAIRS_SHIFT                                                  = 0
	TMC4671_MOTOR_TYPE_MASK                                                     = 0x00FF0000
	TMC4671_MOTOR_TYPE_SHIFT                                                    = 16
	TMC4671_PHI_E_EXT_MASK                                                      = 0x0000FFFF
	TMC4671_PHI_E_EXT_SHIFT                                                     = 0
	TMC4671_OPENLOOP_PHI_DIRECTION_MASK                                         = 0x00001000
	TMC4671_OPENLOOP_PHI_DIRECTION_SHIFT                                        = 12
	TMC4671_OPENLOOP_ACCELERATION_MASK                                          = 0x000FFFFF
	TMC4671_OPENLOOP_ACCELERATION_SHIFT                                         = 0
	TMC4671_OPENLOOP_VELOCITY_TARGET_MASK                                       = 0xFFFFFFFF
	TMC4671_OPENLOOP_VELOCITY_TARGET_SHIFT                                      = 0
	TMC4671_OPENLOOP_VELOCITY_ACTUAL_MASK                                       = 0xFFFFFFFF
	TMC4671_OPENLOOP_VELOCITY_ACTUAL_SHIFT                                      = 0
	TMC4671_OPENLOOP_PHI_MASK                                                   = 0x0000FFFF
	TMC4671_OPENLOOP_PHI_SHIFT                                                  = 0
	TMC4671_UD_EXT_MASK                                                         = 0x0000FFFF
	TMC4671_UD_EXT_SHIFT                                                        = 0
	TMC4671_UQ_EXT_MASK                                                         = 0xFFFF0000
	TMC4671_UQ_EXT_SHIFT                                                        = 16
	TMC4671_ABN_APOL_MASK                                                       = 0x00000001
	TMC4671_ABN_APOL_SHIFT                                                      = 0
	TMC4671_ABN_BPOL_MASK                                                       = 0x00000002
	TMC4671_ABN_BPOL_SHIFT                                                      = 1
	TMC4671_ABN_NPOL_MASK                                                       = 0x00000004
	TMC4671_ABN_NPOL_SHIFT                                                      = 2
	TMC4671_USE_ABN_AS_N_MASK                                                   = 0x00000008
	TMC4671_USE_ABN_AS_N_SHIFT                                                  = 3
	TMC4671_ABN_CLN_MASK                                                        = 0x00000100
	TMC4671_ABN_CLN_SHIFT                                                       = 8
	TMC4671_ABN_DIRECTION_MASK                                                  = 0x00001000
	TMC4671_ABN_DIRECTION_SHIFT                                                 = 12
	TMC4671_ABN_DECODER_PPR_MASK                                                = 0x00FFFFFF
	TMC4671_ABN_DECODER_PPR_SHIFT                                               = 0
	TMC4671_ABN_DECODER_COUNT_MASK                                              = 0x00FFFFFF
	TMC4671_ABN_DECODER_COUNT_SHIFT                                             = 0
	TMC4671_ABN_DECODER_COUNT_N_MASK                                            = 0x00FFFFFF
	TMC4671_ABN_DECODER_COUNT_N_SHIFT                                           = 0
	TMC4671_ABN_DECODER_PHI_M_OFFSET_MASK                                       = 0x0000FFFF
	TMC4671_ABN_DECODER_PHI_M_OFFSET_SHIFT                                      = 0
	TMC4671_ABN_DECODER_PHI_E_OFFSET_MASK                                       = 0xFFFF0000
	TMC4671_ABN_DECODER_PHI_E_OFFSET_SHIFT                                      = 16
	TMC4671_ABN_DECODER_PHI_M_MASK                                              = 0x0000FFFF
	TMC4671_ABN_DECODER_PHI_M_SHIFT                                             = 0
	TMC4671_ABN_DECODER_PHI_E_MASK                                              = 0xFFFF0000
	TMC4671_ABN_DECODER_PHI_E_SHIFT                                             = 16
	TMC4671_ABN_2_APOL_MASK                                                     = 0x00000001
	TMC4671_ABN_2_APOL_SHIFT                                                    = 0
	TMC4671_ABN_2_BPOL_MASK                                                     = 0x00000002
	TMC4671_ABN_2_BPOL_SHIFT                                                    = 1
	TMC4671_ABN_2_NPOL_MASK                                                     = 0x00000004
	TMC4671_ABN_2_NPOL_SHIFT                                                    = 2
	TMC4671_USE_ABN_2_AS_N_MASK                                                 = 0x00000008
	TMC4671_USE_ABN_2_AS_N_SHIFT                                                = 3
	TMC4671_ABN_2_CLN_MASK                                                      = 0x00000100
	TMC4671_ABN_2_CLN_SHIFT                                                     = 8
	TMC4671_ABN_2_DIRECTION_MASK                                                = 0x00001000
	TMC4671_ABN_2_DIRECTION_SHIFT                                               = 12
	TMC4671_ABN_2_DECODER_PPR_MASK                                              = 0x00FFFFFF
	TMC4671_ABN_2_DECODER_PPR_SHIFT                                             = 0
	TMC4671_ABN_2_DECODER_COUNT_MASK                                            = 0x00FFFFFF
	TMC4671_ABN_2_DECODER_COUNT_SHIFT                                           = 0
	TMC4671_ABN_2_DECODER_COUNT_N_MASK                                          = 0x00FFFFFF
	TMC4671_ABN_2_DECODER_COUNT_N_SHIFT                                         = 0
	TMC4671_ABN_2_DECODER_PHI_M_OFFSET_MASK                                     = 0x0000FFFF
	TMC4671_ABN_2_DECODER_PHI_M_OFFSET_SHIFT                                    = 0
	TMC4671_ABN_2_DECODER_PHI_M_MASK                                            = 0x0000FFFF
	TMC4671_ABN_2_DECODER_PHI_M_SHIFT                                           = 0
	TMC4671_HALL_POLARITY_MASK                                                  = 0x00000001
	TMC4671_HALL_POLARITY_SHIFT                                                 = 0
	TMC4671_HALL_SYNCHRONOUS_PWM_SAMPLING_MASK                                  = 0x00000010
	TMC4671_HALL_SYNCHRONOUS_PWM_SAMPLING_SHIFT                                 = 4
	TMC4671_HALL_INTERPOLATION_MASK                                             = 0x00000100
	TMC4671_HALL_INTERPOLATION_SHIFT                                            = 8
	TMC4671_HALL_DIRECTION_MASK                                                 = 0x00001000
	TMC4671_HALL_DIRECTION_SHIFT                                                = 12
	TMC4671_HALL_BLANK_MASK                                                     = 0x0FFF0000
	TMC4671_HALL_BLANK_SHIFT                                                    = 16
	TMC4671_HALL_POSITION_000_MASK                                              = 0x0000FFFF
	TMC4671_HALL_POSITION_000_SHIFT                                             = 0
	TMC4671_HALL_POSITION_060_MASK                                              = 0xFFFF0000
	TMC4671_HALL_POSITION_060_SHIFT                                             = 16
	TMC4671_HALL_POSITION_120_MASK                                              = 0x0000FFFF
	TMC4671_HALL_POSITION_120_SHIFT                                             = 0
	TMC4671_HALL_POSITION_180_MASK                                              = 0xFFFF0000
	TMC4671_HALL_POSITION_180_SHIFT                                             = 16
	TMC4671_HALL_POSITION_240_MASK                                              = 0x0000FFFF
	TMC4671_HALL_POSITION_240_SHIFT                                             = 0
	TMC4671_HALL_POSITION_300_MASK                                              = 0xFFFF0000
	TMC4671_HALL_POSITION_300_SHIFT                                             = 16
	TMC4671_HALL_PHI_M_OFFSET_MASK                                              = 0x0000FFFF
	TMC4671_HALL_PHI_M_OFFSET_SHIFT                                             = 0
	TMC4671_HALL_PHI_E_OFFSET_MASK                                              = 0xFFFF0000
	TMC4671_HALL_PHI_E_OFFSET_SHIFT                                             = 16
	TMC4671_HALL_DPHI_MAX_MASK                                                  = 0x0000FFFF
	TMC4671_HALL_DPHI_MAX_SHIFT                                                 = 0
	TMC4671_HALL_PHI_E_MASK                                                     = 0x0000FFFF
	TMC4671_HALL_PHI_E_SHIFT                                                    = 0
	TMC4671_HALL_PHI_E_INTERPOLATED_MASK                                        = 0xFFFF0000
	TMC4671_HALL_PHI_E_INTERPOLATED_SHIFT                                       = 16
	TMC4671_HALL_PHI_M_MASK                                                     = 0x0000FFFF
	TMC4671_HALL_PHI_M_SHIFT                                                    = 0
	TMC4671_AENC_DECODER_MODE_0_MASK                                            = 0x00000001
	TMC4671_AENC_DECODER_MODE_0_SHIFT                                           = 0
	TMC4671_AENC_DECODER_MODE_12_MASK                                           = 0x00001000
	TMC4671_AENC_DECODER_MODE_12_SHIFT                                          = 12
	TMC4671_AENC_DECODER_N_THRESHOLD_MASK                                       = 0x0000FFFF
	TMC4671_AENC_DECODER_N_THRESHOLD_SHIFT                                      = 0
	TMC4671_AENC_DECODER_PHI_A_RAW_MASK                                         = 0x0000FFFF
	TMC4671_AENC_DECODER_PHI_A_RAW_SHIFT                                        = 0
	TMC4671_AENC_DECODER_PHI_A_OFFSET_MASK                                      = 0x0000FFFF
	TMC4671_AENC_DECODER_PHI_A_OFFSET_SHIFT                                     = 0
	TMC4671_AENC_DECODER_PHI_A_MASK                                             = 0x0000FFFF
	TMC4671_AENC_DECODER_PHI_A_SHIFT                                            = 0
	TMC4671_AENC_DECODER_PPR_MASK                                               = 0x0000FFFF
	TMC4671_AENC_DECODER_PPR_SHIFT                                              = 0
	TMC4671_AENC_DECODER_COUNT_MASK                                             = 0xFFFFFFFF
	TMC4671_AENC_DECODER_COUNT_SHIFT                                            = 0
	TMC4671_AENC_DECODER_COUNT_N_MASK                                           = 0xFFFFFFFF
	TMC4671_AENC_DECODER_COUNT_N_SHIFT                                          = 0
	TMC4671_AENC_DECODER_PHI_M_OFFSET_MASK                                      = 0x0000FFFF
	TMC4671_AENC_DECODER_PHI_M_OFFSET_SHIFT                                     = 0
	TMC4671_AENC_DECODER_PHI_E_OFFSET_MASK                                      = 0xFFFF0000
	TMC4671_AENC_DECODER_PHI_E_OFFSET_SHIFT                                     = 16
	TMC4671_AENC_DECODER_PHI_M_MASK                                             = 0x0000FFFF
	TMC4671_AENC_DECODER_PHI_M_SHIFT                                            = 0
	TMC4671_AENC_DECODER_PHI_E_MASK                                             = 0xFFFF0000
	TMC4671_AENC_DECODER_PHI_E_SHIFT                                            = 16
	TMC4671_BIQUAD_X_A_1_MASK                                                   = 0xFFFFFFFF
	TMC4671_BIQUAD_X_A_1_SHIFT                                                  = 0
	TMC4671_BIQUAD_X_A_2_MASK                                                   = 0xFFFFFFFF
	TMC4671_BIQUAD_X_A_2_SHIFT                                                  = 0
	TMC4671_BIQUAD_X_B_0_MASK                                                   = 0xFFFFFFFF
	TMC4671_BIQUAD_X_B_0_SHIFT                                                  = 0
	TMC4671_BIQUAD_X_B_1_MASK                                                   = 0xFFFFFFFF
	TMC4671_BIQUAD_X_B_1_SHIFT                                                  = 0
	TMC4671_BIQUAD_X_B_2_MASK                                                   = 0xFFFFFFFF
	TMC4671_BIQUAD_X_B_2_SHIFT                                                  = 0
	TMC4671_BIQUAD_X_ENABLE_MASK                                                = 0xFFFFFFFF
	TMC4671_BIQUAD_X_ENABLE_SHIFT                                               = 0
	TMC4671_BIQUAD_V_A_1_MASK                                                   = 0xFFFFFFFF
	TMC4671_BIQUAD_V_A_1_SHIFT                                                  = 0
	TMC4671_BIQUAD_V_A_2_MASK                                                   = 0xFFFFFFFF
	TMC4671_BIQUAD_V_A_2_SHIFT                                                  = 0
	TMC4671_BIQUAD_V_B_0_MASK                                                   = 0xFFFFFFFF
	TMC4671_BIQUAD_V_B_0_SHIFT                                                  = 0
	TMC4671_BIQUAD_V_B_1_MASK                                                   = 0xFFFFFFFF
	TMC4671_BIQUAD_V_B_1_SHIFT                                                  = 0
	TMC4671_BIQUAD_V_B_2_MASK                                                   = 0xFFFFFFFF
	TMC4671_BIQUAD_V_B_2_SHIFT                                                  = 0
	TMC4671_BIQUAD_V_ENABLE_MASK                                                = 0xFFFFFFFF
	TMC4671_BIQUAD_V_ENABLE_SHIFT                                               = 0
	TMC4671_BIQUAD_T_A_1_MASK                                                   = 0xFFFFFFFF
	TMC4671_BIQUAD_T_A_1_SHIFT                                                  = 0
	TMC4671_BIQUAD_T_A_2_MASK                                                   = 0xFFFFFFFF
	TMC4671_BIQUAD_T_A_2_SHIFT                                                  = 0
	TMC4671_BIQUAD_T_B_0_MASK                                                   = 0xFFFFFFFF
	TMC4671_BIQUAD_T_B_0_SHIFT                                                  = 0
	TMC4671_BIQUAD_T_B_1_MASK                                                   = 0xFFFFFFFF
	TMC4671_BIQUAD_T_B_1_SHIFT                                                  = 0
	TMC4671_BIQUAD_T_B_2_MASK                                                   = 0xFFFFFFFF
	TMC4671_BIQUAD_T_B_2_SHIFT                                                  = 0
	TMC4671_BIQUAD_T_ENABLE_MASK                                                = 0xFFFFFFFF
	TMC4671_BIQUAD_T_ENABLE_SHIFT                                               = 0
	TMC4671_BIQUAD_F_A_1_MASK                                                   = 0xFFFFFFFF
	TMC4671_BIQUAD_F_A_1_SHIFT                                                  = 0
	TMC4671_BIQUAD_F_A_2_MASK                                                   = 0xFFFFFFFF
	TMC4671_BIQUAD_F_A_2_SHIFT                                                  = 0
	TMC4671_BIQUAD_F_B_0_MASK                                                   = 0xFFFFFFFF
	TMC4671_BIQUAD_F_B_0_SHIFT                                                  = 0
	TMC4671_BIQUAD_F_B_1_MASK                                                   = 0xFFFFFFFF
	TMC4671_BIQUAD_F_B_1_SHIFT                                                  = 0
	TMC4671_BIQUAD_F_B_2_MASK                                                   = 0xFFFFFFFF
	TMC4671_BIQUAD_F_B_2_SHIFT                                                  = 0
	TMC4671_BIQUAD_F_ENABLE_MASK                                                = 0xFFFFFFFF
	TMC4671_BIQUAD_F_ENABLE_SHIFT                                               = 0
	TMC4671_PRBS_AMPLITUDE_MASK                                                 = 0xFFFFFFFF
	TMC4671_PRBS_AMPLITUDE_SHIFT                                                = 0
	TMC4671_PRBS_DOWN_SAMPLING_RATIO_MASK                                       = 0xFFFFFFFF
	TMC4671_PRBS_DOWN_SAMPLING_RATIO_SHIFT                                      = 0
	TMC4671_REF_SWITCH_ENABLE_MASK                                              = 0x00000001
	TMC4671_REF_SWITCH_ENABLE_SHIFT                                             = 0
	TMC4671_HOME_SWITCH_POLARITY_MASK                                           = 0x00000002
	TMC4671_HOME_SWITCH_POLARITY_SHIFT                                          = 1
	TMC4671_LEFT_SWITCH_POLARITY_MASK                                           = 0x00000004
	TMC4671_LEFT_SWITCH_POLARITY_SHIFT                                          = 2
	TMC4671_RIGHT_SWITCH_POLARITY_MASK                                          = 0x00000008
	TMC4671_RIGHT_SWITCH_POLARITY_SHIFT                                         = 3
	TMC4671_ENCODER_INIT_HALL_ENABLE_MASK                                       = 0x00000001
	TMC4671_ENCODER_INIT_HALL_ENABLE_SHIFT                                      = 0
	TMC4671_SINGLE_PIN_IF_CFG_MASK                                              = 0x000000FF
	TMC4671_SINGLE_PIN_IF_CFG_SHIFT                                             = 0
	TMC4671_SINGLE_PIN_IF_STATUS_MASK                                           = 0xFFFF0000
	TMC4671_SINGLE_PIN_IF_STATUS_SHIFT                                          = 16
	TMC4671_SINGLE_PIN_IF_OFFSET_MASK                                           = 0x0000FFFF
	TMC4671_SINGLE_PIN_IF_OFFSET_SHIFT                                          = 0
	TMC4671_SINGLE_PIN_IF_SCALE_MASK                                            = 0xFFFF0000
	TMC4671_SINGLE_PIN_IF_SCALE_SHIFT                                           = 16
	TMC4671_CURRENT_I_NQ8_8_Q4_12_MASK                                          = 0x00000001
	TMC4671_CURRENT_I_NQ8_8_Q4_12_SHIFT                                         = 0
	TMC4671_CURRENT_P_NQ8_8_Q4_12_MASK                                          = 0x00000002
	TMC4671_CURRENT_P_NQ8_8_Q4_12_SHIFT                                         = 1
	TMC4671_VELOCITY_I_NQ8_8_Q4_12_MASK                                         = 0x00000004
	TMC4671_VELOCITY_I_NQ8_8_Q4_12_SHIFT                                        = 2
	TMC4671_VELOCITY_P_NQ8_8_Q4_12_MASK                                         = 0x00000008
	TMC4671_VELOCITY_P_NQ8_8_Q4_12_SHIFT                                        = 3
	TMC4671_POSITION_I_NQ8_8_Q4_12_MASK                                         = 0x00000010
	TMC4671_POSITION_I_NQ8_8_Q4_12_SHIFT                                        = 4
	TMC4671_POSITION_P_NQ8_8_Q4_12_MASK                                         = 0x00000020
	TMC4671_POSITION_P_NQ8_8_Q4_12_SHIFT                                        = 5
	TMC4671_CONFIG_ADDR_MASK                                                    = 0xFFFFFFFF
	TMC4671_CONFIG_ADDR_SHIFT                                                   = 0
	TMC4671_VELOCITY_SELECTION_MASK                                             = 0x000000FF
	TMC4671_VELOCITY_SELECTION_SHIFT                                            = 0
	TMC4671_VELOCITY_METER_SELECTION_MASK                                       = 0x0000FF00
	TMC4671_VELOCITY_METER_SELECTION_SHIFT                                      = 8
	TMC4671_POSITION_SELECTION_MASK                                             = 0x000000FF
	TMC4671_POSITION_SELECTION_SHIFT                                            = 0
	TMC4671_PHI_E_SELECTION_MASK                                                = 0x000000FF
	TMC4671_PHI_E_SELECTION_SHIFT                                               = 0
	TMC4671_PHI_E_MASK                                                          = 0x0000FFFF
	TMC4671_PHI_E_SHIFT                                                         = 0
	TMC4671_PID_FLUX_I_MASK                                                     = 0x0000FFFF
	TMC4671_PID_FLUX_I_SHIFT                                                    = 0
	TMC4671_PID_FLUX_P_MASK                                                     = 0xFFFF0000
	TMC4671_PID_FLUX_P_SHIFT                                                    = 16
	TMC4671_PID_TORQUE_I_MASK                                                   = 0x0000FFFF
	TMC4671_PID_TORQUE_I_SHIFT                                                  = 0
	TMC4671_PID_TORQUE_P_MASK                                                   = 0xFFFF0000
	TMC4671_PID_TORQUE_P_SHIFT                                                  = 16
	TMC4671_PID_VELOCITY_I_MASK                                                 = 0x0000FFFF
	TMC4671_PID_VELOCITY_I_SHIFT                                                = 0
	TMC4671_PID_VELOCITY_P_MASK                                                 = 0xFFFF0000
	TMC4671_PID_VELOCITY_P_SHIFT                                                = 16
	TMC4671_PID_POSITION_I_MASK                                                 = 0x0000FFFF
	TMC4671_PID_POSITION_I_SHIFT                                                = 0
	TMC4671_PID_POSITION_P_MASK                                                 = 0xFFFF0000
	TMC4671_PID_POSITION_P_SHIFT                                                = 16
	TMC4671_PIDOUT_UQ_UD_LIMITS_MASK                                            = 0x0000FFFF
	TMC4671_PIDOUT_UQ_UD_LIMITS_SHIFT                                           = 0
	TMC4671_PID_TORQUE_FLUX_LIMITS_MASK                                         = 0x0000FFFF
	TMC4671_PID_TORQUE_FLUX_LIMITS_SHIFT                                        = 0
	TMC4671_PID_VELOCITY_LIMIT_MASK                                             = 0xFFFFFFFF
	TMC4671_PID_VELOCITY_LIMIT_SHIFT                                            = 0
	TMC4671_PID_POSITION_LIMIT_LOW_MASK                                         = 0xFFFFFFFF
	TMC4671_PID_POSITION_LIMIT_LOW_SHIFT                                        = 0
	TMC4671_PID_POSITION_LIMIT_HIGH_MASK                                        = 0xFFFFFFFF
	TMC4671_PID_POSITION_LIMIT_HIGH_SHIFT                                       = 0
	TMC4671_MODE_MOTION_MASK                                                    = 0x000000FF
	TMC4671_MODE_MOTION_SHIFT                                                   = 0
	TMC4671_MODE_RAMP_MASK                                                      = 0x0000FF00
	TMC4671_MODE_RAMP_SHIFT                                                     = 8
	TMC4671_MODE_FF_MASK                                                        = 0x00FF0000
	TMC4671_MODE_FF_SHIFT                                                       = 16
	TMC4671_MODE_PID_SMPL_MASK                                                  = 0x7F000000
	TMC4671_MODE_PID_SMPL_SHIFT                                                 = 24
	TMC4671_MODE_PID_TYPE_MASK                                                  = 0x80000000
	TMC4671_MODE_PID_TYPE_SHIFT                                                 = 31
	TMC4671_PID_FLUX_TARGET_MASK                                                = 0x0000FFFF
	TMC4671_PID_FLUX_TARGET_SHIFT                                               = 0
	TMC4671_PID_TORQUE_TARGET_MASK                                              = 0xFFFF0000
	TMC4671_PID_TORQUE_TARGET_SHIFT                                             = 16
	TMC4671_PID_FLUX_OFFSET_MASK                                                = 0x0000FFFF
	TMC4671_PID_FLUX_OFFSET_SHIFT                                               = 0
	TMC4671_PID_TORQUE_OFFSET_MASK                                              = 0xFFFF0000
	TMC4671_PID_TORQUE_OFFSET_SHIFT                                             = 16
	TMC4671_PID_VELOCITY_TARGET_MASK                                            = 0xFFFFFFFF
	TMC4671_PID_VELOCITY_TARGET_SHIFT                                           = 0
	TMC4671_PID_VELOCITY_OFFSET_MASK                                            = 0xFFFFFFFF
	TMC4671_PID_VELOCITY_OFFSET_SHIFT                                           = 0
	TMC4671_PID_POSITION_TARGET_MASK                                            = 0xFFFFFFFF
	TMC4671_PID_POSITION_TARGET_SHIFT                                           = 0
	TMC4671_PID_FLUX_ACTUAL_MASK                                                = 0x0000FFFF
	TMC4671_PID_FLUX_ACTUAL_SHIFT                                               = 0
	TMC4671_PID_TORQUE_ACTUAL_MASK                                              = 0xFFFF0000
	TMC4671_PID_TORQUE_ACTUAL_SHIFT                                             = 16
	TMC4671_PID_VELOCITY_ACTUAL_MASK                                            = 0xFFFFFFFF
	TMC4671_PID_VELOCITY_ACTUAL_SHIFT                                           = 0
	TMC4671_PID_POSITION_ACTUAL_MASK                                            = 0xFFFFFFFF
	TMC4671_PID_POSITION_ACTUAL_SHIFT                                           = 0
	TMC4671_PID_TORQUE_ERROR_MASK                                               = 0xFFFFFFFF
	TMC4671_PID_TORQUE_ERROR_SHIFT                                              = 0
	TMC4671_PID_FLUX_ERROR_MASK                                                 = 0xFFFFFFFF
	TMC4671_PID_FLUX_ERROR_SHIFT                                                = 0
	TMC4671_PID_VELOCITY_ERROR_MASK                                             = 0xFFFFFFFF
	TMC4671_PID_VELOCITY_ERROR_SHIFT                                            = 0
	TMC4671_PID_POSITION_ERROR_MASK                                             = 0xFFFFFFFF
	TMC4671_PID_POSITION_ERROR_SHIFT                                            = 0
	TMC4671_PID_TORQUE_ERROR_SUM_MASK                                           = 0xFFFFFFFF
	TMC4671_PID_TORQUE_ERROR_SUM_SHIFT                                          = 0
	TMC4671_PID_FLUX_ERROR_SUM_MASK                                             = 0xFFFFFFFF
	TMC4671_PID_FLUX_ERROR_SUM_SHIFT                                            = 0
	TMC4671_PID_VELOCITY_ERROR_SUM_MASK                                         = 0xFFFFFFFF
	TMC4671_PID_VELOCITY_ERROR_SUM_SHIFT                                        = 0
	TMC4671_PID_POSITION_ERROR_SUM_MASK                                         = 0xFFFFFFFF
	TMC4671_PID_POSITION_ERROR_SUM_SHIFT                                        = 0
	TMC4671_PID_ERROR_ADDR_MASK                                                 = 0x000000FF
	TMC4671_PID_ERROR_ADDR_SHIFT                                                = 0
	TMC4671_PIDIN_TARGET_TORQUE_MASK                                            = 0xFFFFFFFF
	TMC4671_PIDIN_TARGET_TORQUE_SHIFT                                           = 0
	TMC4671_PIDIN_TARGET_FLUX_MASK                                              = 0xFFFFFFFF
	TMC4671_PIDIN_TARGET_FLUX_SHIFT                                             = 0
	TMC4671_PIDIN_TARGET_VELOCITY_MASK                                          = 0xFFFFFFFF
	TMC4671_PIDIN_TARGET_VELOCITY_SHIFT                                         = 0
	TMC4671_PIDIN_TARGET_POSITION_MASK                                          = 0xFFFFFFFF
	TMC4671_PIDIN_TARGET_POSITION_SHIFT                                         = 0
	TMC4671_PIDOUT_TARGET_TORQUE_MASK                                           = 0xFFFFFFFF
	TMC4671_PIDOUT_TARGET_TORQUE_SHIFT                                          = 0
	TMC4671_PIDOUT_TARGET_FLUX_MASK                                             = 0xFFFFFFFF
	TMC4671_PIDOUT_TARGET_FLUX_SHIFT                                            = 0
	TMC4671_PIDOUT_TARGET_VELOCITY_MASK                                         = 0xFFFFFFFF
	TMC4671_PIDOUT_TARGET_VELOCITY_SHIFT                                        = 0
	TMC4671_PIDOUT_TARGET_POSITION_MASK                                         = 0xFFFFFFFF
	TMC4671_PIDOUT_TARGET_POSITION_SHIFT                                        = 0
	TMC4671_FOC_IUX_MASK                                                        = 0x0000FFFF
	TMC4671_FOC_IUX_SHIFT                                                       = 0
	TMC4671_FOC_IWY_MASK                                                        = 0xFFFF0000
	TMC4671_FOC_IWY_SHIFT                                                       = 16
	TMC4671_FOC_IV_MASK                                                         = 0x0000FFFF
	TMC4671_FOC_IV_SHIFT                                                        = 0
	TMC4671_FOC_IA_MASK                                                         = 0x0000FFFF
	TMC4671_FOC_IA_SHIFT                                                        = 0
	TMC4671_FOC_IB_MASK                                                         = 0xFFFF0000
	TMC4671_FOC_IB_SHIFT                                                        = 16
	TMC4671_FOC_ID_MASK                                                         = 0x0000FFFF
	TMC4671_FOC_ID_SHIFT                                                        = 0
	TMC4671_FOC_IQ_MASK                                                         = 0xFFFF0000
	TMC4671_FOC_IQ_SHIFT                                                        = 16
	TMC4671_FOC_UD_MASK                                                         = 0x0000FFFF
	TMC4671_FOC_UD_SHIFT                                                        = 0
	TMC4671_FOC_UQ_MASK                                                         = 0xFFFF0000
	TMC4671_FOC_UQ_SHIFT                                                        = 16
	TMC4671_FOC_UD_LIMITED_MASK                                                 = 0x0000FFFF
	TMC4671_FOC_UD_LIMITED_SHIFT                                                = 0
	TMC4671_FOC_UQ_LIMITED_MASK                                                 = 0xFFFF0000
	TMC4671_FOC_UQ_LIMITED_SHIFT                                                = 16
	TMC4671_FOC_UA_MASK                                                         = 0x0000FFFF
	TMC4671_FOC_UA_SHIFT                                                        = 0
	TMC4671_FOC_UB_MASK                                                         = 0xFFFF0000
	TMC4671_FOC_UB_SHIFT                                                        = 16
	TMC4671_FOC_UUX_MASK                                                        = 0x0000FFFF
	TMC4671_FOC_UUX_SHIFT                                                       = 0
	TMC4671_FOC_UWY_MASK                                                        = 0xFFFF0000
	TMC4671_FOC_UWY_SHIFT                                                       = 16
	TMC4671_FOC_UV_MASK                                                         = 0x0000FFFF
	TMC4671_FOC_UV_SHIFT                                                        = 0
	TMC4671_PWM_UX_MASK                                                         = 0x0000FFFF
	TMC4671_PWM_UX_SHIFT                                                        = 0
	TMC4671_PWM_WY_MASK                                                         = 0xFFFF0000
	TMC4671_PWM_WY_SHIFT                                                        = 16
	TMC4671_PWM_V_MASK                                                          = 0x0000FFFF
	TMC4671_PWM_V_SHIFT                                                         = 0
	TMC4671_ADC_I_0_MASK                                                        = 0x0000FFFF
	TMC4671_ADC_I_0_SHIFT                                                       = 0
	TMC4671_ADC_I_1_MASK                                                        = 0xFFFF0000
	TMC4671_ADC_I_1_SHIFT                                                       = 16
	TMC4671_PID_FLUX_ACTUAL_DIV256_MASK                                         = 0x000000FF
	TMC4671_PID_FLUX_ACTUAL_DIV256_SHIFT                                        = 0
	TMC4671_PID_TORQUE_ACTUAL_DIV256_MASK                                       = 0x0000FF00
	TMC4671_PID_TORQUE_ACTUAL_DIV256_SHIFT                                      = 8
	TMC4671_PID_FLUX_TARGET_DIV256_MASK                                         = 0x00FF0000
	TMC4671_PID_FLUX_TARGET_DIV256_SHIFT                                        = 16
	TMC4671_PID_TORQUE_TARGET_DIV256_MASK                                       = 0xFF000000
	TMC4671_PID_TORQUE_TARGET_DIV256_SHIFT                                      = 24
	TMC4671_PID_VELOCITY_ACTUAL_DIV256_MASK                                     = 0x0000FFFF
	TMC4671_PID_VELOCITY_ACTUAL_DIV256_SHIFT                                    = 0
	TMC4671_PID_VELOCITY_TARGET_DIV256_MASK                                     = 0xFFFF0000
	TMC4671_PID_VELOCITY_TARGET_DIV256_SHIFT                                    = 16
	TMC4671_PID_VELOCITY_ACTUAL_LSB_MASK                                        = 0x0000FFFF
	TMC4671_PID_VELOCITY_ACTUAL_LSB_SHIFT                                       = 0
	TMC4671_PID_VELOCITY_TARGET_LSB_MASK                                        = 0xFFFF0000
	TMC4671_PID_VELOCITY_TARGET_LSB_SHIFT                                       = 16
	TMC4671_PID_POSITION_ACTUAL_DIV256_MASK                                     = 0x0000FFFF
	TMC4671_PID_POSITION_ACTUAL_DIV256_SHIFT                                    = 0
	TMC4671_PID_POSITION_TARGET_DIV256_MASK                                     = 0xFFFF0000
	TMC4671_PID_POSITION_TARGET_DIV256_SHIFT                                    = 16
	TMC4671_PID_POSITION_ACTUAL_LSB_MASK                                        = 0x0000FFFF
	TMC4671_PID_POSITION_ACTUAL_LSB_SHIFT                                       = 0
	TMC4671_PID_POSITION_TARGET_LSB_MASK                                        = 0xFFFF0000
	TMC4671_PID_POSITION_TARGET_LSB_SHIFT                                       = 16
	TMC4671_FF_VELOCITY_MASK                                                    = 0xFFFFFFFF
	TMC4671_FF_VELOCITY_SHIFT                                                   = 0
	TMC4671_FF_TORQUE_MASK                                                      = 0x0000FFFF
	TMC4671_FF_TORQUE_SHIFT                                                     = 0
	TMC4671_ACTUAL_VELOCITY_PPTM_MASK                                           = 0xFFFFFFFF
	TMC4671_ACTUAL_VELOCITY_PPTM_SHIFT                                          = 0
	TMC4671_HOME_SWITCH_STATUS_MASK                                             = 0x00000001
	TMC4671_HOME_SWITCH_STATUS_SHIFT                                            = 0
	TMC4671_LEFT_SWITCH_STATUS_MASK                                             = 0x00000002
	TMC4671_LEFT_SWITCH_STATUS_SHIFT                                            = 1
	TMC4671_RIGHT_SWITCH_STATUS_MASK                                            = 0x00000004
	TMC4671_RIGHT_SWITCH_STATUS_SHIFT                                           = 2
	TMC4671_HOME_POSITION_MASK                                                  = 0xFFFFFFFF
	TMC4671_HOME_POSITION_SHIFT                                                 = 0
	TMC4671_LEFT_POSITION_MASK                                                  = 0xFFFFFFFF
	TMC4671_LEFT_POSITION_SHIFT                                                 = 0
	TMC4671_RIGHT_POSITION_MASK                                                 = 0xFFFFFFFF
	TMC4671_RIGHT_POSITION_SHIFT                                                = 0
	TMC4671_ENC_INIT_HALL_STATUS_MASK                                           = 0x0000FFFF
	TMC4671_ENC_INIT_HALL_STATUS_SHIFT                                          = 0
	TMC4671_ENC_INIT_HALL_PHI_E_ABN_OFFSET_MASK                                 = 0x0000FFFF
	TMC4671_ENC_INIT_HALL_PHI_E_ABN_OFFSET_SHIFT                                = 0
	TMC4671_ENC_INIT_HALL_PHI_E_AENC_OFFSET_MASK                                = 0x0000FFFF
	TMC4671_ENC_INIT_HALL_PHI_E_AENC_OFFSET_SHIFT                               = 0
	TMC4671_ENC_INIT_HALL_PHI_A_AENC_OFFSET_MASK                                = 0x0000FFFF
	TMC4671_ENC_INIT_HALL_PHI_A_AENC_OFFSET_SHIFT                               = 0
	TMC4671_ENC_INIT_MINI_MOVE_STATUS_MASK                                      = 0x0000FFFF
	TMC4671_ENC_INIT_MINI_MOVE_STATUS_SHIFT                                     = 0
	TMC4671_ENC_INIT_MINI_MOVE_U_D_MASK                                         = 0xFFFF0000
	TMC4671_ENC_INIT_MINI_MOVE_U_D_SHIFT                                        = 16
	TMC4671_ENC_INIT_MINI_MOVE_PHI_E_OFFSET_MASK                                = 0x0000FFFF
	TMC4671_ENC_INIT_MINI_MOVE_PHI_E_OFFSET_SHIFT                               = 0
	TMC4671_ENC_INIT_MINI_MOVE_PHI_E_MASK                                       = 0xFFFF0000
	TMC4671_ENC_INIT_MINI_MOVE_PHI_E_SHIFT                                      = 16
	TMC4671_SINGLE_PIN_IF_TARGET_TORQUE_MASK                                    = 0x0000FFFF
	TMC4671_SINGLE_PIN_IF_TARGET_TORQUE_SHIFT                                   = 0
	TMC4671_SINGLE_PIN_IF_PWM_DUTY_CYCLE_MASK                                   = 0xFFFF0000
	TMC4671_SINGLE_PIN_IF_PWM_DUTY_CYCLE_SHIFT                                  = 16
	TMC4671_SINGLE_PIN_IF_TARGET_VELOCITY_MASK                                  = 0xFFFFFFFF
	TMC4671_SINGLE_PIN_IF_TARGET_VELOCITY_SHIFT                                 = 0
	TMC4671_SINGLE_PIN_IF_TARGET_POSITION_MASK                                  = 0xFFFFFFFF
	TMC4671_SINGLE_PIN_IF_TARGET_POSITION_SHIFT                                 = 0
	TMC4671_INTERIM_ADDR_MASK                                                   = 0x000000FF
	TMC4671_INTERIM_ADDR_SHIFT                                                  = 0
	TMC4671_WATCHDOG_CFG_MASK                                                   = 0x00000003
	TMC4671_WATCHDOG_CFG_SHIFT                                                  = 0
	TMC4671_ADC_VM_LIMIT_LOW_MASK                                               = 0x0000FFFF
	TMC4671_ADC_VM_LIMIT_LOW_SHIFT                                              = 0
	TMC4671_ADC_VM_LIMIT_HIGH_MASK                                              = 0xFFFF0000
	TMC4671_ADC_VM_LIMIT_HIGH_SHIFT                                             = 16
	TMC4671_A_OF_ABN_RAW_MASK                                                   = 0x00000001
	TMC4671_A_OF_ABN_RAW_SHIFT                                                  = 0
	TMC4671_B_OF_ABN_RAW_MASK                                                   = 0x00000002
	TMC4671_B_OF_ABN_RAW_SHIFT                                                  = 1
	TMC4671_N_OF_ABN_RAW_MASK                                                   = 0x00000004
	TMC4671_N_OF_ABN_RAW_SHIFT                                                  = 2
	TMC4671_A_OF_ABN_2_RAW_MASK                                                 = 0x00000010
	TMC4671_A_OF_ABN_2_RAW_SHIFT                                                = 4
	TMC4671_B_OF_ABN_2_RAW_MASK                                                 = 0x00000020
	TMC4671_B_OF_ABN_2_RAW_SHIFT                                                = 5
	TMC4671_N_OF_ABN_2_RAW_MASK                                                 = 0x00000040
	TMC4671_N_OF_ABN_2_RAW_SHIFT                                                = 6
	TMC4671_HALL_UX_OF_HALL_RAW_MASK                                            = 0x00000100
	TMC4671_HALL_UX_OF_HALL_RAW_SHIFT                                           = 8
	TMC4671_HALL_V_OF_HALL_RAW_MASK                                             = 0x00000200
	TMC4671_HALL_V_OF_HALL_RAW_SHIFT                                            = 9
	TMC4671_HALL_WY_OF_HALL_RAW_MASK                                            = 0x00000400
	TMC4671_HALL_WY_OF_HALL_RAW_SHIFT                                           = 10
	TMC4671_REF_SW_R_RAW_MASK                                                   = 0x00001000
	TMC4671_REF_SW_R_RAW_SHIFT                                                  = 12
	TMC4671_REF_SW_L_RAW_MASK                                                   = 0x00004000
	TMC4671_REF_SW_L_RAW_SHIFT                                                  = 14
	TMC4671_ENABLE_IN_RAW_MASK                                                  = 0x00008000
	TMC4671_ENABLE_IN_RAW_SHIFT                                                 = 15
	TMC4671_STP_OF_DIRSTP_RAW_MASK                                              = 0x00010000
	TMC4671_STP_OF_DIRSTP_RAW_SHIFT                                             = 16
	TMC4671_DIR_OF_DIRSTP_RAW_MASK                                              = 0x00020000
	TMC4671_DIR_OF_DIRSTP_RAW_SHIFT                                             = 17
	TMC4671_PWM_IN_RAW_MASK                                                     = 0x00040000
	TMC4671_PWM_IN_RAW_SHIFT                                                    = 18
	TMC4671_HALL_UX_FILT_MASK                                                   = 0x00100000
	TMC4671_HALL_UX_FILT_SHIFT                                                  = 20
	TMC4671_HALL_V_FILT_MASK                                                    = 0x00200000
	TMC4671_HALL_V_FILT_SHIFT                                                   = 21
	TMC4671_HALL_WY_FILT_MASK                                                   = 0x00400000
	TMC4671_HALL_WY_FILT_SHIFT                                                  = 22
	TMC4671_PWM_IDLE_L_RAW_MASK                                                 = 0x10000000
	TMC4671_PWM_IDLE_L_RAW_SHIFT                                                = 28
	TMC4671_OUTPUTS_RAW_0_MASK                                                  = 0x00000001
	TMC4671_OUTPUTS_RAW_0_SHIFT                                                 = 0
	TMC4671_OUTPUTS_RAW_1_MASK                                                  = 0x00000002
	TMC4671_OUTPUTS_RAW_1_SHIFT                                                 = 1
	TMC4671_OUTPUTS_RAW_2_MASK                                                  = 0x00000004
	TMC4671_OUTPUTS_RAW_2_SHIFT                                                 = 2
	TMC4671_OUTPUTS_RAW_3_MASK                                                  = 0x00000008
	TMC4671_OUTPUTS_RAW_3_SHIFT                                                 = 3
	TMC4671_OUTPUTS_RAW_4_MASK                                                  = 0x00000010
	TMC4671_OUTPUTS_RAW_4_SHIFT                                                 = 4
	TMC4671_OUTPUTS_RAW_5_MASK                                                  = 0x00000020
	TMC4671_OUTPUTS_RAW_5_SHIFT                                                 = 5
	TMC4671_OUTPUTS_RAW_6_MASK                                                  = 0x00000040
	TMC4671_OUTPUTS_RAW_6_SHIFT                                                 = 6
	TMC4671_OUTPUTS_RAW_7_MASK                                                  = 0x00000080
	TMC4671_OUTPUTS_RAW_7_SHIFT                                                 = 7
	TMC4671_STEP_WIDTH_MASK                                                     = 0xFFFFFFFF
	TMC4671_STEP_WIDTH_SHIFT                                                    = 0
	TMC4671_UART_BPS_MASK                                                       = 0x03FFFFFF
	TMC4671_UART_BPS_SHIFT                                                      = 0
	TMC4671_ADDR_A_MASK                                                         = 0x000000FF
	TMC4671_ADDR_A_SHIFT                                                        = 0
	TMC4671_ADDR_B_MASK                                                         = 0x0000FF00
	TMC4671_ADDR_B_SHIFT                                                        = 8
	TMC4671_ADDR_C_MASK                                                         = 0x00FF0000
	TMC4671_ADDR_C_SHIFT                                                        = 16
	TMC4671_ADDR_D_MASK                                                         = 0xFF000000
	TMC4671_ADDR_D_SHIFT                                                        = 24
	TMC4671_GPIO_DSADCI_CONFIG_0_MASK                                           = 0x00000001
	TMC4671_GPIO_DSADCI_CONFIG_0_SHIFT                                          = 0
	TMC4671_GPIO_DSADCI_CONFIG_1_MASK                                           = 0x00000002
	TMC4671_GPIO_DSADCI_CONFIG_1_SHIFT                                          = 1
	TMC4671_GPIO_DSADCI_CONFIG_2_MASK                                           = 0x00000004
	TMC4671_GPIO_DSADCI_CONFIG_2_SHIFT                                          = 2
	TMC4671_GPIO_DSADCI_CONFIG_3_MASK                                           = 0x00000008
	TMC4671_GPIO_DSADCI_CONFIG_3_SHIFT                                          = 3
	TMC4671_GPIO_DSADCI_CONFIG_4_MASK                                           = 0x00000010
	TMC4671_GPIO_DSADCI_CONFIG_4_SHIFT                                          = 4
	TMC4671_GPIO_DSADCI_CONFIG_5_MASK                                           = 0x00000020
	TMC4671_GPIO_DSADCI_CONFIG_5_SHIFT                                          = 5
	TMC4671_GPIO_DSADCI_CONFIG_6_MASK                                           = 0x00000040
	TMC4671_GPIO_DSADCI_CONFIG_6_SHIFT                                          = 6
	TMC4671_GPO_MASK                                                            = 0x00FF0000
	TMC4671_GPO_SHIFT                                                           = 16
	TMC4671_GPI_MASK                                                            = 0xFF000000
	TMC4671_GPI_SHIFT                                                           = 24
	TMC4671_STATUS_FLAGS_0_MASK                                                 = 0x00000001
	TMC4671_STATUS_FLAGS_0_SHIFT                                                = 0
	TMC4671_STATUS_FLAGS_1_MASK                                                 = 0x00000002
	TMC4671_STATUS_FLAGS_1_SHIFT                                                = 1
	TMC4671_STATUS_FLAGS_2_MASK                                                 = 0x00000004
	TMC4671_STATUS_FLAGS_2_SHIFT                                                = 2
	TMC4671_STATUS_FLAGS_3_MASK                                                 = 0x00000008
	TMC4671_STATUS_FLAGS_3_SHIFT                                                = 3
	TMC4671_STATUS_FLAGS_4_MASK                                                 = 0x00000010
	TMC4671_STATUS_FLAGS_4_SHIFT                                                = 4
	TMC4671_STATUS_FLAGS_5_MASK                                                 = 0x00000020
	TMC4671_STATUS_FLAGS_5_SHIFT                                                = 5
	TMC4671_STATUS_FLAGS_6_MASK                                                 = 0x00000040
	TMC4671_STATUS_FLAGS_6_SHIFT                                                = 6
	TMC4671_STATUS_FLAGS_7_MASK                                                 = 0x00000080
	TMC4671_STATUS_FLAGS_7_SHIFT                                                = 7
	TMC4671_STATUS_FLAGS_8_MASK                                                 = 0x00000100
	TMC4671_STATUS_FLAGS_8_SHIFT                                                = 8
	TMC4671_STATUS_FLAGS_9_MASK                                                 = 0x00000200
	TMC4671_STATUS_FLAGS_9_SHIFT                                                = 9
	TMC4671_STATUS_FLAGS_10_MASK                                                = 0x00000400
	TMC4671_STATUS_FLAGS_10_SHIFT                                               = 10
	TMC4671_STATUS_FLAGS_11_MASK                                                = 0x00000800
	TMC4671_STATUS_FLAGS_11_SHIFT                                               = 11
	TMC4671_STATUS_FLAGS_12_MASK                                                = 0x00001000
	TMC4671_STATUS_FLAGS_12_SHIFT                                               = 12
	TMC4671_STATUS_FLAGS_13_MASK                                                = 0x00002000
	TMC4671_STATUS_FLAGS_13_SHIFT                                               = 13
	TMC4671_STATUS_FLAGS_14_MASK                                                = 0x00004000
	TMC4671_STATUS_FLAGS_14_SHIFT                                               = 14
	TMC4671_STATUS_FLAGS_15_MASK                                                = 0x00008000
	TMC4671_STATUS_FLAGS_15_SHIFT                                               = 15
	TMC4671_STATUS_FLAGS_16_MASK                                                = 0x00010000
	TMC4671_STATUS_FLAGS_16_SHIFT                                               = 16
	TMC4671_STATUS_FLAGS_17_MASK                                                = 0x00020000
	TMC4671_STATUS_FLAGS_17_SHIFT                                               = 17
	TMC4671_STATUS_FLAGS_18_MASK                                                = 0x00040000
	TMC4671_STATUS_FLAGS_18_SHIFT                                               = 18
	TMC4671_STATUS_FLAGS_19_MASK                                                = 0x00080000
	TMC4671_STATUS_FLAGS_19_SHIFT                                               = 19
	TMC4671_STATUS_FLAGS_20_MASK                                                = 0x00100000
	TMC4671_STATUS_FLAGS_20_SHIFT                                               = 20
	TMC4671_STATUS_FLAGS_21_MASK                                                = 0x00200000
	TMC4671_STATUS_FLAGS_21_SHIFT                                               = 21
	TMC4671_STATUS_FLAGS_22_MASK                                                = 0x00400000
	TMC4671_STATUS_FLAGS_22_SHIFT                                               = 22
	TMC4671_STATUS_FLAGS_23_MASK                                                = 0x00800000
	TMC4671_STATUS_FLAGS_23_SHIFT                                               = 23
	TMC4671_STATUS_FLAGS_24_MASK                                                = 0x01000000
	TMC4671_STATUS_FLAGS_24_SHIFT                                               = 24
	TMC4671_STATUS_FLAGS_25_MASK                                                = 0x02000000
	TMC4671_STATUS_FLAGS_25_SHIFT                                               = 25
	TMC4671_STATUS_FLAGS_26_MASK                                                = 0x04000000
	TMC4671_STATUS_FLAGS_26_SHIFT                                               = 26
	TMC4671_STATUS_FLAGS_27_MASK                                                = 0x08000000
	TMC4671_STATUS_FLAGS_27_SHIFT                                               = 27
	TMC4671_STATUS_FLAGS_28_MASK                                                = 0x10000000
	TMC4671_STATUS_FLAGS_28_SHIFT                                               = 28
	TMC4671_STATUS_FLAGS_29_MASK                                                = 0x20000000
	TMC4671_STATUS_FLAGS_29_SHIFT                                               = 29
	TMC4671_STATUS_FLAGS_30_MASK                                                = 0x40000000
	TMC4671_STATUS_FLAGS_30_SHIFT                                               = 30
	TMC4671_STATUS_FLAGS_31_MASK                                                = 0x80000000
	TMC4671_STATUS_FLAGS_31_SHIFT                                               = 31
	TMC4671_WARNING_MASK_MASK                                                   = 0xFFFFFFFF
	TMC4671_WARNING_MASK_SHIFT                                                  = 0
)

// Register Fields
var (
	TMC4671_SI_TYPE_FIELD = RegisterField{
		Mask:     TMC4671_SI_TYPE_MASK,
		Shift:    TMC4671_SI_TYPE_SHIFT,
		Register: TMC4671_CHIPINFO_DATA,
		IsSigned: false,
	}
	TMC4671_SI_VERSION_FIELD = RegisterField{
		Mask:     TMC4671_SI_VERSION_MASK,
		Shift:    TMC4671_SI_VERSION_SHIFT,
		Register: TMC4671_CHIPINFO_DATA,
		IsSigned: false,
	}
	TMC4671_SI_DATE_FIELD = RegisterField{
		Mask:     TMC4671_SI_DATE_MASK,
		Shift:    TMC4671_SI_DATE_SHIFT,
		Register: TMC4671_CHIPINFO_DATA,
		IsSigned: false,
	}
	TMC4671_SI_TIME_FIELD = RegisterField{
		Mask:     TMC4671_SI_TIME_MASK,
		Shift:    TMC4671_SI_TIME_SHIFT,
		Register: TMC4671_CHIPINFO_DATA,
		IsSigned: false,
	}
	TMC4671_SI_VARIANT_FIELD = RegisterField{
		Mask:     TMC4671_SI_VARIANT_MASK,
		Shift:    TMC4671_SI_VARIANT_SHIFT,
		Register: TMC4671_CHIPINFO_DATA,
		IsSigned: false,
	}
	TMC4671_SI_BUILD_FIELD = RegisterField{
		Mask:     TMC4671_SI_BUILD_MASK,
		Shift:    TMC4671_SI_BUILD_SHIFT,
		Register: TMC4671_CHIPINFO_DATA,
		IsSigned: false,
	}
	TMC4671_CHIP_INFO_ADDRESS_FIELD = RegisterField{
		Mask:     TMC4671_CHIP_INFO_ADDRESS_MASK,
		Shift:    TMC4671_CHIP_INFO_ADDRESS_SHIFT,
		Register: TMC4671_CHIPINFO_ADDR,
		IsSigned: false,
	}
	TMC4671_ADC_I0_RAW_FIELD = RegisterField{
		Mask:     TMC4671_ADC_I0_RAW_MASK,
		Shift:    TMC4671_ADC_I0_RAW_SHIFT,
		Register: TMC4671_ADC_RAW_DATA,
		IsSigned: false,
	}
	TMC4671_ADC_I1_RAW_FIELD = RegisterField{
		Mask:     TMC4671_ADC_I1_RAW_MASK,
		Shift:    TMC4671_ADC_I1_RAW_SHIFT,
		Register: TMC4671_ADC_RAW_DATA,
		IsSigned: false,
	}
	TMC4671_ADC_VM_RAW_FIELD = RegisterField{
		Mask:     TMC4671_ADC_VM_RAW_MASK,
		Shift:    TMC4671_ADC_VM_RAW_SHIFT,
		Register: TMC4671_ADC_RAW_DATA,
		IsSigned: false,
	}
	TMC4671_ADC_AGPI_A_RAW_FIELD = RegisterField{
		Mask:     TMC4671_ADC_AGPI_A_RAW_MASK,
		Shift:    TMC4671_ADC_AGPI_A_RAW_SHIFT,
		Register: TMC4671_ADC_RAW_DATA,
		IsSigned: false,
	}
	TMC4671_ADC_AGPI_B_RAW_FIELD = RegisterField{
		Mask:     TMC4671_ADC_AGPI_B_RAW_MASK,
		Shift:    TMC4671_ADC_AGPI_B_RAW_SHIFT,
		Register: TMC4671_ADC_RAW_DATA,
		IsSigned: false,
	}
	TMC4671_ADC_AENC_UX_RAW_FIELD = RegisterField{
		Mask:     TMC4671_ADC_AENC_UX_RAW_MASK,
		Shift:    TMC4671_ADC_AENC_UX_RAW_SHIFT,
		Register: TMC4671_ADC_RAW_DATA,
		IsSigned: false,
	}
	TMC4671_ADC_AENC_VN_RAW_FIELD = RegisterField{
		Mask:     TMC4671_ADC_AENC_VN_RAW_MASK,
		Shift:    TMC4671_ADC_AENC_VN_RAW_SHIFT,
		Register: TMC4671_ADC_RAW_DATA,
		IsSigned: false,
	}
	TMC4671_ADC_AENC_WY_RAW_FIELD = RegisterField{
		Mask:     TMC4671_ADC_AENC_WY_RAW_MASK,
		Shift:    TMC4671_ADC_AENC_WY_RAW_SHIFT,
		Register: TMC4671_ADC_RAW_DATA,
		IsSigned: false,
	}
	TMC4671_ADC_RAW_ADDR_FIELD = RegisterField{
		Mask:     TMC4671_ADC_RAW_ADDR_MASK,
		Shift:    TMC4671_ADC_RAW_ADDR_SHIFT,
		Register: TMC4671_ADC_RAW_ADDR,
		IsSigned: false,
	}
	TMC4671_CFG_DSMODULATOR_A_FIELD = RegisterField{
		Mask:     TMC4671_CFG_DSMODULATOR_A_MASK,
		Shift:    TMC4671_CFG_DSMODULATOR_A_SHIFT,
		Register: TMC4671_DSADC_MCFG_B_MCFG_A,
		IsSigned: false,
	}
	TMC4671_MCLK_POLARITY_A_FIELD = RegisterField{
		Mask:     TMC4671_MCLK_POLARITY_A_MASK,
		Shift:    TMC4671_MCLK_POLARITY_A_SHIFT,
		Register: TMC4671_DSADC_MCFG_B_MCFG_A,
		IsSigned: false,
	}
	TMC4671_MDAT_POLARITY_A_FIELD = RegisterField{
		Mask:     TMC4671_MDAT_POLARITY_A_MASK,
		Shift:    TMC4671_MDAT_POLARITY_A_SHIFT,
		Register: TMC4671_DSADC_MCFG_B_MCFG_A,
		IsSigned: false,
	}
	TMC4671_SEL_NCLK_MCLK_I_A_FIELD = RegisterField{
		Mask:     TMC4671_SEL_NCLK_MCLK_I_A_MASK,
		Shift:    TMC4671_SEL_NCLK_MCLK_I_A_SHIFT,
		Register: TMC4671_DSADC_MCFG_B_MCFG_A,
		IsSigned: false,
	}
	TMC4671_BLANKING_A_FIELD = RegisterField{
		Mask:     TMC4671_BLANKING_A_MASK,
		Shift:    TMC4671_BLANKING_A_SHIFT,
		Register: TMC4671_DSADC_MCFG_B_MCFG_A,
		IsSigned: false,
	}
	TMC4671_CFG_DSMODULATOR_B_FIELD = RegisterField{
		Mask:     TMC4671_CFG_DSMODULATOR_B_MASK,
		Shift:    TMC4671_CFG_DSMODULATOR_B_SHIFT,
		Register: TMC4671_DSADC_MCFG_B_MCFG_A,
		IsSigned: false,
	}
	TMC4671_MCLK_POLARITY_B_FIELD = RegisterField{
		Mask:     TMC4671_MCLK_POLARITY_B_MASK,
		Shift:    TMC4671_MCLK_POLARITY_B_SHIFT,
		Register: TMC4671_DSADC_MCFG_B_MCFG_A,
		IsSigned: false,
	}
	TMC4671_MDAT_POLARITY_B_FIELD = RegisterField{
		Mask:     TMC4671_MDAT_POLARITY_B_MASK,
		Shift:    TMC4671_MDAT_POLARITY_B_SHIFT,
		Register: TMC4671_DSADC_MCFG_B_MCFG_A,
		IsSigned: false,
	}
	TMC4671_SEL_NCLK_MCLK_I_B_FIELD = RegisterField{
		Mask:     TMC4671_SEL_NCLK_MCLK_I_B_MASK,
		Shift:    TMC4671_SEL_NCLK_MCLK_I_B_SHIFT,
		Register: TMC4671_DSADC_MCFG_B_MCFG_A,
		IsSigned: false,
	}
	TMC4671_BLANKING_B_FIELD = RegisterField{
		Mask:     TMC4671_BLANKING_B_MASK,
		Shift:    TMC4671_BLANKING_B_SHIFT,
		Register: TMC4671_DSADC_MCFG_B_MCFG_A,
		IsSigned: false,
	}
	TMC4671_DSADC_MCLK_A_FIELD = RegisterField{
		Mask:     TMC4671_DSADC_MCLK_A_MASK,
		Shift:    TMC4671_DSADC_MCLK_A_SHIFT,
		Register: TMC4671_DSADC_MCLK_A,
		IsSigned: false,
	}
	TMC4671_DSADC_MCLK_B_FIELD = RegisterField{
		Mask:     TMC4671_DSADC_MCLK_B_MASK,
		Shift:    TMC4671_DSADC_MCLK_B_SHIFT,
		Register: TMC4671_DSADC_MCLK_B,
		IsSigned: false,
	}
	TMC4671_DSADC_MDEC_A_FIELD = RegisterField{
		Mask:     TMC4671_DSADC_MDEC_A_MASK,
		Shift:    TMC4671_DSADC_MDEC_A_SHIFT,
		Register: TMC4671_DSADC_MDEC_B_MDEC_A,
		IsSigned: false,
	}
	TMC4671_DSADC_MDEC_B_FIELD = RegisterField{
		Mask:     TMC4671_DSADC_MDEC_B_MASK,
		Shift:    TMC4671_DSADC_MDEC_B_SHIFT,
		Register: TMC4671_DSADC_MDEC_B_MDEC_A,
		IsSigned: false,
	}
	TMC4671_ADC_I1_OFFSET_FIELD = RegisterField{
		Mask:     TMC4671_ADC_I1_OFFSET_MASK,
		Shift:    TMC4671_ADC_I1_OFFSET_SHIFT,
		Register: TMC4671_ADC_I1_SCALE_OFFSET,
		IsSigned: false,
	}
	TMC4671_ADC_I1_SCALE_FIELD = RegisterField{
		Mask:     TMC4671_ADC_I1_SCALE_MASK,
		Shift:    TMC4671_ADC_I1_SCALE_SHIFT,
		Register: TMC4671_ADC_I1_SCALE_OFFSET,
		IsSigned: true,
	}
	TMC4671_ADC_I0_OFFSET_FIELD = RegisterField{
		Mask:     TMC4671_ADC_I0_OFFSET_MASK,
		Shift:    TMC4671_ADC_I0_OFFSET_SHIFT,
		Register: TMC4671_ADC_I0_SCALE_OFFSET,
		IsSigned: false,
	}
	TMC4671_ADC_I0_SCALE_FIELD = RegisterField{
		Mask:     TMC4671_ADC_I0_SCALE_MASK,
		Shift:    TMC4671_ADC_I0_SCALE_SHIFT,
		Register: TMC4671_ADC_I0_SCALE_OFFSET,
		IsSigned: true,
	}
	TMC4671_ADC_I0_SELECT_FIELD = RegisterField{
		Mask:     TMC4671_ADC_I0_SELECT_MASK,
		Shift:    TMC4671_ADC_I0_SELECT_SHIFT,
		Register: TMC4671_ADC_I_SELECT,
		IsSigned: false,
	}
	TMC4671_ADC_I1_SELECT_FIELD = RegisterField{
		Mask:     TMC4671_ADC_I1_SELECT_MASK,
		Shift:    TMC4671_ADC_I1_SELECT_SHIFT,
		Register: TMC4671_ADC_I_SELECT,
		IsSigned: false,
	}
	TMC4671_ADC_I_UX_SELECT_FIELD = RegisterField{
		Mask:     TMC4671_ADC_I_UX_SELECT_MASK,
		Shift:    TMC4671_ADC_I_UX_SELECT_SHIFT,
		Register: TMC4671_ADC_I_SELECT,
		IsSigned: false,
	}
	TMC4671_ADC_I_V_SELECT_FIELD = RegisterField{
		Mask:     TMC4671_ADC_I_V_SELECT_MASK,
		Shift:    TMC4671_ADC_I_V_SELECT_SHIFT,
		Register: TMC4671_ADC_I_SELECT,
		IsSigned: false,
	}
	TMC4671_ADC_I_WY_SELECT_FIELD = RegisterField{
		Mask:     TMC4671_ADC_I_WY_SELECT_MASK,
		Shift:    TMC4671_ADC_I_WY_SELECT_SHIFT,
		Register: TMC4671_ADC_I_SELECT,
		IsSigned: false,
	}
	TMC4671_ADC_I0_EXT_FIELD = RegisterField{
		Mask:     TMC4671_ADC_I0_EXT_MASK,
		Shift:    TMC4671_ADC_I0_EXT_SHIFT,
		Register: TMC4671_ADC_I1_I0_EXT,
		IsSigned: false,
	}
	TMC4671_ADC_I1_EXT_FIELD = RegisterField{
		Mask:     TMC4671_ADC_I1_EXT_MASK,
		Shift:    TMC4671_ADC_I1_EXT_SHIFT,
		Register: TMC4671_ADC_I1_I0_EXT,
		IsSigned: false,
	}
	TMC4671_ADC_I0_FIELD = RegisterField{
		Mask:     TMC4671_ADC_I0_MASK,
		Shift:    TMC4671_ADC_I0_SHIFT,
		Register: TMC4671_DS_ANALOG_INPUT_STAGE_CFG,
		IsSigned: false,
	}
	TMC4671_ADC_I1_FIELD = RegisterField{
		Mask:     TMC4671_ADC_I1_MASK,
		Shift:    TMC4671_ADC_I1_SHIFT,
		Register: TMC4671_DS_ANALOG_INPUT_STAGE_CFG,
		IsSigned: false,
	}
	TMC4671_ADC_VM_FIELD = RegisterField{
		Mask:     TMC4671_ADC_VM_MASK,
		Shift:    TMC4671_ADC_VM_SHIFT,
		Register: TMC4671_DS_ANALOG_INPUT_STAGE_CFG,
		IsSigned: false,
	}
	TMC4671_ADC_AGPI_A_FIELD = RegisterField{
		Mask:     TMC4671_ADC_AGPI_A_MASK,
		Shift:    TMC4671_ADC_AGPI_A_SHIFT,
		Register: TMC4671_DS_ANALOG_INPUT_STAGE_CFG,
		IsSigned: false,
	}
	TMC4671_ADC_AGPI_B_FIELD = RegisterField{
		Mask:     TMC4671_ADC_AGPI_B_MASK,
		Shift:    TMC4671_ADC_AGPI_B_SHIFT,
		Register: TMC4671_DS_ANALOG_INPUT_STAGE_CFG,
		IsSigned: false,
	}
	TMC4671_ADC_AENC_UX_FIELD = RegisterField{
		Mask:     TMC4671_ADC_AENC_UX_MASK,
		Shift:    TMC4671_ADC_AENC_UX_SHIFT,
		Register: TMC4671_DS_ANALOG_INPUT_STAGE_CFG,
		IsSigned: false,
	}
	TMC4671_ADC_AENC_VN_FIELD = RegisterField{
		Mask:     TMC4671_ADC_AENC_VN_MASK,
		Shift:    TMC4671_ADC_AENC_VN_SHIFT,
		Register: TMC4671_DS_ANALOG_INPUT_STAGE_CFG,
		IsSigned: false,
	}
	TMC4671_ADC_AENC_WY_FIELD = RegisterField{
		Mask:     TMC4671_ADC_AENC_WY_MASK,
		Shift:    TMC4671_ADC_AENC_WY_SHIFT,
		Register: TMC4671_DS_ANALOG_INPUT_STAGE_CFG,
		IsSigned: false,
	}
	TMC4671_AENC_0_OFFSET_FIELD = RegisterField{
		Mask:     TMC4671_AENC_0_OFFSET_MASK,
		Shift:    TMC4671_AENC_0_OFFSET_SHIFT,
		Register: TMC4671_AENC_0_SCALE_OFFSET,
		IsSigned: false,
	}
	TMC4671_AENC_0_SCALE_FIELD = RegisterField{
		Mask:     TMC4671_AENC_0_SCALE_MASK,
		Shift:    TMC4671_AENC_0_SCALE_SHIFT,
		Register: TMC4671_AENC_0_SCALE_OFFSET,
		IsSigned: true,
	}
	TMC4671_AENC_1_OFFSET_FIELD = RegisterField{
		Mask:     TMC4671_AENC_1_OFFSET_MASK,
		Shift:    TMC4671_AENC_1_OFFSET_SHIFT,
		Register: TMC4671_AENC_1_SCALE_OFFSET,
		IsSigned: false,
	}
	TMC4671_AENC_1_SCALE_FIELD = RegisterField{
		Mask:     TMC4671_AENC_1_SCALE_MASK,
		Shift:    TMC4671_AENC_1_SCALE_SHIFT,
		Register: TMC4671_AENC_1_SCALE_OFFSET,
		IsSigned: true,
	}
	TMC4671_AENC_2_OFFSET_FIELD = RegisterField{
		Mask:     TMC4671_AENC_2_OFFSET_MASK,
		Shift:    TMC4671_AENC_2_OFFSET_SHIFT,
		Register: TMC4671_AENC_2_SCALE_OFFSET,
		IsSigned: false,
	}
	TMC4671_AENC_2_SCALE_FIELD = RegisterField{
		Mask:     TMC4671_AENC_2_SCALE_MASK,
		Shift:    TMC4671_AENC_2_SCALE_SHIFT,
		Register: TMC4671_AENC_2_SCALE_OFFSET,
		IsSigned: true,
	}
	TMC4671_AENC_0_SELECT_FIELD = RegisterField{
		Mask:     TMC4671_AENC_0_SELECT_MASK,
		Shift:    TMC4671_AENC_0_SELECT_SHIFT,
		Register: TMC4671_AENC_SELECT,
		IsSigned: false,
	}
	TMC4671_AENC_1_SELECT_FIELD = RegisterField{
		Mask:     TMC4671_AENC_1_SELECT_MASK,
		Shift:    TMC4671_AENC_1_SELECT_SHIFT,
		Register: TMC4671_AENC_SELECT,
		IsSigned: false,
	}
	TMC4671_AENC_2_SELECT_FIELD = RegisterField{
		Mask:     TMC4671_AENC_2_SELECT_MASK,
		Shift:    TMC4671_AENC_2_SELECT_SHIFT,
		Register: TMC4671_AENC_SELECT,
		IsSigned: false,
	}
	TMC4671_ADC_IUX_FIELD = RegisterField{
		Mask:     TMC4671_ADC_IUX_MASK,
		Shift:    TMC4671_ADC_IUX_SHIFT,
		Register: TMC4671_ADC_IWY_IUX,
		IsSigned: true,
	}
	TMC4671_ADC_IWY_FIELD = RegisterField{
		Mask:     TMC4671_ADC_IWY_MASK,
		Shift:    TMC4671_ADC_IWY_SHIFT,
		Register: TMC4671_ADC_IWY_IUX,
		IsSigned: true,
	}
	TMC4671_ADC_IV_FIELD = RegisterField{
		Mask:     TMC4671_ADC_IV_MASK,
		Shift:    TMC4671_ADC_IV_SHIFT,
		Register: TMC4671_ADC_IV,
		IsSigned: true,
	}
	TMC4671_AENC_UX_FIELD = RegisterField{
		Mask:     TMC4671_AENC_UX_MASK,
		Shift:    TMC4671_AENC_UX_SHIFT,
		Register: TMC4671_AENC_WY_UX,
		IsSigned: true,
	}
	TMC4671_AENC_WY_FIELD = RegisterField{
		Mask:     TMC4671_AENC_WY_MASK,
		Shift:    TMC4671_AENC_WY_SHIFT,
		Register: TMC4671_AENC_WY_UX,
		IsSigned: true,
	}
	TMC4671_AENC_VN_FIELD = RegisterField{
		Mask:     TMC4671_AENC_VN_MASK,
		Shift:    TMC4671_AENC_VN_SHIFT,
		Register: TMC4671_AENC_VN,
		IsSigned: true,
	}
	TMC4671_PWM_POLARITIES_0_FIELD = RegisterField{
		Mask:     TMC4671_PWM_POLARITIES_0_MASK,
		Shift:    TMC4671_PWM_POLARITIES_0_SHIFT,
		Register: TMC4671_PWM_POLARITIES,
		IsSigned: false,
	}
	TMC4671_PWM_POLARITIES_1_FIELD = RegisterField{
		Mask:     TMC4671_PWM_POLARITIES_1_MASK,
		Shift:    TMC4671_PWM_POLARITIES_1_SHIFT,
		Register: TMC4671_PWM_POLARITIES,
		IsSigned: false,
	}
	TMC4671_PWM_MAXCNT_FIELD = RegisterField{
		Mask:     TMC4671_PWM_MAXCNT_MASK,
		Shift:    TMC4671_PWM_MAXCNT_SHIFT,
		Register: TMC4671_PWM_MAXCNT,
		IsSigned: false,
	}
	TMC4671_PWM_CHOP_FIELD = RegisterField{
		Mask:     TMC4671_PWM_CHOP_MASK,
		Shift:    TMC4671_PWM_CHOP_SHIFT,
		Register: TMC4671_PWM_SV_CHOP,
		IsSigned: false,
	}
	TMC4671_PWM_SV_FIELD = RegisterField{
		Mask:     TMC4671_PWM_SV_MASK,
		Shift:    TMC4671_PWM_SV_SHIFT,
		Register: TMC4671_PWM_SV_CHOP,
		IsSigned: false,
	}
	TMC4671_N_POLE_PAIRS_FIELD = RegisterField{
		Mask:     TMC4671_N_POLE_PAIRS_MASK,
		Shift:    TMC4671_N_POLE_PAIRS_SHIFT,
		Register: TMC4671_MOTOR_TYPE_N_POLE_PAIRS,
		IsSigned: false,
	}
	TMC4671_MOTOR_TYPE_FIELD = RegisterField{
		Mask:     TMC4671_MOTOR_TYPE_MASK,
		Shift:    TMC4671_MOTOR_TYPE_SHIFT,
		Register: TMC4671_MOTOR_TYPE_N_POLE_PAIRS,
		IsSigned: false,
	}
	TMC4671_PHI_E_EXT_FIELD = RegisterField{
		Mask:     TMC4671_PHI_E_EXT_MASK,
		Shift:    TMC4671_PHI_E_EXT_SHIFT,
		Register: TMC4671_PHI_E_EXT,
		IsSigned: true,
	}
	TMC4671_OPENLOOP_PHI_DIRECTION_FIELD = RegisterField{
		Mask:     TMC4671_OPENLOOP_PHI_DIRECTION_MASK,
		Shift:    TMC4671_OPENLOOP_PHI_DIRECTION_SHIFT,
		Register: TMC4671_OPENLOOP_MODE,
		IsSigned: false,
	}
	TMC4671_OPENLOOP_ACCELERATION_FIELD = RegisterField{
		Mask:     TMC4671_OPENLOOP_ACCELERATION_MASK,
		Shift:    TMC4671_OPENLOOP_ACCELERATION_SHIFT,
		Register: TMC4671_OPENLOOP_ACCELERATION,
		IsSigned: false,
	}
	TMC4671_OPENLOOP_VELOCITY_TARGET_FIELD = RegisterField{
		Mask:     TMC4671_OPENLOOP_VELOCITY_TARGET_MASK,
		Shift:    TMC4671_OPENLOOP_VELOCITY_TARGET_SHIFT,
		Register: TMC4671_OPENLOOP_VELOCITY_TARGET,
		IsSigned: true,
	}
	TMC4671_OPENLOOP_VELOCITY_ACTUAL_FIELD = RegisterField{
		Mask:     TMC4671_OPENLOOP_VELOCITY_ACTUAL_MASK,
		Shift:    TMC4671_OPENLOOP_VELOCITY_ACTUAL_SHIFT,
		Register: TMC4671_OPENLOOP_VELOCITY_ACTUAL,
		IsSigned: true,
	}
	TMC4671_OPENLOOP_PHI_FIELD = RegisterField{
		Mask:     TMC4671_OPENLOOP_PHI_MASK,
		Shift:    TMC4671_OPENLOOP_PHI_SHIFT,
		Register: TMC4671_OPENLOOP_PHI,
		IsSigned: true,
	}
	TMC4671_UD_EXT_FIELD = RegisterField{
		Mask:     TMC4671_UD_EXT_MASK,
		Shift:    TMC4671_UD_EXT_SHIFT,
		Register: TMC4671_UQ_UD_EXT,
		IsSigned: true,
	}
	TMC4671_UQ_EXT_FIELD = RegisterField{
		Mask:     TMC4671_UQ_EXT_MASK,
		Shift:    TMC4671_UQ_EXT_SHIFT,
		Register: TMC4671_UQ_UD_EXT,
		IsSigned: true,
	}
	TMC4671_ABN_APOL_FIELD = RegisterField{
		Mask:     TMC4671_ABN_APOL_MASK,
		Shift:    TMC4671_ABN_APOL_SHIFT,
		Register: TMC4671_ABN_DECODER_MODE,
		IsSigned: false,
	}
	TMC4671_ABN_BPOL_FIELD = RegisterField{
		Mask:     TMC4671_ABN_BPOL_MASK,
		Shift:    TMC4671_ABN_BPOL_SHIFT,
		Register: TMC4671_ABN_DECODER_MODE,
		IsSigned: false,
	}
	TMC4671_ABN_NPOL_FIELD = RegisterField{
		Mask:     TMC4671_ABN_NPOL_MASK,
		Shift:    TMC4671_ABN_NPOL_SHIFT,
		Register: TMC4671_ABN_DECODER_MODE,
		IsSigned: false,
	}
	TMC4671_USE_ABN_AS_N_FIELD = RegisterField{
		Mask:     TMC4671_USE_ABN_AS_N_MASK,
		Shift:    TMC4671_USE_ABN_AS_N_SHIFT,
		Register: TMC4671_ABN_DECODER_MODE,
		IsSigned: false,
	}
	TMC4671_ABN_CLN_FIELD = RegisterField{
		Mask:     TMC4671_ABN_CLN_MASK,
		Shift:    TMC4671_ABN_CLN_SHIFT,
		Register: TMC4671_ABN_DECODER_MODE,
		IsSigned: false,
	}
	TMC4671_ABN_DIRECTION_FIELD = RegisterField{
		Mask:     TMC4671_ABN_DIRECTION_MASK,
		Shift:    TMC4671_ABN_DIRECTION_SHIFT,
		Register: TMC4671_ABN_DECODER_MODE,
		IsSigned: false,
	}
	TMC4671_ABN_DECODER_PPR_FIELD = RegisterField{
		Mask:     TMC4671_ABN_DECODER_PPR_MASK,
		Shift:    TMC4671_ABN_DECODER_PPR_SHIFT,
		Register: TMC4671_ABN_DECODER_PPR,
		IsSigned: false,
	}
	TMC4671_ABN_DECODER_COUNT_FIELD = RegisterField{
		Mask:     TMC4671_ABN_DECODER_COUNT_MASK,
		Shift:    TMC4671_ABN_DECODER_COUNT_SHIFT,
		Register: TMC4671_ABN_DECODER_COUNT,
		IsSigned: false,
	}
	TMC4671_ABN_DECODER_COUNT_N_FIELD = RegisterField{
		Mask:     TMC4671_ABN_DECODER_COUNT_N_MASK,
		Shift:    TMC4671_ABN_DECODER_COUNT_N_SHIFT,
		Register: TMC4671_ABN_DECODER_COUNT_N,
		IsSigned: false,
	}
	TMC4671_ABN_DECODER_PHI_M_OFFSET_FIELD = RegisterField{
		Mask:     TMC4671_ABN_DECODER_PHI_M_OFFSET_MASK,
		Shift:    TMC4671_ABN_DECODER_PHI_M_OFFSET_SHIFT,
		Register: TMC4671_ABN_DECODER_PHI_E_PHI_M_OFFSET,
		IsSigned: true,
	}
	TMC4671_ABN_DECODER_PHI_E_OFFSET_FIELD = RegisterField{
		Mask:     TMC4671_ABN_DECODER_PHI_E_OFFSET_MASK,
		Shift:    TMC4671_ABN_DECODER_PHI_E_OFFSET_SHIFT,
		Register: TMC4671_ABN_DECODER_PHI_E_PHI_M_OFFSET,
		IsSigned: true,
	}
	TMC4671_ABN_DECODER_PHI_M_FIELD = RegisterField{
		Mask:     TMC4671_ABN_DECODER_PHI_M_MASK,
		Shift:    TMC4671_ABN_DECODER_PHI_M_SHIFT,
		Register: TMC4671_ABN_DECODER_PHI_E_PHI_M,
		IsSigned: true,
	}
	TMC4671_ABN_DECODER_PHI_E_FIELD = RegisterField{
		Mask:     TMC4671_ABN_DECODER_PHI_E_MASK,
		Shift:    TMC4671_ABN_DECODER_PHI_E_SHIFT,
		Register: TMC4671_ABN_DECODER_PHI_E_PHI_M,
		IsSigned: true,
	}
	TMC4671_ABN_2_APOL_FIELD = RegisterField{
		Mask:     TMC4671_ABN_2_APOL_MASK,
		Shift:    TMC4671_ABN_2_APOL_SHIFT,
		Register: TMC4671_ABN_2_DECODER_MODE,
		IsSigned: false,
	}
	TMC4671_ABN_2_BPOL_FIELD = RegisterField{
		Mask:     TMC4671_ABN_2_BPOL_MASK,
		Shift:    TMC4671_ABN_2_BPOL_SHIFT,
		Register: TMC4671_ABN_2_DECODER_MODE,
		IsSigned: false,
	}
	TMC4671_ABN_2_NPOL_FIELD = RegisterField{
		Mask:     TMC4671_ABN_2_NPOL_MASK,
		Shift:    TMC4671_ABN_2_NPOL_SHIFT,
		Register: TMC4671_ABN_2_DECODER_MODE,
		IsSigned: false,
	}
	TMC4671_USE_ABN_2_AS_N_FIELD = RegisterField{
		Mask:     TMC4671_USE_ABN_2_AS_N_MASK,
		Shift:    TMC4671_USE_ABN_2_AS_N_SHIFT,
		Register: TMC4671_ABN_2_DECODER_MODE,
		IsSigned: false,
	}
	TMC4671_ABN_2_CLN_FIELD = RegisterField{
		Mask:     TMC4671_ABN_2_CLN_MASK,
		Shift:    TMC4671_ABN_2_CLN_SHIFT,
		Register: TMC4671_ABN_2_DECODER_MODE,
		IsSigned: false,
	}
	TMC4671_ABN_2_DIRECTION_FIELD = RegisterField{
		Mask:     TMC4671_ABN_2_DIRECTION_MASK,
		Shift:    TMC4671_ABN_2_DIRECTION_SHIFT,
		Register: TMC4671_ABN_2_DECODER_MODE,
		IsSigned: false,
	}
	TMC4671_ABN_2_DECODER_PPR_FIELD = RegisterField{
		Mask:     TMC4671_ABN_2_DECODER_PPR_MASK,
		Shift:    TMC4671_ABN_2_DECODER_PPR_SHIFT,
		Register: TMC4671_ABN_2_DECODER_PPR,
		IsSigned: false,
	}
	TMC4671_ABN_2_DECODER_COUNT_FIELD = RegisterField{
		Mask:     TMC4671_ABN_2_DECODER_COUNT_MASK,
		Shift:    TMC4671_ABN_2_DECODER_COUNT_SHIFT,
		Register: TMC4671_ABN_2_DECODER_COUNT,
		IsSigned: false,
	}
	TMC4671_ABN_2_DECODER_COUNT_N_FIELD = RegisterField{
		Mask:     TMC4671_ABN_2_DECODER_COUNT_N_MASK,
		Shift:    TMC4671_ABN_2_DECODER_COUNT_N_SHIFT,
		Register: TMC4671_ABN_2_DECODER_COUNT_N,
		IsSigned: false,
	}
	TMC4671_ABN_2_DECODER_PHI_M_OFFSET_FIELD = RegisterField{
		Mask:     TMC4671_ABN_2_DECODER_PHI_M_OFFSET_MASK,
		Shift:    TMC4671_ABN_2_DECODER_PHI_M_OFFSET_SHIFT,
		Register: TMC4671_ABN_2_DECODER_PHI_M_OFFSET,
		IsSigned: true,
	}
	TMC4671_ABN_2_DECODER_PHI_M_FIELD = RegisterField{
		Mask:     TMC4671_ABN_2_DECODER_PHI_M_MASK,
		Shift:    TMC4671_ABN_2_DECODER_PHI_M_SHIFT,
		Register: TMC4671_ABN_2_DECODER_PHI_M,
		IsSigned: true,
	}
	TMC4671_HALL_POLARITY_FIELD = RegisterField{
		Mask:     TMC4671_HALL_POLARITY_MASK,
		Shift:    TMC4671_HALL_POLARITY_SHIFT,
		Register: TMC4671_HALL_MODE,
		IsSigned: false,
	}
	TMC4671_HALL_SYNCHRONOUS_PWM_SAMPLING_FIELD = RegisterField{
		Mask:     TMC4671_HALL_SYNCHRONOUS_PWM_SAMPLING_MASK,
		Shift:    TMC4671_HALL_SYNCHRONOUS_PWM_SAMPLING_SHIFT,
		Register: TMC4671_HALL_MODE,
		IsSigned: false,
	}
	TMC4671_HALL_INTERPOLATION_FIELD = RegisterField{
		Mask:     TMC4671_HALL_INTERPOLATION_MASK,
		Shift:    TMC4671_HALL_INTERPOLATION_SHIFT,
		Register: TMC4671_HALL_MODE,
		IsSigned: false,
	}
	TMC4671_HALL_DIRECTION_FIELD = RegisterField{
		Mask:     TMC4671_HALL_DIRECTION_MASK,
		Shift:    TMC4671_HALL_DIRECTION_SHIFT,
		Register: TMC4671_HALL_MODE,
		IsSigned: false,
	}
	TMC4671_HALL_BLANK_FIELD = RegisterField{
		Mask:     TMC4671_HALL_BLANK_MASK,
		Shift:    TMC4671_HALL_BLANK_SHIFT,
		Register: TMC4671_HALL_MODE,
		IsSigned: false,
	}
	TMC4671_HALL_POSITION_000_FIELD = RegisterField{
		Mask:     TMC4671_HALL_POSITION_000_MASK,
		Shift:    TMC4671_HALL_POSITION_000_SHIFT,
		Register: TMC4671_HALL_POSITION_060_000,
		IsSigned: true,
	}
	TMC4671_HALL_POSITION_060_FIELD = RegisterField{
		Mask:     TMC4671_HALL_POSITION_060_MASK,
		Shift:    TMC4671_HALL_POSITION_060_SHIFT,
		Register: TMC4671_HALL_POSITION_060_000,
		IsSigned: true,
	}
	TMC4671_HALL_POSITION_120_FIELD = RegisterField{
		Mask:     TMC4671_HALL_POSITION_120_MASK,
		Shift:    TMC4671_HALL_POSITION_120_SHIFT,
		Register: TMC4671_HALL_POSITION_180_120,
		IsSigned: true,
	}
	TMC4671_HALL_POSITION_180_FIELD = RegisterField{
		Mask:     TMC4671_HALL_POSITION_180_MASK,
		Shift:    TMC4671_HALL_POSITION_180_SHIFT,
		Register: TMC4671_HALL_POSITION_180_120,
		IsSigned: true,
	}
	TMC4671_HALL_POSITION_240_FIELD = RegisterField{
		Mask:     TMC4671_HALL_POSITION_240_MASK,
		Shift:    TMC4671_HALL_POSITION_240_SHIFT,
		Register: TMC4671_HALL_POSITION_300_240,
		IsSigned: true,
	}
	TMC4671_HALL_POSITION_300_FIELD = RegisterField{
		Mask:     TMC4671_HALL_POSITION_300_MASK,
		Shift:    TMC4671_HALL_POSITION_300_SHIFT,
		Register: TMC4671_HALL_POSITION_300_240,
		IsSigned: true,
	}
	TMC4671_HALL_PHI_M_OFFSET_FIELD = RegisterField{
		Mask:     TMC4671_HALL_PHI_M_OFFSET_MASK,
		Shift:    TMC4671_HALL_PHI_M_OFFSET_SHIFT,
		Register: TMC4671_HALL_PHI_E_PHI_M_OFFSET,
		IsSigned: true,
	}
	TMC4671_HALL_PHI_E_OFFSET_FIELD = RegisterField{
		Mask:     TMC4671_HALL_PHI_E_OFFSET_MASK,
		Shift:    TMC4671_HALL_PHI_E_OFFSET_SHIFT,
		Register: TMC4671_HALL_PHI_E_PHI_M_OFFSET,
		IsSigned: true,
	}
	TMC4671_HALL_DPHI_MAX_FIELD = RegisterField{
		Mask:     TMC4671_HALL_DPHI_MAX_MASK,
		Shift:    TMC4671_HALL_DPHI_MAX_SHIFT,
		Register: TMC4671_HALL_DPHI_MAX,
		IsSigned: false,
	}
	TMC4671_HALL_PHI_E_FIELD = RegisterField{
		Mask:     TMC4671_HALL_PHI_E_MASK,
		Shift:    TMC4671_HALL_PHI_E_SHIFT,
		Register: TMC4671_HALL_PHI_E_INTERPOLATED_PHI_E,
		IsSigned: true,
	}
	TMC4671_HALL_PHI_E_INTERPOLATED_FIELD = RegisterField{
		Mask:     TMC4671_HALL_PHI_E_INTERPOLATED_MASK,
		Shift:    TMC4671_HALL_PHI_E_INTERPOLATED_SHIFT,
		Register: TMC4671_HALL_PHI_E_INTERPOLATED_PHI_E,
		IsSigned: true,
	}
	TMC4671_HALL_PHI_M_FIELD = RegisterField{
		Mask:     TMC4671_HALL_PHI_M_MASK,
		Shift:    TMC4671_HALL_PHI_M_SHIFT,
		Register: TMC4671_HALL_PHI_M,
		IsSigned: true,
	}
	TMC4671_AENC_DECODER_MODE_0_FIELD = RegisterField{
		Mask:     TMC4671_AENC_DECODER_MODE_0_MASK,
		Shift:    TMC4671_AENC_DECODER_MODE_0_SHIFT,
		Register: TMC4671_AENC_DECODER_MODE,
		IsSigned: false,
	}
	TMC4671_AENC_DECODER_MODE_12_FIELD = RegisterField{
		Mask:     TMC4671_AENC_DECODER_MODE_12_MASK,
		Shift:    TMC4671_AENC_DECODER_MODE_12_SHIFT,
		Register: TMC4671_AENC_DECODER_MODE,
		IsSigned: false,
	}
	TMC4671_AENC_DECODER_N_THRESHOLD_FIELD = RegisterField{
		Mask:     TMC4671_AENC_DECODER_N_THRESHOLD_MASK,
		Shift:    TMC4671_AENC_DECODER_N_THRESHOLD_SHIFT,
		Register: TMC4671_AENC_DECODER_N_THRESHOLD,
		IsSigned: false,
	}
	TMC4671_AENC_DECODER_PHI_A_RAW_FIELD = RegisterField{
		Mask:     TMC4671_AENC_DECODER_PHI_A_RAW_MASK,
		Shift:    TMC4671_AENC_DECODER_PHI_A_RAW_SHIFT,
		Register: TMC4671_AENC_DECODER_PHI_A_RAW,
		IsSigned: true,
	}
	TMC4671_AENC_DECODER_PHI_A_OFFSET_FIELD = RegisterField{
		Mask:     TMC4671_AENC_DECODER_PHI_A_OFFSET_MASK,
		Shift:    TMC4671_AENC_DECODER_PHI_A_OFFSET_SHIFT,
		Register: TMC4671_AENC_DECODER_PHI_A_OFFSET,
		IsSigned: true,
	}
	TMC4671_AENC_DECODER_PHI_A_FIELD = RegisterField{
		Mask:     TMC4671_AENC_DECODER_PHI_A_MASK,
		Shift:    TMC4671_AENC_DECODER_PHI_A_SHIFT,
		Register: TMC4671_AENC_DECODER_PHI_A,
		IsSigned: true,
	}
	TMC4671_AENC_DECODER_PPR_FIELD = RegisterField{
		Mask:     TMC4671_AENC_DECODER_PPR_MASK,
		Shift:    TMC4671_AENC_DECODER_PPR_SHIFT,
		Register: TMC4671_AENC_DECODER_PPR,
		IsSigned: true,
	}
	TMC4671_AENC_DECODER_COUNT_FIELD = RegisterField{
		Mask:     TMC4671_AENC_DECODER_COUNT_MASK,
		Shift:    TMC4671_AENC_DECODER_COUNT_SHIFT,
		Register: TMC4671_AENC_DECODER_COUNT,
		IsSigned: true,
	}
	TMC4671_AENC_DECODER_COUNT_N_FIELD = RegisterField{
		Mask:     TMC4671_AENC_DECODER_COUNT_N_MASK,
		Shift:    TMC4671_AENC_DECODER_COUNT_N_SHIFT,
		Register: TMC4671_AENC_DECODER_COUNT_N,
		IsSigned: true,
	}
	TMC4671_AENC_DECODER_PHI_M_OFFSET_FIELD = RegisterField{
		Mask:     TMC4671_AENC_DECODER_PHI_M_OFFSET_MASK,
		Shift:    TMC4671_AENC_DECODER_PHI_M_OFFSET_SHIFT,
		Register: TMC4671_AENC_DECODER_PHI_E_PHI_M_OFFSET,
		IsSigned: true,
	}
	TMC4671_AENC_DECODER_PHI_E_OFFSET_FIELD = RegisterField{
		Mask:     TMC4671_AENC_DECODER_PHI_E_OFFSET_MASK,
		Shift:    TMC4671_AENC_DECODER_PHI_E_OFFSET_SHIFT,
		Register: TMC4671_AENC_DECODER_PHI_E_PHI_M_OFFSET,
		IsSigned: true,
	}
	TMC4671_AENC_DECODER_PHI_M_FIELD = RegisterField{
		Mask:     TMC4671_AENC_DECODER_PHI_M_MASK,
		Shift:    TMC4671_AENC_DECODER_PHI_M_SHIFT,
		Register: TMC4671_AENC_DECODER_PHI_E_PHI_M,
		IsSigned: true,
	}
	TMC4671_AENC_DECODER_PHI_E_FIELD = RegisterField{
		Mask:     TMC4671_AENC_DECODER_PHI_E_MASK,
		Shift:    TMC4671_AENC_DECODER_PHI_E_SHIFT,
		Register: TMC4671_AENC_DECODER_PHI_E_PHI_M,
		IsSigned: true,
	}
	TMC4671_BIQUAD_X_A_1_FIELD = RegisterField{
		Mask:     TMC4671_BIQUAD_X_A_1_MASK,
		Shift:    TMC4671_BIQUAD_X_A_1_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: true,
	}
	TMC4671_BIQUAD_X_A_2_FIELD = RegisterField{
		Mask:     TMC4671_BIQUAD_X_A_2_MASK,
		Shift:    TMC4671_BIQUAD_X_A_2_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: true,
	}
	TMC4671_BIQUAD_X_B_0_FIELD = RegisterField{
		Mask:     TMC4671_BIQUAD_X_B_0_MASK,
		Shift:    TMC4671_BIQUAD_X_B_0_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: true,
	}
	TMC4671_BIQUAD_X_B_1_FIELD = RegisterField{
		Mask:     TMC4671_BIQUAD_X_B_1_MASK,
		Shift:    TMC4671_BIQUAD_X_B_1_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: true,
	}
	TMC4671_BIQUAD_X_B_2_FIELD = RegisterField{
		Mask:     TMC4671_BIQUAD_X_B_2_MASK,
		Shift:    TMC4671_BIQUAD_X_B_2_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: true,
	}
	TMC4671_BIQUAD_X_ENABLE_FIELD = RegisterField{
		Mask:     TMC4671_BIQUAD_X_ENABLE_MASK,
		Shift:    TMC4671_BIQUAD_X_ENABLE_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: false,
	}
	TMC4671_BIQUAD_V_A_1_FIELD = RegisterField{
		Mask:     TMC4671_BIQUAD_V_A_1_MASK,
		Shift:    TMC4671_BIQUAD_V_A_1_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: true,
	}
	TMC4671_BIQUAD_V_A_2_FIELD = RegisterField{
		Mask:     TMC4671_BIQUAD_V_A_2_MASK,
		Shift:    TMC4671_BIQUAD_V_A_2_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: true,
	}
	TMC4671_BIQUAD_V_B_0_FIELD = RegisterField{
		Mask:     TMC4671_BIQUAD_V_B_0_MASK,
		Shift:    TMC4671_BIQUAD_V_B_0_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: true,
	}
	TMC4671_BIQUAD_V_B_1_FIELD = RegisterField{
		Mask:     TMC4671_BIQUAD_V_B_1_MASK,
		Shift:    TMC4671_BIQUAD_V_B_1_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: true,
	}
	TMC4671_BIQUAD_V_B_2_FIELD = RegisterField{
		Mask:     TMC4671_BIQUAD_V_B_2_MASK,
		Shift:    TMC4671_BIQUAD_V_B_2_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: true,
	}
	TMC4671_BIQUAD_V_ENABLE_FIELD = RegisterField{
		Mask:     TMC4671_BIQUAD_V_ENABLE_MASK,
		Shift:    TMC4671_BIQUAD_V_ENABLE_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: false,
	}
	TMC4671_BIQUAD_T_A_1_FIELD = RegisterField{
		Mask:     TMC4671_BIQUAD_T_A_1_MASK,
		Shift:    TMC4671_BIQUAD_T_A_1_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: true,
	}
	TMC4671_BIQUAD_T_A_2_FIELD = RegisterField{
		Mask:     TMC4671_BIQUAD_T_A_2_MASK,
		Shift:    TMC4671_BIQUAD_T_A_2_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: true,
	}
	TMC4671_BIQUAD_T_B_0_FIELD = RegisterField{
		Mask:     TMC4671_BIQUAD_T_B_0_MASK,
		Shift:    TMC4671_BIQUAD_T_B_0_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: true,
	}
	TMC4671_BIQUAD_T_B_1_FIELD = RegisterField{
		Mask:     TMC4671_BIQUAD_T_B_1_MASK,
		Shift:    TMC4671_BIQUAD_T_B_1_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: true,
	}
	TMC4671_BIQUAD_T_B_2_FIELD = RegisterField{
		Mask:     TMC4671_BIQUAD_T_B_2_MASK,
		Shift:    TMC4671_BIQUAD_T_B_2_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: true,
	}
	TMC4671_BIQUAD_T_ENABLE_FIELD = RegisterField{
		Mask:     TMC4671_BIQUAD_T_ENABLE_MASK,
		Shift:    TMC4671_BIQUAD_T_ENABLE_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: false,
	}
	TMC4671_BIQUAD_F_A_1_FIELD = RegisterField{
		Mask:     TMC4671_BIQUAD_F_A_1_MASK,
		Shift:    TMC4671_BIQUAD_F_A_1_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: true,
	}
	TMC4671_BIQUAD_F_A_2_FIELD = RegisterField{
		Mask:     TMC4671_BIQUAD_F_A_2_MASK,
		Shift:    TMC4671_BIQUAD_F_A_2_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: true,
	}
	TMC4671_BIQUAD_F_B_0_FIELD = RegisterField{
		Mask:     TMC4671_BIQUAD_F_B_0_MASK,
		Shift:    TMC4671_BIQUAD_F_B_0_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: true,
	}
	TMC4671_BIQUAD_F_B_1_FIELD = RegisterField{
		Mask:     TMC4671_BIQUAD_F_B_1_MASK,
		Shift:    TMC4671_BIQUAD_F_B_1_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: true,
	}
	TMC4671_BIQUAD_F_B_2_FIELD = RegisterField{
		Mask:     TMC4671_BIQUAD_F_B_2_MASK,
		Shift:    TMC4671_BIQUAD_F_B_2_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: true,
	}
	TMC4671_BIQUAD_F_ENABLE_FIELD = RegisterField{
		Mask:     TMC4671_BIQUAD_F_ENABLE_MASK,
		Shift:    TMC4671_BIQUAD_F_ENABLE_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: false,
	}
	TMC4671_PRBS_AMPLITUDE_FIELD = RegisterField{
		Mask:     TMC4671_PRBS_AMPLITUDE_MASK,
		Shift:    TMC4671_PRBS_AMPLITUDE_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: true,
	}
	TMC4671_PRBS_DOWN_SAMPLING_RATIO_FIELD = RegisterField{
		Mask:     TMC4671_PRBS_DOWN_SAMPLING_RATIO_MASK,
		Shift:    TMC4671_PRBS_DOWN_SAMPLING_RATIO_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: true,
	}
	TMC4671_REF_SWITCH_ENABLE_FIELD = RegisterField{
		Mask:     TMC4671_REF_SWITCH_ENABLE_MASK,
		Shift:    TMC4671_REF_SWITCH_ENABLE_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: false,
	}
	TMC4671_HOME_SWITCH_POLARITY_FIELD = RegisterField{
		Mask:     TMC4671_HOME_SWITCH_POLARITY_MASK,
		Shift:    TMC4671_HOME_SWITCH_POLARITY_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: false,
	}
	TMC4671_LEFT_SWITCH_POLARITY_FIELD = RegisterField{
		Mask:     TMC4671_LEFT_SWITCH_POLARITY_MASK,
		Shift:    TMC4671_LEFT_SWITCH_POLARITY_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: false,
	}
	TMC4671_RIGHT_SWITCH_POLARITY_FIELD = RegisterField{
		Mask:     TMC4671_RIGHT_SWITCH_POLARITY_MASK,
		Shift:    TMC4671_RIGHT_SWITCH_POLARITY_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: false,
	}
	TMC4671_ENCODER_INIT_HALL_ENABLE_FIELD = RegisterField{
		Mask:     TMC4671_ENCODER_INIT_HALL_ENABLE_MASK,
		Shift:    TMC4671_ENCODER_INIT_HALL_ENABLE_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: false,
	}
	TMC4671_SINGLE_PIN_IF_CFG_FIELD = RegisterField{
		Mask:     TMC4671_SINGLE_PIN_IF_CFG_MASK,
		Shift:    TMC4671_SINGLE_PIN_IF_CFG_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: false,
	}
	TMC4671_SINGLE_PIN_IF_STATUS_FIELD = RegisterField{
		Mask:     TMC4671_SINGLE_PIN_IF_STATUS_MASK,
		Shift:    TMC4671_SINGLE_PIN_IF_STATUS_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: false,
	}
	TMC4671_SINGLE_PIN_IF_OFFSET_FIELD = RegisterField{
		Mask:     TMC4671_SINGLE_PIN_IF_OFFSET_MASK,
		Shift:    TMC4671_SINGLE_PIN_IF_OFFSET_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: false,
	}
	TMC4671_SINGLE_PIN_IF_SCALE_FIELD = RegisterField{
		Mask:     TMC4671_SINGLE_PIN_IF_SCALE_MASK,
		Shift:    TMC4671_SINGLE_PIN_IF_SCALE_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: true,
	}
	TMC4671_CURRENT_I_NQ8_8_Q4_12_FIELD = RegisterField{
		Mask:     TMC4671_CURRENT_I_NQ8_8_Q4_12_MASK,
		Shift:    TMC4671_CURRENT_I_NQ8_8_Q4_12_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: false,
	}
	TMC4671_CURRENT_P_NQ8_8_Q4_12_FIELD = RegisterField{
		Mask:     TMC4671_CURRENT_P_NQ8_8_Q4_12_MASK,
		Shift:    TMC4671_CURRENT_P_NQ8_8_Q4_12_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: false,
	}
	TMC4671_VELOCITY_I_NQ8_8_Q4_12_FIELD = RegisterField{
		Mask:     TMC4671_VELOCITY_I_NQ8_8_Q4_12_MASK,
		Shift:    TMC4671_VELOCITY_I_NQ8_8_Q4_12_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: false,
	}
	TMC4671_VELOCITY_P_NQ8_8_Q4_12_FIELD = RegisterField{
		Mask:     TMC4671_VELOCITY_P_NQ8_8_Q4_12_MASK,
		Shift:    TMC4671_VELOCITY_P_NQ8_8_Q4_12_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: false,
	}
	TMC4671_POSITION_I_NQ8_8_Q4_12_FIELD = RegisterField{
		Mask:     TMC4671_POSITION_I_NQ8_8_Q4_12_MASK,
		Shift:    TMC4671_POSITION_I_NQ8_8_Q4_12_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: false,
	}
	TMC4671_POSITION_P_NQ8_8_Q4_12_FIELD = RegisterField{
		Mask:     TMC4671_POSITION_P_NQ8_8_Q4_12_MASK,
		Shift:    TMC4671_POSITION_P_NQ8_8_Q4_12_SHIFT,
		Register: TMC4671_CONFIG_DATA,
		IsSigned: false,
	}
	TMC4671_CONFIG_ADDR_FIELD = RegisterField{
		Mask:     TMC4671_CONFIG_ADDR_MASK,
		Shift:    TMC4671_CONFIG_ADDR_SHIFT,
		Register: TMC4671_CONFIG_ADDR,
		IsSigned: false,
	}
	TMC4671_VELOCITY_SELECTION_FIELD = RegisterField{
		Mask:     TMC4671_VELOCITY_SELECTION_MASK,
		Shift:    TMC4671_VELOCITY_SELECTION_SHIFT,
		Register: TMC4671_VELOCITY_SELECTION,
		IsSigned: false,
	}
	TMC4671_VELOCITY_METER_SELECTION_FIELD = RegisterField{
		Mask:     TMC4671_VELOCITY_METER_SELECTION_MASK,
		Shift:    TMC4671_VELOCITY_METER_SELECTION_SHIFT,
		Register: TMC4671_VELOCITY_SELECTION,
		IsSigned: false,
	}
	TMC4671_POSITION_SELECTION_FIELD = RegisterField{
		Mask:     TMC4671_POSITION_SELECTION_MASK,
		Shift:    TMC4671_POSITION_SELECTION_SHIFT,
		Register: TMC4671_POSITION_SELECTION,
		IsSigned: false,
	}
	TMC4671_PHI_E_SELECTION_FIELD = RegisterField{
		Mask:     TMC4671_PHI_E_SELECTION_MASK,
		Shift:    TMC4671_PHI_E_SELECTION_SHIFT,
		Register: TMC4671_PHI_E_SELECTION,
		IsSigned: false,
	}
	TMC4671_PHI_E_FIELD = RegisterField{
		Mask:     TMC4671_PHI_E_MASK,
		Shift:    TMC4671_PHI_E_SHIFT,
		Register: TMC4671_PHI_E,
		IsSigned: true,
	}
	TMC4671_PID_FLUX_I_FIELD = RegisterField{
		Mask:     TMC4671_PID_FLUX_I_MASK,
		Shift:    TMC4671_PID_FLUX_I_SHIFT,
		Register: TMC4671_PID_FLUX_P_FLUX_I,
		IsSigned: true,
	}
	TMC4671_PID_FLUX_P_FIELD = RegisterField{
		Mask:     TMC4671_PID_FLUX_P_MASK,
		Shift:    TMC4671_PID_FLUX_P_SHIFT,
		Register: TMC4671_PID_FLUX_P_FLUX_I,
		IsSigned: true,
	}
	TMC4671_PID_TORQUE_I_FIELD = RegisterField{
		Mask:     TMC4671_PID_TORQUE_I_MASK,
		Shift:    TMC4671_PID_TORQUE_I_SHIFT,
		Register: TMC4671_PID_TORQUE_P_TORQUE_I,
		IsSigned: true,
	}
	TMC4671_PID_TORQUE_P_FIELD = RegisterField{
		Mask:     TMC4671_PID_TORQUE_P_MASK,
		Shift:    TMC4671_PID_TORQUE_P_SHIFT,
		Register: TMC4671_PID_TORQUE_P_TORQUE_I,
		IsSigned: true,
	}
	TMC4671_PID_VELOCITY_I_FIELD = RegisterField{
		Mask:     TMC4671_PID_VELOCITY_I_MASK,
		Shift:    TMC4671_PID_VELOCITY_I_SHIFT,
		Register: TMC4671_PID_VELOCITY_P_VELOCITY_I,
		IsSigned: true,
	}
	TMC4671_PID_VELOCITY_P_FIELD = RegisterField{
		Mask:     TMC4671_PID_VELOCITY_P_MASK,
		Shift:    TMC4671_PID_VELOCITY_P_SHIFT,
		Register: TMC4671_PID_VELOCITY_P_VELOCITY_I,
		IsSigned: true,
	}
	TMC4671_PID_POSITION_I_FIELD = RegisterField{
		Mask:     TMC4671_PID_POSITION_I_MASK,
		Shift:    TMC4671_PID_POSITION_I_SHIFT,
		Register: TMC4671_PID_POSITION_P_POSITION_I,
		IsSigned: true,
	}
	TMC4671_PID_POSITION_P_FIELD = RegisterField{
		Mask:     TMC4671_PID_POSITION_P_MASK,
		Shift:    TMC4671_PID_POSITION_P_SHIFT,
		Register: TMC4671_PID_POSITION_P_POSITION_I,
		IsSigned: true,
	}
	TMC4671_PIDOUT_UQ_UD_LIMITS_FIELD = RegisterField{
		Mask:     TMC4671_PIDOUT_UQ_UD_LIMITS_MASK,
		Shift:    TMC4671_PIDOUT_UQ_UD_LIMITS_SHIFT,
		Register: TMC4671_PIDOUT_UQ_UD_LIMITS,
		IsSigned: false,
	}
	TMC4671_PID_TORQUE_FLUX_LIMITS_FIELD = RegisterField{
		Mask:     TMC4671_PID_TORQUE_FLUX_LIMITS_MASK,
		Shift:    TMC4671_PID_TORQUE_FLUX_LIMITS_SHIFT,
		Register: TMC4671_PID_TORQUE_FLUX_LIMITS,
		IsSigned: false,
	}
	TMC4671_PID_VELOCITY_LIMIT_FIELD = RegisterField{
		Mask:     TMC4671_PID_VELOCITY_LIMIT_MASK,
		Shift:    TMC4671_PID_VELOCITY_LIMIT_SHIFT,
		Register: TMC4671_PID_VELOCITY_LIMIT,
		IsSigned: false,
	}
	TMC4671_PID_POSITION_LIMIT_LOW_FIELD = RegisterField{
		Mask:     TMC4671_PID_POSITION_LIMIT_LOW_MASK,
		Shift:    TMC4671_PID_POSITION_LIMIT_LOW_SHIFT,
		Register: TMC4671_PID_POSITION_LIMIT_LOW,
		IsSigned: true,
	}
	TMC4671_PID_POSITION_LIMIT_HIGH_FIELD = RegisterField{
		Mask:     TMC4671_PID_POSITION_LIMIT_HIGH_MASK,
		Shift:    TMC4671_PID_POSITION_LIMIT_HIGH_SHIFT,
		Register: TMC4671_PID_POSITION_LIMIT_HIGH,
		IsSigned: true,
	}
	TMC4671_MODE_MOTION_FIELD = RegisterField{
		Mask:     TMC4671_MODE_MOTION_MASK,
		Shift:    TMC4671_MODE_MOTION_SHIFT,
		Register: TMC4671_MODE_RAMP_MODE_MOTION,
		IsSigned: false,
	}
	TMC4671_MODE_RAMP_FIELD = RegisterField{
		Mask:     TMC4671_MODE_RAMP_MASK,
		Shift:    TMC4671_MODE_RAMP_SHIFT,
		Register: TMC4671_MODE_RAMP_MODE_MOTION,
		IsSigned: false,
	}
	TMC4671_MODE_FF_FIELD = RegisterField{
		Mask:     TMC4671_MODE_FF_MASK,
		Shift:    TMC4671_MODE_FF_SHIFT,
		Register: TMC4671_MODE_RAMP_MODE_MOTION,
		IsSigned: false,
	}
	TMC4671_MODE_PID_SMPL_FIELD = RegisterField{
		Mask:     TMC4671_MODE_PID_SMPL_MASK,
		Shift:    TMC4671_MODE_PID_SMPL_SHIFT,
		Register: TMC4671_MODE_RAMP_MODE_MOTION,
		IsSigned: false,
	}
	TMC4671_MODE_PID_TYPE_FIELD = RegisterField{
		Mask:     TMC4671_MODE_PID_TYPE_MASK,
		Shift:    TMC4671_MODE_PID_TYPE_SHIFT,
		Register: TMC4671_MODE_RAMP_MODE_MOTION,
		IsSigned: false,
	}
	TMC4671_PID_FLUX_TARGET_FIELD = RegisterField{
		Mask:     TMC4671_PID_FLUX_TARGET_MASK,
		Shift:    TMC4671_PID_FLUX_TARGET_SHIFT,
		Register: TMC4671_PID_TORQUE_FLUX_TARGET,
		IsSigned: true,
	}
	TMC4671_PID_TORQUE_TARGET_FIELD = RegisterField{
		Mask:     TMC4671_PID_TORQUE_TARGET_MASK,
		Shift:    TMC4671_PID_TORQUE_TARGET_SHIFT,
		Register: TMC4671_PID_TORQUE_FLUX_TARGET,
		IsSigned: true,
	}
	TMC4671_PID_FLUX_OFFSET_FIELD = RegisterField{
		Mask:     TMC4671_PID_FLUX_OFFSET_MASK,
		Shift:    TMC4671_PID_FLUX_OFFSET_SHIFT,
		Register: TMC4671_PID_TORQUE_FLUX_OFFSET,
		IsSigned: true,
	}
	TMC4671_PID_TORQUE_OFFSET_FIELD = RegisterField{
		Mask:     TMC4671_PID_TORQUE_OFFSET_MASK,
		Shift:    TMC4671_PID_TORQUE_OFFSET_SHIFT,
		Register: TMC4671_PID_TORQUE_FLUX_OFFSET,
		IsSigned: true,
	}
	TMC4671_PID_VELOCITY_TARGET_FIELD = RegisterField{
		Mask:     TMC4671_PID_VELOCITY_TARGET_MASK,
		Shift:    TMC4671_PID_VELOCITY_TARGET_SHIFT,
		Register: TMC4671_PID_VELOCITY_TARGET,
		IsSigned: true,
	}
	TMC4671_PID_VELOCITY_OFFSET_FIELD = RegisterField{
		Mask:     TMC4671_PID_VELOCITY_OFFSET_MASK,
		Shift:    TMC4671_PID_VELOCITY_OFFSET_SHIFT,
		Register: TMC4671_PID_VELOCITY_OFFSET,
		IsSigned: true,
	}
	TMC4671_PID_POSITION_TARGET_FIELD = RegisterField{
		Mask:     TMC4671_PID_POSITION_TARGET_MASK,
		Shift:    TMC4671_PID_POSITION_TARGET_SHIFT,
		Register: TMC4671_PID_POSITION_TARGET,
		IsSigned: true,
	}
	TMC4671_PID_FLUX_ACTUAL_FIELD = RegisterField{
		Mask:     TMC4671_PID_FLUX_ACTUAL_MASK,
		Shift:    TMC4671_PID_FLUX_ACTUAL_SHIFT,
		Register: TMC4671_PID_TORQUE_FLUX_ACTUAL,
		IsSigned: true,
	}
	TMC4671_PID_TORQUE_ACTUAL_FIELD = RegisterField{
		Mask:     TMC4671_PID_TORQUE_ACTUAL_MASK,
		Shift:    TMC4671_PID_TORQUE_ACTUAL_SHIFT,
		Register: TMC4671_PID_TORQUE_FLUX_ACTUAL,
		IsSigned: true,
	}
	TMC4671_PID_VELOCITY_ACTUAL_FIELD = RegisterField{
		Mask:     TMC4671_PID_VELOCITY_ACTUAL_MASK,
		Shift:    TMC4671_PID_VELOCITY_ACTUAL_SHIFT,
		Register: TMC4671_PID_VELOCITY_ACTUAL,
		IsSigned: true,
	}
	TMC4671_PID_POSITION_ACTUAL_FIELD = RegisterField{
		Mask:     TMC4671_PID_POSITION_ACTUAL_MASK,
		Shift:    TMC4671_PID_POSITION_ACTUAL_SHIFT,
		Register: TMC4671_PID_POSITION_ACTUAL,
		IsSigned: true,
	}
	TMC4671_PID_TORQUE_ERROR_FIELD = RegisterField{
		Mask:     TMC4671_PID_TORQUE_ERROR_MASK,
		Shift:    TMC4671_PID_TORQUE_ERROR_SHIFT,
		Register: TMC4671_PID_ERROR_DATA,
		IsSigned: true,
	}
	TMC4671_PID_FLUX_ERROR_FIELD = RegisterField{
		Mask:     TMC4671_PID_FLUX_ERROR_MASK,
		Shift:    TMC4671_PID_FLUX_ERROR_SHIFT,
		Register: TMC4671_PID_ERROR_DATA,
		IsSigned: true,
	}
	TMC4671_PID_VELOCITY_ERROR_FIELD = RegisterField{
		Mask:     TMC4671_PID_VELOCITY_ERROR_MASK,
		Shift:    TMC4671_PID_VELOCITY_ERROR_SHIFT,
		Register: TMC4671_PID_ERROR_DATA,
		IsSigned: true,
	}
	TMC4671_PID_POSITION_ERROR_FIELD = RegisterField{
		Mask:     TMC4671_PID_POSITION_ERROR_MASK,
		Shift:    TMC4671_PID_POSITION_ERROR_SHIFT,
		Register: TMC4671_PID_ERROR_DATA,
		IsSigned: true,
	}
	TMC4671_PID_TORQUE_ERROR_SUM_FIELD = RegisterField{
		Mask:     TMC4671_PID_TORQUE_ERROR_SUM_MASK,
		Shift:    TMC4671_PID_TORQUE_ERROR_SUM_SHIFT,
		Register: TMC4671_PID_ERROR_DATA,
		IsSigned: true,
	}
	TMC4671_PID_FLUX_ERROR_SUM_FIELD = RegisterField{
		Mask:     TMC4671_PID_FLUX_ERROR_SUM_MASK,
		Shift:    TMC4671_PID_FLUX_ERROR_SUM_SHIFT,
		Register: TMC4671_PID_ERROR_DATA,
		IsSigned: true,
	}
	TMC4671_PID_VELOCITY_ERROR_SUM_FIELD = RegisterField{
		Mask:     TMC4671_PID_VELOCITY_ERROR_SUM_MASK,
		Shift:    TMC4671_PID_VELOCITY_ERROR_SUM_SHIFT,
		Register: TMC4671_PID_ERROR_DATA,
		IsSigned: true,
	}
	TMC4671_PID_POSITION_ERROR_SUM_FIELD = RegisterField{
		Mask:     TMC4671_PID_POSITION_ERROR_SUM_MASK,
		Shift:    TMC4671_PID_POSITION_ERROR_SUM_SHIFT,
		Register: TMC4671_PID_ERROR_DATA,
		IsSigned: true,
	}
	TMC4671_PID_ERROR_ADDR_FIELD = RegisterField{
		Mask:     TMC4671_PID_ERROR_ADDR_MASK,
		Shift:    TMC4671_PID_ERROR_ADDR_SHIFT,
		Register: TMC4671_PID_ERROR_ADDR,
		IsSigned: false,
	}
	TMC4671_PIDIN_TARGET_TORQUE_FIELD = RegisterField{
		Mask:     TMC4671_PIDIN_TARGET_TORQUE_MASK,
		Shift:    TMC4671_PIDIN_TARGET_TORQUE_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_PIDIN_TARGET_FLUX_FIELD = RegisterField{
		Mask:     TMC4671_PIDIN_TARGET_FLUX_MASK,
		Shift:    TMC4671_PIDIN_TARGET_FLUX_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_PIDIN_TARGET_VELOCITY_FIELD = RegisterField{
		Mask:     TMC4671_PIDIN_TARGET_VELOCITY_MASK,
		Shift:    TMC4671_PIDIN_TARGET_VELOCITY_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_PIDIN_TARGET_POSITION_FIELD = RegisterField{
		Mask:     TMC4671_PIDIN_TARGET_POSITION_MASK,
		Shift:    TMC4671_PIDIN_TARGET_POSITION_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_PIDOUT_TARGET_TORQUE_FIELD = RegisterField{
		Mask:     TMC4671_PIDOUT_TARGET_TORQUE_MASK,
		Shift:    TMC4671_PIDOUT_TARGET_TORQUE_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_PIDOUT_TARGET_FLUX_FIELD = RegisterField{
		Mask:     TMC4671_PIDOUT_TARGET_FLUX_MASK,
		Shift:    TMC4671_PIDOUT_TARGET_FLUX_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_PIDOUT_TARGET_VELOCITY_FIELD = RegisterField{
		Mask:     TMC4671_PIDOUT_TARGET_VELOCITY_MASK,
		Shift:    TMC4671_PIDOUT_TARGET_VELOCITY_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_PIDOUT_TARGET_POSITION_FIELD = RegisterField{
		Mask:     TMC4671_PIDOUT_TARGET_POSITION_MASK,
		Shift:    TMC4671_PIDOUT_TARGET_POSITION_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_FOC_IUX_FIELD = RegisterField{
		Mask:     TMC4671_FOC_IUX_MASK,
		Shift:    TMC4671_FOC_IUX_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_FOC_IWY_FIELD = RegisterField{
		Mask:     TMC4671_FOC_IWY_MASK,
		Shift:    TMC4671_FOC_IWY_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_FOC_IV_FIELD = RegisterField{
		Mask:     TMC4671_FOC_IV_MASK,
		Shift:    TMC4671_FOC_IV_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_FOC_IA_FIELD = RegisterField{
		Mask:     TMC4671_FOC_IA_MASK,
		Shift:    TMC4671_FOC_IA_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_FOC_IB_FIELD = RegisterField{
		Mask:     TMC4671_FOC_IB_MASK,
		Shift:    TMC4671_FOC_IB_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_FOC_ID_FIELD = RegisterField{
		Mask:     TMC4671_FOC_ID_MASK,
		Shift:    TMC4671_FOC_ID_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_FOC_IQ_FIELD = RegisterField{
		Mask:     TMC4671_FOC_IQ_MASK,
		Shift:    TMC4671_FOC_IQ_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_FOC_UD_FIELD = RegisterField{
		Mask:     TMC4671_FOC_UD_MASK,
		Shift:    TMC4671_FOC_UD_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_FOC_UQ_FIELD = RegisterField{
		Mask:     TMC4671_FOC_UQ_MASK,
		Shift:    TMC4671_FOC_UQ_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_FOC_UD_LIMITED_FIELD = RegisterField{
		Mask:     TMC4671_FOC_UD_LIMITED_MASK,
		Shift:    TMC4671_FOC_UD_LIMITED_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_FOC_UQ_LIMITED_FIELD = RegisterField{
		Mask:     TMC4671_FOC_UQ_LIMITED_MASK,
		Shift:    TMC4671_FOC_UQ_LIMITED_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_FOC_UA_FIELD = RegisterField{
		Mask:     TMC4671_FOC_UA_MASK,
		Shift:    TMC4671_FOC_UA_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_FOC_UB_FIELD = RegisterField{
		Mask:     TMC4671_FOC_UB_MASK,
		Shift:    TMC4671_FOC_UB_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_FOC_UUX_FIELD = RegisterField{
		Mask:     TMC4671_FOC_UUX_MASK,
		Shift:    TMC4671_FOC_UUX_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_FOC_UWY_FIELD = RegisterField{
		Mask:     TMC4671_FOC_UWY_MASK,
		Shift:    TMC4671_FOC_UWY_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_FOC_UV_FIELD = RegisterField{
		Mask:     TMC4671_FOC_UV_MASK,
		Shift:    TMC4671_FOC_UV_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_PWM_UX_FIELD = RegisterField{
		Mask:     TMC4671_PWM_UX_MASK,
		Shift:    TMC4671_PWM_UX_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_PWM_WY_FIELD = RegisterField{
		Mask:     TMC4671_PWM_WY_MASK,
		Shift:    TMC4671_PWM_WY_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_PWM_V_FIELD = RegisterField{
		Mask:     TMC4671_PWM_V_MASK,
		Shift:    TMC4671_PWM_V_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_ADC_I_0_FIELD = RegisterField{
		Mask:     TMC4671_ADC_I_0_MASK,
		Shift:    TMC4671_ADC_I_0_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_ADC_I_1_FIELD = RegisterField{
		Mask:     TMC4671_ADC_I_1_MASK,
		Shift:    TMC4671_ADC_I_1_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_PID_FLUX_ACTUAL_DIV256_FIELD = RegisterField{
		Mask:     TMC4671_PID_FLUX_ACTUAL_DIV256_MASK,
		Shift:    TMC4671_PID_FLUX_ACTUAL_DIV256_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_PID_TORQUE_ACTUAL_DIV256_FIELD = RegisterField{
		Mask:     TMC4671_PID_TORQUE_ACTUAL_DIV256_MASK,
		Shift:    TMC4671_PID_TORQUE_ACTUAL_DIV256_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_PID_FLUX_TARGET_DIV256_FIELD = RegisterField{
		Mask:     TMC4671_PID_FLUX_TARGET_DIV256_MASK,
		Shift:    TMC4671_PID_FLUX_TARGET_DIV256_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_PID_TORQUE_TARGET_DIV256_FIELD = RegisterField{
		Mask:     TMC4671_PID_TORQUE_TARGET_DIV256_MASK,
		Shift:    TMC4671_PID_TORQUE_TARGET_DIV256_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_PID_VELOCITY_ACTUAL_DIV256_FIELD = RegisterField{
		Mask:     TMC4671_PID_VELOCITY_ACTUAL_DIV256_MASK,
		Shift:    TMC4671_PID_VELOCITY_ACTUAL_DIV256_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_PID_VELOCITY_TARGET_DIV256_FIELD = RegisterField{
		Mask:     TMC4671_PID_VELOCITY_TARGET_DIV256_MASK,
		Shift:    TMC4671_PID_VELOCITY_TARGET_DIV256_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_PID_VELOCITY_ACTUAL_LSB_FIELD = RegisterField{
		Mask:     TMC4671_PID_VELOCITY_ACTUAL_LSB_MASK,
		Shift:    TMC4671_PID_VELOCITY_ACTUAL_LSB_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_PID_VELOCITY_TARGET_LSB_FIELD = RegisterField{
		Mask:     TMC4671_PID_VELOCITY_TARGET_LSB_MASK,
		Shift:    TMC4671_PID_VELOCITY_TARGET_LSB_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_PID_POSITION_ACTUAL_DIV256_FIELD = RegisterField{
		Mask:     TMC4671_PID_POSITION_ACTUAL_DIV256_MASK,
		Shift:    TMC4671_PID_POSITION_ACTUAL_DIV256_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_PID_POSITION_TARGET_DIV256_FIELD = RegisterField{
		Mask:     TMC4671_PID_POSITION_TARGET_DIV256_MASK,
		Shift:    TMC4671_PID_POSITION_TARGET_DIV256_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_PID_POSITION_ACTUAL_LSB_FIELD = RegisterField{
		Mask:     TMC4671_PID_POSITION_ACTUAL_LSB_MASK,
		Shift:    TMC4671_PID_POSITION_ACTUAL_LSB_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_PID_POSITION_TARGET_LSB_FIELD = RegisterField{
		Mask:     TMC4671_PID_POSITION_TARGET_LSB_MASK,
		Shift:    TMC4671_PID_POSITION_TARGET_LSB_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_FF_VELOCITY_FIELD = RegisterField{
		Mask:     TMC4671_FF_VELOCITY_MASK,
		Shift:    TMC4671_FF_VELOCITY_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_FF_TORQUE_FIELD = RegisterField{
		Mask:     TMC4671_FF_TORQUE_MASK,
		Shift:    TMC4671_FF_TORQUE_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_ACTUAL_VELOCITY_PPTM_FIELD = RegisterField{
		Mask:     TMC4671_ACTUAL_VELOCITY_PPTM_MASK,
		Shift:    TMC4671_ACTUAL_VELOCITY_PPTM_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_HOME_SWITCH_STATUS_FIELD = RegisterField{
		Mask:     TMC4671_HOME_SWITCH_STATUS_MASK,
		Shift:    TMC4671_HOME_SWITCH_STATUS_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: false,
	}
	TMC4671_LEFT_SWITCH_STATUS_FIELD = RegisterField{
		Mask:     TMC4671_LEFT_SWITCH_STATUS_MASK,
		Shift:    TMC4671_LEFT_SWITCH_STATUS_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: false,
	}
	TMC4671_RIGHT_SWITCH_STATUS_FIELD = RegisterField{
		Mask:     TMC4671_RIGHT_SWITCH_STATUS_MASK,
		Shift:    TMC4671_RIGHT_SWITCH_STATUS_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: false,
	}
	TMC4671_HOME_POSITION_FIELD = RegisterField{
		Mask:     TMC4671_HOME_POSITION_MASK,
		Shift:    TMC4671_HOME_POSITION_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_LEFT_POSITION_FIELD = RegisterField{
		Mask:     TMC4671_LEFT_POSITION_MASK,
		Shift:    TMC4671_LEFT_POSITION_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_RIGHT_POSITION_FIELD = RegisterField{
		Mask:     TMC4671_RIGHT_POSITION_MASK,
		Shift:    TMC4671_RIGHT_POSITION_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_ENC_INIT_HALL_STATUS_FIELD = RegisterField{
		Mask:     TMC4671_ENC_INIT_HALL_STATUS_MASK,
		Shift:    TMC4671_ENC_INIT_HALL_STATUS_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: false,
	}
	TMC4671_ENC_INIT_HALL_PHI_E_ABN_OFFSET_FIELD = RegisterField{
		Mask:     TMC4671_ENC_INIT_HALL_PHI_E_ABN_OFFSET_MASK,
		Shift:    TMC4671_ENC_INIT_HALL_PHI_E_ABN_OFFSET_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: false,
	}
	TMC4671_ENC_INIT_HALL_PHI_E_AENC_OFFSET_FIELD = RegisterField{
		Mask:     TMC4671_ENC_INIT_HALL_PHI_E_AENC_OFFSET_MASK,
		Shift:    TMC4671_ENC_INIT_HALL_PHI_E_AENC_OFFSET_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: false,
	}
	TMC4671_ENC_INIT_HALL_PHI_A_AENC_OFFSET_FIELD = RegisterField{
		Mask:     TMC4671_ENC_INIT_HALL_PHI_A_AENC_OFFSET_MASK,
		Shift:    TMC4671_ENC_INIT_HALL_PHI_A_AENC_OFFSET_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: false,
	}
	TMC4671_ENC_INIT_MINI_MOVE_STATUS_FIELD = RegisterField{
		Mask:     TMC4671_ENC_INIT_MINI_MOVE_STATUS_MASK,
		Shift:    TMC4671_ENC_INIT_MINI_MOVE_STATUS_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: false,
	}
	TMC4671_ENC_INIT_MINI_MOVE_U_D_FIELD = RegisterField{
		Mask:     TMC4671_ENC_INIT_MINI_MOVE_U_D_MASK,
		Shift:    TMC4671_ENC_INIT_MINI_MOVE_U_D_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_ENC_INIT_MINI_MOVE_PHI_E_OFFSET_FIELD = RegisterField{
		Mask:     TMC4671_ENC_INIT_MINI_MOVE_PHI_E_OFFSET_MASK,
		Shift:    TMC4671_ENC_INIT_MINI_MOVE_PHI_E_OFFSET_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: false,
	}
	TMC4671_ENC_INIT_MINI_MOVE_PHI_E_FIELD = RegisterField{
		Mask:     TMC4671_ENC_INIT_MINI_MOVE_PHI_E_MASK,
		Shift:    TMC4671_ENC_INIT_MINI_MOVE_PHI_E_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: false,
	}
	TMC4671_SINGLE_PIN_IF_TARGET_TORQUE_FIELD = RegisterField{
		Mask:     TMC4671_SINGLE_PIN_IF_TARGET_TORQUE_MASK,
		Shift:    TMC4671_SINGLE_PIN_IF_TARGET_TORQUE_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_SINGLE_PIN_IF_PWM_DUTY_CYCLE_FIELD = RegisterField{
		Mask:     TMC4671_SINGLE_PIN_IF_PWM_DUTY_CYCLE_MASK,
		Shift:    TMC4671_SINGLE_PIN_IF_PWM_DUTY_CYCLE_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_SINGLE_PIN_IF_TARGET_VELOCITY_FIELD = RegisterField{
		Mask:     TMC4671_SINGLE_PIN_IF_TARGET_VELOCITY_MASK,
		Shift:    TMC4671_SINGLE_PIN_IF_TARGET_VELOCITY_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_SINGLE_PIN_IF_TARGET_POSITION_FIELD = RegisterField{
		Mask:     TMC4671_SINGLE_PIN_IF_TARGET_POSITION_MASK,
		Shift:    TMC4671_SINGLE_PIN_IF_TARGET_POSITION_SHIFT,
		Register: TMC4671_INTERIM_DATA,
		IsSigned: true,
	}
	TMC4671_INTERIM_ADDR_FIELD = RegisterField{
		Mask:     TMC4671_INTERIM_ADDR_MASK,
		Shift:    TMC4671_INTERIM_ADDR_SHIFT,
		Register: TMC4671_INTERIM_ADDR,
		IsSigned: false,
	}
	TMC4671_WATCHDOG_CFG_FIELD = RegisterField{
		Mask:     TMC4671_WATCHDOG_CFG_MASK,
		Shift:    TMC4671_WATCHDOG_CFG_SHIFT,
		Register: TMC4671_WATCHDOG_CFG,
		IsSigned: false,
	}
	TMC4671_ADC_VM_LIMIT_LOW_FIELD = RegisterField{
		Mask:     TMC4671_ADC_VM_LIMIT_LOW_MASK,
		Shift:    TMC4671_ADC_VM_LIMIT_LOW_SHIFT,
		Register: TMC4671_ADC_VM_LIMITS,
		IsSigned: false,
	}
	TMC4671_ADC_VM_LIMIT_HIGH_FIELD = RegisterField{
		Mask:     TMC4671_ADC_VM_LIMIT_HIGH_MASK,
		Shift:    TMC4671_ADC_VM_LIMIT_HIGH_SHIFT,
		Register: TMC4671_ADC_VM_LIMITS,
		IsSigned: false,
	}
	TMC4671_A_OF_ABN_RAW_FIELD = RegisterField{
		Mask:     TMC4671_A_OF_ABN_RAW_MASK,
		Shift:    TMC4671_A_OF_ABN_RAW_SHIFT,
		Register: TMC4671_INPUTS_RAW,
		IsSigned: false,
	}
	TMC4671_B_OF_ABN_RAW_FIELD = RegisterField{
		Mask:     TMC4671_B_OF_ABN_RAW_MASK,
		Shift:    TMC4671_B_OF_ABN_RAW_SHIFT,
		Register: TMC4671_INPUTS_RAW,
		IsSigned: false,
	}
	TMC4671_N_OF_ABN_RAW_FIELD = RegisterField{
		Mask:     TMC4671_N_OF_ABN_RAW_MASK,
		Shift:    TMC4671_N_OF_ABN_RAW_SHIFT,
		Register: TMC4671_INPUTS_RAW,
		IsSigned: false,
	}
	TMC4671_A_OF_ABN_2_RAW_FIELD = RegisterField{
		Mask:     TMC4671_A_OF_ABN_2_RAW_MASK,
		Shift:    TMC4671_A_OF_ABN_2_RAW_SHIFT,
		Register: TMC4671_INPUTS_RAW,
		IsSigned: false,
	}
	TMC4671_B_OF_ABN_2_RAW_FIELD = RegisterField{
		Mask:     TMC4671_B_OF_ABN_2_RAW_MASK,
		Shift:    TMC4671_B_OF_ABN_2_RAW_SHIFT,
		Register: TMC4671_INPUTS_RAW,
		IsSigned: false,
	}
	TMC4671_N_OF_ABN_2_RAW_FIELD = RegisterField{
		Mask:     TMC4671_N_OF_ABN_2_RAW_MASK,
		Shift:    TMC4671_N_OF_ABN_2_RAW_SHIFT,
		Register: TMC4671_INPUTS_RAW,
		IsSigned: false,
	}
	TMC4671_HALL_UX_OF_HALL_RAW_FIELD = RegisterField{
		Mask:     TMC4671_HALL_UX_OF_HALL_RAW_MASK,
		Shift:    TMC4671_HALL_UX_OF_HALL_RAW_SHIFT,
		Register: TMC4671_INPUTS_RAW,
		IsSigned: false,
	}
	TMC4671_HALL_V_OF_HALL_RAW_FIELD = RegisterField{
		Mask:     TMC4671_HALL_V_OF_HALL_RAW_MASK,
		Shift:    TMC4671_HALL_V_OF_HALL_RAW_SHIFT,
		Register: TMC4671_INPUTS_RAW,
		IsSigned: false,
	}
	TMC4671_HALL_WY_OF_HALL_RAW_FIELD = RegisterField{
		Mask:     TMC4671_HALL_WY_OF_HALL_RAW_MASK,
		Shift:    TMC4671_HALL_WY_OF_HALL_RAW_SHIFT,
		Register: TMC4671_INPUTS_RAW,
		IsSigned: false,
	}
	TMC4671_REF_SW_R_RAW_FIELD = RegisterField{
		Mask:     TMC4671_REF_SW_R_RAW_MASK,
		Shift:    TMC4671_REF_SW_R_RAW_SHIFT,
		Register: TMC4671_INPUTS_RAW,
		IsSigned: false,
	}
	TMC4671_REF_SW_L_RAW_FIELD = RegisterField{
		Mask:     TMC4671_REF_SW_L_RAW_MASK,
		Shift:    TMC4671_REF_SW_L_RAW_SHIFT,
		Register: TMC4671_INPUTS_RAW,
		IsSigned: false,
	}
	TMC4671_ENABLE_IN_RAW_FIELD = RegisterField{
		Mask:     TMC4671_ENABLE_IN_RAW_MASK,
		Shift:    TMC4671_ENABLE_IN_RAW_SHIFT,
		Register: TMC4671_INPUTS_RAW,
		IsSigned: false,
	}
	TMC4671_STP_OF_DIRSTP_RAW_FIELD = RegisterField{
		Mask:     TMC4671_STP_OF_DIRSTP_RAW_MASK,
		Shift:    TMC4671_STP_OF_DIRSTP_RAW_SHIFT,
		Register: TMC4671_INPUTS_RAW,
		IsSigned: false,
	}
	TMC4671_DIR_OF_DIRSTP_RAW_FIELD = RegisterField{
		Mask:     TMC4671_DIR_OF_DIRSTP_RAW_MASK,
		Shift:    TMC4671_DIR_OF_DIRSTP_RAW_SHIFT,
		Register: TMC4671_INPUTS_RAW,
		IsSigned: false,
	}
	TMC4671_PWM_IN_RAW_FIELD = RegisterField{
		Mask:     TMC4671_PWM_IN_RAW_MASK,
		Shift:    TMC4671_PWM_IN_RAW_SHIFT,
		Register: TMC4671_INPUTS_RAW,
		IsSigned: false,
	}
	TMC4671_HALL_UX_FILT_FIELD = RegisterField{
		Mask:     TMC4671_HALL_UX_FILT_MASK,
		Shift:    TMC4671_HALL_UX_FILT_SHIFT,
		Register: TMC4671_INPUTS_RAW,
		IsSigned: false,
	}
	TMC4671_HALL_V_FILT_FIELD = RegisterField{
		Mask:     TMC4671_HALL_V_FILT_MASK,
		Shift:    TMC4671_HALL_V_FILT_SHIFT,
		Register: TMC4671_INPUTS_RAW,
		IsSigned: false,
	}
	TMC4671_HALL_WY_FILT_FIELD = RegisterField{
		Mask:     TMC4671_HALL_WY_FILT_MASK,
		Shift:    TMC4671_HALL_WY_FILT_SHIFT,
		Register: TMC4671_INPUTS_RAW,
		IsSigned: false,
	}
	TMC4671_PWM_IDLE_L_RAW_FIELD = RegisterField{
		Mask:     TMC4671_PWM_IDLE_L_RAW_MASK,
		Shift:    TMC4671_PWM_IDLE_L_RAW_SHIFT,
		Register: TMC4671_INPUTS_RAW,
		IsSigned: false,
	}
	TMC4671_OUTPUTS_RAW_0_FIELD = RegisterField{
		Mask:     TMC4671_OUTPUTS_RAW_0_MASK,
		Shift:    TMC4671_OUTPUTS_RAW_0_SHIFT,
		Register: TMC4671_OUTPUTS_RAW,
		IsSigned: false,
	}
	TMC4671_OUTPUTS_RAW_1_FIELD = RegisterField{
		Mask:     TMC4671_OUTPUTS_RAW_1_MASK,
		Shift:    TMC4671_OUTPUTS_RAW_1_SHIFT,
		Register: TMC4671_OUTPUTS_RAW,
		IsSigned: false,
	}
	TMC4671_OUTPUTS_RAW_2_FIELD = RegisterField{
		Mask:     TMC4671_OUTPUTS_RAW_2_MASK,
		Shift:    TMC4671_OUTPUTS_RAW_2_SHIFT,
		Register: TMC4671_OUTPUTS_RAW,
		IsSigned: false,
	}
	TMC4671_OUTPUTS_RAW_3_FIELD = RegisterField{
		Mask:     TMC4671_OUTPUTS_RAW_3_MASK,
		Shift:    TMC4671_OUTPUTS_RAW_3_SHIFT,
		Register: TMC4671_OUTPUTS_RAW,
		IsSigned: false,
	}
	TMC4671_OUTPUTS_RAW_4_FIELD = RegisterField{
		Mask:     TMC4671_OUTPUTS_RAW_4_MASK,
		Shift:    TMC4671_OUTPUTS_RAW_4_SHIFT,
		Register: TMC4671_OUTPUTS_RAW,
		IsSigned: false,
	}
	TMC4671_OUTPUTS_RAW_5_FIELD = RegisterField{
		Mask:     TMC4671_OUTPUTS_RAW_5_MASK,
		Shift:    TMC4671_OUTPUTS_RAW_5_SHIFT,
		Register: TMC4671_OUTPUTS_RAW,
		IsSigned: false,
	}
	TMC4671_OUTPUTS_RAW_6_FIELD = RegisterField{
		Mask:     TMC4671_OUTPUTS_RAW_6_MASK,
		Shift:    TMC4671_OUTPUTS_RAW_6_SHIFT,
		Register: TMC4671_OUTPUTS_RAW,
		IsSigned: false,
	}
	TMC4671_OUTPUTS_RAW_7_FIELD = RegisterField{
		Mask:     TMC4671_OUTPUTS_RAW_7_MASK,
		Shift:    TMC4671_OUTPUTS_RAW_7_SHIFT,
		Register: TMC4671_OUTPUTS_RAW,
		IsSigned: false,
	}
	TMC4671_STEP_WIDTH_FIELD = RegisterField{
		Mask:     TMC4671_STEP_WIDTH_MASK,
		Shift:    TMC4671_STEP_WIDTH_SHIFT,
		Register: TMC4671_STEP_WIDTH,
		IsSigned: true,
	}
	TMC4671_UART_BPS_FIELD = RegisterField{
		Mask:     TMC4671_UART_BPS_MASK,
		Shift:    TMC4671_UART_BPS_SHIFT,
		Register: TMC4671_UART_BPS,
		IsSigned: false,
	}
	TMC4671_ADDR_A_FIELD = RegisterField{
		Mask:     TMC4671_ADDR_A_MASK,
		Shift:    TMC4671_ADDR_A_SHIFT,
		Register: TMC4671_UART_ADDRS,
		IsSigned: false,
	}
	TMC4671_ADDR_B_FIELD = RegisterField{
		Mask:     TMC4671_ADDR_B_MASK,
		Shift:    TMC4671_ADDR_B_SHIFT,
		Register: TMC4671_UART_ADDRS,
		IsSigned: false,
	}
	TMC4671_ADDR_C_FIELD = RegisterField{
		Mask:     TMC4671_ADDR_C_MASK,
		Shift:    TMC4671_ADDR_C_SHIFT,
		Register: TMC4671_UART_ADDRS,
		IsSigned: false,
	}
	TMC4671_ADDR_D_FIELD = RegisterField{
		Mask:     TMC4671_ADDR_D_MASK,
		Shift:    TMC4671_ADDR_D_SHIFT,
		Register: TMC4671_UART_ADDRS,
		IsSigned: false,
	}
	TMC4671_GPIO_DSADCI_CONFIG_0_FIELD = RegisterField{
		Mask:     TMC4671_GPIO_DSADCI_CONFIG_0_MASK,
		Shift:    TMC4671_GPIO_DSADCI_CONFIG_0_SHIFT,
		Register: TMC4671_GPIO_DSADCI_CONFIG,
		IsSigned: false,
	}
	TMC4671_GPIO_DSADCI_CONFIG_1_FIELD = RegisterField{
		Mask:     TMC4671_GPIO_DSADCI_CONFIG_1_MASK,
		Shift:    TMC4671_GPIO_DSADCI_CONFIG_1_SHIFT,
		Register: TMC4671_GPIO_DSADCI_CONFIG,
		IsSigned: false,
	}
	TMC4671_GPIO_DSADCI_CONFIG_2_FIELD = RegisterField{
		Mask:     TMC4671_GPIO_DSADCI_CONFIG_2_MASK,
		Shift:    TMC4671_GPIO_DSADCI_CONFIG_2_SHIFT,
		Register: TMC4671_GPIO_DSADCI_CONFIG,
		IsSigned: false,
	}
	TMC4671_GPIO_DSADCI_CONFIG_3_FIELD = RegisterField{
		Mask:     TMC4671_GPIO_DSADCI_CONFIG_3_MASK,
		Shift:    TMC4671_GPIO_DSADCI_CONFIG_3_SHIFT,
		Register: TMC4671_GPIO_DSADCI_CONFIG,
		IsSigned: false,
	}
	TMC4671_GPIO_DSADCI_CONFIG_4_FIELD = RegisterField{
		Mask:     TMC4671_GPIO_DSADCI_CONFIG_4_MASK,
		Shift:    TMC4671_GPIO_DSADCI_CONFIG_4_SHIFT,
		Register: TMC4671_GPIO_DSADCI_CONFIG,
		IsSigned: false,
	}
	TMC4671_GPIO_DSADCI_CONFIG_5_FIELD = RegisterField{
		Mask:     TMC4671_GPIO_DSADCI_CONFIG_5_MASK,
		Shift:    TMC4671_GPIO_DSADCI_CONFIG_5_SHIFT,
		Register: TMC4671_GPIO_DSADCI_CONFIG,
		IsSigned: false,
	}
	TMC4671_GPIO_DSADCI_CONFIG_6_FIELD = RegisterField{
		Mask:     TMC4671_GPIO_DSADCI_CONFIG_6_MASK,
		Shift:    TMC4671_GPIO_DSADCI_CONFIG_6_SHIFT,
		Register: TMC4671_GPIO_DSADCI_CONFIG,
		IsSigned: false,
	}
	TMC4671_GPO_FIELD = RegisterField{
		Mask:     TMC4671_GPO_MASK,
		Shift:    TMC4671_GPO_SHIFT,
		Register: TMC4671_GPIO_DSADCI_CONFIG,
		IsSigned: false,
	}
	TMC4671_GPI_FIELD = RegisterField{
		Mask:     TMC4671_GPI_MASK,
		Shift:    TMC4671_GPI_SHIFT,
		Register: TMC4671_GPIO_DSADCI_CONFIG,
		IsSigned: false,
	}
	TMC4671_STATUS_FLAGS_0_FIELD = RegisterField{
		Mask:     TMC4671_STATUS_FLAGS_0_MASK,
		Shift:    TMC4671_STATUS_FLAGS_0_SHIFT,
		Register: TMC4671_STATUS_FLAGS,
		IsSigned: false,
	}
	TMC4671_STATUS_FLAGS_1_FIELD = RegisterField{
		Mask:     TMC4671_STATUS_FLAGS_1_MASK,
		Shift:    TMC4671_STATUS_FLAGS_1_SHIFT,
		Register: TMC4671_STATUS_FLAGS,
		IsSigned: false,
	}
	TMC4671_STATUS_FLAGS_2_FIELD = RegisterField{
		Mask:     TMC4671_STATUS_FLAGS_2_MASK,
		Shift:    TMC4671_STATUS_FLAGS_2_SHIFT,
		Register: TMC4671_STATUS_FLAGS,
		IsSigned: false,
	}
	TMC4671_STATUS_FLAGS_3_FIELD = RegisterField{
		Mask:     TMC4671_STATUS_FLAGS_3_MASK,
		Shift:    TMC4671_STATUS_FLAGS_3_SHIFT,
		Register: TMC4671_STATUS_FLAGS,
		IsSigned: false,
	}
	TMC4671_STATUS_FLAGS_4_FIELD = RegisterField{
		Mask:     TMC4671_STATUS_FLAGS_4_MASK,
		Shift:    TMC4671_STATUS_FLAGS_4_SHIFT,
		Register: TMC4671_STATUS_FLAGS,
		IsSigned: false,
	}
	TMC4671_STATUS_FLAGS_5_FIELD = RegisterField{
		Mask:     TMC4671_STATUS_FLAGS_5_MASK,
		Shift:    TMC4671_STATUS_FLAGS_5_SHIFT,
		Register: TMC4671_STATUS_FLAGS,
		IsSigned: false,
	}
	TMC4671_STATUS_FLAGS_6_FIELD = RegisterField{
		Mask:     TMC4671_STATUS_FLAGS_6_MASK,
		Shift:    TMC4671_STATUS_FLAGS_6_SHIFT,
		Register: TMC4671_STATUS_FLAGS,
		IsSigned: false,
	}
	TMC4671_STATUS_FLAGS_7_FIELD = RegisterField{
		Mask:     TMC4671_STATUS_FLAGS_7_MASK,
		Shift:    TMC4671_STATUS_FLAGS_7_SHIFT,
		Register: TMC4671_STATUS_FLAGS,
		IsSigned: false,
	}
	TMC4671_STATUS_FLAGS_8_FIELD = RegisterField{
		Mask:     TMC4671_STATUS_FLAGS_8_MASK,
		Shift:    TMC4671_STATUS_FLAGS_8_SHIFT,
		Register: TMC4671_STATUS_FLAGS,
		IsSigned: false,
	}
	TMC4671_STATUS_FLAGS_9_FIELD = RegisterField{
		Mask:     TMC4671_STATUS_FLAGS_9_MASK,
		Shift:    TMC4671_STATUS_FLAGS_9_SHIFT,
		Register: TMC4671_STATUS_FLAGS,
		IsSigned: false,
	}
	TMC4671_STATUS_FLAGS_10_FIELD = RegisterField{
		Mask:     TMC4671_STATUS_FLAGS_10_MASK,
		Shift:    TMC4671_STATUS_FLAGS_10_SHIFT,
		Register: TMC4671_STATUS_FLAGS,
		IsSigned: false,
	}
	TMC4671_STATUS_FLAGS_11_FIELD = RegisterField{
		Mask:     TMC4671_STATUS_FLAGS_11_MASK,
		Shift:    TMC4671_STATUS_FLAGS_11_SHIFT,
		Register: TMC4671_STATUS_FLAGS,
		IsSigned: false,
	}
	TMC4671_STATUS_FLAGS_12_FIELD = RegisterField{
		Mask:     TMC4671_STATUS_FLAGS_12_MASK,
		Shift:    TMC4671_STATUS_FLAGS_12_SHIFT,
		Register: TMC4671_STATUS_FLAGS,
		IsSigned: false,
	}
	TMC4671_STATUS_FLAGS_13_FIELD = RegisterField{
		Mask:     TMC4671_STATUS_FLAGS_13_MASK,
		Shift:    TMC4671_STATUS_FLAGS_13_SHIFT,
		Register: TMC4671_STATUS_FLAGS,
		IsSigned: false,
	}
	TMC4671_STATUS_FLAGS_14_FIELD = RegisterField{
		Mask:     TMC4671_STATUS_FLAGS_14_MASK,
		Shift:    TMC4671_STATUS_FLAGS_14_SHIFT,
		Register: TMC4671_STATUS_FLAGS,
		IsSigned: false,
	}
	TMC4671_STATUS_FLAGS_15_FIELD = RegisterField{
		Mask:     TMC4671_STATUS_FLAGS_15_MASK,
		Shift:    TMC4671_STATUS_FLAGS_15_SHIFT,
		Register: TMC4671_STATUS_FLAGS,
		IsSigned: false,
	}
	TMC4671_STATUS_FLAGS_16_FIELD = RegisterField{
		Mask:     TMC4671_STATUS_FLAGS_16_MASK,
		Shift:    TMC4671_STATUS_FLAGS_16_SHIFT,
		Register: TMC4671_STATUS_FLAGS,
		IsSigned: false,
	}
	TMC4671_STATUS_FLAGS_17_FIELD = RegisterField{
		Mask:     TMC4671_STATUS_FLAGS_17_MASK,
		Shift:    TMC4671_STATUS_FLAGS_17_SHIFT,
		Register: TMC4671_STATUS_FLAGS,
		IsSigned: false,
	}
	TMC4671_STATUS_FLAGS_18_FIELD = RegisterField{
		Mask:     TMC4671_STATUS_FLAGS_18_MASK,
		Shift:    TMC4671_STATUS_FLAGS_18_SHIFT,
		Register: TMC4671_STATUS_FLAGS,
		IsSigned: false,
	}
	TMC4671_STATUS_FLAGS_19_FIELD = RegisterField{
		Mask:     TMC4671_STATUS_FLAGS_19_MASK,
		Shift:    TMC4671_STATUS_FLAGS_19_SHIFT,
		Register: TMC4671_STATUS_FLAGS,
		IsSigned: false,
	}
	TMC4671_STATUS_FLAGS_20_FIELD = RegisterField{
		Mask:     TMC4671_STATUS_FLAGS_20_MASK,
		Shift:    TMC4671_STATUS_FLAGS_20_SHIFT,
		Register: TMC4671_STATUS_FLAGS,
		IsSigned: false,
	}
	TMC4671_STATUS_FLAGS_21_FIELD = RegisterField{
		Mask:     TMC4671_STATUS_FLAGS_21_MASK,
		Shift:    TMC4671_STATUS_FLAGS_21_SHIFT,
		Register: TMC4671_STATUS_FLAGS,
		IsSigned: false,
	}
	TMC4671_STATUS_FLAGS_22_FIELD = RegisterField{
		Mask:     TMC4671_STATUS_FLAGS_22_MASK,
		Shift:    TMC4671_STATUS_FLAGS_22_SHIFT,
		Register: TMC4671_STATUS_FLAGS,
		IsSigned: false,
	}
	TMC4671_STATUS_FLAGS_23_FIELD = RegisterField{
		Mask:     TMC4671_STATUS_FLAGS_23_MASK,
		Shift:    TMC4671_STATUS_FLAGS_23_SHIFT,
		Register: TMC4671_STATUS_FLAGS,
		IsSigned: false,
	}
	TMC4671_STATUS_FLAGS_24_FIELD = RegisterField{
		Mask:     TMC4671_STATUS_FLAGS_24_MASK,
		Shift:    TMC4671_STATUS_FLAGS_24_SHIFT,
		Register: TMC4671_STATUS_FLAGS,
		IsSigned: false,
	}
	TMC4671_STATUS_FLAGS_25_FIELD = RegisterField{
		Mask:     TMC4671_STATUS_FLAGS_25_MASK,
		Shift:    TMC4671_STATUS_FLAGS_25_SHIFT,
		Register: TMC4671_STATUS_FLAGS,
		IsSigned: false,
	}
	TMC4671_STATUS_FLAGS_26_FIELD = RegisterField{
		Mask:     TMC4671_STATUS_FLAGS_26_MASK,
		Shift:    TMC4671_STATUS_FLAGS_26_SHIFT,
		Register: TMC4671_STATUS_FLAGS,
		IsSigned: false,
	}
	TMC4671_STATUS_FLAGS_27_FIELD = RegisterField{
		Mask:     TMC4671_STATUS_FLAGS_27_MASK,
		Shift:    TMC4671_STATUS_FLAGS_27_SHIFT,
		Register: TMC4671_STATUS_FLAGS,
		IsSigned: false,
	}
	TMC4671_STATUS_FLAGS_28_FIELD = RegisterField{
		Mask:     TMC4671_STATUS_FLAGS_28_MASK,
		Shift:    TMC4671_STATUS_FLAGS_28_SHIFT,
		Register: TMC4671_STATUS_FLAGS,
		IsSigned: false,
	}
	TMC4671_STATUS_FLAGS_29_FIELD = RegisterField{
		Mask:     TMC4671_STATUS_FLAGS_29_MASK,
		Shift:    TMC4671_STATUS_FLAGS_29_SHIFT,
		Register: TMC4671_STATUS_FLAGS,
		IsSigned: false,
	}
	TMC4671_STATUS_FLAGS_30_FIELD = RegisterField{
		Mask:     TMC4671_STATUS_FLAGS_30_MASK,
		Shift:    TMC4671_STATUS_FLAGS_30_SHIFT,
		Register: TMC4671_STATUS_FLAGS,
		IsSigned: false,
	}
	TMC4671_STATUS_FLAGS_31_FIELD = RegisterField{
		Mask:     TMC4671_STATUS_FLAGS_31_MASK,
		Shift:    TMC4671_STATUS_FLAGS_31_SHIFT,
		Register: TMC4671_STATUS_FLAGS,
		IsSigned: false,
	}
	TMC4671_WARNING_MASK_FIELD = RegisterField{
		Mask:     TMC4671_WARNING_MASK_MASK,
		Shift:    TMC4671_WARNING_MASK_SHIFT,
		Register: TMC4671_STATUS_MASK,
		IsSigned: false,
	}
)
