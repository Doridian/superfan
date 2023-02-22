package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/FoxDenHome/superfan/drivers/control"
	"github.com/FoxDenHome/superfan/drivers/curve"
	"github.com/FoxDenHome/superfan/drivers/thermal"
)

func registerShutdownSignals(callback func()) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		callback()
	}()
}

func main() {
	ctrl := control.X10IPMIDriver{
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
	therm := thermal.LMSensorsDriver{
		Aggregation: thermal.AGGREGATE_MAX,
		NameFilter:  nameFilterMap,
	}
	err = therm.Init()
	if err != nil {
		panic(err)
	}
	defer therm.Close()

	curve := curve.FixedCurveDriver{
		Thresholds: []*curve.FixedThreshold{
			{
				Temperature: 10.0,
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

	shouldRun := true
	go func() {
		<-sigs
		shouldRun = false
	}()

	for shouldRun {
		time.Sleep(5 * time.Second)

		temp, err := therm.GetTemperature()
		if err != nil {
			log.Printf("Error getting temperature: %v", err)
			continue
		}
		speed, err := curve.GetFanSpeedFor(temp)
		if err != nil {
			log.Printf("Error getting target fan speed: %v", err)
			continue
		}
		currentSpeed, err := ctrl.GetFanSpeed()
		if err != nil {
			log.Printf("Error getting current fan speed: %v", err)
			continue
		}

		if currentSpeed == speed {
			log.Printf("[NOP] Temperature: %.0f, Fan speed: %.f%%", temp, currentSpeed*100)
			continue
		}

		err = ctrl.SetFanSpeed(speed)
		if err != nil {
			log.Printf("Error setting fan speed: %v", err)
			continue
		}
		log.Printf("[SET] Temperature: %.0f, Fan speed: %.0f%%", temp, speed*100)
	}
}
