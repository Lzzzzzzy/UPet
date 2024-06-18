import os
import yaml
import requests
from loguru import logger


config = None


def get_apollo_config():
    apollo_domain = os.environ.get("APOLLO_CONFIGSERVICE")
    if not apollo_domain:
        return {}
    url = f"{apollo_domain}/configfiles/json/test-case-manager/default/application"
    response = requests.get(url)
    return response.json()


def init_config():
    apollo_config = get_apollo_config()
    logger.info(f"apollo_config:{apollo_config}")
    if not apollo_config:
        current_dir = os.path.dirname(os.path.abspath(__file__))
        project_root = os.path.abspath(os.path.join(current_dir, os.pardir))
        filename = f"{project_root}/config/config_dev.yml"
        with open(filename, "r") as f:
            config_file = yaml.safe_load(f)
            conf = config_file
    else:
        conf = apollo_config
    global config
    config = conf


def get_config(key: str) -> str:
    """
    通过Key获取配置
    """
    return config.get(key) or ""


def get_action_script_folder_path():
    current_directory = os.getcwd()
    action_scripts_path = get_config("action_scripts_path")

    action_scripts_folder_path = os.path.join(current_directory, action_scripts_path)
    return action_scripts_folder_path


def get_case_script_folder_path():
    current_directory = os.getcwd()
    action_scripts_path = get_config("case_scripts_path")

    action_scripts_folder_path = os.path.join(current_directory, action_scripts_path)
    return action_scripts_folder_path


def get_comfirmed_case_script_folder_path():
    current_directory = os.getcwd()
    action_scripts_path = get_config("confirmed_case_scripts_path")

    action_scripts_folder_path = os.path.join(current_directory, action_scripts_path)
    return action_scripts_folder_path
