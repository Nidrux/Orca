FROM node:22

# Set `DEVCONTAINER` environment variable early to leverage cache
ENV DEVCONTAINER=true

# Install pnpm first (less likely to change frequently)
RUN npm i -g pnpm@latest

# Install basic development tools (keep apt cache clean)
RUN apt update && apt install -y less man-db sudo && rm -rf /var/lib/apt/lists/*

# Ensure default `node` user has access to `sudo`
ARG USERNAME=node
RUN echo "$USERNAME ALL=(root) NOPASSWD:ALL" > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME
