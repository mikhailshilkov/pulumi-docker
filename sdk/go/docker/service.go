// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package docker

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-docker/blob/master/website/docs/r/service.html.markdown.
type Service struct {
	s *pulumi.ResourceState
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOpt) (*Service, error) {
	if args == nil || args.TaskSpec == nil {
		return nil, errors.New("missing required argument 'TaskSpec'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["auth"] = nil
		inputs["convergeConfig"] = nil
		inputs["endpointSpec"] = nil
		inputs["labels"] = nil
		inputs["mode"] = nil
		inputs["name"] = nil
		inputs["rollbackConfig"] = nil
		inputs["taskSpec"] = nil
		inputs["updateConfig"] = nil
	} else {
		inputs["auth"] = args.Auth
		inputs["convergeConfig"] = args.ConvergeConfig
		inputs["endpointSpec"] = args.EndpointSpec
		inputs["labels"] = args.Labels
		inputs["mode"] = args.Mode
		inputs["name"] = args.Name
		inputs["rollbackConfig"] = args.RollbackConfig
		inputs["taskSpec"] = args.TaskSpec
		inputs["updateConfig"] = args.UpdateConfig
	}
	s, err := ctx.RegisterResource("docker:index/service:Service", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Service{s: s}, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ServiceState, opts ...pulumi.ResourceOpt) (*Service, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["auth"] = state.Auth
		inputs["convergeConfig"] = state.ConvergeConfig
		inputs["endpointSpec"] = state.EndpointSpec
		inputs["labels"] = state.Labels
		inputs["mode"] = state.Mode
		inputs["name"] = state.Name
		inputs["rollbackConfig"] = state.RollbackConfig
		inputs["taskSpec"] = state.TaskSpec
		inputs["updateConfig"] = state.UpdateConfig
	}
	s, err := ctx.ReadResource("docker:index/service:Service", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Service{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Service) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Service) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// See Auth below for details.
func (r *Service) Auth() *pulumi.Output {
	return r.s.State["auth"]
}

// See Converge Config below for details.
func (r *Service) ConvergeConfig() *pulumi.Output {
	return r.s.State["convergeConfig"]
}

// See EndpointSpec below for details.
func (r *Service) EndpointSpec() *pulumi.Output {
	return r.s.State["endpointSpec"]
}

// User-defined key/value metadata
func (r *Service) Labels() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["labels"])
}

// The mode of resolution to use for internal load balancing between tasks. `(vip|dnsrr)`. Default: `vip`.
func (r *Service) Mode() *pulumi.Output {
	return r.s.State["mode"]
}

// A random name for the port.
func (r *Service) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// See RollbackConfig below for details.
func (r *Service) RollbackConfig() *pulumi.Output {
	return r.s.State["rollbackConfig"]
}

// See TaskSpec below for details.
func (r *Service) TaskSpec() *pulumi.Output {
	return r.s.State["taskSpec"]
}

// See UpdateConfig below for details.
func (r *Service) UpdateConfig() *pulumi.Output {
	return r.s.State["updateConfig"]
}

// Input properties used for looking up and filtering Service resources.
type ServiceState struct {
	// See Auth below for details.
	Auth interface{}
	// See Converge Config below for details.
	ConvergeConfig interface{}
	// See EndpointSpec below for details.
	EndpointSpec interface{}
	// User-defined key/value metadata
	Labels interface{}
	// The mode of resolution to use for internal load balancing between tasks. `(vip|dnsrr)`. Default: `vip`.
	Mode interface{}
	// A random name for the port.
	Name interface{}
	// See RollbackConfig below for details.
	RollbackConfig interface{}
	// See TaskSpec below for details.
	TaskSpec interface{}
	// See UpdateConfig below for details.
	UpdateConfig interface{}
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// See Auth below for details.
	Auth interface{}
	// See Converge Config below for details.
	ConvergeConfig interface{}
	// See EndpointSpec below for details.
	EndpointSpec interface{}
	// User-defined key/value metadata
	Labels interface{}
	// The mode of resolution to use for internal load balancing between tasks. `(vip|dnsrr)`. Default: `vip`.
	Mode interface{}
	// A random name for the port.
	Name interface{}
	// See RollbackConfig below for details.
	RollbackConfig interface{}
	// See TaskSpec below for details.
	TaskSpec interface{}
	// See UpdateConfig below for details.
	UpdateConfig interface{}
}
