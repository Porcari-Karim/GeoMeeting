import secrets
import sys

key_lenght = 35

if len(sys.argv) > 1:
    try:
        key_lenght = int(sys.argv[1])
    except ValueError:
        print("ERROR : The specified lenght (Args[1]) must be a number \n The value is converted to Int if a décimal is specified")
        sys.exit(1)


key = secrets.token_hex(key_lenght)
print(key)