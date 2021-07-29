package main

import (
	"fmt"
	"time"
)

type sensorData interface {
	getMeasurementTime() time.Time
}

type stream interface {
	getSensorData() sensorData
	isEmpty() bool
	getBatch(startTime time.Time) []sensorData
}

//ensure that sensor streams meets interface
var _ stream = &positionStream{}
var _ stream = &temperatureStream{}
var _ stream = &powerStream{}

type positionData struct {
	deviceId        int
	measurementTime time.Time
	latitude        string
	longitude       string
}

func (p *positionData) getMeasurementTime() time.Time {
	return p.measurementTime
}

func (p *positionData) String() string {
	return fmt.Sprintf("Position sensor:\tmeasurementTime: %s, deviceId: %d, latitude: %s, longitude: %s \n", p.measurementTime.String(), p.deviceId, p.latitude, p.longitude)
}

type temperatureData struct {
	deviceId        int
	measurementTime time.Time
	temperature     string
}

func (t temperatureData) getMeasurementTime() time.Time {
	return t.measurementTime
}

func (p *temperatureData) String() string {
	return fmt.Sprintf("Temperature sensor:\tmeasurementTime: %s, deviceId: %d, temperature: %s\n", p.measurementTime.String(), p.deviceId, p.temperature)
}

type powerData struct {
	deviceId        int
	measurementTime time.Time
	power           string
}

func (p powerData) getMeasurementTime() time.Time {
	return p.measurementTime
}

func (p *powerData) String() string {
	return fmt.Sprintf("Power sensor:\t\tmeasurementTime: %s, deviceId: %d, power: %s\n", p.measurementTime.String(), p.deviceId, p.power)
}

type positionStream struct {
	positions []*positionData
}

func (p *positionStream) getSensorData() sensorData {
	sensor := p.positions[0]
	p.positions = p.positions[1:]
	return sensor
}

func (p *positionStream) isEmpty() bool {
	return len(p.positions) == 0
}

func (p *positionStream) getBatch(startTime time.Time) []sensorData {
	positionBatch := []sensorData{}

	for {
		if p.isEmpty() {
			fmt.Println("waiting from position sensor stream")
			time.Sleep(5 * time.Second)
			continue
		} else if p.positions[0].measurementTime.Before(startTime.Add(batchDuration)) {
			positionBatch = append(positionBatch, p.getSensorData())
		} else {
			return positionBatch
		}
	}
}

type temperatureStream struct {
	temperatures []*temperatureData
}

func (t *temperatureStream) getSensorData() sensorData {
	sensor := t.temperatures[0]
	t.temperatures = t.temperatures[1:]
	return sensor
}

func (t *temperatureStream) isEmpty() bool {
	return len(t.temperatures) == 0
}

func (t *temperatureStream) getBatch(startTime time.Time) []sensorData {
	positionBatch := []sensorData{}

	for {
		if t.isEmpty() {
			fmt.Println("waiting from temperature sensor stream")
			time.Sleep(5 * time.Second)
			continue
		} else if t.temperatures[0].measurementTime.Before(startTime.Add(batchDuration)) {
			positionBatch = append(positionBatch, t.getSensorData())
		} else {
			return positionBatch
		}
	}
}

type powerStream struct {
	powers []*powerData
}

func (p *powerStream) getSensorData() sensorData {
	sensor := p.powers[0]
	p.powers = p.powers[1:]
	return sensor
}

func (p *powerStream) isEmpty() bool {
	return len(p.powers) == 0
}

func (p *powerStream) getBatch(startTime time.Time) []sensorData {
	positionBatch := []sensorData{}

	for {
		if p.isEmpty() {
			fmt.Println("waiting from power sensor stream")
			time.Sleep(5 * time.Second)
			continue
		} else if p.powers[0].measurementTime.Before(startTime.Add(batchDuration)) {
			positionBatch = append(positionBatch, p.getSensorData())
		} else {
			return positionBatch
		}
	}
}

type outputStream struct {
	out []sensorData
}

func (out *outputStream) print() {
	for _, pos := range out.out {
		fmt.Printf("%+v", pos)
	}
}
