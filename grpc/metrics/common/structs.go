package common

import (
	"time"

	"encoding/json"

	"github.com/shirou/gopsutil/load"
)

//
// MetricsLoadEntity type
//
type MetricsLoadEntity struct {
	Avg  *load.AvgStat  `json:"avg"`
	Misc *load.MiscStat `json:"misc"`
}

//
// String method
//
func (rcv *MetricsLoadEntity) String() string {
	data, _ := json.Marshal(rcv)
	return string(data)
}

//
// MetricsSection type
//
type MetricsSection struct {
	Name string        `json:"name"`
	Docs []interface{} `json:"docs"`
}

//
// NewMetricsSection constructor
//
func NewMetricsSection(n string) *MetricsSection {
	return &MetricsSection{Name: n, Docs: make([]interface{}, 0)}
}

//
// MetricsModel type
//
type MetricsModel struct {
	Sections []*MetricsSection `json:"sections"`
}

//
// NewMetricsModel constructor
//
func NewMetricsModel() *MetricsModel {
	return &MetricsModel{
		Sections: make([]*MetricsSection, 0),
	}
}

//
// String method
//
func (rcv *MetricsModel) String() string {
	//data, _ := json.Marshal(rcv)
	data, _ := json.MarshalIndent(rcv, "", " ")
	return string(data)
}

//
// MetricsModelDevice type
//
type MetricsModelDevice struct {
	DeviceID  string            `json:"device_id"`
	HostUUID  string            `json:"host_uuid"`
	CreatedAt time.Time         `json:"created_at"`
	Sections  []*MetricsSection `json:"sections"`
}

//
// NewMetricsModelDevice constructor
//
func NewMetricsModelDevice(d, h string, t time.Time) *MetricsModelDevice {
	return &MetricsModelDevice{
		DeviceID:  d,
		HostUUID:  h,
		CreatedAt: t,
		Sections:  make([]*MetricsSection, 0),
	}
}

//
// String method
//
func (rcv *MetricsModelDevice) String() string {
	data, _ := json.Marshal(rcv)
	return string(data)
}
