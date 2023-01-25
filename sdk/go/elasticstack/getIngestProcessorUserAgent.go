// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `userAgent` processor extracts details from the user agent string a browser sends with its web requests. This processor adds this information by default under the `userAgent` field.
//
// The ingest-user-agent module ships by default with the regexes.yaml made available by uap-java with an Apache 2.0 license. For more details see https://github.com/ua-parser/uap-core.
//
// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/user-agent-processor.html
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
// 		agent, err := elasticstack.GetIngestProcessorUserAgent(ctx, &elasticstack.GetIngestProcessorUserAgentArgs{
// 			Field: "agent",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = elasticstack.NewIngestPipeline(ctx, "myIngestPipeline", &elasticstack.IngestPipelineArgs{
// 			Processors: pulumi.StringArray{
// 				*pulumi.String(agent.Json),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetIngestProcessorUserAgent(ctx *pulumi.Context, args *GetIngestProcessorUserAgentArgs, opts ...pulumi.InvokeOption) (*GetIngestProcessorUserAgentResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetIngestProcessorUserAgentResult
	err := ctx.Invoke("elasticstack:index/getIngestProcessorUserAgent:getIngestProcessorUserAgent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIngestProcessorUserAgent.
type GetIngestProcessorUserAgentArgs struct {
	// Extracts device type from the user agent string on a best-effort basis. Supported only starting from Elasticsearch version **8.0**
	ExtractDeviceType *bool `pulumi:"extractDeviceType"`
	// The field containing the user agent string.
	Field string `pulumi:"field"`
	// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
	IgnoreMissing *bool `pulumi:"ignoreMissing"`
	// Controls what properties are added to `targetField`.
	Properties []string `pulumi:"properties"`
	// The name of the file in the `config/ingest-user-agent` directory containing the regular expressions for parsing the user agent string.
	RegexFile *string `pulumi:"regexFile"`
	// The field that will be filled with the user agent details.
	TargetField *string `pulumi:"targetField"`
}

// A collection of values returned by getIngestProcessorUserAgent.
type GetIngestProcessorUserAgentResult struct {
	// Extracts device type from the user agent string on a best-effort basis. Supported only starting from Elasticsearch version **8.0**
	ExtractDeviceType *bool `pulumi:"extractDeviceType"`
	// The field containing the user agent string.
	Field string `pulumi:"field"`
	// Internal identifier of the resource.
	Id string `pulumi:"id"`
	// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
	IgnoreMissing *bool `pulumi:"ignoreMissing"`
	// JSON representation of this data source.
	Json string `pulumi:"json"`
	// Controls what properties are added to `targetField`.
	Properties []string `pulumi:"properties"`
	// The name of the file in the `config/ingest-user-agent` directory containing the regular expressions for parsing the user agent string.
	RegexFile *string `pulumi:"regexFile"`
	// The field that will be filled with the user agent details.
	TargetField *string `pulumi:"targetField"`
}

func GetIngestProcessorUserAgentOutput(ctx *pulumi.Context, args GetIngestProcessorUserAgentOutputArgs, opts ...pulumi.InvokeOption) GetIngestProcessorUserAgentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIngestProcessorUserAgentResult, error) {
			args := v.(GetIngestProcessorUserAgentArgs)
			r, err := GetIngestProcessorUserAgent(ctx, &args, opts...)
			var s GetIngestProcessorUserAgentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIngestProcessorUserAgentResultOutput)
}

// A collection of arguments for invoking getIngestProcessorUserAgent.
type GetIngestProcessorUserAgentOutputArgs struct {
	// Extracts device type from the user agent string on a best-effort basis. Supported only starting from Elasticsearch version **8.0**
	ExtractDeviceType pulumi.BoolPtrInput `pulumi:"extractDeviceType"`
	// The field containing the user agent string.
	Field pulumi.StringInput `pulumi:"field"`
	// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
	IgnoreMissing pulumi.BoolPtrInput `pulumi:"ignoreMissing"`
	// Controls what properties are added to `targetField`.
	Properties pulumi.StringArrayInput `pulumi:"properties"`
	// The name of the file in the `config/ingest-user-agent` directory containing the regular expressions for parsing the user agent string.
	RegexFile pulumi.StringPtrInput `pulumi:"regexFile"`
	// The field that will be filled with the user agent details.
	TargetField pulumi.StringPtrInput `pulumi:"targetField"`
}

func (GetIngestProcessorUserAgentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIngestProcessorUserAgentArgs)(nil)).Elem()
}

// A collection of values returned by getIngestProcessorUserAgent.
type GetIngestProcessorUserAgentResultOutput struct{ *pulumi.OutputState }

func (GetIngestProcessorUserAgentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIngestProcessorUserAgentResult)(nil)).Elem()
}

func (o GetIngestProcessorUserAgentResultOutput) ToGetIngestProcessorUserAgentResultOutput() GetIngestProcessorUserAgentResultOutput {
	return o
}

func (o GetIngestProcessorUserAgentResultOutput) ToGetIngestProcessorUserAgentResultOutputWithContext(ctx context.Context) GetIngestProcessorUserAgentResultOutput {
	return o
}

// Extracts device type from the user agent string on a best-effort basis. Supported only starting from Elasticsearch version **8.0**
func (o GetIngestProcessorUserAgentResultOutput) ExtractDeviceType() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorUserAgentResult) *bool { return v.ExtractDeviceType }).(pulumi.BoolPtrOutput)
}

// The field containing the user agent string.
func (o GetIngestProcessorUserAgentResultOutput) Field() pulumi.StringOutput {
	return o.ApplyT(func(v GetIngestProcessorUserAgentResult) string { return v.Field }).(pulumi.StringOutput)
}

// Internal identifier of the resource.
func (o GetIngestProcessorUserAgentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIngestProcessorUserAgentResult) string { return v.Id }).(pulumi.StringOutput)
}

// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
func (o GetIngestProcessorUserAgentResultOutput) IgnoreMissing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorUserAgentResult) *bool { return v.IgnoreMissing }).(pulumi.BoolPtrOutput)
}

// JSON representation of this data source.
func (o GetIngestProcessorUserAgentResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v GetIngestProcessorUserAgentResult) string { return v.Json }).(pulumi.StringOutput)
}

// Controls what properties are added to `targetField`.
func (o GetIngestProcessorUserAgentResultOutput) Properties() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIngestProcessorUserAgentResult) []string { return v.Properties }).(pulumi.StringArrayOutput)
}

// The name of the file in the `config/ingest-user-agent` directory containing the regular expressions for parsing the user agent string.
func (o GetIngestProcessorUserAgentResultOutput) RegexFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorUserAgentResult) *string { return v.RegexFile }).(pulumi.StringPtrOutput)
}

// The field that will be filled with the user agent details.
func (o GetIngestProcessorUserAgentResultOutput) TargetField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIngestProcessorUserAgentResult) *string { return v.TargetField }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIngestProcessorUserAgentResultOutput{})
}
