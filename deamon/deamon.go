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
	// config file path
	configPath string
	// config details
	config *Config
}

func NewStatServer(source *api.DataSource, threshholds *api.MetricFilter, configPath string) *StatServer {
	return &StatServer{2 * time.Minute, source, threshholds, configPath, nil}
}

func (this *StatServer) Start() error {
	fmt.Println(" Server: started, config: ", *this)
	this.config = ParseConfig(this.configPath)
	fmt.Println(" Server: parsed config: ", *this.config)
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
