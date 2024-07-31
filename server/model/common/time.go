package common

import (
	"database/sql/driver"
	"errors"
	"time"
)

type CustomTime struct {
	time.Time
}

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = s[1 : len(s)-1]                         // 去除双引号
	t, err := time.Parse("2006-01-02 15:04", s) // 自定义时间格式
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}

// Implement driver.Valuer interface for CustomTime
func (ct CustomTime) Value() (driver.Value, error) {
	return ct.Time, nil
}

// Implement sql.Scanner interface for CustomTime
func (ct *CustomTime) Scan(value interface{}) error {
	if value == nil {
		ct.Time = time.Time{}
		return nil
	}

	parsedTime, ok := value.(time.Time)
	if !ok {
		return errors.New("invalid type for customTime")
	}

	ct.Time = parsedTime
	return nil
}
