// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Elasticstack.Inputs
{

    public sealed class IndexLifecycleColdSetPriorityArgs : global::Pulumi.ResourceArgs
    {
        [Input("priority", required: true)]
        public Input<int> Priority { get; set; } = null!;

        public IndexLifecycleColdSetPriorityArgs()
        {
        }
        public static new IndexLifecycleColdSetPriorityArgs Empty => new IndexLifecycleColdSetPriorityArgs();
    }
}