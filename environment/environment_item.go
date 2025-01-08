package environment

// Define an enum using iota
type EnvironmentItemType int

const (
	Temp EnvironmentItemType = iota
	PowerIn
	PowerOut
	PowerStatus
	PowerCapacity
	PowerMode
	FanStatus
	FanSpeed
)

type EnvironmentItem struct {
	Name   string
	Detail string
	Status string
	Type   EnvironmentItemType
	Value  float64
}

// Root structure
type RootEnvironment struct {
	FanDetails FanDetails `json:"fandetails"`
	PowerSup   PowerSup   `json:"powersup"`
	TempInfo   TempInfo   `json:"TABLE_tempinfo"`
}

type FanDetails struct {
	FanInfo         FanInfo      `json:"TABLE_faninfo"`
	FanZoneSpeed    FanZoneSpeed `json:"TABLE_fan_zone_speed"`
	FanFilterStatus string       `json:"fan_filter_status"`
}

type FanInfo struct {
	RowFanInfo []FanInfoRow `json:"ROW_faninfo"`
}

type FanInfoRow struct {
	FanName   string `json:"fanname"`
	FanModel  string `json:"fanmodel"`
	FanHwVer  string `json:"fanhwver"`
	FanDir    string `json:"fandir"`
	FanStatus string `json:"fanstatus"`
}

type FanZoneSpeed struct {
	RowFanZoneSpeed FanZoneSpeedRow `json:"ROW_fan_zone_speed"`
}

type FanZoneSpeedRow struct {
	Zone      string `json:"zone"`
	ZoneSpeed string `json:"zonespeed"`
}

type PowerSup struct {
	VoltageLevel string       `json:"voltage_level"`
	PSInfo       PSInfo       `json:"TABLE_psinfo"`
	PowerSummary PowerSummary `json:"power_summary"`
}

type PSInfo struct {
	RowPSInfo []PSInfoRow `json:"ROW_psinfo"`
}

type PSInfoRow struct {
	PSNum       string `json:"psnum"`
	PSModel     string `json:"psmodel"`
	ActualOut   string `json:"actual_out"`
	ActualInput string `json:"actual_input"`
	TotalCapa   string `json:"tot_capa"`
	PSStatus    string `json:"ps_status"`
}

type PowerSummary struct {
	PSRedunMode           string `json:"ps_redun_mode"`
	PSOperMode            string `json:"ps_oper_mode"`
	TotalPowerCapacity    string `json:"tot_pow_capacity"`
	TotalGridACapacity    string `json:"tot_gridA_capacity"`
	TotalGridBCapacity    string `json:"tot_gridB_capacity"`
	CumulativePower       string `json:"cumulative_power"`
	TotalPowerOutActual   string `json:"tot_pow_out_actual_draw"`
	TotalPowerInputActual string `json:"tot_pow_input_actual_draw"`
	TotalPowerAllocBudget string `json:"tot_pow_alloc_budgeted"`
	AvailablePower        string `json:"available_pow"`
}

type TempInfo struct {
	RowTempInfo []TempInfoRow `json:"ROW_tempinfo"`
}

type TempInfoRow struct {
	TempMod     string `json:"tempmod"`
	Sensor      string `json:"sensor"`
	MajThres    string `json:"majthres"`
	MinThres    string `json:"minthres"`
	CurTemp     string `json:"curtemp"`
	AlarmStatus string `json:"alarmstatus"`
}
