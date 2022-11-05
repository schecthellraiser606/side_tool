import base64

user = b'admin'

def tamper(payload, **kwargs):
    auth = user + payload.encode()
    return base64.b64encode(auth).decode()