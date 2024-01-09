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

// Reads the contents of a file into a string and returns the SHA256 hash of it.
func Filesha256(ctx *pulumi.Context, args *Filesha256Args, opts ...pulumi.InvokeOption) (*Filesha256Result, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv Filesha256Result
	err := ctx.Invoke("std:index:filesha256", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type Filesha256Args struct {
	Input string `pulumi:"input"`
}

type Filesha256Result struct {
	Result string `pulumi:"result"`
}

func Filesha256Output(ctx *pulumi.Context, args Filesha256OutputArgs, opts ...pulumi.InvokeOption) Filesha256ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (Filesha256Result, error) {
			args := v.(Filesha256Args)
			r, err := Filesha256(ctx, &args, opts...)
			var s Filesha256Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(Filesha256ResultOutput)
}

type Filesha256OutputArgs struct {
	Input pulumi.StringInput `pulumi:"input"`
}

func (Filesha256OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Filesha256Args)(nil)).Elem()
}

type Filesha256ResultOutput struct{ *pulumi.OutputState }

func (Filesha256ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Filesha256Result)(nil)).Elem()
}

func (o Filesha256ResultOutput) ToFilesha256ResultOutput() Filesha256ResultOutput {
	return o
}

func (o Filesha256ResultOutput) ToFilesha256ResultOutputWithContext(ctx context.Context) Filesha256ResultOutput {
	return o
}

func (o Filesha256ResultOutput) ToOutput(ctx context.Context) pulumix.Output[Filesha256Result] {
	return pulumix.Output[Filesha256Result]{
		OutputState: o.OutputState,
	}
}

func (o Filesha256ResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v Filesha256Result) string { return v.Result }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(Filesha256ResultOutput{})
}
