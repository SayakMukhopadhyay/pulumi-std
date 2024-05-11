// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Unixtorfc3999
    {
        /// <summary>
        /// Converts a RFC3999 formatted timestamp into a Unix timestamp with milliseconds.
        /// </summary>
        public static Task<Unixtorfc3999Result> InvokeAsync(Unixtorfc3999Args args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Unixtorfc3999Result>("std:index:unixtorfc3999", args ?? new Unixtorfc3999Args(), options.WithDefaults());

        /// <summary>
        /// Converts a RFC3999 formatted timestamp into a Unix timestamp with milliseconds.
        /// </summary>
        public static Output<Unixtorfc3999Result> Invoke(Unixtorfc3999InvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Unixtorfc3999Result>("std:index:unixtorfc3999", args ?? new Unixtorfc3999InvokeArgs(), options.WithDefaults());
    }


    public sealed class Unixtorfc3999Args : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public string Input { get; set; } = null!;

        public Unixtorfc3999Args()
        {
        }
        public static new Unixtorfc3999Args Empty => new Unixtorfc3999Args();
    }

    public sealed class Unixtorfc3999InvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public Input<string> Input { get; set; } = null!;

        public Unixtorfc3999InvokeArgs()
        {
        }
        public static new Unixtorfc3999InvokeArgs Empty => new Unixtorfc3999InvokeArgs();
    }


    [OutputType]
    public sealed class Unixtorfc3999Result
    {
        public readonly int Result;

        [OutputConstructor]
        private Unixtorfc3999Result(int result)
        {
            Result = result;
        }
    }
}