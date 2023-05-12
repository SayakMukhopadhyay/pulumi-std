# coding=utf-8
# *** WARNING: this file was generated by pulumi. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'LengthResult',
    'AwaitableLengthResult',
    'length',
    'length_output',
]

@pulumi.output_type
class LengthResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, int):
            raise TypeError("Expected argument 'result' to be a int")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> int:
        return pulumi.get(self, "result")


class AwaitableLengthResult(LengthResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return LengthResult(
            result=self.result)


def length(input: Optional[Any] = None,
           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableLengthResult:
    """
    Determines the length of a given list, map, or string.
    """
    __args__ = dict()
    __args__['input'] = input
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:length', __args__, opts=opts, typ=LengthResult).value

    return AwaitableLengthResult(
        result=__ret__.result)


@_utilities.lift_output_func(length)
def length_output(input: Optional[Any] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[LengthResult]:
    """
    Determines the length of a given list, map, or string.
    """
    ...