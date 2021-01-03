package common

type Config struct {
	FritzHost           string `properties:"fb.host,default=http://fritz.box"`
	FritzPassword       string `properties:"fb.password"`
	CsvActive           bool   `properties:"csv.active,default=true"`
	CsvFile             string `properties:"csv.filename"`
	InfluxActive        bool   `properties:"influxdb.active,default=false"`
	InfluxHost          string `properties:"influxdb.host"`
	InfluxDB            string `properties:"influxdb.db"`
	InfluxUser          string `properties:"influxdb.user"`
	InfluxPassword      string `properties:"influxdb.password"`
	InfluxTcMeasurement string `properties:"influxdb.traffic.measurement"`
	InfluxTcReceived    string `properties:"influxdb.traffic.received"`
	InfluxTcSent        string `properties:"influxdb.traffic.sent"`
	InfluxImportFile    string `properties:"influxdb.importfile"`
}
