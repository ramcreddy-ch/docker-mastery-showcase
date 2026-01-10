# LEVEL 1: THE FOUNDATION
# This is a standard Dockerfile most beginners start with.
# It's functional, but has flaws: large size, root user, and poor caching.

# 1. Base Image (Thick image, includes everything)
FROM node:18

# 2. Set working directory
WORKDIR /usr/src/app

# 3. Copy EVERYTHING (Inefficient for caching)
COPY . .

# 4. Install dependencies
RUN npm install

# 5. Expose port
EXPOSE 3000

# 6. Start the app
CMD ["node", "app.js"]
