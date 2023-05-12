// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class SortPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final SortPlainArgs Empty = new SortPlainArgs();

    @Import(name="input", required=true)
    private List<String> input;

    public List<String> input() {
        return this.input;
    }

    private SortPlainArgs() {}

    private SortPlainArgs(SortPlainArgs $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SortPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SortPlainArgs $;

        public Builder() {
            $ = new SortPlainArgs();
        }

        public Builder(SortPlainArgs defaults) {
            $ = new SortPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(List<String> input) {
            $.input = input;
            return this;
        }

        public Builder input(String... input) {
            return input(List.of(input));
        }

        public SortPlainArgs build() {
            $.input = Objects.requireNonNull($.input, "expected parameter 'input' to be non-null");
            return $;
        }
    }

}