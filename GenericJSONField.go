package db_sqlserver

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// GenericJSONField is used to handle generic json data in postgres
type GenericJSONField map[string]interface{}

// Scan convert the json field into our type
func (v *GenericJSONField) Scan(src interface{}) error {
	var b []byte
	switch src.(type) {
	case []byte:
		b = src.([]byte)
	case string:
		b = []byte(src.(string))
	case nil:
		b = make([]byte, 0)
	default:
		return errors.New("unsupported type")
	}

	return json.Unmarshal(b, v)
}

// Value try to get the string slice representation in database
func (v GenericJSONField) Value() (driver.Value, error) {
	return json.Marshal(v)
}
