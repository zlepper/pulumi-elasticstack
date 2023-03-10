// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Runs an ingest processor on each element of an array or object.
 *
 * All ingest processors can run on array or object elements. However, if the number of elements is unknown, it can be cumbersome to process each one in the same way.
 *
 * The `foreach` processor lets you specify a `field` containing array or object values and a `processor` to run on each element in the field.
 *
 * See: https://www.elastic.co/guide/en/elasticsearch/reference/current/foreach-processor.html
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as elasticstack from "@pulumi/elasticstack";
 *
 * const convert = elasticstack.getIngestProcessorConvert({
 *     field: "_ingest._value",
 *     type: "integer",
 * });
 * const foreach = convert.then(convert => elasticstack.getIngestProcessorForeach({
 *     field: "values",
 *     processor: convert.json,
 * }));
 * const myIngestPipeline = new elasticstack.IngestPipeline("myIngestPipeline", {processors: [foreach.then(foreach => foreach.json)]});
 * ```
 */
export function getIngestProcessorForeach(args: GetIngestProcessorForeachArgs, opts?: pulumi.InvokeOptions): Promise<GetIngestProcessorForeachResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("elasticstack:index/getIngestProcessorForeach:getIngestProcessorForeach", {
        "description": args.description,
        "field": args.field,
        "if": args.if,
        "ignoreFailure": args.ignoreFailure,
        "ignoreMissing": args.ignoreMissing,
        "onFailures": args.onFailures,
        "processor": args.processor,
        "tag": args.tag,
    }, opts);
}

/**
 * A collection of arguments for invoking getIngestProcessorForeach.
 */
export interface GetIngestProcessorForeachArgs {
    /**
     * Description of the processor.
     */
    description?: string;
    /**
     * Field containing array or object values.
     */
    field: string;
    /**
     * Conditionally execute the processor
     */
    if?: string;
    /**
     * Ignore failures for the processor.
     */
    ignoreFailure?: boolean;
    /**
     * If `true`, the processor silently exits without changing the document if the `field` is `null` or missing.
     */
    ignoreMissing?: boolean;
    /**
     * Handle failures for the processor.
     */
    onFailures?: string[];
    /**
     * Ingest processor to run on each element.
     */
    processor: string;
    /**
     * Identifier for the processor.
     */
    tag?: string;
}

/**
 * A collection of values returned by getIngestProcessorForeach.
 */
export interface GetIngestProcessorForeachResult {
    /**
     * Description of the processor.
     */
    readonly description?: string;
    /**
     * Field containing array or object values.
     */
    readonly field: string;
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
     * If `true`, the processor silently exits without changing the document if the `field` is `null` or missing.
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
     * Ingest processor to run on each element.
     */
    readonly processor: string;
    /**
     * Identifier for the processor.
     */
    readonly tag?: string;
}
/**
 * Runs an ingest processor on each element of an array or object.
 *
 * All ingest processors can run on array or object elements. However, if the number of elements is unknown, it can be cumbersome to process each one in the same way.
 *
 * The `foreach` processor lets you specify a `field` containing array or object values and a `processor` to run on each element in the field.
 *
 * See: https://www.elastic.co/guide/en/elasticsearch/reference/current/foreach-processor.html
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as elasticstack from "@pulumi/elasticstack";
 *
 * const convert = elasticstack.getIngestProcessorConvert({
 *     field: "_ingest._value",
 *     type: "integer",
 * });
 * const foreach = convert.then(convert => elasticstack.getIngestProcessorForeach({
 *     field: "values",
 *     processor: convert.json,
 * }));
 * const myIngestPipeline = new elasticstack.IngestPipeline("myIngestPipeline", {processors: [foreach.then(foreach => foreach.json)]});
 * ```
 */
export function getIngestProcessorForeachOutput(args: GetIngestProcessorForeachOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIngestProcessorForeachResult> {
    return pulumi.output(args).apply((a: any) => getIngestProcessorForeach(a, opts))
}

/**
 * A collection of arguments for invoking getIngestProcessorForeach.
 */
export interface GetIngestProcessorForeachOutputArgs {
    /**
     * Description of the processor.
     */
    description?: pulumi.Input<string>;
    /**
     * Field containing array or object values.
     */
    field: pulumi.Input<string>;
    /**
     * Conditionally execute the processor
     */
    if?: pulumi.Input<string>;
    /**
     * Ignore failures for the processor.
     */
    ignoreFailure?: pulumi.Input<boolean>;
    /**
     * If `true`, the processor silently exits without changing the document if the `field` is `null` or missing.
     */
    ignoreMissing?: pulumi.Input<boolean>;
    /**
     * Handle failures for the processor.
     */
    onFailures?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Ingest processor to run on each element.
     */
    processor: pulumi.Input<string>;
    /**
     * Identifier for the processor.
     */
    tag?: pulumi.Input<string>;
}
