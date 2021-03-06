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
	case "kubernetes:storage.k8s.io/v1:CSIDriver":
		r, err = NewCSIDriver(ctx, name, nil, pulumi.URN_(urn))
	case "kubernetes:storage.k8s.io/v1:CSIDriverList":
		r, err = NewCSIDriverList(ctx, name, nil, pulumi.URN_(urn))
	case "kubernetes:storage.k8s.io/v1:CSINode":
		r, err = NewCSINode(ctx, name, nil, pulumi.URN_(urn))
	case "kubernetes:storage.k8s.io/v1:CSINodeList":
		r, err = NewCSINodeList(ctx, name, nil, pulumi.URN_(urn))
	case "kubernetes:storage.k8s.io/v1:StorageClass":
		r, err = NewStorageClass(ctx, name, nil, pulumi.URN_(urn))
	case "kubernetes:storage.k8s.io/v1:StorageClassList":
		r, err = NewStorageClassList(ctx, name, nil, pulumi.URN_(urn))
	case "kubernetes:storage.k8s.io/v1:VolumeAttachment":
		r, err = NewVolumeAttachment(ctx, name, nil, pulumi.URN_(urn))
	case "kubernetes:storage.k8s.io/v1:VolumeAttachmentList":
		r, err = NewVolumeAttachmentList(ctx, name, nil, pulumi.URN_(urn))
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
		"storage.k8s.io/v1",
		&module{version},
	)
}
