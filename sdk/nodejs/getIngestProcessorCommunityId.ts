// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Helper data source to which can be used to create a processor to compute the Community ID for network flow data as defined in the [Community ID Specification](https://github.com/corelight/community-id-spec).
 * You can use a community ID to correlate network events related to a single flow.
 *
 * The community ID processor reads network flow data from related [Elastic Common Schema (ECS)](https://www.elastic.co/guide/en/ecs/1.12) fields by default. If you use the ECS, no configuration is required.
 *
 * See: https://www.elastic.co/guide/en/elasticsearch/reference/current/community-id-processor.html
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as elasticstack from "@pulumi/elasticstack";
 *
 * const community = elasticstack.getIngestProcessorCommunityId({});
 * const myIngestPipeline = new elasticstack.IngestPipeline("myIngestPipeline", {processors: [community.then(community => community.json)]});
 * ```
 */
export function getIngestProcessorCommunityId(args?: GetIngestProcessorCommunityIdArgs, opts?: pulumi.InvokeOptions): Promise<GetIngestProcessorCommunityIdResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("elasticstack:index/getIngestProcessorCommunityId:getIngestProcessorCommunityId", {
        "description": args.description,
        "destinationIp": args.destinationIp,
        "destinationPort": args.destinationPort,
        "ianaNumber": args.ianaNumber,
        "icmpCode": args.icmpCode,
        "icmpType": args.icmpType,
        "if": args.if,
        "ignoreFailure": args.ignoreFailure,
        "ignoreMissing": args.ignoreMissing,
        "onFailures": args.onFailures,
        "seed": args.seed,
        "sourceIp": args.sourceIp,
        "sourcePort": args.sourcePort,
        "tag": args.tag,
        "targetField": args.targetField,
        "transport": args.transport,
    }, opts);
}

/**
 * A collection of arguments for invoking getIngestProcessorCommunityId.
 */
export interface GetIngestProcessorCommunityIdArgs {
    /**
     * Description of the processor.
     */
    description?: string;
    /**
     * Field containing the destination IP address.
     */
    destinationIp?: string;
    /**
     * Field containing the destination port.
     */
    destinationPort?: number;
    /**
     * Field containing the IANA number.
     */
    ianaNumber?: number;
    /**
     * Field containing the ICMP code.
     */
    icmpCode?: number;
    /**
     * Field containing the ICMP type.
     */
    icmpType?: number;
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
     * Handle failures for the processor.
     */
    onFailures?: string[];
    /**
     * Seed for the community ID hash. Must be between 0 and 65535 (inclusive). The seed can prevent hash collisions between network domains, such as a staging and production network that use the same addressing scheme.
     */
    seed?: number;
    /**
     * Field containing the source IP address.
     */
    sourceIp?: string;
    /**
     * Field containing the source port.
     */
    sourcePort?: number;
    /**
     * Identifier for the processor.
     */
    tag?: string;
    /**
     * Output field for the community ID.
     */
    targetField?: string;
    /**
     * Field containing the transport protocol. Used only when the `ianaNumber` field is not present.
     */
    transport?: string;
}

/**
 * A collection of values returned by getIngestProcessorCommunityId.
 */
export interface GetIngestProcessorCommunityIdResult {
    /**
     * Description of the processor.
     */
    readonly description?: string;
    /**
     * Field containing the destination IP address.
     */
    readonly destinationIp?: string;
    /**
     * Field containing the destination port.
     */
    readonly destinationPort?: number;
    /**
     * Field containing the IANA number.
     */
    readonly ianaNumber?: number;
    /**
     * Field containing the ICMP code.
     */
    readonly icmpCode?: number;
    /**
     * Field containing the ICMP type.
     */
    readonly icmpType?: number;
    /**
     * Internal identifier of the resource
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
     * JSON representation of this data source.
     */
    readonly json: string;
    /**
     * Handle failures for the processor.
     */
    readonly onFailures?: string[];
    /**
     * Seed for the community ID hash. Must be between 0 and 65535 (inclusive). The seed can prevent hash collisions between network domains, such as a staging and production network that use the same addressing scheme.
     */
    readonly seed?: number;
    /**
     * Field containing the source IP address.
     */
    readonly sourceIp?: string;
    /**
     * Field containing the source port.
     */
    readonly sourcePort?: number;
    /**
     * Identifier for the processor.
     */
    readonly tag?: string;
    /**
     * Output field for the community ID.
     */
    readonly targetField?: string;
    /**
     * Field containing the transport protocol. Used only when the `ianaNumber` field is not present.
     */
    readonly transport?: string;
}
/**
 * Helper data source to which can be used to create a processor to compute the Community ID for network flow data as defined in the [Community ID Specification](https://github.com/corelight/community-id-spec).
 * You can use a community ID to correlate network events related to a single flow.
 *
 * The community ID processor reads network flow data from related [Elastic Common Schema (ECS)](https://www.elastic.co/guide/en/ecs/1.12) fields by default. If you use the ECS, no configuration is required.
 *
 * See: https://www.elastic.co/guide/en/elasticsearch/reference/current/community-id-processor.html
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as elasticstack from "@pulumi/elasticstack";
 *
 * const community = elasticstack.getIngestProcessorCommunityId({});
 * const myIngestPipeline = new elasticstack.IngestPipeline("myIngestPipeline", {processors: [community.then(community => community.json)]});
 * ```
 */
export function getIngestProcessorCommunityIdOutput(args?: GetIngestProcessorCommunityIdOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIngestProcessorCommunityIdResult> {
    return pulumi.output(args).apply((a: any) => getIngestProcessorCommunityId(a, opts))
}

/**
 * A collection of arguments for invoking getIngestProcessorCommunityId.
 */
export interface GetIngestProcessorCommunityIdOutputArgs {
    /**
     * Description of the processor.
     */
    description?: pulumi.Input<string>;
    /**
     * Field containing the destination IP address.
     */
    destinationIp?: pulumi.Input<string>;
    /**
     * Field containing the destination port.
     */
    destinationPort?: pulumi.Input<number>;
    /**
     * Field containing the IANA number.
     */
    ianaNumber?: pulumi.Input<number>;
    /**
     * Field containing the ICMP code.
     */
    icmpCode?: pulumi.Input<number>;
    /**
     * Field containing the ICMP type.
     */
    icmpType?: pulumi.Input<number>;
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
     * Handle failures for the processor.
     */
    onFailures?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Seed for the community ID hash. Must be between 0 and 65535 (inclusive). The seed can prevent hash collisions between network domains, such as a staging and production network that use the same addressing scheme.
     */
    seed?: pulumi.Input<number>;
    /**
     * Field containing the source IP address.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Field containing the source port.
     */
    sourcePort?: pulumi.Input<number>;
    /**
     * Identifier for the processor.
     */
    tag?: pulumi.Input<string>;
    /**
     * Output field for the community ID.
     */
    targetField?: pulumi.Input<string>;
    /**
     * Field containing the transport protocol. Used only when the `ianaNumber` field is not present.
     */
    transport?: pulumi.Input<string>;
}
