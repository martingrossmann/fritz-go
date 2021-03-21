package fritz

import (
	"time"
)

// Stores general information about the FritzBox
type FritzBox struct {
	Host     string
	User     string
	Passw    string
	Insecure bool // If host is https allow insecure/invalid TLS certificates
	session  SessionInfo
}

// Stores the current FritzBox session
type SessionInfo struct {
	SID       string `xml:"SID"`
	Challenge string `xml:"Challenge"`
	Users     []User `xml:"Users"`
}

type User struct {
	// TODO does not work
	//LastUsed int    `xml:"last,attr,default=0"`
	User string `xml:"User"`
}

type OnlineCounter struct {
	DataSent     int
	DataReceived int
	DayOfData    time.Time
}
