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

// Returns the element at the specified index.
func Element(ctx *pulumi.Context, args *ElementArgs, opts ...pulumi.InvokeOption) (*ElementResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ElementResult
	err := ctx.Invoke("std:index:element", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ElementArgs struct {
	Index int           `pulumi:"index"`
	Input []interface{} `pulumi:"input"`
}

type ElementResult struct {
	Result interface{} `pulumi:"result"`
}

func ElementOutput(ctx *pulumi.Context, args ElementOutputArgs, opts ...pulumi.InvokeOption) ElementResultOutput {
	outputResult := pulumix.ApplyErr[*ElementArgs](args.ToOutput(), func(plainArgs *ElementArgs) (*ElementResult, error) {
		return Element(ctx, plainArgs, opts...)
	})

	return pulumix.Cast[ElementResultOutput, *ElementResult](outputResult)
}

type ElementOutputArgs struct {
	Index pulumix.Input[int]           `pulumi:"index"`
	Input pulumix.Input[[]interface{}] `pulumi:"input"`
}

func (args ElementOutputArgs) ToOutput() pulumix.Output[*ElementArgs] {
	allArgs := pulumix.All(
		args.Index.ToOutput(context.Background()).AsAny(),
		args.Input.ToOutput(context.Background()).AsAny())
	return pulumix.Apply[[]any](allArgs, func(resolvedArgs []interface{}) *ElementArgs {
		return &ElementArgs{
			Index: resolvedArgs[0].(int),
			Input: resolvedArgs[1].([]interface{}),
		}
	})
}

type ElementResultOutput struct{ *pulumi.OutputState }

func (ElementResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ElementResult)(nil)).Elem()
}

func (o ElementResultOutput) ToOutput(context.Context) pulumix.Output[*ElementResult] {
	return pulumix.Output[*ElementResult]{
		OutputState: o.OutputState,
	}
}

func (o ElementResultOutput) Result() pulumix.Output[interface{}] {
	return pulumix.Apply[*ElementResult](o, func(v *ElementResult) interface{} { return v.Result })
}
