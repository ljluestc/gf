// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gtime_test

import (
	"github.com/gogf/gf/v2/os/gtime"
)

func init() {
	// Set the local timezone to Asia/Shanghai for all tests in this package.
	// This ensures that tests are consistent across different environments.
	_ = gtime.SetTimeZone("Asia/Shanghai")
}
