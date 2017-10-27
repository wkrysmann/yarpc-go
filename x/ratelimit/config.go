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

import (
	"fmt"

	"go.uber.org/config"
)

const (
	// ConfigurationKey is the portion of the service configuration that
	// the ratelimiting middleware reads
	ConfigurationKey = "ratelimit"
)

type ratelmiterConfig []InboundThrottleConfig

// InboundThrottleConfig describes how to configure a ratelimiter
type InboundThrottleConfig struct {
	// RPS is the maximum requests per second, after which the inbound will
	// throttle inbound requests with a ResourceExhaustedError of "rate limit
	// exceeded".
	RPS int `config:"rps"`
	// BurstLimit determines how much slack the rate limiter will tolerate for
	// a burst of requests from an idle state before throttling.
	// The default is 10. A burstLimit of 0 implies the default.
	// Use "noSlack" to configure a rate limiter without slack.
	BurstLimit int `config:"burstLimit"`
	// NoSlack configures the rate limiter without any slack, even after idling
	// indefinitely.
	NoSlack bool `config:"noSlack"`

	// Procedure name to throttle
	Procedure string `config:"procedure"`

	// TODO(apeatsbond): change config format so the below can be explicitly set elsewhere
	// GlobalThrottle indicates if this limiter configuration is for the global rate
	GlobalThrottle bool `config:"global"`
	// DefaultThrottle indicates if this limiter is the default for unspecified
	//  procedures
	DefaultThrottle bool `config:"default"`
}

// Params defines dependencies for a new Unary Inbound Middleware
type Params struct{ Config config.Provider }

// Build creates a unary inbound rate limit middleware, or returns an error if
// the configuration is invalid.
func (c InboundThrottleConfig) Build() (*Throttle, error) {
	var opts []Option
	if c.NoSlack && c.BurstLimit > 0 {
		return nil, fmt.Errorf("unary inbound rate limit middleware configured with contradictory noSlack and non-zero BurstLimit (%d)", c.BurstLimit)
	}
	if c.NoSlack {
		opts = append(opts, WithoutSlack)
	}
	if c.BurstLimit > 0 {
		opts = append(opts, WithBurstLimit(c.BurstLimit))
	}

	// TODO(apeatsbond): create throttle wrapper so we can keep track of what
	// each throttle should apply to (ie service/caller/procedure name)
	return NewThrottle(c.RPS, opts...)
}

func New(p Params) (*UnaryInboundMiddleware, error) {
	// TODO(apeatsbond): report name/version, similar to retryfx

	var cfg ratelmiterConfig
	if err := p.Config.Get(ConfigurationKey).Populate(&cfg); err != nil {
		return nil, err
	}

	return CreateMiddlewareFromConfig(cfg)
}

func CreateMiddlewareFromConfig(cfg ratelmiterConfig) (*UnaryInboundMiddleware, error) {
	var (
		globalThrottle  *Throttle
		defaultThrottle *Throttle
	)
	throttles := make([]*Throttle, 0, len(cfg))

	// iterate over limiter configurations and create them
	for _, tCfg := range cfg {
		if tCfg.DefaultThrottle && defaultThrottle != nil {
			// TODO(apeatsbond): maybe just log/ignore these cases and take the first one as valid
			return nil, fmt.Errorf("more than one default limiter cannot be initialized")
		}
		if tCfg.GlobalThrottle && globalThrottle != nil {
			return nil, fmt.Errorf("more than one global limiter cannot be initialiized")
		}

		t, err := tCfg.Build()
		if err != nil {
			return nil, err
		}

		if tCfg.DefaultThrottle {
			defaultThrottle = t
		} else if tCfg.GlobalThrottle {
			globalThrottle = t
		} else {
			throttles = append(throttles, t)
		}
	}

	if globalThrottle == nil {
		// TODO(apeatsbond): ¯\_(ツ)_/¯
		// determine this cumulative rate
	}
	if defaultThrottle == nil {
		// TODO(apeatsbond): ¯\_(ツ)_/¯
	}

	return &UnaryInboundMiddleware{
		throttles:       throttles,
		defaultThrottle: defaultThrottle,
		globalThrottle:  globalThrottle,
	}, nil
}
