// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Similar to the Grok Processor, dissect also extracts structured fields out of a single text field within a document. However unlike the Grok Processor, dissect does not use Regular Expressions. This allows dissect’s syntax to be simple and for some cases faster than the Grok Processor.
//
// Dissect matches a single text field against a defined pattern.
//
// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/dissect-processor.html
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/zlepper/pulumi-elasticstack/sdk/go/elasticstack"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		dissect, err := elasticstack.GetIngestProcessorDissect(ctx, &elasticstack.GetIngestProcessorDissectArgs{
// 			Field:   "message",
// 			Pattern: fmt.Sprintf("%vclientip} %vident} %vauth} [%v@timestamp}] \"%vverb} %vrequest} HTTP/%vhttpversion}\" %vstatus} %vsize}", "%{", "%{", "%{", "%{", "%{", "%{", "%{", "%{", "%{"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = elasticstack.NewIngestPipeline(ctx, "myIngestPipeline", &elasticstack.IngestPipelineArgs{
// 			Processors: pulumi.StringArray{
// 				*pulumi.String(dissect.Json),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetIngestProcessorDissect(ctx *pulumi.Context, args *GetIngestProcessorDissectArgs, opts ...pulumi.InvokeOption) (*GetIngestProcessorDissectResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetIngestProcessorDissectResult
	err := ctx.Invoke("elasticstack:index/getIngestProcessorDissect:getIngestProcessorDissect", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIngestProcessorDissect.
type GetIngestProcessorDissectArgs struct {
	// The character(s) that separate the appended fields.
	AppendSeparator *string `pulumi:"appendSeparator"`
	// Description of the processor.
	Description *string `pulumi:"description"`
	// The field to dissect.
	Field string `pulumi:"field"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
	IgnoreMissing *bool `pulumi:"ignoreMissing"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// The pattern to apply to the field.
	Pattern string `pulumi:"pattern"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
}

// A collection of values returned by getIngestProcessorDissect.
type GetIngestProcessorDissectResult struct {
	// The character(s) that separate the appended fields.
	AppendSeparator *string `pulumi:"appendSeparator"`
	// Description of the processor.
	Description *string `pulumi:"description"`
	// The field to dissect.
	Field string `pulumi:"field"`
	// Internal identifier of the resource
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
	// The pattern to apply to the field.
	Pattern string `pulumi:"pattern"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
}

func GetIngestProcessorDissectOutput(ctx *pulumi.Context, args GetIngestProcessorDissectOutputArgs, opts ...pulumi.InvokeOption) GetIngestProcessorDissectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIngestProcessorDissectResult, error) {
			args := v.(GetIngestProcessorDissectArgs)
			r, err := GetIngestProcessorDissect(ctx, &args, opts...)
			var s GetIngestProcessorDissectResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIngestProcessorDissectResultOutput)
}

// A collection of arguments for invoking getIngestProcessorDissect.
type GetIngestProcessorDissectOutputArgs struct {
	// The character(s) that separate the appended fields.
	AppendSeparator pulumi.StringPtrInput `pulumi:"appendSeparator"`
	// Description of the processor.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The field to dissect.
	Field pulumi.StringInput `pulumi:"field"`
	// Conditionally execute the processor
	If pulumi.StringPtrInput `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure pulumi.BoolPtrInput `pulumi:"ignoreFailure"`
	// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
	IgnoreMissing pulumi.BoolPtrInput `pulumi:"ignoreMissing"`
	// Handle failures for the processor.
	OnFailures pulumi.StringArrayInput `pulumi:"onFailures"`
	// The pattern to apply to the field.
	Pattern pulumi.StringInput `pulumi:"pattern"`
	// Identifier for the processor.
	Tag pulumi.StringPtrInput `pulumi:"tag"`
}

func (GetIngestProcessorDissectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIngestProcessorDissectArgs)(nil)).Elem()
}

// A collection of values returned by getIngestProcessorDissect.
type GetIngestProcessorDissectResultOutput struct{ *pulumi.OutputState }

func (GetIngestProcessorDissectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIngestProcessorDissectResult)(nil)).Elem()
}

func (o GetIngestProcessorDissectResultOutput) ToGetIngestProcessorDissectResultOutput() GetIngestProcessorDissectResultOutput {
	return o
}

func (o GetIngestProcessorDissectResultOutput) ToGetIngestProcessorDissectResultOutputWithContext(ctx context.Context) GetIngestProcessorDissectResultOutput {
	return o
}

// The character(s) that separate the appended fields.
func (o GetIngestProcessorDissectResultOutput) AppendSeparator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorDissectResult) *string { return v.AppendSeparator }).(pulumi.StringPtrOutput)
}

// Description of the processor.
func (o GetIngestProcessorDissectResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorDissectResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The field to dissect.
func (o GetIngestProcessorDissectResultOutput) Field() pulumi.StringOutput {
	return o.ApplyT(func(v GetIngestProcessorDissectResult) string { return v.Field }).(pulumi.StringOutput)
}

// Internal identifier of the resource
func (o GetIngestProcessorDissectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIngestProcessorDissectResult) string { return v.Id }).(pulumi.StringOutput)
}

// Conditionally execute the processor
func (o GetIngestProcessorDissectResultOutput) If() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorDissectResult) *string { return v.If }).(pulumi.StringPtrOutput)
}

// Ignore failures for the processor.
func (o GetIngestProcessorDissectResultOutput) IgnoreFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorDissectResult) *bool { return v.IgnoreFailure }).(pulumi.BoolPtrOutput)
}

// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
func (o GetIngestProcessorDissectResultOutput) IgnoreMissing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorDissectResult) *bool { return v.IgnoreMissing }).(pulumi.BoolPtrOutput)
}

// JSON representation of this data source.
func (o GetIngestProcessorDissectResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v GetIngestProcessorDissectResult) string { return v.Json }).(pulumi.StringOutput)
}

// Handle failures for the processor.
func (o GetIngestProcessorDissectResultOutput) OnFailures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIngestProcessorDissectResult) []string { return v.OnFailures }).(pulumi.StringArrayOutput)
}

// The pattern to apply to the field.
func (o GetIngestProcessorDissectResultOutput) Pattern() pulumi.StringOutput {
	return o.ApplyT(func(v GetIngestProcessorDissectResult) string { return v.Pattern }).(pulumi.StringOutput)
}

// Identifier for the processor.
func (o GetIngestProcessorDissectResultOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorDissectResult) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIngestProcessorDissectResultOutput{})
}
