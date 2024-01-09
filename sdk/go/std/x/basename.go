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

// Returns the last element of the input path.
func Basename(ctx *pulumi.Context, args *BasenameArgs, opts ...pulumi.InvokeOption) (*BasenameResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv BasenameResult
	err := ctx.Invoke("std:index:basename", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type BasenameArgs struct {
	Input string `pulumi:"input"`
}

type BasenameResult struct {
	Result string `pulumi:"result"`
}

func BasenameOutput(ctx *pulumi.Context, args BasenameOutputArgs, opts ...pulumi.InvokeOption) BasenameResultOutput {
	outputResult := pulumix.ApplyErr[*BasenameArgs](args.ToOutput(), func(plainArgs *BasenameArgs) (*BasenameResult, error) {
		return Basename(ctx, plainArgs, opts...)
	})

	return pulumix.Cast[BasenameResultOutput, *BasenameResult](outputResult)
}

type BasenameOutputArgs struct {
	Input pulumix.Input[string] `pulumi:"input"`
}

func (args BasenameOutputArgs) ToOutput() pulumix.Output[*BasenameArgs] {
	allArgs := pulumix.All(
		args.Input.ToOutput(context.Background()).AsAny())
	return pulumix.Apply[[]any](allArgs, func(resolvedArgs []interface{}) *BasenameArgs {
		return &BasenameArgs{
			Input: resolvedArgs[0].(string),
		}
	})
}

type BasenameResultOutput struct{ *pulumi.OutputState }

func (BasenameResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BasenameResult)(nil)).Elem()
}

func (o BasenameResultOutput) ToOutput(context.Context) pulumix.Output[*BasenameResult] {
	return pulumix.Output[*BasenameResult]{
		OutputState: o.OutputState,
	}
}

func (o BasenameResultOutput) Result() pulumix.Output[string] {
	return pulumix.Apply[*BasenameResult](o, func(v *BasenameResult) string { return v.Result })
}
