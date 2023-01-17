// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class Sha1Args extends com.pulumi.resources.InvokeArgs {

    public static final Sha1Args Empty = new Sha1Args();

    @Import(name="input", required=true)
    private Output<String> input;

    public Output<String> input() {
        return this.input;
    }

    private Sha1Args() {}

    private Sha1Args(Sha1Args $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(Sha1Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Sha1Args $;

        public Builder() {
            $ = new Sha1Args();
        }

        public Builder(Sha1Args defaults) {
            $ = new Sha1Args(Objects.requireNonNull(defaults));
        }

        public Builder input(Output<String> input) {
            $.input = input;
            return this;
        }

        public Builder input(String input) {
            return input(Output.of(input));
        }

        public Sha1Args build() {
            $.input = Objects.requireNonNull($.input, "expected parameter 'input' to be non-null");
            return $;
        }
    }

}
