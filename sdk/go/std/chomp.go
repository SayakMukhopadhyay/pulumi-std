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

// Removes one or more newline characters from the end of the given string.
func Chomp(ctx *pulumi.Context, args *ChompArgs, opts ...pulumi.InvokeOption) (*ChompResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ChompResult
	err := ctx.Invoke("std:index:chomp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ChompArgs struct {
	Input string `pulumi:"input"`
}

type ChompResult struct {
	Result string `pulumi:"result"`
}

func ChompOutput(ctx *pulumi.Context, args ChompOutputArgs, opts ...pulumi.InvokeOption) ChompResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ChompResult, error) {
			args := v.(ChompArgs)
			r, err := Chomp(ctx, &args, opts...)
			var s ChompResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ChompResultOutput)
}

type ChompOutputArgs struct {
	Input pulumi.StringInput `pulumi:"input"`
}

func (ChompOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ChompArgs)(nil)).Elem()
}

type ChompResultOutput struct{ *pulumi.OutputState }

func (ChompResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ChompResult)(nil)).Elem()
}

func (o ChompResultOutput) ToChompResultOutput() ChompResultOutput {
	return o
}

func (o ChompResultOutput) ToChompResultOutputWithContext(ctx context.Context) ChompResultOutput {
	return o
}

func (o ChompResultOutput) ToOutput(ctx context.Context) pulumix.Output[ChompResult] {
	return pulumix.Output[ChompResult]{
		OutputState: o.OutputState,
	}
}

func (o ChompResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v ChompResult) string { return v.Result }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ChompResultOutput{})
}
