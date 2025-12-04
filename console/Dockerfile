# ==============================================
# Multi-stage build for Agent Manager Console
# ==============================================

# Stage 1: Build stage
FROM node:20.19-alpine AS builder

# Install Git (required by Rush)
RUN apk add --no-cache git

# Install pnpm globally
RUN npm install -g pnpm@9.12.3

# Install Rush globally
RUN npm install -g @microsoft/rush@5.157.0

WORKDIR /app

# Copy Rush configuration and monorepo structure
COPY rush.json ./
COPY common/ common/
COPY apps/ apps/
COPY workspaces/ workspaces/

# Configure pnpm for Docker builds to avoid cross-device link errors
# Use copy instead of hard links for pnpm
ENV RUSH_PNPM_STORE_PATH=/app/.pnpm-store
RUN pnpm config set store-dir /app/.pnpm-store && \
    pnpm config set package-import-method copy

# Install all dependencies using Rush
# Clean temp directory completely to avoid store path mismatch and recycler EXDEV errors
RUN rm -rf /app/common/temp && \
    rush update --full

# Build all workspace packages and the webapp
RUN rush build --verbose

# Build the webapp for production
WORKDIR /app/apps/webapp
RUN rushx build

# ==============================================
# Stage 2: Runtime stage with Nginx
# ==============================================
FROM nginx:alpine

# Non-root user configuration
ARG USERNAME=appuser
ARG USER_UID=10000
ARG USER_GID=$USER_UID

WORKDIR /app

# Create non-root user and set up writable directories
RUN mkdir -p /tmp && \
    addgroup -g $USER_GID $USERNAME && \
    adduser -D -u $USER_UID -G $USERNAME $USERNAME && \
    chown -R $USER_UID:$USER_GID /app && chmod -R 755 /app && \
    chown -R $USERNAME:$USERNAME /tmp && \
    chown -R $USERNAME:$USERNAME /var/cache/nginx && \
    chown -R $USERNAME:$USERNAME /var/log/nginx && \
    chown -R $USERNAME:$USERNAME /etc/nginx/conf.d && \
    touch /tmp/nginx.pid && \
    chown -R $USERNAME:$USERNAME /tmp/nginx.pid && \
    chmod -R 666 /tmp/nginx.pid

# Copy nginx configuration
COPY docker/nginx.conf /etc/nginx/nginx.conf

# Copy entrypoint script
COPY docker/entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh

# Copy built static files from builder
# Note: config.js is mounted via ConfigMap in Kubernetes
COPY --chown=appuser:appuser --from=builder /app/apps/webapp/dist /app/

# Switch to non-root user
USER appuser

# Expose non-privileged port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --quiet --tries=1 --spider http://localhost:8080/health || exit 1

# Stop signal
STOPSIGNAL SIGQUIT

# Start nginx via entrypoint
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
