# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'ParseintResult',
    'AwaitableParseintResult',
    'parseint',
    'parseint_output',
]

@pulumi.output_type
class ParseintResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, int):
            raise TypeError("Expected argument 'result' to be a int")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> int:
        return pulumi.get(self, "result")


class AwaitableParseintResult(ParseintResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ParseintResult(
            result=self.result)


def parseint(base: Optional[int] = None,
             input: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableParseintResult:
    """
    Parses the given string as a representation of an integer in the specified base
    and returns the resulting number. The base must be between 2 and 62 inclusive.
    	.
    """
    __args__ = dict()
    __args__['base'] = base
    __args__['input'] = input
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:parseint', __args__, opts=opts, typ=ParseintResult).value

    return AwaitableParseintResult(
        result=pulumi.get(__ret__, 'result'))


@_utilities.lift_output_func(parseint)
def parseint_output(base: Optional[pulumi.Input[Optional[int]]] = None,
                    input: Optional[pulumi.Input[str]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ParseintResult]:
    """
    Parses the given string as a representation of an integer in the specified base
    and returns the resulting number. The base must be between 2 and 62 inclusive.
    	.
    """
    ...
