// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Elasticstack.Inputs
{

    public sealed class SnapshotRepositoryS3GetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the path to the repository data within its bucket.
        /// </summary>
        [Input("basePath")]
        public Input<string>? BasePath { get; set; }

        /// <summary>
        /// Name of the S3 bucket to use for snapshots.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// Minimum threshold below which the chunk is uploaded using a single request.
        /// </summary>
        [Input("bufferSize")]
        public Input<string>? BufferSize { get; set; }

        /// <summary>
        /// The S3 repository supports all S3 canned ACLs.
        /// </summary>
        [Input("cannedAcl")]
        public Input<string>? CannedAcl { get; set; }

        /// <summary>
        /// Maximum size of files in snapshots.
        /// </summary>
        [Input("chunkSize")]
        public Input<string>? ChunkSize { get; set; }

        /// <summary>
        /// The name of the S3 client to use to connect to S3.
        /// </summary>
        [Input("client")]
        public Input<string>? Client { get; set; }

        /// <summary>
        /// If true, metadata files, such as index mappings and settings, are compressed in snapshots.
        /// </summary>
        [Input("compress")]
        public Input<bool>? Compress { get; set; }

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

        /// <summary>
        /// When true, files are encrypted server-side using AES-256 algorithm.
        /// </summary>
        [Input("serverSideEncryption")]
        public Input<bool>? ServerSideEncryption { get; set; }

        /// <summary>
        /// Sets the S3 storage class for objects stored in the snapshot repository.
        /// </summary>
        [Input("storageClass")]
        public Input<string>? StorageClass { get; set; }

        public SnapshotRepositoryS3GetArgs()
        {
        }
        public static new SnapshotRepositoryS3GetArgs Empty => new SnapshotRepositoryS3GetArgs();
    }
}
