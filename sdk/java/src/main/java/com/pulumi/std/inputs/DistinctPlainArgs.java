// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Object;
import java.util.List;
import java.util.Objects;


public final class DistinctPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final DistinctPlainArgs Empty = new DistinctPlainArgs();

    @Import(name="input", required=true)
    private List<Object> input;

    public List<Object> input() {
        return this.input;
    }

    private DistinctPlainArgs() {}

    private DistinctPlainArgs(DistinctPlainArgs $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DistinctPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DistinctPlainArgs $;

        public Builder() {
            $ = new DistinctPlainArgs();
        }

        public Builder(DistinctPlainArgs defaults) {
            $ = new DistinctPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(List<Object> input) {
            $.input = input;
            return this;
        }

        public Builder input(Object... input) {
            return input(List.of(input));
        }

        public DistinctPlainArgs build() {
            $.input = Objects.requireNonNull($.input, "expected parameter 'input' to be non-null");
            return $;
        }
    }

}
