// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "kubernetes:rbac.authorization.k8s.io/v1:ClusterRole":
		r, err = NewClusterRole(ctx, name, nil, pulumi.URN_(urn))
	case "kubernetes:rbac.authorization.k8s.io/v1:ClusterRoleBinding":
		r, err = NewClusterRoleBinding(ctx, name, nil, pulumi.URN_(urn))
	case "kubernetes:rbac.authorization.k8s.io/v1:ClusterRoleBindingList":
		r, err = NewClusterRoleBindingList(ctx, name, nil, pulumi.URN_(urn))
	case "kubernetes:rbac.authorization.k8s.io/v1:ClusterRoleList":
		r, err = NewClusterRoleList(ctx, name, nil, pulumi.URN_(urn))
	case "kubernetes:rbac.authorization.k8s.io/v1:Role":
		r, err = NewRole(ctx, name, nil, pulumi.URN_(urn))
	case "kubernetes:rbac.authorization.k8s.io/v1:RoleBinding":
		r, err = NewRoleBinding(ctx, name, nil, pulumi.URN_(urn))
	case "kubernetes:rbac.authorization.k8s.io/v1:RoleBindingList":
		r, err = NewRoleBindingList(ctx, name, nil, pulumi.URN_(urn))
	case "kubernetes:rbac.authorization.k8s.io/v1:RoleList":
		r, err = NewRoleList(ctx, name, nil, pulumi.URN_(urn))
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	return
}

func init() {
	version, err := kubernetes.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"kubernetes",
		"rbac.authorization.k8s.io/v1",
		&module{version},
	)
}
