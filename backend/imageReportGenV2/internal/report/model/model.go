package model

import "time"

type HeaderFields struct {
	IncludeClient   bool      `json:"includeClient"`
	ClientName      string    `json:"clientName"`
	IncludeLocation bool      `json:"includeLocation"`
	City            string    `json:"city"`
	State           string    `json:"state"`
	IncludeDate     bool      `json:"includeDate"`
	CompletionDate  time.Time `json:"completionDate"`
}

type LayoutPreset string

const (
	LayoutDefault LayoutPreset = "default"
	LayoutCustom  LayoutPreset = "custom"
)

type Grid struct {
	Rows int `json:"rows"`
	Cols int `json:"cols"`
}

type LayoutConfig struct {
	Preset    LayoutPreset `json:"preset"`
	Landscape Grid         `json:"landscape"`
	Portrait  Grid         `json:"portrait"`
}

type ReportConfig struct {
	TemplatePath string       `json:"templatePath"`
	OutputPath   string       `json:"outputPath"`
	Header       HeaderFields `json:"header"`
	Layout       LayoutConfig `json:"layout"`
}

type ImageInfo struct {
	Path        string `json:"path"`
	Orientation string `json:"orientation"`
}
