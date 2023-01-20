// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// URL-decodes a string. If the field is an array of strings, all members of the array will be decoded.
//
// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/urldecode-processor.html
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
// 		urldecode, err := elasticstack.GetIngestProcessorUrldecode(ctx, &elasticstack.GetIngestProcessorUrldecodeArgs{
// 			Field: "my_url_to_decode",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = elasticstack.NewIngestPipeline(ctx, "myIngestPipeline", &elasticstack.IngestPipelineArgs{
// 			Processors: pulumi.StringArray{
// 				*pulumi.String(urldecode.Json),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetIngestProcessorUrldecode(ctx *pulumi.Context, args *GetIngestProcessorUrldecodeArgs, opts ...pulumi.InvokeOption) (*GetIngestProcessorUrldecodeResult, error) {
	var rv GetIngestProcessorUrldecodeResult
	err := ctx.Invoke("elasticstack:index/getIngestProcessorUrldecode:getIngestProcessorUrldecode", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIngestProcessorUrldecode.
type GetIngestProcessorUrldecodeArgs struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// The field to decode
	Field string `pulumi:"field"`
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
	// The field to assign the converted value to, by default `field` is updated in-place.
	TargetField *string `pulumi:"targetField"`
}

// A collection of values returned by getIngestProcessorUrldecode.
type GetIngestProcessorUrldecodeResult struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// The field to decode
	Field string `pulumi:"field"`
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
	// The field to assign the converted value to, by default `field` is updated in-place.
	TargetField *string `pulumi:"targetField"`
}

func GetIngestProcessorUrldecodeOutput(ctx *pulumi.Context, args GetIngestProcessorUrldecodeOutputArgs, opts ...pulumi.InvokeOption) GetIngestProcessorUrldecodeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIngestProcessorUrldecodeResult, error) {
			args := v.(GetIngestProcessorUrldecodeArgs)
			r, err := GetIngestProcessorUrldecode(ctx, &args, opts...)
			var s GetIngestProcessorUrldecodeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIngestProcessorUrldecodeResultOutput)
}

// A collection of arguments for invoking getIngestProcessorUrldecode.
type GetIngestProcessorUrldecodeOutputArgs struct {
	// Description of the processor.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The field to decode
	Field pulumi.StringInput `pulumi:"field"`
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
	// The field to assign the converted value to, by default `field` is updated in-place.
	TargetField pulumi.StringPtrInput `pulumi:"targetField"`
}

func (GetIngestProcessorUrldecodeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIngestProcessorUrldecodeArgs)(nil)).Elem()
}

// A collection of values returned by getIngestProcessorUrldecode.
type GetIngestProcessorUrldecodeResultOutput struct{ *pulumi.OutputState }

func (GetIngestProcessorUrldecodeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIngestProcessorUrldecodeResult)(nil)).Elem()
}

func (o GetIngestProcessorUrldecodeResultOutput) ToGetIngestProcessorUrldecodeResultOutput() GetIngestProcessorUrldecodeResultOutput {
	return o
}

func (o GetIngestProcessorUrldecodeResultOutput) ToGetIngestProcessorUrldecodeResultOutputWithContext(ctx context.Context) GetIngestProcessorUrldecodeResultOutput {
	return o
}

// Description of the processor.
func (o GetIngestProcessorUrldecodeResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorUrldecodeResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The field to decode
func (o GetIngestProcessorUrldecodeResultOutput) Field() pulumi.StringOutput {
	return o.ApplyT(func(v GetIngestProcessorUrldecodeResult) string { return v.Field }).(pulumi.StringOutput)
}

// Internal identifier of the resource.
func (o GetIngestProcessorUrldecodeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIngestProcessorUrldecodeResult) string { return v.Id }).(pulumi.StringOutput)
}

// Conditionally execute the processor
func (o GetIngestProcessorUrldecodeResultOutput) If() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorUrldecodeResult) *string { return v.If }).(pulumi.StringPtrOutput)
}

// Ignore failures for the processor.
func (o GetIngestProcessorUrldecodeResultOutput) IgnoreFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorUrldecodeResult) *bool { return v.IgnoreFailure }).(pulumi.BoolPtrOutput)
}

// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
func (o GetIngestProcessorUrldecodeResultOutput) IgnoreMissing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorUrldecodeResult) *bool { return v.IgnoreMissing }).(pulumi.BoolPtrOutput)
}

// JSON representation of this data source.
func (o GetIngestProcessorUrldecodeResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v GetIngestProcessorUrldecodeResult) string { return v.Json }).(pulumi.StringOutput)
}

// Handle failures for the processor.
func (o GetIngestProcessorUrldecodeResultOutput) OnFailures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIngestProcessorUrldecodeResult) []string { return v.OnFailures }).(pulumi.StringArrayOutput)
}

// Identifier for the processor.
func (o GetIngestProcessorUrldecodeResultOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorUrldecodeResult) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

// The field to assign the converted value to, by default `field` is updated in-place.
func (o GetIngestProcessorUrldecodeResultOutput) TargetField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorUrldecodeResult) *string { return v.TargetField }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIngestProcessorUrldecodeResultOutput{})
}
