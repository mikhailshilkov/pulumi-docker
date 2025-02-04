// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package docker

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages the configuration of a Docker service in a swarm.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-docker/blob/master/website/docs/r/config.html.markdown.
type Config struct {
	s *pulumi.ResourceState
}

// NewConfig registers a new resource with the given unique name, arguments, and options.
func NewConfig(ctx *pulumi.Context,
	name string, args *ConfigArgs, opts ...pulumi.ResourceOpt) (*Config, error) {
	if args == nil || args.Data == nil {
		return nil, errors.New("missing required argument 'Data'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["data"] = nil
		inputs["name"] = nil
	} else {
		inputs["data"] = args.Data
		inputs["name"] = args.Name
	}
	s, err := ctx.RegisterResource("docker:index/config:Config", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Config{s: s}, nil
}

// GetConfig gets an existing Config resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfig(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ConfigState, opts ...pulumi.ResourceOpt) (*Config, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["data"] = state.Data
		inputs["name"] = state.Name
	}
	s, err := ctx.ReadResource("docker:index/config:Config", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Config{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Config) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Config) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The base64 encoded data of the config.
func (r *Config) Data() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["data"])
}

// The name of the Docker config.
func (r *Config) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Input properties used for looking up and filtering Config resources.
type ConfigState struct {
	// The base64 encoded data of the config.
	Data interface{}
	// The name of the Docker config.
	Name interface{}
}

// The set of arguments for constructing a Config resource.
type ConfigArgs struct {
	// The base64 encoded data of the config.
	Data interface{}
	// The name of the Docker config.
	Name interface{}
}
