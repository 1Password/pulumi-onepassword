# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = [
    'ItemPasswordRecipeArgs',
    'ItemPasswordRecipeArgsDict',
    'ItemSectionArgs',
    'ItemSectionArgsDict',
    'ItemSectionFieldArgs',
    'ItemSectionFieldArgsDict',
    'ItemSectionFieldPasswordRecipeArgs',
    'ItemSectionFieldPasswordRecipeArgsDict',
    'GetItemFileArgs',
    'GetItemFileArgsDict',
    'GetItemSectionArgs',
    'GetItemSectionArgsDict',
    'GetItemSectionFieldArgs',
    'GetItemSectionFieldArgsDict',
    'GetItemSectionFileArgs',
    'GetItemSectionFileArgsDict',
]

MYPY = False

if not MYPY:
    class ItemPasswordRecipeArgsDict(TypedDict):
        digits: NotRequired[pulumi.Input[bool]]
        """
        Use digits [0-9] when generating the password.
        """
        length: NotRequired[pulumi.Input[int]]
        """
        The length of the password to be generated.
        """
        letters: NotRequired[pulumi.Input[bool]]
        """
        Use letters [a-zA-Z] when generating the password.
        """
        symbols: NotRequired[pulumi.Input[bool]]
        """
        Use symbols [!@.-_*] when generating the password.
        """
elif False:
    ItemPasswordRecipeArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ItemPasswordRecipeArgs:
    def __init__(__self__, *,
                 digits: Optional[pulumi.Input[bool]] = None,
                 length: Optional[pulumi.Input[int]] = None,
                 letters: Optional[pulumi.Input[bool]] = None,
                 symbols: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[bool] digits: Use digits [0-9] when generating the password.
        :param pulumi.Input[int] length: The length of the password to be generated.
        :param pulumi.Input[bool] letters: Use letters [a-zA-Z] when generating the password.
        :param pulumi.Input[bool] symbols: Use symbols [!@.-_*] when generating the password.
        """
        if digits is not None:
            pulumi.set(__self__, "digits", digits)
        if length is not None:
            pulumi.set(__self__, "length", length)
        if letters is not None:
            pulumi.set(__self__, "letters", letters)
        if symbols is not None:
            pulumi.set(__self__, "symbols", symbols)

    @property
    @pulumi.getter
    def digits(self) -> Optional[pulumi.Input[bool]]:
        """
        Use digits [0-9] when generating the password.
        """
        return pulumi.get(self, "digits")

    @digits.setter
    def digits(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "digits", value)

    @property
    @pulumi.getter
    def length(self) -> Optional[pulumi.Input[int]]:
        """
        The length of the password to be generated.
        """
        return pulumi.get(self, "length")

    @length.setter
    def length(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "length", value)

    @property
    @pulumi.getter
    def letters(self) -> Optional[pulumi.Input[bool]]:
        """
        Use letters [a-zA-Z] when generating the password.
        """
        return pulumi.get(self, "letters")

    @letters.setter
    def letters(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "letters", value)

    @property
    @pulumi.getter
    def symbols(self) -> Optional[pulumi.Input[bool]]:
        """
        Use symbols [!@.-_*] when generating the password.
        """
        return pulumi.get(self, "symbols")

    @symbols.setter
    def symbols(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "symbols", value)


if not MYPY:
    class ItemSectionArgsDict(TypedDict):
        label: pulumi.Input[str]
        """
        The label for the section.
        """
        fields: NotRequired[pulumi.Input[Sequence[pulumi.Input['ItemSectionFieldArgsDict']]]]
        """
        A list of custom fields in the section.
        """
        id: NotRequired[pulumi.Input[str]]
        """
        A unique identifier for the section.
        """
elif False:
    ItemSectionArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ItemSectionArgs:
    def __init__(__self__, *,
                 label: pulumi.Input[str],
                 fields: Optional[pulumi.Input[Sequence[pulumi.Input['ItemSectionFieldArgs']]]] = None,
                 id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] label: The label for the section.
        :param pulumi.Input[Sequence[pulumi.Input['ItemSectionFieldArgs']]] fields: A list of custom fields in the section.
        :param pulumi.Input[str] id: A unique identifier for the section.
        """
        pulumi.set(__self__, "label", label)
        if fields is not None:
            pulumi.set(__self__, "fields", fields)
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def label(self) -> pulumi.Input[str]:
        """
        The label for the section.
        """
        return pulumi.get(self, "label")

    @label.setter
    def label(self, value: pulumi.Input[str]):
        pulumi.set(self, "label", value)

    @property
    @pulumi.getter
    def fields(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ItemSectionFieldArgs']]]]:
        """
        A list of custom fields in the section.
        """
        return pulumi.get(self, "fields")

    @fields.setter
    def fields(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ItemSectionFieldArgs']]]]):
        pulumi.set(self, "fields", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        A unique identifier for the section.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)


if not MYPY:
    class ItemSectionFieldArgsDict(TypedDict):
        label: pulumi.Input[str]
        """
        The label for the field.
        """
        id: NotRequired[pulumi.Input[str]]
        """
        A unique identifier for the field.
        """
        password_recipe: NotRequired[pulumi.Input['ItemSectionFieldPasswordRecipeArgsDict']]
        """
        The recipe used to generate a new value for a password.
        """
        purpose: NotRequired[pulumi.Input[str]]
        """
        Purpose indicates this is a special field: a username, password, or notes field. One of ["USERNAME" "PASSWORD" "NOTES"]
        """
        type: NotRequired[pulumi.Input[str]]
        """
        The type of value stored in the field. One of ["STRING" "CONCEALED" "EMAIL" "URL" "OTP" "DATE" "MONTH_YEAR" "MENU"]
        """
        value: NotRequired[pulumi.Input[str]]
        """
        The value of the field.
        """
elif False:
    ItemSectionFieldArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ItemSectionFieldArgs:
    def __init__(__self__, *,
                 label: pulumi.Input[str],
                 id: Optional[pulumi.Input[str]] = None,
                 password_recipe: Optional[pulumi.Input['ItemSectionFieldPasswordRecipeArgs']] = None,
                 purpose: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] label: The label for the field.
        :param pulumi.Input[str] id: A unique identifier for the field.
        :param pulumi.Input['ItemSectionFieldPasswordRecipeArgs'] password_recipe: The recipe used to generate a new value for a password.
        :param pulumi.Input[str] purpose: Purpose indicates this is a special field: a username, password, or notes field. One of ["USERNAME" "PASSWORD" "NOTES"]
        :param pulumi.Input[str] type: The type of value stored in the field. One of ["STRING" "CONCEALED" "EMAIL" "URL" "OTP" "DATE" "MONTH_YEAR" "MENU"]
        :param pulumi.Input[str] value: The value of the field.
        """
        pulumi.set(__self__, "label", label)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if password_recipe is not None:
            pulumi.set(__self__, "password_recipe", password_recipe)
        if purpose is not None:
            pulumi.set(__self__, "purpose", purpose)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def label(self) -> pulumi.Input[str]:
        """
        The label for the field.
        """
        return pulumi.get(self, "label")

    @label.setter
    def label(self, value: pulumi.Input[str]):
        pulumi.set(self, "label", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        A unique identifier for the field.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="passwordRecipe")
    def password_recipe(self) -> Optional[pulumi.Input['ItemSectionFieldPasswordRecipeArgs']]:
        """
        The recipe used to generate a new value for a password.
        """
        return pulumi.get(self, "password_recipe")

    @password_recipe.setter
    def password_recipe(self, value: Optional[pulumi.Input['ItemSectionFieldPasswordRecipeArgs']]):
        pulumi.set(self, "password_recipe", value)

    @property
    @pulumi.getter
    def purpose(self) -> Optional[pulumi.Input[str]]:
        """
        Purpose indicates this is a special field: a username, password, or notes field. One of ["USERNAME" "PASSWORD" "NOTES"]
        """
        return pulumi.get(self, "purpose")

    @purpose.setter
    def purpose(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "purpose", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of value stored in the field. One of ["STRING" "CONCEALED" "EMAIL" "URL" "OTP" "DATE" "MONTH_YEAR" "MENU"]
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        The value of the field.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


if not MYPY:
    class ItemSectionFieldPasswordRecipeArgsDict(TypedDict):
        digits: NotRequired[pulumi.Input[bool]]
        """
        Use digits [0-9] when generating the password.
        """
        length: NotRequired[pulumi.Input[int]]
        """
        The length of the password to be generated.
        """
        letters: NotRequired[pulumi.Input[bool]]
        """
        Use letters [a-zA-Z] when generating the password.
        """
        symbols: NotRequired[pulumi.Input[bool]]
        """
        Use symbols [!@.-_*] when generating the password.
        """
elif False:
    ItemSectionFieldPasswordRecipeArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ItemSectionFieldPasswordRecipeArgs:
    def __init__(__self__, *,
                 digits: Optional[pulumi.Input[bool]] = None,
                 length: Optional[pulumi.Input[int]] = None,
                 letters: Optional[pulumi.Input[bool]] = None,
                 symbols: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[bool] digits: Use digits [0-9] when generating the password.
        :param pulumi.Input[int] length: The length of the password to be generated.
        :param pulumi.Input[bool] letters: Use letters [a-zA-Z] when generating the password.
        :param pulumi.Input[bool] symbols: Use symbols [!@.-_*] when generating the password.
        """
        if digits is not None:
            pulumi.set(__self__, "digits", digits)
        if length is not None:
            pulumi.set(__self__, "length", length)
        if letters is not None:
            pulumi.set(__self__, "letters", letters)
        if symbols is not None:
            pulumi.set(__self__, "symbols", symbols)

    @property
    @pulumi.getter
    def digits(self) -> Optional[pulumi.Input[bool]]:
        """
        Use digits [0-9] when generating the password.
        """
        return pulumi.get(self, "digits")

    @digits.setter
    def digits(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "digits", value)

    @property
    @pulumi.getter
    def length(self) -> Optional[pulumi.Input[int]]:
        """
        The length of the password to be generated.
        """
        return pulumi.get(self, "length")

    @length.setter
    def length(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "length", value)

    @property
    @pulumi.getter
    def letters(self) -> Optional[pulumi.Input[bool]]:
        """
        Use letters [a-zA-Z] when generating the password.
        """
        return pulumi.get(self, "letters")

    @letters.setter
    def letters(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "letters", value)

    @property
    @pulumi.getter
    def symbols(self) -> Optional[pulumi.Input[bool]]:
        """
        Use symbols [!@.-_*] when generating the password.
        """
        return pulumi.get(self, "symbols")

    @symbols.setter
    def symbols(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "symbols", value)


if not MYPY:
    class GetItemFileArgsDict(TypedDict):
        content: str
        """
        The content of the file.
        """
        content_base64: str
        """
        The content of the file in base64 encoding. (Use this for binary files.)
        """
        id: str
        """
        The UUID of the file.
        """
        name: str
        """
        The name of the file.
        """
elif False:
    GetItemFileArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class GetItemFileArgs:
    def __init__(__self__, *,
                 content: str,
                 content_base64: str,
                 id: str,
                 name: str):
        """
        :param str content: The content of the file.
        :param str content_base64: The content of the file in base64 encoding. (Use this for binary files.)
        :param str id: The UUID of the file.
        :param str name: The name of the file.
        """
        pulumi.set(__self__, "content", content)
        pulumi.set(__self__, "content_base64", content_base64)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def content(self) -> str:
        """
        The content of the file.
        """
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: str):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter(name="contentBase64")
    def content_base64(self) -> str:
        """
        The content of the file in base64 encoding. (Use this for binary files.)
        """
        return pulumi.get(self, "content_base64")

    @content_base64.setter
    def content_base64(self, value: str):
        pulumi.set(self, "content_base64", value)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The UUID of the file.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: str):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the file.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: str):
        pulumi.set(self, "name", value)


if not MYPY:
    class GetItemSectionArgsDict(TypedDict):
        id: str
        """
        A unique identifier for the section.
        """
        label: str
        """
        The label for the section.
        """
        fields: NotRequired[Sequence['GetItemSectionFieldArgsDict']]
        files: NotRequired[Sequence['GetItemSectionFileArgsDict']]
        """
        A list of files attached to the section.
        """
elif False:
    GetItemSectionArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class GetItemSectionArgs:
    def __init__(__self__, *,
                 id: str,
                 label: str,
                 fields: Optional[Sequence['GetItemSectionFieldArgs']] = None,
                 files: Optional[Sequence['GetItemSectionFileArgs']] = None):
        """
        :param str id: A unique identifier for the section.
        :param str label: The label for the section.
        :param Sequence['GetItemSectionFileArgs'] files: A list of files attached to the section.
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "label", label)
        if fields is not None:
            pulumi.set(__self__, "fields", fields)
        if files is not None:
            pulumi.set(__self__, "files", files)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        A unique identifier for the section.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: str):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def label(self) -> str:
        """
        The label for the section.
        """
        return pulumi.get(self, "label")

    @label.setter
    def label(self, value: str):
        pulumi.set(self, "label", value)

    @property
    @pulumi.getter
    def fields(self) -> Optional[Sequence['GetItemSectionFieldArgs']]:
        return pulumi.get(self, "fields")

    @fields.setter
    def fields(self, value: Optional[Sequence['GetItemSectionFieldArgs']]):
        pulumi.set(self, "fields", value)

    @property
    @pulumi.getter
    def files(self) -> Optional[Sequence['GetItemSectionFileArgs']]:
        """
        A list of files attached to the section.
        """
        return pulumi.get(self, "files")

    @files.setter
    def files(self, value: Optional[Sequence['GetItemSectionFileArgs']]):
        pulumi.set(self, "files", value)


if not MYPY:
    class GetItemSectionFieldArgsDict(TypedDict):
        id: str
        """
        A unique identifier for the field.
        """
        label: str
        """
        The label for the field.
        """
        purpose: str
        """
        Purpose indicates this is a special field: a username, password, or notes field. One of ["USERNAME" "PASSWORD" "NOTES"]
        """
        type: str
        """
        The type of value stored in the field. One of ["STRING" "CONCEALED" "EMAIL" "URL" "OTP" "DATE" "MONTH_YEAR" "MENU"]
        """
        value: str
        """
        The value of the field.
        """
elif False:
    GetItemSectionFieldArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class GetItemSectionFieldArgs:
    def __init__(__self__, *,
                 id: str,
                 label: str,
                 purpose: str,
                 type: str,
                 value: str):
        """
        :param str id: A unique identifier for the field.
        :param str label: The label for the field.
        :param str purpose: Purpose indicates this is a special field: a username, password, or notes field. One of ["USERNAME" "PASSWORD" "NOTES"]
        :param str type: The type of value stored in the field. One of ["STRING" "CONCEALED" "EMAIL" "URL" "OTP" "DATE" "MONTH_YEAR" "MENU"]
        :param str value: The value of the field.
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "label", label)
        pulumi.set(__self__, "purpose", purpose)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        A unique identifier for the field.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: str):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def label(self) -> str:
        """
        The label for the field.
        """
        return pulumi.get(self, "label")

    @label.setter
    def label(self, value: str):
        pulumi.set(self, "label", value)

    @property
    @pulumi.getter
    def purpose(self) -> str:
        """
        Purpose indicates this is a special field: a username, password, or notes field. One of ["USERNAME" "PASSWORD" "NOTES"]
        """
        return pulumi.get(self, "purpose")

    @purpose.setter
    def purpose(self, value: str):
        pulumi.set(self, "purpose", value)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of value stored in the field. One of ["STRING" "CONCEALED" "EMAIL" "URL" "OTP" "DATE" "MONTH_YEAR" "MENU"]
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: str):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The value of the field.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: str):
        pulumi.set(self, "value", value)


if not MYPY:
    class GetItemSectionFileArgsDict(TypedDict):
        content: str
        """
        The content of the file.
        """
        content_base64: str
        """
        The content of the file in base64 encoding. (Use this for binary files.)
        """
        id: str
        """
        The UUID of the file.
        """
        name: str
        """
        The name of the file.
        """
elif False:
    GetItemSectionFileArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class GetItemSectionFileArgs:
    def __init__(__self__, *,
                 content: str,
                 content_base64: str,
                 id: str,
                 name: str):
        """
        :param str content: The content of the file.
        :param str content_base64: The content of the file in base64 encoding. (Use this for binary files.)
        :param str id: The UUID of the file.
        :param str name: The name of the file.
        """
        pulumi.set(__self__, "content", content)
        pulumi.set(__self__, "content_base64", content_base64)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def content(self) -> str:
        """
        The content of the file.
        """
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: str):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter(name="contentBase64")
    def content_base64(self) -> str:
        """
        The content of the file in base64 encoding. (Use this for binary files.)
        """
        return pulumi.get(self, "content_base64")

    @content_base64.setter
    def content_base64(self, value: str):
        pulumi.set(self, "content_base64", value)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The UUID of the file.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: str):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the file.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: str):
        pulumi.set(self, "name", value)


