package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

func (unit TemperatureUnit) String() string {
	return []string{"°C", "°F"}[unit]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (temperature Temperature) String() string {
	return fmt.Sprintf("%d %s", temperature.degree, temperature.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

func (unit SpeedUnit) String() string {
	return []string{"km/h", "mph"}[unit]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

func (speed Speed) String() string {
	return fmt.Sprintf("%d %s", speed.magnitude, speed.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

func (data MeteorologyData) String() string {
	return fmt.Sprintf("%s: %s, Wind %s at %s, %d%% Humidity",
		data.location, data.temperature, data.windDirection, data.windSpeed, data.humidity)
}
