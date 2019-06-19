// Code generated by "enumer -type=CUpti_EventCollectionMode -json -text -yaml -sql"; DO NOT EDIT.

//
package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

const (
	_CUpti_EventCollectionModeName_0 = "CUPTI_EVENT_COLLECTION_MODE_CONTINUOUSCUPTI_EVENT_COLLECTION_MODE_KERNEL"
	_CUpti_EventCollectionModeName_1 = "CUPTI_EVENT_COLLECTION_MODE_FORCE_INT"
)

var (
	_CUpti_EventCollectionModeIndex_0 = [...]uint8{0, 38, 72}
	_CUpti_EventCollectionModeIndex_1 = [...]uint8{0, 37}
)

func (i CUpti_EventCollectionMode) String() string {
	switch {
	case 0 <= i && i <= 1:
		return _CUpti_EventCollectionModeName_0[_CUpti_EventCollectionModeIndex_0[i]:_CUpti_EventCollectionModeIndex_0[i+1]]
	case i == 2147483647:
		return _CUpti_EventCollectionModeName_1
	default:
		return fmt.Sprintf("CUpti_EventCollectionMode(%d)", i)
	}
}

var _CUpti_EventCollectionModeValues = []CUpti_EventCollectionMode{0, 1, 2147483647}

var _CUpti_EventCollectionModeNameToValueMap = map[string]CUpti_EventCollectionMode{
	_CUpti_EventCollectionModeName_0[0:38]:  0,
	_CUpti_EventCollectionModeName_0[38:72]: 1,
	_CUpti_EventCollectionModeName_1[0:37]:  2147483647,
}

// CUpti_EventCollectionModeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func CUpti_EventCollectionModeString(s string) (CUpti_EventCollectionMode, error) {
	if val, ok := _CUpti_EventCollectionModeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to CUpti_EventCollectionMode values", s)
}

// CUpti_EventCollectionModeValues returns all values of the enum
func CUpti_EventCollectionModeValues() []CUpti_EventCollectionMode {
	return _CUpti_EventCollectionModeValues
}

// IsACUpti_EventCollectionMode returns "true" if the value is listed in the enum definition. "false" otherwise
func (i CUpti_EventCollectionMode) IsACUpti_EventCollectionMode() bool {
	for _, v := range _CUpti_EventCollectionModeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for CUpti_EventCollectionMode
func (i CUpti_EventCollectionMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for CUpti_EventCollectionMode
func (i *CUpti_EventCollectionMode) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("CUpti_EventCollectionMode should be a string, got %s", data)
	}

	var err error
	*i, err = CUpti_EventCollectionModeString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for CUpti_EventCollectionMode
func (i CUpti_EventCollectionMode) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for CUpti_EventCollectionMode
func (i *CUpti_EventCollectionMode) UnmarshalText(text []byte) error {
	var err error
	*i, err = CUpti_EventCollectionModeString(string(text))
	return err
}

// MarshalYAML implements a YAML Marshaler for CUpti_EventCollectionMode
func (i CUpti_EventCollectionMode) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for CUpti_EventCollectionMode
func (i *CUpti_EventCollectionMode) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = CUpti_EventCollectionModeString(s)
	return err
}

func (i CUpti_EventCollectionMode) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *CUpti_EventCollectionMode) Scan(value interface{}) error {
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

	val, err := CUpti_EventCollectionModeString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
