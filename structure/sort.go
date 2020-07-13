package structure

import "sort"

//sort
type Replicas []*Replica

func (s Replicas) Len() int {
	return len(s)
}

func (s Replicas) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type ByReplica struct {
	Replicas
}

func (s ByReplica) Less(i, j int) bool {
	return  s.Replicas[i].Replica > s.Replicas[j].Replica
}

type Groups []*Group

func (s Groups) Len() int {
	return len(s)
}

func (s Groups) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type ByGroup struct {
	Groups
}

func (s ByGroup) Less(i, j int) bool {
	return  s.Groups[i].GroupNumber > s.Groups[j].GroupNumber
}

func (clickhouse ClickHouse) Sort() {
	//ipx index
	sort.Strings(clickhouse.Zookeeper.Ips)
	//port for replica 1 2 3
	sort.Ints(clickhouse.Shard.Port)
	//sort by gourp number
	clickhouse.SortGroups()
	//sort bt replica number
	clickhouse.SortReplicas()
}

func (clickhouse ClickHouse) SortReplicas() {
	sort.Sort(ByReplica{clickhouse.Replicas})
}

func (clickhouse ClickHouse) SortGroups() {
	sort.Sort(ByGroup{clickhouse.Shard.Groups})
}

