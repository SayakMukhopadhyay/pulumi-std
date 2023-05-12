// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Converts the first letter of each word in the given string to uppercase.
 */
export function title(args: TitleArgs, opts?: pulumi.InvokeOptions): Promise<TitleResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:title", {
        "input": args.input,
    }, opts);
}

export interface TitleArgs {
    input: string;
}

export interface TitleResult {
    readonly result: string;
}
/**
 * Converts the first letter of each word in the given string to uppercase.
 */
export function titleOutput(args: TitleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<TitleResult> {
    return pulumi.output(args).apply((a: any) => title(a, opts))
}

export interface TitleOutputArgs {
    input: pulumi.Input<string>;
}