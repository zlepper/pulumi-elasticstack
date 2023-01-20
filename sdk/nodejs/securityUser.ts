// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Adds and updates users in the native realm. These users are commonly referred to as native users. See, https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-user.html
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as elasticstack from "@pulumi/elasticstack";
 *
 * const user = new elasticstack.SecurityUser("user", {
 *     username: "testuser",
 *     passwordHash: `$2a$10$rMZe6TdsUwBX/TA8vRDz0OLwKAZeCzXM4jT3tfCjpSTB8HoFuq8xO`,
 *     roles: ["kibana_user"],
 *     metadata: JSON.stringify({
 *         env: "testing",
 *         open: false,
 *         number: 49,
 *     }),
 *     elasticsearchConnection: {
 *         endpoints: ["http://localhost:9200"],
 *         username: "elastic",
 *         password: "changeme",
 *     },
 * });
 * const dev = new elasticstack.SecurityUser("dev", {
 *     username: "devuser",
 *     password: "1234567890",
 *     roles: ["kibana_user"],
 *     metadata: JSON.stringify({
 *         env: "testing",
 *         open: false,
 *         number: 49,
 *     }),
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import elasticstack:index/securityUser:SecurityUser user <cluster_uuid>/elastic
 * ```
 */
export class SecurityUser extends pulumi.CustomResource {
    /**
     * Get an existing SecurityUser resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecurityUserState, opts?: pulumi.CustomResourceOptions): SecurityUser {
        return new SecurityUser(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'elasticstack:index/securityUser:SecurityUser';

    /**
     * Returns true if the given object is an instance of SecurityUser.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurityUser {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurityUser.__pulumiType;
    }

    /**
     * Elasticsearch connection configuration block.
     */
    public readonly elasticsearchConnection!: pulumi.Output<outputs.SecurityUserElasticsearchConnection | undefined>;
    /**
     * The email of the user.
     */
    public readonly email!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether the user is enabled. The default value is true.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The full name of the user.
     */
    public readonly fullName!: pulumi.Output<string | undefined>;
    /**
     * Arbitrary metadata that you want to associate with the user.
     */
    public readonly metadata!: pulumi.Output<string>;
    /**
     * The user’s password. Passwords must be at least 6 characters long.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * A hash of the user’s password. This must be produced using the same hashing algorithm as has been configured for password storage (see https://www.elastic.co/guide/en/elasticsearch/reference/current/security-settings.html#hashing-settings).
     */
    public readonly passwordHash!: pulumi.Output<string | undefined>;
    /**
     * A set of roles the user has. The roles determine the user’s access permissions. Default is [].
     */
    public readonly roles!: pulumi.Output<string[]>;
    /**
     * An identifier for the user (see https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-user.html#security-api-put-user-path-params).
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a SecurityUser resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecurityUserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurityUserArgs | SecurityUserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecurityUserState | undefined;
            resourceInputs["elasticsearchConnection"] = state ? state.elasticsearchConnection : undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["fullName"] = state ? state.fullName : undefined;
            resourceInputs["metadata"] = state ? state.metadata : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["passwordHash"] = state ? state.passwordHash : undefined;
            resourceInputs["roles"] = state ? state.roles : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as SecurityUserArgs | undefined;
            if ((!args || args.roles === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roles'");
            }
            if ((!args || args.username === undefined) && !opts.urn) {
                throw new Error("Missing required property 'username'");
            }
            resourceInputs["elasticsearchConnection"] = args ? args.elasticsearchConnection : undefined;
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["fullName"] = args ? args.fullName : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["passwordHash"] = args?.passwordHash ? pulumi.secret(args.passwordHash) : undefined;
            resourceInputs["roles"] = args ? args.roles : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password", "passwordHash"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SecurityUser.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecurityUser resources.
 */
export interface SecurityUserState {
    /**
     * Elasticsearch connection configuration block.
     */
    elasticsearchConnection?: pulumi.Input<inputs.SecurityUserElasticsearchConnection>;
    /**
     * The email of the user.
     */
    email?: pulumi.Input<string>;
    /**
     * Specifies whether the user is enabled. The default value is true.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The full name of the user.
     */
    fullName?: pulumi.Input<string>;
    /**
     * Arbitrary metadata that you want to associate with the user.
     */
    metadata?: pulumi.Input<string>;
    /**
     * The user’s password. Passwords must be at least 6 characters long.
     */
    password?: pulumi.Input<string>;
    /**
     * A hash of the user’s password. This must be produced using the same hashing algorithm as has been configured for password storage (see https://www.elastic.co/guide/en/elasticsearch/reference/current/security-settings.html#hashing-settings).
     */
    passwordHash?: pulumi.Input<string>;
    /**
     * A set of roles the user has. The roles determine the user’s access permissions. Default is [].
     */
    roles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An identifier for the user (see https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-user.html#security-api-put-user-path-params).
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecurityUser resource.
 */
export interface SecurityUserArgs {
    /**
     * Elasticsearch connection configuration block.
     */
    elasticsearchConnection?: pulumi.Input<inputs.SecurityUserElasticsearchConnection>;
    /**
     * The email of the user.
     */
    email?: pulumi.Input<string>;
    /**
     * Specifies whether the user is enabled. The default value is true.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The full name of the user.
     */
    fullName?: pulumi.Input<string>;
    /**
     * Arbitrary metadata that you want to associate with the user.
     */
    metadata?: pulumi.Input<string>;
    /**
     * The user’s password. Passwords must be at least 6 characters long.
     */
    password?: pulumi.Input<string>;
    /**
     * A hash of the user’s password. This must be produced using the same hashing algorithm as has been configured for password storage (see https://www.elastic.co/guide/en/elasticsearch/reference/current/security-settings.html#hashing-settings).
     */
    passwordHash?: pulumi.Input<string>;
    /**
     * A set of roles the user has. The roles determine the user’s access permissions. Default is [].
     */
    roles: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An identifier for the user (see https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-user.html#security-api-put-user-path-params).
     */
    username: pulumi.Input<string>;
}
