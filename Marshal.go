package main

import (
	"encoding/xml"
	"log"
	logger_ "ClickHouse_ConfigServer/log"
)



var storageconfiguration = `
<yandex>
    <storage_configuration>
        <disks>
            <disk0>
                <path>/data0/jdolap/clickhouse/data/</path>
            </disk0>
            <disk1>
                <path>/data1/jdolap/clickhouse/data/</path>
            </disk1>
            <disk6>
                <path>/data6/jdolap/clickhouse/data1/</path>
            </disk6>
            <disk7>
                <path>/data7/jdolap/clickhouse/data1/</path>
            </disk7>
        </disks>

        <policies>
            <jdob_ha>
                <volumes>
                    <hot>
                        <disk>disk0</disk>
                        <disk>disk1</disk>
						<max_data_part_size_bytes>3298534883328</max_data_part_size_bytes>
                    </hot>
                    <cold>
                        <disk>disk6</disk>
                        <disk>disk7</disk>
                    </cold>
                </volumes>
                <move_factor>0.25</move_factor>
            </jdob_ha>
        </policies>
    </storage_configuration>
</yandex>`

type Yandex struct {
	XMLName              xml.Name             `xml:"yandex"`
	StorageConfiguration StorageConfiguration `xml:"storage_configuration"`
}

type StorageConfiguration struct {
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
	Max_data_part_size_bytes string   `xml:"max_data_part_size_bytes,omitempty"`
}

func GetLocalTypeName(name string) xml.Name {
	return xml.Name{Local: name}
}

func (this Yandex) Marshal() {
	anyObjectMarshal(this)
}

func anyObjectMarshal(v interface{}) {
	Xml, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("xml: \n", string(Xml))
}

func main() {
	/*Yandex := Yandex{
		GetLocalTypeName("yandex"),
		StorageConfiguration{
			Disks: &Disks{
				Disks: &[]Disk{
					{
						XMLName: GetLocalTypeName("diks0"),
						Path:    "/data0/jdolap/clickhouse/data",
					}, {
						XMLName: GetLocalTypeName("disk1"),
						Path:    "/data1/jdolap/clickhouse/data",
					},
					{
						XMLName: GetLocalTypeName("diks6"),
						Path:    "/data6/jdolap/clickhouse/data0",
					}, {
						XMLName: GetLocalTypeName("disk7"),
						Path:    "/data7/jdolap/clickhouse/data0",
					},
				},
			},
			Policies: &Policies{
				Policies: &[]Policie{
					{
						XMLName: GetLocalTypeName("jdob"),
						Volumes: &Volumes{
							Volumes: &[]Volume{
								{
									XMLName: GetLocalTypeName("hot"),
									Disks: []string{
										"disk0",
										"disk1",
									},
									Max_data_part_size_bytes: "0.25",
								},
								{
									XMLName: GetLocalTypeName("cold"),
									Disks: []string{
										"disk6",
										"disk7",
									},
								},
							},
						},
					},
				},
			},
		},
	}*/

	Yandex := Yandex{
		GetLocalTypeName("yandex"),
		StorageConfiguration{
			Disks: &Disks{
				Disks: &[]Disk{
					{
						XMLName: GetLocalTypeName("diks0"),
						Path:    "/data0/jdolap/clickhouse/data",
					}, {
						XMLName: GetLocalTypeName("disk1"),
						Path:    "/data1/jdolap/clickhouse/data",
					},
					{
						XMLName: GetLocalTypeName("diks6"),
						Path:    "/data6/jdolap/clickhouse/data0",
					}, {
						XMLName: GetLocalTypeName("disk7"),
						Path:    "/data7/jdolap/clickhouse/data0",
					},
				},
			},
			Policies: &Policies{
				Policies: &[]Policie{
					{
						XMLName: GetLocalTypeName("jdob"),
						Volumes: &Volumes{
							Volumes: &[]Volume{
								{
									XMLName: GetLocalTypeName("hot"),
									Disks: []string{
										"disk0",
										"disk1",
									},
									Max_data_part_size_bytes: "0.25",
								},
								{
									XMLName: GetLocalTypeName("cold"),
									Disks: []string{
										"disk6",
										"disk7",
									},
								},
							},
						},
					},
				},
			},
		},
	}
	Yandex.Marshal()


}
