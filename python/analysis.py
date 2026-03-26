import json
import csv
import requests

resp = requests.get("https://jsonplaceholder.typicode.com/users")
users = resp.json()

biz_users = [u for u in users if u["email"].endswith(".biz")]

with open("analysis_users.json", "w") as f:
    json.dump(biz_users, f, indent=2)

with open("analysis_users.csv", "w", newline="") as f:
    writer = csv.writer(f)
    writer.writerow(["name", "email"])
    for u in biz_users:
        writer.writerow([u["name"], u["email"]])

print(f"{len(biz_users)} usuarios procesados y exportados")