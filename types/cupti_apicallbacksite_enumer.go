// Code generated by "enumer -type=CUpti_ApiCallbackSite -json -text -yaml -sql"; DO NOT EDIT.

//
package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

const (
	_CUpti_ApiCallbackSiteName_0 = "CUPTI_API_ENTERCUPTI_API_EXIT"
	_CUpti_ApiCallbackSiteName_1 = "CUPTI_API_CBSITE_FORCE_INT"
)

var (
	_CUpti_ApiCallbackSiteIndex_0 = [...]uint8{0, 15, 29}
	_CUpti_ApiCallbackSiteIndex_1 = [...]uint8{0, 26}
)

func (i CUpti_ApiCallbackSite) String() string {
	switch {
	case 0 <= i && i <= 1:
		return _CUpti_ApiCallbackSiteName_0[_CUpti_ApiCallbackSiteIndex_0[i]:_CUpti_ApiCallbackSiteIndex_0[i+1]]
	case i == 2147483647:
		return _CUpti_ApiCallbackSiteName_1
	default:
		return fmt.Sprintf("CUpti_ApiCallbackSite(%d)", i)
	}
}

var _CUpti_ApiCallbackSiteValues = []CUpti_ApiCallbackSite{0, 1, 2147483647}

var _CUpti_ApiCallbackSiteNameToValueMap = map[string]CUpti_ApiCallbackSite{
	_CUpti_ApiCallbackSiteName_0[0:15]:  0,
	_CUpti_ApiCallbackSiteName_0[15:29]: 1,
	_CUpti_ApiCallbackSiteName_1[0:26]:  2147483647,
}

// CUpti_ApiCallbackSiteString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func CUpti_ApiCallbackSiteString(s string) (CUpti_ApiCallbackSite, error) {
	if val, ok := _CUpti_ApiCallbackSiteNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to CUpti_ApiCallbackSite values", s)
}

// CUpti_ApiCallbackSiteValues returns all values of the enum
func CUpti_ApiCallbackSiteValues() []CUpti_ApiCallbackSite {
	return _CUpti_ApiCallbackSiteValues
}

// IsACUpti_ApiCallbackSite returns "true" if the value is listed in the enum definition. "false" otherwise
func (i CUpti_ApiCallbackSite) IsACUpti_ApiCallbackSite() bool {
	for _, v := range _CUpti_ApiCallbackSiteValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for CUpti_ApiCallbackSite
func (i CUpti_ApiCallbackSite) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for CUpti_ApiCallbackSite
func (i *CUpti_ApiCallbackSite) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("CUpti_ApiCallbackSite should be a string, got %s", data)
	}

	var err error
	*i, err = CUpti_ApiCallbackSiteString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for CUpti_ApiCallbackSite
func (i CUpti_ApiCallbackSite) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for CUpti_ApiCallbackSite
func (i *CUpti_ApiCallbackSite) UnmarshalText(text []byte) error {
	var err error
	*i, err = CUpti_ApiCallbackSiteString(string(text))
	return err
}

// MarshalYAML implements a YAML Marshaler for CUpti_ApiCallbackSite
func (i CUpti_ApiCallbackSite) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for CUpti_ApiCallbackSite
func (i *CUpti_ApiCallbackSite) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = CUpti_ApiCallbackSiteString(s)
	return err
}

func (i CUpti_ApiCallbackSite) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *CUpti_ApiCallbackSite) Scan(value interface{}) error {
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

	val, err := CUpti_ApiCallbackSiteString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
