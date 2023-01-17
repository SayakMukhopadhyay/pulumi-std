// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class TimeaddArgs extends com.pulumi.resources.InvokeArgs {

    public static final TimeaddArgs Empty = new TimeaddArgs();

    @Import(name="duration", required=true)
    private Output<String> duration;

    public Output<String> duration() {
        return this.duration;
    }

    @Import(name="timestamp", required=true)
    private Output<String> timestamp;

    public Output<String> timestamp() {
        return this.timestamp;
    }

    private TimeaddArgs() {}

    private TimeaddArgs(TimeaddArgs $) {
        this.duration = $.duration;
        this.timestamp = $.timestamp;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TimeaddArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TimeaddArgs $;

        public Builder() {
            $ = new TimeaddArgs();
        }

        public Builder(TimeaddArgs defaults) {
            $ = new TimeaddArgs(Objects.requireNonNull(defaults));
        }

        public Builder duration(Output<String> duration) {
            $.duration = duration;
            return this;
        }

        public Builder duration(String duration) {
            return duration(Output.of(duration));
        }

        public Builder timestamp(Output<String> timestamp) {
            $.timestamp = timestamp;
            return this;
        }

        public Builder timestamp(String timestamp) {
            return timestamp(Output.of(timestamp));
        }

        public TimeaddArgs build() {
            $.duration = Objects.requireNonNull($.duration, "expected parameter 'duration' to be non-null");
            $.timestamp = Objects.requireNonNull($.timestamp, "expected parameter 'timestamp' to be non-null");
            return $;
        }
    }

}
