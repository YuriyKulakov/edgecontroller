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

package gorilla

import (
	"context"
	"fmt"
	"net/http"

	cce "github.com/smartedgemec/controller-ce"
)

func checkDBDeleteApps(
	ctx context.Context,
	ps cce.PersistenceService,
	id string,
) (statusCode int, err error) {
	var es []cce.Entity

	if es, err = ps.Filter(
		ctx,
		&cce.DNSAppAlias{},
		[]cce.Filter{
			{
				Field: "app_id",
				Value: id,
			},
		},
	); err != nil {
		return http.StatusInternalServerError, err
	}

	if len(es) != 0 {
		return http.StatusUnprocessableEntity, fmt.Errorf(
			"cannot delete app_id %s: record in use in "+
				"dns_app_aliases",
			id)
	}

	return 0, nil
}

func checkDBDeleteVNFs(
	ctx context.Context,
	ps cce.PersistenceService,
	id string,
) (statusCode int, err error) {
	var es []cce.Entity

	if es, err = ps.Filter(
		ctx,
		&cce.DNSVNFAlias{},
		[]cce.Filter{
			{
				Field: "vnf_id",
				Value: id,
			},
		},
	); err != nil {
		return http.StatusInternalServerError, err
	}

	if len(es) != 0 {
		return http.StatusUnprocessableEntity, fmt.Errorf(
			"cannot delete vnf_id %s: record in use in "+
				"dns_vnf_aliases",
			id)
	}

	return 0, nil
}

func checkDBDeleteDNSConfigs(
	ctx context.Context,
	ps cce.PersistenceService,
	id string,
) (statusCode int, err error) {
	var es []cce.Entity

	if es, err = ps.Filter(
		ctx,
		&cce.DNSConfigDNSAppAlias{},
		[]cce.Filter{
			{
				Field: "dns_config_id",
				Value: id,
			},
		},
	); err != nil {
		return http.StatusInternalServerError, err
	}

	if len(es) != 0 {
		return http.StatusUnprocessableEntity, fmt.Errorf(
			"cannot delete dns_config_id %s: record in use in "+
				"dns_configs_dns_app_aliases",
			id)
	}

	if es, err = ps.Filter(
		ctx,
		&cce.DNSConfigDNSVNFAlias{},
		[]cce.Filter{
			{
				Field: "dns_config_id",
				Value: id,
			},
		},
	); err != nil {
		return http.StatusInternalServerError, err
	}

	if len(es) != 0 {
		return http.StatusUnprocessableEntity, fmt.Errorf(
			"cannot delete dns_config_id %s: record in use in "+
				"dns_configs_dns_vnf_aliases",
			id)
	}

	return 0, nil
}

func checkDBDeleteDNSAppAliases(
	ctx context.Context,
	ps cce.PersistenceService,
	id string,
) (statusCode int, err error) {
	var es []cce.Entity

	if es, err = ps.Filter(
		ctx,
		&cce.DNSConfigDNSAppAlias{},
		[]cce.Filter{
			{
				Field: "dns_app_alias_id",
				Value: id,
			},
		},
	); err != nil {
		return http.StatusInternalServerError, err
	}

	if len(es) != 0 {
		return http.StatusUnprocessableEntity, fmt.Errorf(
			"cannot delete dns_app_alias_id %s: record in use in "+
				"dns_configs_dns_app_aliases",
			id)
	}

	return 0, nil
}

func checkDBDeleteDNSVNFAliases(
	ctx context.Context,
	ps cce.PersistenceService,
	id string,
) (statusCode int, err error) {
	var es []cce.Entity

	if es, err = ps.Filter(
		ctx,
		&cce.DNSConfigDNSVNFAlias{},
		[]cce.Filter{
			{
				Field: "dns_vnf_alias_id",
				Value: id,
			},
		},
	); err != nil {
		return http.StatusInternalServerError, err
	}

	if len(es) != 0 {
		return http.StatusUnprocessableEntity, fmt.Errorf(
			"cannot delete dns_vnf_alias_id %s: record in use in "+
				"dns_configs_dns_vnf_aliases",
			id)
	}

	return 0, nil
}