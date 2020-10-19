// Copyright 2020 colynn.liu
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package httpclient

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConstantBackoffNextTime(t *testing.T) {
	constantBackoff := NewConstantBackoff(100*time.Millisecond, 0*time.Millisecond)
	assert.Equal(t, 100*time.Millisecond, constantBackoff.Next(0))
	assert.Equal(t, 100*time.Millisecond, constantBackoff.Next(1))
	assert.Equal(t, 100*time.Millisecond, constantBackoff.Next(2))
	assert.Equal(t, 100*time.Millisecond, constantBackoff.Next(3))
}
