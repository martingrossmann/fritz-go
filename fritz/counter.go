package fritz

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
)

const (
	XPATH_SENT_DATA_YESTERDAY_DE     string = "//tr[td[text() = 'Gestern']]/td[@datalabel='Datenvolumen gesendet(MB)']/text()"
	XPATH_RECEIVED_DATA_YESTERDAY_DE string = "//tr[td[text() = 'Gestern']]/td[@datalabel='Datenvolumen empfangen(MB)']/text()"
)

func (fb *FritzBox) ReadOnlineCounter() (OnlineCounter, error) {

	client := fb.getHTTPClient()

	counter, err := fetchCounterInfo(client, fb.Host+"/data.lua", fb)
	if err != nil {
		return OnlineCounter{}, err
	}

	return counter, nil
}

func fetchCounterInfo(client *http.Client, url string, fb *FritzBox) (OnlineCounter, error) {

	counter := OnlineCounter{
		DayOfData: time.Now().AddDate(0, 0, -1),
	}
	payload := bytes.NewBufferString("xhr=1&sid=" + fb.session.SID + "&lang=de&page=netCnt&no_sidrenew=")
	resp, err := client.Post(url, "application/x-www-form-urlencoded", payload)
	if err != nil {
		return OnlineCounter{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return OnlineCounter{}, err
	}

	doc, err := htmlquery.Parse(strings.NewReader(string(body)))
	if err != nil {
		return OnlineCounter{}, err
	}

	dataSentNode := htmlquery.FindOne(doc, XPATH_SENT_DATA_YESTERDAY_DE)
	counter.DataSent = dataSentNode.Data

	dataReceivedNode := htmlquery.FindOne(doc, XPATH_RECEIVED_DATA_YESTERDAY_DE)
	counter.DataReceived = dataReceivedNode.Data

	return counter, nil
}
