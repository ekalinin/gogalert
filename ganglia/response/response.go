package response

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

// Full metric description
type MetricFlat struct {
	GMetric
	GHost
	GCluster
	GGrid
}

// Response from Ganglia's Meta deamon
type GMetaResponse struct {
	XMLName xml.Name `xml:"GANGLIA_XML"`
	Version string   `xml:"VERSION,attr"`
	Source  string   `xml:"SOURCE,attr"`
	Grids   []GGrid  `xml:"GRID"`
}

func (self GMetaResponse) String() string {
	return fmt.Sprintf("Resp: Version=%s Source=%s Grids=%s",
		self.Version, self.Source, self.Grids)
}

// Options for Find method
type MetricFilter struct {
	MetricName  string
	HostName    string
	ClusterName string
	Condition   string
	Threshhold  int
}

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

					if filter.Condition != "" && filter.Condition != "" {
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

// List of grid in the Ganglia's Meta deamon's response
type GGrid struct {
	XMLName   xml.Name   `xml:"GRID"`
	Name      string     `xml:"NAME,attr"`
	Authority string     `xml:"AUTHORITY,attr"`
	Localtime string     `xml:"LOCALTIME,attr"`
	Clusters  []GCluster `xml:"CLUSTER"`
}

func (self GGrid) PrintNested() string {
	return fmt.Sprintf("\n\t Grid: Name=%s Auth=%s Time=%s Clusters=%s",
		self.Name, self.Authority, self.Localtime, self.Clusters)
}

func (self GGrid) String() string {
	return fmt.Sprintf("Grid: Name=%s Auth=%s Time=%s",
		self.Name, self.Authority, self.Localtime)
}

// List of cluster in each grid
type GCluster struct {
	XMLName   xml.Name `xml:"CLUSTER"`
	Name      string   `xml:"NAME,attr"`
	Localtime string   `xml:"LOCALTIME,attr"`
	Owner     string   `xml:"OWNER,attr"`
	LatLong   string   `xml:"LATLONG,attr"`
	Url       string   `xml:"URL,attr"`
	Hosts     []GHost  `xml:"HOST"`
}

func (self GCluster) PrintNested() string {
	return fmt.Sprintf("\n\t\t Cluster: Name=%s Time=%s Url=%s Hosts=%s",
		self.Name, self.Localtime, self.Url, self.Hosts)
}

func (self GCluster) String() string {
	return fmt.Sprintf("Cluster: Name=%s Time=%s Url=%s",
		self.Name, self.Localtime, self.Url)
}

// List of hosts in each cluster
type GHost struct {
	XMLName     xml.Name  `xml:"HOST"`
	Name        string    `xml:"NAME,attr"`
	IP          string    `xml:"IP,attr"`
	Reported    string    `xml:"REPORTED,attr"`
	Tn          string    `xml:"TN,attr"`
	Tmax        string    `xml:"TMAX,attr"`
	Dmax        string    `xml:"DMAX,attr"`
	Location    string    `xml:"LOCATION,attr"`
	GMonStarted string    `xml:"GMOND_STARTED,attr"`
	Tags        string    `xml:"TAGS,attr"`
	Metrics     []GMetric `xml:"METRIC"`
}

func (self GHost) PrintNested() string {
	return fmt.Sprintf("\n\t\t\t Host: Name=%s IP=%s Time=%s Tn=%s Metrics=%s",
		self.Name, self.IP, self.Reported, self.Tn, self.Metrics)
}

func (self GHost) String() string {
	return fmt.Sprintf("Host: Name=%s IP=%s Reported=%s Location=%s Tn=%s",
		self.Name, self.IP, self.Reported, self.Location, self.Tn)
}

// List of metrics in each cluster
type GMetric struct {
	XMLName xml.Name `xml:"METRIC"`
	Name    string   `xml:"NAME,attr"`
	Val     string   `xml:"VAL,attr"`
	Type    string   `xml:"TYPE,attr"`
	Units   string   `xml:"UNITS,attr"`
	Tn      string   `xml:"TN,attr"`
	Tmax    string   `xml:"TMAX,attr"`
	Dmax    string   `xml:"DMAX,attr"`
	Slope   string   `xml:"SLOPE,attr"`
	Source  string   `xml:"SOURCE,attr"`
	Extra   GExtra   `xml:"EXTRA_DATA"`
}

func (self GMetric) PrintNested() string {
	return fmt.Sprintf("\n\t\t\t\t Metric: Name=%s Val=%s Type=%s Units=%s",
		self.Name, self.Val, self.Type, self.Units)
}

func (self GMetric) String() string {
	return fmt.Sprintf("Metric: Name=%s Val=%s Type=%s Units=%s",
		self.Name, self.Val, self.Type, self.Units)
}

// Extra values for each metric
type GExtra struct {
	XMLName  xml.Name   `xml:"EXTRA_DATA"`
	Elements []GElement `xml:"EXTRA_ELEMENT"`
}

// Element of each extra value
type GElement struct {
	XMLName xml.Name `xml:"EXTRA_ELEMENT"`
	Name    string   `xml:"NAME,attr"`
	Val     string   `xml:"VAL,attr"`
}
