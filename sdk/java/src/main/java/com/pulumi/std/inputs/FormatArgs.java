// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class FormatArgs extends com.pulumi.resources.InvokeArgs {

    public static final FormatArgs Empty = new FormatArgs();

    @Import(name="args", required=true)
    private Output<List<Object>> args;

    public Output<List<Object>> args() {
        return this.args;
    }

    @Import(name="input", required=true)
    private Output<String> input;

    public Output<String> input() {
        return this.input;
    }

    private FormatArgs() {}

    private FormatArgs(FormatArgs $) {
        this.args = $.args;
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FormatArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FormatArgs $;

        public Builder() {
            $ = new FormatArgs();
        }

        public Builder(FormatArgs defaults) {
            $ = new FormatArgs(Objects.requireNonNull(defaults));
        }

        public Builder args(Output<List<Object>> args) {
            $.args = args;
            return this;
        }

        public Builder args(List<Object> args) {
            return args(Output.of(args));
        }

        public Builder args(Object... args) {
            return args(List.of(args));
        }

        public Builder input(Output<String> input) {
            $.input = input;
            return this;
        }

        public Builder input(String input) {
            return input(Output.of(input));
        }

        public FormatArgs build() {
            $.args = Objects.requireNonNull($.args, "expected parameter 'args' to be non-null");
            $.input = Objects.requireNonNull($.input, "expected parameter 'input' to be non-null");
            return $;
        }
    }

}
