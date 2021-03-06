package unidev

// UAP is a Unifi Access Point.
type UAP struct {
	/* This was auto generated and then slowly edited by hand
	   to get all the data types right and graphable.
	   No ones feelings will be hurt if you want to break this
		 up into multiple structs, and/or make it better in general.
	*/
	ID           string  `json:"_id"`
	UUptime      float64 `json:"_uptime"`
	AdoptIP      string  `json:"adopt_ip,omitempty"`
	AdoptURL     string  `json:"adopt_url,omitempty"`
	Adopted      bool    `json:"adopted"`
	AntennaTable []struct {
		ID        float64 `json:"id"`
		Name      string  `json:"name"`
		Wifi0Gain float64 `json:"wifi0_gain"`
		Wifi1Gain float64 `json:"wifi1_gain"`
	} `json:"antenna_table"`
	BandsteeringMode string  `json:"bandsteering_mode,omitempty"`
	BoardRev         int     `json:"board_rev"`
	Bytes            float64 `json:"bytes"`
	BytesD           float64 `json:"bytes-d"`
	BytesR           float64 `json:"bytes-r"`
	Cfgversion       string  `json:"cfgversion"`
	ConfigNetwork    struct {
		IP   string `json:"ip"`
		Type string `json:"type"`
	} `json:"config_network"`
	ConnectRequestIP   string        `json:"connect_request_ip"`
	ConnectRequestPort string        `json:"connect_request_port"`
	ConsideredLostAt   float64       `json:"considered_lost_at"`
	CountrycodeTable   []float64     `json:"countrycode_table"`
	Default            bool          `json:"default,omitempty"`
	DeviceID           string        `json:"device_id"`
	DiscoveredVia      string        `json:"discovered_via,omitempty"`
	DownlinkTable      []interface{} `json:"downlink_table"`
	EthernetTable      []struct {
		Mac     string  `json:"mac"`
		Name    string  `json:"name"`
		NumPort float64 `json:"num_port"`
	} `json:"ethernet_table"`
	FwCaps          int     `json:"fw_caps"`
	GuestNumSta     int     `json:"guest-num_sta"`
	GuestToken      string  `json:"guest_token"`
	HasEth1         bool    `json:"has_eth1"`
	HasSpeaker      bool    `json:"has_speaker"`
	InformIP        string  `json:"inform_ip"`
	InformURL       string  `json:"inform_url"`
	IP              string  `json:"ip"`
	Isolated        bool    `json:"isolated"`
	KnownCfgversion string  `json:"known_cfgversion"`
	LastSeen        float64 `json:"last_seen"`
	LastUplink      struct {
		UplinkMac        string `json:"uplink_mac"`
		UplinkRemotePort int    `json:"uplink_remote_port"`
	} `json:"last_uplink"`
	LedOverride         string  `json:"led_override"`
	Locating            bool    `json:"locating"`
	Mac                 string  `json:"mac"`
	Model               string  `json:"model"`
	Name                string  `json:"name"`
	NextHeartbeatAt     float64 `json:"next_heartbeat_at"`
	NumSta              float64 `json:"num_sta"`
	OutdoorModeOverride string  `json:"outdoor_mode_override"`
	PortTable           []struct {
		AggregatedBy bool    `json:"aggregated_by"`
		AttrNoEdit   bool    `json:"attr_no_edit,omitempty"`
		Autoneg      bool    `json:"autoneg"`
		BytesR       float64 `json:"bytes-r"`
		Enable       bool    `json:"enable"`
		FlowctrlRx   bool    `json:"flowctrl_rx"`
		FlowctrlTx   bool    `json:"flowctrl_tx"`
		FullDuplex   bool    `json:"full_duplex"`
		IsUplink     bool    `json:"is_uplink"`
		Jumbo        bool    `json:"jumbo"`
		MacTable     []struct {
			Age    float64 `json:"age"`
			Mac    string  `json:"mac"`
			Static bool    `json:"static"`
			Uptime float64 `json:"uptime"`
			Vlan   float64 `json:"vlan"`
		} `json:"mac_table"`
		Masked    bool    `json:"masked"`
		Media     string  `json:"media"`
		Name      string  `json:"name"`
		OpMode    string  `json:"op_mode"`
		PoeCaps   float64 `json:"poe_caps"`
		PortDelta struct {
			RxBytes   float64 `json:"rx_bytes"`
			RxPackets float64 `json:"rx_packets"`
			TimeDelta float64 `json:"time_delta"`
			TxBytes   float64 `json:"tx_bytes"`
			TxPackets float64 `json:"tx_packets"`
		} `json:"port_delta"`
		PortIdx     float64 `json:"port_idx"`
		PortPoe     bool    `json:"port_poe"`
		PortconfID  string  `json:"portconf_id"`
		RxBroadcast float64 `json:"rx_broadcast"`
		RxBytes     float64 `json:"rx_bytes"`
		RxBytesR    float64 `json:"rx_bytes-r"`
		RxDropped   float64 `json:"rx_dropped"`
		RxErrors    float64 `json:"rx_errors"`
		RxMulticast float64 `json:"rx_multicast"`
		RxPackets   float64 `json:"rx_packets"`
		Speed       float64 `json:"speed"`
		StpPathcost float64 `json:"stp_pathcost"`
		StpState    string  `json:"stp_state"`
		TxBroadcast float64 `json:"tx_broadcast"`
		TxBytes     float64 `json:"tx_bytes"`
		TxBytesR    float64 `json:"tx_bytes-r"`
		TxDropped   float64 `json:"tx_dropped"`
		TxErrors    float64 `json:"tx_errors"`
		TxMulticast float64 `json:"tx_multicast"`
		TxPackets   float64 `json:"tx_packets"`
		Up          bool    `json:"up"`
	} `json:"port_table"`
	RadioTable []struct {
		BuiltinAntGain     float64 `json:"builtin_ant_gain"`
		BuiltinAntenna     bool    `json:"builtin_antenna"`
		Channel            FlexInt `json:"channel"`
		CurrentAntennaGain float64 `json:"current_antenna_gain"`
		Ht                 string  `json:"ht"`
		MaxTxpower         float64 `json:"max_txpower"`
		MinRssiEnabled     bool    `json:"min_rssi_enabled"`
		MinTxpower         float64 `json:"min_txpower"`
		Name               string  `json:"name"`
		Nss                float64 `json:"nss"`
		Radio              string  `json:"radio"`
		RadioCaps          float64 `json:"radio_caps"`
		TxPower            FlexInt `json:"tx_power"`
		TxPowerMode        string  `json:"tx_power_mode"`
		WlangroupID        string  `json:"wlangroup_id"`
		HasDfs             bool    `json:"has_dfs,omitempty"`
		HasFccdfs          bool    `json:"has_fccdfs,omitempty"`
		Is11Ac             bool    `json:"is_11ac,omitempty"`
	} `json:"radio_table"`
	RadioTableStats []struct {
		AstBeXmit   float64 `json:"ast_be_xmit"`
		AstCst      float64 `json:"ast_cst"`
		AstTxto     float64 `json:"ast_txto"`
		Channel     float64 `json:"channel"`
		CuSelfRx    float64 `json:"cu_self_rx"`
		CuSelfTx    float64 `json:"cu_self_tx"`
		CuTotal     float64 `json:"cu_total"`
		Extchannel  float64 `json:"extchannel"`
		Gain        float64 `json:"gain"`
		GuestNumSta float64 `json:"guest-num_sta"`
		Name        string  `json:"name"`
		NumSta      float64 `json:"num_sta"`
		Radio       string  `json:"radio"`
		State       string  `json:"state"`
		TxPackets   float64 `json:"tx_packets"`
		TxPower     float64 `json:"tx_power"`
		TxRetries   float64 `json:"tx_retries"`
		UserNumSta  float64 `json:"user-num_sta"`
	} `json:"radio_table_stats"`
	Rollupgrade      bool          `json:"rollupgrade"`
	RxBytes          float64       `json:"rx_bytes"`
	RxBytesD         float64       `json:"rx_bytes-d"`
	ScanRadioTable   []interface{} `json:"scan_radio_table"`
	Scanning         bool          `json:"scanning"`
	Serial           string        `json:"serial"`
	SiteID           string        `json:"site_id"`
	SpectrumScanning bool          `json:"spectrum_scanning"`
	SSHSessionTable  []interface{} `json:"ssh_session_table"`
	Stat             struct {
		Ap                 string  `json:"ap"`
		Bytes              float64 `json:"bytes"`
		Datetime           string  `json:"datetime"`
		Duration           float64 `json:"duration"`
		GuestRxBytes       float64 `json:"guest-rx_bytes"`
		GuestRxCrypts      float64 `json:"guest-rx_crypts"`
		GuestRxDropped     float64 `json:"guest-rx_dropped"`
		GuestRxErrors      float64 `json:"guest-rx_errors"`
		GuestRxFrags       float64 `json:"guest-rx_frags"`
		GuestRxPackets     float64 `json:"guest-rx_packets"`
		GuestTxBytes       float64 `json:"guest-tx_bytes"`
		GuestTxDropped     float64 `json:"guest-tx_dropped"`
		GuestTxErrors      float64 `json:"guest-tx_errors"`
		GuestTxPackets     float64 `json:"guest-tx_packets"`
		GuestTxRetries     float64 `json:"guest-tx_retries"`
		O                  string  `json:"o"`
		Oid                string  `json:"oid"`
		Port1RxBroadcast   float64 `json:"port_1-rx_broadcast"`
		Port1RxBytes       float64 `json:"port_1-rx_bytes"`
		Port1RxMulticast   float64 `json:"port_1-rx_multicast"`
		Port1RxPackets     float64 `json:"port_1-rx_packets"`
		Port1TxBroadcast   float64 `json:"port_1-tx_broadcast"`
		Port1TxBytes       float64 `json:"port_1-tx_bytes"`
		Port1TxMulticast   float64 `json:"port_1-tx_multicast"`
		Port1TxPackets     float64 `json:"port_1-tx_packets"`
		RxBytes            float64 `json:"rx_bytes"`
		RxCrypts           float64 `json:"rx_crypts"`
		RxDropped          float64 `json:"rx_dropped"`
		RxErrors           float64 `json:"rx_errors"`
		RxFrags            float64 `json:"rx_frags"`
		RxPackets          float64 `json:"rx_packets"`
		SiteID             string  `json:"site_id"`
		Time               float64 `json:"time"`
		TxBytes            float64 `json:"tx_bytes"`
		TxDropped          float64 `json:"tx_dropped"`
		TxErrors           float64 `json:"tx_errors"`
		TxPackets          float64 `json:"tx_packets"`
		TxRetries          float64 `json:"tx_retries"`
		UserRxBytes        float64 `json:"user-rx_bytes"`
		UserRxCrypts       float64 `json:"user-rx_crypts"`
		UserRxDropped      float64 `json:"user-rx_dropped"`
		UserRxErrors       float64 `json:"user-rx_errors"`
		UserRxFrags        float64 `json:"user-rx_frags"`
		UserRxPackets      float64 `json:"user-rx_packets"`
		UserTxBytes        float64 `json:"user-tx_bytes"`
		UserTxDropped      float64 `json:"user-tx_dropped"`
		UserTxErrors       float64 `json:"user-tx_errors"`
		UserTxPackets      float64 `json:"user-tx_packets"`
		UserTxRetries      float64 `json:"user-tx_retries"`
		UserWifi0RxBytes   float64 `json:"user-wifi0-rx_bytes"`
		UserWifi0RxCrypts  float64 `json:"user-wifi0-rx_crypts"`
		UserWifi0RxDropped float64 `json:"user-wifi0-rx_dropped"`
		UserWifi0RxErrors  float64 `json:"user-wifi0-rx_errors"`
		UserWifi0RxFrags   float64 `json:"user-wifi0-rx_frags"`
		UserWifi0RxPackets float64 `json:"user-wifi0-rx_packets"`
		UserWifi0TxBytes   float64 `json:"user-wifi0-tx_bytes"`
		UserWifi0TxDropped float64 `json:"user-wifi0-tx_dropped"`
		UserWifi0TxErrors  float64 `json:"user-wifi0-tx_errors"`
		UserWifi0TxPackets float64 `json:"user-wifi0-tx_packets"`
		UserWifi0TxRetries float64 `json:"user-wifi0-tx_retries"`
		UserWifi1RxBytes   float64 `json:"user-wifi1-rx_bytes"`
		UserWifi1RxCrypts  float64 `json:"user-wifi1-rx_crypts"`
		UserWifi1RxDropped float64 `json:"user-wifi1-rx_dropped"`
		UserWifi1RxErrors  float64 `json:"user-wifi1-rx_errors"`
		UserWifi1RxFrags   float64 `json:"user-wifi1-rx_frags"`
		UserWifi1RxPackets float64 `json:"user-wifi1-rx_packets"`
		UserWifi1TxBytes   float64 `json:"user-wifi1-tx_bytes"`
		UserWifi1TxDropped float64 `json:"user-wifi1-tx_dropped"`
		UserWifi1TxErrors  float64 `json:"user-wifi1-tx_errors"`
		UserWifi1TxPackets float64 `json:"user-wifi1-tx_packets"`
		UserWifi1TxRetries float64 `json:"user-wifi1-tx_retries"`
		Wifi0RxBytes       float64 `json:"wifi0-rx_bytes"`
		Wifi0RxCrypts      float64 `json:"wifi0-rx_crypts"`
		Wifi0RxDropped     float64 `json:"wifi0-rx_dropped"`
		Wifi0RxErrors      float64 `json:"wifi0-rx_errors"`
		Wifi0RxFrags       float64 `json:"wifi0-rx_frags"`
		Wifi0RxPackets     float64 `json:"wifi0-rx_packets"`
		Wifi0TxBytes       float64 `json:"wifi0-tx_bytes"`
		Wifi0TxDropped     float64 `json:"wifi0-tx_dropped"`
		Wifi0TxErrors      float64 `json:"wifi0-tx_errors"`
		Wifi0TxPackets     float64 `json:"wifi0-tx_packets"`
		Wifi0TxRetries     float64 `json:"wifi0-tx_retries"`
		Wifi1RxBytes       float64 `json:"wifi1-rx_bytes"`
		Wifi1RxCrypts      float64 `json:"wifi1-rx_crypts"`
		Wifi1RxDropped     float64 `json:"wifi1-rx_dropped"`
		Wifi1RxErrors      float64 `json:"wifi1-rx_errors"`
		Wifi1RxFrags       float64 `json:"wifi1-rx_frags"`
		Wifi1RxPackets     float64 `json:"wifi1-rx_packets"`
		Wifi1TxBytes       float64 `json:"wifi1-tx_bytes"`
		Wifi1TxDropped     float64 `json:"wifi1-tx_dropped"`
		Wifi1TxErrors      float64 `json:"wifi1-tx_errors"`
		Wifi1TxPackets     float64 `json:"wifi1-tx_packets"`
		Wifi1TxRetries     float64 `json:"wifi1-tx_retries"`
	} `json:"stat"`
	State    int `json:"state"`
	SysStats struct {
		Loadavg1  float64 `json:"loadavg_1,string"`
		Loadavg15 float64 `json:"loadavg_15,string"`
		Loadavg5  float64 `json:"loadavg_5,string"`
		MemBuffer float64 `json:"mem_buffer"`
		MemTotal  float64 `json:"mem_total"`
		MemUsed   float64 `json:"mem_used"`
	} `json:"sys_stats"`
	SystemStats struct {
		CPU    float64 `json:"cpu,string"`
		Mem    float64 `json:"mem,string"`
		Uptime float64 `json:"uptime,string"`
	} `json:"system-stats"`
	TxBytes    float64 `json:"tx_bytes"`
	TxBytesD   float64 `json:"tx_bytes-d"`
	Type       string  `json:"type"`
	Upgradable bool    `json:"upgradable"`
	Uplink     struct {
		FullDuplex       bool    `json:"full_duplex"`
		IP               string  `json:"ip"`
		Mac              string  `json:"mac"`
		MaxSpeed         int     `json:"max_speed"`
		MaxVlan          int     `json:"max_vlan"`
		Media            string  `json:"media"`
		Name             string  `json:"name"`
		Netmask          string  `json:"netmask"`
		NumPort          int     `json:"num_port"`
		RxBytes          float64 `json:"rx_bytes"`
		RxBytesR         float64 `json:"rx_bytes-r"`
		RxDropped        float64 `json:"rx_dropped"`
		RxErrors         float64 `json:"rx_errors"`
		RxMulticast      float64 `json:"rx_multicast"`
		RxPackets        float64 `json:"rx_packets"`
		Speed            float64 `json:"speed"`
		TxBytes          float64 `json:"tx_bytes"`
		TxBytesR         float64 `json:"tx_bytes-r"`
		TxDropped        float64 `json:"tx_dropped"`
		TxErrors         float64 `json:"tx_errors"`
		TxPackets        float64 `json:"tx_packets"`
		Type             string  `json:"type"`
		Up               bool    `json:"up"`
		UplinkMac        string  `json:"uplink_mac"`
		UplinkRemotePort int     `json:"uplink_remote_port"`
	} `json:"uplink"`
	UplinkTable []interface{} `json:"uplink_table"`
	Uptime      FlexInt       `json:"uptime"`
	UserNumSta  int           `json:"user-num_sta"`
	VapTable    []struct {
		ApMac               string  `json:"ap_mac"`
		Bssid               string  `json:"bssid"`
		Ccq                 int     `json:"ccq"`
		Channel             int     `json:"channel"`
		Essid               string  `json:"essid"`
		Extchannel          int     `json:"extchannel"`
		ID                  string  `json:"id"`
		IsGuest             bool    `json:"is_guest"`
		IsWep               bool    `json:"is_wep"`
		MacFilterRejections int     `json:"mac_filter_rejections"`
		MapID               string  `json:"map_id"`
		Name                string  `json:"name"`
		NumSta              int     `json:"num_sta"`
		Radio               string  `json:"radio"`
		RadioName           string  `json:"radio_name"`
		RxBytes             float64 `json:"rx_bytes"`
		RxCrypts            float64 `json:"rx_crypts"`
		RxDropped           float64 `json:"rx_dropped"`
		RxErrors            float64 `json:"rx_errors"`
		RxFrags             float64 `json:"rx_frags"`
		RxNwids             float64 `json:"rx_nwids"`
		RxPackets           float64 `json:"rx_packets"`
		SiteID              string  `json:"site_id"`
		State               string  `json:"state"`
		T                   string  `json:"t"`
		TxBytes             float64 `json:"tx_bytes"`
		TxDropped           float64 `json:"tx_dropped"`
		TxErrors            float64 `json:"tx_errors"`
		TxLatencyAvg        float64 `json:"tx_latency_avg"`
		TxLatencyMax        float64 `json:"tx_latency_max"`
		TxLatencyMin        float64 `json:"tx_latency_min"`
		TxPackets           float64 `json:"tx_packets"`
		TxPower             int     `json:"tx_power"`
		TxRetries           int     `json:"tx_retries"`
		Up                  bool    `json:"up"`
		Usage               string  `json:"usage"`
		WlanconfID          string  `json:"wlanconf_id"`
	} `json:"vap_table"`
	Version             string        `json:"version"`
	VersionIncompatible bool          `json:"version_incompatible"`
	VwireEnabled        bool          `json:"vwireEnabled"`
	VwireTable          []interface{} `json:"vwire_table"`
	VwireVapTable       []struct {
		Bssid     string `json:"bssid"`
		Radio     string `json:"radio"`
		RadioName string `json:"radio_name"`
		State     string `json:"state"`
	} `json:"vwire_vap_table"`
	WifiCaps int `json:"wifi_caps"`
}
