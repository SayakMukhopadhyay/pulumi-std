// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class Unixtorfc3999PlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final Unixtorfc3999PlainArgs Empty = new Unixtorfc3999PlainArgs();

    @Import(name="input", required=true)
    private String input;

    public String input() {
        return this.input;
    }

    private Unixtorfc3999PlainArgs() {}

    private Unixtorfc3999PlainArgs(Unixtorfc3999PlainArgs $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(Unixtorfc3999PlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Unixtorfc3999PlainArgs $;

        public Builder() {
            $ = new Unixtorfc3999PlainArgs();
        }

        public Builder(Unixtorfc3999PlainArgs defaults) {
            $ = new Unixtorfc3999PlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(String input) {
            $.input = input;
            return this;
        }

        public Unixtorfc3999PlainArgs build() {
            $.input = Objects.requireNonNull($.input, "expected parameter 'input' to be non-null");
            return $;
        }
    }

}
