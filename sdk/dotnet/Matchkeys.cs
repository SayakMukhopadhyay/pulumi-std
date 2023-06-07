// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Matchkeys
    {
        /// <summary>
        /// For two lists values and keys of equal length,
        /// returns all elements from values where the corresponding element from keys exists in the searchset list.
        /// </summary>
        public static Task<MatchkeysResult> InvokeAsync(MatchkeysArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<MatchkeysResult>("std:index:matchkeys", args ?? new MatchkeysArgs(), options.WithDefaults());

        /// <summary>
        /// For two lists values and keys of equal length,
        /// returns all elements from values where the corresponding element from keys exists in the searchset list.
        /// </summary>
        public static Output<MatchkeysResult> Invoke(MatchkeysInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<MatchkeysResult>("std:index:matchkeys", args ?? new MatchkeysInvokeArgs(), options.WithDefaults());
    }


    public sealed class MatchkeysArgs : global::Pulumi.InvokeArgs
    {
        [Input("searchList", required: true)]
        private List<string>? _searchList;
        public List<string> SearchList
        {
            get => _searchList ?? (_searchList = new List<string>());
            set => _searchList = value;
        }

        [Input("values", required: true)]
        private List<string>? _values;
        public List<string> Values
        {
            get => _values ?? (_values = new List<string>());
            set => _values = value;
        }

        public MatchkeysArgs()
        {
        }
        public static new MatchkeysArgs Empty => new MatchkeysArgs();
    }

    public sealed class MatchkeysInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("searchList", required: true)]
        private InputList<string>? _searchList;
        public InputList<string> SearchList
        {
            get => _searchList ?? (_searchList = new InputList<string>());
            set => _searchList = value;
        }

        [Input("values", required: true)]
        private InputList<string>? _values;
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public MatchkeysInvokeArgs()
        {
        }
        public static new MatchkeysInvokeArgs Empty => new MatchkeysInvokeArgs();
    }


    [OutputType]
    public sealed class MatchkeysResult
    {
        public readonly ImmutableArray<string> Result;

        [OutputConstructor]
        private MatchkeysResult(ImmutableArray<string> result)
        {
            Result = result;
        }
    }
}
