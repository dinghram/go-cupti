// Code generated by "enumer -type=CUpti_MetricValueUtilizationLevel -json -text -yaml -sql"; DO NOT EDIT.

//
package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

const (
	_CUpti_MetricValueUtilizationLevelName_0 = "CUPTI_METRIC_VALUE_UTILIZATION_IDLE"
	_CUpti_MetricValueUtilizationLevelName_1 = "CUPTI_METRIC_VALUE_UTILIZATION_LOW"
	_CUpti_MetricValueUtilizationLevelName_2 = "CUPTI_METRIC_VALUE_UTILIZATION_MID"
	_CUpti_MetricValueUtilizationLevelName_3 = "CUPTI_METRIC_VALUE_UTILIZATION_HIGH"
	_CUpti_MetricValueUtilizationLevelName_4 = "CUPTI_METRIC_VALUE_UTILIZATION_MAX"
	_CUpti_MetricValueUtilizationLevelName_5 = "CUPTI_METRIC_VALUE_UTILIZATION_FORCE_INT"
)

var (
	_CUpti_MetricValueUtilizationLevelIndex_0 = [...]uint8{0, 35}
	_CUpti_MetricValueUtilizationLevelIndex_1 = [...]uint8{0, 34}
	_CUpti_MetricValueUtilizationLevelIndex_2 = [...]uint8{0, 34}
	_CUpti_MetricValueUtilizationLevelIndex_3 = [...]uint8{0, 35}
	_CUpti_MetricValueUtilizationLevelIndex_4 = [...]uint8{0, 34}
	_CUpti_MetricValueUtilizationLevelIndex_5 = [...]uint8{0, 40}
)

func (i CUpti_MetricValueUtilizationLevel) String() string {
	switch {
	case i == 0:
		return _CUpti_MetricValueUtilizationLevelName_0
	case i == 2:
		return _CUpti_MetricValueUtilizationLevelName_1
	case i == 5:
		return _CUpti_MetricValueUtilizationLevelName_2
	case i == 8:
		return _CUpti_MetricValueUtilizationLevelName_3
	case i == 10:
		return _CUpti_MetricValueUtilizationLevelName_4
	case i == 2147483647:
		return _CUpti_MetricValueUtilizationLevelName_5
	default:
		return fmt.Sprintf("CUpti_MetricValueUtilizationLevel(%d)", i)
	}
}

var _CUpti_MetricValueUtilizationLevelValues = []CUpti_MetricValueUtilizationLevel{0, 2, 5, 8, 10, 2147483647}

var _CUpti_MetricValueUtilizationLevelNameToValueMap = map[string]CUpti_MetricValueUtilizationLevel{
	_CUpti_MetricValueUtilizationLevelName_0[0:35]: 0,
	_CUpti_MetricValueUtilizationLevelName_1[0:34]: 2,
	_CUpti_MetricValueUtilizationLevelName_2[0:34]: 5,
	_CUpti_MetricValueUtilizationLevelName_3[0:35]: 8,
	_CUpti_MetricValueUtilizationLevelName_4[0:34]: 10,
	_CUpti_MetricValueUtilizationLevelName_5[0:40]: 2147483647,
}

// CUpti_MetricValueUtilizationLevelString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func CUpti_MetricValueUtilizationLevelString(s string) (CUpti_MetricValueUtilizationLevel, error) {
	if val, ok := _CUpti_MetricValueUtilizationLevelNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to CUpti_MetricValueUtilizationLevel values", s)
}

// CUpti_MetricValueUtilizationLevelValues returns all values of the enum
func CUpti_MetricValueUtilizationLevelValues() []CUpti_MetricValueUtilizationLevel {
	return _CUpti_MetricValueUtilizationLevelValues
}

// IsACUpti_MetricValueUtilizationLevel returns "true" if the value is listed in the enum definition. "false" otherwise
func (i CUpti_MetricValueUtilizationLevel) IsACUpti_MetricValueUtilizationLevel() bool {
	for _, v := range _CUpti_MetricValueUtilizationLevelValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for CUpti_MetricValueUtilizationLevel
func (i CUpti_MetricValueUtilizationLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for CUpti_MetricValueUtilizationLevel
func (i *CUpti_MetricValueUtilizationLevel) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("CUpti_MetricValueUtilizationLevel should be a string, got %s", data)
	}

	var err error
	*i, err = CUpti_MetricValueUtilizationLevelString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for CUpti_MetricValueUtilizationLevel
func (i CUpti_MetricValueUtilizationLevel) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for CUpti_MetricValueUtilizationLevel
func (i *CUpti_MetricValueUtilizationLevel) UnmarshalText(text []byte) error {
	var err error
	*i, err = CUpti_MetricValueUtilizationLevelString(string(text))
	return err
}

// MarshalYAML implements a YAML Marshaler for CUpti_MetricValueUtilizationLevel
func (i CUpti_MetricValueUtilizationLevel) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for CUpti_MetricValueUtilizationLevel
func (i *CUpti_MetricValueUtilizationLevel) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = CUpti_MetricValueUtilizationLevelString(s)
	return err
}

func (i CUpti_MetricValueUtilizationLevel) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *CUpti_MetricValueUtilizationLevel) Scan(value interface{}) error {
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

	val, err := CUpti_MetricValueUtilizationLevelString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}