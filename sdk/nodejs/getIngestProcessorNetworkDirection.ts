// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Calculates the network direction given a source IP address, destination IP address, and a list of internal networks.
 *
 * The network direction processor reads IP addresses from Elastic Common Schema (ECS) fields by default. If you use the ECS, only the `internalNetworks` option must be specified.
 *
 * One of either `internalNetworks` or `internalNetworksField` must be specified. If `internalNetworksField` is specified, it follows the behavior specified by `ignoreMissing`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as elasticstack from "@pulumi/elasticstack";
 *
 * const networkDirection = elasticstack.getIngestProcessorNetworkDirection({
 *     internalNetworks: ["private"],
 * });
 * const myIngestPipeline = new elasticstack.IngestPipeline("myIngestPipeline", {processors: [networkDirection.then(networkDirection => networkDirection.json)]});
 * ```
 */
export function getIngestProcessorNetworkDirection(args?: GetIngestProcessorNetworkDirectionArgs, opts?: pulumi.InvokeOptions): Promise<GetIngestProcessorNetworkDirectionResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("elasticstack:index/getIngestProcessorNetworkDirection:getIngestProcessorNetworkDirection", {
        "description": args.description,
        "destinationIp": args.destinationIp,
        "if": args.if,
        "ignoreFailure": args.ignoreFailure,
        "ignoreMissing": args.ignoreMissing,
        "internalNetworks": args.internalNetworks,
        "internalNetworksField": args.internalNetworksField,
        "onFailures": args.onFailures,
        "sourceIp": args.sourceIp,
        "tag": args.tag,
        "targetField": args.targetField,
    }, opts);
}

/**
 * A collection of arguments for invoking getIngestProcessorNetworkDirection.
 */
export interface GetIngestProcessorNetworkDirectionArgs {
    /**
     * Description of the processor.
     */
    description?: string;
    /**
     * Field containing the destination IP address.
     */
    destinationIp?: string;
    /**
     * Conditionally execute the processor
     */
    if?: string;
    /**
     * Ignore failures for the processor.
     */
    ignoreFailure?: boolean;
    /**
     * If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
     */
    ignoreMissing?: boolean;
    /**
     * List of internal networks.
     */
    internalNetworks?: string[];
    /**
     * A field on the given document to read the internalNetworks configuration from.
     */
    internalNetworksField?: string;
    /**
     * Handle failures for the processor.
     */
    onFailures?: string[];
    /**
     * Field containing the source IP address.
     */
    sourceIp?: string;
    /**
     * Identifier for the processor.
     */
    tag?: string;
    /**
     * Output field for the network direction.
     */
    targetField?: string;
}

/**
 * A collection of values returned by getIngestProcessorNetworkDirection.
 */
export interface GetIngestProcessorNetworkDirectionResult {
    /**
     * Description of the processor.
     */
    readonly description?: string;
    /**
     * Field containing the destination IP address.
     */
    readonly destinationIp?: string;
    /**
     * Internal identifier of the resource.
     */
    readonly id: string;
    /**
     * Conditionally execute the processor
     */
    readonly if?: string;
    /**
     * Ignore failures for the processor.
     */
    readonly ignoreFailure?: boolean;
    /**
     * If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
     */
    readonly ignoreMissing?: boolean;
    /**
     * List of internal networks.
     */
    readonly internalNetworks?: string[];
    /**
     * A field on the given document to read the internalNetworks configuration from.
     */
    readonly internalNetworksField?: string;
    /**
     * JSON representation of this data source.
     */
    readonly json: string;
    /**
     * Handle failures for the processor.
     */
    readonly onFailures?: string[];
    /**
     * Field containing the source IP address.
     */
    readonly sourceIp?: string;
    /**
     * Identifier for the processor.
     */
    readonly tag?: string;
    /**
     * Output field for the network direction.
     */
    readonly targetField?: string;
}
/**
 * Calculates the network direction given a source IP address, destination IP address, and a list of internal networks.
 *
 * The network direction processor reads IP addresses from Elastic Common Schema (ECS) fields by default. If you use the ECS, only the `internalNetworks` option must be specified.
 *
 * One of either `internalNetworks` or `internalNetworksField` must be specified. If `internalNetworksField` is specified, it follows the behavior specified by `ignoreMissing`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as elasticstack from "@pulumi/elasticstack";
 *
 * const networkDirection = elasticstack.getIngestProcessorNetworkDirection({
 *     internalNetworks: ["private"],
 * });
 * const myIngestPipeline = new elasticstack.IngestPipeline("myIngestPipeline", {processors: [networkDirection.then(networkDirection => networkDirection.json)]});
 * ```
 */
export function getIngestProcessorNetworkDirectionOutput(args?: GetIngestProcessorNetworkDirectionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIngestProcessorNetworkDirectionResult> {
    return pulumi.output(args).apply((a: any) => getIngestProcessorNetworkDirection(a, opts))
}

/**
 * A collection of arguments for invoking getIngestProcessorNetworkDirection.
 */
export interface GetIngestProcessorNetworkDirectionOutputArgs {
    /**
     * Description of the processor.
     */
    description?: pulumi.Input<string>;
    /**
     * Field containing the destination IP address.
     */
    destinationIp?: pulumi.Input<string>;
    /**
     * Conditionally execute the processor
     */
    if?: pulumi.Input<string>;
    /**
     * Ignore failures for the processor.
     */
    ignoreFailure?: pulumi.Input<boolean>;
    /**
     * If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
     */
    ignoreMissing?: pulumi.Input<boolean>;
    /**
     * List of internal networks.
     */
    internalNetworks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A field on the given document to read the internalNetworks configuration from.
     */
    internalNetworksField?: pulumi.Input<string>;
    /**
     * Handle failures for the processor.
     */
    onFailures?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Field containing the source IP address.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Identifier for the processor.
     */
    tag?: pulumi.Input<string>;
    /**
     * Output field for the network direction.
     */
    targetField?: pulumi.Input<string>;
}