// Copyright 2017 Politecnico di Torino
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package config

import (
	"fmt"

	l "github.com/op/go-logging"
)

var Hover = "http://localhost:5002"

var Debug = false
var Info = true

var log = l.MustGetLogger("hovercli")

func PrintConfig() {
	fmt.Printf("-----Hover Command Line Interface----------------------------\n")
	fmt.Printf("connecting to hover....")
	fmt.Printf("%8s:%s\n", "Hover", Hover)
	fmt.Printf("--------------------------------------------------------------\n\n")
}
