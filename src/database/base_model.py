from flask import request
from database import db
from sqlalchemy.ext.declarative import declarative_base
from datetime import datetime

Base = declarative_base()


class CreatedModel(db.Model):
    __abstract__ = True
    in_dtm = db.Column(db.DateTime, default=datetime.now, comment="创建时间")
    in_user_id = db.Column(db.Integer, comment="创建人")

    def __init__(self, user_id, in_dtm=None) -> None:
        self.in_user_id = user_id
        self.in_dtm = in_dtm or datetime.now()

    @property
    def in_user(self):
        from user.model import User

        return User.query.get(self.in_user_id)


class UpdateModel(db.Model):
    __abstract__ = True
    edit_dtm = db.Column(
        db.DateTime, default=datetime.now, onupdate=datetime.now, comment="更新时间"
    )
    edit_user_id = db.Column(db.Integer, comment="更新人")

    def __init__(self, user_id, edit_dtm=None) -> None:
        self.edit_user_id = user_id
        self.edit_dtm = edit_dtm or datetime.now()

    def update_model(self, user_id):
        self.edit_user_id = user_id
        self.edit_dtm = datetime.now()

    @property
    def edit_user(self):
        from user.model import User

        return User.query.get(self.edit_user_id)


class DeleteModel(db.Model):
    __abstract__ = True
    remove_user_id = db.Column(db.Integer, comment="刪除人", default=0)
    remove_dtm = db.Column(
        db.DateTime, comment="删除时间", default="1900-01-01 00:00:00"
    )

    def __init__(self, user_id=0, remove_dtm=None) -> None:
        super().__init__()
        self.remove_user_id = user_id
        self.remove_dtm = remove_dtm or "1900-01-01 00:00:00"

    @property
    def remove_user(self):
        from user.model import User

        return User.query.get(self.remove_user_id)


class SoftDeleteModel(DeleteModel):
    __abstract__ = True
    is_deleted = db.Column(
        db.Boolean, default=False, comment="是否删除, 0:未删除, 1:已删除"
    )

    def soft_delete(self, user_id):
        self.is_deleted = 1
        self.remove_dtm = datetime.now()
        self.remove_user_id = user_id

    def rollback(self):
        self.is_deleted = 0
