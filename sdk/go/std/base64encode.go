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

// Returns a base64-encoded representation of the given string.
func Base64encode(ctx *pulumi.Context, args *Base64encodeArgs, opts ...pulumi.InvokeOption) (*Base64encodeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv Base64encodeResult
	err := ctx.Invoke("std:index:base64encode", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type Base64encodeArgs struct {
	Input string `pulumi:"input"`
}

type Base64encodeResult struct {
	Result string `pulumi:"result"`
}

func Base64encodeOutput(ctx *pulumi.Context, args Base64encodeOutputArgs, opts ...pulumi.InvokeOption) Base64encodeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (Base64encodeResult, error) {
			args := v.(Base64encodeArgs)
			r, err := Base64encode(ctx, &args, opts...)
			var s Base64encodeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(Base64encodeResultOutput)
}

type Base64encodeOutputArgs struct {
	Input pulumi.StringInput `pulumi:"input"`
}

func (Base64encodeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Base64encodeArgs)(nil)).Elem()
}

type Base64encodeResultOutput struct{ *pulumi.OutputState }

func (Base64encodeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Base64encodeResult)(nil)).Elem()
}

func (o Base64encodeResultOutput) ToBase64encodeResultOutput() Base64encodeResultOutput {
	return o
}

func (o Base64encodeResultOutput) ToBase64encodeResultOutputWithContext(ctx context.Context) Base64encodeResultOutput {
	return o
}

func (o Base64encodeResultOutput) ToOutput(ctx context.Context) pulumix.Output[Base64encodeResult] {
	return pulumix.Output[Base64encodeResult]{
		OutputState: o.OutputState,
	}
}

func (o Base64encodeResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v Base64encodeResult) string { return v.Result }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(Base64encodeResultOutput{})
}
