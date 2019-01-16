/*
Copyright 2018 BlackRock, Inc.

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

package sensor

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestValidateSensor(t *testing.T) {
	convey.Convey("Given a sensor", t, func() {
		sensor, err := getSensor()
		convey.So(err, convey.ShouldBeNil)
		convey.Convey("Validate", func() {
			err := ValidateSensor(sensor)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}
