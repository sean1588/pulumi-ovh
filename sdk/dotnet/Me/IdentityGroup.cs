// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Me
{
    /// <summary>
    /// Creates an identity group.
    /// 
    /// ## Example Usage
    /// </summary>
    [OvhResourceType("ovh:Me/identityGroup:IdentityGroup")]
    public partial class IdentityGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Creation date of this group.
        /// </summary>
        [Output("creation")]
        public Output<string> Creation { get; private set; } = null!;

        /// <summary>
        /// Is the group a default and immutable one.
        /// </summary>
        [Output("defaultGroup")]
        public Output<bool> DefaultGroup { get; private set; } = null!;

        /// <summary>
        /// Group description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Date of the last update of this group.
        /// </summary>
        [Output("lastUpdate")]
        public Output<string> LastUpdate { get; private set; } = null!;

        /// <summary>
        /// Group name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Role associated with the group. Valid roles are ADMIN, REGULAR, UNPRIVILEGED, and NONE.
        /// </summary>
        [Output("role")]
        public Output<string?> Role { get; private set; } = null!;

        /// <summary>
        /// URN of the user group, used when writing IAM policies
        /// </summary>
        [Output("urn")]
        public Output<string> Urn { get; private set; } = null!;


        /// <summary>
        /// Create a IdentityGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IdentityGroup(string name, IdentityGroupArgs? args = null, CustomResourceOptions? options = null)
            : base("ovh:Me/identityGroup:IdentityGroup", name, args ?? new IdentityGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IdentityGroup(string name, Input<string> id, IdentityGroupState? state = null, CustomResourceOptions? options = null)
            : base("ovh:Me/identityGroup:IdentityGroup", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/ovh/pulumi-ovh",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IdentityGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IdentityGroup Get(string name, Input<string> id, IdentityGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new IdentityGroup(name, id, state, options);
        }
    }

    public sealed class IdentityGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Group description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Role associated with the group. Valid roles are ADMIN, REGULAR, UNPRIVILEGED, and NONE.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public IdentityGroupArgs()
        {
        }
        public static new IdentityGroupArgs Empty => new IdentityGroupArgs();
    }

    public sealed class IdentityGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Creation date of this group.
        /// </summary>
        [Input("creation")]
        public Input<string>? Creation { get; set; }

        /// <summary>
        /// Is the group a default and immutable one.
        /// </summary>
        [Input("defaultGroup")]
        public Input<bool>? DefaultGroup { get; set; }

        /// <summary>
        /// Group description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Date of the last update of this group.
        /// </summary>
        [Input("lastUpdate")]
        public Input<string>? LastUpdate { get; set; }

        /// <summary>
        /// Group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Role associated with the group. Valid roles are ADMIN, REGULAR, UNPRIVILEGED, and NONE.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// URN of the user group, used when writing IAM policies
        /// </summary>
        [Input("urn")]
        public Input<string>? Urn { get; set; }

        public IdentityGroupState()
        {
        }
        public static new IdentityGroupState Empty => new IdentityGroupState();
    }
}
