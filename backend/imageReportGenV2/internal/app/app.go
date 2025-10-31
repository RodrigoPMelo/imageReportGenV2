package app

import (
	"context"
	"time"

	"imageReportGenV2/internal/report"
	"imageReportGenV2/internal/report/model"
)

type ImageReportApp struct {
	ctx context.Context
}

func NewImageReportApp() *ImageReportApp {
	return &ImageReportApp{}
}

func (a *ImageReportApp) Startup(ctx context.Context) {
	a.ctx = ctx
}

type GenerateRequest struct {
	TemplatePath    string   `json:"templatePath"`
	OutputPath      string   `json:"outputPath"`
	ClientName      string   `json:"clientName"`
	City            string   `json:"city"`
	State           string   `json:"state"`
	IncludeClient   bool     `json:"includeClient"`
	IncludeLocation bool     `json:"includeLocation"`
	IncludeDate     bool     `json:"includeDate"`
	CompletionDate  string   `json:"completionDate"` // yyyy-mm-dd
	Images          []string `json:"images"`
}

func (a *ImageReportApp) GenerateReport(req GenerateRequest) (string, error) {
	cfg := report.NewDefaultConfig()
	cfg.TemplatePath = req.TemplatePath
	if req.OutputPath != "" {
		cfg.OutputPath = req.OutputPath
	}
	cfg.Header = model.HeaderFields{
		IncludeClient:   req.IncludeClient,
		ClientName:      req.ClientName,
		IncludeLocation: req.IncludeLocation,
		City:            req.City,
		State:           req.State,
		IncludeDate:     req.IncludeDate,
	}

	// parse data se vier
	if req.IncludeDate && req.CompletionDate != "" {
		if t, err := time.Parse("2006-01-02", req.CompletionDate); err == nil {
			cfg.Header.CompletionDate = t
		}
	}

	// resolve orientação das imagens
	imgInfos, err := report.BuildImageInfos(req.Images)
	if err != nil {
		return "", err
	}

	if err := report.RenderDOCX(cfg, imgInfos); err != nil {
		return "", err
	}

	return cfg.OutputPath, nil
}
