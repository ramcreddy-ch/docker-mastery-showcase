# LEVEL 7: DOCKER FOR AI / ML
# This Dockerfile shows how to leverage GPUs and handle large model weights.

# 1. Use NVIDIA Base Image (CUDA support)
# This allows the container to talk to the host's GPU
FROM nvidia/cuda:12.0.1-base-ubuntu22.04

# 2. Prevent interactive prompts during build
ENV DEBIAN_FRONTEND=noninteractive

# 3. Install Python and ML dependencies
RUN apt-get update && apt-get install -y \
    python3-pip \
    python3-dev \
    git \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# 4. Large Model Optimization Strategy:
# Pro-tip: Don't COPY 10GB models into the image. Use Volumes or 'docker pull' scripts.
# But for small models, use a specific layer:
COPY requirements.txt .
RUN pip3 install --no-cache-dir -r requirements.txt

# 5. Copy inference code
COPY inference.py .

# 6. Optimized Layering:
# Keep code that changes frequently at the bottom to avoid rebuilding ML layers.
COPY . .

# 7. GPU Support check (Will show info if NVIDIA toolkit is set up on host)
CMD ["nvidia-smi"]
