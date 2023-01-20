// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Drops the document without raising any errors. This is useful to prevent the document from getting indexed based on some condition.
 *
 * See: https://www.elastic.co/guide/en/elasticsearch/reference/current/drop-processor.html
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as elasticstack from "@pulumi/elasticstack";
 *
 * const drop = elasticstack.getIngestProcessorDrop({
 *     "if": "ctx.network_name == 'Guest'",
 * });
 * const myIngestPipeline = new elasticstack.IngestPipeline("myIngestPipeline", {processors: [drop.then(drop => drop.json)]});
 * ```
 */
export function getIngestProcessorDrop(args?: GetIngestProcessorDropArgs, opts?: pulumi.InvokeOptions): Promise<GetIngestProcessorDropResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("elasticstack:index/getIngestProcessorDrop:getIngestProcessorDrop", {
        "description": args.description,
        "if": args.if,
        "ignoreFailure": args.ignoreFailure,
        "onFailures": args.onFailures,
        "tag": args.tag,
    }, opts);
}

/**
 * A collection of arguments for invoking getIngestProcessorDrop.
 */
export interface GetIngestProcessorDropArgs {
    /**
     * Description of the processor.
     */
    description?: string;
    /**
     * Conditionally execute the processor
     */
    if?: string;
    /**
     * Ignore failures for the processor.
     */
    ignoreFailure?: boolean;
    /**
     * Handle failures for the processor.
     */
    onFailures?: string[];
    /**
     * Identifier for the processor.
     */
    tag?: string;
}

/**
 * A collection of values returned by getIngestProcessorDrop.
 */
export interface GetIngestProcessorDropResult {
    /**
     * Description of the processor.
     */
    readonly description?: string;
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
     * JSON representation of this data source.
     */
    readonly json: string;
    /**
     * Handle failures for the processor.
     */
    readonly onFailures?: string[];
    /**
     * Identifier for the processor.
     */
    readonly tag?: string;
}
/**
 * Drops the document without raising any errors. This is useful to prevent the document from getting indexed based on some condition.
 *
 * See: https://www.elastic.co/guide/en/elasticsearch/reference/current/drop-processor.html
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as elasticstack from "@pulumi/elasticstack";
 *
 * const drop = elasticstack.getIngestProcessorDrop({
 *     "if": "ctx.network_name == 'Guest'",
 * });
 * const myIngestPipeline = new elasticstack.IngestPipeline("myIngestPipeline", {processors: [drop.then(drop => drop.json)]});
 * ```
 */
export function getIngestProcessorDropOutput(args?: GetIngestProcessorDropOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIngestProcessorDropResult> {
    return pulumi.output(args).apply((a: any) => getIngestProcessorDrop(a, opts))
}

/**
 * A collection of arguments for invoking getIngestProcessorDrop.
 */
export interface GetIngestProcessorDropOutputArgs {
    /**
     * Description of the processor.
     */
    description?: pulumi.Input<string>;
    /**
     * Conditionally execute the processor
     */
    if?: pulumi.Input<string>;
    /**
     * Ignore failures for the processor.
     */
    ignoreFailure?: pulumi.Input<boolean>;
    /**
     * Handle failures for the processor.
     */
    onFailures?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Identifier for the processor.
     */
    tag?: pulumi.Input<string>;
}