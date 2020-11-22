package cupti

import (
	"testing"

	"github.com/c3sr/go-cupti/types"
)

func Test1(t *testing.T) {
	getActivityObjectKindId(types.CUPTI_ACTIVITY_OBJECT_PROCESS, nil)
}
