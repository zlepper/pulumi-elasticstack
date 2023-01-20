// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Elasticstack
{
    public static class GetIngestProcessorDate
    {
        /// <summary>
        /// Helper data source to which can be used to parse dates from fields, and then uses the date or timestamp as the timestamp for the document. 
        /// By default, the date processor adds the parsed date as a new field called `@timestamp`. You can specify a different field by setting the `target_field` configuration parameter. Multiple date formats are supported as part of the same date processor definition. They will be used sequentially to attempt parsing the date field, in the same order they were defined as part of the processor definition.
        /// 
        /// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/date-processor.html
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Here is an example that adds the parsed date to the `timestamp` field based on the `initial_date` field:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Elasticstack = Pulumi.Elasticstack;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var date = Elasticstack.GetIngestProcessorDate.Invoke(new()
        ///     {
        ///         Field = "initial_date",
        ///         TargetField = "timestamp",
        ///         Formats = new[]
        ///         {
        ///             "dd/MM/yyyy HH:mm:ss",
        ///         },
        ///         Timezone = "Europe/Amsterdam",
        ///     });
        /// 
        ///     var myIngestPipeline = new Elasticstack.IngestPipeline("myIngestPipeline", new()
        ///     {
        ///         Processors = new[]
        ///         {
        ///             date.Apply(getIngestProcessorDateResult =&gt; getIngestProcessorDateResult.Json),
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetIngestProcessorDateResult> InvokeAsync(GetIngestProcessorDateArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIngestProcessorDateResult>("elasticstack:index/getIngestProcessorDate:getIngestProcessorDate", args ?? new GetIngestProcessorDateArgs(), options.WithDefaults());

        /// <summary>
        /// Helper data source to which can be used to parse dates from fields, and then uses the date or timestamp as the timestamp for the document. 
        /// By default, the date processor adds the parsed date as a new field called `@timestamp`. You can specify a different field by setting the `target_field` configuration parameter. Multiple date formats are supported as part of the same date processor definition. They will be used sequentially to attempt parsing the date field, in the same order they were defined as part of the processor definition.
        /// 
        /// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/date-processor.html
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Here is an example that adds the parsed date to the `timestamp` field based on the `initial_date` field:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Elasticstack = Pulumi.Elasticstack;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var date = Elasticstack.GetIngestProcessorDate.Invoke(new()
        ///     {
        ///         Field = "initial_date",
        ///         TargetField = "timestamp",
        ///         Formats = new[]
        ///         {
        ///             "dd/MM/yyyy HH:mm:ss",
        ///         },
        ///         Timezone = "Europe/Amsterdam",
        ///     });
        /// 
        ///     var myIngestPipeline = new Elasticstack.IngestPipeline("myIngestPipeline", new()
        ///     {
        ///         Processors = new[]
        ///         {
        ///             date.Apply(getIngestProcessorDateResult =&gt; getIngestProcessorDateResult.Json),
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetIngestProcessorDateResult> Invoke(GetIngestProcessorDateInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIngestProcessorDateResult>("elasticstack:index/getIngestProcessorDate:getIngestProcessorDate", args ?? new GetIngestProcessorDateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIngestProcessorDateArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Description of the processor.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// The field to get the date from.
        /// </summary>
        [Input("field", required: true)]
        public string Field { get; set; } = null!;

        [Input("formats", required: true)]
        private List<string>? _formats;

        /// <summary>
        /// An array of the expected date formats.
        /// </summary>
        public List<string> Formats
        {
            get => _formats ?? (_formats = new List<string>());
            set => _formats = value;
        }

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
        /// The locale to use when parsing the date, relevant when parsing month names or week days.
        /// </summary>
        [Input("locale")]
        public string? Locale { get; set; }

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
        /// The format to use when writing the date to `target_field`.
        /// </summary>
        [Input("outputFormat")]
        public string? OutputFormat { get; set; }

        /// <summary>
        /// Identifier for the processor.
        /// </summary>
        [Input("tag")]
        public string? Tag { get; set; }

        /// <summary>
        /// The field that will hold the parsed date.
        /// </summary>
        [Input("targetField")]
        public string? TargetField { get; set; }

        /// <summary>
        /// The timezone to use when parsing the date.
        /// </summary>
        [Input("timezone")]
        public string? Timezone { get; set; }

        public GetIngestProcessorDateArgs()
        {
        }
        public static new GetIngestProcessorDateArgs Empty => new GetIngestProcessorDateArgs();
    }

    public sealed class GetIngestProcessorDateInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Description of the processor.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The field to get the date from.
        /// </summary>
        [Input("field", required: true)]
        public Input<string> Field { get; set; } = null!;

        [Input("formats", required: true)]
        private InputList<string>? _formats;

        /// <summary>
        /// An array of the expected date formats.
        /// </summary>
        public InputList<string> Formats
        {
            get => _formats ?? (_formats = new InputList<string>());
            set => _formats = value;
        }

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
        /// The locale to use when parsing the date, relevant when parsing month names or week days.
        /// </summary>
        [Input("locale")]
        public Input<string>? Locale { get; set; }

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
        /// The format to use when writing the date to `target_field`.
        /// </summary>
        [Input("outputFormat")]
        public Input<string>? OutputFormat { get; set; }

        /// <summary>
        /// Identifier for the processor.
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        /// <summary>
        /// The field that will hold the parsed date.
        /// </summary>
        [Input("targetField")]
        public Input<string>? TargetField { get; set; }

        /// <summary>
        /// The timezone to use when parsing the date.
        /// </summary>
        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        public GetIngestProcessorDateInvokeArgs()
        {
        }
        public static new GetIngestProcessorDateInvokeArgs Empty => new GetIngestProcessorDateInvokeArgs();
    }


    [OutputType]
    public sealed class GetIngestProcessorDateResult
    {
        /// <summary>
        /// Description of the processor.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The field to get the date from.
        /// </summary>
        public readonly string Field;
        /// <summary>
        /// An array of the expected date formats.
        /// </summary>
        public readonly ImmutableArray<string> Formats;
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
        /// The locale to use when parsing the date, relevant when parsing month names or week days.
        /// </summary>
        public readonly string? Locale;
        /// <summary>
        /// Handle failures for the processor.
        /// </summary>
        public readonly ImmutableArray<string> OnFailures;
        /// <summary>
        /// The format to use when writing the date to `target_field`.
        /// </summary>
        public readonly string? OutputFormat;
        /// <summary>
        /// Identifier for the processor.
        /// </summary>
        public readonly string? Tag;
        /// <summary>
        /// The field that will hold the parsed date.
        /// </summary>
        public readonly string? TargetField;
        /// <summary>
        /// The timezone to use when parsing the date.
        /// </summary>
        public readonly string? Timezone;

        [OutputConstructor]
        private GetIngestProcessorDateResult(
            string? description,

            string field,

            ImmutableArray<string> formats,

            string id,

            string? @if,

            bool? ignoreFailure,

            string json,

            string? locale,

            ImmutableArray<string> onFailures,

            string? outputFormat,

            string? tag,

            string? targetField,

            string? timezone)
        {
            Description = description;
            Field = field;
            Formats = formats;
            Id = id;
            If = @if;
            IgnoreFailure = ignoreFailure;
            Json = json;
            Locale = locale;
            OnFailures = onFailures;
            OutputFormat = outputFormat;
            Tag = tag;
            TargetField = targetField;
            Timezone = timezone;
        }
    }
}
