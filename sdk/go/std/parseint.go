// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Parses the given string as a representation of an integer in the specified base
// and returns the resulting number. The base must be between 2 and 62 inclusive.
//
//	.
func Parseint(ctx *pulumi.Context, args *ParseintArgs, opts ...pulumi.InvokeOption) (*ParseintResult, error) {
	var rv ParseintResult
	err := ctx.Invoke("std:index:parseint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ParseintArgs struct {
	Base  *int   `pulumi:"base"`
	Input string `pulumi:"input"`
}

type ParseintResult struct {
	Result int `pulumi:"result"`
}

func ParseintOutput(ctx *pulumi.Context, args ParseintOutputArgs, opts ...pulumi.InvokeOption) ParseintResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ParseintResult, error) {
			args := v.(ParseintArgs)
			r, err := Parseint(ctx, &args, opts...)
			var s ParseintResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ParseintResultOutput)
}

type ParseintOutputArgs struct {
	Base  pulumi.IntPtrInput `pulumi:"base"`
	Input pulumi.StringInput `pulumi:"input"`
}

func (ParseintOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParseintArgs)(nil)).Elem()
}

type ParseintResultOutput struct{ *pulumi.OutputState }

func (ParseintResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParseintResult)(nil)).Elem()
}

func (o ParseintResultOutput) ToParseintResultOutput() ParseintResultOutput {
	return o
}

func (o ParseintResultOutput) ToParseintResultOutputWithContext(ctx context.Context) ParseintResultOutput {
	return o
}

func (o ParseintResultOutput) Result() pulumi.IntOutput {
	return o.ApplyT(func(v ParseintResult) int { return v.Result }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(ParseintResultOutput{})
}