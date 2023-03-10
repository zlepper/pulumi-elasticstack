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
    'GetIngestProcessorCsvResult',
    'AwaitableGetIngestProcessorCsvResult',
    'get_ingest_processor_csv',
    'get_ingest_processor_csv_output',
]

@pulumi.output_type
class GetIngestProcessorCsvResult:
    """
    A collection of values returned by getIngestProcessorCsv.
    """
    def __init__(__self__, description=None, empty_value=None, field=None, id=None, if_=None, ignore_failure=None, ignore_missing=None, json=None, on_failures=None, quote=None, separator=None, tag=None, target_fields=None, trim=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if empty_value and not isinstance(empty_value, str):
            raise TypeError("Expected argument 'empty_value' to be a str")
        pulumi.set(__self__, "empty_value", empty_value)
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
        if ignore_missing and not isinstance(ignore_missing, bool):
            raise TypeError("Expected argument 'ignore_missing' to be a bool")
        pulumi.set(__self__, "ignore_missing", ignore_missing)
        if json and not isinstance(json, str):
            raise TypeError("Expected argument 'json' to be a str")
        pulumi.set(__self__, "json", json)
        if on_failures and not isinstance(on_failures, list):
            raise TypeError("Expected argument 'on_failures' to be a list")
        pulumi.set(__self__, "on_failures", on_failures)
        if quote and not isinstance(quote, str):
            raise TypeError("Expected argument 'quote' to be a str")
        pulumi.set(__self__, "quote", quote)
        if separator and not isinstance(separator, str):
            raise TypeError("Expected argument 'separator' to be a str")
        pulumi.set(__self__, "separator", separator)
        if tag and not isinstance(tag, str):
            raise TypeError("Expected argument 'tag' to be a str")
        pulumi.set(__self__, "tag", tag)
        if target_fields and not isinstance(target_fields, list):
            raise TypeError("Expected argument 'target_fields' to be a list")
        pulumi.set(__self__, "target_fields", target_fields)
        if trim and not isinstance(trim, bool):
            raise TypeError("Expected argument 'trim' to be a bool")
        pulumi.set(__self__, "trim", trim)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Description of the processor.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="emptyValue")
    def empty_value(self) -> Optional[str]:
        """
        Value used to fill empty fields, empty fields will be skipped if this is not provided.
        """
        return pulumi.get(self, "empty_value")

    @property
    @pulumi.getter
    def field(self) -> str:
        """
        The field to extract data from.
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
    @pulumi.getter(name="ignoreMissing")
    def ignore_missing(self) -> Optional[bool]:
        """
        If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
        """
        return pulumi.get(self, "ignore_missing")

    @property
    @pulumi.getter
    def json(self) -> str:
        """
        JSON representation of this data source.
        """
        return pulumi.get(self, "json")

    @property
    @pulumi.getter(name="onFailures")
    def on_failures(self) -> Optional[Sequence[str]]:
        """
        Handle failures for the processor.
        """
        return pulumi.get(self, "on_failures")

    @property
    @pulumi.getter
    def quote(self) -> Optional[str]:
        """
        Quote used in CSV, has to be single character string
        """
        return pulumi.get(self, "quote")

    @property
    @pulumi.getter
    def separator(self) -> Optional[str]:
        """
        Separator used in CSV, has to be single character string.
        """
        return pulumi.get(self, "separator")

    @property
    @pulumi.getter
    def tag(self) -> Optional[str]:
        """
        Identifier for the processor.
        """
        return pulumi.get(self, "tag")

    @property
    @pulumi.getter(name="targetFields")
    def target_fields(self) -> Sequence[str]:
        """
        The array of fields to assign extracted values to.
        """
        return pulumi.get(self, "target_fields")

    @property
    @pulumi.getter
    def trim(self) -> Optional[bool]:
        """
        Trim whitespaces in unquoted fields.
        """
        return pulumi.get(self, "trim")


class AwaitableGetIngestProcessorCsvResult(GetIngestProcessorCsvResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIngestProcessorCsvResult(
            description=self.description,
            empty_value=self.empty_value,
            field=self.field,
            id=self.id,
            if_=self.if_,
            ignore_failure=self.ignore_failure,
            ignore_missing=self.ignore_missing,
            json=self.json,
            on_failures=self.on_failures,
            quote=self.quote,
            separator=self.separator,
            tag=self.tag,
            target_fields=self.target_fields,
            trim=self.trim)


def get_ingest_processor_csv(description: Optional[str] = None,
                             empty_value: Optional[str] = None,
                             field: Optional[str] = None,
                             if_: Optional[str] = None,
                             ignore_failure: Optional[bool] = None,
                             ignore_missing: Optional[bool] = None,
                             on_failures: Optional[Sequence[str]] = None,
                             quote: Optional[str] = None,
                             separator: Optional[str] = None,
                             tag: Optional[str] = None,
                             target_fields: Optional[Sequence[str]] = None,
                             trim: Optional[bool] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIngestProcessorCsvResult:
    """
    Helper data source to which can be used to extract fields from CSV line out of a single text field within a document. Any empty field in CSV will be skipped.

    See: https://www.elastic.co/guide/en/elasticsearch/reference/current/csv-processor.html

    ## Example Usage

    ```python
    import pulumi
    import pulumi_elasticstack as elasticstack

    csv = elasticstack.get_ingest_processor_csv(field="my_field",
        target_fields=[
            "field1",
            "field2",
        ])
    my_ingest_pipeline = elasticstack.IngestPipeline("myIngestPipeline", processors=[csv.json])
    ```

    If the `trim` option is enabled then any whitespace in the beginning and in the end of each unquoted field will be trimmed. For example with configuration above, a value of A, B will result in field field2 having value {nbsp}B (with space at the beginning). If trim is enabled A, B will result in field field2 having value B (no whitespace). Quoted fields will be left untouched.


    :param str description: Description of the processor.
    :param str empty_value: Value used to fill empty fields, empty fields will be skipped if this is not provided.
    :param str field: The field to extract data from.
    :param str if_: Conditionally execute the processor
    :param bool ignore_failure: Ignore failures for the processor.
    :param bool ignore_missing: If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
    :param Sequence[str] on_failures: Handle failures for the processor.
    :param str quote: Quote used in CSV, has to be single character string
    :param str separator: Separator used in CSV, has to be single character string.
    :param str tag: Identifier for the processor.
    :param Sequence[str] target_fields: The array of fields to assign extracted values to.
    :param bool trim: Trim whitespaces in unquoted fields.
    """
    __args__ = dict()
    __args__['description'] = description
    __args__['emptyValue'] = empty_value
    __args__['field'] = field
    __args__['if'] = if_
    __args__['ignoreFailure'] = ignore_failure
    __args__['ignoreMissing'] = ignore_missing
    __args__['onFailures'] = on_failures
    __args__['quote'] = quote
    __args__['separator'] = separator
    __args__['tag'] = tag
    __args__['targetFields'] = target_fields
    __args__['trim'] = trim
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('elasticstack:index/getIngestProcessorCsv:getIngestProcessorCsv', __args__, opts=opts, typ=GetIngestProcessorCsvResult).value

    return AwaitableGetIngestProcessorCsvResult(
        description=__ret__.description,
        empty_value=__ret__.empty_value,
        field=__ret__.field,
        id=__ret__.id,
        if_=__ret__.if_,
        ignore_failure=__ret__.ignore_failure,
        ignore_missing=__ret__.ignore_missing,
        json=__ret__.json,
        on_failures=__ret__.on_failures,
        quote=__ret__.quote,
        separator=__ret__.separator,
        tag=__ret__.tag,
        target_fields=__ret__.target_fields,
        trim=__ret__.trim)


@_utilities.lift_output_func(get_ingest_processor_csv)
def get_ingest_processor_csv_output(description: Optional[pulumi.Input[Optional[str]]] = None,
                                    empty_value: Optional[pulumi.Input[Optional[str]]] = None,
                                    field: Optional[pulumi.Input[str]] = None,
                                    if_: Optional[pulumi.Input[Optional[str]]] = None,
                                    ignore_failure: Optional[pulumi.Input[Optional[bool]]] = None,
                                    ignore_missing: Optional[pulumi.Input[Optional[bool]]] = None,
                                    on_failures: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                    quote: Optional[pulumi.Input[Optional[str]]] = None,
                                    separator: Optional[pulumi.Input[Optional[str]]] = None,
                                    tag: Optional[pulumi.Input[Optional[str]]] = None,
                                    target_fields: Optional[pulumi.Input[Sequence[str]]] = None,
                                    trim: Optional[pulumi.Input[Optional[bool]]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIngestProcessorCsvResult]:
    """
    Helper data source to which can be used to extract fields from CSV line out of a single text field within a document. Any empty field in CSV will be skipped.

    See: https://www.elastic.co/guide/en/elasticsearch/reference/current/csv-processor.html

    ## Example Usage

    ```python
    import pulumi
    import pulumi_elasticstack as elasticstack

    csv = elasticstack.get_ingest_processor_csv(field="my_field",
        target_fields=[
            "field1",
            "field2",
        ])
    my_ingest_pipeline = elasticstack.IngestPipeline("myIngestPipeline", processors=[csv.json])
    ```

    If the `trim` option is enabled then any whitespace in the beginning and in the end of each unquoted field will be trimmed. For example with configuration above, a value of A, B will result in field field2 having value {nbsp}B (with space at the beginning). If trim is enabled A, B will result in field field2 having value B (no whitespace). Quoted fields will be left untouched.


    :param str description: Description of the processor.
    :param str empty_value: Value used to fill empty fields, empty fields will be skipped if this is not provided.
    :param str field: The field to extract data from.
    :param str if_: Conditionally execute the processor
    :param bool ignore_failure: Ignore failures for the processor.
    :param bool ignore_missing: If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
    :param Sequence[str] on_failures: Handle failures for the processor.
    :param str quote: Quote used in CSV, has to be single character string
    :param str separator: Separator used in CSV, has to be single character string.
    :param str tag: Identifier for the processor.
    :param Sequence[str] target_fields: The array of fields to assign extracted values to.
    :param bool trim: Trim whitespaces in unquoted fields.
    """
    ...
