package main

import (
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/dave/jennifer/jen"
	"github.com/liamawhite/protoc-gen-terraform/pkg/generate"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			generateFile(gen, f)
		}
		return nil
	})
}

// generateFile generates a _ascii.pb.go file containing gRPC service definitions.
func generateFile(gen *protogen.Plugin, file *protogen.File) {
	filename := file.GeneratedFilenamePrefix + "_terraform.go"

	f := jen.NewFile(string(file.GoPackageName))
	for _, m := range file.Messages {
		generate.Scheme(f, m)
	}

	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-terraform. DO NOT EDIT.")
	g.P(f.GoString())
}