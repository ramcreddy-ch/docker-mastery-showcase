# 🎯 The "Never Sleep" Docker CLI Cheatsheet
> For the Senior Developer who needs to move fast.

---

## 🏗️ Building & Managing Images

| Task | Command | Why? |
| :--- | :--- | :--- |
| **Build** | `docker build -t name:tag .` | Standard build. |
| **Build (No Cache)** | `docker build --no-cache -t name .` | Force clean build when layers get "stuck". |
| **Deep Inspect** | `docker inspect image_name` | See metadata, layers, and env vars. |
| **History** | `docker history image_name` | Debug layer sizes and see Dockerfile steps. |
| **Prune** | `docker image prune -a` | Clean up dangling and unused images. |

---

## 🏃 Running & Debugging Containers

| Task | Command | Why? |
| :--- | :--- | :--- |
| **Interactive** | `docker run -it --name test ubuntu /bin/bash` | Jump inside a shell for debugging. |
| **Exec** | `docker exec -it container_id sh` | Enter a RUNNING container. |
| **Logs** | `docker logs -f --tail 100 name` | Follow logs in real-time. |
| **Stats** | `docker stats` | Live CPU/Memory usage (Crucial for AI/Prod). |
| **Copy File** | `docker cp ./local_file name:/app/file` | Move files in/out of a container without volumes. |
| **Rename** | `docker rename old_name new_name` | Fix a typo in a production container name. |

---

## 🌐 Networking & Volumes

| Task | Command | Why? |
| :--- | :--- | :--- |
| **List Nets** | `docker network ls` | See all virtual networks. |
| **Connect** | `docker network connect net_name cont` | Hot-connect a container to a new network. |
| **Net Inspect** | `docker network inspect bridge` | Find IP addresses of all containers in a net. |
| **Volume LS** | `docker volume ls` | List all persistent storage. |
| **Volumes Prune**| `docker volume prune` | **WARNING:** Cleans up unused volumes (Data loss risk!). |

---

## 🛠️ Advanced Senior Hacks

### 1. The "Clean Slate" Command
```bash
# Wipe everything: Containers, Networks, Images (unused), and Cache
docker system prune -a --volumes
```

### 2. Finding "Fat" Containers
```bash
# Sort containers by size on disk
docker ps -as --format "{{.ID}}: {{.Size}} ({{.Names}})"
```

### 3. Debugging Distroless (Sidecar pattern)
Since Distroless has no shell, use `docker run --network container:<id>` to attach a debugger to the same network namespace.

---
*Note: Use `docker --help` for the basics, use this guide for the win.*
