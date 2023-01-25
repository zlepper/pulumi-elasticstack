// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Converts a JSON string into a structured JSON object.
//
// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/json-processor.html
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
// 		jsonProc, err := elasticstack.GetIngestProcessorJson(ctx, &elasticstack.GetIngestProcessorJsonArgs{
// 			Field:       "string_source",
// 			TargetField: pulumi.StringRef("json_target"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = elasticstack.NewIngestPipeline(ctx, "myIngestPipeline", &elasticstack.IngestPipelineArgs{
// 			Processors: pulumi.StringArray{
// 				*pulumi.String(jsonProc.Json),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetIngestProcessorJson(ctx *pulumi.Context, args *GetIngestProcessorJsonArgs, opts ...pulumi.InvokeOption) (*GetIngestProcessorJsonResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetIngestProcessorJsonResult
	err := ctx.Invoke("elasticstack:index/getIngestProcessorJson:getIngestProcessorJson", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIngestProcessorJson.
type GetIngestProcessorJsonArgs struct {
	// Flag that forces the parsed JSON to be added at the top level of the document. `targetField` must not be set when this option is chosen.
	AddToRoot *bool `pulumi:"addToRoot"`
	// When set to `replace`, root fields that conflict with fields from the parsed JSON will be overridden. When set to `merge`, conflicting fields will be merged. Only applicable if `addToRoot` is set to `true`.
	AddToRootConflictStrategy *string `pulumi:"addToRootConflictStrategy"`
	// When set to `true`, the JSON parser will not fail if the JSON contains duplicate keys. Instead, the last encountered value for any duplicate key wins.
	AllowDuplicateKeys *bool `pulumi:"allowDuplicateKeys"`
	// Description of the processor.
	Description *string `pulumi:"description"`
	// The field to be parsed.
	Field string `pulumi:"field"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
	// The field that the converted structured object will be written into. Any existing content in this field will be overwritten.
	TargetField *string `pulumi:"targetField"`
}

// A collection of values returned by getIngestProcessorJson.
type GetIngestProcessorJsonResult struct {
	// Flag that forces the parsed JSON to be added at the top level of the document. `targetField` must not be set when this option is chosen.
	AddToRoot *bool `pulumi:"addToRoot"`
	// When set to `replace`, root fields that conflict with fields from the parsed JSON will be overridden. When set to `merge`, conflicting fields will be merged. Only applicable if `addToRoot` is set to `true`.
	AddToRootConflictStrategy *string `pulumi:"addToRootConflictStrategy"`
	// When set to `true`, the JSON parser will not fail if the JSON contains duplicate keys. Instead, the last encountered value for any duplicate key wins.
	AllowDuplicateKeys *bool `pulumi:"allowDuplicateKeys"`
	// Description of the processor.
	Description *string `pulumi:"description"`
	// The field to be parsed.
	Field string `pulumi:"field"`
	// Internal identifier of the resource.
	Id string `pulumi:"id"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// JSON representation of this data source.
	Json string `pulumi:"json"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
	// The field that the converted structured object will be written into. Any existing content in this field will be overwritten.
	TargetField *string `pulumi:"targetField"`
}

func GetIngestProcessorJsonOutput(ctx *pulumi.Context, args GetIngestProcessorJsonOutputArgs, opts ...pulumi.InvokeOption) GetIngestProcessorJsonResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIngestProcessorJsonResult, error) {
			args := v.(GetIngestProcessorJsonArgs)
			r, err := GetIngestProcessorJson(ctx, &args, opts...)
			var s GetIngestProcessorJsonResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIngestProcessorJsonResultOutput)
}

// A collection of arguments for invoking getIngestProcessorJson.
type GetIngestProcessorJsonOutputArgs struct {
	// Flag that forces the parsed JSON to be added at the top level of the document. `targetField` must not be set when this option is chosen.
	AddToRoot pulumi.BoolPtrInput `pulumi:"addToRoot"`
	// When set to `replace`, root fields that conflict with fields from the parsed JSON will be overridden. When set to `merge`, conflicting fields will be merged. Only applicable if `addToRoot` is set to `true`.
	AddToRootConflictStrategy pulumi.StringPtrInput `pulumi:"addToRootConflictStrategy"`
	// When set to `true`, the JSON parser will not fail if the JSON contains duplicate keys. Instead, the last encountered value for any duplicate key wins.
	AllowDuplicateKeys pulumi.BoolPtrInput `pulumi:"allowDuplicateKeys"`
	// Description of the processor.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The field to be parsed.
	Field pulumi.StringInput `pulumi:"field"`
	// Conditionally execute the processor
	If pulumi.StringPtrInput `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure pulumi.BoolPtrInput `pulumi:"ignoreFailure"`
	// Handle failures for the processor.
	OnFailures pulumi.StringArrayInput `pulumi:"onFailures"`
	// Identifier for the processor.
	Tag pulumi.StringPtrInput `pulumi:"tag"`
	// The field that the converted structured object will be written into. Any existing content in this field will be overwritten.
	TargetField pulumi.StringPtrInput `pulumi:"targetField"`
}

func (GetIngestProcessorJsonOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIngestProcessorJsonArgs)(nil)).Elem()
}

// A collection of values returned by getIngestProcessorJson.
type GetIngestProcessorJsonResultOutput struct{ *pulumi.OutputState }

func (GetIngestProcessorJsonResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIngestProcessorJsonResult)(nil)).Elem()
}

func (o GetIngestProcessorJsonResultOutput) ToGetIngestProcessorJsonResultOutput() GetIngestProcessorJsonResultOutput {
	return o
}

func (o GetIngestProcessorJsonResultOutput) ToGetIngestProcessorJsonResultOutputWithContext(ctx context.Context) GetIngestProcessorJsonResultOutput {
	return o
}

// Flag that forces the parsed JSON to be added at the top level of the document. `targetField` must not be set when this option is chosen.
func (o GetIngestProcessorJsonResultOutput) AddToRoot() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorJsonResult) *bool { return v.AddToRoot }).(pulumi.BoolPtrOutput)
}

// When set to `replace`, root fields that conflict with fields from the parsed JSON will be overridden. When set to `merge`, conflicting fields will be merged. Only applicable if `addToRoot` is set to `true`.
func (o GetIngestProcessorJsonResultOutput) AddToRootConflictStrategy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorJsonResult) *string { return v.AddToRootConflictStrategy }).(pulumi.StringPtrOutput)
}

// When set to `true`, the JSON parser will not fail if the JSON contains duplicate keys. Instead, the last encountered value for any duplicate key wins.
func (o GetIngestProcessorJsonResultOutput) AllowDuplicateKeys() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorJsonResult) *bool { return v.AllowDuplicateKeys }).(pulumi.BoolPtrOutput)
}

// Description of the processor.
func (o GetIngestProcessorJsonResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorJsonResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The field to be parsed.
func (o GetIngestProcessorJsonResultOutput) Field() pulumi.StringOutput {
	return o.ApplyT(func(v GetIngestProcessorJsonResult) string { return v.Field }).(pulumi.StringOutput)
}

// Internal identifier of the resource.
func (o GetIngestProcessorJsonResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIngestProcessorJsonResult) string { return v.Id }).(pulumi.StringOutput)
}

// Conditionally execute the processor
func (o GetIngestProcessorJsonResultOutput) If() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorJsonResult) *string { return v.If }).(pulumi.StringPtrOutput)
}

// Ignore failures for the processor.
func (o GetIngestProcessorJsonResultOutput) IgnoreFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorJsonResult) *bool { return v.IgnoreFailure }).(pulumi.BoolPtrOutput)
}

// JSON representation of this data source.
func (o GetIngestProcessorJsonResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v GetIngestProcessorJsonResult) string { return v.Json }).(pulumi.StringOutput)
}

// Handle failures for the processor.
func (o GetIngestProcessorJsonResultOutput) OnFailures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIngestProcessorJsonResult) []string { return v.OnFailures }).(pulumi.StringArrayOutput)
}

// Identifier for the processor.
func (o GetIngestProcessorJsonResultOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorJsonResult) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

// The field that the converted structured object will be written into. Any existing content in this field will be overwritten.
func (o GetIngestProcessorJsonResultOutput) TargetField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorJsonResult) *string { return v.TargetField }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIngestProcessorJsonResultOutput{})
}
