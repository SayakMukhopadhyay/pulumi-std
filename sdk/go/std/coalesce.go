// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-std/sdk/go/std/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Returns the first non-empty value from the given arguments.
func Coalesce(ctx *pulumi.Context, args *CoalesceArgs, opts ...pulumi.InvokeOption) (*CoalesceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv CoalesceResult
	err := ctx.Invoke("std:index:coalesce", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type CoalesceArgs struct {
	Input []string `pulumi:"input"`
}

type CoalesceResult struct {
	Result string `pulumi:"result"`
}

func CoalesceOutput(ctx *pulumi.Context, args CoalesceOutputArgs, opts ...pulumi.InvokeOption) CoalesceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (CoalesceResult, error) {
			args := v.(CoalesceArgs)
			r, err := Coalesce(ctx, &args, opts...)
			var s CoalesceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(CoalesceResultOutput)
}

type CoalesceOutputArgs struct {
	Input pulumi.StringArrayInput `pulumi:"input"`
}

func (CoalesceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CoalesceArgs)(nil)).Elem()
}

type CoalesceResultOutput struct{ *pulumi.OutputState }

func (CoalesceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CoalesceResult)(nil)).Elem()
}

func (o CoalesceResultOutput) ToCoalesceResultOutput() CoalesceResultOutput {
	return o
}

func (o CoalesceResultOutput) ToCoalesceResultOutputWithContext(ctx context.Context) CoalesceResultOutput {
	return o
}

func (o CoalesceResultOutput) ToOutput(ctx context.Context) pulumix.Output[CoalesceResult] {
	return pulumix.Output[CoalesceResult]{
		OutputState: o.OutputState,
	}
}

func (o CoalesceResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v CoalesceResult) string { return v.Result }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CoalesceResultOutput{})
}
