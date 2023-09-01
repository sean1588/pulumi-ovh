// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudprojectdatabase

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the list of managed databases of a public cloud project.
//
// ## Example Usage
//
// To get the list of database clusters service for a given engine:
func GetDatabases(ctx *pulumi.Context, args *GetDatabasesArgs, opts ...pulumi.InvokeOption) (*GetDatabasesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDatabasesResult
	err := ctx.Invoke("ovh:CloudProjectDatabase/getDatabases:getDatabases", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatabases.
type GetDatabasesArgs struct {
	// The database engine you want to list. To get a full list of available engine visit:
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	Engine string `pulumi:"engine"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getDatabases.
type GetDatabasesResult struct {
	// The list of managed databases ids of the project.
	ClusterIds []string `pulumi:"clusterIds"`
	// See Argument Reference above.
	Engine string `pulumi:"engine"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	ServiceName string `pulumi:"serviceName"`
}

func GetDatabasesOutput(ctx *pulumi.Context, args GetDatabasesOutputArgs, opts ...pulumi.InvokeOption) GetDatabasesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDatabasesResult, error) {
			args := v.(GetDatabasesArgs)
			r, err := GetDatabases(ctx, &args, opts...)
			var s GetDatabasesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDatabasesResultOutput)
}

// A collection of arguments for invoking getDatabases.
type GetDatabasesOutputArgs struct {
	// The database engine you want to list. To get a full list of available engine visit:
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	Engine pulumi.StringInput `pulumi:"engine"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetDatabasesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabasesArgs)(nil)).Elem()
}

// A collection of values returned by getDatabases.
type GetDatabasesResultOutput struct{ *pulumi.OutputState }

func (GetDatabasesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabasesResult)(nil)).Elem()
}

func (o GetDatabasesResultOutput) ToGetDatabasesResultOutput() GetDatabasesResultOutput {
	return o
}

func (o GetDatabasesResultOutput) ToGetDatabasesResultOutputWithContext(ctx context.Context) GetDatabasesResultOutput {
	return o
}

// The list of managed databases ids of the project.
func (o GetDatabasesResultOutput) ClusterIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDatabasesResult) []string { return v.ClusterIds }).(pulumi.StringArrayOutput)
}

// See Argument Reference above.
func (o GetDatabasesResultOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabasesResult) string { return v.Engine }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDatabasesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabasesResult) string { return v.Id }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetDatabasesResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabasesResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDatabasesResultOutput{})
}
