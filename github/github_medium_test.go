// +build medium

/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2017 Intel Corporation

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
	"testing"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGithubCollector(t *testing.T) {
	gc := NewCollector()

	Convey("Test GitHubCollector", t, func() {
		Convey("Collect login", func() {
			metrics := []plugin.Metric{
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "github", "user", "sarahjhh", "login"),
					Config:    plugin.Config{"user": "sarahjhh"},
				},
			}
			mts, err := gc.CollectMetrics(metrics)
			So(mts, ShouldNotBeEmpty)
			So(err, ShouldBeNil)
			So(mts[0].Data, ShouldEqual, "sarahjhh")
		})

		Convey("Collect name", func() {
			metrics := []plugin.Metric{
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "github", "user", "sarahjhh", "name"),
					Config:    plugin.Config{"user": "sarahjhh"},
				},
			}
			mts, err := gc.CollectMetrics(metrics)
			So(mts, ShouldNotBeEmpty)
			So(err, ShouldBeNil)
			So(mts[0].Data, ShouldEqual, "Sarah")
		})
	})
}
