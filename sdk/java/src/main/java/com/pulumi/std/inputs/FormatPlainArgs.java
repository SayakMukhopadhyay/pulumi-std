// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class FormatPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final FormatPlainArgs Empty = new FormatPlainArgs();

    @Import(name="args", required=true)
    private List<Object> args;

    public List<Object> args() {
        return this.args;
    }

    @Import(name="input", required=true)
    private String input;

    public String input() {
        return this.input;
    }

    private FormatPlainArgs() {}

    private FormatPlainArgs(FormatPlainArgs $) {
        this.args = $.args;
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FormatPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FormatPlainArgs $;

        public Builder() {
            $ = new FormatPlainArgs();
        }

        public Builder(FormatPlainArgs defaults) {
            $ = new FormatPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder args(List<Object> args) {
            $.args = args;
            return this;
        }

        public Builder args(Object... args) {
            return args(List.of(args));
        }

        public Builder input(String input) {
            $.input = input;
            return this;
        }

        public FormatPlainArgs build() {
            $.args = Objects.requireNonNull($.args, "expected parameter 'args' to be non-null");
            $.input = Objects.requireNonNull($.input, "expected parameter 'input' to be non-null");
            return $;
        }
    }

}