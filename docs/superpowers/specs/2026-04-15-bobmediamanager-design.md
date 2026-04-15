# BobMediaManager — Design Spec
**Date:** 2026-04-15  
**Port:** 48040  
**Stack:** SvelteKit + DaisyUI v5 · Go · Docker

---

## 1. Overview

BobMediaManager is a self-hosted media file organizer and metadata manager — a FileBot-style tool built as a web application. It renames, sorts, and enriches media files (movies, TV series, anime, music) by fetching metadata, artwork, NFO companion files, and subtitles from multiple online databases. All operations run server-side via a Go backend; the SvelteKit frontend provides a split-view workflow UI.

---

## 2. Scope

| In scope | Out of scope |
|---|---|
| File rename / move / copy | Media playback |
| Folder rename | Transcoding |
| Metadata fetch (TMDB, TVDB, AniDB, IMDb, MusicBrainz, Last.fm, custom) | Download automation |
| Artwork download (poster, fanart, extrafanart, thumb, banner, logo, clearart, disc) | Torrent management |
| NFO generation (Kodi/Jellyfin-compatible XML) | |
| Subtitle grabbing (OpenSubtitles) | |
| File signature / hash verification | |
| Multi-select, Ctrl/Shift click | |
| DaisyUI theming (25 themes) | |
| Custom metadata source registration | |

---

## 3. Architecture

```
Browser (SvelteKit)
        │  HTTP/WebSocket
        ▼
Go HTTP Server (:48040)
  ├── API router (chi)
  ├── File scanner service
  ├── Matcher service  (fuzzy name → DB query)
  ├── Metadata service (per-source adapters)
  ├── Rename engine    (Move / Copy / Rename-in-place)
  ├── NFO generator    (Go templates, Kodi schema)
  ├── Image downloader (parallel, typed: poster/fanart/extrafanart/…)
  ├── Subtitle grabber (OpenSubtitles REST API)
  ├── Verify service   (SHA256 / MD5)
  └── Config store     (JSON file, custom sources, theme pref)
```

WebSocket stream is used for long-running operations (rename batch, image downloads) to push progress to the frontend without polling.

---

## 4. Frontend (SvelteKit)

### 4.1 Layout

```
┌─ Topbar: logo · tabs · theme picker (DaisyUI) · Execute ─────────────┐
├─ Left: source folder tree (no destination shown here)                  │
├─ Center:                                                                │
│   ├─ Toolbar: media type · DB badges · search · Pick ID               │
│   ├─ Op strip: Move/Copy/Rename-in-place · Destination* · Browse      │
│   │            Rename folders too checkbox                              │
│   ├─ Format row: rename template · Presets · Tokens                   │
│   └─ Split view:                                                        │
│       [Source col + search] → [Arrow col] → [Renamed col + search]    │
└─ Right panel (collapsible): Metadata / Images / Companions / NFO tabs ┘
└─ Status bar: per-source connection dots · file counts · port           │
```

### 4.2 Tabs

| Tab | Purpose |
|---|---|
| Rename & Sort | Main workflow — source → renamed split view |
| Metadata | Fetch / edit metadata for a selected item |
| NFO | NFO template editor + preview |
| Images | Browse and download artwork per type |
| Subtitles | Search and grab subtitles (OpenSubtitles) |
| Verify | Hash verification (SHA256/MD5) of files |
| History | Operation log with undo |
| ⚙ Sources | Custom metadata source management |

### 4.3 ID Picker Modal

- Size: `min(92vw, 1080px) × min(88vh, 700px)`
- DB tab toggles: AniDB / TVDB / TMDB / IMDb / MusicBrainz
- Manual ID input accepting `SOURCE:ID` format (e.g. `TMDB:693134`, `ANIDB:17243`, `IMDB:tt14491254`)
- Search field with ranked results
- Each result shows: poster thumbnail, title, year, plot excerpt, all matched IDs as tags, confidence %, per-DB link-back buttons (`↗ AniDB`, `↗ TVDB`, etc.)
- Footer: selected ID display string + Confirm / Cancel

### 4.4 Theme System

25 DaisyUI themes applied via `data-theme` on `<html>`. Theme picker dropdown shows two sections (☀️ Light / 🌙 Dark), each theme rendered as a 4-color strip (bg · primary · secondary · accent). Theme preference persisted in localStorage and synced to backend config.

### 4.5 File Selection

- Single click: select one row
- Ctrl/Cmd+click: toggle add/remove from selection
- Shift+click: range select from last clicked
- Context menu (right-click): Search/Override ID, Re-match, Edit renamed path, Generate NFO, Download images, Grab subtitles, Verify signature, Select all matched, Execute selected, Copy paths, Remove from list

### 4.6 Companions

Per-file companion generation is fully configurable via checkboxes:

- NFO: `tvshow.nfo`, `episode.nfo`, `movie.nfo`, `artist.nfo`, `album.nfo`
- Series images: `poster.jpg`, `fanart.jpg`, `extrafanart/` (configurable max count), `banner.jpg`, `logo.png`, `clearart.png`, `landscape.jpg`, `disc.png`
- Season images: `season{N}-poster.jpg`, `season{N}-fanart.jpg`, `season{N}-banner.jpg`
- Episode images: `{ep}-thumb.jpg`
- Music: `folder.jpg`, `artist.jpg`
- Subtitles: `{ep}.{lang}.srt` per selected language

---

## 5. Backend (Go)

### 5.1 API Endpoints

```
GET  /api/sources            list configured metadata sources
POST /api/sources            add custom source
PUT  /api/sources/:id        update
DELETE /api/sources/:id      remove

POST /api/scan               scan a directory, return file list
POST /api/match              run fuzzy match on file list
POST /api/search             search a specific DB for a query
GET  /api/metadata/:source/:id   fetch metadata for a resolved ID
POST /api/rename             execute rename/move/copy batch
POST /api/nfo                generate NFO files for matched items
POST /api/images             download images for matched items
POST /api/subtitles/search   search subtitles
POST /api/subtitles/download download selected subtitles
POST /api/verify             compute hashes for files

WS   /ws/progress            stream operation progress events
```

### 5.2 Metadata Sources

Each source implements a common interface:

```go
type MetadataSource interface {
    Search(ctx, query string, mediaType MediaType) ([]SearchResult, error)
    GetByID(ctx, id string) (*MediaMetadata, error)
    Name() string
    Slug() string
    SupportedTypes() []MediaType
}
```

Built-in sources: `AniDB`, `TVDB`, `TMDB`, `IMDb`, `MusicBrainz`, `LastFM`, `OpenSubtitles`.

Custom sources are stored in `config.json` and loaded at startup as generic REST adapters with a user-defined field mapping.

### 5.3 Rename Engine

Operations: `Move`, `Copy`, `RenameInPlace`.

Format tokens (FileBot-compatible subset):
`{n}` title · `{y}` year · `{s}` season · `{e}` episode · `{s00}` zero-padded season · `{e00}` zero-padded episode · `{t}` episode title · `{artist}` · `{album}` · `{track00}` · `{lang}` · `{ext}`.

Folder renaming: when "Rename folders too" is enabled, the engine renames parent directories using the same template's directory segments.

### 5.4 NFO Generator

Go templates producing Kodi-compatible XML. Templates are user-editable. Supported NFO types: `tvshow.nfo`, `episodedetails`, `movie`, `musicvideo`, `artist`, `album`.

### 5.5 Image Downloader

Downloads run concurrently (worker pool, configurable concurrency). Image types: poster, fanart, extrafanart (saved to `extrafanart/` subfolder), thumb, banner, logo, clearart, disc, characterart. Sources: TMDB image API, TVDB artwork API, Fanart.tv (optional).

### 5.6 Subtitle Grabber

OpenSubtitles REST API v1. Search by file hash + title fallback. Download `.srt` / `.ass` to `{episode-base}.{lang}.srt` alongside the video file.

### 5.7 Verify Service

SHA256 and MD5 hash computation with progress streaming. Results shown inline on the file row. Optional: compare against `.nfo`-embedded hash or torrent file.

---

## 6. Data Models

```go
type MediaFile struct {
    Path        string
    Name        string
    Size        int64
    IsFolder    bool
    MatchStatus MatchStatus  // matched | pending | unmatched
    Confidence  float64
    ResolvedIDs map[string]string  // {"anidb":"17243","tmdb":"209867",...}
    Metadata    *MediaMetadata
    RenamedPath string
}

type MediaMetadata struct {
    Title       string
    OrigTitle   string
    Year        int
    Overview    string
    Rating      float64
    Genres      []string
    Studio      string
    EpisodeNum  int
    SeasonNum   int
    EpisodeTitle string
    AiredDate   string
    IDs         map[string]string
    Images      map[ImageType][]ImageEntry
}

type CustomSource struct {
    ID          string
    Name        string
    BaseURL     string
    APIKey      string
    MediaTypes  []MediaType
    FieldMap    map[string]string  // our field → JSON path in response
}
```

---

## 7. Configuration

`config.json` (persisted in Docker volume):
```json
{
  "port": 48040,
  "theme": "dark",
  "sources": { "anidb": true, "tvdb": true, "tmdb": true, "imdb": true },
  "apiKeys": { "tvdb": "...", "tmdb": "...", "opensubtitles": "..." },
  "customSources": [],
  "companions": { "poster": true, "fanart": true, "extrafanart": true, "maxExtrafanart": 5, "nfo": true },
  "rename": { "defaultOperation": "move", "renameFolders": false }
}
```

---

## 8. Docker

```yaml
services:
  bobmediamanager:
    build: .
    ports:
      - "48040:48040"
    volumes:
      - ./config:/config
      - /media:/media
    restart: unless-stopped
```

Single-stage build: Go binary embedded with SvelteKit static output (via `embed.FS`).

---

## 9. Non-Goals

- No user authentication (single-user, self-hosted)
- No database — all state is in-memory per session; config persisted to JSON
- No media player / streaming
