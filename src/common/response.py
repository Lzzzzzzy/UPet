from common.message_code import MessageCodes


def marshal_response(data: dict = None, msg_code: MessageCodes = MessageCodes.SUCCESS):
    """
    拼接response返回值
    """
    code_info = msg_code.value
    result = {"msg_code": code_info[0], "msg": code_info[1]}
    if data is not None:
        result["data"] = data
    return result
