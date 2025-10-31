package report

import "imageReportGenV2/internal/report/model"

func NewDefaultConfig() model.ReportConfig {
	return model.ReportConfig{
		TemplatePath: "",
		OutputPath:   "relatorio_gerado.docx",
		Header:       model.HeaderFields{},
		Layout: model.LayoutConfig{
			Preset:    model.LayoutDefault,
			Landscape: model.Grid{Rows: 3, Cols: 1}, // 3 imagens/página p/ paisagem
			Portrait:  model.Grid{Rows: 2, Cols: 2}, // 4 imagens/página p/ retrato
		},
	}
}
