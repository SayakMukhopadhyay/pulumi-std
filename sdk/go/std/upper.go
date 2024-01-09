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

// Converts all cased letters in the given string to uppercase.
func Upper(ctx *pulumi.Context, args *UpperArgs, opts ...pulumi.InvokeOption) (*UpperResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv UpperResult
	err := ctx.Invoke("std:index:upper", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type UpperArgs struct {
	Input string `pulumi:"input"`
}

type UpperResult struct {
	Result string `pulumi:"result"`
}

func UpperOutput(ctx *pulumi.Context, args UpperOutputArgs, opts ...pulumi.InvokeOption) UpperResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (UpperResult, error) {
			args := v.(UpperArgs)
			r, err := Upper(ctx, &args, opts...)
			var s UpperResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(UpperResultOutput)
}

type UpperOutputArgs struct {
	Input pulumi.StringInput `pulumi:"input"`
}

func (UpperOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UpperArgs)(nil)).Elem()
}

type UpperResultOutput struct{ *pulumi.OutputState }

func (UpperResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UpperResult)(nil)).Elem()
}

func (o UpperResultOutput) ToUpperResultOutput() UpperResultOutput {
	return o
}

func (o UpperResultOutput) ToUpperResultOutputWithContext(ctx context.Context) UpperResultOutput {
	return o
}

func (o UpperResultOutput) ToOutput(ctx context.Context) pulumix.Output[UpperResult] {
	return pulumix.Output[UpperResult]{
		OutputState: o.OutputState,
	}
}

func (o UpperResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v UpperResult) string { return v.Result }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(UpperResultOutput{})
}
