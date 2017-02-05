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
// Inline command line interface for debug purposes
// in future this cli will be a separate go program that connects to the main iovisor-ovn daemon
// in future this cli will use a cli go library (e.g. github.com/urfave/cli )

package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mbertrone/hovercli/hoverctl"
)

func Cli(dataplaneref *hoverctl.Dataplane) {
	//db := &mainlogic.Mon.DB
	dataplane := dataplaneref
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("cli@hovercli$ ")
		line, _ := reader.ReadString('\n')

		line = TrimSuffix(line, "\n")
		args := strings.Split(line, " ")

		if len(args) >= 1 {
			switch args[0] {
			case "interfaces", "i":
				fmt.Printf("\nInterfaces\n\n")
				_, external_interfaces := hoverctl.ExternalInterfacesListGET(dataplane)
				hoverctl.ExternalInterfacesListPrint(external_interfaces)
			case "modules", "m":
				if len(args) >= 2 {
					switch args[1] {
					case "get":
						switch len(args) {
						case 2:
							fmt.Printf("\nModules GET\n\n")
							_, modules := hoverctl.ModuleListGET(dataplane)
							hoverctl.ModuleListPrint(modules)
						case 3:
							fmt.Printf("\nModules GET\n\n")
							_, module := hoverctl.ModuleGET(dataplane, args[2])
							hoverctl.ModulePrint(module)
						default:
							PrintModulesUsage()
						}
					case "post":
						switch len(args) {
						case 3:
							fmt.Printf("\nModules POST\n\n")
							if args[2] == "switch" {
								// _, module := hoverctl.ModulePOST(dataplane, "bpf", "Switch", l2switch.SwitchSecurityPolicy)
								// hoverctl.ModulePrint(module)
							} else {
								//TODO Print modules list
							}
						default:
							PrintModulesUsage()
						}
					case "delete":
						switch len(args) {
						case 3:
							fmt.Printf("\nModules DELETE\n\n")
							hoverctl.ModuleDELETE(dataplane, args[2])
						default:
							PrintModulesUsage()
						}
					default:
						PrintModulesUsage()
					}
				} else {
					PrintModulesUsage()
				}
			case "links", "l":
				if len(args) >= 2 {
					switch args[1] {
					case "get":
						switch len(args) {
						case 2:
							fmt.Printf("\nLinks GET\n\n")
							_, links := hoverctl.LinkListGet(dataplane)
							hoverctl.LinkListPrint(links)
						case 3:
							fmt.Printf("\nLinks GET\n\n")
							_, link := hoverctl.LinkGET(dataplane, args[2])
							hoverctl.LinkPrint(link)
						default:
							PrintLinksUsage()
						}
					case "post":
						switch len(args) {
						case 4:
							fmt.Printf("\nLinks POST\n\n")
							_, link := hoverctl.LinkPOST(dataplane, args[2], args[3])
							hoverctl.LinkPrint(link)
						default:
							PrintLinksUsage()
						}
					case "delete":
						switch len(args) {
						case 3:
							fmt.Printf("\nLinks DELETE\n\n")
							hoverctl.LinkDELETE(dataplane, args[2])
						default:
							PrintLinksUsage()
						}
					default:
						PrintLinksUsage()
					}
				} else {
					PrintLinksUsage()
				}
			case "table", "t":
				if len(args) >= 2 {
					switch args[1] {
					case "get":
						switch len(args) {
						case 2:
							fmt.Printf("\nTable GET\n\n")
							_, modules := hoverctl.ModuleListGET(dataplane)
							for moduleName, _ := range modules {
								fmt.Printf("**MODULE** -> %s\n", moduleName)
								_, tables := hoverctl.TableListGET(dataplane, moduleName)
								for _, tablename := range tables {
									fmt.Printf("Table *%s*\n", tablename)
									_, table := hoverctl.TableGET(dataplane, moduleName, tablename)
									hoverctl.TablePrint(table)
								}
							}
						case 3:
							fmt.Printf("\nTable GET\n\n")
							_, tables := hoverctl.TableListGET(dataplane, args[2])
							for _, tablename := range tables {
								fmt.Printf("Table *%s*\n", tablename)
								_, table := hoverctl.TableGET(dataplane, args[2], tablename)
								hoverctl.TablePrint(table)
							}
						case 4:
							fmt.Printf("\nTable GET\n\n")
							_, table := hoverctl.TableGET(dataplane, args[2], args[3])
							hoverctl.TablePrint(table)
						case 5:
							fmt.Printf("\nTable GET\n\n")
							_, tableEntry := hoverctl.TableEntryGET(dataplane, args[2], args[3], args[4])
							hoverctl.TableEntryPrint(tableEntry)
						default:
							PrintTableUsage()
						}
					case "put":
						if len(args) == 6 {
							fmt.Printf("\nTable PUT\n\n")
							_, tableEntry := hoverctl.TableEntryPUT(dataplane, args[2], args[3], args[4], args[5])
							hoverctl.TableEntryPrint(tableEntry)
						} else {
							PrintTableUsage()
						}
					case "post":
						if len(args) == 6 {
							fmt.Printf("\nTable POST\n\n")
							_, tableEntry := hoverctl.TableEntryPOST(dataplane, args[2], args[3], args[4], args[5])
							hoverctl.TableEntryPrint(tableEntry)
						} else {
							PrintTableUsage()
						}
					case "delete":
						if len(args) == 5 {
							fmt.Printf("\nTable DELETE\n\n")
							hoverctl.TableEntryDELETE(dataplane, args[2], args[3], args[4])
						} else {
							PrintTableUsage()
						}
					default:
						PrintTableUsage()
					}
				} else {
					PrintTableUsage()
				}
			case "help", "h":
				PrintHelp()

			case "":
			default:
				fmt.Printf("\nInvalid Command\n\n")
				PrintHelp()
			}
		}
		fmt.Printf("\n")
	}
}

func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}
