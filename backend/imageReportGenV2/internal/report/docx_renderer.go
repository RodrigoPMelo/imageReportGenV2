package report

import (
	"fmt"
	"strings"

	"imageReportGenV2/internal/report/model"

	"github.com/gomutex/godocx"
	"github.com/gomutex/godocx/common/units"
	"github.com/gomutex/godocx/docx"
)

const (
	a4WidthIn     = 8.27
	a4HeightIn    = 11.69
	defaultMargin = 0.5
)

type cellSize struct {
	WidthIn  float64
	HeightIn float64
}

func RenderDOCX(cfg model.ReportConfig, images []model.ImageInfo) error {
	doc, err := godocx.OpenDocument(cfg.TemplatePath)
	if err != nil {
		return fmt.Errorf("open template: %w", err)
	}
	defer doc.Close()

	// header opcional
	if hasHeader(cfg.Header) {
		p := doc.AddEmptyParagraph()
		p.AddText(buildHeaderLine(cfg.Header))
		doc.AddPageBreak()
	}

	// separa por orientação
	var landscape, portrait []string
	for _, img := range images {
		if img.Orientation == "Paisagem" {
			landscape = append(landscape, img.Path)
		} else {
			portrait = append(portrait, img.Path)
		}
	}

	landGrid := cfg.Layout.Landscape
	portGrid := cfg.Layout.Portrait
	if cfg.Layout.Preset == model.LayoutDefault {
		landGrid = model.Grid{Rows: 3, Cols: 1}
		portGrid = model.Grid{Rows: 2, Cols: 2}
	}

	if len(landscape) > 0 {
		if err := addGridPages(doc, landscape, landGrid); err != nil {
			return err
		}
		if len(portrait) > 0 {
			doc.AddPageBreak()
		}
	}

	if len(portrait) > 0 {
		if err := addGridPages(doc, portrait, portGrid); err != nil {
			return err
		}
	}

	out := strings.TrimSpace(cfg.OutputPath)
	if out == "" {
		out = "relatorio_gerado.docx"
	}

	return doc.SaveTo(out)
}

func addGridPages(doc *docx.RootDoc, paths []string, grid model.Grid) error {
	sz := calcCellSize(grid)
	w := units.Inch(sz.WidthIn)
	h := units.Inch(sz.HeightIn)

	i := 0
	for i < len(paths) {
		if i > 0 {
			doc.AddPageBreak()
		}
		t := doc.AddTable()
		for r := 0; r < grid.Rows && i < len(paths); r++ {
			row := t.AddRow()
			for c := 0; c < grid.Cols && i < len(paths); c++ {
				cell := row.AddCell()
				if _, err := cell.AddEmptyPara().AddPicture(paths[i], w, h); err != nil {
					return err
				}
				i++
			}
		}
	}
	return nil
}

func calcCellSize(grid model.Grid) cellSize {
	cols := grid.Cols
	if cols < 1 {
		cols = 1
	}
	rows := grid.Rows
	if rows < 1 {
		rows = 1
	}
	usableW := a4WidthIn - 2*defaultMargin
	usableH := a4HeightIn - 2*defaultMargin
	return cellSize{
		WidthIn:  usableW / float64(cols),
		HeightIn: usableH / float64(rows),
	}
}

func hasHeader(h model.HeaderFields) bool {
	return (h.IncludeClient && h.ClientName != "") ||
		(h.IncludeLocation && (h.City != "" || h.State != "")) ||
		h.IncludeDate
}

func buildHeaderLine(h model.HeaderFields) string {
	parts := []string{}
	if h.IncludeClient && h.ClientName != "" {
		parts = append(parts, "Cliente: "+h.ClientName)
	}
	if h.IncludeLocation && (h.City != "" || h.State != "") {
		loc := strings.TrimSpace(strings.Join([]string{h.City, h.State}, "/"))
		parts = append(parts, "Local: "+loc)
	}
	if h.IncludeDate {
		parts = append(parts, "Conclusão: "+h.CompletionDate.Format("2006-01-02"))
	}
	return strings.Join(parts, " | ")
}
