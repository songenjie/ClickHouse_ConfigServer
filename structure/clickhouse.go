package structure

type ClickHouse struct {
	Cluster   string `json:"cluster"`
	Zookeeper struct {
		Port int      `json:"port"`
		Ips  []string `json:"ips"`
	} `json:"zookeeper"`
	Shard Shard `json:"shard"`
	User  struct {
		LoadBalancing             string `json:"load_balancing"`
		MaxMemoryUsage            string `json:"max_memory_usage"`
		MaxMemoryUsageForAllQuery string `json:"max_memory_usage_for_all_query"`
		SkipUnabvailiableShard    bool   `json:"skip_unabvailiable_shard"`
	} `json:"user"`
	Replicas []*Replica `json:"replicas"`
}

type Shard struct {
	Port   []int    `json:"port"`
	Groups []*Group `json:"groups"`
}

/*
Num1:shard = GroupNumber*3+1
Num2:shard = GroupNumber*3+2
Num3:shard = GroupNumber*3+3
*/

type Group struct {
	GroupNumber      int  `json:"group_number"`
	InterReplication bool `json:"inter_replication"`
	Weight           int  `json:"weight"`
	Ips              struct {
		Num1 string `json:"1"`
		Num2 string `json:"2"`
		Num3 string `json:"3"`
	} `json:"ips"`
}

type Replica struct {
	HTTPPort            string `json:"http_port"`
	TCPPort             string `json:"tcp_port"`
	MysqlPort           string `json:"mysql_port"`
	InterserverHTTPPort string `json:"interserver_http_port"`
	PrometheusPort      string `json:"prometheus_port"`
	LogDir              string `json:"log_dir"`
	MetaDataDir         string `json:"meta_data_dir"`
	SslDir              string `json:"ssl_dir"`
	TmpDir              string `json:"tmp_dir"`
	ConfDir             string `json:"conf_dir"`
	Replica             int    `json:"replica"`
	DataDir             struct {
		Hot struct {
			Diska string `json:"diska"`
			Diskb string `json:"diskb"`
		} `json:"hot"`
		Cold struct {
			Diskc string `json:"diskc"`
			Diskd string `json:"diskd"`
		} `json:"cold"`
		MaxDataPartSizeBytes string `json:"max_data_part_size_bytes"`
		MoveFactor           string `json:"move_factor"`
	} `json:"data_dir"`
}