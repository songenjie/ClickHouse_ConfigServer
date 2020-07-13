package structure

import (
	"encoding/xml"
	"errors"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"os"
	"path/filepath"
)

var (
	app                  = kingpin.New(filepath.Base(os.Args[0]), "The jdolap-api-server")
	ClickHouseConfigFile = app.Flag("clickhouseconfigfile", "clickhouse config file of json fullpath").Default("./conf/clickhouse.json").String()
	DataDir              = app.Flag("datadir", "fromat all cluster ip config id datadir").Default("./data").String()
	ClickHouses          = []*ClickHouse{}
	ClickHouesConfigs    = map[string]*ClickhouseConfig{}
	FileNames            = []string{"config", "metrika", "users", "schema", "zookeeper"}
)

type ClickhouseConfig struct {
	Configs  map[string]*Config
	Metrikas map[string]*Metrika
	Users
	Zookeeper
	Schema
}

//Format File
func anyObjectMarshal(v interface{}) string {
	Xml, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Println(err)
		return ""
	}
	log.Println("xml: \n", string(Xml))
	return string(Xml)
}

const (
	CONFIG    = "config"
	METRIKA   = "metrika"
	USERS     = "users"
	SCHEMA    = "shema"
	ZOOKEEPER = "zookeeper"
)

type Yandex interface {
	FormatFile(ClickHouse) error
}

type ConfigServerYandex struct {
	Yandex
	*ClickHouse
}

func (this ClickHouse) GetFileObject(filename string) (ConfigServerYandex, error) {
	switch filename {
	case CONFIG:
		return ConfigServerYandex{
			Yandex:     Config{},
			ClickHouse: &this,
		}, nil
	case METRIKA:
		return ConfigServerYandex{
			Yandex:     Metrika{},
			ClickHouse: &this,
		}, nil
	case USERS:
		return ConfigServerYandex{
			Yandex:     Users{},
			ClickHouse: &this,
		}, nil
	case SCHEMA:
		return ConfigServerYandex{
			Yandex:     Schema{},
			ClickHouse: &this,
		}, nil
	case ZOOKEEPER:
		return ConfigServerYandex{
			Yandex:     Zookeeper{},
			ClickHouse: &this,
		}, nil
	default:
		return ConfigServerYandex{}, errors.New("not support init : " + filename + " file !")
	}
	return ConfigServerYandex{}, nil
}

func (configServerYandex ConfigServerYandex) FormatFile() error {
	return configServerYandex.Yandex.FormatFile(*configServerYandex.ClickHouse)
}

func (clickhouse ClickHouse) FormatFile() error {
	for _, filename := range FileNames {
		ConfigServerYandex, err := clickhouse.GetFileObject(filename)
		if err != nil {
			return err
		}
		if err = ConfigServerYandex.FormatFile(); err != nil {
			return err
		}
	}
	return nil
}

func ParserError(err error) {
	if err != nil {
		log.Println(err)
	}
}
