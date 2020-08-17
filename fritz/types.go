package fritz

import "time"

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
}

type OnlineCounter struct {
	DataSent     string
	DataReceived string
	DayOfData    time.Time
}
