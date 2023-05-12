// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns a copy of the string with all Unicode letters mapped to their lower case.
func Lower(ctx *pulumi.Context, args *LowerArgs, opts ...pulumi.InvokeOption) (*LowerResult, error) {
	var rv LowerResult
	err := ctx.Invoke("std:index:lower", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LowerArgs struct {
	Input string `pulumi:"input"`
}

type LowerResult struct {
	Result string `pulumi:"result"`
}

func LowerOutput(ctx *pulumi.Context, args LowerOutputArgs, opts ...pulumi.InvokeOption) LowerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LowerResult, error) {
			args := v.(LowerArgs)
			r, err := Lower(ctx, &args, opts...)
			var s LowerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LowerResultOutput)
}

type LowerOutputArgs struct {
	Input pulumi.StringInput `pulumi:"input"`
}

func (LowerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LowerArgs)(nil)).Elem()
}

type LowerResultOutput struct{ *pulumi.OutputState }

func (LowerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LowerResult)(nil)).Elem()
}

func (o LowerResultOutput) ToLowerResultOutput() LowerResultOutput {
	return o
}

func (o LowerResultOutput) ToLowerResultOutputWithContext(ctx context.Context) LowerResultOutput {
	return o
}

func (o LowerResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v LowerResult) string { return v.Result }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LowerResultOutput{})
}