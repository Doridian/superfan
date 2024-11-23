package main

import (
	"log"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/FoxDenHome/superfan/drivers/control"
	"github.com/FoxDenHome/superfan/drivers/curve"
	"github.com/FoxDenHome/superfan/drivers/thermal"
)

func runLoop(therm thermal.Driver, curve curve.Curve, ctrl control.Driver) bool {
	temp, err := therm.GetTemperature()
	if err != nil {
		log.Printf("Error getting temperature: %v", err)
		return false
	}
	speed, err := curve.GetFanSpeedFor(temp)
	if err != nil {
		log.Printf("Error getting target fan speed: %v", err)
		return false
	}
	currentSpeed, err := ctrl.GetFanSpeed()
	if err != nil {
		log.Printf("Error getting current fan speed: %v", err)
		return false
	}
	log.Printf("[GET] Temperature: %.0f", temp)

	speedDiff := math.Abs(speed - currentSpeed)
	if speedDiff < 0.01 {
		return false
	}

	err = ctrl.SetFanSpeed(speed)
	if err != nil {
		log.Printf("Error setting fan speed: %v", err)
		return false
	}
	log.Printf("[SET] Fan speed: %.0f%%", speed*100)
	return true
}

const shortSleep = 5 * time.Second
const longSleep = 15 * time.Second

func main() {
	ctrl := &control.X10IPMIDriver{
		IPMIDriver: control.IPMIDriver{
			DeviceIndex: 0,
		},
	}
	err := ctrl.Init()
	if err != nil {
		panic(err)
	}
	defer ctrl.Close()

	nameFilterMap := make(map[string]bool)
	nameFilterMap["Package id 0"] = true
	nameFilterMap["Package id 1"] = true
	therm := &thermal.LMSensorsDriver{
		Aggregation: thermal.AGGREGATE_MAX,
		NameFilter:  nameFilterMap,
	}
	err = therm.Init()
	if err != nil {
		panic(err)
	}
	defer therm.Close()

	curve := &curve.LinearInterpolated{
		Points: []*curve.Point{
			{
				Temperature: 30.0,
				Speed:       0.03,
			},
			{
				Temperature: 55.0,
				Speed:       0.08,
			},
			{
				Temperature: 65.0,
				Speed:       0.19,
			},
			{
				Temperature: 70.0,
				Speed:       0.63,
			},
			{
				Temperature: 75.0,
				Speed:       1.00,
			},
		},
	}
	err = curve.Init()
	if err != nil {
		panic(err)
	}
	defer curve.Close()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	time.Sleep(100 * time.Millisecond)
runLoopFor:
	for {
		sleepTime := shortSleep
		didChange := runLoop(therm, curve, ctrl)
		if didChange {
			sleepTime = longSleep
		}

		select {
		case <-sigs:
			break runLoopFor
		case <-time.After(sleepTime):
		}
	}
	time.Sleep(100 * time.Millisecond)
}
