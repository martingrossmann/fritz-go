# fritz-go

Access online counter information without TR064 using Golang. The results can stored in
 
 * a CSV file in the format `<date in ISO layout>, <data received in MB>, <data sent in MB>` _and/or_
 * to an InfluxDB and can visualized with Grafana.

## Why not TR064?
TR064 is a great protocol to interact with your Fritz.Box. Unfortunately in case of online counter metrics like sent and recieved traffic data TR064 has limitations.

Thats why [Mirco Ropic](http://www.apfel-z.net/spezial/kontakt/) shows a way to get all traffic data via http post requests: http://www.apfel-z.net/artikel/Fritz_Box_API_via_curl_wget/. 

## Installation

### InfluxDB/Grafana

* Check out the complete repository
* Copy the ``docker/.env_templ`` to ``docker/.env``
* Fill out the needed values for credentials and other stuff
* Run
  ````
  cd docker
  docker-compose up -d
  ````
* Setup in Grafana the datasource for the InfluxDB 
* You can import the Grafana dashboard ``docker/grafana_dashboard/dashboard.json``

### Build the app

* Run ``go build`` to create an executable file for your system _or_
* Run ``build.cmd`` (sorry, it's a Windows batch) to create 
  * fritz-go.exe
  * fritz-go as Linux/ARM binary.
  
The Linux/ARM binary can used on Synology NAS systems build on ARM (tested with DS218play).

## Usage

* Rename the file ``settings_tmpl.conf`` to ``settings.conf`` and fill out the correct values.
* Start the application with ``fritz-go``. The app connects to your Fritz box, identify your traffic and store the data to CSV and/or to an InfluxDB.

### Parameter

``fritz-go convert`` creates an import file for InfluxDB based on existing CSV file (see ``settings.conf`` for filename).

You can import that file to InfluxDB with the following command:

````shell
curl -i -XPOST 'http://localhost:8086/api/v2/write?bucket=fritzdata&precision=s' --header 'Authorization: Token fritz:<passwpord>' --data-binary @influxdbdata
````

where ``fritzdata`` is your InfluxDB db (=bucket) and ``influxdbdata`` your import file

## Thanks to

https://github.com/tisba/fritz-tls for inspirations.
