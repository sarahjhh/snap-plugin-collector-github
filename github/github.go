/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2015 Intel Corporation

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

package github

import (
	"strings"

	"time"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
)

// Collector struct type
type Collector struct {
	client Client
}

// NewCollector creates a instance of Github collector.
func NewCollector() Collector {
	return Collector{client: NewClient(time.Second * 5)}
}

/*  CollectMetrics collects metrics for testing.

CollectMetrics() will be called by Snap when a task that collects one of the metrics returned from this plugins
GetMetricTypes() is started. The input will include a slice of all the metric types being collected.

The output is the collected metrics as plugin.Metric and an error.
*/
func (c Collector) CollectMetrics(mts []plugin.Metric) ([]plugin.Metric, error) {
	metrics := []plugin.Metric{}

	user, err := mts[0].Config.GetString("user")
	if err != nil {
		return nil, err
	}

	mp, err := c.client.getMetricData(user)
	for _, mt := range mts {
		ns := mt.Namespace.Strings()
		ns[3] = user
		nss := strings.Join(ns, "/")
		metrics = append(metrics, mp[nss])
	}

	return metrics, nil
}

/*
	GetMetricTypes returns metric types for testing.
	GetMetricTypes() will be called when your plugin is loaded in order to populate the metric catalog(where snaps stores all
	available metrics).

	Config info is passed in. This config information would come from global config snap settings.

	The metrics returned will be advertised to users who list all the metrics and will become targetable by tasks.
*/
func (c Collector) GetMetricTypes(cfg plugin.Config) ([]plugin.Metric, error) {
	user, err := cfg.GetString("user")
	if err != nil {
		return nil, err
	}

	return c.client.getMetricTypes(user)
}

/*
	GetConfigPolicy() returns the configPolicy for your plugin.

	A config policy is how users can provide configuration info to
	plugin. Here you define what sorts of config info your plugin
	needs and/or requires.
*/
func (c Collector) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()
	policy.AddNewStringRule([]string{""}, "user", false, plugin.SetDefaultString("sarahjhh"))

	return *policy, nil
}
