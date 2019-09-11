// Copyright 2019 Intel Corporation. All rights reserved
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

package cce

import (
	"errors"
	"fmt"
	"net"
	"strings"

	"github.com/otcshare/edgecontroller/uuid"
)

// TrafficPolicyKubeOVN is an application or interface traffic policy.
type TrafficPolicyKubeOVN struct {
	ID      string         `json:"id"`
	Name    string         `json:"name"`
	Ingress []*IngressRule `json:"ingress_rules"`
	Egress  []*EgressRule  `json:"egress_rules"`
}

// GetTableName returns the name of the persistence table.
func (*TrafficPolicyKubeOVN) GetTableName() string {
	return "traffic_policies_kube_ovn"
}

// GetID gets the ID.
func (tp *TrafficPolicyKubeOVN) GetID() string {
	return tp.ID
}

// SetID sets the ID.
func (tp *TrafficPolicyKubeOVN) SetID(id string) {
	tp.ID = id
}

// Validate validates the model.
func (tp *TrafficPolicyKubeOVN) Validate() error {
	if !uuid.IsValid(tp.ID) {
		return errors.New("id not a valid UUID")
	}

	if tp.Name == "" {
		return errors.New("name cannot be empty")
	}

	if len(tp.Ingress) == 0 && len(tp.Egress) == 0 {
		return errors.New("ingress and egress cannot be empty")
	}

	for i, rule := range tp.Ingress {
		if err := rule.Validate(); err != nil {
			return fmt.Errorf("Ingress[%d].%s", i, err.Error())
		}
	}
	for i, rule := range tp.Egress {
		if err := rule.Validate(); err != nil {
			return fmt.Errorf("Egress[%d].%s", i, err.Error())
		}
	}

	return nil
}

// FilterFields returns the filterable fields for this model.
func (*TrafficPolicyKubeOVN) FilterFields() []string {
	return []string{}
}

func (tp *TrafficPolicyKubeOVN) String() string {
	var ingress, egress string

	for i, rule := range tp.Ingress {
		ingress += rule.String()
		if i < len(tp.Ingress)-1 {
			ingress += "\n		"
		}
	}

	for i, rule := range tp.Egress {
		egress += rule.String()
		if i < len(tp.Egress)-1 {
			egress += "\n		"
		}
	}

	return fmt.Sprintf(strings.TrimSpace(`
TrafficPolicyKubeOVN[
	ID: %s,
	Name: %s,
	Ingress: [
		%s
	], 
	Egress: [
		%s
	]
]`),
		tp.ID,
		tp.Name,
		ingress,
		egress)
}

// IngressRule is the model for a ingress rule.
type IngressRule struct {
	Description string     `json:"description"`
	From        []*IPBlock `json:"from"`
	Ports       []*Port    `json:"ports"`
}

// Validate validates the model.
func (ir *IngressRule) Validate() error {
	for i, block := range ir.From {
		if err := block.Validate(); err != nil {
			return fmt.Errorf("From[%d].%s", i, err.Error())
		}
	}

	for i, port := range ir.Ports {
		if err := port.Validate(); err != nil {
			return fmt.Errorf("Ports[%d].%s", i, err.Error())
		}
	}

	return nil
}

func (ir *IngressRule) String() string {
	var from, ports string
	for i, block := range ir.From {
		from += block.String()
		if i < len(ir.From)-1 {
			from += "\n				"
		}
	}
	for i, port := range ir.Ports {
		ports += port.String()
		if i < len(ir.Ports)-1 {
			ports += "\n				"
		}
	}
	return fmt.Sprintf(strings.TrimSpace(`
		IngressRule[
			Description: %s,
			From: [
				%s
			],
			Ports: [
				%s
			]
		]`),
		ir.Description,
		from,
		ports)
}

// EgressRule is the model for a egress rule.
type EgressRule struct {
	Description string     `json:"description"`
	To          []*IPBlock `json:"to"`
	Ports       []*Port    `json:"ports"`
}

// Validate validates the model.
func (er *EgressRule) Validate() error {
	for i, block := range er.To {
		if err := block.Validate(); err != nil {
			return fmt.Errorf("To[%d].%s", i, err.Error())
		}
	}

	for i, port := range er.Ports {
		if err := port.Validate(); err != nil {
			return fmt.Errorf("Ports[%d].%s", i, err.Error())
		}
	}

	return nil
}

func (er *EgressRule) String() string {
	var to, ports string
	for i, block := range er.To {
		to += block.String()
		if i < len(er.To)-1 {
			to += "\n				"
		}
	}
	for i, port := range er.Ports {
		ports += port.String()
		if i < len(er.Ports)-1 {
			ports += "\n				"
		}
	}
	return fmt.Sprintf(strings.TrimSpace(`
		EgressRule[
			Description: %s,
			To: [
				%s
			],
			Ports: [
				%s
			]
		]`),
		er.Description,
		to,
		ports)
}

// IPBlock is the model for a ip block.
type IPBlock struct {
	CIDR   string   `json:"cidr"`
	Except []string `json:"except"`
}

// Validate validates the model.
func (ipb *IPBlock) Validate() error {
	_, n, err := net.ParseCIDR(ipb.CIDR)
	if err != nil {
		return fmt.Errorf("Invalid CIDR: %s", err.Error())
	}

	for i, exceptCIDR := range ipb.Except {
		exceptIP, exceptNet, err := net.ParseCIDR(exceptCIDR)
		if err != nil {
			return fmt.Errorf("Except[%d].Invalid CIDR: %s", i, err.Error())
		}

		if exceptNet.String() == n.String() {
			return fmt.Errorf("Except[%d].CIDR(%s) is the same as CIDR(%s)", i, exceptCIDR, ipb.CIDR)
		}

		if !n.IP.Equal(exceptIP.Mask(n.Mask)) {
			return fmt.Errorf("Except[%d].CIDR(%s) is not in CIDR(%s)", i, exceptCIDR, ipb.CIDR)
		}

		eS, _ := exceptNet.Mask.Size()
		nS, _ := n.Mask.Size()
		if eS <= nS {
			return fmt.Errorf("Except[%d].CIDR(%s) mask is invalid", i, exceptCIDR)
		}
	}

	return nil
}

func (ipb *IPBlock) String() string {
	var except string
	for i, e := range ipb.Except {
		except += e
		if i < len(ipb.Except)-1 {
			except += "\n						"
		}
	}
	return fmt.Sprintf(strings.TrimSpace(`
				IPBlock[
					CIDR: %s,
					Except:	[
						%s
					]
				]`),
		ipb.CIDR,
		except)
}

// Port is the model for a ip block.
type Port struct {
	Port     uint16 `json:"port"`
	Protocol string `json:"protocol"`
}

// Validate validates the model.
func (p *Port) Validate() error {
	if p.Protocol != "tcp" && p.Protocol != "udp" && p.Protocol != "sctp" {
		return fmt.Errorf("Not supported protocol: %s", p.Protocol)
	}

	return nil
}

func (p *Port) String() string {
	return fmt.Sprintf(strings.TrimSpace(`
				Port[
					Port: %d,
					Protocol: %s
				]`),
		p.Port,
		p.Protocol)
}