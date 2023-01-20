// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Elasticstack.Outputs
{

    [OutputType]
    public sealed class IndexTemplateTemplate
    {
        /// <summary>
        /// Alias to add.
        /// </summary>
        public readonly ImmutableArray<Outputs.IndexTemplateTemplateAlias> Aliases;
        /// <summary>
        /// Mapping for fields in the index.
        /// </summary>
        public readonly string? Mappings;
        /// <summary>
        /// Configuration options for the index. See, https://www.elastic.co/guide/en/elasticsearch/reference/current/index-modules.html#index-modules-settings
        /// </summary>
        public readonly string? Settings;

        [OutputConstructor]
        private IndexTemplateTemplate(
            ImmutableArray<Outputs.IndexTemplateTemplateAlias> aliases,

            string? mappings,

            string? settings)
        {
            Aliases = aliases;
            Mappings = mappings;
            Settings = settings;
        }
    }
}