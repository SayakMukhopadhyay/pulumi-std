// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Uuid
    {
        /// <summary>
        /// Returns a unique identifier string, generated and formatted as required by RFC 4122.
        /// </summary>
        public static Task<UuidResult> InvokeAsync(UuidArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<UuidResult>("std:index:uuid", args ?? new UuidArgs(), options.WithDefaults());

        /// <summary>
        /// Returns a unique identifier string, generated and formatted as required by RFC 4122.
        /// </summary>
        public static Output<UuidResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<UuidResult>("std:index:uuid", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class UuidArgs : global::Pulumi.InvokeArgs
    {
        public UuidArgs()
        {
        }
        public static new UuidArgs Empty => new UuidArgs();
    }


    [OutputType]
    public sealed class UuidResult
    {
        public readonly string Result;

        [OutputConstructor]
        private UuidResult(string result)
        {
            Result = result;
        }
    }
}
