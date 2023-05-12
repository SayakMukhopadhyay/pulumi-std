// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Max
    {
        /// <summary>
        /// Returns the largest of the floats.
        /// </summary>
        public static Task<MaxResult> InvokeAsync(MaxArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<MaxResult>("std:index:max", args ?? new MaxArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the largest of the floats.
        /// </summary>
        public static Output<MaxResult> Invoke(MaxInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<MaxResult>("std:index:max", args ?? new MaxInvokeArgs(), options.WithDefaults());
    }


    public sealed class MaxArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        private List<double>? _input;
        public List<double> Input
        {
            get => _input ?? (_input = new List<double>());
            set => _input = value;
        }

        public MaxArgs()
        {
        }
        public static new MaxArgs Empty => new MaxArgs();
    }

    public sealed class MaxInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        private InputList<double>? _input;
        public InputList<double> Input
        {
            get => _input ?? (_input = new InputList<double>());
            set => _input = value;
        }

        public MaxInvokeArgs()
        {
        }
        public static new MaxInvokeArgs Empty => new MaxInvokeArgs();
    }


    [OutputType]
    public sealed class MaxResult
    {
        public readonly double Result;

        [OutputConstructor]
        private MaxResult(double result)
        {
            Result = result;
        }
    }
}