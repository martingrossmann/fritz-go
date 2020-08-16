package fritz

import (
	"crypto/md5"
	"crypto/tls"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// PerformLogin performs a login and returns SessionInfo including
// the session id (SID) on success
func (fb *FritzBox) PerformLogin(adminPassword string) error {
	client := fb.getHTTPClient()

	session, err := fetchSessionInfo(client, fb.Host+"/login_sid.lua")
	if err != nil {
		return err
	}

	response := buildAuthHash(session.Challenge, adminPassword)

	url, err := url.Parse(fb.Host)
	if err != nil {
		return err
	}
	user := url.User.Username()
	url.User = nil

	session, err = fetchSessionInfo(client, url.String()+"/login_sid.lua?&username="+user+"&response="+response)
	if err != nil {
		return err
	}
	if session.SID == "0000000000000000" {
		return errors.New("Login not successful")
	}

	fb.session = session

	return nil
}

func fetchSessionInfo(client *http.Client, url string) (SessionInfo, error) {
	resp, err := client.Get(url)
	if err != nil {
		return SessionInfo{}, err
	}

	defer resp.Body.Close() // nolint: errcheck

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return SessionInfo{}, err
	}

	var sessionInfo SessionInfo
	err = xml.Unmarshal(body, &sessionInfo)
	if err != nil {
		return SessionInfo{}, err
	}

	return sessionInfo, nil
}

// Create challenge/password hash
func buildAuthHash(challenge string, password string) string {
	challengePassword := utf8ToUtf16(challenge + "-" + password)

	md5Response := md5.Sum([]byte(challengePassword)) // nolint: gas

	return challenge + "-" + fmt.Sprintf("%x", md5Response)
}

// FritzBox use UTF16 LittleEndian (aka UCS-2LE)
func utf8ToUtf16(input string) string {
	e := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	t := e.NewEncoder()

	outstr, _, err := transform.String(t, input)
	if err != nil {
		log.Fatal(err)
	}

	return outstr
}

func (fb *FritzBox) getHTTPClient() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: false,
	}

	if fb.Insecure {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true} // nolint: gas
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   2 * time.Minute,
	}

	return client
}
