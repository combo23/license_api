import requests
import os
import hwid

LICENSE_KEY = "TEST"

def get_hwid():
    return hwid.get_hwid()

if __name__ == '__main__':
    print("Checking license...")

    req = requests.post(
        "http://127.0.0.1:8080/verify", 
        headers={"Content-Type": "application/json"}, 
        json={"hwid": get_hwid(), "license_key": LICENSE_KEY}
    )

    if req.status_code == 200:
        print(req.json()["message"])
        # verified
    else:
        print(req.json()["error"])
        os._exit(1)