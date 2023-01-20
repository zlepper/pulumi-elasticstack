// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Elasticstack
{
    public static class GetIngestProcessorForeach
    {
        /// <summary>
        /// Runs an ingest processor on each element of an array or object.
        /// 
        /// All ingest processors can run on array or object elements. However, if the number of elements is unknown, it can be cumbersome to process each one in the same way.
        /// 
        /// The `foreach` processor lets you specify a `field` containing array or object values and a `processor` to run on each element in the field.
        /// 
        /// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/foreach-processor.html
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
        ///     var convert = Elasticstack.GetIngestProcessorConvert.Invoke(new()
        ///     {
        ///         Field = "_ingest._value",
        ///         Type = "integer",
        ///     });
        /// 
        ///     var @foreach = Elasticstack.GetIngestProcessorForeach.Invoke(new()
        ///     {
        ///         Field = "values",
        ///         Processor = convert.Apply(getIngestProcessorConvertResult =&gt; getIngestProcessorConvertResult.Json),
        ///     });
        /// 
        ///     var myIngestPipeline = new Elasticstack.IngestPipeline("myIngestPipeline", new()
        ///     {
        ///         Processors = new[]
        ///         {
        ///             @foreach.Apply(getIngestProcessorForeachResult =&gt; getIngestProcessorForeachResult).Apply(@foreach =&gt; @foreach.Apply(getIngestProcessorForeachResult =&gt; getIngestProcessorForeachResult.Json)),
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetIngestProcessorForeachResult> InvokeAsync(GetIngestProcessorForeachArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIngestProcessorForeachResult>("elasticstack:index/getIngestProcessorForeach:getIngestProcessorForeach", args ?? new GetIngestProcessorForeachArgs(), options.WithDefaults());

        /// <summary>
        /// Runs an ingest processor on each element of an array or object.
        /// 
        /// All ingest processors can run on array or object elements. However, if the number of elements is unknown, it can be cumbersome to process each one in the same way.
        /// 
        /// The `foreach` processor lets you specify a `field` containing array or object values and a `processor` to run on each element in the field.
        /// 
        /// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/foreach-processor.html
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
        ///     var convert = Elasticstack.GetIngestProcessorConvert.Invoke(new()
        ///     {
        ///         Field = "_ingest._value",
        ///         Type = "integer",
        ///     });
        /// 
        ///     var @foreach = Elasticstack.GetIngestProcessorForeach.Invoke(new()
        ///     {
        ///         Field = "values",
        ///         Processor = convert.Apply(getIngestProcessorConvertResult =&gt; getIngestProcessorConvertResult.Json),
        ///     });
        /// 
        ///     var myIngestPipeline = new Elasticstack.IngestPipeline("myIngestPipeline", new()
        ///     {
        ///         Processors = new[]
        ///         {
        ///             @foreach.Apply(getIngestProcessorForeachResult =&gt; getIngestProcessorForeachResult).Apply(@foreach =&gt; @foreach.Apply(getIngestProcessorForeachResult =&gt; getIngestProcessorForeachResult.Json)),
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetIngestProcessorForeachResult> Invoke(GetIngestProcessorForeachInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIngestProcessorForeachResult>("elasticstack:index/getIngestProcessorForeach:getIngestProcessorForeach", args ?? new GetIngestProcessorForeachInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIngestProcessorForeachArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Description of the processor.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// Field containing array or object values.
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
        /// If `true`, the processor silently exits without changing the document if the `field` is `null` or missing.
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
        /// Ingest processor to run on each element.
        /// </summary>
        [Input("processor", required: true)]
        public string Processor { get; set; } = null!;

        /// <summary>
        /// Identifier for the processor.
        /// </summary>
        [Input("tag")]
        public string? Tag { get; set; }

        public GetIngestProcessorForeachArgs()
        {
        }
        public static new GetIngestProcessorForeachArgs Empty => new GetIngestProcessorForeachArgs();
    }

    public sealed class GetIngestProcessorForeachInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Description of the processor.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Field containing array or object values.
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
        /// If `true`, the processor silently exits without changing the document if the `field` is `null` or missing.
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
        /// Ingest processor to run on each element.
        /// </summary>
        [Input("processor", required: true)]
        public Input<string> Processor { get; set; } = null!;

        /// <summary>
        /// Identifier for the processor.
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        public GetIngestProcessorForeachInvokeArgs()
        {
        }
        public static new GetIngestProcessorForeachInvokeArgs Empty => new GetIngestProcessorForeachInvokeArgs();
    }


    [OutputType]
    public sealed class GetIngestProcessorForeachResult
    {
        /// <summary>
        /// Description of the processor.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Field containing array or object values.
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
        /// If `true`, the processor silently exits without changing the document if the `field` is `null` or missing.
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
        /// Ingest processor to run on each element.
        /// </summary>
        public readonly string Processor;
        /// <summary>
        /// Identifier for the processor.
        /// </summary>
        public readonly string? Tag;

        [OutputConstructor]
        private GetIngestProcessorForeachResult(
            string? description,

            string field,

            string id,

            string? @if,

            bool? ignoreFailure,

            bool? ignoreMissing,

            string json,

            ImmutableArray<string> onFailures,

            string processor,

            string? tag)
        {
            Description = description;
            Field = field;
            Id = id;
            If = @if;
            IgnoreFailure = ignoreFailure;
            IgnoreMissing = ignoreMissing;
            Json = json;
            OnFailures = onFailures;
            Processor = processor;
            Tag = tag;
        }
    }
}