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
    'Sha512Result',
    'AwaitableSha512Result',
    'sha512',
    'sha512_output',
]

@pulumi.output_type
class Sha512Result:
    def __init__(__self__, result=None):
        if result and not isinstance(result, str):
            raise TypeError("Expected argument 'result' to be a str")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> str:
        return pulumi.get(self, "result")


class AwaitableSha512Result(Sha512Result):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return Sha512Result(
            result=self.result)


def sha512(input: Optional[str] = None,
           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableSha512Result:
    """
    Returns a hexadecimal representation of the SHA-512 hash of the given string.
    """
    __args__ = dict()
    __args__['input'] = input
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:sha512', __args__, opts=opts, typ=Sha512Result).value

    return AwaitableSha512Result(
        result=pulumi.get(__ret__, 'result'))


@_utilities.lift_output_func(sha512)
def sha512_output(input: Optional[pulumi.Input[str]] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[Sha512Result]:
    """
    Returns a hexadecimal representation of the SHA-512 hash of the given string.
    """
    ...
