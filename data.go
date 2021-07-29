package main

import "time"

func getSeed() (positionStream, temperatureStream, powerStream) {

	positionSensors := positionStream{}
	for i := 0; i < 37; i++ {
		period := 5 * time.Second
		position := &positionData{
			deviceId:        1,
			measurementTime: time.Date(2021, time.July, 1, 0, 0, 0, 0, time.UTC).Add(time.Duration(i) * period),
			latitude:        "20",
			longitude:       "20",
		}
		positionSensors.positions = append(positionSensors.positions, position)
	}

	tempertureSensors := temperatureStream{}
	for i := 0; i < 20; i++ {
		period := 20 * time.Second
		temperature := &temperatureData{
			deviceId:        1,
			measurementTime: time.Date(2021, time.July, 1, 0, 0, 0, 0, time.UTC).Add(time.Duration(i) * period),
			temperature:     "20 C",
		}
		tempertureSensors.temperatures = append(tempertureSensors.temperatures, temperature)
	}

	powerSensors := powerStream{}
	for i := 0; i < 20; i++ {
		period := 37 * time.Second
		power := &powerData{
			deviceId:        1,
			measurementTime: time.Date(2021, time.July, 1, 0, 0, 0, 0, time.UTC).Add(time.Duration(i) * period),
			power:           "100 W",
		}
		powerSensors.powers = append(powerSensors.powers, power)
	}

	return positionSensors, tempertureSensors, powerSensors
}
