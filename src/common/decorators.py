from flask import request
from loguru import logger
from functools import wraps


def log_request_error(func):

    @wraps(func)
    def wrapper(*args, **kwargs):
        try:
            logger.info(f"request url: {request.path} / {request.method}")
            return func(*args, **kwargs)
        except Exception as e:
            logger.error(f"request url: {request.path} error")
            logger.exception(e)

    return wrapper
