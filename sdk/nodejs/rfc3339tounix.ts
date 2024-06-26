// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Converts a RFC3999 formatted timestamp into a Unix timestamp with milliseconds.
 */
export function rfc3339tounix(args: Rfc3339tounixArgs, opts?: pulumi.InvokeOptions): Promise<Rfc3339tounixResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:rfc3339tounix", {
        "input": args.input,
    }, opts);
}

export interface Rfc3339tounixArgs {
    input: string;
}

export interface Rfc3339tounixResult {
    readonly result: number;
}
/**
 * Converts a RFC3999 formatted timestamp into a Unix timestamp with milliseconds.
 */
export function rfc3339tounixOutput(args: Rfc3339tounixOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<Rfc3339tounixResult> {
    return pulumi.output(args).apply((a: any) => rfc3339tounix(a, opts))
}

export interface Rfc3339tounixOutputArgs {
    input: pulumi.Input<string>;
}
