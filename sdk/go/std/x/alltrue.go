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

// Returns true if all elements in a given collection are true or \"true\".
// It also returns true if the collection is empty.
func Alltrue(ctx *pulumi.Context, args *AlltrueArgs, opts ...pulumi.InvokeOption) (*AlltrueResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv AlltrueResult
	err := ctx.Invoke("std:index:alltrue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type AlltrueArgs struct {
	Input []interface{} `pulumi:"input"`
}

type AlltrueResult struct {
	Result bool `pulumi:"result"`
}

func AlltrueOutput(ctx *pulumi.Context, args AlltrueOutputArgs, opts ...pulumi.InvokeOption) AlltrueResultOutput {
	outputResult := pulumix.ApplyErr[*AlltrueArgs](args.ToOutput(), func(plainArgs *AlltrueArgs) (*AlltrueResult, error) {
		return Alltrue(ctx, plainArgs, opts...)
	})

	return pulumix.Cast[AlltrueResultOutput, *AlltrueResult](outputResult)
}

type AlltrueOutputArgs struct {
	Input pulumix.Input[[]interface{}] `pulumi:"input"`
}

func (args AlltrueOutputArgs) ToOutput() pulumix.Output[*AlltrueArgs] {
	allArgs := pulumix.All(
		args.Input.ToOutput(context.Background()).AsAny())
	return pulumix.Apply[[]any](allArgs, func(resolvedArgs []interface{}) *AlltrueArgs {
		return &AlltrueArgs{
			Input: resolvedArgs[0].([]interface{}),
		}
	})
}

type AlltrueResultOutput struct{ *pulumi.OutputState }

func (AlltrueResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlltrueResult)(nil)).Elem()
}

func (o AlltrueResultOutput) ToOutput(context.Context) pulumix.Output[*AlltrueResult] {
	return pulumix.Output[*AlltrueResult]{
		OutputState: o.OutputState,
	}
}

func (o AlltrueResultOutput) Result() pulumix.Output[bool] {
	return pulumix.Apply[*AlltrueResult](o, func(v *AlltrueResult) bool { return v.Result })
}
