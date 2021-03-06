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
	case "kubernetes:authorization.k8s.io/v1:LocalSubjectAccessReview":
		r, err = NewLocalSubjectAccessReview(ctx, name, nil, pulumi.URN_(urn))
	case "kubernetes:authorization.k8s.io/v1:SelfSubjectAccessReview":
		r, err = NewSelfSubjectAccessReview(ctx, name, nil, pulumi.URN_(urn))
	case "kubernetes:authorization.k8s.io/v1:SelfSubjectRulesReview":
		r, err = NewSelfSubjectRulesReview(ctx, name, nil, pulumi.URN_(urn))
	case "kubernetes:authorization.k8s.io/v1:SubjectAccessReview":
		r, err = NewSubjectAccessReview(ctx, name, nil, pulumi.URN_(urn))
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
		"authorization.k8s.io/v1",
		&module{version},
	)
}
