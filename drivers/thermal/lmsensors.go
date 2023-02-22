package thermal

import (
	"encoding/json"
	"errors"
	"os/exec"
)

const (
	AGGREGATE_AVERAGE = iota
	AGGREGATE_MIN
	AGGREGATE_MAX
)

type LMSensorsDriver struct {
	Aggregation int
	NameFilter  map[string]bool
}

type lmSensorsSensor = map[string]float64
type lmSensorsDevice = map[string]interface{}
type lmSensorsJSON = map[string]lmSensorsDevice

func (d *LMSensorsDriver) Init() error {
	return nil
}

func (d *LMSensorsDriver) Close() error {
	return nil
}

func (d *LMSensorsDriver) GetTemperature() (float64, error) {
	rawData, err := exec.Command("sensors", "-j").Output()
	if err != nil {
		return 0, err
	}

	var data lmSensorsJSON
	err = json.Unmarshal(rawData, &data)
	if err != nil {
		return 0, err
	}

	var val float64 = -1
	var count int = 0
	for _, device := range data {
		for sensorName, sensorRaw := range device {
			if len(d.NameFilter) > 0 && !d.NameFilter[sensorName] {
				continue
			}

			sensorWant, ok := sensorRaw.(map[string]interface{})
			if !ok {
				continue
			}

			var sensorValue float64 = -1
			for sensorField, sensorFieldValue := range sensorWant {
				if sensorField[len(sensorField)-5:] != "_input" {
					continue
				}
				if sensorField[:4] != "temp" {
					continue
				}
				sensorValue, ok = sensorFieldValue.(float64)
				if !ok {
					sensorValue = -1
					continue
				}
				break
			}

			if sensorValue < 0 {
				continue
			}

			count++
			switch d.Aggregation {
			case AGGREGATE_AVERAGE:
				val += sensorValue
			case AGGREGATE_MIN:
				if sensorValue < val || val < 0 {
					val = sensorValue
				}
			case AGGREGATE_MAX:
				if sensorValue > val || val < 0 {
					val = sensorValue
				}
			}
		}
	}

	if count == 0 {
		return 0, errors.New("no sensors matched")
	}

	if d.Aggregation == AGGREGATE_AVERAGE {
		val /= float64(count)
	}

	return val, nil
}
