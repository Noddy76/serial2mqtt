# serial2mqtt

![GitHub license](https://img.shields.io/github/license/Noddy76/serial2mqtt.svg)
![Go](https://github.com/Noddy76/serial2mqtt/workflows/Go/badge.svg)

A simple daemon to relay lines from a serial port to a MQTT topic.

## Usage
	
```
./serial2mqtt -serial /dev/ttyUSB0 -baud 9600 -broker tcp://mqtt.home:1883
```

| Option   | Default              | Description                    |
| -------- | -------------------- | ------------------------------ |
| serial   |                      | Serial port                    |
| baud     | 115200               | Serial baud rate               |
| bits     | 8                    | Serial data bits               |
| parity   | N                    | Serial parity                  |
| stopBits | 1                    | Serial stop bits               |
| broker   | tcp://localhost:1883 | MQTT Broker connext string     |
| clientid | serial2mqtt          | MQTT client ID                 |
| topic    | serial2mqtt          | MQTT topic to publish lines on |
