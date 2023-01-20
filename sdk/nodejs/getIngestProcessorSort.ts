// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Sorts the elements of an array ascending or descending. Homogeneous arrays of numbers will be sorted numerically, while arrays of strings or heterogeneous arrays of strings + numbers will be sorted lexicographically. Throws an error when the field is not an array.
 *
 * See: https://www.elastic.co/guide/en/elasticsearch/reference/current/sort-processor.html
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as elasticstack from "@pulumi/elasticstack";
 *
 * const sort = elasticstack.getIngestProcessorSort({
 *     field: "array_field_to_sort",
 *     order: "desc",
 * });
 * const myIngestPipeline = new elasticstack.IngestPipeline("myIngestPipeline", {processors: [sort.then(sort => sort.json)]});
 * ```
 */
export function getIngestProcessorSort(args: GetIngestProcessorSortArgs, opts?: pulumi.InvokeOptions): Promise<GetIngestProcessorSortResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("elasticstack:index/getIngestProcessorSort:getIngestProcessorSort", {
        "description": args.description,
        "field": args.field,
        "if": args.if,
        "ignoreFailure": args.ignoreFailure,
        "onFailures": args.onFailures,
        "order": args.order,
        "tag": args.tag,
        "targetField": args.targetField,
    }, opts);
}

/**
 * A collection of arguments for invoking getIngestProcessorSort.
 */
export interface GetIngestProcessorSortArgs {
    /**
     * Description of the processor.
     */
    description?: string;
    /**
     * The field to be sorted
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
     * Handle failures for the processor.
     */
    onFailures?: string[];
    /**
     * The sort order to use. Accepts `asc` or `desc`.
     */
    order?: string;
    /**
     * Identifier for the processor.
     */
    tag?: string;
    /**
     * The field to assign the sorted value to, by default `field` is updated in-place
     */
    targetField?: string;
}

/**
 * A collection of values returned by getIngestProcessorSort.
 */
export interface GetIngestProcessorSortResult {
    /**
     * Description of the processor.
     */
    readonly description?: string;
    /**
     * The field to be sorted
     */
    readonly field: string;
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
     * JSON representation of this data source.
     */
    readonly json: string;
    /**
     * Handle failures for the processor.
     */
    readonly onFailures?: string[];
    /**
     * The sort order to use. Accepts `asc` or `desc`.
     */
    readonly order?: string;
    /**
     * Identifier for the processor.
     */
    readonly tag?: string;
    /**
     * The field to assign the sorted value to, by default `field` is updated in-place
     */
    readonly targetField?: string;
}
/**
 * Sorts the elements of an array ascending or descending. Homogeneous arrays of numbers will be sorted numerically, while arrays of strings or heterogeneous arrays of strings + numbers will be sorted lexicographically. Throws an error when the field is not an array.
 *
 * See: https://www.elastic.co/guide/en/elasticsearch/reference/current/sort-processor.html
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as elasticstack from "@pulumi/elasticstack";
 *
 * const sort = elasticstack.getIngestProcessorSort({
 *     field: "array_field_to_sort",
 *     order: "desc",
 * });
 * const myIngestPipeline = new elasticstack.IngestPipeline("myIngestPipeline", {processors: [sort.then(sort => sort.json)]});
 * ```
 */
export function getIngestProcessorSortOutput(args: GetIngestProcessorSortOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIngestProcessorSortResult> {
    return pulumi.output(args).apply((a: any) => getIngestProcessorSort(a, opts))
}

/**
 * A collection of arguments for invoking getIngestProcessorSort.
 */
export interface GetIngestProcessorSortOutputArgs {
    /**
     * Description of the processor.
     */
    description?: pulumi.Input<string>;
    /**
     * The field to be sorted
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
     * Handle failures for the processor.
     */
    onFailures?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The sort order to use. Accepts `asc` or `desc`.
     */
    order?: pulumi.Input<string>;
    /**
     * Identifier for the processor.
     */
    tag?: pulumi.Input<string>;
    /**
     * The field to assign the sorted value to, by default `field` is updated in-place
     */
    targetField?: pulumi.Input<string>;
}
