// Code generated by 'goexports archive/zip'. DO NOT EDIT.

// +build go1.12,!go1.13

package stdlib

import (
	"archive/zip"
	"reflect"
)

func init() {
	Symbols["archive/zip"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Deflate":              reflect.ValueOf(zip.Deflate),
		"ErrAlgorithm":         reflect.ValueOf(&zip.ErrAlgorithm).Elem(),
		"ErrChecksum":          reflect.ValueOf(&zip.ErrChecksum).Elem(),
		"ErrFormat":            reflect.ValueOf(&zip.ErrFormat).Elem(),
		"FileInfoHeader":       reflect.ValueOf(zip.FileInfoHeader),
		"NewReader":            reflect.ValueOf(zip.NewReader),
		"NewWriter":            reflect.ValueOf(zip.NewWriter),
		"OpenReader":           reflect.ValueOf(zip.OpenReader),
		"RegisterCompressor":   reflect.ValueOf(zip.RegisterCompressor),
		"RegisterDecompressor": reflect.ValueOf(zip.RegisterDecompressor),
		"Store":                reflect.ValueOf(zip.Store),

		// type definitions
		"Compressor":   reflect.ValueOf((*zip.Compressor)(nil)),
		"Decompressor": reflect.ValueOf((*zip.Decompressor)(nil)),
		"File":         reflect.ValueOf((*zip.File)(nil)),
		"FileHeader":   reflect.ValueOf((*zip.FileHeader)(nil)),
		"ReadCloser":   reflect.ValueOf((*zip.ReadCloser)(nil)),
		"Reader":       reflect.ValueOf((*zip.Reader)(nil)),
		"Writer":       reflect.ValueOf((*zip.Writer)(nil)),
	}
}
