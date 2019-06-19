// Code generated by "enumer -type=CUpti_DevType -json -text -yaml -sql"; DO NOT EDIT.

//
package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

const (
	_CUpti_DevTypeName_0 = "CUPTI_DEV_TYPE_INVALIDCUPTI_DEV_TYPE_GPUCUPTI_DEV_TYPE_NPU"
	_CUpti_DevTypeName_1 = "CUPTI_DEV_TYPE_FORCE_INT"
)

var (
	_CUpti_DevTypeIndex_0 = [...]uint8{0, 22, 40, 58}
	_CUpti_DevTypeIndex_1 = [...]uint8{0, 24}
)

func (i CUpti_DevType) String() string {
	switch {
	case 0 <= i && i <= 2:
		return _CUpti_DevTypeName_0[_CUpti_DevTypeIndex_0[i]:_CUpti_DevTypeIndex_0[i+1]]
	case i == 2147483647:
		return _CUpti_DevTypeName_1
	default:
		return fmt.Sprintf("CUpti_DevType(%d)", i)
	}
}

var _CUpti_DevTypeValues = []CUpti_DevType{0, 1, 2, 2147483647}

var _CUpti_DevTypeNameToValueMap = map[string]CUpti_DevType{
	_CUpti_DevTypeName_0[0:22]:  0,
	_CUpti_DevTypeName_0[22:40]: 1,
	_CUpti_DevTypeName_0[40:58]: 2,
	_CUpti_DevTypeName_1[0:24]:  2147483647,
}

// CUpti_DevTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func CUpti_DevTypeString(s string) (CUpti_DevType, error) {
	if val, ok := _CUpti_DevTypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to CUpti_DevType values", s)
}

// CUpti_DevTypeValues returns all values of the enum
func CUpti_DevTypeValues() []CUpti_DevType {
	return _CUpti_DevTypeValues
}

// IsACUpti_DevType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i CUpti_DevType) IsACUpti_DevType() bool {
	for _, v := range _CUpti_DevTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for CUpti_DevType
func (i CUpti_DevType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for CUpti_DevType
func (i *CUpti_DevType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("CUpti_DevType should be a string, got %s", data)
	}

	var err error
	*i, err = CUpti_DevTypeString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for CUpti_DevType
func (i CUpti_DevType) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for CUpti_DevType
func (i *CUpti_DevType) UnmarshalText(text []byte) error {
	var err error
	*i, err = CUpti_DevTypeString(string(text))
	return err
}

// MarshalYAML implements a YAML Marshaler for CUpti_DevType
func (i CUpti_DevType) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for CUpti_DevType
func (i *CUpti_DevType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = CUpti_DevTypeString(s)
	return err
}

func (i CUpti_DevType) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *CUpti_DevType) Scan(value interface{}) error {
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

	val, err := CUpti_DevTypeString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
