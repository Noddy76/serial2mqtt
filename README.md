# serial2mqtt

A simple daemon to relay lines from a serial port to a MQTT topic.

## Usage
	
```
./serial2mqtt -serial /dev/ttyUSB0 -baud 9600 -broker tcp://mqtt.home:1883
```

| Option | Default | Description |
| --- | --- | --- |
| serial |  | Serial port |
| baud | 115200 | Serial baud rate |
| bits | 8 | Serial data bits |
| parity | parity | N | Serial parity |
| stopBits | stopBits | 1 | Serial stop bits |
| mqttBroker | broker | tcp://localhost:1883 | MQTT Broker connext string |
| mqttClientID | clientid | serial2mqtt | MQTT client ID |
| mqttTopic | topic | serial2mqtt | MQTT topic to publish lines on |
