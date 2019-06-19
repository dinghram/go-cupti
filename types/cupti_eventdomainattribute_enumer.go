// Code generated by "enumer -type=CUpti_EventDomainAttribute -json -text -yaml -sql"; DO NOT EDIT.

//
package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

const (
	_CUpti_EventDomainAttributeName_0 = "CUPTI_EVENT_DOMAIN_ATTR_NAMECUPTI_EVENT_DOMAIN_ATTR_INSTANCE_COUNT"
	_CUpti_EventDomainAttributeName_1 = "CUPTI_EVENT_DOMAIN_ATTR_TOTAL_INSTANCE_COUNTCUPTI_EVENT_DOMAIN_ATTR_COLLECTION_METHOD"
	_CUpti_EventDomainAttributeName_2 = "CUPTI_EVENT_DOMAIN_ATTR_FORCE_INT"
)

var (
	_CUpti_EventDomainAttributeIndex_0 = [...]uint8{0, 28, 66}
	_CUpti_EventDomainAttributeIndex_1 = [...]uint8{0, 44, 85}
	_CUpti_EventDomainAttributeIndex_2 = [...]uint8{0, 33}
)

func (i CUpti_EventDomainAttribute) String() string {
	switch {
	case 0 <= i && i <= 1:
		return _CUpti_EventDomainAttributeName_0[_CUpti_EventDomainAttributeIndex_0[i]:_CUpti_EventDomainAttributeIndex_0[i+1]]
	case 3 <= i && i <= 4:
		i -= 3
		return _CUpti_EventDomainAttributeName_1[_CUpti_EventDomainAttributeIndex_1[i]:_CUpti_EventDomainAttributeIndex_1[i+1]]
	case i == 2147483647:
		return _CUpti_EventDomainAttributeName_2
	default:
		return fmt.Sprintf("CUpti_EventDomainAttribute(%d)", i)
	}
}

var _CUpti_EventDomainAttributeValues = []CUpti_EventDomainAttribute{0, 1, 3, 4, 2147483647}

var _CUpti_EventDomainAttributeNameToValueMap = map[string]CUpti_EventDomainAttribute{
	_CUpti_EventDomainAttributeName_0[0:28]:  0,
	_CUpti_EventDomainAttributeName_0[28:66]: 1,
	_CUpti_EventDomainAttributeName_1[0:44]:  3,
	_CUpti_EventDomainAttributeName_1[44:85]: 4,
	_CUpti_EventDomainAttributeName_2[0:33]:  2147483647,
}

// CUpti_EventDomainAttributeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func CUpti_EventDomainAttributeString(s string) (CUpti_EventDomainAttribute, error) {
	if val, ok := _CUpti_EventDomainAttributeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to CUpti_EventDomainAttribute values", s)
}

// CUpti_EventDomainAttributeValues returns all values of the enum
func CUpti_EventDomainAttributeValues() []CUpti_EventDomainAttribute {
	return _CUpti_EventDomainAttributeValues
}

// IsACUpti_EventDomainAttribute returns "true" if the value is listed in the enum definition. "false" otherwise
func (i CUpti_EventDomainAttribute) IsACUpti_EventDomainAttribute() bool {
	for _, v := range _CUpti_EventDomainAttributeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for CUpti_EventDomainAttribute
func (i CUpti_EventDomainAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for CUpti_EventDomainAttribute
func (i *CUpti_EventDomainAttribute) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("CUpti_EventDomainAttribute should be a string, got %s", data)
	}

	var err error
	*i, err = CUpti_EventDomainAttributeString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for CUpti_EventDomainAttribute
func (i CUpti_EventDomainAttribute) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for CUpti_EventDomainAttribute
func (i *CUpti_EventDomainAttribute) UnmarshalText(text []byte) error {
	var err error
	*i, err = CUpti_EventDomainAttributeString(string(text))
	return err
}

// MarshalYAML implements a YAML Marshaler for CUpti_EventDomainAttribute
func (i CUpti_EventDomainAttribute) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for CUpti_EventDomainAttribute
func (i *CUpti_EventDomainAttribute) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = CUpti_EventDomainAttributeString(s)
	return err
}

func (i CUpti_EventDomainAttribute) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *CUpti_EventDomainAttribute) Scan(value interface{}) error {
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

	val, err := CUpti_EventDomainAttributeString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}