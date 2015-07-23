package main

import (
	"./ganglia/api"
	"./ganglia/response"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	node        = kingpin.Flag("node", "Node name to search").String()
	cluster     = kingpin.Flag("cluster", "Cluster name to search").String()
	metric      = kingpin.Flag("metric", "Metric name to search").String()
	threshhold  = kingpin.Flag("threshhold", "Threshhold value").Int()
	condition   = kingpin.Flag("condition", "Condition").Enum("eq", "gt", "ge", "lt", "le")
	listMetric  = kingpin.Flag("list-metrics", "List metrics").Bool()
	listNodes   = kingpin.Flag("list-nodes", "List nodes").Bool()
	lisClusters = kingpin.Flag("list-clusters", "List clusters").Bool()
	localPath   = kingpin.Flag("file", "Read gmeta response from local file").ExistingFile()
	remoteIP    = kingpin.Flag("host", "Read gmeta response from host, default: 127.0.0.1").Default("127.0.0.1").String()
	remotePort  = kingpin.Flag("port", "Read gmeta response from port, default: 8651").Default("8651").Int()
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

	for _, metric := range gResp.Find(&filter) {
		fmt.Println(metric.GCluster)
		fmt.Println(metric.GHost)
		fmt.Println(metric.GMetric)
		fmt.Println()
	}

}
