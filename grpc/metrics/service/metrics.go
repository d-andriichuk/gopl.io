package service

import (
	"context"
	"log"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"google.golang.org/grpc"
	"gopl.io/grpc/metrics/common"
	pb "gopl.io/grpc/metrics/metrics"
)

//
// Metrics type
//
type Metrics struct {
	// env contracts.IEnv

	// cloud contracts.ICloud
}

//
// NewMetrics constructor
//
func NewMetrics() *Metrics {
	// func NewMetrics(e contracts.IEnv, c contracts.ICloud) contracts.IMetrics {
	//return &Metrics{env: e, cloud: c}
	return &Metrics{}
}

//
// Name method
//
func (rcv *Metrics) Name() string {
	return "metrics"
}

//
// Statistics method
//
func (rcv *Metrics) Statistics() *common.MetricsModel {
	metricsModel := common.NewMetricsModel()

	//
	// IO section
	//
	{
		metricsSection := common.NewMetricsSection("io")

		metrics, _ := disk.IOCounters()

		metricsSection.Docs = append(metricsSection.Docs, metrics)
		metricsModel.Sections = append(metricsModel.Sections, metricsSection)
	}

	//
	// CPU section
	//
	{
		metricsSection := common.NewMetricsSection("cpu")

		metrics, _ := cpu.Times(true)

		metricsSection.Docs = append(metricsSection.Docs, metrics)
		metricsModel.Sections = append(metricsModel.Sections, metricsSection)
	}

	//
	// Mem section
	//
	{
		metricsSection := common.NewMetricsSection("mem")

		metrics, _ := mem.VirtualMemory()

		metricsSection.Docs = append(metricsSection.Docs, metrics)
		metricsModel.Sections = append(metricsModel.Sections, metricsSection)
	}

	//
	// Net section
	//
	{
		metricsSection := common.NewMetricsSection("net")

		metrics, _ := net.IOCounters(true)

		metricsSection.Docs = append(metricsSection.Docs, metrics)
		metricsModel.Sections = append(metricsModel.Sections, metricsSection)
	}

	//
	// Disk section
	//
	{
		metricsSection := common.NewMetricsSection("disk")

		metrics, _ := disk.Usage("/")

		metricsSection.Docs = append(metricsSection.Docs, metrics)
		metricsModel.Sections = append(metricsModel.Sections, metricsSection)
	}

	//
	// Swap section
	//
	{
		metricsSection := common.NewMetricsSection("swap")

		metrics, _ := mem.SwapMemory()

		metricsSection.Docs = append(metricsSection.Docs, metrics)
		metricsModel.Sections = append(metricsModel.Sections, metricsSection)
	}

	//
	// Load section
	//
	{
		metricsSection := common.NewMetricsSection("load")

		metrics := &common.MetricsLoadEntity{}

		metrics.Avg, _ = load.Avg()
		metrics.Misc, _ = load.Misc()

		metricsSection.Docs = append(metricsSection.Docs, metrics)
		metricsModel.Sections = append(metricsModel.Sections, metricsSection)
	}

	return metricsModel
}

//
// Exchange method
//
func (rcv *Metrics) Exchange() error {
	// metricsModel := rcv.Statistics()

	// todo: send metrics statisctics to cloud
	// rcv.cloud.Connection()...
	//return nil

	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewMetricsClient(conn)

	_, err = client.Exchange(context.Background(), &pb.MetricsModel{Metrics: rcv.Statistics().String()})
	if err != nil {
		return err
	}

	return nil
}

//
// Initialize method
//
func (rcv *Metrics) Initialize() {
	log.Printf("banjo: register service metrics for statistics gathering...")
}

//
// Finalize method
//
func (rcv *Metrics) Finalize() {
	log.Printf("banjo: unregister service metrics for statistics gathering...")
}
