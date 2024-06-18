import re


def find_enum_member(enum_class, value):
    """
    通过枚举值找到对应的枚举成员
    """
    for member in enum_class.__members__.values():
        if member.value == value:
            return member
    return None


def camel_to_snake(input_string):
    """
    驼峰命名转蛇形命名
    """
    formatted_string = re.sub(r"([A-Z])", r"_\1", input_string).lower()
    if formatted_string.startswith("_"):
        formatted_string = formatted_string[1:]
    return formatted_string


def snake_to_pascal(snake_str):
    """
    蛇形命名转大驼峰
    """
    components = snake_str.split("_")
    pascal_str = "".join(x.title() for x in components)
    return pascal_str
