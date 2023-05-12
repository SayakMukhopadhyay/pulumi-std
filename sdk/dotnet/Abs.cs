// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Abs
    {
        /// <summary>
        /// Returns the absolute value of a given float. 
        /// Example: abs(1) returns 1, and abs(-1) would also return 1, whereas abs(-3.14) would return 3.14.
        /// </summary>
        public static Task<AbsResult> InvokeAsync(AbsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<AbsResult>("std:index:abs", args ?? new AbsArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the absolute value of a given float. 
        /// Example: abs(1) returns 1, and abs(-1) would also return 1, whereas abs(-3.14) would return 3.14.
        /// </summary>
        public static Output<AbsResult> Invoke(AbsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<AbsResult>("std:index:abs", args ?? new AbsInvokeArgs(), options.WithDefaults());
    }


    public sealed class AbsArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public double Input { get; set; }

        public AbsArgs()
        {
        }
        public static new AbsArgs Empty => new AbsArgs();
    }

    public sealed class AbsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public Input<double> Input { get; set; } = null!;

        public AbsInvokeArgs()
        {
        }
        public static new AbsInvokeArgs Empty => new AbsInvokeArgs();
    }


    [OutputType]
    public sealed class AbsResult
    {
        public readonly double Result;

        [OutputConstructor]
        private AbsResult(double result)
        {
            Result = result;
        }
    }
}