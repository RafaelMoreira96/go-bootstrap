package utils

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// Date é um tipo para representar datas sem a parte do tempo
type Date struct {
	time.Time
}

const DateFormat = "2006-01-02"

// Scan implementa a interface sql.Scanner
func (d *Date) Scan(value interface{}) error {
	if value == nil {
		*d = Date{Time: time.Time{}}
		return nil
	}
	parsedTime, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("can't scan value into Date: %v", value)
	}
	*d = Date{Time: parsedTime}
	return nil
}

// Value implementa a interface driver.Valuer
func (d Date) Value() (driver.Value, error) {
	return d.Time, nil
}

// MarshalJSON implementa a interface json.Marshaler
func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Format(DateFormat))
}

// UnmarshalJSON implementa a interface json.Unmarshaler
func (d *Date) UnmarshalJSON(data []byte) error {
	var err error
	parsedTime, err := time.Parse(`"`+DateFormat+`"`, string(data))
	if err != nil {
		return err
	}
	d.Time = parsedTime
	return nil
}

func Today() Date {
	return Date{Time: time.Now()}
}
