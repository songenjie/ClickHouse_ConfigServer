package structure

import "encoding/xml"

type Metrika struct {
	XMLName                 xml.Name                 `xml:"yandex"`
	ClickhouseRemoteServers *[]ClickhouseRemoteServer `xml:"clickhouse_remote_servers"`
	Macros                  struct {
		Text    string `xml:",chardata"`
		Shard   string `xml:"shard"`
		Replica string `xml:"replica"`
	} `xml:"macros"`
	ZookeeperServers struct {
		Node []struct {
			Index string `xml:"index,attr"`
			Host  string `xml:"host"`
			Port  string `xml:"port"`
		} `xml:"node"`
	} `xml:"zookeeper-servers"`
	Networks struct {
		Ip string `xml:"ip"`
	} `xml:"networks"`
	ClickhouseCompression struct {
		Case struct {
			MinPartSize      string `xml:"min_part_size"`
			MinPartSizeRatio string `xml:"min_part_size_ratio"`
			Method           string `xml:"method"`
		} `xml:"case"`
	} `xml:"clickhouse_compression"`
}

type ClickhouseRemoteServers struct {
	ClickhouseRemoteServers *[]ClickhouseRemoteServers
}

type ClickhouseRemoteServer struct {
	XMLName xml.Name
	Shard   []struct {
		InternalReplication string `xml:"internal_replication"`
		Weight              string `xml:"weight"`
		Replica             []struct {
			Host string `xml:"host"`
			Port string `xml:"port"`
		} `xml:"replica"`
	} `xml:"shard"`
}
