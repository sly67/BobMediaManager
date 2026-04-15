# BobMediaManager

Self-hosted media file organizer — rename, sort, fetch metadata, artwork, NFO files and subtitles from TMDB, TVDB, AniDB, IMDb, MusicBrainz and custom sources.

**Port:** `48040`  
**Stack:** SvelteKit + DaisyUI v5 · Go · Docker

## Features

- Rename & sort movies, TV series, anime and music with customizable format templates
- Fetch metadata from TMDB, TVDB, AniDB, IMDb, MusicBrainz, Last.fm
- Custom metadata source registration (any REST API)
- Download all artwork types: poster, fanart, extrafanart, thumb, banner, logo, clearart, disc
- Generate Kodi/Jellyfin-compatible NFO companion files
- Grab subtitles via OpenSubtitles
- File hash verification (SHA256/MD5)
- Move / Copy / Rename-in-place operations with folder renaming support
- Manual ID override (`TMDB:693134`, `ANIDB:17243`, `IMDB:tt14491254`)
- 25 DaisyUI themes, Ctrl/Shift multi-select, collapsible inspector panel

## Quick start

```bash
docker compose up --build -d
```

Open `http://localhost:48040`

## Development

```bash
make dev
```

Requires Go 1.23+ and Node 22+.

## Design & Plan

See [`docs/superpowers/specs/`](docs/superpowers/specs/) for the full design spec and implementation plan.
