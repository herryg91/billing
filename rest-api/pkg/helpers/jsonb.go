package helpers

import (
	"database/sql/driver"
	"encoding/json"
)

type JSONB struct{}

func (m JSONB) Value() (driver.Value, error) {
	val, err := json.Marshal(m)
	return string(val), err
}
func (JSONB) GormDataType() string {
	return "JSONB"
}
