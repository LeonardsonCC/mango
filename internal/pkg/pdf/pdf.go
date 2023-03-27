package pdf

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

type ErrFailedToAddImage struct {
	Page int
}

func (e *ErrFailedToAddImage) Error() string {
	return fmt.Sprintf("failed to add page %d to manga pdf", e.Page)
}

type Warnings struct {
	Warns []error
}

func (e *Warnings) Error() string {
	strs := make([]string, len(e.Warns))
	for i, w := range e.Warns {
		strs[i] = w.Error()
	}

	return strings.Join(strs, "\n")
}

func GeneratePdf(pages map[int][]byte, w io.Writer) error {
	// generate pdf
	conf := model.NewDefaultConfiguration()
	conf.Cmd = model.IMPORTIMAGES
	imp := pdfcpu.DefaultImportConfig()

	var (
		ctx *model.Context
		err error
	)

	ctx, err = pdfcpu.CreateContextWithXRefTable(conf, imp.PageDim)
	if err != nil {
		return err
	}

	pagesIndRef, err := ctx.Pages()
	if err != nil {
		return err
	}

	pagesDict, err := ctx.DereferenceDict(*pagesIndRef)
	if err != nil {
		return err
	}

	imgWarnings := new(Warnings)

	for i := 1; i < len(pages); i++ {
		if page, has := pages[i]; has && page != nil {
			indRef, err := pdfcpu.NewPageForImage(ctx.XRefTable, bytes.NewReader(page), pagesIndRef, imp)
			// just ignore the image errors
			if err != nil {
				imgWarnings.Warns = append(imgWarnings.Warns, &ErrFailedToAddImage{i})
				continue
			}

			if err = model.AppendPageTree(indRef, 1, pagesDict); err != nil {
				return err
			}

			ctx.PageCount++
		}
	}

	api.WriteContext(ctx, w)

	return imgWarnings
}
