// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This processor helps automatically parse messages (or specific event fields) which are of the `foo=bar` variety.
//
// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/kv-processor.html
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
// 		kv, err := elasticstack.GetIngestProcessorKv(ctx, &elasticstack.GetIngestProcessorKvArgs{
// 			Field:      "message",
// 			FieldSplit: " ",
// 			ValueSplit: "=",
// 			ExcludeKeys: []string{
// 				"tags",
// 			},
// 			Prefix: pulumi.StringRef("setting_"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = elasticstack.NewIngestPipeline(ctx, "myIngestPipeline", &elasticstack.IngestPipelineArgs{
// 			Processors: pulumi.StringArray{
// 				*pulumi.String(kv.Json),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetIngestProcessorKv(ctx *pulumi.Context, args *GetIngestProcessorKvArgs, opts ...pulumi.InvokeOption) (*GetIngestProcessorKvResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetIngestProcessorKvResult
	err := ctx.Invoke("elasticstack:index/getIngestProcessorKv:getIngestProcessorKv", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIngestProcessorKv.
type GetIngestProcessorKvArgs struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// List of keys to exclude from document
	ExcludeKeys []string `pulumi:"excludeKeys"`
	// The field to be parsed. Supports template snippets.
	Field string `pulumi:"field"`
	// Regex pattern to use for splitting key-value pairs.
	FieldSplit string `pulumi:"fieldSplit"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
	IgnoreMissing *bool `pulumi:"ignoreMissing"`
	// List of keys to filter and insert into document. Defaults to including all keys
	IncludeKeys []string `pulumi:"includeKeys"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// Prefix to be added to extracted keys.
	Prefix *string `pulumi:"prefix"`
	// If `true` strip brackets `()`, `<>`, `[]` as well as quotes `'` and `"` from extracted values.
	StripBrackets *bool `pulumi:"stripBrackets"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
	// The field to insert the extracted keys into. Defaults to the root of the document.
	TargetField *string `pulumi:"targetField"`
	// String of characters to trim from extracted keys.
	TrimKey *string `pulumi:"trimKey"`
	// String of characters to trim from extracted values.
	TrimValue *string `pulumi:"trimValue"`
	// Regex pattern to use for splitting the key from the value within a key-value pair.
	ValueSplit string `pulumi:"valueSplit"`
}

// A collection of values returned by getIngestProcessorKv.
type GetIngestProcessorKvResult struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// List of keys to exclude from document
	ExcludeKeys []string `pulumi:"excludeKeys"`
	// The field to be parsed. Supports template snippets.
	Field string `pulumi:"field"`
	// Regex pattern to use for splitting key-value pairs.
	FieldSplit string `pulumi:"fieldSplit"`
	// Internal identifier of the resource
	Id string `pulumi:"id"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
	IgnoreMissing *bool `pulumi:"ignoreMissing"`
	// List of keys to filter and insert into document. Defaults to including all keys
	IncludeKeys []string `pulumi:"includeKeys"`
	// JSON representation of this data source.
	Json string `pulumi:"json"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// Prefix to be added to extracted keys.
	Prefix *string `pulumi:"prefix"`
	// If `true` strip brackets `()`, `<>`, `[]` as well as quotes `'` and `"` from extracted values.
	StripBrackets *bool `pulumi:"stripBrackets"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
	// The field to insert the extracted keys into. Defaults to the root of the document.
	TargetField *string `pulumi:"targetField"`
	// String of characters to trim from extracted keys.
	TrimKey *string `pulumi:"trimKey"`
	// String of characters to trim from extracted values.
	TrimValue *string `pulumi:"trimValue"`
	// Regex pattern to use for splitting the key from the value within a key-value pair.
	ValueSplit string `pulumi:"valueSplit"`
}

func GetIngestProcessorKvOutput(ctx *pulumi.Context, args GetIngestProcessorKvOutputArgs, opts ...pulumi.InvokeOption) GetIngestProcessorKvResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIngestProcessorKvResult, error) {
			args := v.(GetIngestProcessorKvArgs)
			r, err := GetIngestProcessorKv(ctx, &args, opts...)
			var s GetIngestProcessorKvResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIngestProcessorKvResultOutput)
}

// A collection of arguments for invoking getIngestProcessorKv.
type GetIngestProcessorKvOutputArgs struct {
	// Description of the processor.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// List of keys to exclude from document
	ExcludeKeys pulumi.StringArrayInput `pulumi:"excludeKeys"`
	// The field to be parsed. Supports template snippets.
	Field pulumi.StringInput `pulumi:"field"`
	// Regex pattern to use for splitting key-value pairs.
	FieldSplit pulumi.StringInput `pulumi:"fieldSplit"`
	// Conditionally execute the processor
	If pulumi.StringPtrInput `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure pulumi.BoolPtrInput `pulumi:"ignoreFailure"`
	// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
	IgnoreMissing pulumi.BoolPtrInput `pulumi:"ignoreMissing"`
	// List of keys to filter and insert into document. Defaults to including all keys
	IncludeKeys pulumi.StringArrayInput `pulumi:"includeKeys"`
	// Handle failures for the processor.
	OnFailures pulumi.StringArrayInput `pulumi:"onFailures"`
	// Prefix to be added to extracted keys.
	Prefix pulumi.StringPtrInput `pulumi:"prefix"`
	// If `true` strip brackets `()`, `<>`, `[]` as well as quotes `'` and `"` from extracted values.
	StripBrackets pulumi.BoolPtrInput `pulumi:"stripBrackets"`
	// Identifier for the processor.
	Tag pulumi.StringPtrInput `pulumi:"tag"`
	// The field to insert the extracted keys into. Defaults to the root of the document.
	TargetField pulumi.StringPtrInput `pulumi:"targetField"`
	// String of characters to trim from extracted keys.
	TrimKey pulumi.StringPtrInput `pulumi:"trimKey"`
	// String of characters to trim from extracted values.
	TrimValue pulumi.StringPtrInput `pulumi:"trimValue"`
	// Regex pattern to use for splitting the key from the value within a key-value pair.
	ValueSplit pulumi.StringInput `pulumi:"valueSplit"`
}

func (GetIngestProcessorKvOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIngestProcessorKvArgs)(nil)).Elem()
}

// A collection of values returned by getIngestProcessorKv.
type GetIngestProcessorKvResultOutput struct{ *pulumi.OutputState }

func (GetIngestProcessorKvResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIngestProcessorKvResult)(nil)).Elem()
}

func (o GetIngestProcessorKvResultOutput) ToGetIngestProcessorKvResultOutput() GetIngestProcessorKvResultOutput {
	return o
}

func (o GetIngestProcessorKvResultOutput) ToGetIngestProcessorKvResultOutputWithContext(ctx context.Context) GetIngestProcessorKvResultOutput {
	return o
}

// Description of the processor.
func (o GetIngestProcessorKvResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorKvResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// List of keys to exclude from document
func (o GetIngestProcessorKvResultOutput) ExcludeKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIngestProcessorKvResult) []string { return v.ExcludeKeys }).(pulumi.StringArrayOutput)
}

// The field to be parsed. Supports template snippets.
func (o GetIngestProcessorKvResultOutput) Field() pulumi.StringOutput {
	return o.ApplyT(func(v GetIngestProcessorKvResult) string { return v.Field }).(pulumi.StringOutput)
}

// Regex pattern to use for splitting key-value pairs.
func (o GetIngestProcessorKvResultOutput) FieldSplit() pulumi.StringOutput {
	return o.ApplyT(func(v GetIngestProcessorKvResult) string { return v.FieldSplit }).(pulumi.StringOutput)
}

// Internal identifier of the resource
func (o GetIngestProcessorKvResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIngestProcessorKvResult) string { return v.Id }).(pulumi.StringOutput)
}

// Conditionally execute the processor
func (o GetIngestProcessorKvResultOutput) If() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorKvResult) *string { return v.If }).(pulumi.StringPtrOutput)
}

// Ignore failures for the processor.
func (o GetIngestProcessorKvResultOutput) IgnoreFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorKvResult) *bool { return v.IgnoreFailure }).(pulumi.BoolPtrOutput)
}

// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
func (o GetIngestProcessorKvResultOutput) IgnoreMissing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorKvResult) *bool { return v.IgnoreMissing }).(pulumi.BoolPtrOutput)
}

// List of keys to filter and insert into document. Defaults to including all keys
func (o GetIngestProcessorKvResultOutput) IncludeKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIngestProcessorKvResult) []string { return v.IncludeKeys }).(pulumi.StringArrayOutput)
}

// JSON representation of this data source.
func (o GetIngestProcessorKvResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v GetIngestProcessorKvResult) string { return v.Json }).(pulumi.StringOutput)
}

// Handle failures for the processor.
func (o GetIngestProcessorKvResultOutput) OnFailures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIngestProcessorKvResult) []string { return v.OnFailures }).(pulumi.StringArrayOutput)
}

// Prefix to be added to extracted keys.
func (o GetIngestProcessorKvResultOutput) Prefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorKvResult) *string { return v.Prefix }).(pulumi.StringPtrOutput)
}

// If `true` strip brackets `()`, `<>`, `[]` as well as quotes `'` and `"` from extracted values.
func (o GetIngestProcessorKvResultOutput) StripBrackets() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorKvResult) *bool { return v.StripBrackets }).(pulumi.BoolPtrOutput)
}

// Identifier for the processor.
func (o GetIngestProcessorKvResultOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorKvResult) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

// The field to insert the extracted keys into. Defaults to the root of the document.
func (o GetIngestProcessorKvResultOutput) TargetField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorKvResult) *string { return v.TargetField }).(pulumi.StringPtrOutput)
}

// String of characters to trim from extracted keys.
func (o GetIngestProcessorKvResultOutput) TrimKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorKvResult) *string { return v.TrimKey }).(pulumi.StringPtrOutput)
}

// String of characters to trim from extracted values.
func (o GetIngestProcessorKvResultOutput) TrimValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorKvResult) *string { return v.TrimValue }).(pulumi.StringPtrOutput)
}

// Regex pattern to use for splitting the key from the value within a key-value pair.
func (o GetIngestProcessorKvResultOutput) ValueSplit() pulumi.StringOutput {
	return o.ApplyT(func(v GetIngestProcessorKvResult) string { return v.ValueSplit }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIngestProcessorKvResultOutput{})
}
