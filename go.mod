module github.com/pulumi/pulumi-docker

go 1.12

require (
	github.com/Azure/go-autorest/autorest/azure/auth v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/to v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.2.0 // indirect
	github.com/containerd/containerd v1.2.8 // indirect
	github.com/hashicorp/terraform v0.12.6
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v1.0.0
	github.com/pulumi/pulumi-terraform v0.18.4-0.20190828172748-3f206601e7a1
	github.com/stretchr/testify v1.3.1-0.20190311161405-34c6fa2dc709
	github.com/terraform-providers/terraform-provider-docker v0.0.0-20190822140719-8a5b696b491c
)

replace (
	git.apache.org/thrift.git => github.com/apache/thrift v0.12.0
	github.com/Azure/azure-sdk-for-go => github.com/Azure/azure-sdk-for-go v31.1.0+incompatible
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.4.3-0.20190807103436-de736cf91b92
	github.com/docker/docker => github.com/docker/docker v0.7.3-0.20190827134902-c33872e3f4dc
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
	github.com/zclconf/go-cty => github.com/zclconf/go-cty v1.1.0
)
