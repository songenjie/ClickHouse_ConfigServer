package structure

import "encoding/xml"

type Users struct {
	XMLName  xml.Name `xml:"yandex"`
	Text     string   `xml:",chardata"`
	Profiles struct {
		Text    string `xml:",chardata"`
		Default struct {
			Text                 string `xml:",chardata"`
			MaxMemoryUsage       string `xml:"max_memory_usage"`
			UseUncompressedCache string `xml:"use_uncompressed_cache"`
			LoadBalancing        string `xml:"load_balancing"`
		} `xml:"default"`
		Readonly struct {
			Text          string `xml:",chardata"`
			LoadBalancing string `xml:"load_balancing"`
			Readonly      string `xml:"readonly"`
		} `xml:"readonly"`
	} `xml:"profiles"`
	Users struct {
		Text    string `xml:",chardata"`
		Default struct {
			Text     string `xml:",chardata"`
			Password string `xml:"password"`
			Networks struct {
				Text    string `xml:",chardata"`
				Incl    string `xml:"incl,attr"`
				Replace string `xml:"replace,attr"`
				Ip      string `xml:"ip"`
			} `xml:"networks"`
			Profile string `xml:"profile"`
			Quota   string `xml:"quota"`
		} `xml:"default"`
	} `xml:"users"`
	Quotas struct {
		Text    string `xml:",chardata"`
		Default struct {
			Text     string `xml:",chardata"`
			Interval struct {
				Text          string `xml:",chardata"`
				Duration      string `xml:"duration"`
				Queries       string `xml:"queries"`
				Errors        string `xml:"errors"`
				ResultRows    string `xml:"result_rows"`
				ReadRows      string `xml:"read_rows"`
				ExecutionTime string `xml:"execution_time"`
			} `xml:"interval"`
		} `xml:"default"`
	} `xml:"quotas"`
}