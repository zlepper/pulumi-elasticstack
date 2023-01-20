// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Elasticstack.Inputs
{

    public sealed class SnapshotRepositoryFsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum size of files in snapshots.
        /// </summary>
        [Input("chunkSize")]
        public Input<string>? ChunkSize { get; set; }

        /// <summary>
        /// If true, metadata files, such as index mappings and settings, are compressed in snapshots.
        /// </summary>
        [Input("compress")]
        public Input<bool>? Compress { get; set; }

        /// <summary>
        /// Location of the shared filesystem used to store and retrieve snapshots.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Maximum number of snapshots the repository can contain.
        /// </summary>
        [Input("maxNumberOfSnapshots")]
        public Input<int>? MaxNumberOfSnapshots { get; set; }

        /// <summary>
        /// Maximum snapshot restore rate per node.
        /// </summary>
        [Input("maxRestoreBytesPerSec")]
        public Input<string>? MaxRestoreBytesPerSec { get; set; }

        /// <summary>
        /// Maximum snapshot creation rate per node.
        /// </summary>
        [Input("maxSnapshotBytesPerSec")]
        public Input<string>? MaxSnapshotBytesPerSec { get; set; }

        /// <summary>
        /// If true, the repository is read-only.
        /// </summary>
        [Input("readonly")]
        public Input<bool>? Readonly { get; set; }

        public SnapshotRepositoryFsGetArgs()
        {
        }
        public static new SnapshotRepositoryFsGetArgs Empty => new SnapshotRepositoryFsGetArgs();
    }
}
