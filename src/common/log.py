from loguru import logger

from common.config import get_config


def init_logger():
    log_filepath = get_config("log_filepath")
    logger.add(
        f"{log_filepath}/info.log", rotation="00:00", retention="7 days", level="INFO"
    )
    logger.add(
        f"{log_filepath}/error.log",
        rotation="00:00",
        retention="7 days",
        level="ERROR",
        backtrace=True,
    )
