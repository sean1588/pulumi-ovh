// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject
{
    /// <summary>
    /// Attaches a failover IP address to a compute instance
    /// 
    /// ## Example Usage
    /// </summary>
    [OvhResourceType("ovh:CloudProject/failoverIpAttach:FailoverIpAttach")]
    public partial class FailoverIpAttach : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The IP block
        /// </summary>
        [Output("block")]
        public Output<string> Block { get; private set; } = null!;

        /// <summary>
        /// Ip continent
        /// </summary>
        [Output("continentCode")]
        public Output<string> ContinentCode { get; private set; } = null!;

        /// <summary>
        /// Ip location
        /// </summary>
        [Output("geoLoc")]
        public Output<string> GeoLoc { get; private set; } = null!;

        /// <summary>
        /// The failover ip address to attach
        /// </summary>
        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;

        /// <summary>
        /// Current operation progress in percent
        /// </summary>
        [Output("progress")]
        public Output<int> Progress { get; private set; } = null!;

        /// <summary>
        /// The GUID of an instance to which the failover IP address is be attached
        /// </summary>
        [Output("routedTo")]
        public Output<string> RoutedTo { get; private set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// Ip status, can be `ok` or `operationPending`
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// IP sub type
        /// </summary>
        [Output("subType")]
        public Output<string> SubType { get; private set; } = null!;


        /// <summary>
        /// Create a FailoverIpAttach resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FailoverIpAttach(string name, FailoverIpAttachArgs args, CustomResourceOptions? options = null)
            : base("ovh:CloudProject/failoverIpAttach:FailoverIpAttach", name, args ?? new FailoverIpAttachArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FailoverIpAttach(string name, Input<string> id, FailoverIpAttachState? state = null, CustomResourceOptions? options = null)
            : base("ovh:CloudProject/failoverIpAttach:FailoverIpAttach", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FailoverIpAttach resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FailoverIpAttach Get(string name, Input<string> id, FailoverIpAttachState? state = null, CustomResourceOptions? options = null)
        {
            return new FailoverIpAttach(name, id, state, options);
        }
    }

    public sealed class FailoverIpAttachArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IP block
        /// </summary>
        [Input("block")]
        public Input<string>? Block { get; set; }

        /// <summary>
        /// Ip continent
        /// </summary>
        [Input("continentCode")]
        public Input<string>? ContinentCode { get; set; }

        /// <summary>
        /// Ip location
        /// </summary>
        [Input("geoLoc")]
        public Input<string>? GeoLoc { get; set; }

        /// <summary>
        /// The failover ip address to attach
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// The GUID of an instance to which the failover IP address is be attached
        /// </summary>
        [Input("routedTo")]
        public Input<string>? RoutedTo { get; set; }

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public FailoverIpAttachArgs()
        {
        }
        public static new FailoverIpAttachArgs Empty => new FailoverIpAttachArgs();
    }

    public sealed class FailoverIpAttachState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IP block
        /// </summary>
        [Input("block")]
        public Input<string>? Block { get; set; }

        /// <summary>
        /// Ip continent
        /// </summary>
        [Input("continentCode")]
        public Input<string>? ContinentCode { get; set; }

        /// <summary>
        /// Ip location
        /// </summary>
        [Input("geoLoc")]
        public Input<string>? GeoLoc { get; set; }

        /// <summary>
        /// The failover ip address to attach
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Current operation progress in percent
        /// </summary>
        [Input("progress")]
        public Input<int>? Progress { get; set; }

        /// <summary>
        /// The GUID of an instance to which the failover IP address is be attached
        /// </summary>
        [Input("routedTo")]
        public Input<string>? RoutedTo { get; set; }

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Ip status, can be `ok` or `operationPending`
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// IP sub type
        /// </summary>
        [Input("subType")]
        public Input<string>? SubType { get; set; }

        public FailoverIpAttachState()
        {
        }
        public static new FailoverIpAttachState Empty => new FailoverIpAttachState();
    }
}
