package api

import (
	"errors"
	"fmt"
	"github.com/ekalinin/gogalert/gmeta"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
)

type DataSource struct {
	Path string
	Host string
	Port int
}

func (this *DataSource) Read() ([]byte, error) {

	if this.Path != "" {
		xmlFile, err := os.Open(this.Path)
		CheckError(err)
		defer xmlFile.Close()
		return ioutil.ReadAll(xmlFile)
	}

	if this.Host != "" {
		connectString := []string{this.Host, strconv.Itoa(this.Port)}
		conn, err := net.Dial("tcp", strings.Join(connectString, ":"))
		CheckError(err)
		defer conn.Close()
		return ioutil.ReadAll(conn)
	}

	return []byte{}, errors.New("Not a file nor a socket!")
}

type GMetaWrapper struct {
	gmeta.GMetaResponse
}

// Full metric description
type MetricFlat struct {
	gmeta.GMetric
	gmeta.GHost
	gmeta.GCluster
	gmeta.GGrid
}

// Options for Find method
type MetricFilter struct {
	MetricName  string
	HostName    string
	ClusterName string
	Condition   string
	Threshhold  string
}

func NewMetricFilter(filter []string) *MetricFilter {
	return &MetricFilter{
		/* MetricName */ filter[2],
		/* HostName	*/ strings.Replace(filter[1], "*", "", -1),
		/* ClusterName */ strings.Replace(filter[0], "*", "", -1),
		/* Condition */ filter[3],
		/* Threshhold */ filter[4]}
}

// Parses XML fetched from source
func Parse(source *DataSource) *GMetaWrapper {
	xmlData, err := source.Read()
	CheckError(err)
	resp := gmeta.Parse(xmlData)

	return (&GMetaWrapper{resp})
}

func (self *GMetaWrapper) Find(filter *MetricFilter) []MetricFlat {
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
						currVal, err := strconv.ParseFloat(currMetric.Val, 32)
						if err != nil {
							panic(fmt.Sprintf("cannot parse float %s: %v", currMetric.Val, err))
						}
						threshhold, err := strconv.ParseFloat(filter.Threshhold, 32)
						if err != nil {
							panic(fmt.Sprintf("cannot parse float %s: %v", filter.Threshhold, err))
						}
						if !ValidateCondition(float32(currVal), filter.Condition, float32(threshhold)) {
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

// Validate condition
func ValidateCondition(val float32, condition string, threshhold float32) bool {
	res := false
	switch condition {
	case "eq":
		res = val == threshhold
	case "==":
		res = val == threshhold
	case "gt":
		res = val > threshhold
	case ">":
		res = val > threshhold
	case "ge":
		res = val >= threshhold
	case ">=":
		res = val >= threshhold
	case "lt":
		res = val < threshhold
	case "<":
		res = val < threshhold
	case "le":
		res = val <= threshhold
	case "<=":
		res = val <= threshhold
	}
	//fmt.Printf("%v %v %v --> %v\n", val, condition, threshhold, res)
	return res
}

// Check error, panic if it is
func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
