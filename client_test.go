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
	"fmt"
	"testing"
	"time"
)

func TestNewHClient(t *testing.T) {
	NewHClient()

	timeout := 1000 * time.Millisecond
	client := NewHClient(WithHTTPTimeout(timeout))

	// Use the clients GET method to create and execute the request
	res, err := client.Get("https://github.com", nil)
	if err != nil {
		panic(err)
	}

	// Heimdall returns the standard *http.Response object
	// body, err := ioutil.ReadAll(res.Body)
	// fmt.Println(string(body))
	fmt.Println(res.Status)
}

func TestNewHClientWithRetry(t *testing.T) {
	// First set a backoff mechanism. Constant backoff increases the backoff at a constant rate
	backoffInterval := 2 * time.Millisecond
	// Define a maximum jitter interval. It must be more than 1*time.Millisecond
	maximumJitterInterval := 5 * time.Millisecond

	backoff := NewConstantBackoff(backoffInterval, maximumJitterInterval)

	// Create a new retry mechanism with the backoff
	retrier := NewRetrier(backoff)

	timeout := 10000 * time.Millisecond
	// Create a new client, sets the retry mechanism, and the number of times you would like to retry
	client := NewHClient(
		WithHTTPTimeout(timeout),
		WithRetrier(retrier),
		WithRetryCount(4),
	)

	_, err := client.Get("https://github.com", nil)
	if err != nil {
		panic(err)
	}
}
