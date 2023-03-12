package pdf

import (
	"bytes"
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func GeneratePdf(pages map[int][]byte, w io.Writer) {
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
		panic(err)
	}

	pagesIndRef, err := ctx.Pages()
	if err != nil {
		panic(err)
	}

	pagesDict, err := ctx.DereferenceDict(*pagesIndRef)
	if err != nil {
		panic(err)
	}

	for i := 1; i < len(pages); i++ {
		if page, has := pages[i]; has {
			r := bytes.NewReader(page)
			indRef, _ := pdfcpu.NewPageForImage(ctx.XRefTable, r, pagesIndRef, imp)

			if err = model.AppendPageTree(indRef, 1, pagesDict); err != nil {
				panic(err)
			}

			ctx.PageCount++
		}
	}

	api.WriteContext(ctx, w)

}
