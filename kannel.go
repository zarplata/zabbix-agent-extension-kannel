package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	hierr "github.com/reconquest/hierr-go"
)

// Stats kannel
type Stats struct {
	Status    string      `xml:"status"`
	WPD       WPDStats    `xml:"wpd"`
	SMS       SMSStats    `xml:"sms"`
	DLR       DLRStats    `xml:"dlr"`
	SMSCCount int         `xml:"smscs>count"`
	SMSC      []SMSCStats `xml:"smscs>smsc"`
}

// WPDStats struct
type WPDStats struct {
	Recv       int `xml:"received>total"`
	RecvQueued int `xml:"received>queued"`
	Sent       int `xml:"sent>total"`
	SentQueued int `xml:"sent>queued"`
}

// SMSStats struct
type SMSStats struct {
	Recv       int    `xml:"received>total"`
	RecvQueued int    `xml:"received>queued"`
	Sent       int    `xml:"sent>total"`
	SentQueued int    `xml:"sent>queued"`
	StoreSize  int    `xml:"storesize"`
	Inbound    string `xml:"inbound"`
	Outbound   string `xml:"outbound"`
}

// DLRStats struct
type DLRStats struct {
	Queued  int    `xml:"queued"`
	Storage string `xml:"storage"`
}

// SMSCStats struct
type SMSCStats struct {
	Name   string `xml:"name"`
	ID     string `xml:"id"`
	Status string `xml:"status"`
	Uptime float64
	Recv   int `xml:"received"`
	Sent   int `xml:"sent"`
	Failed int `xml:"failed"`
	Queued int `xml:"queued"`
}

func getKannelStats(kannel string) (*Stats, error) {

	var (
		stats    Stats
		provider SMSCStats
		result   Stats
	)

	timeout := time.Duration(3 * time.Second)
	client := http.Client{Timeout: timeout}

	ret, err := client.Get(kannel)
	if err != nil {
		return &result, hierr.Errorf(
			err,
			"can`t get kannel stats %s",
			kannel,
		)
	}

	defer ret.Body.Close()

	if ret.StatusCode != http.StatusOK {
		return &result, fmt.Errorf(
			"can`t get status, return %d HTTP code, expected %d HTTP code",
			ret.StatusCode,
			http.StatusOK,
		)
	}

	body, err := ioutil.ReadAll(ret.Body)
	if err != nil {
		return &result, hierr.Errorf(err, "can't read response body")
	}

	err = xml.Unmarshal(body, &result)
	if err != nil {
		return &result, hierr.Errorf(err, "can't unmarshal body")
	}

	//cut all after comma "running, uptime 33d 2h 5m 18s"
	result.Status = strings.Split(result.Status, ", ")[0]

	for _, provider = range result.SMSC {
		//"online 15356s": field[0] - status, field[1] - uptime
		if strings.Contains(provider.Status, "online") {
			uptime, _ := time.ParseDuration(strings.Fields(provider.Status)[1])
			provider.Uptime = uptime.Seconds()
			provider.Status = strings.Fields(provider.Status)[0]
		}
		stats.SMSC = append(stats.SMSC, provider)
	}
	result.SMSC = stats.SMSC

	return &result, nil
}
