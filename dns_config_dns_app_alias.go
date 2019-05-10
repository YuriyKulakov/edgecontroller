// Copyright 2019 Smart-Edge.com, Inc. All rights reserved.
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
	"strings"

	"github.com/smartedgemec/controller-ce/uuid"
)

// DNSConfigDNSAppAlias represents an association between a DNSConfig
// and a DNSAppAlias.
type DNSConfigDNSAppAlias struct {
	ID            string `json:"id"`
	DNSConfigID   string `json:"dns_config_id"`
	DNSAppAliasID string `json:"dns_app_alias_id"`
}

// GetTableName returns the name of the persistence table.
func (*DNSConfigDNSAppAlias) GetTableName() string {
	return "dns_configs_dns_app_aliases"
}

// GetID gets the ID.
func (cfg_alias *DNSConfigDNSAppAlias) GetID() string {
	return cfg_alias.ID
}

// SetID sets the ID.
func (cfg_alias *DNSConfigDNSAppAlias) SetID(id string) {
	cfg_alias.ID = id
}

// Validate validates the model.
func (cfg_alias *DNSConfigDNSAppAlias) Validate() error {
	if !uuid.IsValid(cfg_alias.ID) {
		return errors.New("id not a valid uuid")
	}
	if !uuid.IsValid(cfg_alias.DNSConfigID) {
		return errors.New("dns_config_id not a valid uuid")
	}
	if !uuid.IsValid(cfg_alias.DNSAppAliasID) {
		return errors.New("dns_app_alias_id not a valid uuid")
	}

	return nil
}

func (cfg_alias *DNSConfigDNSAppAlias) String() string {
	return fmt.Sprintf(strings.TrimSpace(`
DNSConfigDNSAppAlias[
    ID: %s
    DNSConfigID: %s
    DNSAppAliasID: %s
]`),
		cfg_alias.ID,
		cfg_alias.DNSConfigID,
		cfg_alias.DNSAppAliasID)
}