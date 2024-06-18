import datetime
from loguru import logger
import requests

from common.config import get_config
from common.jwt import encode_jwt
from user.model import User


def generate_auth_token(user: User) -> str:
    """
    生成登录token
    """
    payload = {
        "user_id": user.id,
        "email": user.email,
        "exp": datetime.datetime.now() + datetime.timedelta(days=5),
    }
    return encode_jwt(payload)


def get_github_access_token(code: str) -> str:
    """
    获取github的access token
    """
    client_id = get_config("github.client_id")
    client_secret = get_config("github.client_secret")
    url = get_config("github.access_token_url")
    data = {
        "client_id": client_id,
        "client_secret": client_secret,
        "code": code,
    }
    response = requests.post(url, data=data)
    logger.info(f"get_github_access_token response:{response.text}")
    # response text 格式： access_token=XX&expires_in=XX&refresh_token=XX&refresh_token_expires_in=XX&scope=&token_type=bearer
    token_parts = response.text.split("&")
    access_token = None

    for part in token_parts:
        if "access_token" in part:
            access_token = part.split("=")[1]
            break
    return access_token


def get_user_info_from_github(access_token: str) -> dict:
    """
    从github获取用户信息
    """
    user_info_url = get_config("github.user_info_url")
    user_email_url = get_config("github.user_email_url")

    headers = {"Authorization": f"Bearer {access_token}"}
    user_info_response = requests.get(user_info_url, headers=headers)
    user_data = user_info_response.json()
    name = user_data.get("name")
    avatar = user_data.get("avatar_url")
    email = user_data.get("email")
    if not email:
        user_email_response = requests.get(user_email_url, headers=headers)
        user_email_data = user_email_response.json()
        email = user_email_data[0].get("email")
    return {"email": email, "name": name, "avatar": avatar}
