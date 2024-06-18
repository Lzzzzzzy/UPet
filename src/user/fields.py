from flask_restful import fields

user_info_fields = {
    "id": fields.Integer,
    "email": fields.String,
    "name": fields.String,
    "is_admin": fields.Boolean,
}
