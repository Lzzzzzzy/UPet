from datetime import datetime
from flask import request
from flask_restful import marshal

from auth import bp
from auth.service import (
    generate_auth_token,
    get_github_access_token,
    get_user_info_from_github,
)
from common.config import get_config
from common.decorators import log_request_error
from common.message_code import MessageCodes
from common.response import marshal_response
from database import db
from user.fields import user_info_fields
from user.repository import get_user_by_email, create_user


@bp.route("/github/oauth", methods=["POST"])
@log_request_error
def github_callback():
    """
    github OAuth登录
    """
    code = request.json.get("code")

    access_token = get_github_access_token(code)
    if not access_token:
        return marshal_response(msg_code=MessageCodes.GITHUB_LOGIN_ERROR)

    github_user_info = get_user_info_from_github(access_token)

    user = get_user_by_email(github_user_info["email"])
    if not user:
        email = github_user_info["email"]
        username = github_user_info.get("name")
        if not username:
            username = email.split("@")[0]
        user = create_user(username, email)
        user.is_admin = True
    user.last_login_time = datetime.now()
    db.session.commit()

    user_info = dict(marshal(user, user_info_fields))
    auth_jwt = generate_auth_token(user)
    user_info.update(github_user_info)
    return marshal_response({"user_info": user_info, "auth_token": auth_jwt})


@bp.route("/github/url", methods=["GET"])
@log_request_error
def get_third_party_login_url_for_github():
    """
    获取github OAuth 登录地址
    """
    auth_url = get_config("github.auth_url")
    client_id = get_config("github.client_id")
    redirect_url = get_config("github.redict_url")
    return marshal_response(
        {"login_url": f"{auth_url}?client_id={client_id}&redirect_uri={redirect_url}"}
    )
