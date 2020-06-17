// toopenapi3 relies on `swagger2openapi` to convert
// Swagger 2.0 specs to OpenAPI 3.0 specs.
package toopenapi3

import (
	"errors"
	"path/filepath"
	"strings"

	"github.com/grokify/gotilla/cmd/cmdutil"
	"github.com/grokify/gotilla/errors/errorsutil"
	"github.com/grokify/gotilla/type/stringsutil"
)

const cmdSwagger2OpenAPI = "swagger2openapi"

func Convert(filenames []string, outdir string, renameFunc func(string) string) errorsutil.ErrorInfos {
	errinfos := errorsutil.ErrorInfos{}
	for _, srcpath := range filenames {
		_, srcfile := filepath.Split(srcpath)
		outfile := renameFunc(srcfile)
		outpath := filepath.Join(outdir, outfile)

		qtr := stringsutil.Quoter{Beg: "", End: ""}
		cmd := strings.Join([]string{
			cmdSwagger2OpenAPI,
			qtr.Quote(srcpath)}, " ")

		_, stderr, err := cmdutil.ExecToFiles(cmd, outpath, "", 0644)
		if err != nil {
			if err.Error() == "exit status 1" {
				ei := errorsutil.ErrorInfo{
					Input:   srcpath,
					Correct: outpath,
					Display: cmd,
					Error:   errors.New(stderr.String())}
				errinfos = append(errinfos, &ei)
			} else {
				ei := errorsutil.ErrorInfo{
					Input:   srcpath,
					Correct: outpath,
					Display: cmd,
					Error:   err}
				errinfos = append(errinfos, &ei)
			}
		} else {
			ei := errorsutil.ErrorInfo{
				Input:   srcpath,
				Correct: outpath,
				Display: cmd,
				Error:   nil}
			errinfos = append(errinfos, &ei)
		}
	}
	errinfos.Inflate()
	return errinfos
}