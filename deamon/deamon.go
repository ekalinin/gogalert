package deamon

import (
	"fmt"
	"github.com/ekalinin/gogalert/api"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type StatServer struct {
	// Period to refresh data from ganglia
	refreshRate int
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
	return &StatServer{60, source, threshholds, configPath, nil}
}

func (this *StatServer) Start() error {
	fmt.Println(" Server: started, config: ", *this)
	this.config = ParseConfig(this.configPath)
	this.refreshRate = this.config.Interval
	fmt.Println(" Server: parsed config: ", *this.config)

	go this.CollectStats()

	signalChan := make(chan os.Signal, 1)
	// http://adampresley.com/2015/02/16/waiting-for-goroutines-to-finish-running-before-exiting.html
	// https://gist.github.com/reiki4040/be3705f307d3cd136e85
	signal.Notify(signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	needStop := false
	for {
		select {
		case <-time.After(time.Duration(this.refreshRate) * time.Second):
			go this.CollectStats()
		case <-signalChan:
			fmt.Println("Stopping...")
			needStop = true
		}

		if needStop {
			break
		}
	}

	return nil
}

func (this *StatServer) CollectStats() {
	fmt.Println(" Server: collecting data ...")
	for _, metric := range api.Parse(this.source).Find(this.threshholds) {
		fmt.Println(" >> Alert: ", metric.GMetric)
	}
}
