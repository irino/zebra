// Copyright 2016, 2017 Zebra Project
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

package rib

import (
	"fmt"
	"net"
	"testing"

	"github.com/coreswitch/netutil"
)

func SendMessage(conn net.Conn, msg *Message) {
	s, _ := msg.Serialize()
	conn.Write(s)
}

func Dial() (net.Conn, error) {
	// Connect to server
	conn, err := net.Dial("unix", "/var/run/zserv.api")
	if err != nil {
		return nil, err
	}
	return conn, err
}

// ZAPI version 2.
func TestV2Hello(t *testing.T) {
	fmt.Println("Hello")

	conn, err := Dial()
	if err != nil {
		t.Errorf("Dial failed %v\n", err)
	}
	defer conn.Close()
	msg := Message{
		Header: Header{Command: HELLO},
	}
	SendMessage(conn, &msg)
}

func TestV2Nexthop(t *testing.T) {
	fmt.Println("Nexthop IPv4 Lookup")

	conn, err := Dial()
	if err != nil {
		t.Errorf("Dial failed %v\n", err)
	}
	defer conn.Close()

	addr := netutil.ParseIPv4("10.0.0.1")
	body := &IPv4NexthopLookupBody{Addr: addr}

	msg := Message{
		Header: Header{Command: IPV4_NEXTHOP_LOOKUP},
		Body:   body,
	}
	SendMessage(conn, &msg)
}
