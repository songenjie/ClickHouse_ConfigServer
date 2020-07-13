package main

import (
	"ClickHouse_ConfigServer/fsnotify"
	"ClickHouse_ConfigServer/logger"
	"ClickHouse_ConfigServer/structure"
	"encoding/json"
	"io/ioutil"
	"log"
)

func init() {

	ParserClickHouse()
	//notify
	//watch, err := fsnotify.NewNotifyFile()
	_, err := fsnotify.NewNotifyFile()
	if err != nil {
		return
	}
	//watch.WatchDir(*structure.DataDir)
}

func ParserClickHouse() {

	//1 use define clickhouse.json in apollo
	//2 configserver get clickhouse.json by key:cluster clickhouse_json
	body, err := ioutil.ReadFile(*structure.ClickHouseConfigFile)
	if err != nil {
		return
	}
	//3 parse to struct
	if err = json.Unmarshal(body, structure.ClickHouses); err != nil {
		return
	}
	//4 init file
	for _, clickhouse := range structure.ClickHouses {
		if len(clickhouse.Cluster) == 0 {
			return
		}
		log.Println(clickhouse.Cluster)
		clickhouse.Sort()
		clickhouse.FormatFile()
	}
}

func main() {
	log.Println("hello")
	logger.TestLevelRotate()
	//logger.TestingLevelFileLog()
	//logger.TestRotatingFileLog()
}
