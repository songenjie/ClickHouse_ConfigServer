package structure

import "encoding/xml"

type Schema struct {
	XMLName       xml.Name `xml:"yandex"`
	Text          string   `xml:",chardata"`
	RemoteServers struct {
		Text string `xml:",chardata"`
		Ck06 struct {
			Text  string `xml:",chardata"`
			Shard []struct {
				Text                string `xml:",chardata"`
				InternalReplication string `xml:"internal_replication"`
				Replica             struct {
					Text string `xml:",chardata"`
					Host string `xml:"host"`
					Port string `xml:"port"`
				} `xml:"replica"`
				Weight string `xml:"weight"`
			} `xml:"shard"`
		} `xml:"ck_06"`
	} `xml:"remote_servers"`
	MaxWorkers   string `xml:"max_workers"`
	SettingsPull struct {
		Text     string `xml:",chardata"`
		Readonly string `xml:"readonly"`
	} `xml:"settings_pull"`
	SettingsPush struct {
		Text     string `xml:",chardata"`
		Readonly string `xml:"readonly"`
	} `xml:"settings_push"`
	Settings struct {
		Text                  string `xml:",chardata"`
		ConnectTimeout        string `xml:"connect_timeout"`
		InsertDistributedSync string `xml:"insert_distributed_sync"`
	} `xml:"settings"`
	Tables struct {
		Text      string `xml:",chardata"`
		Tabletmp2 struct {
			Text              string `xml:",chardata"`
			ClusterPull       string `xml:"cluster_pull"`
			DatabasePull      string `xml:"database_pull"`
			TablePull         string `xml:"table_pull"`
			ClusterPush       string `xml:"cluster_push"`
			DatabasePush      string `xml:"database_push"`
			TablePush         string `xml:"table_push"`
			Engine            string `xml:"engine"`
			ShardingKey       string `xml:"sharding_key"`
			EnabledPartitions struct {
				Text      string `xml:",chardata"`
				Partition string `xml:"partition"`
			} `xml:"enabled_partitions"`
		} `xml:"tabletmp2"`
	} `xml:"tables"`
}