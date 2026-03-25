import json

# Leer archivo JSON generado por Go
with open("users.json") as f:
    users = json.load(f)

print("Análisis en Python:")

domains = {}
for u in users:
    domain = u["email"].split("@")[1]
    domains[domain] = domains.get(domain, 0) + 1

print(f"Total de usuarios: {len(users)}")
print("Usuarios por dominio:")
for d, c in domains.items():
    print(f"- {d}: {c}")