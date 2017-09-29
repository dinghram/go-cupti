// Code generated by "enumer -type=CUpti_ActivityPCSamplingPeriod -json"; DO NOT EDIT

package types

import (
	"encoding/json"
	"fmt"
)

const (
	_CUpti_ActivityPCSamplingPeriod_name_0 = "CUPTI_ACTIVITY_PC_SAMPLING_PERIOD_INVALIDCUPTI_ACTIVITY_PC_SAMPLING_PERIOD_MINCUPTI_ACTIVITY_PC_SAMPLING_PERIOD_LOWCUPTI_ACTIVITY_PC_SAMPLING_PERIOD_MIDCUPTI_ACTIVITY_PC_SAMPLING_PERIOD_HIGHCUPTI_ACTIVITY_PC_SAMPLING_PERIOD_MAX"
	_CUpti_ActivityPCSamplingPeriod_name_1 = "CUPTI_ACTIVITY_PC_SAMPLING_PERIOD_FORCE_INT"
)

var (
	_CUpti_ActivityPCSamplingPeriod_index_0 = [...]uint8{0, 41, 78, 115, 152, 190, 227}
	_CUpti_ActivityPCSamplingPeriod_index_1 = [...]uint8{0, 43}
)

func (i CUpti_ActivityPCSamplingPeriod) String() string {
	switch {
	case 0 <= i && i <= 5:
		return _CUpti_ActivityPCSamplingPeriod_name_0[_CUpti_ActivityPCSamplingPeriod_index_0[i]:_CUpti_ActivityPCSamplingPeriod_index_0[i+1]]
	case i == 2147483647:
		return _CUpti_ActivityPCSamplingPeriod_name_1
	default:
		return fmt.Sprintf("CUpti_ActivityPCSamplingPeriod(%d)", i)
	}
}

var _CUpti_ActivityPCSamplingPeriodNameToValue_map = map[string]CUpti_ActivityPCSamplingPeriod{
	_CUpti_ActivityPCSamplingPeriod_name_0[0:41]:    0,
	_CUpti_ActivityPCSamplingPeriod_name_0[41:78]:   1,
	_CUpti_ActivityPCSamplingPeriod_name_0[78:115]:  2,
	_CUpti_ActivityPCSamplingPeriod_name_0[115:152]: 3,
	_CUpti_ActivityPCSamplingPeriod_name_0[152:190]: 4,
	_CUpti_ActivityPCSamplingPeriod_name_0[190:227]: 5,
	_CUpti_ActivityPCSamplingPeriod_name_1[0:43]:    2147483647,
}

func CUpti_ActivityPCSamplingPeriodString(s string) (CUpti_ActivityPCSamplingPeriod, error) {
	if val, ok := _CUpti_ActivityPCSamplingPeriodNameToValue_map[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to CUpti_ActivityPCSamplingPeriod values", s)
}

func (i CUpti_ActivityPCSamplingPeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

func (i *CUpti_ActivityPCSamplingPeriod) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("CUpti_ActivityPCSamplingPeriod should be a string, got %s", data)
	}

	var err error
	*i, err = CUpti_ActivityPCSamplingPeriodString(s)
	return err
}