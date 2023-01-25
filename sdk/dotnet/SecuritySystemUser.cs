// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Elasticstack
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Elasticstack = Pulumi.Elasticstack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var kibanaSystem = new Elasticstack.SecuritySystemUser("kibanaSystem", new()
    ///     {
    ///         ElasticsearchConnection = new Elasticstack.Inputs.SecuritySystemUserElasticsearchConnectionArgs
    ///         {
    ///             Endpoints = new[]
    ///             {
    ///                 "http://localhost:9200",
    ///             },
    ///             Password = "changeme",
    ///             Username = "elastic",
    ///         },
    ///         PasswordHash = "$2a$10$rMZe6TdsUwBX/TA8vRDz0OLwKAZeCzXM4jT3tfCjpSTB8HoFuq8xO",
    ///         Username = "kibana_system",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [ElasticstackResourceType("elasticstack:index/securitySystemUser:SecuritySystemUser")]
    public partial class SecuritySystemUser : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Elasticsearch connection configuration block.
        /// </summary>
        [Output("elasticsearchConnection")]
        public Output<Outputs.SecuritySystemUserElasticsearchConnection?> ElasticsearchConnection { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the user is enabled. The default value is true.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// The user’s password. Passwords must be at least 6 characters long.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// A hash of the user’s password. This must be produced using the same hashing algorithm as has been configured for password storage (see https://www.elastic.co/guide/en/elasticsearch/reference/current/security-settings.html#hashing-settings).
        /// </summary>
        [Output("passwordHash")]
        public Output<string?> PasswordHash { get; private set; } = null!;

        /// <summary>
        /// An identifier for the system user (see https://www.elastic.co/guide/en/elasticsearch/reference/current/built-in-users.html).
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a SecuritySystemUser resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecuritySystemUser(string name, SecuritySystemUserArgs args, CustomResourceOptions? options = null)
            : base("elasticstack:index/securitySystemUser:SecuritySystemUser", name, args ?? new SecuritySystemUserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecuritySystemUser(string name, Input<string> id, SecuritySystemUserState? state = null, CustomResourceOptions? options = null)
            : base("elasticstack:index/securitySystemUser:SecuritySystemUser", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/zlepper",
                AdditionalSecretOutputs =
                {
                    "password",
                    "passwordHash",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SecuritySystemUser resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecuritySystemUser Get(string name, Input<string> id, SecuritySystemUserState? state = null, CustomResourceOptions? options = null)
        {
            return new SecuritySystemUser(name, id, state, options);
        }
    }

    public sealed class SecuritySystemUserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Elasticsearch connection configuration block.
        /// </summary>
        [Input("elasticsearchConnection")]
        public Input<Inputs.SecuritySystemUserElasticsearchConnectionArgs>? ElasticsearchConnection { get; set; }

        /// <summary>
        /// Specifies whether the user is enabled. The default value is true.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// The user’s password. Passwords must be at least 6 characters long.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("passwordHash")]
        private Input<string>? _passwordHash;

        /// <summary>
        /// A hash of the user’s password. This must be produced using the same hashing algorithm as has been configured for password storage (see https://www.elastic.co/guide/en/elasticsearch/reference/current/security-settings.html#hashing-settings).
        /// </summary>
        public Input<string>? PasswordHash
        {
            get => _passwordHash;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _passwordHash = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// An identifier for the system user (see https://www.elastic.co/guide/en/elasticsearch/reference/current/built-in-users.html).
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public SecuritySystemUserArgs()
        {
        }
        public static new SecuritySystemUserArgs Empty => new SecuritySystemUserArgs();
    }

    public sealed class SecuritySystemUserState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Elasticsearch connection configuration block.
        /// </summary>
        [Input("elasticsearchConnection")]
        public Input<Inputs.SecuritySystemUserElasticsearchConnectionGetArgs>? ElasticsearchConnection { get; set; }

        /// <summary>
        /// Specifies whether the user is enabled. The default value is true.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// The user’s password. Passwords must be at least 6 characters long.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("passwordHash")]
        private Input<string>? _passwordHash;

        /// <summary>
        /// A hash of the user’s password. This must be produced using the same hashing algorithm as has been configured for password storage (see https://www.elastic.co/guide/en/elasticsearch/reference/current/security-settings.html#hashing-settings).
        /// </summary>
        public Input<string>? PasswordHash
        {
            get => _passwordHash;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _passwordHash = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// An identifier for the system user (see https://www.elastic.co/guide/en/elasticsearch/reference/current/built-in-users.html).
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public SecuritySystemUserState()
        {
        }
        public static new SecuritySystemUserState Empty => new SecuritySystemUserState();
    }
}
