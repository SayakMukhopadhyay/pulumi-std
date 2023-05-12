// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Object;
import java.util.List;
import java.util.Objects;


public final class MapPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final MapPlainArgs Empty = new MapPlainArgs();

    @Import(name="args", required=true)
    private List<Object> args;

    public List<Object> args() {
        return this.args;
    }

    private MapPlainArgs() {}

    private MapPlainArgs(MapPlainArgs $) {
        this.args = $.args;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MapPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MapPlainArgs $;

        public Builder() {
            $ = new MapPlainArgs();
        }

        public Builder(MapPlainArgs defaults) {
            $ = new MapPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder args(List<Object> args) {
            $.args = args;
            return this;
        }

        public Builder args(Object... args) {
            return args(List.of(args));
        }

        public MapPlainArgs build() {
            $.args = Objects.requireNonNull($.args, "expected parameter 'args' to be non-null");
            return $;
        }
    }

}