// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbaas

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about a DBaas logs output graylog stream.
//
// ## Example Usage
func LookupLogsOutputGraylogStream(ctx *pulumi.Context, args *LookupLogsOutputGraylogStreamArgs, opts ...pulumi.InvokeOption) (*LookupLogsOutputGraylogStreamResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLogsOutputGraylogStreamResult
	err := ctx.Invoke("ovh:Dbaas/getLogsOutputGraylogStream:getLogsOutputGraylogStream", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLogsOutputGraylogStream.
type LookupLogsOutputGraylogStreamArgs struct {
	// The service name. It's the ID of your Logs Data Platform instance.
	ServiceName string `pulumi:"serviceName"`
	// Stream description
	Title string `pulumi:"title"`
}

// A collection of values returned by getLogsOutputGraylogStream.
type LookupLogsOutputGraylogStreamResult struct {
	CanAlert bool `pulumi:"canAlert"`
	// Cold storage compression method
	ColdStorageCompression string `pulumi:"coldStorageCompression"`
	// ColdStorage content
	ColdStorageContent string `pulumi:"coldStorageContent"`
	// Is Cold storage enabled?
	ColdStorageEnabled bool `pulumi:"coldStorageEnabled"`
	// Notify on new Cold storage archive
	ColdStorageNotifyEnabled bool `pulumi:"coldStorageNotifyEnabled"`
	// Cold storage retention in year
	ColdStorageRetention int `pulumi:"coldStorageRetention"`
	// ColdStorage destination
	ColdStorageTarget string `pulumi:"coldStorageTarget"`
	// Stream creation
	CreatedAt string `pulumi:"createdAt"`
	// Stream description
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Enable ES indexing
	IndexingEnabled bool `pulumi:"indexingEnabled"`
	// Maximum indexing size (in GB)
	IndexingMaxSize int `pulumi:"indexingMaxSize"`
	// If set, notify when size is near 80, 90 or 100 % of the maximum configured setting
	IndexingNotifyEnabled bool `pulumi:"indexingNotifyEnabled"`
	// Indicates if you are allowed to edit entry
	IsEditable bool `pulumi:"isEditable"`
	// Indicates if you are allowed to share entry
	IsShareable bool `pulumi:"isShareable"`
	// Number of alert condition
	NbAlertCondition int `pulumi:"nbAlertCondition"`
	// Number of coldstored archives
	NbArchive int `pulumi:"nbArchive"`
	// Parent stream ID
	ParentStreamId string `pulumi:"parentStreamId"`
	// If set, pause indexing when maximum size is reach
	PauseIndexingOnMaxSize bool `pulumi:"pauseIndexingOnMaxSize"`
	// Retention ID
	RetentionId string `pulumi:"retentionId"`
	ServiceName string `pulumi:"serviceName"`
	// Stream ID
	StreamId string `pulumi:"streamId"`
	Title    string `pulumi:"title"`
	// Stream last update
	UpdatedAt string `pulumi:"updatedAt"`
	// Enable Websocket
	WebSocketEnabled bool `pulumi:"webSocketEnabled"`
}

func LookupLogsOutputGraylogStreamOutput(ctx *pulumi.Context, args LookupLogsOutputGraylogStreamOutputArgs, opts ...pulumi.InvokeOption) LookupLogsOutputGraylogStreamResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLogsOutputGraylogStreamResult, error) {
			args := v.(LookupLogsOutputGraylogStreamArgs)
			r, err := LookupLogsOutputGraylogStream(ctx, &args, opts...)
			var s LookupLogsOutputGraylogStreamResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLogsOutputGraylogStreamResultOutput)
}

// A collection of arguments for invoking getLogsOutputGraylogStream.
type LookupLogsOutputGraylogStreamOutputArgs struct {
	// The service name. It's the ID of your Logs Data Platform instance.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Stream description
	Title pulumi.StringInput `pulumi:"title"`
}

func (LookupLogsOutputGraylogStreamOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLogsOutputGraylogStreamArgs)(nil)).Elem()
}

// A collection of values returned by getLogsOutputGraylogStream.
type LookupLogsOutputGraylogStreamResultOutput struct{ *pulumi.OutputState }

func (LookupLogsOutputGraylogStreamResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLogsOutputGraylogStreamResult)(nil)).Elem()
}

func (o LookupLogsOutputGraylogStreamResultOutput) ToLookupLogsOutputGraylogStreamResultOutput() LookupLogsOutputGraylogStreamResultOutput {
	return o
}

func (o LookupLogsOutputGraylogStreamResultOutput) ToLookupLogsOutputGraylogStreamResultOutputWithContext(ctx context.Context) LookupLogsOutputGraylogStreamResultOutput {
	return o
}

func (o LookupLogsOutputGraylogStreamResultOutput) CanAlert() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLogsOutputGraylogStreamResult) bool { return v.CanAlert }).(pulumi.BoolOutput)
}

// Cold storage compression method
func (o LookupLogsOutputGraylogStreamResultOutput) ColdStorageCompression() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLogsOutputGraylogStreamResult) string { return v.ColdStorageCompression }).(pulumi.StringOutput)
}

// ColdStorage content
func (o LookupLogsOutputGraylogStreamResultOutput) ColdStorageContent() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLogsOutputGraylogStreamResult) string { return v.ColdStorageContent }).(pulumi.StringOutput)
}

// Is Cold storage enabled?
func (o LookupLogsOutputGraylogStreamResultOutput) ColdStorageEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLogsOutputGraylogStreamResult) bool { return v.ColdStorageEnabled }).(pulumi.BoolOutput)
}

// Notify on new Cold storage archive
func (o LookupLogsOutputGraylogStreamResultOutput) ColdStorageNotifyEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLogsOutputGraylogStreamResult) bool { return v.ColdStorageNotifyEnabled }).(pulumi.BoolOutput)
}

// Cold storage retention in year
func (o LookupLogsOutputGraylogStreamResultOutput) ColdStorageRetention() pulumi.IntOutput {
	return o.ApplyT(func(v LookupLogsOutputGraylogStreamResult) int { return v.ColdStorageRetention }).(pulumi.IntOutput)
}

// ColdStorage destination
func (o LookupLogsOutputGraylogStreamResultOutput) ColdStorageTarget() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLogsOutputGraylogStreamResult) string { return v.ColdStorageTarget }).(pulumi.StringOutput)
}

// Stream creation
func (o LookupLogsOutputGraylogStreamResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLogsOutputGraylogStreamResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// Stream description
func (o LookupLogsOutputGraylogStreamResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLogsOutputGraylogStreamResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupLogsOutputGraylogStreamResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLogsOutputGraylogStreamResult) string { return v.Id }).(pulumi.StringOutput)
}

// Enable ES indexing
func (o LookupLogsOutputGraylogStreamResultOutput) IndexingEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLogsOutputGraylogStreamResult) bool { return v.IndexingEnabled }).(pulumi.BoolOutput)
}

// Maximum indexing size (in GB)
func (o LookupLogsOutputGraylogStreamResultOutput) IndexingMaxSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupLogsOutputGraylogStreamResult) int { return v.IndexingMaxSize }).(pulumi.IntOutput)
}

// If set, notify when size is near 80, 90 or 100 % of the maximum configured setting
func (o LookupLogsOutputGraylogStreamResultOutput) IndexingNotifyEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLogsOutputGraylogStreamResult) bool { return v.IndexingNotifyEnabled }).(pulumi.BoolOutput)
}

// Indicates if you are allowed to edit entry
func (o LookupLogsOutputGraylogStreamResultOutput) IsEditable() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLogsOutputGraylogStreamResult) bool { return v.IsEditable }).(pulumi.BoolOutput)
}

// Indicates if you are allowed to share entry
func (o LookupLogsOutputGraylogStreamResultOutput) IsShareable() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLogsOutputGraylogStreamResult) bool { return v.IsShareable }).(pulumi.BoolOutput)
}

// Number of alert condition
func (o LookupLogsOutputGraylogStreamResultOutput) NbAlertCondition() pulumi.IntOutput {
	return o.ApplyT(func(v LookupLogsOutputGraylogStreamResult) int { return v.NbAlertCondition }).(pulumi.IntOutput)
}

// Number of coldstored archives
func (o LookupLogsOutputGraylogStreamResultOutput) NbArchive() pulumi.IntOutput {
	return o.ApplyT(func(v LookupLogsOutputGraylogStreamResult) int { return v.NbArchive }).(pulumi.IntOutput)
}

// Parent stream ID
func (o LookupLogsOutputGraylogStreamResultOutput) ParentStreamId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLogsOutputGraylogStreamResult) string { return v.ParentStreamId }).(pulumi.StringOutput)
}

// If set, pause indexing when maximum size is reach
func (o LookupLogsOutputGraylogStreamResultOutput) PauseIndexingOnMaxSize() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLogsOutputGraylogStreamResult) bool { return v.PauseIndexingOnMaxSize }).(pulumi.BoolOutput)
}

// Retention ID
func (o LookupLogsOutputGraylogStreamResultOutput) RetentionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLogsOutputGraylogStreamResult) string { return v.RetentionId }).(pulumi.StringOutput)
}

func (o LookupLogsOutputGraylogStreamResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLogsOutputGraylogStreamResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// Stream ID
func (o LookupLogsOutputGraylogStreamResultOutput) StreamId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLogsOutputGraylogStreamResult) string { return v.StreamId }).(pulumi.StringOutput)
}

func (o LookupLogsOutputGraylogStreamResultOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLogsOutputGraylogStreamResult) string { return v.Title }).(pulumi.StringOutput)
}

// Stream last update
func (o LookupLogsOutputGraylogStreamResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLogsOutputGraylogStreamResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

// Enable Websocket
func (o LookupLogsOutputGraylogStreamResultOutput) WebSocketEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLogsOutputGraylogStreamResult) bool { return v.WebSocketEnabled }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLogsOutputGraylogStreamResultOutput{})
}
