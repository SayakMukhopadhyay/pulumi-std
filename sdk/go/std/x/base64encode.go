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
	outputResult := pulumix.ApplyErr[*Base64encodeArgs](args.ToOutput(), func(plainArgs *Base64encodeArgs) (*Base64encodeResult, error) {
		return Base64encode(ctx, plainArgs, opts...)
	})

	return pulumix.Cast[Base64encodeResultOutput, *Base64encodeResult](outputResult)
}

type Base64encodeOutputArgs struct {
	Input pulumix.Input[string] `pulumi:"input"`
}

func (args Base64encodeOutputArgs) ToOutput() pulumix.Output[*Base64encodeArgs] {
	allArgs := pulumix.All(
		args.Input.ToOutput(context.Background()).AsAny())
	return pulumix.Apply[[]any](allArgs, func(resolvedArgs []interface{}) *Base64encodeArgs {
		return &Base64encodeArgs{
			Input: resolvedArgs[0].(string),
		}
	})
}

type Base64encodeResultOutput struct{ *pulumi.OutputState }

func (Base64encodeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Base64encodeResult)(nil)).Elem()
}

func (o Base64encodeResultOutput) ToOutput(context.Context) pulumix.Output[*Base64encodeResult] {
	return pulumix.Output[*Base64encodeResult]{
		OutputState: o.OutputState,
	}
}

func (o Base64encodeResultOutput) Result() pulumix.Output[string] {
	return pulumix.Apply[*Base64encodeResult](o, func(v *Base64encodeResult) string { return v.Result })
}
