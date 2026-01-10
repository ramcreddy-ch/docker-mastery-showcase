# Level 3: Security & Hardening
# This level focuses on reducing the attack surface by running as a non-root user 
# and adding production-grade reliability features.

# --- STAGE 1: Build ---
FROM node:18-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm ci --only=production # Use 'npm ci' for reliable, fast production installs
COPY . .

# --- STAGE 2: Run ---
FROM node:18-alpine

# 1. Use Production Environment
ENV NODE_ENV=production

WORKDIR /app

# 2. Security: Run as a non-privileged user (node user is built-in to node images)
# Even if someone escapes the app, they aren't 'root' inside the container.
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/app.js ./app.js
COPY --from=builder /app/package.json ./package.json

# Ensure the files are owned by the non-root user
RUN chown -R node:node /app

# 3. Healthcheck: Let Docker heartbeat the application
HEALTHCHECK --interval=30s --timeout=3s --retries=3 \
    CMD wget --quiet --tries=1 --spider http://localhost:3000/ || exit 1

# 4. Switch to the non-root user
USER node

EXPOSE 3000
CMD ["node", "app.js"]
