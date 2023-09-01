// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the user details of a previously created public cloud project user.
//
// ## Example Usage
func LookupUser(ctx *pulumi.Context, args *LookupUserArgs, opts ...pulumi.InvokeOption) (*LookupUserResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupUserResult
	err := ctx.Invoke("ovh:CloudProject/getUser:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUser.
type LookupUserArgs struct {
	// The ID of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
	// The ID of a public cloud project's user.
	UserId string `pulumi:"userId"`
}

// A collection of values returned by getUser.
type LookupUserResult struct {
	// the date the user was created.
	CreationDate string `pulumi:"creationDate"`
	// description of the role
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of roles associated with the user.
	Roles       []GetUserRole `pulumi:"roles"`
	ServiceName string        `pulumi:"serviceName"`
	// the status of the user. should be normally set to 'ok'.
	Status string `pulumi:"status"`
	UserId string `pulumi:"userId"`
	// the username generated for the user. This username can be used with
	// the Openstack API.
	Username string `pulumi:"username"`
}

func LookupUserOutput(ctx *pulumi.Context, args LookupUserOutputArgs, opts ...pulumi.InvokeOption) LookupUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserResult, error) {
			args := v.(LookupUserArgs)
			r, err := LookupUser(ctx, &args, opts...)
			var s LookupUserResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserResultOutput)
}

// A collection of arguments for invoking getUser.
type LookupUserOutputArgs struct {
	// The ID of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// The ID of a public cloud project's user.
	UserId pulumi.StringInput `pulumi:"userId"`
}

func (LookupUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserArgs)(nil)).Elem()
}

// A collection of values returned by getUser.
type LookupUserResultOutput struct{ *pulumi.OutputState }

func (LookupUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserResult)(nil)).Elem()
}

func (o LookupUserResultOutput) ToLookupUserResultOutput() LookupUserResultOutput {
	return o
}

func (o LookupUserResultOutput) ToLookupUserResultOutputWithContext(ctx context.Context) LookupUserResultOutput {
	return o
}

// the date the user was created.
func (o LookupUserResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

// description of the role
func (o LookupUserResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of roles associated with the user.
func (o LookupUserResultOutput) Roles() GetUserRoleArrayOutput {
	return o.ApplyT(func(v LookupUserResult) []GetUserRole { return v.Roles }).(GetUserRoleArrayOutput)
}

func (o LookupUserResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// the status of the user. should be normally set to 'ok'.
func (o LookupUserResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupUserResultOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.UserId }).(pulumi.StringOutput)
}

// the username generated for the user. This username can be used with
// the Openstack API.
func (o LookupUserResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserResultOutput{})
}
