package structure

import (
	"log"
	"strconv"
)

/*
/cluster/shard/replica/files
example
./jdolap_ck_04/01/01/ip
./jdolap_ck_04/01/01/file_size
./jdolap_ck_04/01/01/config.xml
./jdolap_ck_04/01/01/users.xml
./jdolap_ck_04/01/01/metrika.xml
./jdolap_ck_04/01/01/zookeeper.xml
./jdolap_ck_04/01/01/schema.xml
*/

func (clickhouse ClickHouse) FormatConfig(config Config) error {
	for _, group := range clickhouse.Shard.Groups {
		log.Println(group.GroupNumber)
		log.Println(group.Ips.Num1)
		log.Println(group.Ips.Num2)
		log.Println(group.Ips.Num3)
		ClickHouesConfigs[clickhouse.Cluster].Configs[strconv.Itoa(group.GroupNumber*3+1)] = nil;
	}
	return nil
}
func (clickhouse ClickHouse) FormatMetrika(metrika Metrika) error {
	//structure.ClickHouesConfigs[clickhouse.Cluster].Metrika.ClickhouseRemoteServers =
	//structure.ClickHouesConfigs[clickhouse.Cluster].Metrikas
	for _, group := range clickhouse.Shard.Groups {
		log.Println(group.GroupNumber)
		log.Println(group.Ips.Num1)
		log.Println(group.Ips.Num2)
		log.Println(group.Ips.Num3)
		ClickHouesConfigs[clickhouse.Cluster].Metrikas[strconv.Itoa(group.GroupNumber*3+1)] = nil;
	}
	return nil

}
func (clickhouse ClickHouse) FormatSchema(schema Schema) error {
	return nil

}
func (clickhouse ClickHouse) FormatUsers(users Users) error {
	ClickHouesConfigs[clickhouse.Cluster].Users.Profiles.Readonly.LoadBalancing = clickhouse.User.LoadBalancing
	ClickHouesConfigs[clickhouse.Cluster].Users.Profiles.Default.LoadBalancing = clickhouse.User.MaxMemoryUsage
	ClickHouesConfigs[clickhouse.Cluster].Users.Profiles.Default.MaxMemoryUsage = clickhouse.User.MaxMemoryUsage
	//structure.ClickHouesConfigs[clickhouse.Cluster].Users.Profiles.Default. = clickhouse.User.MaxMemoryUsageForAllQuery
	//structure.ClickHouesConfigs[clickhouse.Cluster].Users.Profiles.Readonly.LoadBalancing = clickhouse.User.SkipUnabvailiableShard
	return nil
}
func (clickhouse ClickHouse) FormatZookeeper(zookeeper Zookeeper) error {
	for index, zookeeperip := range clickhouse.Zookeeper.Ips {
		ClickHouesConfigs[clickhouse.Cluster].Zookeeper.Zookeeper.Node = append(
			ClickHouesConfigs[clickhouse.Cluster].Zookeeper.Zookeeper.Node,
			&Node{
				Index: strconv.Itoa(index),
				Host:  zookeeperip,
				Port:  clickhouse.Zookeeper.Port,
			})
	}
	return nil
}

func (this Config) FormatFile(clickhouse ClickHouse) error {
	return clickhouse.FormatConfig(this)
}

func (this Metrika) FormatFile(clickhouse ClickHouse) error {
	return clickhouse.FormatMetrika(this)
}

func (this Users) FormatFile(clickhouse ClickHouse) error {
	return clickhouse.FormatUsers(this)
}

func (this Schema) FormatFile(clickhouse ClickHouse) error {
	return clickhouse.FormatSchema(this)
}

func (this Zookeeper) FormatFile(clickhouse ClickHouse) error {
	return clickhouse.FormatZookeeper(this)
}
