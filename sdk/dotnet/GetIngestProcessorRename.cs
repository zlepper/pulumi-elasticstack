// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Elasticstack
{
    public static class GetIngestProcessorRename
    {
        /// <summary>
        /// Renames an existing field. If the field doesn’t exist or the new name is already used, an exception will be thrown.
        /// 
        /// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/rename-processor.html
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
        ///     var rename = Elasticstack.GetIngestProcessorRename.Invoke(new()
        ///     {
        ///         Field = "provider",
        ///         TargetField = "cloud.provider",
        ///     });
        /// 
        ///     var myIngestPipeline = new Elasticstack.IngestPipeline("myIngestPipeline", new()
        ///     {
        ///         Processors = new[]
        ///         {
        ///             rename.Apply(getIngestProcessorRenameResult =&gt; getIngestProcessorRenameResult.Json),
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetIngestProcessorRenameResult> InvokeAsync(GetIngestProcessorRenameArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIngestProcessorRenameResult>("elasticstack:index/getIngestProcessorRename:getIngestProcessorRename", args ?? new GetIngestProcessorRenameArgs(), options.WithDefaults());

        /// <summary>
        /// Renames an existing field. If the field doesn’t exist or the new name is already used, an exception will be thrown.
        /// 
        /// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/rename-processor.html
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
        ///     var rename = Elasticstack.GetIngestProcessorRename.Invoke(new()
        ///     {
        ///         Field = "provider",
        ///         TargetField = "cloud.provider",
        ///     });
        /// 
        ///     var myIngestPipeline = new Elasticstack.IngestPipeline("myIngestPipeline", new()
        ///     {
        ///         Processors = new[]
        ///         {
        ///             rename.Apply(getIngestProcessorRenameResult =&gt; getIngestProcessorRenameResult.Json),
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetIngestProcessorRenameResult> Invoke(GetIngestProcessorRenameInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIngestProcessorRenameResult>("elasticstack:index/getIngestProcessorRename:getIngestProcessorRename", args ?? new GetIngestProcessorRenameInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIngestProcessorRenameArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Description of the processor.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// The field to be renamed.
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
        /// The new name of the field.
        /// </summary>
        [Input("targetField", required: true)]
        public string TargetField { get; set; } = null!;

        public GetIngestProcessorRenameArgs()
        {
        }
        public static new GetIngestProcessorRenameArgs Empty => new GetIngestProcessorRenameArgs();
    }

    public sealed class GetIngestProcessorRenameInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Description of the processor.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The field to be renamed.
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
        /// The new name of the field.
        /// </summary>
        [Input("targetField", required: true)]
        public Input<string> TargetField { get; set; } = null!;

        public GetIngestProcessorRenameInvokeArgs()
        {
        }
        public static new GetIngestProcessorRenameInvokeArgs Empty => new GetIngestProcessorRenameInvokeArgs();
    }


    [OutputType]
    public sealed class GetIngestProcessorRenameResult
    {
        /// <summary>
        /// Description of the processor.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The field to be renamed.
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
        /// The new name of the field.
        /// </summary>
        public readonly string TargetField;

        [OutputConstructor]
        private GetIngestProcessorRenameResult(
            string? description,

            string field,

            string id,

            string? @if,

            bool? ignoreFailure,

            bool? ignoreMissing,

            string json,

            ImmutableArray<string> onFailures,

            string? tag,

            string targetField)
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
