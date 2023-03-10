// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Elasticstack
{
    public static class GetIngestProcessorUrldecode
    {
        /// <summary>
        /// URL-decodes a string. If the field is an array of strings, all members of the array will be decoded.
        /// 
        /// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/urldecode-processor.html
        /// 
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Elasticstack = Pulumi.Elasticstack;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var urldecode = Elasticstack.GetIngestProcessorUrldecode.Invoke(new()
        ///     {
        ///         Field = "my_url_to_decode",
        ///     });
        /// 
        ///     var myIngestPipeline = new Elasticstack.IngestPipeline("myIngestPipeline", new()
        ///     {
        ///         Processors = new[]
        ///         {
        ///             urldecode.Apply(getIngestProcessorUrldecodeResult =&gt; getIngestProcessorUrldecodeResult.Json),
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetIngestProcessorUrldecodeResult> InvokeAsync(GetIngestProcessorUrldecodeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIngestProcessorUrldecodeResult>("elasticstack:index/getIngestProcessorUrldecode:getIngestProcessorUrldecode", args ?? new GetIngestProcessorUrldecodeArgs(), options.WithDefaults());

        /// <summary>
        /// URL-decodes a string. If the field is an array of strings, all members of the array will be decoded.
        /// 
        /// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/urldecode-processor.html
        /// 
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Elasticstack = Pulumi.Elasticstack;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var urldecode = Elasticstack.GetIngestProcessorUrldecode.Invoke(new()
        ///     {
        ///         Field = "my_url_to_decode",
        ///     });
        /// 
        ///     var myIngestPipeline = new Elasticstack.IngestPipeline("myIngestPipeline", new()
        ///     {
        ///         Processors = new[]
        ///         {
        ///             urldecode.Apply(getIngestProcessorUrldecodeResult =&gt; getIngestProcessorUrldecodeResult.Json),
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetIngestProcessorUrldecodeResult> Invoke(GetIngestProcessorUrldecodeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIngestProcessorUrldecodeResult>("elasticstack:index/getIngestProcessorUrldecode:getIngestProcessorUrldecode", args ?? new GetIngestProcessorUrldecodeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIngestProcessorUrldecodeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Description of the processor.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// The field to decode
        /// </summary>
        [Input("field", required: true)]
        public string Field { get; set; } = null!;

        /// <summary>
        /// Conditionally execute the processor
        /// </summary>
        [Input("if")]
        public string? If { get; set; }

        /// <summary>
        /// Ignore failures for the processor.
        /// </summary>
        [Input("ignoreFailure")]
        public bool? IgnoreFailure { get; set; }

        /// <summary>
        /// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
        /// </summary>
        [Input("ignoreMissing")]
        public bool? IgnoreMissing { get; set; }

        [Input("onFailures")]
        private List<string>? _onFailures;

        /// <summary>
        /// Handle failures for the processor.
        /// </summary>
        public List<string> OnFailures
        {
            get => _onFailures ?? (_onFailures = new List<string>());
            set => _onFailures = value;
        }

        /// <summary>
        /// Identifier for the processor.
        /// </summary>
        [Input("tag")]
        public string? Tag { get; set; }

        /// <summary>
        /// The field to assign the converted value to, by default `field` is updated in-place.
        /// </summary>
        [Input("targetField")]
        public string? TargetField { get; set; }

        public GetIngestProcessorUrldecodeArgs()
        {
        }
        public static new GetIngestProcessorUrldecodeArgs Empty => new GetIngestProcessorUrldecodeArgs();
    }

    public sealed class GetIngestProcessorUrldecodeInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Description of the processor.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The field to decode
        /// </summary>
        [Input("field", required: true)]
        public Input<string> Field { get; set; } = null!;

        /// <summary>
        /// Conditionally execute the processor
        /// </summary>
        [Input("if")]
        public Input<string>? If { get; set; }

        /// <summary>
        /// Ignore failures for the processor.
        /// </summary>
        [Input("ignoreFailure")]
        public Input<bool>? IgnoreFailure { get; set; }

        /// <summary>
        /// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
        /// </summary>
        [Input("ignoreMissing")]
        public Input<bool>? IgnoreMissing { get; set; }

        [Input("onFailures")]
        private InputList<string>? _onFailures;

        /// <summary>
        /// Handle failures for the processor.
        /// </summary>
        public InputList<string> OnFailures
        {
            get => _onFailures ?? (_onFailures = new InputList<string>());
            set => _onFailures = value;
        }

        /// <summary>
        /// Identifier for the processor.
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        /// <summary>
        /// The field to assign the converted value to, by default `field` is updated in-place.
        /// </summary>
        [Input("targetField")]
        public Input<string>? TargetField { get; set; }

        public GetIngestProcessorUrldecodeInvokeArgs()
        {
        }
        public static new GetIngestProcessorUrldecodeInvokeArgs Empty => new GetIngestProcessorUrldecodeInvokeArgs();
    }


    [OutputType]
    public sealed class GetIngestProcessorUrldecodeResult
    {
        /// <summary>
        /// Description of the processor.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The field to decode
        /// </summary>
        public readonly string Field;
        /// <summary>
        /// Internal identifier of the resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Conditionally execute the processor
        /// </summary>
        public readonly string? If;
        /// <summary>
        /// Ignore failures for the processor.
        /// </summary>
        public readonly bool? IgnoreFailure;
        /// <summary>
        /// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
        /// </summary>
        public readonly bool? IgnoreMissing;
        /// <summary>
        /// JSON representation of this data source.
        /// </summary>
        public readonly string Json;
        /// <summary>
        /// Handle failures for the processor.
        /// </summary>
        public readonly ImmutableArray<string> OnFailures;
        /// <summary>
        /// Identifier for the processor.
        /// </summary>
        public readonly string? Tag;
        /// <summary>
        /// The field to assign the converted value to, by default `field` is updated in-place.
        /// </summary>
        public readonly string? TargetField;

        [OutputConstructor]
        private GetIngestProcessorUrldecodeResult(
            string? description,

            string field,

            string id,

            string? @if,

            bool? ignoreFailure,

            bool? ignoreMissing,

            string json,

            ImmutableArray<string> onFailures,

            string? tag,

            string? targetField)
        {
            Description = description;
            Field = field;
            Id = id;
            If = @if;
            IgnoreFailure = ignoreFailure;
            IgnoreMissing = ignoreMissing;
            Json = json;
            OnFailures = onFailures;
            Tag = tag;
            TargetField = targetField;
        }
    }
}
