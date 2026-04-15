package models

// MediaType identifies the category of a media file.
type MediaType string

const (
	MediaTypeMovie  MediaType = "movie"
	MediaTypeSeries MediaType = "series"
	MediaTypeAnime  MediaType = "anime"
	MediaTypeMusic  MediaType = "music"
)

// MatchStatus reflects the matching state of a scanned file.
type MatchStatus string

const (
	MatchStatusPending   MatchStatus = "pending"
	MatchStatusMatched   MatchStatus = "matched"
	MatchStatusUnmatched MatchStatus = "unmatched"
)

// ImageType enumerates all supported artwork variants.
type ImageType string

const (
	ImageTypePoster       ImageType = "poster"
	ImageTypeFanart       ImageType = "fanart"
	ImageTypeExtrafanart  ImageType = "extrafanart"
	ImageTypeThumb        ImageType = "thumb"
	ImageTypeBanner       ImageType = "banner"
	ImageTypeLogo         ImageType = "logo"
	ImageTypeClearArt     ImageType = "clearart"
	ImageTypeDisc         ImageType = "disc"
	ImageTypeCharacterArt ImageType = "characterart"
)

// ImageEntry holds the URL and metadata for a single artwork image.
type ImageEntry struct {
	URL      string    `json:"url"`
	Type     ImageType `json:"type"`
	Language string    `json:"language,omitempty"`
	Width    int       `json:"width,omitempty"`
	Height   int       `json:"height,omitempty"`
	Rating   float64   `json:"rating,omitempty"`
}

// MediaFile represents a scanned file on disk.
type MediaFile struct {
	ID          string            `json:"id"`
	Path        string            `json:"path"`
	Name        string            `json:"name"`
	Size        int64             `json:"size"`
	IsFolder    bool              `json:"isFolder"`
	MediaType   MediaType         `json:"mediaType"`
	MatchStatus MatchStatus       `json:"matchStatus"`
	Confidence  float64           `json:"confidence"`
	ResolvedIDs map[string]string `json:"resolvedIds"` // e.g. {"anidb":"17243","tmdb":"209867"}
	Metadata    *MediaMetadata    `json:"metadata,omitempty"`
	RenamedPath string            `json:"renamedPath,omitempty"`
	Hashes      *FileHashes       `json:"hashes,omitempty"`
}

// MediaMetadata holds enriched information fetched from a metadata source.
type MediaMetadata struct {
	Title        string                       `json:"title"`
	OrigTitle    string                       `json:"origTitle,omitempty"`
	Year         int                          `json:"year,omitempty"`
	Overview     string                       `json:"overview,omitempty"`
	Rating       float64                      `json:"rating,omitempty"`
	Genres       []string                     `json:"genres,omitempty"`
	Studio       string                       `json:"studio,omitempty"`
	SeasonNum    int                          `json:"seasonNum,omitempty"`
	EpisodeNum   int                          `json:"episodeNum,omitempty"`
	EpisodeTitle string                       `json:"episodeTitle,omitempty"`
	AiredDate    string                       `json:"airedDate,omitempty"`
	IDs          map[string]string            `json:"ids,omitempty"`
	Images       map[ImageType][]ImageEntry   `json:"images,omitempty"`
	// Music fields
	Artist       string `json:"artist,omitempty"`
	Album        string `json:"album,omitempty"`
	TrackNum     int    `json:"trackNum,omitempty"`
}

// SearchResult is returned from source search queries.
type SearchResult struct {
	ID         string            `json:"id"`
	SourceSlug string            `json:"sourceSlug"`
	Title      string            `json:"title"`
	Year       int               `json:"year,omitempty"`
	Overview   string            `json:"overview,omitempty"`
	PosterURL  string            `json:"posterUrl,omitempty"`
	IDs        map[string]string `json:"ids,omitempty"`
	Confidence float64           `json:"confidence"`
}

// FileHashes holds computed checksums for a media file.
type FileHashes struct {
	SHA256 string `json:"sha256,omitempty"`
	MD5    string `json:"md5,omitempty"`
}

// ProgressEvent is sent over the WebSocket connection.
type ProgressEvent struct {
	Operation string  `json:"operation"` // rename | images | nfo | verify | subtitles
	FileID    string  `json:"fileId,omitempty"`
	Total     int     `json:"total"`
	Done      int     `json:"done"`
	Error     string  `json:"error,omitempty"`
	Status    string  `json:"status"` // running | completed | error
}

// OperationLog records a rename/move/copy action for history & undo.
type OperationLog struct {
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Operation string `json:"operation"` // move | copy | rename
	Source    string `json:"source"`
	Dest      string `json:"dest"`
}
