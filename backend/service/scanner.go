package service

import (
	"os"
	"path/filepath"
	"strings"

	"bobmediamanager/models"
)

var videoExts = map[string]bool{
	".mkv": true, ".mp4": true, ".avi": true, ".mov": true,
	".wmv": true, ".m4v": true, ".ts":  true, ".m2ts": true,
	".webm": true,
}

var audioExts = map[string]bool{
	".flac": true, ".mp3": true, ".m4a": true, ".ogg": true,
	".opus": true, ".wav": true, ".aac": true, ".wma": true,
}

// ScanRequest holds parameters for a directory scan.
type ScanRequest struct {
	Path      string           `json:"path"`
	Recursive bool             `json:"recursive"`
	MediaType models.MediaType `json:"mediaType"`
}

// Scan walks a directory and returns matching media files.
func Scan(req ScanRequest) ([]models.MediaFile, error) {
	var files []models.MediaFile

	walkFn := func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil // skip unreadable entries
		}
		if d.IsDir() {
			if path != req.Path && !req.Recursive {
				return filepath.SkipDir
			}
			return nil
		}

		ext := strings.ToLower(filepath.Ext(path))
		if !isMatchingExt(ext, req.MediaType) {
			return nil
		}

		info, err := d.Info()
		if err != nil {
			return nil
		}

		files = append(files, models.MediaFile{
			ID:          generateID(path),
			Path:        path,
			Name:        d.Name(),
			Size:        info.Size(),
			IsFolder:    false,
			MediaType:   req.MediaType,
			MatchStatus: models.MatchStatusPending,
		})
		return nil
	}

	if err := filepath.WalkDir(req.Path, walkFn); err != nil {
		return nil, err
	}
	return files, nil
}

func isMatchingExt(ext string, mt models.MediaType) bool {
	switch mt {
	case models.MediaTypeMusic:
		return audioExts[ext]
	default:
		return videoExts[ext]
	}
}

func generateID(path string) string {
	// Simple deterministic ID from path (hash not needed at scaffold stage)
	sum := 0
	for _, c := range path {
		sum = sum*31 + int(c)
	}
	if sum < 0 {
		sum = -sum
	}
	return filepath.Base(path) + "_" + itoa(sum)
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	buf := [20]byte{}
	pos := len(buf)
	for n > 0 {
		pos--
		buf[pos] = byte('0' + n%10)
		n /= 10
	}
	return string(buf[pos:])
}
