// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Removes existing fields. If one field doesn’t exist, an exception will be thrown.
//
// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/remove-processor.html
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/zlepper/pulumi-elasticstack/sdk/go/elasticstack"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		remove, err := elasticstack.GetIngestProcessorRemove(ctx, &elasticstack.GetIngestProcessorRemoveArgs{
// 			Fields: []string{
// 				"user_agent",
// 				"url",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = elasticstack.NewIngestPipeline(ctx, "myIngestPipeline", &elasticstack.IngestPipelineArgs{
// 			Processors: pulumi.StringArray{
// 				*pulumi.String(remove.Json),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetIngestProcessorRemove(ctx *pulumi.Context, args *GetIngestProcessorRemoveArgs, opts ...pulumi.InvokeOption) (*GetIngestProcessorRemoveResult, error) {
	var rv GetIngestProcessorRemoveResult
	err := ctx.Invoke("elasticstack:index/getIngestProcessorRemove:getIngestProcessorRemove", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIngestProcessorRemove.
type GetIngestProcessorRemoveArgs struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// Fields to be removed.
	Fields []string `pulumi:"fields"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
	IgnoreMissing *bool `pulumi:"ignoreMissing"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
}

// A collection of values returned by getIngestProcessorRemove.
type GetIngestProcessorRemoveResult struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// Fields to be removed.
	Fields []string `pulumi:"fields"`
	// Internal identifier of the resource.
	Id string `pulumi:"id"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
	IgnoreMissing *bool `pulumi:"ignoreMissing"`
	// JSON representation of this data source.
	Json string `pulumi:"json"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
}

func GetIngestProcessorRemoveOutput(ctx *pulumi.Context, args GetIngestProcessorRemoveOutputArgs, opts ...pulumi.InvokeOption) GetIngestProcessorRemoveResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIngestProcessorRemoveResult, error) {
			args := v.(GetIngestProcessorRemoveArgs)
			r, err := GetIngestProcessorRemove(ctx, &args, opts...)
			var s GetIngestProcessorRemoveResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIngestProcessorRemoveResultOutput)
}

// A collection of arguments for invoking getIngestProcessorRemove.
type GetIngestProcessorRemoveOutputArgs struct {
	// Description of the processor.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Fields to be removed.
	Fields pulumi.StringArrayInput `pulumi:"fields"`
	// Conditionally execute the processor
	If pulumi.StringPtrInput `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure pulumi.BoolPtrInput `pulumi:"ignoreFailure"`
	// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
	IgnoreMissing pulumi.BoolPtrInput `pulumi:"ignoreMissing"`
	// Handle failures for the processor.
	OnFailures pulumi.StringArrayInput `pulumi:"onFailures"`
	// Identifier for the processor.
	Tag pulumi.StringPtrInput `pulumi:"tag"`
}

func (GetIngestProcessorRemoveOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIngestProcessorRemoveArgs)(nil)).Elem()
}

// A collection of values returned by getIngestProcessorRemove.
type GetIngestProcessorRemoveResultOutput struct{ *pulumi.OutputState }

func (GetIngestProcessorRemoveResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIngestProcessorRemoveResult)(nil)).Elem()
}

func (o GetIngestProcessorRemoveResultOutput) ToGetIngestProcessorRemoveResultOutput() GetIngestProcessorRemoveResultOutput {
	return o
}

func (o GetIngestProcessorRemoveResultOutput) ToGetIngestProcessorRemoveResultOutputWithContext(ctx context.Context) GetIngestProcessorRemoveResultOutput {
	return o
}

// Description of the processor.
func (o GetIngestProcessorRemoveResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorRemoveResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Fields to be removed.
func (o GetIngestProcessorRemoveResultOutput) Fields() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIngestProcessorRemoveResult) []string { return v.Fields }).(pulumi.StringArrayOutput)
}

// Internal identifier of the resource.
func (o GetIngestProcessorRemoveResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIngestProcessorRemoveResult) string { return v.Id }).(pulumi.StringOutput)
}

// Conditionally execute the processor
func (o GetIngestProcessorRemoveResultOutput) If() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorRemoveResult) *string { return v.If }).(pulumi.StringPtrOutput)
}

// Ignore failures for the processor.
func (o GetIngestProcessorRemoveResultOutput) IgnoreFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorRemoveResult) *bool { return v.IgnoreFailure }).(pulumi.BoolPtrOutput)
}

// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
func (o GetIngestProcessorRemoveResultOutput) IgnoreMissing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorRemoveResult) *bool { return v.IgnoreMissing }).(pulumi.BoolPtrOutput)
}

// JSON representation of this data source.
func (o GetIngestProcessorRemoveResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v GetIngestProcessorRemoveResult) string { return v.Json }).(pulumi.StringOutput)
}

// Handle failures for the processor.
func (o GetIngestProcessorRemoveResultOutput) OnFailures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIngestProcessorRemoveResult) []string { return v.OnFailures }).(pulumi.StringArrayOutput)
}

// Identifier for the processor.
func (o GetIngestProcessorRemoveResultOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorRemoveResult) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIngestProcessorRemoveResultOutput{})
}
