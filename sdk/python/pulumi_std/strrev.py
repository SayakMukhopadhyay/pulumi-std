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
    'StrrevResult',
    'AwaitableStrrevResult',
    'strrev',
    'strrev_output',
]

@pulumi.output_type
class StrrevResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, str):
            raise TypeError("Expected argument 'result' to be a str")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> str:
        return pulumi.get(self, "result")


class AwaitableStrrevResult(StrrevResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return StrrevResult(
            result=self.result)


def strrev(input: Optional[str] = None,
           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableStrrevResult:
    """
    Returns the given string with all of its Unicode characters in reverse order.
    """
    __args__ = dict()
    __args__['input'] = input
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:strrev', __args__, opts=opts, typ=StrrevResult).value

    return AwaitableStrrevResult(
        result=pulumi.get(__ret__, 'result'))


@_utilities.lift_output_func(strrev)
def strrev_output(input: Optional[pulumi.Input[str]] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[StrrevResult]:
    """
    Returns the given string with all of its Unicode characters in reverse order.
    """
    ...
