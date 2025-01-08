package facts

type VersionInfo struct {
	HeaderStr       string `json:"header_str"`
	BiosVerStr      string `json:"bios_ver_str"`
	KickstartVerStr string `json:"kickstart_ver_str"`
	NxosVerStr      string `json:"nxos_ver_str"`
	BiosCmplTime    string `json:"bios_cmpl_time"`
	KickFileName    string `json:"kick_file_name"`
	NxosFileName    string `json:"nxos_file_name"`
	KickCmplTime    string `json:"kick_cmpl_time"`
	NxosCmplTime    string `json:"nxos_cmpl_time"`
	KickTmstmp      string `json:"kick_tmstmp"`
	NxosTmstmp      string `json:"nxos_tmstmp"`
	ChassisID       string `json:"chassis_id"`
	CPUName         string `json:"cpu_name"`
	Memory          int64  `json:"memory,string,omitempty"`
	MemType         string `json:"mem_type"`
	ProcBoardID     string `json:"proc_board_id"`
	HostName        string `json:"host_name"`
	BootflashSize   int64  `json:"bootflash_size,string,omitempty"`
	KernUptmDays    int64  `json:"kern_uptm_days,string,omitempty"`
	KernUptmHrs     int64  `json:"kern_uptm_hrs,string,omitempty"`
	KernUptmMins    int64  `json:"kern_uptm_mins,string,omitempty"`
	KernUptmSecs    int64  `json:"kern_uptm_secs,string,omitempty"`
	RrUsecs         int64  `json:"rr_usecs,string,omitempty"`
	RrCtime         string `json:"rr_ctime"`
	RrReason        string `json:"rr_reason"`
	RrSysVer        string `json:"rr_sys_ver"`
	RrService       string `json:"rr_service"`
	Plugins         string `json:"plugins"`
	Manufacturer    string `json:"manufacturer"`
}

type VersionFact struct {
	Version string
}

type MemoryFact struct {
	Type  string
	Total float64
	Used  float64
	Free  float64
}

type CPUFact struct {
	FiveSeconds float64
	Interrupts  float64
	OneMinute   float64
	FiveMinutes float64
}
