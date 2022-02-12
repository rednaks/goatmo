package main

import (
  "flag"
  "net/http"
  "os"
  "os/signal"
  "syscall"
  "time"

  "github.com/prometheus/client_golang/prometheus"
  "github.com/prometheus/client_golang/prometheus/promauto"
  "github.com/prometheus/client_golang/prometheus/promhttp"
  dht "github.com/d2r2/go-dht"
  logger "github.com/d2r2/go-logger"
)

var lg = logger.NewPackageLogger("main",
logger.InfoLevel,
)

func getSensorValues() (hum, temp float32) {
  sensorType := dht.DHT22
  temp, hum, retried, err := dht.ReadDHTxxWithRetry(sensorType, *dht_pin, false, 10)
  if err != nil {
    lg.Info(err)
    return 0.0, 0.0
  }

  lg.Infof("Sensor = %v: Temperature = %v*C, Humidity = %v%%, (retried %d times)",
  sensorType, temp, hum, retried)
  return
}

func recordMetrics(humidity prometheus.Gauge, temperature prometheus.Gauge) {
  go func() {
    for {
      hum, temp := getSensorValues()
      humidity.Set(float64(hum))
      temperature.Set(float64(temp))
      time.Sleep(10 * time.Second)
    }
  }()
}

var addr = flag.String("listen-address", "0.0.0.0:9000", "The address to listen on for HTTP requests.")
var prefix = flag.String("metric-prefix", "goatmo", "The prefix wanted for the metric")
var dht_pin = flag.Int("dht-pin", 4, "Pin where DHT is plugged in")


func initMetrics() (h prometheus.Gauge, t prometheus.Gauge) {
  var (
    humidity = promauto.NewGauge(prometheus.GaugeOpts{
      Name: *prefix + "_room_humidity",
      Help: "Room Humidity",
    })

    temperature = promauto.NewGauge(prometheus.GaugeOpts{
      Name: *prefix + "_room_temperature",
      Help: "Room Temperature",
    })
  )

  return humidity, temperature
}

func initHandleInterrupt() {
  c := make(chan os.Signal)
  signal.Notify(c, os.Interrupt, syscall.SIGTERM)
  go func() {
    <-c
    lg.Infof("Exitinng ...")
    os.Exit(1)
  }()
}

func main() {

  flag.Parse()
  initHandleInterrupt()
  recordMetrics(initMetrics())

  http.Handle("/metrics", promhttp.Handler())
  http.ListenAndServe(*addr, nil)
}
