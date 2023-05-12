// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;


public final class MergePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final MergePlainArgs Empty = new MergePlainArgs();

    @Import(name="input", required=true)
    private List<Map<String,Object>> input;

    public List<Map<String,Object>> input() {
        return this.input;
    }

    private MergePlainArgs() {}

    private MergePlainArgs(MergePlainArgs $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MergePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MergePlainArgs $;

        public Builder() {
            $ = new MergePlainArgs();
        }

        public Builder(MergePlainArgs defaults) {
            $ = new MergePlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(List<Map<String,Object>> input) {
            $.input = input;
            return this;
        }

        public Builder input(Map<String,Object>... input) {
            return input(List.of(input));
        }

        public MergePlainArgs build() {
            $.input = Objects.requireNonNull($.input, "expected parameter 'input' to be non-null");
            return $;
        }
    }

}