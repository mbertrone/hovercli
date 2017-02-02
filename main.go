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
package main

import (
	"flag"
	"os"

	"github.com/mbertrone/hovercli/cli"
	"github.com/mbertrone/hovercli/config"
	"github.com/mbertrone/hovercli/hoverctl"
	l "github.com/op/go-logging"

	"github.com/mbertrone/hovercli/common"
)

var log = l.MustGetLogger("hovercli")

func init() {

	flag.BoolVar(&config.Debug, "debug", false, "enable DEBUG level in logger")
	flag.BoolVar(&config.Info, "info", true, "enable INFO  level in logger")
	flag.StringVar(&config.Hover, "hover", config.Hover, "hover url")
}

func main() {

	flag.Parse()
	common.LogInit()
	config.PrintConfig()

	var dataplane *hoverctl.Dataplane
	dataplane = hoverctl.NewDataplane()
	if err := dataplane.Init(config.Hover); err != nil {
		log.Errorf("unable to conect to Hover %s\n%s\n", config.Hover, err)
		os.Exit(1)
	}

	cli.Cli(dataplane)
}
