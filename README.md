## description
Prometheus exporter for the DHT22 written in go.

## usage
In order to access GPIO on rasberry pi you need to execute the exporter as root.
```bash
./goatmo --help

  -dht-pin int
    	Pin where DHT is plugged in (default 4)
  -listen-address string
    	The address to listen on for HTTP requests. (default "0.0.0.0:9000")
  -metric-prefix string
    	The prefix wanted for the metric (default "goatmo")

```

example:

```bash
sudo ./goatmo -listen-address 0.0.0.0:9100 -metric-prefix home
```

### metrics
Two gauge metrics are available `goatmo_room_temperature` and `goatmo_room_humidity`

```
# HELP goatmo_room_humidity Room Humidity
# TYPE goatmo_room_humidity gauge
goatmo_room_humidity 59.20000076293945
# HELP goatmo_room_temperature Room Temperature
# TYPE goatmo_room_temperature gauge
goatmo_room_temperature 18.100000381469727
```

# license MIT
