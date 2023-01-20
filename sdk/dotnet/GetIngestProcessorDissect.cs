// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Elasticstack
{
    public static class GetIngestProcessorDissect
    {
        /// <summary>
        /// Similar to the Grok Processor, dissect also extracts structured fields out of a single text field within a document. However unlike the Grok Processor, dissect does not use Regular Expressions. This allows dissect’s syntax to be simple and for some cases faster than the Grok Processor.
        /// 
        /// Dissect matches a single text field against a defined pattern.
        /// 
        /// 
        /// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/dissect-processor.html
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
        ///     var dissect = Elasticstack.GetIngestProcessorDissect.Invoke(new()
        ///     {
        ///         Field = "message",
        ///         Pattern = "%{clientip} %{ident} %{auth} [%{@timestamp}] \"%{verb} %{request} HTTP/%{httpversion}\" %{status} %{size}",
        ///     });
        /// 
        ///     var myIngestPipeline = new Elasticstack.IngestPipeline("myIngestPipeline", new()
        ///     {
        ///         Processors = new[]
        ///         {
        ///             dissect.Apply(getIngestProcessorDissectResult =&gt; getIngestProcessorDissectResult.Json),
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetIngestProcessorDissectResult> InvokeAsync(GetIngestProcessorDissectArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIngestProcessorDissectResult>("elasticstack:index/getIngestProcessorDissect:getIngestProcessorDissect", args ?? new GetIngestProcessorDissectArgs(), options.WithDefaults());

        /// <summary>
        /// Similar to the Grok Processor, dissect also extracts structured fields out of a single text field within a document. However unlike the Grok Processor, dissect does not use Regular Expressions. This allows dissect’s syntax to be simple and for some cases faster than the Grok Processor.
        /// 
        /// Dissect matches a single text field against a defined pattern.
        /// 
        /// 
        /// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/dissect-processor.html
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
        ///     var dissect = Elasticstack.GetIngestProcessorDissect.Invoke(new()
        ///     {
        ///         Field = "message",
        ///         Pattern = "%{clientip} %{ident} %{auth} [%{@timestamp}] \"%{verb} %{request} HTTP/%{httpversion}\" %{status} %{size}",
        ///     });
        /// 
        ///     var myIngestPipeline = new Elasticstack.IngestPipeline("myIngestPipeline", new()
        ///     {
        ///         Processors = new[]
        ///         {
        ///             dissect.Apply(getIngestProcessorDissectResult =&gt; getIngestProcessorDissectResult.Json),
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetIngestProcessorDissectResult> Invoke(GetIngestProcessorDissectInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIngestProcessorDissectResult>("elasticstack:index/getIngestProcessorDissect:getIngestProcessorDissect", args ?? new GetIngestProcessorDissectInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIngestProcessorDissectArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The character(s) that separate the appended fields.
        /// </summary>
        [Input("appendSeparator")]
        public string? AppendSeparator { get; set; }

        /// <summary>
        /// Description of the processor.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// The field to dissect.
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
        /// The pattern to apply to the field.
        /// </summary>
        [Input("pattern", required: true)]
        public string Pattern { get; set; } = null!;

        /// <summary>
        /// Identifier for the processor.
        /// </summary>
        [Input("tag")]
        public string? Tag { get; set; }

        public GetIngestProcessorDissectArgs()
        {
        }
        public static new GetIngestProcessorDissectArgs Empty => new GetIngestProcessorDissectArgs();
    }

    public sealed class GetIngestProcessorDissectInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The character(s) that separate the appended fields.
        /// </summary>
        [Input("appendSeparator")]
        public Input<string>? AppendSeparator { get; set; }

        /// <summary>
        /// Description of the processor.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The field to dissect.
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
        /// The pattern to apply to the field.
        /// </summary>
        [Input("pattern", required: true)]
        public Input<string> Pattern { get; set; } = null!;

        /// <summary>
        /// Identifier for the processor.
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        public GetIngestProcessorDissectInvokeArgs()
        {
        }
        public static new GetIngestProcessorDissectInvokeArgs Empty => new GetIngestProcessorDissectInvokeArgs();
    }


    [OutputType]
    public sealed class GetIngestProcessorDissectResult
    {
        /// <summary>
        /// The character(s) that separate the appended fields.
        /// </summary>
        public readonly string? AppendSeparator;
        /// <summary>
        /// Description of the processor.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The field to dissect.
        /// </summary>
        public readonly string Field;
        /// <summary>
        /// Internal identifier of the resource
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
        /// The pattern to apply to the field.
        /// </summary>
        public readonly string Pattern;
        /// <summary>
        /// Identifier for the processor.
        /// </summary>
        public readonly string? Tag;

        [OutputConstructor]
        private GetIngestProcessorDissectResult(
            string? appendSeparator,

            string? description,

            string field,

            string id,

            string? @if,

            bool? ignoreFailure,

            bool? ignoreMissing,

            string json,

            ImmutableArray<string> onFailures,

            string pattern,

            string? tag)
        {
            AppendSeparator = appendSeparator;
            Description = description;
            Field = field;
            Id = id;
            If = @if;
            IgnoreFailure = ignoreFailure;
            IgnoreMissing = ignoreMissing;
            Json = json;
            OnFailures = onFailures;
            Pattern = pattern;
            Tag = tag;
        }
    }
}