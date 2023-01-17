// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;


public final class ValuesArgs extends com.pulumi.resources.InvokeArgs {

    public static final ValuesArgs Empty = new ValuesArgs();

    @Import(name="input", required=true)
    private Output<Map<String,Object>> input;

    public Output<Map<String,Object>> input() {
        return this.input;
    }

    private ValuesArgs() {}

    private ValuesArgs(ValuesArgs $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ValuesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ValuesArgs $;

        public Builder() {
            $ = new ValuesArgs();
        }

        public Builder(ValuesArgs defaults) {
            $ = new ValuesArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(Output<Map<String,Object>> input) {
            $.input = input;
            return this;
        }

        public Builder input(Map<String,Object> input) {
            return input(Output.of(input));
        }

        public ValuesArgs build() {
            $.input = Objects.requireNonNull($.input, "expected parameter 'input' to be non-null");
            return $;
        }
    }

}
