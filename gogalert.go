package main

import (
	"./ganglia/api"
	"./ganglia/response"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	// filters
	node    = kingpin.Flag("node", "Node name to search").String()
	cluster = kingpin.Flag("cluster", "Cluster name to search").String()
	metric  = kingpin.Flag("metric", "Metric name to search").String()
	// conditions
	threshhold = kingpin.Flag("threshhold", "Threshhold value").Int()
	condition  = kingpin.Flag("condition", "Condition").Enum("eq", "gt", "ge", "lt", "le")
	// source of the xml
	localPath  = kingpin.Flag("file", "Read gmeta response from local file").ExistingFile()
	remoteIP   = kingpin.Flag("host", "Read gmeta response from host, default: 127.0.0.1").Default("127.0.0.1").String()
	remotePort = kingpin.Flag("port", "Read gmeta response from port, default: 8651").Default("8651").Int()
	// list objects
	listMetric   = kingpin.Flag("list-metrics", "List metrics").Bool()
	listNodes    = kingpin.Flag("list-nodes", "List nodes").Bool()
	listClusters = kingpin.Flag("list-clusters", "List clusters").Bool()
)

func main() {

	kingpin.Parse()

	gResp := response.GMetaResponse{}
	filter := response.MetricFilter{*metric, *node, *cluster, *condition, *threshhold}

	if *localPath != "" {
		gResp = api.ParseFile(*localPath)
	} else {
		gResp = api.ParseSocket(*remoteIP, *remotePort)
	}

	s := api.NewGSet()
	for _, metric := range gResp.Find(&filter) {
		if *listClusters {
			s.PrintIfNotInSet(metric.GCluster.Name)
		}
		if *listNodes {
			s.PrintIfNotInSet(metric.GHost.Name)
		}
		if *listMetric {
			s.PrintIfNotInSet(metric.GMetric.Name)
		}
	}

}
