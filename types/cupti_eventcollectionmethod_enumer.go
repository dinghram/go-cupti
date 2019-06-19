// Code generated by "enumer -type=CUpti_EventCollectionMethod -json -text -yaml -sql"; DO NOT EDIT.

//
package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

const (
	_CUpti_EventCollectionMethodName_0 = "CUPTI_EVENT_COLLECTION_METHOD_PMCUPTI_EVENT_COLLECTION_METHOD_SMCUPTI_EVENT_COLLECTION_METHOD_INSTRUMENTEDCUPTI_EVENT_COLLECTION_METHOD_NVLINK_TC"
	_CUpti_EventCollectionMethodName_1 = "CUPTI_EVENT_COLLECTION_METHOD_FORCE_INT"
)

var (
	_CUpti_EventCollectionMethodIndex_0 = [...]uint8{0, 32, 64, 106, 145}
	_CUpti_EventCollectionMethodIndex_1 = [...]uint8{0, 39}
)

func (i CUpti_EventCollectionMethod) String() string {
	switch {
	case 0 <= i && i <= 3:
		return _CUpti_EventCollectionMethodName_0[_CUpti_EventCollectionMethodIndex_0[i]:_CUpti_EventCollectionMethodIndex_0[i+1]]
	case i == 2147483647:
		return _CUpti_EventCollectionMethodName_1
	default:
		return fmt.Sprintf("CUpti_EventCollectionMethod(%d)", i)
	}
}

var _CUpti_EventCollectionMethodValues = []CUpti_EventCollectionMethod{0, 1, 2, 3, 2147483647}

var _CUpti_EventCollectionMethodNameToValueMap = map[string]CUpti_EventCollectionMethod{
	_CUpti_EventCollectionMethodName_0[0:32]:    0,
	_CUpti_EventCollectionMethodName_0[32:64]:   1,
	_CUpti_EventCollectionMethodName_0[64:106]:  2,
	_CUpti_EventCollectionMethodName_0[106:145]: 3,
	_CUpti_EventCollectionMethodName_1[0:39]:    2147483647,
}

// CUpti_EventCollectionMethodString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func CUpti_EventCollectionMethodString(s string) (CUpti_EventCollectionMethod, error) {
	if val, ok := _CUpti_EventCollectionMethodNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to CUpti_EventCollectionMethod values", s)
}

// CUpti_EventCollectionMethodValues returns all values of the enum
func CUpti_EventCollectionMethodValues() []CUpti_EventCollectionMethod {
	return _CUpti_EventCollectionMethodValues
}

// IsACUpti_EventCollectionMethod returns "true" if the value is listed in the enum definition. "false" otherwise
func (i CUpti_EventCollectionMethod) IsACUpti_EventCollectionMethod() bool {
	for _, v := range _CUpti_EventCollectionMethodValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for CUpti_EventCollectionMethod
func (i CUpti_EventCollectionMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for CUpti_EventCollectionMethod
func (i *CUpti_EventCollectionMethod) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("CUpti_EventCollectionMethod should be a string, got %s", data)
	}

	var err error
	*i, err = CUpti_EventCollectionMethodString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for CUpti_EventCollectionMethod
func (i CUpti_EventCollectionMethod) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for CUpti_EventCollectionMethod
func (i *CUpti_EventCollectionMethod) UnmarshalText(text []byte) error {
	var err error
	*i, err = CUpti_EventCollectionMethodString(string(text))
	return err
}

// MarshalYAML implements a YAML Marshaler for CUpti_EventCollectionMethod
func (i CUpti_EventCollectionMethod) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for CUpti_EventCollectionMethod
func (i *CUpti_EventCollectionMethod) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = CUpti_EventCollectionMethodString(s)
	return err
}

func (i CUpti_EventCollectionMethod) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *CUpti_EventCollectionMethod) Scan(value interface{}) error {
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

	val, err := CUpti_EventCollectionMethodString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
