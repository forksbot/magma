/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package registry

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	platform_registry "magma/orc8r/lib/go/registry"
)

// CloudRegistry is aa adaptor between platform registry and gateway registry
// it's used for legacy FeG gateway support and may be replaced by gateway/service_registry in the future
type CloudRegistry interface {
	// platform's ServiceRegistry methods
	AddServices(locations ...platform_registry.ServiceLocation)
	AddService(location platform_registry.ServiceLocation)
	GetServiceAddress(service string) (string, error)
	GetServicePort(service string) (int, error)
	GetServiceProxyAliases(service string) (map[string]int, error)
	ListAllServices() []string

	// Gateway specific methods
	//
	// GetConnection provides a gRPC connection to a service in the registry.
	GetConnection(service string) (*grpc.ClientConn, error)
	GetConnectionImpl(ctx context.Context, service string, opts ...grpc.DialOption) (*grpc.ClientConn, error)

	// GetCloudConnection Creates and returns a new GRPC service connection to the service.
	// GetCloudConnection always creates a new connection & it's responsibility of the caller to close it.
	GetCloudConnection(service string) (*grpc.ClientConn, error)
}

// NewCloudRegistry returns a new instance of gateway's cloud registry
func NewCloudRegistry() *platform_registry.ServiceRegistry {
	return platform_registry.Get()
}
