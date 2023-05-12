// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Double;
import java.util.Objects;


public final class PowPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final PowPlainArgs Empty = new PowPlainArgs();

    @Import(name="base", required=true)
    private Double base;

    public Double base() {
        return this.base;
    }

    @Import(name="exponent", required=true)
    private Double exponent;

    public Double exponent() {
        return this.exponent;
    }

    private PowPlainArgs() {}

    private PowPlainArgs(PowPlainArgs $) {
        this.base = $.base;
        this.exponent = $.exponent;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PowPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PowPlainArgs $;

        public Builder() {
            $ = new PowPlainArgs();
        }

        public Builder(PowPlainArgs defaults) {
            $ = new PowPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder base(Double base) {
            $.base = base;
            return this;
        }

        public Builder exponent(Double exponent) {
            $.exponent = exponent;
            return this;
        }

        public PowPlainArgs build() {
            $.base = Objects.requireNonNull($.base, "expected parameter 'base' to be non-null");
            $.exponent = Objects.requireNonNull($.exponent, "expected parameter 'exponent' to be non-null");
            return $;
        }
    }

}