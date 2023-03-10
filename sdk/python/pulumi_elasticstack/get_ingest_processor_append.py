# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetIngestProcessorAppendResult',
    'AwaitableGetIngestProcessorAppendResult',
    'get_ingest_processor_append',
    'get_ingest_processor_append_output',
]

@pulumi.output_type
class GetIngestProcessorAppendResult:
    """
    A collection of values returned by getIngestProcessorAppend.
    """
    def __init__(__self__, allow_duplicates=None, description=None, field=None, id=None, if_=None, ignore_failure=None, json=None, media_type=None, on_failures=None, tag=None, values=None):
        if allow_duplicates and not isinstance(allow_duplicates, bool):
            raise TypeError("Expected argument 'allow_duplicates' to be a bool")
        pulumi.set(__self__, "allow_duplicates", allow_duplicates)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if field and not isinstance(field, str):
            raise TypeError("Expected argument 'field' to be a str")
        pulumi.set(__self__, "field", field)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if if_ and not isinstance(if_, str):
            raise TypeError("Expected argument 'if_' to be a str")
        pulumi.set(__self__, "if_", if_)
        if ignore_failure and not isinstance(ignore_failure, bool):
            raise TypeError("Expected argument 'ignore_failure' to be a bool")
        pulumi.set(__self__, "ignore_failure", ignore_failure)
        if json and not isinstance(json, str):
            raise TypeError("Expected argument 'json' to be a str")
        pulumi.set(__self__, "json", json)
        if media_type and not isinstance(media_type, str):
            raise TypeError("Expected argument 'media_type' to be a str")
        pulumi.set(__self__, "media_type", media_type)
        if on_failures and not isinstance(on_failures, list):
            raise TypeError("Expected argument 'on_failures' to be a list")
        pulumi.set(__self__, "on_failures", on_failures)
        if tag and not isinstance(tag, str):
            raise TypeError("Expected argument 'tag' to be a str")
        pulumi.set(__self__, "tag", tag)
        if values and not isinstance(values, list):
            raise TypeError("Expected argument 'values' to be a list")
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter(name="allowDuplicates")
    def allow_duplicates(self) -> Optional[bool]:
        """
        If `false`, the processor does not append values already present in the field.
        """
        return pulumi.get(self, "allow_duplicates")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Description of the processor.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def field(self) -> str:
        """
        The field to be appended to.
        """
        return pulumi.get(self, "field")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Internal identifier of the resource
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="if")
    def if_(self) -> Optional[str]:
        """
        Conditionally execute the processor
        """
        return pulumi.get(self, "if_")

    @property
    @pulumi.getter(name="ignoreFailure")
    def ignore_failure(self) -> Optional[bool]:
        """
        Ignore failures for the processor.
        """
        return pulumi.get(self, "ignore_failure")

    @property
    @pulumi.getter
    def json(self) -> str:
        """
        JSON representation of this data source.
        """
        return pulumi.get(self, "json")

    @property
    @pulumi.getter(name="mediaType")
    def media_type(self) -> Optional[str]:
        """
        The media type for encoding value. Applies only when value is a template snippet. Must be one of `application/json`, `text/plain`, or `application/x-www-form-urlencoded`. Supported only from Elasticsearch version **7.15**.
        """
        return pulumi.get(self, "media_type")

    @property
    @pulumi.getter(name="onFailures")
    def on_failures(self) -> Optional[Sequence[str]]:
        """
        Handle failures for the processor.
        """
        return pulumi.get(self, "on_failures")

    @property
    @pulumi.getter
    def tag(self) -> Optional[str]:
        """
        Identifier for the processor.
        """
        return pulumi.get(self, "tag")

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        """
        The value to be appended.
        """
        return pulumi.get(self, "values")


class AwaitableGetIngestProcessorAppendResult(GetIngestProcessorAppendResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIngestProcessorAppendResult(
            allow_duplicates=self.allow_duplicates,
            description=self.description,
            field=self.field,
            id=self.id,
            if_=self.if_,
            ignore_failure=self.ignore_failure,
            json=self.json,
            media_type=self.media_type,
            on_failures=self.on_failures,
            tag=self.tag,
            values=self.values)


def get_ingest_processor_append(allow_duplicates: Optional[bool] = None,
                                description: Optional[str] = None,
                                field: Optional[str] = None,
                                if_: Optional[str] = None,
                                ignore_failure: Optional[bool] = None,
                                media_type: Optional[str] = None,
                                on_failures: Optional[Sequence[str]] = None,
                                tag: Optional[str] = None,
                                values: Optional[Sequence[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIngestProcessorAppendResult:
    """
    Helper data source to which can be used to create a processor to append one or more values to an existing array if the field already exists and it is an array.
    Converts a scalar to an array and appends one or more values to it if the field exists and it is a scalar. Creates an array containing the provided values if the field doesn???t exist.

    See: https://www.elastic.co/guide/en/elasticsearch/reference/current/append-processor.html

    ## Example Usage

    ```python
    import pulumi
    import pulumi_elasticstack as elasticstack

    tags = elasticstack.get_ingest_processor_append(field="tags",
        values=[
            "production",
            "{{{app}}}",
            "{{{owner}}}",
        ])
    my_ingest_pipeline = elasticstack.IngestPipeline("myIngestPipeline", processors=[tags.json])
    ```


    :param bool allow_duplicates: If `false`, the processor does not append values already present in the field.
    :param str description: Description of the processor.
    :param str field: The field to be appended to.
    :param str if_: Conditionally execute the processor
    :param bool ignore_failure: Ignore failures for the processor.
    :param str media_type: The media type for encoding value. Applies only when value is a template snippet. Must be one of `application/json`, `text/plain`, or `application/x-www-form-urlencoded`. Supported only from Elasticsearch version **7.15**.
    :param Sequence[str] on_failures: Handle failures for the processor.
    :param str tag: Identifier for the processor.
    :param Sequence[str] values: The value to be appended.
    """
    __args__ = dict()
    __args__['allowDuplicates'] = allow_duplicates
    __args__['description'] = description
    __args__['field'] = field
    __args__['if'] = if_
    __args__['ignoreFailure'] = ignore_failure
    __args__['mediaType'] = media_type
    __args__['onFailures'] = on_failures
    __args__['tag'] = tag
    __args__['values'] = values
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('elasticstack:index/getIngestProcessorAppend:getIngestProcessorAppend', __args__, opts=opts, typ=GetIngestProcessorAppendResult).value

    return AwaitableGetIngestProcessorAppendResult(
        allow_duplicates=__ret__.allow_duplicates,
        description=__ret__.description,
        field=__ret__.field,
        id=__ret__.id,
        if_=__ret__.if_,
        ignore_failure=__ret__.ignore_failure,
        json=__ret__.json,
        media_type=__ret__.media_type,
        on_failures=__ret__.on_failures,
        tag=__ret__.tag,
        values=__ret__.values)


@_utilities.lift_output_func(get_ingest_processor_append)
def get_ingest_processor_append_output(allow_duplicates: Optional[pulumi.Input[Optional[bool]]] = None,
                                       description: Optional[pulumi.Input[Optional[str]]] = None,
                                       field: Optional[pulumi.Input[str]] = None,
                                       if_: Optional[pulumi.Input[Optional[str]]] = None,
                                       ignore_failure: Optional[pulumi.Input[Optional[bool]]] = None,
                                       media_type: Optional[pulumi.Input[Optional[str]]] = None,
                                       on_failures: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                       tag: Optional[pulumi.Input[Optional[str]]] = None,
                                       values: Optional[pulumi.Input[Sequence[str]]] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIngestProcessorAppendResult]:
    """
    Helper data source to which can be used to create a processor to append one or more values to an existing array if the field already exists and it is an array.
    Converts a scalar to an array and appends one or more values to it if the field exists and it is a scalar. Creates an array containing the provided values if the field doesn???t exist.

    See: https://www.elastic.co/guide/en/elasticsearch/reference/current/append-processor.html

    ## Example Usage

    ```python
    import pulumi
    import pulumi_elasticstack as elasticstack

    tags = elasticstack.get_ingest_processor_append(field="tags",
        values=[
            "production",
            "{{{app}}}",
            "{{{owner}}}",
        ])
    my_ingest_pipeline = elasticstack.IngestPipeline("myIngestPipeline", processors=[tags.json])
    ```


    :param bool allow_duplicates: If `false`, the processor does not append values already present in the field.
    :param str description: Description of the processor.
    :param str field: The field to be appended to.
    :param str if_: Conditionally execute the processor
    :param bool ignore_failure: Ignore failures for the processor.
    :param str media_type: The media type for encoding value. Applies only when value is a template snippet. Must be one of `application/json`, `text/plain`, or `application/x-www-form-urlencoded`. Supported only from Elasticsearch version **7.15**.
    :param Sequence[str] on_failures: Handle failures for the processor.
    :param str tag: Identifier for the processor.
    :param Sequence[str] values: The value to be appended.
    """
    ...
