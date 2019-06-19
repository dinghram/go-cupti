// Code generated by "enumer -type=CUpti_MetricPropertyID -json -text -yaml -sql"; DO NOT EDIT.

//
package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

const _CUpti_MetricPropertyIDName = "CUPTI_METRIC_PROPERTY_MULTIPROCESSOR_COUNT"

var _CUpti_MetricPropertyIDIndex = [...]uint8{0, 42}

func (i CUpti_MetricPropertyID) String() string {
	if i < 0 || i >= CUpti_MetricPropertyID(len(_CUpti_MetricPropertyIDIndex)-1) {
		return fmt.Sprintf("CUpti_MetricPropertyID(%d)", i)
	}
	return _CUpti_MetricPropertyIDName[_CUpti_MetricPropertyIDIndex[i]:_CUpti_MetricPropertyIDIndex[i+1]]
}

var _CUpti_MetricPropertyIDValues = []CUpti_MetricPropertyID{0}

var _CUpti_MetricPropertyIDNameToValueMap = map[string]CUpti_MetricPropertyID{
	_CUpti_MetricPropertyIDName[0:42]: 0,
}

// CUpti_MetricPropertyIDString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func CUpti_MetricPropertyIDString(s string) (CUpti_MetricPropertyID, error) {
	if val, ok := _CUpti_MetricPropertyIDNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to CUpti_MetricPropertyID values", s)
}

// CUpti_MetricPropertyIDValues returns all values of the enum
func CUpti_MetricPropertyIDValues() []CUpti_MetricPropertyID {
	return _CUpti_MetricPropertyIDValues
}

// IsACUpti_MetricPropertyID returns "true" if the value is listed in the enum definition. "false" otherwise
func (i CUpti_MetricPropertyID) IsACUpti_MetricPropertyID() bool {
	for _, v := range _CUpti_MetricPropertyIDValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for CUpti_MetricPropertyID
func (i CUpti_MetricPropertyID) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for CUpti_MetricPropertyID
func (i *CUpti_MetricPropertyID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("CUpti_MetricPropertyID should be a string, got %s", data)
	}

	var err error
	*i, err = CUpti_MetricPropertyIDString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for CUpti_MetricPropertyID
func (i CUpti_MetricPropertyID) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for CUpti_MetricPropertyID
func (i *CUpti_MetricPropertyID) UnmarshalText(text []byte) error {
	var err error
	*i, err = CUpti_MetricPropertyIDString(string(text))
	return err
}

// MarshalYAML implements a YAML Marshaler for CUpti_MetricPropertyID
func (i CUpti_MetricPropertyID) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for CUpti_MetricPropertyID
func (i *CUpti_MetricPropertyID) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = CUpti_MetricPropertyIDString(s)
	return err
}

func (i CUpti_MetricPropertyID) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *CUpti_MetricPropertyID) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	str, ok := value.(string)
	if !ok {
		bytes, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("value is not a byte slice")
		}

		str = string(bytes[:])
	}

	val, err := CUpti_MetricPropertyIDString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
