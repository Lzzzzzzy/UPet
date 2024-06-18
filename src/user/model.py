from datetime import datetime
from database import db, Base, CreatedModel, SoftDeleteModel, UpdateModel


class User(Base, CreatedModel, SoftDeleteModel, UpdateModel):
    id = db.Column(db.Integer, primary_key=True, autoincrement=True)
    name = db.Column(db.String(255), nullable=False, comment="账号")
    email = db.Column(db.String(255), comment="邮箱")
    last_login_time = db.Column(db.DateTime, comment="最后一次登录时间")
    is_admin = db.Column(db.Boolean, default=False, comment="是否是管理员, 0:否, 1:是")

    def __init__(self, name, email):
        CreatedModel.__init__(self, 0)
        UpdateModel.__init__(self, 0)

        self.name = name
        self.email = email
        self.last_login_time = datetime.now()
