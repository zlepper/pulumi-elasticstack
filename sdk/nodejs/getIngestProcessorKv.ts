// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This processor helps automatically parse messages (or specific event fields) which are of the `foo=bar` variety.
 *
 * See: https://www.elastic.co/guide/en/elasticsearch/reference/current/kv-processor.html
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as elasticstack from "@pulumi/elasticstack";
 *
 * const kv = elasticstack.getIngestProcessorKv({
 *     field: "message",
 *     fieldSplit: " ",
 *     valueSplit: "=",
 *     excludeKeys: ["tags"],
 *     prefix: "setting_",
 * });
 * const myIngestPipeline = new elasticstack.IngestPipeline("myIngestPipeline", {processors: [kv.then(kv => kv.json)]});
 * ```
 */
export function getIngestProcessorKv(args: GetIngestProcessorKvArgs, opts?: pulumi.InvokeOptions): Promise<GetIngestProcessorKvResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("elasticstack:index/getIngestProcessorKv:getIngestProcessorKv", {
        "description": args.description,
        "excludeKeys": args.excludeKeys,
        "field": args.field,
        "fieldSplit": args.fieldSplit,
        "if": args.if,
        "ignoreFailure": args.ignoreFailure,
        "ignoreMissing": args.ignoreMissing,
        "includeKeys": args.includeKeys,
        "onFailures": args.onFailures,
        "prefix": args.prefix,
        "stripBrackets": args.stripBrackets,
        "tag": args.tag,
        "targetField": args.targetField,
        "trimKey": args.trimKey,
        "trimValue": args.trimValue,
        "valueSplit": args.valueSplit,
    }, opts);
}

/**
 * A collection of arguments for invoking getIngestProcessorKv.
 */
export interface GetIngestProcessorKvArgs {
    /**
     * Description of the processor.
     */
    description?: string;
    /**
     * List of keys to exclude from document
     */
    excludeKeys?: string[];
    /**
     * The field to be parsed. Supports template snippets.
     */
    field: string;
    /**
     * Regex pattern to use for splitting key-value pairs.
     */
    fieldSplit: string;
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
     * List of keys to filter and insert into document. Defaults to including all keys
     */
    includeKeys?: string[];
    /**
     * Handle failures for the processor.
     */
    onFailures?: string[];
    /**
     * Prefix to be added to extracted keys.
     */
    prefix?: string;
    /**
     * If `true` strip brackets `()`, `<>`, `[]` as well as quotes `'` and `"` from extracted values.
     */
    stripBrackets?: boolean;
    /**
     * Identifier for the processor.
     */
    tag?: string;
    /**
     * The field to insert the extracted keys into. Defaults to the root of the document.
     */
    targetField?: string;
    /**
     * String of characters to trim from extracted keys.
     */
    trimKey?: string;
    /**
     * String of characters to trim from extracted values.
     */
    trimValue?: string;
    /**
     * Regex pattern to use for splitting the key from the value within a key-value pair.
     */
    valueSplit: string;
}

/**
 * A collection of values returned by getIngestProcessorKv.
 */
export interface GetIngestProcessorKvResult {
    /**
     * Description of the processor.
     */
    readonly description?: string;
    /**
     * List of keys to exclude from document
     */
    readonly excludeKeys?: string[];
    /**
     * The field to be parsed. Supports template snippets.
     */
    readonly field: string;
    /**
     * Regex pattern to use for splitting key-value pairs.
     */
    readonly fieldSplit: string;
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
     * List of keys to filter and insert into document. Defaults to including all keys
     */
    readonly includeKeys?: string[];
    /**
     * JSON representation of this data source.
     */
    readonly json: string;
    /**
     * Handle failures for the processor.
     */
    readonly onFailures?: string[];
    /**
     * Prefix to be added to extracted keys.
     */
    readonly prefix?: string;
    /**
     * If `true` strip brackets `()`, `<>`, `[]` as well as quotes `'` and `"` from extracted values.
     */
    readonly stripBrackets?: boolean;
    /**
     * Identifier for the processor.
     */
    readonly tag?: string;
    /**
     * The field to insert the extracted keys into. Defaults to the root of the document.
     */
    readonly targetField?: string;
    /**
     * String of characters to trim from extracted keys.
     */
    readonly trimKey?: string;
    /**
     * String of characters to trim from extracted values.
     */
    readonly trimValue?: string;
    /**
     * Regex pattern to use for splitting the key from the value within a key-value pair.
     */
    readonly valueSplit: string;
}
/**
 * This processor helps automatically parse messages (or specific event fields) which are of the `foo=bar` variety.
 *
 * See: https://www.elastic.co/guide/en/elasticsearch/reference/current/kv-processor.html
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as elasticstack from "@pulumi/elasticstack";
 *
 * const kv = elasticstack.getIngestProcessorKv({
 *     field: "message",
 *     fieldSplit: " ",
 *     valueSplit: "=",
 *     excludeKeys: ["tags"],
 *     prefix: "setting_",
 * });
 * const myIngestPipeline = new elasticstack.IngestPipeline("myIngestPipeline", {processors: [kv.then(kv => kv.json)]});
 * ```
 */
export function getIngestProcessorKvOutput(args: GetIngestProcessorKvOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIngestProcessorKvResult> {
    return pulumi.output(args).apply((a: any) => getIngestProcessorKv(a, opts))
}

/**
 * A collection of arguments for invoking getIngestProcessorKv.
 */
export interface GetIngestProcessorKvOutputArgs {
    /**
     * Description of the processor.
     */
    description?: pulumi.Input<string>;
    /**
     * List of keys to exclude from document
     */
    excludeKeys?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The field to be parsed. Supports template snippets.
     */
    field: pulumi.Input<string>;
    /**
     * Regex pattern to use for splitting key-value pairs.
     */
    fieldSplit: pulumi.Input<string>;
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
     * List of keys to filter and insert into document. Defaults to including all keys
     */
    includeKeys?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Handle failures for the processor.
     */
    onFailures?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Prefix to be added to extracted keys.
     */
    prefix?: pulumi.Input<string>;
    /**
     * If `true` strip brackets `()`, `<>`, `[]` as well as quotes `'` and `"` from extracted values.
     */
    stripBrackets?: pulumi.Input<boolean>;
    /**
     * Identifier for the processor.
     */
    tag?: pulumi.Input<string>;
    /**
     * The field to insert the extracted keys into. Defaults to the root of the document.
     */
    targetField?: pulumi.Input<string>;
    /**
     * String of characters to trim from extracted keys.
     */
    trimKey?: pulumi.Input<string>;
    /**
     * String of characters to trim from extracted values.
     */
    trimValue?: pulumi.Input<string>;
    /**
     * Regex pattern to use for splitting the key from the value within a key-value pair.
     */
    valueSplit: pulumi.Input<string>;
}
