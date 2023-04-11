package model

import "time"

type TbNodeDeviceRecord struct {
	CreateTime    time.Time `json:"create_time"`                             // 时间戳
	CreateTimeStr int64     `json:"create_time" db:"create_time"`            // 时间戳
	ConId         string    `json:"con_id"  db:"con_id"`                     // 归属集中器ID
	Online        int       `json:"online"  db:"online"`                     // 在线状态：1-在线 0-离线
	AlarmState    int       `json:"alarm_state"  db:"alarm_state"`           // 报警状态位(整数，高字节在前)
	PowerState    int       `json:"power_status"  db:"power_status"`         // 电源通电状态：1-通电 0-断电
	Brightness    int       `json:"brightness"  db:"brightness"`             // 亮度 0~100
	Temperature   int       `json:"temperature"  db:"temperature"`           // 温度
	PowerF        int       `json:"pf"  db:"pf"`                             // 功率因数
	Kwh           float64   `json:"kwh" gorm:"kwh" db:"kwh"`                 // 电能总电量 0.001kwh
	InputVoltage  float64   `json:"input_voltage"  db:"input_voltage"`       // 输入电压V
	InputCurrent  float64   `json:"input_current"  db:"input_current"`       // 输入电流A
	InputPower    float64   `json:"input_power"  db:"input_power"`           // 输入功率P
	LeakCurrent   float64   `json:"leak_current"  db:"leak_current"`         // 漏电流值 0.1mA
	RatedPower    int       `json:"rated_power" db:"rated_power"`            // 电源额定功率
	PoleLeakV     float64   `json:"pole_leakage_vol"  db:"pole_leakage_vol"` // 灯杆漏电电压
	PoleLeakC     float64   `json:"pole_leakage_cur"  db:"pole_leakage_cur"` // 灯杆漏电电流
	Tags          TbNodeDeviceRecordTags
}

type TbNodeDeviceRecordTags struct {
	NodeId   string `json:"node_id"  db:"node_id"`     // 节点id
	LedNo    int    `json:"led_no"  db:"led_no"`       // LED编号(从1开始)
	NodeType string `json:"node_type"  db:"node_type"` // 节点类型
}

type TbNodeDeviceRecordDb struct {
	CreateTime   time.Time `json:"create_time"  db:"create_time"`           // 时间戳
	ConId        string    `json:"con_id"  db:"con_id"`                     // 归属集中器ID
	Online       int       `json:"online"  db:"online"`                     // 在线状态：1-在线 0-离线
	AlarmState   int       `json:"alarm_state"  db:"alarm_state"`           // 报警状态位(整数，高字节在前)
	PowerState   int       `json:"power_status"  db:"power_status"`         // 电源通电状态：1-通电 0-断电
	Brightness   int       `json:"brightness"  db:"brightness"`             // 亮度 0~100
	Temperature  int       `json:"temperature"  db:"temperature"`           // 温度
	PowerF       int       `json:"pf"  db:"pf"`                             // 功率因数
	Kwh          float64   `json:"kwh" gorm:"kwh" db:"kwh"`                 // 电能总电量 0.001kwh
	InputVoltage float64   `json:"input_voltage"  db:"input_voltage"`       // 输入电压V
	InputCurrent float64   `json:"input_current"  db:"input_current"`       // 输入电流A
	InputPower   float64   `json:"input_power"  db:"input_power"`           // 输入功率P
	LeakCurrent  float64   `json:"leak_current"  db:"leak_current"`         // 漏电流值 0.1mA
	RatedPower   int       `json:"rated_power" db:"rated_power"`            // 电源额定功率
	PoleLeakV    float64   `json:"pole_leakage_vol"  db:"pole_leakage_vol"` // 灯杆漏电电压
	PoleLeakC    float64   `json:"pole_leakage_cur"  db:"pole_leakage_cur"` // 灯杆漏电电流
	NodeId       string    `json:"node_id"  db:"node_id"`                   // 节点id
	LedNo        int       `json:"led_no"  db:"led_no"`                     // LED编号(从1开始)
	NodeType     string    `json:"node_type"  db:"node_type"`               // 节点类型
}
