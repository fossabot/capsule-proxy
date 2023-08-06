// Copyright 2020-2021 Clastix Labs
// SPDX-License-Identifier: Apache-2.0

package namespaced

import (
	"fmt"

	"github.com/clastix/capsule-proxy/internal/modules"
	"github.com/clastix/capsule-proxy/internal/request"
	"github.com/clastix/capsule-proxy/internal/tenant"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/selection"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type list struct {
	client client.Client
	log    logr.Logger
	gk     schema.GroupVersionKind
}

func List(resource schema.GroupVersionKind, client client.Client) modules.Module {

	return &list{
		client: client,
		log:    ctrl.Log.WithName("namespaced_list"),
		gk:     resource,
	}
}

func (l list) Path() string {
	var path = ""
	if l.gk.Group == "" {
		path = fmt.Sprintf("/api/%s/{endpoint:%s/?}", l.gk.Version, l.gk.Kind)
	} else {
		path = fmt.Sprintf("/apis/%s/%s/{endpoint:%s/?}", l.gk.Group, l.gk.Version, l.gk.Kind)
	}
	return path
}

func (l list) Methods() []string {
	return []string{}
}

func (l list) Handle(proxyTenants []*tenant.ProxyTenant, proxyRequest request.Request) (selector labels.Selector, err error) {
	var sourceTenants []string

	for _, tnt := range proxyTenants {
		sourceTenants = append(sourceTenants, tnt.Tenant.Name)
	}

	var r *labels.Requirement

	switch {
	case len(sourceTenants) > 0:
		r, err = labels.NewRequirement("name", selection.In, sourceTenants)
	default:
		r, err = labels.NewRequirement("dontexistsignoreme", selection.Exists, []string{})
	}

	return labels.NewSelector().Add(*r), nil
}
