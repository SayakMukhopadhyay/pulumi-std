// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class Filebase64sha512Args extends com.pulumi.resources.InvokeArgs {

    public static final Filebase64sha512Args Empty = new Filebase64sha512Args();

    @Import(name="input", required=true)
    private Output<String> input;

    public Output<String> input() {
        return this.input;
    }

    private Filebase64sha512Args() {}

    private Filebase64sha512Args(Filebase64sha512Args $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(Filebase64sha512Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Filebase64sha512Args $;

        public Builder() {
            $ = new Filebase64sha512Args();
        }

        public Builder(Filebase64sha512Args defaults) {
            $ = new Filebase64sha512Args(Objects.requireNonNull(defaults));
        }

        public Builder input(Output<String> input) {
            $.input = input;
            return this;
        }

        public Builder input(String input) {
            return input(Output.of(input));
        }

        public Filebase64sha512Args build() {
            $.input = Objects.requireNonNull($.input, "expected parameter 'input' to be non-null");
            return $;
        }
    }

}
