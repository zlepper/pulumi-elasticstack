// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Elasticstack.Inputs
{

    public sealed class SnapshotRepositoryHdfsArgs : global::Pulumi.ResourceArgs
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
        /// Whether to load the default Hadoop configuration or not.
        /// </summary>
        [Input("loadDefaults")]
        public Input<bool>? LoadDefaults { get; set; }

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
        /// The file path within the filesystem where data is stored/loaded.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// If true, the repository is read-only.
        /// </summary>
        [Input("readonly")]
        public Input<bool>? Readonly { get; set; }

        /// <summary>
        /// The uri address for hdfs. ex: "hdfs://\n\n:\n\n/".
        /// </summary>
        [Input("uri", required: true)]
        public Input<string> Uri { get; set; } = null!;

        public SnapshotRepositoryHdfsArgs()
        {
        }
        public static new SnapshotRepositoryHdfsArgs Empty => new SnapshotRepositoryHdfsArgs();
    }
}
