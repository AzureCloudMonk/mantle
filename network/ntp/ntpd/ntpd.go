// Copyright 2015 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/coreos/mantle/network/ntp"
	"log"
	"net"
)

func main() {
	l, err := net.ListenPacket("udp", ":123")
	if err != nil {
		log.Fatalf("Listen failed: %v", err)
	}

	err = ntp.Server{l}.Serve()
	if err != nil {
		log.Fatalf("Serve failed: %v", err)
	}
}
