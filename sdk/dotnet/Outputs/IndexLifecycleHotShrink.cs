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
    public sealed class IndexLifecycleHotShrink
    {
        public readonly string? MaxPrimaryShardSize;
        public readonly int? NumberOfShards;

        [OutputConstructor]
        private IndexLifecycleHotShrink(
            string? maxPrimaryShardSize,

            int? numberOfShards)
        {
            MaxPrimaryShardSize = maxPrimaryShardSize;
            NumberOfShards = numberOfShards;
        }
    }
}
