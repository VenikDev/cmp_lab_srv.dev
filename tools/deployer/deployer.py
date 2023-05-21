import paramiko
from dotenv import load_dotenv
import os


def main():
    ssh = paramiko.SSHClient()

    # добавляем новый ключ в список известных хостов
    ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())

    # устанавливаем соединение
    ssh.connect(
        os.getenv('SSH_IP'),
        username=os.getenv('SSH_NAME'),
        password=os.getenv('SSH_PASSWORD'))

    # отправляем команды
    ssh.exec_command('cd cmp-srv/')
    ssh.exec_command('git pull')
    stdin, stdout, stderr = ssh.exec_command('make deploy SCALE=3')

    result = stdout.read()
    print(result.decode())
    ssh.close()


if __name__ == "__main__":
    load_dotenv()
    main()
