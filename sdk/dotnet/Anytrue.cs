// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Anytrue
    {
        /// <summary>
        /// Returns true if any of the elements in a given collection are true or \"true\".
        /// It also returns false if the collection is empty.
        /// </summary>
        public static Task<AnytrueResult> InvokeAsync(AnytrueArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<AnytrueResult>("std:index:anytrue", args ?? new AnytrueArgs(), options.WithDefaults());

        /// <summary>
        /// Returns true if any of the elements in a given collection are true or \"true\".
        /// It also returns false if the collection is empty.
        /// </summary>
        public static Output<AnytrueResult> Invoke(AnytrueInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<AnytrueResult>("std:index:anytrue", args ?? new AnytrueInvokeArgs(), options.WithDefaults());
    }


    public sealed class AnytrueArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        private List<object>? _input;
        public List<object> Input
        {
            get => _input ?? (_input = new List<object>());
            set => _input = value;
        }

        public AnytrueArgs()
        {
        }
        public static new AnytrueArgs Empty => new AnytrueArgs();
    }

    public sealed class AnytrueInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        private InputList<object>? _input;
        public InputList<object> Input
        {
            get => _input ?? (_input = new InputList<object>());
            set => _input = value;
        }

        public AnytrueInvokeArgs()
        {
        }
        public static new AnytrueInvokeArgs Empty => new AnytrueInvokeArgs();
    }


    [OutputType]
    public sealed class AnytrueResult
    {
        public readonly bool Result;

        [OutputConstructor]
        private AnytrueResult(bool result)
        {
            Result = result;
        }
    }
}
