// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class BasenamePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final BasenamePlainArgs Empty = new BasenamePlainArgs();

    @Import(name="input", required=true)
    private String input;

    public String input() {
        return this.input;
    }

    private BasenamePlainArgs() {}

    private BasenamePlainArgs(BasenamePlainArgs $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BasenamePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BasenamePlainArgs $;

        public Builder() {
            $ = new BasenamePlainArgs();
        }

        public Builder(BasenamePlainArgs defaults) {
            $ = new BasenamePlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(String input) {
            $.input = input;
            return this;
        }

        public BasenamePlainArgs build() {
            $.input = Objects.requireNonNull($.input, "expected parameter 'input' to be non-null");
            return $;
        }
    }

}