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

// Returns the Blowfish encrypted hash of the string at the given cost.
// A default cost of 10 will be used if not provided.
func Bcrypt(ctx *pulumi.Context, args *BcryptArgs, opts ...pulumi.InvokeOption) (*BcryptResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv BcryptResult
	err := ctx.Invoke("std:index:bcrypt", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type BcryptArgs struct {
	Cost  *int   `pulumi:"cost"`
	Input string `pulumi:"input"`
}

type BcryptResult struct {
	Result string `pulumi:"result"`
}

func BcryptOutput(ctx *pulumi.Context, args BcryptOutputArgs, opts ...pulumi.InvokeOption) BcryptResultOutput {
	outputResult := pulumix.ApplyErr[*BcryptArgs](args.ToOutput(), func(plainArgs *BcryptArgs) (*BcryptResult, error) {
		return Bcrypt(ctx, plainArgs, opts...)
	})

	return pulumix.Cast[BcryptResultOutput, *BcryptResult](outputResult)
}

type BcryptOutputArgs struct {
	Cost  pulumix.Input[*int]   `pulumi:"cost"`
	Input pulumix.Input[string] `pulumi:"input"`
}

func (args BcryptOutputArgs) ToOutput() pulumix.Output[*BcryptArgs] {
	allArgs := pulumix.All(
		args.Cost.ToOutput(context.Background()).AsAny(),
		args.Input.ToOutput(context.Background()).AsAny())
	return pulumix.Apply[[]any](allArgs, func(resolvedArgs []interface{}) *BcryptArgs {
		return &BcryptArgs{
			Cost:  resolvedArgs[0].(*int),
			Input: resolvedArgs[1].(string),
		}
	})
}

type BcryptResultOutput struct{ *pulumi.OutputState }

func (BcryptResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BcryptResult)(nil)).Elem()
}

func (o BcryptResultOutput) ToOutput(context.Context) pulumix.Output[*BcryptResult] {
	return pulumix.Output[*BcryptResult]{
		OutputState: o.OutputState,
	}
}

func (o BcryptResultOutput) Result() pulumix.Output[string] {
	return pulumix.Apply[*BcryptResult](o, func(v *BcryptResult) string { return v.Result })
}
