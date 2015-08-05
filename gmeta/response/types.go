package response

import (
	"encoding/xml"
	"fmt"
)

type GMetaResponse struct {
	XMLName xml.Name `xml:"GANGLIA_XML"`
	Version string   `xml:"VERSION,attr"`
	Source  string   `xml:"SOURCE,attr"`
	Grids   []GGrid  `xml:"GRID"`
}

func (self GMetaResponse) String() string {
	return fmt.Sprintf("Resp: Version=%s Source=%s",
		self.Version, self.Source)
}

// List of grid in the Ganglia's Meta deamon's response
type GGrid struct {
	XMLName   xml.Name   `xml:"GRID"`
	Name      string     `xml:"NAME,attr"`
	Authority string     `xml:"AUTHORITY,attr"`
	Localtime string     `xml:"LOCALTIME,attr"`
	Clusters  []GCluster `xml:"CLUSTER"`
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

func (self GCluster) String() string {
	return fmt.Sprintf("Cluster: Name=%s Time=%s Owner=%s LatLong=%s Url=%s",
		self.Name, self.Localtime, self.Owner, self.LatLong, self.Url)
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

func (self GHost) String() string {
	return fmt.Sprintf("Host: Name=%s IP=%s Reported=%s Tn=%s Tmax=%s Dmax=%s Location=%s GMonStarted=%s Tags=%s",
		self.Name, self.IP, self.Reported, self.Tn, self.Tmax, self.Dmax, self.Location, self.GMonStarted, self.Tags)
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

func (self GMetric) String() string {
	return fmt.Sprintf("Metric: Name=%s Val=%s Type=%s Units=%s Tn=%s Tmax=%s Dmax=%s Slope=%s Source=%s",
		self.Name, self.Val, self.Type, self.Units, self.Tn, self.Tmax, self.Dmax, self.Slope, self.Source)
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
