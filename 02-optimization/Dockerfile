# Level 2: Optimization
# This level focuses on Multi-stage builds and reducing image size.

# --- STAGE 1: Build ---
FROM node:18-alpine AS builder
WORKDIR /app
# Copy only package files first to leverage Docker layer caching
COPY package*.json ./
RUN npm install
# Copy rest of the app
COPY . .

# --- STAGE 2: Run ---
FROM node:18-alpine
WORKDIR /app
# Only copy the necessary files from the builder stage
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/app.js ./app.js
COPY --from=builder /app/package.json ./package.json

EXPOSE 3000
CMD ["node", "app.js"]
