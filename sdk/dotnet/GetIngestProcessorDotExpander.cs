// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Elasticstack
{
    public static class GetIngestProcessorDotExpander
    {
        /// <summary>
        /// Expands a field with dots into an object field. This processor allows fields with dots in the name to be accessible by other processors in the pipeline. Otherwise these fields can’t be accessed by any processor.
        /// 
        /// See: elastic.co/guide/en/elasticsearch/reference/current/dot-expand-processor.html
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
        ///     var dotExpander = Elasticstack.GetIngestProcessorDotExpander.Invoke(new()
        ///     {
        ///         Field = "foo.bar",
        ///     });
        /// 
        ///     var myIngestPipeline = new Elasticstack.IngestPipeline("myIngestPipeline", new()
        ///     {
        ///         Processors = new[]
        ///         {
        ///             dotExpander.Apply(getIngestProcessorDotExpanderResult =&gt; getIngestProcessorDotExpanderResult.Json),
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetIngestProcessorDotExpanderResult> InvokeAsync(GetIngestProcessorDotExpanderArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIngestProcessorDotExpanderResult>("elasticstack:index/getIngestProcessorDotExpander:getIngestProcessorDotExpander", args ?? new GetIngestProcessorDotExpanderArgs(), options.WithDefaults());

        /// <summary>
        /// Expands a field with dots into an object field. This processor allows fields with dots in the name to be accessible by other processors in the pipeline. Otherwise these fields can’t be accessed by any processor.
        /// 
        /// See: elastic.co/guide/en/elasticsearch/reference/current/dot-expand-processor.html
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
        ///     var dotExpander = Elasticstack.GetIngestProcessorDotExpander.Invoke(new()
        ///     {
        ///         Field = "foo.bar",
        ///     });
        /// 
        ///     var myIngestPipeline = new Elasticstack.IngestPipeline("myIngestPipeline", new()
        ///     {
        ///         Processors = new[]
        ///         {
        ///             dotExpander.Apply(getIngestProcessorDotExpanderResult =&gt; getIngestProcessorDotExpanderResult.Json),
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetIngestProcessorDotExpanderResult> Invoke(GetIngestProcessorDotExpanderInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIngestProcessorDotExpanderResult>("elasticstack:index/getIngestProcessorDotExpander:getIngestProcessorDotExpander", args ?? new GetIngestProcessorDotExpanderInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIngestProcessorDotExpanderArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Description of the processor.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// The field to expand into an object field. If set to *, all top-level fields will be expanded.
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
        /// Controls the behavior when there is already an existing nested object that conflicts with the expanded field.
        /// </summary>
        [Input("override")]
        public bool? Override { get; set; }

        /// <summary>
        /// The field that contains the field to expand.
        /// </summary>
        [Input("path")]
        public string? Path { get; set; }

        /// <summary>
        /// Identifier for the processor.
        /// </summary>
        [Input("tag")]
        public string? Tag { get; set; }

        public GetIngestProcessorDotExpanderArgs()
        {
        }
        public static new GetIngestProcessorDotExpanderArgs Empty => new GetIngestProcessorDotExpanderArgs();
    }

    public sealed class GetIngestProcessorDotExpanderInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Description of the processor.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The field to expand into an object field. If set to *, all top-level fields will be expanded.
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
        /// Controls the behavior when there is already an existing nested object that conflicts with the expanded field.
        /// </summary>
        [Input("override")]
        public Input<bool>? Override { get; set; }

        /// <summary>
        /// The field that contains the field to expand.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Identifier for the processor.
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        public GetIngestProcessorDotExpanderInvokeArgs()
        {
        }
        public static new GetIngestProcessorDotExpanderInvokeArgs Empty => new GetIngestProcessorDotExpanderInvokeArgs();
    }


    [OutputType]
    public sealed class GetIngestProcessorDotExpanderResult
    {
        /// <summary>
        /// Description of the processor.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The field to expand into an object field. If set to *, all top-level fields will be expanded.
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
        /// JSON representation of this data source.
        /// </summary>
        public readonly string Json;
        /// <summary>
        /// Handle failures for the processor.
        /// </summary>
        public readonly ImmutableArray<string> OnFailures;
        /// <summary>
        /// Controls the behavior when there is already an existing nested object that conflicts with the expanded field.
        /// </summary>
        public readonly bool? Override;
        /// <summary>
        /// The field that contains the field to expand.
        /// </summary>
        public readonly string? Path;
        /// <summary>
        /// Identifier for the processor.
        /// </summary>
        public readonly string? Tag;

        [OutputConstructor]
        private GetIngestProcessorDotExpanderResult(
            string? description,

            string field,

            string id,

            string? @if,

            bool? ignoreFailure,

            string json,

            ImmutableArray<string> onFailures,

            bool? @override,

            string? path,

            string? tag)
        {
            Description = description;
            Field = field;
            Id = id;
            If = @if;
            IgnoreFailure = ignoreFailure;
            Json = json;
            OnFailures = onFailures;
            Override = @override;
            Path = path;
            Tag = tag;
        }
    }
}
