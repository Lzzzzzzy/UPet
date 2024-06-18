from enum import Enum, unique


@unique
class MessageCodes(Enum):
    SUCCESS = 0, "Success"
    FORM_VALID_FAILED = 1, "FORM VALID FAILED"

    # auth
    NEED_LOGIN = 1000, "Need login"
    GITHUB_LOGIN_ERROR = 1001, "Login from Github error"

    # project
    NO_PROJECT = 1701, "No project"

    # machine
    IP_EXIST = 1101, "Machine is exist"
    NO_MACHINE = 1102, "No machine"
    IP_REQUIRED = 1103, "IP is required"
    PORT_REQUIRED = 1104, "Port is required"
    NOT_ALLOW_EDIT_MACHINE = 1105, "This machine is not allowed to modify"
    NOT_ALLOW_DELETE_MACHINE = 1105, "This machine is not allowed to delete"

    # tag
    CONTENT_REQUIRED = 1201, "Content is required"
    NO_TAG = 1202, "No tag"

    # action
    PROJECT_ID_REQUIRED = 1301, "Project is required"
    PROJECT_NOT_EXIST = 1302, "Project not exist"
    ACTION_NAME_REQUIRED = 1303, "Action name is required"
    ACTION_OPERATE_WRONG = 1304, "Operate not support"
    NO_ACTION = 1305, "No action"
    ACTION_PATH_NOT_EXIST = 1306, "Action script path not exist"

    # case
    CASE_NAME_REQUIRED = 1401, "Case name is required"
    CASE_OPERATE_WRONG = 1402, "Operate not support"
    NO_CASE = 1403, "No case"
    CASE_NAME_MUST_BE_ENGLISH = 1404, "Case name must be English"
    CASE_PARSE_ERROR = 1405, "Error when running case, please check parameters"
    CASE_ENVIRONMENT_ERROR = 1406, "Case in checked environments is not exist"
    CASE_NUMBER_REQUIRED = 1407, "Case number is required"
    CASE_NUMBER_REPEAT = 1408, "Case number is existed"
