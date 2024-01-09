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

// Returns the greatest integer value less than or equal to the argument.
func Log(ctx *pulumi.Context, args *LogArgs, opts ...pulumi.InvokeOption) (*LogResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LogResult
	err := ctx.Invoke("std:index:log", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LogArgs struct {
	Base  float64 `pulumi:"base"`
	Input float64 `pulumi:"input"`
}

type LogResult struct {
	Result float64 `pulumi:"result"`
}

func LogOutput(ctx *pulumi.Context, args LogOutputArgs, opts ...pulumi.InvokeOption) LogResultOutput {
	outputResult := pulumix.ApplyErr[*LogArgs](args.ToOutput(), func(plainArgs *LogArgs) (*LogResult, error) {
		return Log(ctx, plainArgs, opts...)
	})

	return pulumix.Cast[LogResultOutput, *LogResult](outputResult)
}

type LogOutputArgs struct {
	Base  pulumix.Input[float64] `pulumi:"base"`
	Input pulumix.Input[float64] `pulumi:"input"`
}

func (args LogOutputArgs) ToOutput() pulumix.Output[*LogArgs] {
	allArgs := pulumix.All(
		args.Base.ToOutput(context.Background()).AsAny(),
		args.Input.ToOutput(context.Background()).AsAny())
	return pulumix.Apply[[]any](allArgs, func(resolvedArgs []interface{}) *LogArgs {
		return &LogArgs{
			Base:  resolvedArgs[0].(float64),
			Input: resolvedArgs[1].(float64),
		}
	})
}

type LogResultOutput struct{ *pulumi.OutputState }

func (LogResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogResult)(nil)).Elem()
}

func (o LogResultOutput) ToOutput(context.Context) pulumix.Output[*LogResult] {
	return pulumix.Output[*LogResult]{
		OutputState: o.OutputState,
	}
}

func (o LogResultOutput) Result() pulumix.Output[float64] {
	return pulumix.Apply[*LogResult](o, func(v *LogResult) float64 { return v.Result })
}
