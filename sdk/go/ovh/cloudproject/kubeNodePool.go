// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a nodepool in a OVHcloud Managed Kubernetes Service cluster.
//
// ## Example Usage
//
// Create a simple node pool in your Kubernetes cluster:
//
// Create an advanced node pool in your Kubernetes cluster:
//
// ## Import
//
// OVHcloud Managed Kubernetes Service cluster node pool can be imported using the `service_name`, the `id` of the cluster, and the `id` of the nodepool separated by "/" E.g., bash <break><break>```sh<break> $ pulumi import ovh:CloudProject/kubeNodePool:KubeNodePool pool service_name/kube_id/poolid <break>```<break><break>
type KubeNodePool struct {
	pulumi.CustomResourceState

	// should the pool use the anti-affinity feature. Default to `false`. **Changing this value recreates the resource.**
	AntiAffinity pulumi.BoolOutput `pulumi:"antiAffinity"`
	// Enable auto-scaling for the pool. Default to `false`.
	// * ` template  ` - (Optional) Managed Kubernetes nodepool template, which is a complex object constituted by two main nested objects:
	Autoscale pulumi.BoolOutput `pulumi:"autoscale"`
	// Number of nodes which are actually ready in the pool
	AvailableNodes pulumi.IntOutput `pulumi:"availableNodes"`
	// Creation date
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Number of nodes present in the pool
	CurrentNodes pulumi.IntOutput `pulumi:"currentNodes"`
	// number of nodes to start.
	DesiredNodes pulumi.IntOutput `pulumi:"desiredNodes"`
	// Flavor name
	Flavor pulumi.StringOutput `pulumi:"flavor"`
	// a valid OVHcloud public cloud flavor ID in which the nodes will be started. Ex: "b2-7". You can find the list of flavor IDs: https://www.ovhcloud.com/fr/public-cloud/prices/.
	// **Changing this value recreates the resource.**
	FlavorName pulumi.StringOutput `pulumi:"flavorName"`
	// The id of the managed kubernetes cluster. **Changing this value recreates the resource.**
	KubeId pulumi.StringOutput `pulumi:"kubeId"`
	// maximum number of nodes allowed in the pool. Setting `desiredNodes` over this value will raise an error.
	MaxNodes pulumi.IntOutput `pulumi:"maxNodes"`
	// minimum number of nodes allowed in the pool. Setting `desiredNodes` under this value will raise an error.
	MinNodes pulumi.IntOutput `pulumi:"minNodes"`
	// should the nodes be billed on a monthly basis. Default to `false`. **Changing this value recreates the resource.**
	MonthlyBilled pulumi.BoolOutput `pulumi:"monthlyBilled"`
	// The name of the nodepool. Warning: `_` char is not allowed! **Changing this value recreates the resource.**
	Name pulumi.StringOutput `pulumi:"name"`
	// Project id
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Status describing the state between number of nodes wanted and available ones
	SizeStatus pulumi.StringOutput `pulumi:"sizeStatus"`
	// Current status
	Status pulumi.StringOutput `pulumi:"status"`
	// Node pool template
	Template KubeNodePoolTemplatePtrOutput `pulumi:"template"`
	// Number of nodes with the latest version installed in the pool
	UpToDateNodes pulumi.IntOutput `pulumi:"upToDateNodes"`
	// Last update date
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewKubeNodePool registers a new resource with the given unique name, arguments, and options.
func NewKubeNodePool(ctx *pulumi.Context,
	name string, args *KubeNodePoolArgs, opts ...pulumi.ResourceOption) (*KubeNodePool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FlavorName == nil {
		return nil, errors.New("invalid value for required argument 'FlavorName'")
	}
	if args.KubeId == nil {
		return nil, errors.New("invalid value for required argument 'KubeId'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource KubeNodePool
	err := ctx.RegisterResource("ovh:CloudProject/kubeNodePool:KubeNodePool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKubeNodePool gets an existing KubeNodePool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKubeNodePool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KubeNodePoolState, opts ...pulumi.ResourceOption) (*KubeNodePool, error) {
	var resource KubeNodePool
	err := ctx.ReadResource("ovh:CloudProject/kubeNodePool:KubeNodePool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KubeNodePool resources.
type kubeNodePoolState struct {
	// should the pool use the anti-affinity feature. Default to `false`. **Changing this value recreates the resource.**
	AntiAffinity *bool `pulumi:"antiAffinity"`
	// Enable auto-scaling for the pool. Default to `false`.
	// * ` template  ` - (Optional) Managed Kubernetes nodepool template, which is a complex object constituted by two main nested objects:
	Autoscale *bool `pulumi:"autoscale"`
	// Number of nodes which are actually ready in the pool
	AvailableNodes *int `pulumi:"availableNodes"`
	// Creation date
	CreatedAt *string `pulumi:"createdAt"`
	// Number of nodes present in the pool
	CurrentNodes *int `pulumi:"currentNodes"`
	// number of nodes to start.
	DesiredNodes *int `pulumi:"desiredNodes"`
	// Flavor name
	Flavor *string `pulumi:"flavor"`
	// a valid OVHcloud public cloud flavor ID in which the nodes will be started. Ex: "b2-7". You can find the list of flavor IDs: https://www.ovhcloud.com/fr/public-cloud/prices/.
	// **Changing this value recreates the resource.**
	FlavorName *string `pulumi:"flavorName"`
	// The id of the managed kubernetes cluster. **Changing this value recreates the resource.**
	KubeId *string `pulumi:"kubeId"`
	// maximum number of nodes allowed in the pool. Setting `desiredNodes` over this value will raise an error.
	MaxNodes *int `pulumi:"maxNodes"`
	// minimum number of nodes allowed in the pool. Setting `desiredNodes` under this value will raise an error.
	MinNodes *int `pulumi:"minNodes"`
	// should the nodes be billed on a monthly basis. Default to `false`. **Changing this value recreates the resource.**
	MonthlyBilled *bool `pulumi:"monthlyBilled"`
	// The name of the nodepool. Warning: `_` char is not allowed! **Changing this value recreates the resource.**
	Name *string `pulumi:"name"`
	// Project id
	ProjectId *string `pulumi:"projectId"`
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
	ServiceName *string `pulumi:"serviceName"`
	// Status describing the state between number of nodes wanted and available ones
	SizeStatus *string `pulumi:"sizeStatus"`
	// Current status
	Status *string `pulumi:"status"`
	// Node pool template
	Template *KubeNodePoolTemplate `pulumi:"template"`
	// Number of nodes with the latest version installed in the pool
	UpToDateNodes *int `pulumi:"upToDateNodes"`
	// Last update date
	UpdatedAt *string `pulumi:"updatedAt"`
}

type KubeNodePoolState struct {
	// should the pool use the anti-affinity feature. Default to `false`. **Changing this value recreates the resource.**
	AntiAffinity pulumi.BoolPtrInput
	// Enable auto-scaling for the pool. Default to `false`.
	// * ` template  ` - (Optional) Managed Kubernetes nodepool template, which is a complex object constituted by two main nested objects:
	Autoscale pulumi.BoolPtrInput
	// Number of nodes which are actually ready in the pool
	AvailableNodes pulumi.IntPtrInput
	// Creation date
	CreatedAt pulumi.StringPtrInput
	// Number of nodes present in the pool
	CurrentNodes pulumi.IntPtrInput
	// number of nodes to start.
	DesiredNodes pulumi.IntPtrInput
	// Flavor name
	Flavor pulumi.StringPtrInput
	// a valid OVHcloud public cloud flavor ID in which the nodes will be started. Ex: "b2-7". You can find the list of flavor IDs: https://www.ovhcloud.com/fr/public-cloud/prices/.
	// **Changing this value recreates the resource.**
	FlavorName pulumi.StringPtrInput
	// The id of the managed kubernetes cluster. **Changing this value recreates the resource.**
	KubeId pulumi.StringPtrInput
	// maximum number of nodes allowed in the pool. Setting `desiredNodes` over this value will raise an error.
	MaxNodes pulumi.IntPtrInput
	// minimum number of nodes allowed in the pool. Setting `desiredNodes` under this value will raise an error.
	MinNodes pulumi.IntPtrInput
	// should the nodes be billed on a monthly basis. Default to `false`. **Changing this value recreates the resource.**
	MonthlyBilled pulumi.BoolPtrInput
	// The name of the nodepool. Warning: `_` char is not allowed! **Changing this value recreates the resource.**
	Name pulumi.StringPtrInput
	// Project id
	ProjectId pulumi.StringPtrInput
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
	ServiceName pulumi.StringPtrInput
	// Status describing the state between number of nodes wanted and available ones
	SizeStatus pulumi.StringPtrInput
	// Current status
	Status pulumi.StringPtrInput
	// Node pool template
	Template KubeNodePoolTemplatePtrInput
	// Number of nodes with the latest version installed in the pool
	UpToDateNodes pulumi.IntPtrInput
	// Last update date
	UpdatedAt pulumi.StringPtrInput
}

func (KubeNodePoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*kubeNodePoolState)(nil)).Elem()
}

type kubeNodePoolArgs struct {
	// should the pool use the anti-affinity feature. Default to `false`. **Changing this value recreates the resource.**
	AntiAffinity *bool `pulumi:"antiAffinity"`
	// Enable auto-scaling for the pool. Default to `false`.
	// * ` template  ` - (Optional) Managed Kubernetes nodepool template, which is a complex object constituted by two main nested objects:
	Autoscale *bool `pulumi:"autoscale"`
	// number of nodes to start.
	DesiredNodes *int `pulumi:"desiredNodes"`
	// a valid OVHcloud public cloud flavor ID in which the nodes will be started. Ex: "b2-7". You can find the list of flavor IDs: https://www.ovhcloud.com/fr/public-cloud/prices/.
	// **Changing this value recreates the resource.**
	FlavorName string `pulumi:"flavorName"`
	// The id of the managed kubernetes cluster. **Changing this value recreates the resource.**
	KubeId string `pulumi:"kubeId"`
	// maximum number of nodes allowed in the pool. Setting `desiredNodes` over this value will raise an error.
	MaxNodes *int `pulumi:"maxNodes"`
	// minimum number of nodes allowed in the pool. Setting `desiredNodes` under this value will raise an error.
	MinNodes *int `pulumi:"minNodes"`
	// should the nodes be billed on a monthly basis. Default to `false`. **Changing this value recreates the resource.**
	MonthlyBilled *bool `pulumi:"monthlyBilled"`
	// The name of the nodepool. Warning: `_` char is not allowed! **Changing this value recreates the resource.**
	Name *string `pulumi:"name"`
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
	ServiceName string `pulumi:"serviceName"`
	// Node pool template
	Template *KubeNodePoolTemplate `pulumi:"template"`
}

// The set of arguments for constructing a KubeNodePool resource.
type KubeNodePoolArgs struct {
	// should the pool use the anti-affinity feature. Default to `false`. **Changing this value recreates the resource.**
	AntiAffinity pulumi.BoolPtrInput
	// Enable auto-scaling for the pool. Default to `false`.
	// * ` template  ` - (Optional) Managed Kubernetes nodepool template, which is a complex object constituted by two main nested objects:
	Autoscale pulumi.BoolPtrInput
	// number of nodes to start.
	DesiredNodes pulumi.IntPtrInput
	// a valid OVHcloud public cloud flavor ID in which the nodes will be started. Ex: "b2-7". You can find the list of flavor IDs: https://www.ovhcloud.com/fr/public-cloud/prices/.
	// **Changing this value recreates the resource.**
	FlavorName pulumi.StringInput
	// The id of the managed kubernetes cluster. **Changing this value recreates the resource.**
	KubeId pulumi.StringInput
	// maximum number of nodes allowed in the pool. Setting `desiredNodes` over this value will raise an error.
	MaxNodes pulumi.IntPtrInput
	// minimum number of nodes allowed in the pool. Setting `desiredNodes` under this value will raise an error.
	MinNodes pulumi.IntPtrInput
	// should the nodes be billed on a monthly basis. Default to `false`. **Changing this value recreates the resource.**
	MonthlyBilled pulumi.BoolPtrInput
	// The name of the nodepool. Warning: `_` char is not allowed! **Changing this value recreates the resource.**
	Name pulumi.StringPtrInput
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
	ServiceName pulumi.StringInput
	// Node pool template
	Template KubeNodePoolTemplatePtrInput
}

func (KubeNodePoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubeNodePoolArgs)(nil)).Elem()
}

type KubeNodePoolInput interface {
	pulumi.Input

	ToKubeNodePoolOutput() KubeNodePoolOutput
	ToKubeNodePoolOutputWithContext(ctx context.Context) KubeNodePoolOutput
}

func (*KubeNodePool) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeNodePool)(nil)).Elem()
}

func (i *KubeNodePool) ToKubeNodePoolOutput() KubeNodePoolOutput {
	return i.ToKubeNodePoolOutputWithContext(context.Background())
}

func (i *KubeNodePool) ToKubeNodePoolOutputWithContext(ctx context.Context) KubeNodePoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeNodePoolOutput)
}

// KubeNodePoolArrayInput is an input type that accepts KubeNodePoolArray and KubeNodePoolArrayOutput values.
// You can construct a concrete instance of `KubeNodePoolArrayInput` via:
//
//	KubeNodePoolArray{ KubeNodePoolArgs{...} }
type KubeNodePoolArrayInput interface {
	pulumi.Input

	ToKubeNodePoolArrayOutput() KubeNodePoolArrayOutput
	ToKubeNodePoolArrayOutputWithContext(context.Context) KubeNodePoolArrayOutput
}

type KubeNodePoolArray []KubeNodePoolInput

func (KubeNodePoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KubeNodePool)(nil)).Elem()
}

func (i KubeNodePoolArray) ToKubeNodePoolArrayOutput() KubeNodePoolArrayOutput {
	return i.ToKubeNodePoolArrayOutputWithContext(context.Background())
}

func (i KubeNodePoolArray) ToKubeNodePoolArrayOutputWithContext(ctx context.Context) KubeNodePoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeNodePoolArrayOutput)
}

// KubeNodePoolMapInput is an input type that accepts KubeNodePoolMap and KubeNodePoolMapOutput values.
// You can construct a concrete instance of `KubeNodePoolMapInput` via:
//
//	KubeNodePoolMap{ "key": KubeNodePoolArgs{...} }
type KubeNodePoolMapInput interface {
	pulumi.Input

	ToKubeNodePoolMapOutput() KubeNodePoolMapOutput
	ToKubeNodePoolMapOutputWithContext(context.Context) KubeNodePoolMapOutput
}

type KubeNodePoolMap map[string]KubeNodePoolInput

func (KubeNodePoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KubeNodePool)(nil)).Elem()
}

func (i KubeNodePoolMap) ToKubeNodePoolMapOutput() KubeNodePoolMapOutput {
	return i.ToKubeNodePoolMapOutputWithContext(context.Background())
}

func (i KubeNodePoolMap) ToKubeNodePoolMapOutputWithContext(ctx context.Context) KubeNodePoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeNodePoolMapOutput)
}

type KubeNodePoolOutput struct{ *pulumi.OutputState }

func (KubeNodePoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeNodePool)(nil)).Elem()
}

func (o KubeNodePoolOutput) ToKubeNodePoolOutput() KubeNodePoolOutput {
	return o
}

func (o KubeNodePoolOutput) ToKubeNodePoolOutputWithContext(ctx context.Context) KubeNodePoolOutput {
	return o
}

// should the pool use the anti-affinity feature. Default to `false`. **Changing this value recreates the resource.**
func (o KubeNodePoolOutput) AntiAffinity() pulumi.BoolOutput {
	return o.ApplyT(func(v *KubeNodePool) pulumi.BoolOutput { return v.AntiAffinity }).(pulumi.BoolOutput)
}

// Enable auto-scaling for the pool. Default to `false`.
// * ` template  ` - (Optional) Managed Kubernetes nodepool template, which is a complex object constituted by two main nested objects:
func (o KubeNodePoolOutput) Autoscale() pulumi.BoolOutput {
	return o.ApplyT(func(v *KubeNodePool) pulumi.BoolOutput { return v.Autoscale }).(pulumi.BoolOutput)
}

// Number of nodes which are actually ready in the pool
func (o KubeNodePoolOutput) AvailableNodes() pulumi.IntOutput {
	return o.ApplyT(func(v *KubeNodePool) pulumi.IntOutput { return v.AvailableNodes }).(pulumi.IntOutput)
}

// Creation date
func (o KubeNodePoolOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeNodePool) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Number of nodes present in the pool
func (o KubeNodePoolOutput) CurrentNodes() pulumi.IntOutput {
	return o.ApplyT(func(v *KubeNodePool) pulumi.IntOutput { return v.CurrentNodes }).(pulumi.IntOutput)
}

// number of nodes to start.
func (o KubeNodePoolOutput) DesiredNodes() pulumi.IntOutput {
	return o.ApplyT(func(v *KubeNodePool) pulumi.IntOutput { return v.DesiredNodes }).(pulumi.IntOutput)
}

// Flavor name
func (o KubeNodePoolOutput) Flavor() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeNodePool) pulumi.StringOutput { return v.Flavor }).(pulumi.StringOutput)
}

// a valid OVHcloud public cloud flavor ID in which the nodes will be started. Ex: "b2-7". You can find the list of flavor IDs: https://www.ovhcloud.com/fr/public-cloud/prices/.
// **Changing this value recreates the resource.**
func (o KubeNodePoolOutput) FlavorName() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeNodePool) pulumi.StringOutput { return v.FlavorName }).(pulumi.StringOutput)
}

// The id of the managed kubernetes cluster. **Changing this value recreates the resource.**
func (o KubeNodePoolOutput) KubeId() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeNodePool) pulumi.StringOutput { return v.KubeId }).(pulumi.StringOutput)
}

// maximum number of nodes allowed in the pool. Setting `desiredNodes` over this value will raise an error.
func (o KubeNodePoolOutput) MaxNodes() pulumi.IntOutput {
	return o.ApplyT(func(v *KubeNodePool) pulumi.IntOutput { return v.MaxNodes }).(pulumi.IntOutput)
}

// minimum number of nodes allowed in the pool. Setting `desiredNodes` under this value will raise an error.
func (o KubeNodePoolOutput) MinNodes() pulumi.IntOutput {
	return o.ApplyT(func(v *KubeNodePool) pulumi.IntOutput { return v.MinNodes }).(pulumi.IntOutput)
}

// should the nodes be billed on a monthly basis. Default to `false`. **Changing this value recreates the resource.**
func (o KubeNodePoolOutput) MonthlyBilled() pulumi.BoolOutput {
	return o.ApplyT(func(v *KubeNodePool) pulumi.BoolOutput { return v.MonthlyBilled }).(pulumi.BoolOutput)
}

// The name of the nodepool. Warning: `_` char is not allowed! **Changing this value recreates the resource.**
func (o KubeNodePoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeNodePool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Project id
func (o KubeNodePoolOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeNodePool) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
func (o KubeNodePoolOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeNodePool) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Status describing the state between number of nodes wanted and available ones
func (o KubeNodePoolOutput) SizeStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeNodePool) pulumi.StringOutput { return v.SizeStatus }).(pulumi.StringOutput)
}

// Current status
func (o KubeNodePoolOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeNodePool) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Node pool template
func (o KubeNodePoolOutput) Template() KubeNodePoolTemplatePtrOutput {
	return o.ApplyT(func(v *KubeNodePool) KubeNodePoolTemplatePtrOutput { return v.Template }).(KubeNodePoolTemplatePtrOutput)
}

// Number of nodes with the latest version installed in the pool
func (o KubeNodePoolOutput) UpToDateNodes() pulumi.IntOutput {
	return o.ApplyT(func(v *KubeNodePool) pulumi.IntOutput { return v.UpToDateNodes }).(pulumi.IntOutput)
}

// Last update date
func (o KubeNodePoolOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeNodePool) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type KubeNodePoolArrayOutput struct{ *pulumi.OutputState }

func (KubeNodePoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KubeNodePool)(nil)).Elem()
}

func (o KubeNodePoolArrayOutput) ToKubeNodePoolArrayOutput() KubeNodePoolArrayOutput {
	return o
}

func (o KubeNodePoolArrayOutput) ToKubeNodePoolArrayOutputWithContext(ctx context.Context) KubeNodePoolArrayOutput {
	return o
}

func (o KubeNodePoolArrayOutput) Index(i pulumi.IntInput) KubeNodePoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KubeNodePool {
		return vs[0].([]*KubeNodePool)[vs[1].(int)]
	}).(KubeNodePoolOutput)
}

type KubeNodePoolMapOutput struct{ *pulumi.OutputState }

func (KubeNodePoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KubeNodePool)(nil)).Elem()
}

func (o KubeNodePoolMapOutput) ToKubeNodePoolMapOutput() KubeNodePoolMapOutput {
	return o
}

func (o KubeNodePoolMapOutput) ToKubeNodePoolMapOutputWithContext(ctx context.Context) KubeNodePoolMapOutput {
	return o
}

func (o KubeNodePoolMapOutput) MapIndex(k pulumi.StringInput) KubeNodePoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KubeNodePool {
		return vs[0].(map[string]*KubeNodePool)[vs[1].(string)]
	}).(KubeNodePoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KubeNodePoolInput)(nil)).Elem(), &KubeNodePool{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubeNodePoolArrayInput)(nil)).Elem(), KubeNodePoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubeNodePoolMapInput)(nil)).Elem(), KubeNodePoolMap{})
	pulumi.RegisterOutputType(KubeNodePoolOutput{})
	pulumi.RegisterOutputType(KubeNodePoolArrayOutput{})
	pulumi.RegisterOutputType(KubeNodePoolMapOutput{})
}
