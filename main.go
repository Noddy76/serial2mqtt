/*
Copyright 2020 James Grant

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/goburrow/serial"
)

func main() {
	var serialPort = flag.String("serial", "", "Serial port")
	var baudRate = flag.Int("baud", 115200, "Serial baud rate")
	var dataBits = flag.Int("bits", 8, "Serial data bits")
	var parity = flag.String("parity", "N", "Serial parity")
	var stopBits = flag.Int("stopBits", 1, "Serial stop bits")

	var mqttBroker = flag.String("broker", "tcp://localhost:1883", "MQTT Broker connext string")
	var mqttClientID = flag.String("clientid", "serial2mqtt", "MQTT client ID")
	var mqttTopic = flag.String("topic", "serial2mqtt", "MQTT topic to publish lines on")

	flag.Parse()

	opts := mqtt.NewClientOptions().
		AddBroker(*mqttBroker).
		SetClientID(*mqttClientID)
	mqttClient := mqtt.NewClient(opts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	defer mqttClient.Disconnect(250)

	serialConfig := &serial.Config{
		Address:  *serialPort,
		BaudRate: *baudRate,
		DataBits: *dataBits,
		Parity:   *parity,
		StopBits: *stopBits,
	}

	port, err := serial.Open(serialConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	scanner := bufio.NewScanner(port)

	for scanner.Scan() {
		var text = scanner.Text()
		fmt.Println(text)
		token := mqttClient.Publish(*mqttTopic, 0, false, text)
		token.Wait()
	}
}
