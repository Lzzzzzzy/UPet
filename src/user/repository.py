from typing import List, Optional
from sqlalchemy import or_
from database import db
from user.model import User


def create_user(name: str, email: str) -> User:
    """
    创建用户
    """
    user = User(name, email)
    db.session.add(user)
    db.session.flush()
    user.in_user_id = user.id
    user.edit_user_id = user.id
    return user


def get_user_by_email(email: str) -> User:
    """
    根据user email查询用户信息
    """
    return (
        db.session.query(User)
        .filter(User.email == email, User.is_deleted.is_(False))
        .first()
    )


def get_user_by_name_or_email(info: str) -> List[User]:
    """
    根据name或email查询user
    """
    conditions = [User.email.like(f"%{info}%"), User.name.like(f"%{info}%")]
    return db.session.query(User).filter(or_(*conditions)).all()


def get_user_by_id(id: int) -> Optional[User]:
    """
    根据user_id查询user
    """
    return User.query.get(id)
