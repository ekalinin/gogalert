package deamon

import (
	"fmt"
	"github.com/ekalinin/gogalert/api"
	"time"
)

type StatServer struct {
	// Period to refresh data from ganglia
	refreshRate time.Duration
	// Source of the XML data
	source *api.DataSource
	// threshholds
	threshholds *api.MetricFilter
}

func NewStatServer(source *api.DataSource, threshholds *api.MetricFilter) *StatServer {
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
