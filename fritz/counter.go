package fritz

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
)

// FritzOS >= 7.25
// Needed for parsing the JSON part taken from HTML source code
type CounterData struct {
	Yesterday struct {
		Bytessenthigh     string `json:"BytesSentHigh"`
		Bytessentlow      string `json:"BytesSentLow"`
		Bytesreceivedhigh string `json:"BytesReceivedHigh"`
		Bytesreceivedlow  string `json:"BytesReceivedLow"`
	} `json:"Yesterday"`
	Today struct {
		Bytessenthigh     string `json:"BytesSentHigh"`
		Bytessentlow      string `json:"BytesSentLow"`
		Bytesreceivedhigh string `json:"BytesReceivedHigh"`
		Bytesreceivedlow  string `json:"BytesReceivedLow"`
	} `json:"Today"`
}

// FritzOS <= 7.21
// Needed for parsing values from HTML table
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

	// return fetchCounterInfoWithXpath(string(body))

	return fetchCounterInfoWithRegex(string(body))
}

// Up to FritzOS 7.21 the information are part of an HTML table
func fetchCounterInfoWithXpath(content string) (OnlineCounter, error) {
	tY := time.Now().AddDate(0, 0, -1)
	rounded := time.Date(tY.Year(), tY.Month(), tY.Day(), 12, 0, 0, 0, tY.Location())

	counter := OnlineCounter{
		DayOfData: rounded,
	}

	doc, err := htmlquery.Parse(strings.NewReader(content))
	if err != nil {
		return OnlineCounter{}, err
	}

	// htmlquery.FindOne(doc, "//tr[@id = 'uiYesterday']/td[contains(@class, 'vol-outgoing')]")
	dataSentNode := htmlquery.FindOne(doc, XPATH_SENT_DATA_YESTERDAY_DE)
	counter.DataSent, err = strconv.Atoi(dataSentNode.Data)
	if err != nil {
		return OnlineCounter{}, err
	}

	dataReceivedNode := htmlquery.FindOne(doc, XPATH_RECEIVED_DATA_YESTERDAY_DE)
	counter.DataReceived, err = strconv.Atoi(dataReceivedNode.Data)
	if err != nil {
		return OnlineCounter{}, err
	}

	log.Println("Counter data read from Fritz.Box.")
	log.Println("Yesterday:", counter)

	return counter, nil
}

// With FritzOS 7.25 and higher the information are stored in a JavaScript object within the HTML page
func fetchCounterInfoWithRegex(content string) (OnlineCounter, error) {
	tY := time.Now().AddDate(0, 0, -1)
	yesterdaysDate := time.Date(tY.Year(), tY.Month(), tY.Day(), 12, 0, 0, 0, tY.Location())

	counter := OnlineCounter{
		DayOfData: yesterdaysDate,
	}

	re := regexp.MustCompile(`const data =(.*);`)
	result := re.FindStringSubmatch(content)
	if len(result) != 2 {
		return OnlineCounter{}, errors.New("No counter data found.")
	}
	var cData CounterData
	err := json.Unmarshal([]byte(result[1]), &cData)
	if err != nil {
		return OnlineCounter{}, err
	}
	receivedBytesLow, _ := strconv.Atoi(cData.Yesterday.Bytesreceivedlow)
	receivedBytesHigh, _ := strconv.Atoi(cData.Yesterday.Bytesreceivedhigh)
	sentBytesLow, _ := strconv.Atoi(cData.Yesterday.Bytessentlow)
	setnBytesHigh, _ := strconv.Atoi(cData.Yesterday.Bytessenthigh)
	counter.DataReceived = calculateBytes(receivedBytesHigh, receivedBytesLow)
	counter.DataSent = calculateBytes(setnBytesHigh, sentBytesLow)

	return counter, nil
}

// Inspired from FritzBox HTML source code
// Includes MB calculating
func calculateBytes(high int, low int) int {
	b := int64(high)*4294967296 + int64(low)
	return int(b / 1000000)
}
