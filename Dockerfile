# ==============================================================================
# Stage 1: Build React frontend
# ==============================================================================
FROM node:22-alpine AS frontend-builder

WORKDIR /app/web

# Install dependencies first (better caching)
COPY web/package.json web/yarn.lock web/.yarnrc.yml ./
COPY web/.yarn ./.yarn

RUN corepack enable && yarn install --immutable

# Copy source and build
COPY web/ ./
ARG VITE_API_URL=""
ARG VITE_AUTH0_DOMAIN
ARG VITE_AUTH0_CLIENT_ID
ARG VITE_AUTH0_AUDIENCE

ENV VITE_API_URL=$VITE_API_URL
ENV VITE_AUTH0_DOMAIN=$VITE_AUTH0_DOMAIN
ENV VITE_AUTH0_CLIENT_ID=$VITE_AUTH0_CLIENT_ID
ENV VITE_AUTH0_AUDIENCE=$VITE_AUTH0_AUDIENCE

RUN yarn build

# ==============================================================================
# Stage 2: Build Go backend
# ==============================================================================
FROM golang:1.24-alpine AS backend-builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Download dependencies first (better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy source and build
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -o /app/apiserver ./cmd/apiserver

# ==============================================================================
# Stage 3: Production image
# ==============================================================================
FROM alpine:3.20

WORKDIR /app

# Install runtime dependencies
RUN apk add --no-cache ca-certificates tzdata sqlite

# Create non-root user
RUN addgroup -S octoback && adduser -S octoback -G octoback

# Copy binaries and static files
COPY --from=backend-builder /app/apiserver /app/apiserver
COPY --from=frontend-builder /app/web/dist /app/static
COPY migrations /app/migrations

# Set ownership
RUN chown -R octoback:octoback /app

USER octoback

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/ || exit 1

# Run the server
CMD ["/app/apiserver"]
