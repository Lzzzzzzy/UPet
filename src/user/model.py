from datetime import datetime
from database import db, Base, CreatedModel, SoftDeleteModel, UpdateModel


class User(Base, CreatedModel, SoftDeleteModel, UpdateModel):
    id = db.Column(db.Integer, primary_key=True, autoincrement=True)
    name = db.Column(db.String(255), nullable=False, comment="用户名")
    email = db.Column(db.String(255), comment="邮箱")
    phone = db.Column(db.String(255), comment="电话")
    last_login_time = db.Column(db.DateTime, comment="最后一次登录时间")

    def __init__(self, name, email):
        CreatedModel.__init__(self, 0)
        UpdateModel.__init__(self, 0)

        self.name = name
        self.email = email
        self.last_login_time = datetime.now()
