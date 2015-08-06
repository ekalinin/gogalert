package response

import (
	"strconv"
)

func ValidateCondition(val float32, condition string, threshhold float32) bool {
	res := false
	switch condition {
	case "eq":
		res = val == threshhold
	case "gt":
		res = val > threshhold
	case "ge":
		res = val >= threshhold
	case "lt":
		res = val < threshhold
	case "le":
		res = val <= threshhold
	}
	//fmt.Printf("%v %v %v --> %v\n", val, condition, threshhold, res)
	return res
}

// Full metric description
type MetricFlat struct {
	GMetric
	GHost
	GCluster
	GGrid
}

// Options for Find method
type MetricFilter struct {
	MetricName  string
	HostName    string
	ClusterName string
	Condition   string
	Threshhold  int
}

func (self GMetaResponse) Find(filter *MetricFilter) []MetricFlat {
	res := []MetricFlat{}

	for _, g := range self.Grids {
		for c := 0; c < len(g.Clusters); c++ {
			currCluster := g.Clusters[c]
			fCluster := filter.ClusterName

			if fCluster != "" && currCluster.Name != fCluster {
				continue
			}
			for h := 0; h < len(currCluster.Hosts); h++ {
				currHost := currCluster.Hosts[h]
				fHost := filter.HostName

				if fHost != "" && currHost.Name != fHost {
					continue
				}
				for m := 0; m < len(currHost.Metrics); m++ {
					currMetric := currHost.Metrics[m]
					fMetric := filter.MetricName

					if fMetric != "" && currMetric.Name != fMetric {
						continue
					}

					if filter.Condition != "" {
						currVal, _ := strconv.ParseFloat(currMetric.Val, 32)
						threshhold := float32(filter.Threshhold)
						if !ValidateCondition(float32(currVal), filter.Condition, threshhold) {
							continue
						}
					}

					res = append(res, MetricFlat{
						currMetric, currHost, currCluster, g})
				}
			}
		}
	}

	return res
}
