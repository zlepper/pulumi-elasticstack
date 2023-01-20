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
    public sealed class GetSecurityUserElasticsearchConnectionResult
    {
        /// <summary>
        /// API Key to use for authentication to Elasticsearch
        /// </summary>
        public readonly string? ApiKey;
        /// <summary>
        /// PEM-encoded custom Certificate Authority certificate
        /// </summary>
        public readonly string? CaData;
        /// <summary>
        /// Path to a custom Certificate Authority certificate
        /// </summary>
        public readonly string? CaFile;
        /// <summary>
        /// PEM encoded certificate for client auth
        /// </summary>
        public readonly string? CertData;
        /// <summary>
        /// Path to a file containing the PEM encoded certificate for client auth
        /// </summary>
        public readonly string? CertFile;
        public readonly ImmutableArray<string> Endpoints;
        /// <summary>
        /// Disable TLS certificate validation
        /// </summary>
        public readonly bool? Insecure;
        /// <summary>
        /// PEM encoded private key for client auth
        /// </summary>
        public readonly string? KeyData;
        /// <summary>
        /// Path to a file containing the PEM encoded private key for client auth
        /// </summary>
        public readonly string? KeyFile;
        /// <summary>
        /// Password to use for API authentication to Elasticsearch.
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// Username to use for API authentication to Elasticsearch.
        /// </summary>
        public readonly string? Username;

        [OutputConstructor]
        private GetSecurityUserElasticsearchConnectionResult(
            string? apiKey,

            string? caData,

            string? caFile,

            string? certData,

            string? certFile,

            ImmutableArray<string> endpoints,

            bool? insecure,

            string? keyData,

            string? keyFile,

            string? password,

            string? username)
        {
            ApiKey = apiKey;
            CaData = caData;
            CaFile = caFile;
            CertData = certData;
            CertFile = certFile;
            Endpoints = endpoints;
            Insecure = insecure;
            KeyData = keyData;
            KeyFile = keyFile;
            Password = password;
            Username = username;
        }
    }
}
