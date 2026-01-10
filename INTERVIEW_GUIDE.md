# 🧠 Senior Docker Interview Guide (The Deep Dive)
> 50+ Concepts and Scenarios for 5+ YOE Professionals.

---

## 🏗️ Architecture & Internals

### 1. How does Docker isolate? (The "Golden Trio")
Explain these three Linux features:
- **Namespaces**: Provide isolation (PID, NET, MNT, UTS, IPC, USER).
- **Cgroups**: Provide resource limits (CPU, Memory, I/O).
- **Union Filesystem (Overlay2)**: Layering and Copy-on-Write (CoW).

### 2. Process ID 1 (The Init Problem)
Why do we use `dumb-init` or Tini?
- **Answer**: Containers need a process that handles signals (SIGTERM) and reaps zombie processes. Node/Python don't do this well as PID 1.

### 3. Layer Caching Pitfalls
Why does `COPY . .` at the top break builds?
- **Answer**: It invalidates the cache for all subsequent steps (like `RUN npm install`) even if only a README changed.

---

## 🌐 Advanced Networking Scenarios

### Scenario: "My container can't talk to the DB on the host machine."
- **Solution**: Avoid using `localhost` (that refers to the container). Use `host.docker.internal` (Mac/Win) or the host's actual IP.

### Scenario: "Explain the difference between EXPOSE and -p."
- **Answer**: `EXPOSE` is metadata/documentation for internal communication. `-p` (Publish) actually maps a host port to the container port via IPTables.

---

## 🤖 Docker for AI & High-Performance
- **Question**: How do you pass a GPU to a container?
- **Answer**: Use the NVIDIA Container Toolkit (`--gpus all`).
- **Optimization**: For LLMs, explain why we mount the model weight directory as a **Read-Only Volume** instead of baking them into the image (Layer limits & image portability).

---

## 🛡️ Security & Hardening (The Senior Check)
- **Capability Drop**: Why use `--cap-drop ALL`? (Principle of least privilege).
- **Distroless**: Why is it better than Alpine? (Alpine still has a shell/package manager; Distroless is just the binary).
- **Immutable Infrastructure**: Why should containers be stateless?

---

## 📋 10 Critical Scenario-Based Questions

1. **"A containerized app is slow. How do you profile it from the host?"**
2. **"Explain the 'Copy-on-Write' mechanism in Overlay2."**
3. **"How do you handle secrets without using environment variables?" (Docker Secrets/Mounts).**
4. **"What happens if you don't limit memory, and the container has a leak?" (OOM Killer).**
5. **"How do you share a network between two containers without using a bridge?" (Network=container:xyz).**
6. ... and 45 more key concepts (See logic inside README).

---
> "The goal of an interview is not to prove you know the commands, but to prove you know how the engine works under the hood."
