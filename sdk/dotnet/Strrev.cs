// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Strrev
    {
        /// <summary>
        /// Returns the given string with all of its Unicode characters in reverse order.
        /// </summary>
        public static Task<StrrevResult> InvokeAsync(StrrevArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<StrrevResult>("std:index:strrev", args ?? new StrrevArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the given string with all of its Unicode characters in reverse order.
        /// </summary>
        public static Output<StrrevResult> Invoke(StrrevInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<StrrevResult>("std:index:strrev", args ?? new StrrevInvokeArgs(), options.WithDefaults());
    }


    public sealed class StrrevArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public string Input { get; set; } = null!;

        public StrrevArgs()
        {
        }
        public static new StrrevArgs Empty => new StrrevArgs();
    }

    public sealed class StrrevInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public Input<string> Input { get; set; } = null!;

        public StrrevInvokeArgs()
        {
        }
        public static new StrrevInvokeArgs Empty => new StrrevInvokeArgs();
    }


    [OutputType]
    public sealed class StrrevResult
    {
        public readonly string Result;

        [OutputConstructor]
        private StrrevResult(string result)
        {
            Result = result;
        }
    }
}