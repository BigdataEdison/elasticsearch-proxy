/*
Copyright 2016 Medcl (m AT medcl.net)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	"github.com/infinitbyte/framework/core/index"
	"github.com/infinitbyte/framework/core/pipeline"
)

type UpstreamConfig struct {
	Name          string                    `config:"name"`
	QueueName     string                    `config:"queue_name"`
	Enabled       bool                      `config:"enabled"`
	Active        bool                      `config:"active"`
	Timeout       string                    `config:"timeout"`
	Elasticsearch index.ElasticsearchConfig `config:"elasticsearch"`
}

func (v *UpstreamConfig) SafeGetQueueName() string {
	queueName := v.QueueName
	if queueName == "" {
		queueName = v.Name
	}
	return queueName
}

type ProxyConfig struct {
	UIEnabled       bool
	Upstream        []UpstreamConfig `config:"upstream"`
	Algorithm       string
	BasicAuthConfig BasicAuthConfig `config:"basic_auth"`
}

type BasicAuthConfig struct {
	User    User `config:"user"`
	Enabled bool `config:"enabled"`
}

type User struct {
	Username string `config:"username"`
	Password string `config:"password"`
}

const Url pipeline.ParaKey = "url"
const Method pipeline.ParaKey = "method"
const Body pipeline.ParaKey = "body"
const Upstream pipeline.ParaKey = "upstream"
const Response pipeline.ParaKey = "response"
const ResponseSize pipeline.ParaKey = "response_size"
const ResponseStatusCode pipeline.ParaKey = "response_code"

var upstreams map[string]UpstreamConfig = map[string]UpstreamConfig{}

func GetUpstreamConfig(key string) UpstreamConfig {
	v := upstreams[key]
	return v
}
func GetUpstreamConfigs() map[string]UpstreamConfig {
	return upstreams
}

func SetUpstream(ups []UpstreamConfig) {
	for _, v := range ups {
		//default Active is true
		v.Active = true

		//TODO get upstream status from DB, override active field
		upstreams[v.Name] = v
	}
}
