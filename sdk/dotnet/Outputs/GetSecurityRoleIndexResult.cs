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
    public sealed class GetSecurityRoleIndexResult
    {
        public readonly bool AllowRestrictedIndices;
        public readonly ImmutableArray<Outputs.GetSecurityRoleIndexFieldSecurityResult> FieldSecurities;
        public readonly ImmutableArray<string> Names;
        public readonly ImmutableArray<string> Privileges;
        public readonly string Query;

        [OutputConstructor]
        private GetSecurityRoleIndexResult(
            bool allowRestrictedIndices,

            ImmutableArray<Outputs.GetSecurityRoleIndexFieldSecurityResult> fieldSecurities,

            ImmutableArray<string> names,

            ImmutableArray<string> privileges,

            string query)
        {
            AllowRestrictedIndices = allowRestrictedIndices;
            FieldSecurities = fieldSecurities;
            Names = names;
            Privileges = privileges;
            Query = query;
        }
    }
}
