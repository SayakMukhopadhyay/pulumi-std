// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Object;
import java.util.Objects;


public final class JsonencodeArgs extends com.pulumi.resources.InvokeArgs {

    public static final JsonencodeArgs Empty = new JsonencodeArgs();

    @Import(name="input", required=true)
    private Output<Object> input;

    public Output<Object> input() {
        return this.input;
    }

    private JsonencodeArgs() {}

    private JsonencodeArgs(JsonencodeArgs $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(JsonencodeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private JsonencodeArgs $;

        public Builder() {
            $ = new JsonencodeArgs();
        }

        public Builder(JsonencodeArgs defaults) {
            $ = new JsonencodeArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(Output<Object> input) {
            $.input = input;
            return this;
        }

        public Builder input(Object input) {
            return input(Output.of(input));
        }

        public JsonencodeArgs build() {
            $.input = Objects.requireNonNull($.input, "expected parameter 'input' to be non-null");
            return $;
        }
    }

}