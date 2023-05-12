// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Produces a list by dividing a given string at all occurrences of a given separator
func Split(ctx *pulumi.Context, args *SplitArgs, opts ...pulumi.InvokeOption) (*SplitResult, error) {
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
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (SplitResult, error) {
			args := v.(SplitArgs)
			r, err := Split(ctx, &args, opts...)
			var s SplitResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(SplitResultOutput)
}

type SplitOutputArgs struct {
	Separator pulumi.StringInput `pulumi:"separator"`
	Text      pulumi.StringInput `pulumi:"text"`
}

func (SplitOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SplitArgs)(nil)).Elem()
}

type SplitResultOutput struct{ *pulumi.OutputState }

func (SplitResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SplitResult)(nil)).Elem()
}

func (o SplitResultOutput) ToSplitResultOutput() SplitResultOutput {
	return o
}

func (o SplitResultOutput) ToSplitResultOutputWithContext(ctx context.Context) SplitResultOutput {
	return o
}

func (o SplitResultOutput) Result() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SplitResult) []string { return v.Result }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(SplitResultOutput{})
}