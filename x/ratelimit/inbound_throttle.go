// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package ratelimit

import "go.uber.org/yarpc/api/transport"

// TODO(apeatsbond): change to
// - request throttle?
// - specificThrottle
// - constrained throttle..
// - service throttle?
// TODO(apeatsbond): maybe have separate throttlers per thing we're throttling with??
type inboundThrottle struct {
	throttle *Throttle

	//
	service   string
	procedure string
	caller    string
}

// Throttle the request
func (t *inboundThrottle) Throttle() bool {
	return t.throttle.Throttle()
}

// AppliesToRequest determines if this throttle can be responsible for
// throttling the given request
func (t *inboundThrottle) AppliesToRequest(req *transport.Request) bool {
	// TODO(apeatsbond): choose based on service name, caller, procedure whatever
	return true
}
