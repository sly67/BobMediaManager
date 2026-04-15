# Stage 1: Build SvelteKit frontend
FROM node:22-alpine AS frontend
WORKDIR /app/frontend
COPY frontend/package*.json ./
RUN npm ci
COPY frontend/ .
RUN npm run build

# Stage 2: Build Go backend (embed frontend output)
FROM golang:1.23-alpine AS backend
WORKDIR /app
COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY backend/ .
COPY --from=frontend /app/frontend/build ./static/
RUN go build -ldflags="-s -w" -o bobmediamanager .

# Stage 3: Minimal runtime image
FROM alpine:3.20
RUN apk add --no-cache ca-certificates tzdata
COPY --from=backend /app/bobmediamanager /app/bobmediamanager
EXPOSE 48040
ENTRYPOINT ["/app/bobmediamanager"]
