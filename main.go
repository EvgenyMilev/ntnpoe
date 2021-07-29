package main

import (
	"fmt"
	"time"
)

const batchDuration = 1 * time.Minute

func main() {

	positionStream, temperatureStream, powerStream := getSeed()

	startTime := time.Date(2021, time.July, 1, 0, 0, 0, 0, time.UTC)
	batchCount := 0

	for {
		//get 1-minute data batches from streams
		positionBatch := positionStream.getBatch(startTime)
		temperatureBatch := temperatureStream.getBatch(startTime)
		powerBatch := powerStream.getBatch(startTime)
		batches := [][]sensorData{positionBatch, temperatureBatch, powerBatch}

		fmt.Printf("<<<<<<<<<<<<<<<< 1-MINUTE BATCH %d >>>>>>>>>>>>>>>>>>>\n", batchCount)

		outStream := outputStream{}
		//merge the batches and send to out-stream
		outStream.out = merge(batches)
		outStream.print()

		//increase start time to get the next 1-minute data batches
		startTime = startTime.Add(batchDuration)
		batchCount++
		time.Sleep(2 * time.Second)
	}

}

func merge(batches [][]sensorData) []sensorData {
	if len(batches) < 2 {
		return nil
	}

	//finish the recursion
	if len(batches) == 2 {
		return mergePare(batches[0], batches[1])
	}
	//merge batches recursively
	return mergePare(merge(batches[:len(batches)-1]), batches[len(batches)-1])
}

func mergePare(frequentData []sensorData, infrequentData []sensorData) []sensorData {
	offset := 0
	result := make([]sensorData, 0, len(frequentData)+len(infrequentData))
	for j := 0; j < len(infrequentData); j++ {
		//try to find a collection of data in a frequency batch,
		//each element of which is less than the current element in the infrequency batch
		for i := 0; i < len(frequentData); i++ {
			//if the data in the frequency sensor has run out, complete the cycle
			if i+offset == len(frequentData) {
				result = append(result, frequentData[offset:]...)
				result = append(result, infrequentData[j:]...)
				return result
			}
			//copy the collection to the result when a late item of infrequent sensors is found
			if frequentData[i+offset].getMeasurementTime().After(infrequentData[j].getMeasurementTime()) {
				result = append(result, frequentData[offset:i+offset]...)
				result = append(result, infrequentData[j])
				offset += i
				break
			}
		}
	}
	result = append(result, frequentData[offset:]...)
	return result
}
