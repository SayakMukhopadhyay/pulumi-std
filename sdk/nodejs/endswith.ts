// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Determines if the input string ends with the suffix.
 */
export function endswith(args: EndswithArgs, opts?: pulumi.InvokeOptions): Promise<EndswithResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:endswith", {
        "input": args.input,
        "suffix": args.suffix,
    }, opts);
}

export interface EndswithArgs {
    input: string;
    suffix: string;
}

export interface EndswithResult {
    readonly result: boolean;
}
/**
 * Determines if the input string ends with the suffix.
 */
export function endswithOutput(args: EndswithOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<EndswithResult> {
    return pulumi.output(args).apply((a: any) => endswith(a, opts))
}

export interface EndswithOutputArgs {
    input: pulumi.Input<string>;
    suffix: pulumi.Input<string>;
}