// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class Unixtorfc3999Args extends com.pulumi.resources.InvokeArgs {

    public static final Unixtorfc3999Args Empty = new Unixtorfc3999Args();

    @Import(name="input", required=true)
    private Output<String> input;

    public Output<String> input() {
        return this.input;
    }

    private Unixtorfc3999Args() {}

    private Unixtorfc3999Args(Unixtorfc3999Args $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(Unixtorfc3999Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Unixtorfc3999Args $;

        public Builder() {
            $ = new Unixtorfc3999Args();
        }

        public Builder(Unixtorfc3999Args defaults) {
            $ = new Unixtorfc3999Args(Objects.requireNonNull(defaults));
        }

        public Builder input(Output<String> input) {
            $.input = input;
            return this;
        }

        public Builder input(String input) {
            return input(Output.of(input));
        }

        public Unixtorfc3999Args build() {
            $.input = Objects.requireNonNull($.input, "expected parameter 'input' to be non-null");
            return $;
        }
    }

}
