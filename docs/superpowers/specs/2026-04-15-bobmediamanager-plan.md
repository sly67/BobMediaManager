# BobMediaManager — Implementation Plan
**Date:** 2026-04-15  
**Stack:** SvelteKit + DaisyUI v5 · Go · Docker  
**Port:** 48040

---

## Phase 0 — Project Scaffold

### 0.1 Repo structure
```
BobMediaManager/
├── backend/
│   ├── main.go
│   ├── api/          (handlers, router)
│   ├── service/      (scanner, matcher, rename, metadata, nfo, images, subtitles, verify)
│   ├── sources/      (anidb, tvdb, tmdb, imdb, musicbrainz, lastfm, opensubtitles, custom)
│   ├── models/       (MediaFile, MediaMetadata, CustomSource, Config)
│   └── config/       (config.json loader/saver)
├── frontend/
│   ├── src/
│   │   ├── lib/
│   │   │   ├── components/  (FileRow, SplitView, IdPickerModal, SourcesModal, CtxMenu, …)
│   │   │   ├── stores/      (files, selection, config, theme, progress)
│   │   │   └── api.ts       (typed fetch wrappers + WS client)
│   │   └── routes/
│   │       └── +page.svelte (single-page app)
│   ├── static/
│   └── svelte.config.js
├── docs/
├── Dockerfile
├── docker-compose.yml
└── Makefile
```

### 0.2 Dockerfile
- Stage 1: `node:22-alpine` — build SvelteKit (`npm run build`)
- Stage 2: `golang:1.23-alpine` — build Go binary, embed `frontend/build` with `go:embed`
- Final: scratch/alpine, single binary at `/app/bobmediamanager`

### 0.3 Makefile targets
`dev`, `build`, `docker-build`, `docker-up`, `lint`, `test`

---

## Phase 1 — Go Backend Core

### 1.1 HTTP server + router
- `chi` router
- CORS for dev (SvelteKit dev server on :5173)
- Static file serving from embedded `frontend/build`
- WebSocket endpoint `/ws/progress`

### 1.2 Config service
- Load/save `config.json` from `/config/config.json` (Docker volume)
- Default values on first run
- `GET /api/config` · `PUT /api/config`

### 1.3 Models
Define all structs: `MediaFile`, `MediaMetadata`, `ImageEntry`, `ImageType`, `MediaType`, `MatchStatus`, `CustomSource`, `Config`, `ProgressEvent`.

### 1.4 File scanner
- `POST /api/scan` — accepts `{path, recursive, mediaType}`
- Walk directory, return `[]MediaFile` with name, path, size, isFolder
- Filter by extension (`.mkv .mp4 .avi .flac .mp3 .m4a …`)

---

## Phase 2 — Metadata Sources

### 2.1 Source interface
```go
type MetadataSource interface {
    Search(ctx context.Context, q string, t MediaType) ([]SearchResult, error)
    GetByID(ctx context.Context, id string) (*MediaMetadata, error)
    Name() string
    Slug() string
    SupportedTypes() []MediaType
}
```

### 2.2 TMDB adapter
- Auth: Bearer token (API key v4)
- Search: `/search/tv`, `/search/movie`
- Detail: `/tv/{id}`, `/movie/{id}`, `/tv/{id}/season/{n}/episode/{n}`
- Images: `/tv/{id}/images`, `/movie/{id}/images`

### 2.3 TVDB adapter
- Auth: JWT via `/login`
- Search: `/search`
- Series detail: `/series/{id}/extended`
- Episodes: `/series/{id}/episodes/default`
- Artwork: `/series/{id}/artworks`

### 2.4 AniDB adapter
- UDP API (port 9000) for authenticated queries
- HTTP API for unauthenticated anime/episode info
- Handles rate limiting (AniDB is strict: 1 req/2s)

### 2.5 IMDb adapter
- Cinemagoer (Python subprocess) OR direct scrape with `colly`
- Returns title, year, rating, genres, plot

### 2.6 MusicBrainz adapter
- `https://musicbrainz.org/ws/2/`
- Release group search, recording lookup

### 2.7 Last.fm adapter
- Artist/album info, cover art

### 2.8 OpenSubtitles adapter
- REST API v1: `/subtitles` search, `/download` endpoint

### 2.9 Custom source adapter
- Generic REST GET with configurable path, headers, field map (JSONPath)
- Stored in config, instantiated at startup

### 2.10 Source registry
- Registry mapping slug → `MetadataSource`
- `GET /api/sources` — list with enabled/disabled state
- `POST /api/sources` — register custom
- `DELETE /api/sources/:id` — remove custom
- `POST /api/sources/:slug/test` — test connection

---

## Phase 3 — Matcher + Search

### 3.1 Fuzzy matcher
- Parse filename: strip scene tags (`[SubsPlease]`, `WEB-DL`, etc.), extract title, year, season/episode numbers
- Library: custom regex parser + optional `go-fuzz` scoring
- `POST /api/match` — runs matcher across all files, returns confidence + candidate IDs

### 3.2 Manual search
- `POST /api/search` — `{query, sources[], mediaType}`
- Fan-out search across enabled sources, merge + deduplicate results by ID cross-reference
- Return sorted by confidence

### 3.3 ID override
- `PUT /api/files/:fileID/id` — set explicit `SOURCE:ID` string, bypasses matcher

---

## Phase 4 — Rename Engine

### 4.1 Template engine
- Parse format string: `{n}`, `{y}`, `{s}`, `{e}`, `{s00}`, `{e00}`, `{t}`, `{artist}`, `{album}`, `{track00}`, `{ext}`
- Sanitize output for filesystem (strip illegal chars)
- `POST /api/rename/preview` — dry-run, returns source→dest map
- `POST /api/rename/execute` — performs moves/copies, streams progress via WS

### 4.2 Operations
- `Move`: `os.Rename` (cross-device fallback: copy+delete)
- `Copy`: `io.Copy` with progress
- `RenameInPlace`: rename only, no path change

### 4.3 Folder rename
- When `renameFolders: true`, after file moves detect empty old dirs and rename parent folders using template directory segments

### 4.4 Conflict resolution
- If destination exists: Skip / Overwrite / Rename with suffix (user-configurable)

### 4.5 History / undo
- Write operation log to `/config/history.json`
- `GET /api/history` · `POST /api/history/:id/undo`

---

## Phase 5 — NFO Generator

### 5.1 Templates
- Go `text/template` files embedded in binary
- Types: `tvshow`, `episodedetails`, `movie`, `artist`, `album`
- Kodi / Jellyfin schema
- User-editable templates stored in `/config/nfo-templates/`

### 5.2 Generation
- `POST /api/nfo` — `{fileIDs[], types[]}` — generates `.nfo` files next to each media file
- Streams progress via WS

### 5.3 ID fields in NFO
- Writes all resolved IDs: `<anidbid>`, `<tmdbid>`, `<tvdbid>`, `<imdbid>`, etc.

---

## Phase 6 — Image Downloader

### 6.1 Image types
`poster`, `fanart`, `extrafanart`, `thumb`, `banner`, `logo`, `clearart`, `disc`, `characterart`

### 6.2 Download service
- Worker pool (default 4 concurrent)
- Destination naming:
  - `poster.jpg` — series/movie level
  - `fanart.jpg` — series/movie level
  - `extrafanart/1.jpg` … `extrafanart/N.jpg` — up to `maxExtrafanart`
  - `Season NN/season{N}-poster.jpg`
  - `{episode-base}-thumb.jpg`
- `POST /api/images` — `{fileIDs[], types[], maxExtrafanart}`
- Progress via WS

---

## Phase 7 — Subtitles

### 7.1 Search
- `POST /api/subtitles/search` — search by file hash (OpenSubtitles) + title fallback
- Returns ranked subtitle list with language, release, uploader

### 7.2 Download
- `POST /api/subtitles/download` — `{subtitleID, language, targetPath}`
- Saves as `{episode-base}.{lang}.srt` / `.ass`

---

## Phase 8 — Verify Service

### 8.1 Hash computation
- SHA256 + MD5 per file
- Stream progress via WS
- `POST /api/verify` — `{filePaths[]}`
- Results cached in session

### 8.2 UI display
- Hash shown inline on file row after verification
- Compare mode: paste expected hash to verify match

---

## Phase 9 — SvelteKit Frontend

### 9.1 Stores
```
files.ts       — source file list, filter query, selection set
renamed.ts     — renamed preview list, filter query
config.ts      — app config (theme, sources, companions, rename op)
progress.ts    — WS progress events
metadata.ts    — selected file metadata + images
```

### 9.2 Components

| Component | Responsibility |
|---|---|
| `TopBar.svelte` | Tabs, theme picker, Execute button |
| `ThemePicker.svelte` | DaisyUI dropdown, 4-color strip swatches, ☀️/🌙 sections |
| `LeftPanel.svelte` | Source folder tree |
| `Toolbar.svelte` | Media type, DB badges, search, Pick ID |
| `OpStrip.svelte` | Move/Copy/Rename, destination input (required), folder checkbox |
| `FormatRow.svelte` | Rename template, presets, token reference |
| `FileList.svelte` | Virtualized scrollable list with Ctrl/Shift selection |
| `FileRow.svelte` | Single file/folder row, status dot, confidence |
| `SearchBar.svelte` | Live filter input (used in both columns) |
| `ArrowCol.svelte` | Arrow column between source and renamed |
| `RightPanel.svelte` | Collapsible wrapper with `◀/▶` button |
| `MetadataTab.svelte` | Series/episode info, rating, IDs, action buttons |
| `ImagesTab.svelte` | Type selector, image grid, extrafanart, download |
| `CompanionsTab.svelte` | Checkbox grid for all companion types |
| `NfoTab.svelte` | Template preview + edit |
| `IdPickerModal.svelte` | Full ID search/override modal |
| `SourcesModal.svelte` | Custom metadata source management |
| `ContextMenu.svelte` | Right-click menu |
| `StatusBar.svelte` | Connection dots, file counts, port |
| `ProgressToast.svelte` | WS operation progress overlay |

### 9.3 WebSocket client
- Connect on mount, auto-reconnect
- Dispatch `ProgressEvent` to `progress` store
- Show toast/overlay for active operations

### 9.4 Routing
Single-page app at `/`. All tabs rendered conditionally (no SvelteKit routing needed).

---

## Phase 10 — Docker + Deployment

### 10.1 Dockerfile
```dockerfile
FROM node:22-alpine AS frontend
WORKDIR /app/frontend
COPY frontend/package*.json ./
RUN npm ci
COPY frontend/ .
RUN npm run build

FROM golang:1.23-alpine AS backend
WORKDIR /app
COPY backend/ ./backend/
COPY --from=frontend /app/frontend/build ./backend/static/
RUN cd backend && go build -o bobmediamanager .

FROM alpine:3.20
COPY --from=backend /app/backend/bobmediamanager /app/bobmediamanager
EXPOSE 48040
CMD ["/app/bobmediamanager"]
```

### 10.2 docker-compose.yml
```yaml
services:
  bobmediamanager:
    build: .
    container_name: bobmediamanager
    ports:
      - "48040:48040"
    volumes:
      - ./config:/config
      - /media:/media:ro
    restart: unless-stopped
```

### 10.3 Traefik labels (optional, using `zerorouter-headers` middleware)
Standard Bob-stack Traefik setup, no additional CSP headers needed.

---

## Build Order Summary

| Phase | Deliverable | Depends on |
|---|---|---|
| 0 | Project scaffold, Dockerfile, Makefile | — |
| 1 | Go server, config, scanner | 0 |
| 2 | All metadata source adapters | 1 |
| 3 | Matcher + search API | 2 |
| 4 | Rename engine | 3 |
| 5 | NFO generator | 3 |
| 6 | Image downloader | 2 |
| 7 | Subtitles | 2 |
| 8 | Verify service | 1 |
| 9 | SvelteKit frontend (all components) | 1–8 (API contracts) |
| 10 | Docker, deployment | 9 |

Frontend development (Phase 9) can begin in parallel with Phases 2–8 using mocked API responses.
