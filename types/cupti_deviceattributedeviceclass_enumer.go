// Code generated by "enumer -type=CUpti_DeviceAttributeDeviceClass -json -text -yaml -sql"; DO NOT EDIT.

//
package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

const _CUpti_DeviceAttributeDeviceClassName = "CUPTI_DEVICE_ATTR_DEVICE_CLASS_TESLACUPTI_DEVICE_ATTR_DEVICE_CLASS_QUADROCUPTI_DEVICE_ATTR_DEVICE_CLASS_GEFORCECUPTI_DEVICE_ATTR_DEVICE_CLASS_TEGRA"

var _CUpti_DeviceAttributeDeviceClassIndex = [...]uint8{0, 36, 73, 111, 147}

func (i CUpti_DeviceAttributeDeviceClass) String() string {
	if i < 0 || i >= CUpti_DeviceAttributeDeviceClass(len(_CUpti_DeviceAttributeDeviceClassIndex)-1) {
		return fmt.Sprintf("CUpti_DeviceAttributeDeviceClass(%d)", i)
	}
	return _CUpti_DeviceAttributeDeviceClassName[_CUpti_DeviceAttributeDeviceClassIndex[i]:_CUpti_DeviceAttributeDeviceClassIndex[i+1]]
}

var _CUpti_DeviceAttributeDeviceClassValues = []CUpti_DeviceAttributeDeviceClass{0, 1, 2, 3}

var _CUpti_DeviceAttributeDeviceClassNameToValueMap = map[string]CUpti_DeviceAttributeDeviceClass{
	_CUpti_DeviceAttributeDeviceClassName[0:36]:    0,
	_CUpti_DeviceAttributeDeviceClassName[36:73]:   1,
	_CUpti_DeviceAttributeDeviceClassName[73:111]:  2,
	_CUpti_DeviceAttributeDeviceClassName[111:147]: 3,
}

// CUpti_DeviceAttributeDeviceClassString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func CUpti_DeviceAttributeDeviceClassString(s string) (CUpti_DeviceAttributeDeviceClass, error) {
	if val, ok := _CUpti_DeviceAttributeDeviceClassNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to CUpti_DeviceAttributeDeviceClass values", s)
}

// CUpti_DeviceAttributeDeviceClassValues returns all values of the enum
func CUpti_DeviceAttributeDeviceClassValues() []CUpti_DeviceAttributeDeviceClass {
	return _CUpti_DeviceAttributeDeviceClassValues
}

// IsACUpti_DeviceAttributeDeviceClass returns "true" if the value is listed in the enum definition. "false" otherwise
func (i CUpti_DeviceAttributeDeviceClass) IsACUpti_DeviceAttributeDeviceClass() bool {
	for _, v := range _CUpti_DeviceAttributeDeviceClassValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for CUpti_DeviceAttributeDeviceClass
func (i CUpti_DeviceAttributeDeviceClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for CUpti_DeviceAttributeDeviceClass
func (i *CUpti_DeviceAttributeDeviceClass) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("CUpti_DeviceAttributeDeviceClass should be a string, got %s", data)
	}

	var err error
	*i, err = CUpti_DeviceAttributeDeviceClassString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for CUpti_DeviceAttributeDeviceClass
func (i CUpti_DeviceAttributeDeviceClass) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for CUpti_DeviceAttributeDeviceClass
func (i *CUpti_DeviceAttributeDeviceClass) UnmarshalText(text []byte) error {
	var err error
	*i, err = CUpti_DeviceAttributeDeviceClassString(string(text))
	return err
}

// MarshalYAML implements a YAML Marshaler for CUpti_DeviceAttributeDeviceClass
func (i CUpti_DeviceAttributeDeviceClass) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for CUpti_DeviceAttributeDeviceClass
func (i *CUpti_DeviceAttributeDeviceClass) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = CUpti_DeviceAttributeDeviceClassString(s)
	return err
}

func (i CUpti_DeviceAttributeDeviceClass) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *CUpti_DeviceAttributeDeviceClass) Scan(value interface{}) error {
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

	val, err := CUpti_DeviceAttributeDeviceClassString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}