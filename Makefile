.PHONY: dev build docker-build docker-up docker-down prune clean lint test

# ── Local development ────────────────────────────────────────────────────────
# Runs the Go backend and SvelteKit dev server concurrently.
# Requires: Go 1.23+ and Node 22+ installed locally.
dev:
	@echo "→ Starting backend..."
	cd backend && /usr/local/go/bin/go run . &
	@echo "→ Starting frontend dev server..."
	cd frontend && npm run dev

# ── Local build ──────────────────────────────────────────────────────────────
build:
	cd frontend && npm ci && npm run build
	cp -r frontend/build backend/static/
	cd backend && /usr/local/go/bin/go build -o bobmediamanager .

# ── Docker ───────────────────────────────────────────────────────────────────
docker-build:
	docker compose build
	@$(MAKE) prune

docker-up:
	docker compose up --build -d
	@$(MAKE) prune

docker-down:
	docker compose down

# ── Storage hygiene ──────────────────────────────────────────────────────────
# Removes dangling images and stale build cache to prevent storage overflow.
prune:
	@echo "→ Pruning dangling Docker images..."
	docker image prune -f
	@echo "→ Pruning Docker build cache..."
	docker builder prune -f

# ── Maintenance ──────────────────────────────────────────────────────────────
clean:
	rm -rf frontend/build frontend/.svelte-kit backend/bobmediamanager backend/static/index.html
	@$(MAKE) prune

lint:
	cd frontend && npm run check
	cd backend && /usr/local/go/bin/go vet ./...

test:
	cd backend && /usr/local/go/bin/go test ./...
