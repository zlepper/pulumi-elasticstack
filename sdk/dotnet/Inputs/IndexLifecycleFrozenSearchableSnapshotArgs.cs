// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Elasticstack.Inputs
{

    public sealed class IndexLifecycleFrozenSearchableSnapshotArgs : global::Pulumi.ResourceArgs
    {
        [Input("forceMergeIndex")]
        public Input<bool>? ForceMergeIndex { get; set; }

        [Input("snapshotRepository", required: true)]
        public Input<string> SnapshotRepository { get; set; } = null!;

        public IndexLifecycleFrozenSearchableSnapshotArgs()
        {
        }
        public static new IndexLifecycleFrozenSearchableSnapshotArgs Empty => new IndexLifecycleFrozenSearchableSnapshotArgs();
    }
}
