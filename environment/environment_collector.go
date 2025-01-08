package environment

import (
	"github.com/lwlcom/cisco_exporter/rpc"
	"log"
	"strings"

	"github.com/lwlcom/cisco_exporter/collector"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "cisco_environment_"
const command string = "show environment | json"

var (
	temperaturesDesc        *prometheus.Desc
	powerSupplyInDesc       *prometheus.Desc
	powerSupplyOutDesc      *prometheus.Desc
	powerSupplyStatusDesc   *prometheus.Desc
	powerSupplyCapacityDesc *prometheus.Desc
	powerSupplyModeDesc     *prometheus.Desc

	fanStatusDesc *prometheus.Desc
	fanSpeedDesc  *prometheus.Desc
)

func init() {
	temperaturesDesc = prometheus.NewDesc(prefix+"sensor_temp", "Sensor temperatures", []string{"target", "item"}, nil)
	powerSupplyInDesc = prometheus.NewDesc(prefix+"power_in", "Input power in Watts", []string{"target", "item", "detail"}, nil)
	powerSupplyOutDesc = prometheus.NewDesc(prefix+"power_out", "Output power in Watts", []string{"target", "item", "detail"}, nil)
	powerSupplyStatusDesc = prometheus.NewDesc(prefix+"power_up", "Status of power supplies (1 OK, 0 Something is wrong)", []string{"target", "item", "detail"}, nil)
	powerSupplyCapacityDesc = prometheus.NewDesc(prefix+"power_capacity", "PSU capacity in Watts", []string{"target", "item", "detail"}, nil)
	powerSupplyModeDesc = prometheus.NewDesc(prefix+"power_mode", "Power Supply redundancy mode, value is always 0", []string{"target", "item", "status"}, nil)

	fanStatusDesc = prometheus.NewDesc(prefix+"fan_up", "Status of fans (1 OK, 0 Something is wrong)", []string{"target", "item", "detail"}, nil)
	fanSpeedDesc = prometheus.NewDesc(prefix+"fan_speed_percent", "Speed of fans, from 0 percent to 100", []string{"target", "item"}, nil)

}

type environmentCollector struct {
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &environmentCollector{}
}

// Name returns the name of the collector
func (*environmentCollector) Name() string {
	return "Environment"
}

// Describe describes the metrics
func (*environmentCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- temperaturesDesc
	ch <- powerSupplyInDesc
	ch <- powerSupplyOutDesc
	ch <- powerSupplyStatusDesc
	ch <- powerSupplyCapacityDesc
	ch <- powerSupplyModeDesc
}

// Collect collects metrics from Cisco
func (c *environmentCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	out, err := client.RunCommand(command)
	if err != nil {
		return err
	}
	outString := strings.ReplaceAll(out, command, "")
	items, err := c.Parse(client.OSType, outString)
	if err != nil {
		if client.Debug {
			log.Printf("Parse environment for %s: %s\n", labelValues[0], err.Error())
		}
		return nil
	}

	for _, item := range items {
		switch item.Type {
		case Temp:
			ch <- prometheus.MustNewConstMetric(temperaturesDesc, prometheus.GaugeValue, item.Value, append(labelValues, item.Name)...)
		case FanStatus:
			ch <- prometheus.MustNewConstMetric(fanStatusDesc, prometheus.GaugeValue, item.Value, append(labelValues, item.Name, item.Detail)...)

		case FanSpeed:
			ch <- prometheus.MustNewConstMetric(fanSpeedDesc, prometheus.GaugeValue, item.Value, append(labelValues, item.Name)...)

		case PowerIn:
			ch <- prometheus.MustNewConstMetric(powerSupplyInDesc, prometheus.GaugeValue, item.Value, append(labelValues, item.Name, item.Detail)...)

		case PowerOut:
			ch <- prometheus.MustNewConstMetric(powerSupplyOutDesc, prometheus.GaugeValue, item.Value, append(labelValues, item.Name, item.Detail)...)

		case PowerStatus:
			ch <- prometheus.MustNewConstMetric(powerSupplyStatusDesc, prometheus.GaugeValue, item.Value, append(labelValues, item.Name, item.Detail)...)

		case PowerCapacity:
			ch <- prometheus.MustNewConstMetric(powerSupplyCapacityDesc, prometheus.GaugeValue, item.Value, append(labelValues, item.Name, item.Detail)...)

		case PowerMode:
			ch <- prometheus.MustNewConstMetric(powerSupplyModeDesc, prometheus.GaugeValue, item.Value, append(labelValues, item.Name, item.Status)...)
		}
	}

	return nil
}
