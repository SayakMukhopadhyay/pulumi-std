// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Object;
import java.util.List;
import java.util.Objects;


public final class ConcatPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final ConcatPlainArgs Empty = new ConcatPlainArgs();

    @Import(name="input", required=true)
    private List<List<Object>> input;

    public List<List<Object>> input() {
        return this.input;
    }

    private ConcatPlainArgs() {}

    private ConcatPlainArgs(ConcatPlainArgs $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ConcatPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ConcatPlainArgs $;

        public Builder() {
            $ = new ConcatPlainArgs();
        }

        public Builder(ConcatPlainArgs defaults) {
            $ = new ConcatPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(List<List<Object>> input) {
            $.input = input;
            return this;
        }

        public Builder input(List<Object>... input) {
            return input(List.of(input));
        }

        public ConcatPlainArgs build() {
            $.input = Objects.requireNonNull($.input, "expected parameter 'input' to be non-null");
            return $;
        }
    }

}
