package report

import (
	"image"
	"os"

	_ "image/jpeg"
	_ "image/png"
)

func orientationOf(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	cfg, _, err := image.DecodeConfig(f)
	if err != nil {
		return "", err
	}

	if cfg.Width > cfg.Height {
		return "Paisagem", nil
	}
	return "Retrato", nil
}
