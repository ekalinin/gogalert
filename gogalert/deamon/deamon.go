package deamon

import (
	"../ganglia/api"
	"../ganglia/response"
	"fmt"
	"time"
)

type StatServer struct {
	// Period to refresh data from ganglia
	refreshRate time.Duration
	// Source of the XML data
	source *api.DataSource
	// threshholds
	threshholds *response.MetricFilter
}

func NewStatServer(source *api.DataSource, threshholds *response.MetricFilter) *StatServer {
	return &StatServer{2 * time.Minute, source, threshholds}
}

func (this *StatServer) Start() error {
	fmt.Println(" Server: started, config: ", *this)
	for {
		go this.CollectStats()
		time.Sleep(this.refreshRate)
	}

	return nil
}

func (this *StatServer) CollectStats() {
	fmt.Println(" Server: collecting data ...")
	for _, metric := range api.Parse(this.source).Find(this.threshholds) {
		fmt.Println(" >> Alert: ", metric.GMetric)
	}
}
