from flask import Flask, request
from flask_cors import CORS

from common.config import init_config, get_config
from common.jwt import decode_jwt
from database import init_db
from common.blueprint import init_blueprint
from common.response import marshal_response
from common.log import init_logger
from common.message_code import MessageCodes


def init_app():
    app = Flask(__name__)
    init_config()

    app_secret = get_config("app_secret")
    app.config["SECRET_KEY"] = app_secret

    CORS(app, resources={r"/api/*": {"origins": "*"}})

    init_db(app)
    init_logger()
    init_blueprint(app)

    @app.before_request
    def before_request():
        no_need_token_uri = [
            "/api/github/url",
            "/api/github/oauth",
            "/info",
            "/api/projects/init",
        ]
        if request.method == "OPTIONS" or request.path in no_need_token_uri:
            return

        authorization = request.headers.get("Authorization") or ""
        split_authorization = authorization.split(" ")
        if len(split_authorization) != 2 or split_authorization[0] != "Bearer":
            return marshal_response(msg_code=MessageCodes.NEED_LOGIN)
        payload = decode_jwt(split_authorization[1])
        if not payload:
            return marshal_response(msg_code=MessageCodes.NEED_LOGIN)
        request.current_user = payload
        request.current_user_id = payload.get("user_id")

    @app.route("/info")
    def health_check():
        return marshal_response()

    return app


app = init_app()

if __name__ == "__main__":
    debug = True
    port = get_config("port")

    app.run(debug=debug, port=port, host="0.0.0.0")
