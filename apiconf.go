package gotrex

import (
	"encoding/json"
	"time"
)

const PROTOCOL = "https"
const ENDPOINT = "bittrex.com/api"
const APIVER = "v1.1"
const TIMEOUT = 30
const TIMEPARSE = "2006-01-02T15:04:05"

type jsonTime struct {
	time.Time
}

func (jt *jsonTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	t, err := time.Parse(TIMEPARSE, s)
	if err != nil {
		return err
	}
	jt.Time = t
	return nil
}

func (jt jsonTime) MarshalJSON() ([]byte, error) {
	return json.Marshal((*time.Time)(&jt.Time).Format(TIMEPARSE))
}
