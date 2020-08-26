package common

type Config struct {
	FritzHost      string `properties:"fb.host,default=http://fritz.box"`
	FritzPassword  string `properties:"fb.password"`
	CsvFile        string `properties:"csv.filename"`
	InfluxHost     string `properties:"influxdb.host"`
	InfluxDB       string `properties:"influxdb.db"`
	InfluxUser     string `properties:"influxdb.user"`
	InfluxPassword string `properties:"influxdb.password"`
}
