from typing import List
from user.form import QueryUserForm
from user.model import User


def get_users_by_paginate(form: QueryUserForm) -> List[User]:
    """
    分页查询用户信息
    """
    base_query = User.query.filter(User.is_deleted.is_(False)).order_by(User.id.desc())
    if form.email.data:
        base_query = base_query.filter(User.email == form.email.data)
    if form.name.data:
        base_query = base_query.filter(User.name == form.name.data)
    return base_query.paginate(
        page=form.page.data, per_page=form.page_size.data, error_out=False
    )
