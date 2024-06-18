import pymysql
from flask_sqlalchemy import SQLAlchemy
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker, scoped_session
from common.config import get_config

db = SQLAlchemy()


def init_db(app):
    username = get_config("database.username")
    password = get_config("database.password")
    host = get_config("database.host")
    port = get_config("database.port")
    database = get_config("database.database")
    sqlalchemy_track_modifications = False

    database_uri = f"mysql+pymysql://{username}:{password}@{host}:{port}/{database}"
    app.config["SQLALCHEMY_DATABASE_URI"] = database_uri
    app.config["SQLALCHEMY_TRACK_MODIFICATIONS"] = sqlalchemy_track_modifications
    app.config["SQLALCHEMY_POOL_SIZE"] = 10  # 连接池大小
    app.config["SQLALCHEMY_POOL_TIMEOUT"] = 60  # 连接超时时间
    app.config["SQLALCHEMY_POOL_RECYCLE"] = 3600  # 连接回收时间

    global db
    db.init_app(app)
    engine = create_engine(database_uri, pool_size=20, max_overflow=5)
    session_factory = sessionmaker(bind=engine)
    db.session = scoped_session(session_factory)
