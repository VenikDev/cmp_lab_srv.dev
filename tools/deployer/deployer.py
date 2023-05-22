import sys

import paramiko
from dotenv import load_dotenv


def main():
    ssh = paramiko.SSHClient()

    # Automatically add the server's host key
    ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())

    # Connect to the server
    ssh.connect(sys.argv[1],
                username=sys.argv[2],
                password=sys.argv[2],
                disabled_algorithms={'pubkeys': ['rsa-sha2-256', 'rsa-sha2-512']})

    # Send a command
    stdin, stdout, stderr = ssh.exec_command('cd cmp-srv && git pull && make deploy SCALE=3')

    # Print the output of the command
    print(stdout.read().decode())

    # Close the SSH connection
    ssh.close()


if __name__ == "__main__":
    load_dotenv()
    main()
