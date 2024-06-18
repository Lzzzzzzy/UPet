import sqlalchemy


class Enum(sqlalchemy.types.TypeDecorator):
    impl = sqlalchemy.types.SmallInteger
    cache_ok = True

    def __init__(self, enum_class, **kwargs):
        super().__init__(**kwargs)
        self._enum_class = enum_class

    def process_bind_param(self, value, dialect):
        return value.value

    def process_result_value(self, value, dialect):
        if value is not None:
            return self._enum_class(value)
        if value is None:
            if "none" in self._enum_class.__members__:
                return self._enum_class.none
            if None in self._enum_values:
                return self._enum_class(None)

        return None

    @property
    def _enum_values(self):
        return [v.value for v in self._enum_class]

    @property
    def python_type(self):
        return self._enum_class
