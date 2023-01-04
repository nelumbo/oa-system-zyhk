package models

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type XTime struct {
	time.Time
}

type XDate struct {
	time.Time
}

type XForms struct {
	Data     interface{} `json:"data"`
	Total    int64       `json:"total"`
	PageSize int         `json:"pageSize"`
	PageNo   int         `json:"pageNo"`
}

func (t *XTime) UnmarshalJSON(data []byte) error {
	if len(data) == 2 || string(data) == "null" {
		return nil
	}
	var err error
	str := string(data)
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02T15:04:05Z", timeStr)
	*t = XTime{t1}
	return err
}

func (t XTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte(`""`), nil
	}
	output := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(output), nil
}

func (t XTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *XTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = XTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// func (t *XDate) UnmarshalJSON(data []byte) error {
// 	var err error
// 	str := string(data)
// 	timeStr := strings.Trim(str, "\"")
// 	t1, err := time.Parse("2006-01-02", timeStr)
// 	*t = XDate{t1}
// 	return err
// }

func (t XDate) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte(`""`), nil
	}
	output := fmt.Sprintf("\"%s\"", t.Format("2006-01-02"))
	return []byte(output), nil
}

func (t XDate) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *XDate) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = XDate{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
