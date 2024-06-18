from user import bp as user_bp
from auth import bp as auth_bp


def init_blueprint(app):
    app.register_blueprint(user_bp)
    app.register_blueprint(auth_bp)
