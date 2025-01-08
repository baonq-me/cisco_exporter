package environment

import (
	"encoding/json"
	"errors"
	"github.com/lwlcom/cisco_exporter/rpc"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// Parse parses cli output and tries to find oll temperature and power related data
func (c *environmentCollector) Parse(ostype string, output string) ([]EnvironmentItem, error) {
	if ostype != rpc.IOSXE && ostype != rpc.NXOS && ostype != rpc.IOS {
		return nil, errors.New("'show environment' is not implemented for " + ostype)
	}

	var items []EnvironmentItem
	var outputJson RootEnvironment

	/*
		Compile regex to match the desired content between line breaks
		This regex will convert ...

		IronLink# show environment | json
		<json data>
		IronLink# show environment | json

		... to

		<json data>
	*/
	re := regexp.MustCompile(`(?m)^.*\n(.*?)\n.*$`)

	// Extract the content
	match := re.FindStringSubmatch(output)
	if len(match) > 1 {
		// log.Printf("Extracted content:", match[1])
		outputJsonString := match[1]
		// Convert the string to JSON
		err := json.Unmarshal([]byte(outputJsonString), &outputJson)
		if err != nil {
			log.Printf("error decoding json: %v", err)
			if e, ok := err.(*json.SyntaxError); ok {
				log.Printf("syntax error at byte offset %d", e.Offset)
			}
			// log.Printf("raw string: %q", output)
		}
	}

	// Extract temp
	for _, tempRow := range outputJson.TempInfo.RowTempInfo {
		tempValue, _ := strconv.Atoi(tempRow.CurTemp)
		items = append(items, EnvironmentItem{
			Name:   tempRow.Sensor,
			Status: tempRow.AlarmStatus,
			Value:  float64(tempValue),
			Type:   Temp,
		})
	}

	// Extract power
	for _, psRow := range outputJson.PowerSup.PSInfo.RowPSInfo {
		powerInValue, _ := strconv.Atoi(strings.Replace(psRow.ActualInput, " W", "", -1))
		powerOutValue, _ := strconv.Atoi(strings.Replace(psRow.ActualOut, " W", "", -1))
		powerCapacity, _ := strconv.Atoi(strings.Replace(psRow.TotalCapa, " W", "", -1))

		items = append(items, EnvironmentItem{
			Name:   psRow.PSNum,
			Detail: psRow.PSModel,
			Status: psRow.PSStatus,
			Value:  float64(powerInValue),
			Type:   PowerIn,
		})

		items = append(items, EnvironmentItem{
			Name:   psRow.PSNum,
			Detail: psRow.PSModel,
			Status: psRow.PSStatus,
			Value:  float64(powerOutValue),
			Type:   PowerOut,
		})

		items = append(items, EnvironmentItem{
			Name:   psRow.PSNum,
			Detail: psRow.PSModel,
			Status: psRow.PSStatus,
			Value:  float64(powerCapacity),
			Type:   PowerCapacity,
		})

		items = append(items, EnvironmentItem{
			Name:   psRow.PSNum,
			Detail: psRow.PSModel,
			Status: psRow.PSStatus,
			Value:  float64(getStatus(psRow.PSStatus)),
			Type:   PowerStatus,
		})
	}
	items = append(items, EnvironmentItem{
		Name:   "configured",
		Status: outputJson.PowerSup.PowerSummary.PSRedunMode,
		Value:  0,
		Type:   PowerMode,
	})
	items = append(items, EnvironmentItem{
		Name:   "operational",
		Status: outputJson.PowerSup.PowerSummary.PSOperMode,
		Value:  0,
		Type:   PowerMode,
	})

	// Extract fan
	for _, fanRow := range outputJson.FanDetails.FanInfo.RowFanInfo {
		items = append(items, EnvironmentItem{
			Name:   fanRow.FanName,
			Detail: fanRow.FanModel,
			Status: fanRow.FanStatus,
			Value:  float64(getStatus(fanRow.FanStatus)),
			Type:   FanStatus,
		})
	}
	items = append(items, EnvironmentItem{
		Name:  outputJson.FanDetails.FanZoneSpeed.RowFanZoneSpeed.Zone,
		Value: float64(convertFanSpeedToPercent(outputJson.FanDetails.FanZoneSpeed.RowFanZoneSpeed.ZoneSpeed)),
		Type:  FanSpeed,
	})

	return items, nil
}

func getStatus(status string) int {
	lowerStatus := strings.ToLower(status)
	if lowerStatus == "good" || lowerStatus == "ok" || lowerStatus == "normal" {
		return 1
	} else {
		return 0
	}
}

func convertFanSpeedToPercent(speed string) int {

	decimal, err := strconv.ParseInt(strings.TrimPrefix(speed, "0x"), 16, 64)
	if err != nil {
		return -1
	}

	// Ensure the decimal value is within the valid range
	if decimal < 0 || decimal > 255 {
		return -2
	}

	// Calculate the percentage (decimal / 255 * 100)
	return int(float64(decimal) / 255 * 100)

}
