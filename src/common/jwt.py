import jwt

from common.config import get_config


def encode_jwt(payload: dict) -> str:
    """
    生成jwt
    """
    jwt_secret = get_config("jwt_secret")
    jwt_algorithm = get_config("jwt_algorithm")
    token = jwt.encode(payload, jwt_secret, algorithm=jwt_algorithm)
    return token


def decode_jwt(token: str) -> dict:
    """
    解析jwt
    """
    try:
        jwt_secret = get_config("jwt_secret")
        jwt_algorithm = get_config("jwt_algorithm")
        decoded_data = jwt.decode(token, jwt_secret, algorithms=[jwt_algorithm])
        return decoded_data
    except Exception as e:
        return {}
