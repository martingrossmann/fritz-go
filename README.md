# fritz-go

Access online counter information without TR064 using Golang and stores the result in a CSV file.

At the moment the counter data from the day before are stored in the CSV according the following format:

``
<date in ISO layout>, <data received in MB>, <data sent in MB>
``

## Why not TR064?
TR064 is a great protocol to interact with your Fritz.Box. Unfortunately in case of online counter metrics like sent and recieved traffic data TR064 has limitations.

Thats why [Mirco Ropic](http://www.apfel-z.net/spezial/kontakt/) shows a way to get all traffic data via http post requests: http://www.apfel-z.net/artikel/Fritz_Box_API_via_curl_wget/. 

## Usage

* Rename the file `settings_tmpl.conf` to `settings.conf` and fill out the correct values.
* Start the application 


## Thanks to

https://github.com/tisba/fritz-tls for inspirations.