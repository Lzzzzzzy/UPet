from flask import Blueprint

bp = Blueprint("auth", __name__, url_prefix="/api")

from . import api
