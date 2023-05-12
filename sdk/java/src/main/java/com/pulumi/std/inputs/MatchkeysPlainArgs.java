// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class MatchkeysPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final MatchkeysPlainArgs Empty = new MatchkeysPlainArgs();

    @Import(name="searchList", required=true)
    private List<String> searchList;

    public List<String> searchList() {
        return this.searchList;
    }

    @Import(name="values", required=true)
    private List<String> values;

    public List<String> values() {
        return this.values;
    }

    private MatchkeysPlainArgs() {}

    private MatchkeysPlainArgs(MatchkeysPlainArgs $) {
        this.searchList = $.searchList;
        this.values = $.values;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MatchkeysPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MatchkeysPlainArgs $;

        public Builder() {
            $ = new MatchkeysPlainArgs();
        }

        public Builder(MatchkeysPlainArgs defaults) {
            $ = new MatchkeysPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder searchList(List<String> searchList) {
            $.searchList = searchList;
            return this;
        }

        public Builder searchList(String... searchList) {
            return searchList(List.of(searchList));
        }

        public Builder values(List<String> values) {
            $.values = values;
            return this;
        }

        public Builder values(String... values) {
            return values(List.of(values));
        }

        public MatchkeysPlainArgs build() {
            $.searchList = Objects.requireNonNull($.searchList, "expected parameter 'searchList' to be non-null");
            $.values = Objects.requireNonNull($.values, "expected parameter 'values' to be non-null");
            return $;
        }
    }

}