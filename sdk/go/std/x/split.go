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

// Produces a list by dividing a given string at all occurrences of a given separator
func Split(ctx *pulumi.Context, args *SplitArgs, opts ...pulumi.InvokeOption) (*SplitResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv SplitResult
	err := ctx.Invoke("std:index:split", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type SplitArgs struct {
	Separator string `pulumi:"separator"`
	Text      string `pulumi:"text"`
}

type SplitResult struct {
	Result []string `pulumi:"result"`
}

func SplitOutput(ctx *pulumi.Context, args SplitOutputArgs, opts ...pulumi.InvokeOption) SplitResultOutput {
	outputResult := pulumix.ApplyErr[*SplitArgs](args.ToOutput(), func(plainArgs *SplitArgs) (*SplitResult, error) {
		return Split(ctx, plainArgs, opts...)
	})

	return pulumix.Cast[SplitResultOutput, *SplitResult](outputResult)
}

type SplitOutputArgs struct {
	Separator pulumix.Input[string] `pulumi:"separator"`
	Text      pulumix.Input[string] `pulumi:"text"`
}

func (args SplitOutputArgs) ToOutput() pulumix.Output[*SplitArgs] {
	allArgs := pulumix.All(
		args.Separator.ToOutput(context.Background()).AsAny(),
		args.Text.ToOutput(context.Background()).AsAny())
	return pulumix.Apply[[]any](allArgs, func(resolvedArgs []interface{}) *SplitArgs {
		return &SplitArgs{
			Separator: resolvedArgs[0].(string),
			Text:      resolvedArgs[1].(string),
		}
	})
}

type SplitResultOutput struct{ *pulumi.OutputState }

func (SplitResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SplitResult)(nil)).Elem()
}

func (o SplitResultOutput) ToOutput(context.Context) pulumix.Output[*SplitResult] {
	return pulumix.Output[*SplitResult]{
		OutputState: o.OutputState,
	}
}

func (o SplitResultOutput) Result() pulumix.ArrayOutput[string] {
	value := pulumix.Apply[*SplitResult](o, func(v *SplitResult) []string { return v.Result })
	return pulumix.ArrayOutput[string]{
		OutputState: value.OutputState,
	}
}
