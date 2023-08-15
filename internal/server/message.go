package server

type Message struct {
	SensorId     int     `json:"sensor_id"`
	PressurePa   float64 `json:"pressure_pa"`
	HumidityRh   float64 `json:"humidity_rh"`
	TemperatureC float64 `json:"temperature_c"`
	DewPointC    float64 `json:"dew_point_c"`
	Time         string  `json:"time"`
}
