// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Elasticstack
{
    public static class GetIngestProcessorDateIndexName
    {
        /// <summary>
        /// The purpose of this processor is to point documents to the right time based index based on a date or timestamp field in a document by using the date math index name support.
        /// 
        /// The processor sets the _index metadata field with a date math index name expression based on the provided index name prefix, a date or timestamp field in the documents being processed and the provided date rounding.
        /// 
        /// First, this processor fetches the date or timestamp from a field in the document being processed. Optionally, date formatting can be configured on how the field’s value should be parsed into a date. Then this date, the provided index name prefix and the provided date rounding get formatted into a date math index name expression. Also here optionally date formatting can be specified on how the date should be formatted into a date math index name expression.
        /// 
        /// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/date-index-name-processor.html
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
        ///     var dateIndexName = Elasticstack.GetIngestProcessorDateIndexName.Invoke(new()
        ///     {
        ///         Description = "monthly date-time index naming",
        ///         Field = "date1",
        ///         IndexNamePrefix = "my-index-",
        ///         DateRounding = "M",
        ///     });
        /// 
        ///     var myIngestPipeline = new Elasticstack.IngestPipeline("myIngestPipeline", new()
        ///     {
        ///         Processors = new[]
        ///         {
        ///             dateIndexName.Apply(getIngestProcessorDateIndexNameResult =&gt; getIngestProcessorDateIndexNameResult.Json),
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetIngestProcessorDateIndexNameResult> InvokeAsync(GetIngestProcessorDateIndexNameArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIngestProcessorDateIndexNameResult>("elasticstack:index/getIngestProcessorDateIndexName:getIngestProcessorDateIndexName", args ?? new GetIngestProcessorDateIndexNameArgs(), options.WithDefaults());

        /// <summary>
        /// The purpose of this processor is to point documents to the right time based index based on a date or timestamp field in a document by using the date math index name support.
        /// 
        /// The processor sets the _index metadata field with a date math index name expression based on the provided index name prefix, a date or timestamp field in the documents being processed and the provided date rounding.
        /// 
        /// First, this processor fetches the date or timestamp from a field in the document being processed. Optionally, date formatting can be configured on how the field’s value should be parsed into a date. Then this date, the provided index name prefix and the provided date rounding get formatted into a date math index name expression. Also here optionally date formatting can be specified on how the date should be formatted into a date math index name expression.
        /// 
        /// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/date-index-name-processor.html
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
        ///     var dateIndexName = Elasticstack.GetIngestProcessorDateIndexName.Invoke(new()
        ///     {
        ///         Description = "monthly date-time index naming",
        ///         Field = "date1",
        ///         IndexNamePrefix = "my-index-",
        ///         DateRounding = "M",
        ///     });
        /// 
        ///     var myIngestPipeline = new Elasticstack.IngestPipeline("myIngestPipeline", new()
        ///     {
        ///         Processors = new[]
        ///         {
        ///             dateIndexName.Apply(getIngestProcessorDateIndexNameResult =&gt; getIngestProcessorDateIndexNameResult.Json),
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetIngestProcessorDateIndexNameResult> Invoke(GetIngestProcessorDateIndexNameInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIngestProcessorDateIndexNameResult>("elasticstack:index/getIngestProcessorDateIndexName:getIngestProcessorDateIndexName", args ?? new GetIngestProcessorDateIndexNameInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIngestProcessorDateIndexNameArgs : global::Pulumi.InvokeArgs
    {
        [Input("dateFormats")]
        private List<string>? _dateFormats;

        /// <summary>
        /// An array of the expected date formats for parsing dates / timestamps in the document being preprocessed.
        /// </summary>
        public List<string> DateFormats
        {
            get => _dateFormats ?? (_dateFormats = new List<string>());
            set => _dateFormats = value;
        }

        /// <summary>
        /// How to round the date when formatting the date into the index name.
        /// </summary>
        [Input("dateRounding", required: true)]
        public string DateRounding { get; set; } = null!;

        /// <summary>
        /// Description of the processor.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// The field to get the date or timestamp from.
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
        /// The format to be used when printing the parsed date into the index name.
        /// </summary>
        [Input("indexNameFormat")]
        public string? IndexNameFormat { get; set; }

        /// <summary>
        /// A prefix of the index name to be prepended before the printed date.
        /// </summary>
        [Input("indexNamePrefix")]
        public string? IndexNamePrefix { get; set; }

        /// <summary>
        /// The locale to use when parsing the date from the document being preprocessed, relevant when parsing month names or week days.
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
        /// Identifier for the processor.
        /// </summary>
        [Input("tag")]
        public string? Tag { get; set; }

        /// <summary>
        /// The timezone to use when parsing the date and when date math index supports resolves expressions into concrete index names.
        /// </summary>
        [Input("timezone")]
        public string? Timezone { get; set; }

        public GetIngestProcessorDateIndexNameArgs()
        {
        }
        public static new GetIngestProcessorDateIndexNameArgs Empty => new GetIngestProcessorDateIndexNameArgs();
    }

    public sealed class GetIngestProcessorDateIndexNameInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("dateFormats")]
        private InputList<string>? _dateFormats;

        /// <summary>
        /// An array of the expected date formats for parsing dates / timestamps in the document being preprocessed.
        /// </summary>
        public InputList<string> DateFormats
        {
            get => _dateFormats ?? (_dateFormats = new InputList<string>());
            set => _dateFormats = value;
        }

        /// <summary>
        /// How to round the date when formatting the date into the index name.
        /// </summary>
        [Input("dateRounding", required: true)]
        public Input<string> DateRounding { get; set; } = null!;

        /// <summary>
        /// Description of the processor.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The field to get the date or timestamp from.
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
        /// The format to be used when printing the parsed date into the index name.
        /// </summary>
        [Input("indexNameFormat")]
        public Input<string>? IndexNameFormat { get; set; }

        /// <summary>
        /// A prefix of the index name to be prepended before the printed date.
        /// </summary>
        [Input("indexNamePrefix")]
        public Input<string>? IndexNamePrefix { get; set; }

        /// <summary>
        /// The locale to use when parsing the date from the document being preprocessed, relevant when parsing month names or week days.
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
        /// Identifier for the processor.
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        /// <summary>
        /// The timezone to use when parsing the date and when date math index supports resolves expressions into concrete index names.
        /// </summary>
        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        public GetIngestProcessorDateIndexNameInvokeArgs()
        {
        }
        public static new GetIngestProcessorDateIndexNameInvokeArgs Empty => new GetIngestProcessorDateIndexNameInvokeArgs();
    }


    [OutputType]
    public sealed class GetIngestProcessorDateIndexNameResult
    {
        /// <summary>
        /// An array of the expected date formats for parsing dates / timestamps in the document being preprocessed.
        /// </summary>
        public readonly ImmutableArray<string> DateFormats;
        /// <summary>
        /// How to round the date when formatting the date into the index name.
        /// </summary>
        public readonly string DateRounding;
        /// <summary>
        /// Description of the processor.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The field to get the date or timestamp from.
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
        /// The format to be used when printing the parsed date into the index name.
        /// </summary>
        public readonly string? IndexNameFormat;
        /// <summary>
        /// A prefix of the index name to be prepended before the printed date.
        /// </summary>
        public readonly string? IndexNamePrefix;
        /// <summary>
        /// JSON representation of this data source.
        /// </summary>
        public readonly string Json;
        /// <summary>
        /// The locale to use when parsing the date from the document being preprocessed, relevant when parsing month names or week days.
        /// </summary>
        public readonly string? Locale;
        /// <summary>
        /// Handle failures for the processor.
        /// </summary>
        public readonly ImmutableArray<string> OnFailures;
        /// <summary>
        /// Identifier for the processor.
        /// </summary>
        public readonly string? Tag;
        /// <summary>
        /// The timezone to use when parsing the date and when date math index supports resolves expressions into concrete index names.
        /// </summary>
        public readonly string? Timezone;

        [OutputConstructor]
        private GetIngestProcessorDateIndexNameResult(
            ImmutableArray<string> dateFormats,

            string dateRounding,

            string? description,

            string field,

            string id,

            string? @if,

            bool? ignoreFailure,

            string? indexNameFormat,

            string? indexNamePrefix,

            string json,

            string? locale,

            ImmutableArray<string> onFailures,

            string? tag,

            string? timezone)
        {
            DateFormats = dateFormats;
            DateRounding = dateRounding;
            Description = description;
            Field = field;
            Id = id;
            If = @if;
            IgnoreFailure = ignoreFailure;
            IndexNameFormat = indexNameFormat;
            IndexNamePrefix = indexNamePrefix;
            Json = json;
            Locale = locale;
            OnFailures = onFailures;
            Tag = tag;
            Timezone = timezone;
        }
    }
}
