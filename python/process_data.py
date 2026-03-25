import json
import csv
from datetime import datetime

with open("../users.json") as f:
    users = json.load(f)


file_name = f"users_python_{int(datetime.now().timestamp())}.csv"
with open(file_name, "w", newline="") as f:
    writer = csv.DictWriter(f, fieldnames=["name", "email"])
    writer.writeheader()
    writer.writerows(users)

print(f"{len(users)} usuarios exportados por Python a {file_name}")