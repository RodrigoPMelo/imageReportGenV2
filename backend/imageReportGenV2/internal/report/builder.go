package report

import "imageReportGenV2/internal/report/model"

func BuildImageInfos(paths []string) ([]model.ImageInfo, error) {
	result := make([]model.ImageInfo, 0, len(paths))
	for _, p := range paths {
		o, err := orientationOf(p)
		if err != nil {
			return nil, err
		}
		result = append(result, model.ImageInfo{
			Path:        p,
			Orientation: o,
		})
	}
	return result, nil
}
