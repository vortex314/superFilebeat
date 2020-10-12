// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package actions

import (
	"fmt"
	"strings"
	"time"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/elastic/beats/v7/libbeat/processors"
	"github.com/elastic/beats/v7/libbeat/processors/checks"
	"github.com/vjeantet/grok"
)

type grokPatterns struct {
	Patterns   []string
	Timestamps []string
	Grok       *grok.Grok
	logger     *logp.Logger
}

func init() {
	processors.RegisterPlugin("grok",
		checks.ConfigChecked(newGrokPatterns,
			checks.RequireFields("patterns"),
			checks.AllowedFields("patterns", "timestamps", "when")))

}

func newGrokPatterns(c *common.Config) (processors.Processor, error) {
	config := struct {
		Patterns   []string `config:"patterns"`
		Timestamps []string `config:"timestamps"`
	}{}
	err := c.Unpack(&config)
	if err != nil {
		return nil, fmt.Errorf("fail to unpack the grok_Patterns configuration: %s", err)
	}

	grok, _ := grok.NewWithConfig(&grok.Config{NamedCapturesOnly: true})

	f := &grokPatterns{Patterns: config.Patterns,
		Timestamps: config.Timestamps,
		Grok:       grok,
		logger:     logp.NewLogger("grok-processor")}
	return f, nil
}

func (f *grokPatterns) Run(event *beat.Event) (*beat.Event, error) {

	var errors []string

	failed := true
	msg, _ := event.Fields.GetValue("message")
	message := msg.(string)
//    fmt.Println(" ===> text : " ,message)
//    fmt.Println(" ===> patterns : ",f.Patterns)


	for _, pattern := range f.Patterns {
//        fmt.Println(" pattern ===", pattern)

		values, erc := f.Grok.Parse(pattern, message)
 //       fmt.Println(" ===> values :",values)
 //       fmt.Println(" ===> erc :" ,erc)
		if len(values)!=0 || erc!=nil { // testing on erc doesn't work  , always nil
			failed = false
			for k, v := range values {
				if k == "timestamp" {
					for _, timestamp := range f.Timestamps {
//                        fmt.Println(" timestamp pattern  ===", timestamp)
						t, e := time.Parse(timestamp, v)
						if e == nil {
							//							fmt.Println(" timestamp ===", t.Unix())
							//							event.PutValue("timestamp", t.UnixNano()/1000000)
							event.PutValue("timestamp", t)
							break
						}
					}
				} else {
					event.PutValue(k, v)
				}
			}
			break
		}
	}

	if failed {
		f.logger.Warn(" failed to find matching pattern")
		errors = append(errors, " failed to find matching pattern ")
		event.PutValue("@errors", strings.Join(errors, ", "))
	}

	if len(errors) > 0 {
		return event, fmt.Errorf(strings.Join(errors, ", "))
	}
	return event, nil
}

func (f *grokPatterns) String() string {
	return "grok_Patterns=" + strings.Join(f.Patterns, ", ")
}
