from flask import request
from flask_restful import marshal

from common.decorators import log_request_error
from common.response import marshal_response

from user import bp
from user.form import QueryUserForm
from user.server import get_users_by_paginate
from user.fields import user_info_fields


@bp.route("/users", methods=["GET"])
@log_request_error
def query_user_api():
    """
    分页查询用户
    """
    query_user_form = QueryUserForm.from_json(request.args)

    user_pagination = get_users_by_paginate(query_user_form)
    user_infos = marshal(user_pagination.items, user_info_fields)
    resp = {
        "pages": user_pagination.pages,
        "total": user_pagination.total,
        "user_infos": user_infos,
    }
    return marshal_response(resp)
