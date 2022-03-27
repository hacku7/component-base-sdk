package retryutil

import (
	"fmt"
)

var RetryAbleErr = fmt.Errorf("retry")
var TimeoutErr = fmt.Errorf("timeout")
