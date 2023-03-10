// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Elasticstack
{
    public static class GetIngestProcessorPipeline
    {
        /// <summary>
        /// Executes another pipeline.
        /// 
        /// The name of the current pipeline can be accessed from the `_ingest.pipeline` ingest metadata key.
        /// 
        /// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/pipeline-processor.html
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
        ///     var appendTags = Elasticstack.GetIngestProcessorAppend.Invoke(new()
        ///     {
        ///         Field = "tags",
        ///         Values = new[]
        ///         {
        ///             "production",
        ///             "{{{app}}}",
        ///             "{{{owner}}}",
        ///         },
        ///     });
        /// 
        ///     var pipelineA = new Elasticstack.IngestPipeline("pipelineA", new()
        ///     {
        ///         Processors = new[]
        ///         {
        ///             appendTags.Apply(getIngestProcessorAppendResult =&gt; getIngestProcessorAppendResult.Json),
        ///         },
        ///     });
        /// 
        ///     var fingerprint = Elasticstack.GetIngestProcessorFingerprint.Invoke(new()
        ///     {
        ///         Fields = new[]
        ///         {
        ///             "owner",
        ///         },
        ///     });
        /// 
        ///     var pipeline = Elasticstack.GetIngestProcessorPipeline.Invoke(new()
        ///     {
        ///         Name = pipelineA.Name,
        ///     });
        /// 
        ///     var pipelineB = new Elasticstack.IngestPipeline("pipelineB", new()
        ///     {
        ///         Processors = new[]
        ///         {
        ///             pipeline.Apply(getIngestProcessorPipelineResult =&gt; getIngestProcessorPipelineResult.Json),
        ///             fingerprint.Apply(getIngestProcessorFingerprintResult =&gt; getIngestProcessorFingerprintResult.Json),
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetIngestProcessorPipelineResult> InvokeAsync(GetIngestProcessorPipelineArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIngestProcessorPipelineResult>("elasticstack:index/getIngestProcessorPipeline:getIngestProcessorPipeline", args ?? new GetIngestProcessorPipelineArgs(), options.WithDefaults());

        /// <summary>
        /// Executes another pipeline.
        /// 
        /// The name of the current pipeline can be accessed from the `_ingest.pipeline` ingest metadata key.
        /// 
        /// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/pipeline-processor.html
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
        ///     var appendTags = Elasticstack.GetIngestProcessorAppend.Invoke(new()
        ///     {
        ///         Field = "tags",
        ///         Values = new[]
        ///         {
        ///             "production",
        ///             "{{{app}}}",
        ///             "{{{owner}}}",
        ///         },
        ///     });
        /// 
        ///     var pipelineA = new Elasticstack.IngestPipeline("pipelineA", new()
        ///     {
        ///         Processors = new[]
        ///         {
        ///             appendTags.Apply(getIngestProcessorAppendResult =&gt; getIngestProcessorAppendResult.Json),
        ///         },
        ///     });
        /// 
        ///     var fingerprint = Elasticstack.GetIngestProcessorFingerprint.Invoke(new()
        ///     {
        ///         Fields = new[]
        ///         {
        ///             "owner",
        ///         },
        ///     });
        /// 
        ///     var pipeline = Elasticstack.GetIngestProcessorPipeline.Invoke(new()
        ///     {
        ///         Name = pipelineA.Name,
        ///     });
        /// 
        ///     var pipelineB = new Elasticstack.IngestPipeline("pipelineB", new()
        ///     {
        ///         Processors = new[]
        ///         {
        ///             pipeline.Apply(getIngestProcessorPipelineResult =&gt; getIngestProcessorPipelineResult.Json),
        ///             fingerprint.Apply(getIngestProcessorFingerprintResult =&gt; getIngestProcessorFingerprintResult.Json),
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetIngestProcessorPipelineResult> Invoke(GetIngestProcessorPipelineInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIngestProcessorPipelineResult>("elasticstack:index/getIngestProcessorPipeline:getIngestProcessorPipeline", args ?? new GetIngestProcessorPipelineInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIngestProcessorPipelineArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Description of the processor.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

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
        /// The name of the pipeline to execute.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

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

        public GetIngestProcessorPipelineArgs()
        {
        }
        public static new GetIngestProcessorPipelineArgs Empty => new GetIngestProcessorPipelineArgs();
    }

    public sealed class GetIngestProcessorPipelineInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Description of the processor.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

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
        /// The name of the pipeline to execute.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

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

        public GetIngestProcessorPipelineInvokeArgs()
        {
        }
        public static new GetIngestProcessorPipelineInvokeArgs Empty => new GetIngestProcessorPipelineInvokeArgs();
    }


    [OutputType]
    public sealed class GetIngestProcessorPipelineResult
    {
        /// <summary>
        /// Description of the processor.
        /// </summary>
        public readonly string? Description;
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
        /// JSON representation of this data source.
        /// </summary>
        public readonly string Json;
        /// <summary>
        /// The name of the pipeline to execute.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Handle failures for the processor.
        /// </summary>
        public readonly ImmutableArray<string> OnFailures;
        /// <summary>
        /// Identifier for the processor.
        /// </summary>
        public readonly string? Tag;

        [OutputConstructor]
        private GetIngestProcessorPipelineResult(
            string? description,

            string id,

            string? @if,

            bool? ignoreFailure,

            string json,

            string name,

            ImmutableArray<string> onFailures,

            string? tag)
        {
            Description = description;
            Id = id;
            If = @if;
            IgnoreFailure = ignoreFailure;
            Json = json;
            Name = name;
            OnFailures = onFailures;
            Tag = tag;
        }
    }
}
