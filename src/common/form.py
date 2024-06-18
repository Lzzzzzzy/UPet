from flask import request
from flask_wtf import FlaskForm
from flask_restful import fields
from wtforms import IntegerField
from wtforms.utils import unset_value

from common.tools import find_enum_member


class BaseForm(FlaskForm):
    class Meta:
        csrf = False

    @classmethod
    def from_json(cls, json_data):
        return cls(formdata=None, data=json_data)

    @classmethod
    def from_request(cls):
        if request.method == "GET":
            data = request.args
        else:
            data = request.get_json()
        return cls.from_json(data)

    def get_form_error(self):
        errors = self.errors
        errors_data = {}
        for field, errors in self.errors.items():
            errors_data[field] = errors[0].value[1]
        return errors_data


class PagenationForm(BaseForm):
    page = IntegerField(default=1)
    page_size = IntegerField(default=20)


class EnumValueField(fields.Raw):

    def format(self, value):
        try:
            return value.value
        except Exception as e:
            raise e


class EnumNameField(fields.Raw):

    def format(self, value):
        try:
            return value.name.lower()
        except Exception as e:
            raise e


class IntEnumField(IntegerField):

    def __init__(self, enum_cls, label=None, validators=None, **kwargs):
        super().__init__(label, validators, **kwargs)
        self.enum_cls = enum_cls

    def _value(self):
        if self.raw_data:
            return self.raw_data[0]
        if self.data is not None:
            return str(self.data)
        return ""

    def process_data(self, value):
        if value is None or value is unset_value:
            self.data = None
            return

        try:
            self.data = find_enum_member(self.enum_cls, int(value))
        except (ValueError, TypeError) as exc:
            self.data = None
            raise ValueError(self.gettext("Not a valid integer enum value.")) from exc

    def process_formdata(self, valuelist):
        if not valuelist:
            return

        try:
            self.data = find_enum_member(self.enum_cls, int(valuelist[0]))
        except ValueError as exc:
            self.data = None
            raise ValueError(self.gettext("Not a valid integer enum value.")) from exc
