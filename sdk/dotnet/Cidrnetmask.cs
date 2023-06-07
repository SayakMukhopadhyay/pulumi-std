// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Cidrnetmask
    {
        /// <summary>
        /// Takes an IP address range in CIDR notation and returns the address-formatted subnet mask format
        /// that some systems expect for IPv4 interfaces.
        /// For example, cidrnetmask("10.0.0.0/8") returns 255.0.0.0.
        /// Not applicable to IPv6 networks since CIDR notation is the only valid notation for IPv6.
        /// </summary>
        public static Task<CidrnetmaskResult> InvokeAsync(CidrnetmaskArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<CidrnetmaskResult>("std:index:cidrnetmask", args ?? new CidrnetmaskArgs(), options.WithDefaults());

        /// <summary>
        /// Takes an IP address range in CIDR notation and returns the address-formatted subnet mask format
        /// that some systems expect for IPv4 interfaces.
        /// For example, cidrnetmask("10.0.0.0/8") returns 255.0.0.0.
        /// Not applicable to IPv6 networks since CIDR notation is the only valid notation for IPv6.
        /// </summary>
        public static Output<CidrnetmaskResult> Invoke(CidrnetmaskInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<CidrnetmaskResult>("std:index:cidrnetmask", args ?? new CidrnetmaskInvokeArgs(), options.WithDefaults());
    }


    public sealed class CidrnetmaskArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public string Input { get; set; } = null!;

        public CidrnetmaskArgs()
        {
        }
        public static new CidrnetmaskArgs Empty => new CidrnetmaskArgs();
    }

    public sealed class CidrnetmaskInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public Input<string> Input { get; set; } = null!;

        public CidrnetmaskInvokeArgs()
        {
        }
        public static new CidrnetmaskInvokeArgs Empty => new CidrnetmaskInvokeArgs();
    }


    [OutputType]
    public sealed class CidrnetmaskResult
    {
        public readonly string Result;

        [OutputConstructor]
        private CidrnetmaskResult(string result)
        {
            Result = result;
        }
    }
}
