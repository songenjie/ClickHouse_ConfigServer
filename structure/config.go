package structure

import (
	"encoding/xml"
)

type Config struct {
	XMLName xml.Name `xml:"yandex"`
	Text    string   `xml:",chardata"`
	Logger  struct {
		Text     string `xml:",chardata"`
		Level    string `xml:"level"`
		Log      string `xml:"log"`
		Errorlog string `xml:"errorlog"`
		Size     string `xml:"size"`
		Count    string `xml:"count"`
	} `xml:"logger"`
	HTTPPort            string `xml:"http_port"`
	TcpPort             string `xml:"tcp_port"`
	MysqlPort           string `xml:"mysql_port"`
	InterserverHTTPPort string `xml:"interserver_http_port"`
	OpenSSL             struct {
		Text   string `xml:",chardata"`
		Server struct {
			Text                string `xml:",chardata"`
			CertificateFile     string `xml:"certificateFile"`
			PrivateKeyFile      string `xml:"privateKeyFile"`
			DhParamsFile        string `xml:"dhParamsFile"`
			VerificationMode    string `xml:"verificationMode"`
			LoadDefaultCAFile   string `xml:"loadDefaultCAFile"`
			CacheSessions       string `xml:"cacheSessions"`
			DisableProtocols    string `xml:"disableProtocols"`
			PreferServerCiphers string `xml:"preferServerCiphers"`
		} `xml:"server"`
		Client struct {
			Text                      string `xml:",chardata"`
			LoadDefaultCAFile         string `xml:"loadDefaultCAFile"`
			CacheSessions             string `xml:"cacheSessions"`
			DisableProtocols          string `xml:"disableProtocols"`
			PreferServerCiphers       string `xml:"preferServerCiphers"`
			InvalidCertificateHandler struct {
				Text string `xml:",chardata"`
				Name string `xml:"name"`
			} `xml:"invalidCertificateHandler"`
		} `xml:"client"`
	} `xml:"openSSL"`
	ListenHost            string `xml:"listen_host"`
	MaxConnections        string `xml:"max_connections"`
	KeepAliveTimeout      string `xml:"keep_alive_timeout"`
	MaxConcurrentQueries  string `xml:"max_concurrent_queries"`
	UncompressedCacheSize string `xml:"uncompressed_cache_size"`
	MarkCacheSize         string `xml:"mark_cache_size"`
	Path                  string `xml:"path"`
	TmpPath               string `xml:"tmp_path"`
	UserFilesPath         string `xml:"user_files_path"`
	UsersConfig           string `xml:"users_config"`
	DefaultProfile        string `xml:"default_profile"`
	SystemProfile         string `xml:"system_profile"`
	DefaultDatabase       string `xml:"default_database"`
	MlockExecutable       string `xml:"mlock_executable"`
	RemoteServers         struct {
		Text string `xml:",chardata"`
		Incl string `xml:"incl,attr"`
	} `xml:"remote_servers"`
	IncludeFrom string `xml:"include_from"`
	Zookeeper   struct {
		Text     string `xml:",chardata"`
		Incl     string `xml:"incl,attr"`
		Optional string `xml:"optional,attr"`
	} `xml:"zookeeper"`
	Macros struct {
		Text     string `xml:",chardata"`
		Incl     string `xml:"incl,attr"`
		Optional string `xml:"optional,attr"`
	} `xml:"macros"`
	BuiltinDictionariesReloadInterval string `xml:"builtin_dictionaries_reload_interval"`
	MaxSessionTimeout                 string `xml:"max_session_timeout"`
	DefaultSessionTimeout             string `xml:"default_session_timeout"`
	Prometheus                        struct {
		Text                string `xml:",chardata"`
		Endpoint            string `xml:"endpoint"`
		Port                string `xml:"port"`
		Metrics             string `xml:"metrics"`
		Events              string `xml:"events"`
		AsynchronousMetrics string `xml:"asynchronous_metrics"`
	} `xml:"prometheus"`
	QueryLog struct {
		Text                      string `xml:",chardata"`
		Database                  string `xml:"database"`
		Table                     string `xml:"table"`
		PartitionBy               string `xml:"partition_by"`
		FlushIntervalMilliseconds string `xml:"flush_interval_milliseconds"`
	} `xml:"query_log"`
	TraceLog struct {
		Text                      string `xml:",chardata"`
		Database                  string `xml:"database"`
		Table                     string `xml:"table"`
		PartitionBy               string `xml:"partition_by"`
		FlushIntervalMilliseconds string `xml:"flush_interval_milliseconds"`
	} `xml:"trace_log"`
	QueryThreadLog struct {
		Text                      string `xml:",chardata"`
		Database                  string `xml:"database"`
		Table                     string `xml:"table"`
		PartitionBy               string `xml:"partition_by"`
		FlushIntervalMilliseconds string `xml:"flush_interval_milliseconds"`
	} `xml:"query_thread_log"`
	DictionariesConfig string `xml:"dictionaries_config"`
	Compression        struct {
		Text string `xml:",chardata"`
		Incl string `xml:"incl,attr"`
	} `xml:"compression"`
	DistributedDdl struct {
		Text string `xml:",chardata"`
		Path string `xml:"path"`
	} `xml:"distributed_ddl"`
	GraphiteRollupExample struct {
		Text    string `xml:",chardata"`
		Pattern struct {
			Text      string `xml:",chardata"`
			Regexp    string `xml:"regexp"`
			Function  string `xml:"function"`
			Retention []struct {
				Text      string `xml:",chardata"`
				Age       string `xml:"age"`
				Precision string `xml:"precision"`
			} `xml:"retention"`
		} `xml:"pattern"`
		Default struct {
			Text      string `xml:",chardata"`
			Function  string `xml:"function"`
			Retention []struct {
				Text      string `xml:",chardata"`
				Age       string `xml:"age"`
				Precision string `xml:"precision"`
			} `xml:"retention"`
		} `xml:"default"`
	} `xml:"graphite_rollup_example"`
	FormatSchemaPath     string               `xml:"format_schema_path"`
	StorageConfiguration StorageConfiguration `xml:"storage_configuration"`
}

type StorageConfiguration struct {
	//Name  xml.Name `xml:"storage_configuration"`
	Disks    *Disks    `xml:"disks"`
	Policies *Policies `xml:"policies"`
}

/*Disk*/
type Disks struct {
	Disks *[]Disk `xml:",any"`
}
type Disk struct {
	//diskname
	XMLName xml.Name
	Path    string `xml:"path"`
}

/*Policies*/
type Policies struct {
	Policies *[]Policie `xml:",any"`
}

type Policie struct {
	XMLName xml.Name
	Volumes *Volumes `xml:"volumns"`
}

type Volumes struct {
	Volumes *[]Volume `xml:",any"`
}
type Volume struct {
	XMLName                  xml.Name
	Disks                    []string `xml:"disk"`
	Max_data_part_size_bytes string   `xml:"max_data_part_size_bytes"`
}

func GetLocalTypeName(name string) xml.Name {
	return xml.Name{Local: name}
}

func (this Config) Marshal() {
	anyObjectMarshal(this)
}
