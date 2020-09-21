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

package http

import (
	"testing"
	"time"

	"github.com/issue9/assert"
)

func TestRetrierWithExponentialBackoff(t *testing.T) {

	exponentialBackoff := NewExponentialBackoff(2*time.Millisecond, 10*time.Millisecond, 2.0, 1*time.Millisecond)
	exponentialRetrier := NewRetrier(exponentialBackoff)

	assert.True(t, 4*time.Millisecond <= exponentialRetrier.NextInterval(1))
}

func TestRetrierWithConstantBackoff(t *testing.T) {
	backoffInterval := 2 * time.Millisecond
	maximumJitterInterval := 1 * time.Millisecond

	constantBackoff := NewConstantBackoff(backoffInterval, maximumJitterInterval)
	constantRetrier := NewRetrier(constantBackoff)

	assert.True(t, 2*time.Millisecond <= constantRetrier.NextInterval(1))
}

func TestNoRetrier(t *testing.T) {
	noRetrier := NewNoRetrier()
	nextInterval := noRetrier.NextInterval(1)
	assert.Equal(t, time.Duration(0), nextInterval)
}
