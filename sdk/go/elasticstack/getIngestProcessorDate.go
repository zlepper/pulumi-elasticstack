// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Helper data source to which can be used to parse dates from fields, and then uses the date or timestamp as the timestamp for the document.
// By default, the date processor adds the parsed date as a new field called `@timestamp`. You can specify a different field by setting the `targetField` configuration parameter. Multiple date formats are supported as part of the same date processor definition. They will be used sequentially to attempt parsing the date field, in the same order they were defined as part of the processor definition.
//
// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/date-processor.html
//
// ## Example Usage
//
// Here is an example that adds the parsed date to the `timestamp` field based on the `initialDate` field:
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
// 		date, err := elasticstack.GetIngestProcessorDate(ctx, &elasticstack.GetIngestProcessorDateArgs{
// 			Field:       "initial_date",
// 			TargetField: pulumi.StringRef("timestamp"),
// 			Formats: []string{
// 				"dd/MM/yyyy HH:mm:ss",
// 			},
// 			Timezone: pulumi.StringRef("Europe/Amsterdam"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = elasticstack.NewIngestPipeline(ctx, "myIngestPipeline", &elasticstack.IngestPipelineArgs{
// 			Processors: pulumi.StringArray{
// 				*pulumi.String(date.Json),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetIngestProcessorDate(ctx *pulumi.Context, args *GetIngestProcessorDateArgs, opts ...pulumi.InvokeOption) (*GetIngestProcessorDateResult, error) {
	var rv GetIngestProcessorDateResult
	err := ctx.Invoke("elasticstack:index/getIngestProcessorDate:getIngestProcessorDate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIngestProcessorDate.
type GetIngestProcessorDateArgs struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// The field to get the date from.
	Field string `pulumi:"field"`
	// An array of the expected date formats.
	Formats []string `pulumi:"formats"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// The locale to use when parsing the date, relevant when parsing month names or week days.
	Locale *string `pulumi:"locale"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// The format to use when writing the date to `targetField`.
	OutputFormat *string `pulumi:"outputFormat"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
	// The field that will hold the parsed date.
	TargetField *string `pulumi:"targetField"`
	// The timezone to use when parsing the date.
	Timezone *string `pulumi:"timezone"`
}

// A collection of values returned by getIngestProcessorDate.
type GetIngestProcessorDateResult struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// The field to get the date from.
	Field string `pulumi:"field"`
	// An array of the expected date formats.
	Formats []string `pulumi:"formats"`
	// Internal identifier of the resource
	Id string `pulumi:"id"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// JSON representation of this data source.
	Json string `pulumi:"json"`
	// The locale to use when parsing the date, relevant when parsing month names or week days.
	Locale *string `pulumi:"locale"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// The format to use when writing the date to `targetField`.
	OutputFormat *string `pulumi:"outputFormat"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
	// The field that will hold the parsed date.
	TargetField *string `pulumi:"targetField"`
	// The timezone to use when parsing the date.
	Timezone *string `pulumi:"timezone"`
}

func GetIngestProcessorDateOutput(ctx *pulumi.Context, args GetIngestProcessorDateOutputArgs, opts ...pulumi.InvokeOption) GetIngestProcessorDateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIngestProcessorDateResult, error) {
			args := v.(GetIngestProcessorDateArgs)
			r, err := GetIngestProcessorDate(ctx, &args, opts...)
			var s GetIngestProcessorDateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIngestProcessorDateResultOutput)
}

// A collection of arguments for invoking getIngestProcessorDate.
type GetIngestProcessorDateOutputArgs struct {
	// Description of the processor.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The field to get the date from.
	Field pulumi.StringInput `pulumi:"field"`
	// An array of the expected date formats.
	Formats pulumi.StringArrayInput `pulumi:"formats"`
	// Conditionally execute the processor
	If pulumi.StringPtrInput `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure pulumi.BoolPtrInput `pulumi:"ignoreFailure"`
	// The locale to use when parsing the date, relevant when parsing month names or week days.
	Locale pulumi.StringPtrInput `pulumi:"locale"`
	// Handle failures for the processor.
	OnFailures pulumi.StringArrayInput `pulumi:"onFailures"`
	// The format to use when writing the date to `targetField`.
	OutputFormat pulumi.StringPtrInput `pulumi:"outputFormat"`
	// Identifier for the processor.
	Tag pulumi.StringPtrInput `pulumi:"tag"`
	// The field that will hold the parsed date.
	TargetField pulumi.StringPtrInput `pulumi:"targetField"`
	// The timezone to use when parsing the date.
	Timezone pulumi.StringPtrInput `pulumi:"timezone"`
}

func (GetIngestProcessorDateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIngestProcessorDateArgs)(nil)).Elem()
}

// A collection of values returned by getIngestProcessorDate.
type GetIngestProcessorDateResultOutput struct{ *pulumi.OutputState }

func (GetIngestProcessorDateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIngestProcessorDateResult)(nil)).Elem()
}

func (o GetIngestProcessorDateResultOutput) ToGetIngestProcessorDateResultOutput() GetIngestProcessorDateResultOutput {
	return o
}

func (o GetIngestProcessorDateResultOutput) ToGetIngestProcessorDateResultOutputWithContext(ctx context.Context) GetIngestProcessorDateResultOutput {
	return o
}

// Description of the processor.
func (o GetIngestProcessorDateResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorDateResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The field to get the date from.
func (o GetIngestProcessorDateResultOutput) Field() pulumi.StringOutput {
	return o.ApplyT(func(v GetIngestProcessorDateResult) string { return v.Field }).(pulumi.StringOutput)
}

// An array of the expected date formats.
func (o GetIngestProcessorDateResultOutput) Formats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIngestProcessorDateResult) []string { return v.Formats }).(pulumi.StringArrayOutput)
}

// Internal identifier of the resource
func (o GetIngestProcessorDateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIngestProcessorDateResult) string { return v.Id }).(pulumi.StringOutput)
}

// Conditionally execute the processor
func (o GetIngestProcessorDateResultOutput) If() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorDateResult) *string { return v.If }).(pulumi.StringPtrOutput)
}

// Ignore failures for the processor.
func (o GetIngestProcessorDateResultOutput) IgnoreFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorDateResult) *bool { return v.IgnoreFailure }).(pulumi.BoolPtrOutput)
}

// JSON representation of this data source.
func (o GetIngestProcessorDateResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v GetIngestProcessorDateResult) string { return v.Json }).(pulumi.StringOutput)
}

// The locale to use when parsing the date, relevant when parsing month names or week days.
func (o GetIngestProcessorDateResultOutput) Locale() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorDateResult) *string { return v.Locale }).(pulumi.StringPtrOutput)
}

// Handle failures for the processor.
func (o GetIngestProcessorDateResultOutput) OnFailures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIngestProcessorDateResult) []string { return v.OnFailures }).(pulumi.StringArrayOutput)
}

// The format to use when writing the date to `targetField`.
func (o GetIngestProcessorDateResultOutput) OutputFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorDateResult) *string { return v.OutputFormat }).(pulumi.StringPtrOutput)
}

// Identifier for the processor.
func (o GetIngestProcessorDateResultOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorDateResult) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

// The field that will hold the parsed date.
func (o GetIngestProcessorDateResultOutput) TargetField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorDateResult) *string { return v.TargetField }).(pulumi.StringPtrOutput)
}

// The timezone to use when parsing the date.
func (o GetIngestProcessorDateResultOutput) Timezone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorDateResult) *string { return v.Timezone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIngestProcessorDateResultOutput{})
}