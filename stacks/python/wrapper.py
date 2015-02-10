import socket
import sys
import os
import json

YASEN_PATH = os.environ['PROJECTS'] +'/yasen'
UNITS_CONFIG_PATH = YASEN_PATH +'/configs/units.json'

if len(sys.argv) != 2:
    print('Usage: wrapper.py unit')
    exit(1)

unit = sys.argv[1]
config = json.load(open(UNITS_CONFIG_PATH)).get(unit, None)
if not config:
    print('Unknown unit: %s' % unit)
    exit(1)

unit_path = config['impl'].split('.')
unit_func = unit_path.pop()
unit_module = unit_path.pop()
unit_path = '.'.join(unit_path)

sys.path.append(YASEN_PATH)
_units = __import__(unit_path, fromlist=[unit_module])
_unit = getattr(_units, unit_module)
actor = getattr(_unit, unit_func)

def process(input):
    data = json.loads(input.decode('utf-8'))
    new_data = actor(data)
    return json.dumps(new_data).encode('utf-8')

SOCKET_PATH = '/tmp/unit_%s.sock' % unit

# Make sure the socket does not already exist
try: os.unlink(SOCKET_PATH)
except OSError:
    if os.path.exists(SOCKET_PATH): raise

# Create a UDS socket
sock = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)

# Bind the socket to the port
print('Listening on socket %s' % SOCKET_PATH)
sock.bind(SOCKET_PATH)

# Listen for incoming connections
sock.listen(1)

BUF_SIZE = 1024
while True:
    print('Accepting connections')
    conn, _ = sock.accept()
    try:
        print('New connection')
        data = b''
        while True:
            buf = conn.recv(BUF_SIZE)
            data += buf
            if len(buf) < BUF_SIZE:
                break

        print('Received in total %d bytes' % len(data))

        new_data = process(data)
        conn.sendall(new_data)
    finally:
        conn.close()
