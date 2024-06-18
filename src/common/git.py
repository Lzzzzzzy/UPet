import subprocess

from loguru import logger


def clone_repository(git_addr: str, local_path: str, branch_name: str):
    subprocess.run(["git", "clone", "-b", branch_name, git_addr, local_path])


def pull_repository(local_path: str, branch_name: str):
    subprocess.run(["git", "stash"], cwd=local_path)
    subprocess.run(["git", "pull", "origin", branch_name], cwd=local_path)
    subprocess.run(["git", "stash", "pop"], cwd=local_path)


def push_changes_to_repository(
    local_path: str, branch_name: str, commit_msg: str = "commit"
):
    # 添加文件到暂存区
    subprocess.run(["git", "add", "."], cwd=local_path)

    # 提交更改
    subprocess.run(["git", "commit", "-m", commit_msg], cwd=local_path)

    logger.info(f"push to branch:{branch_name}")
    subprocess.run(["git", "push", "origin", branch_name], cwd=local_path)
