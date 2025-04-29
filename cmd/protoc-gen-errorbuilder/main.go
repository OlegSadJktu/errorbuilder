package main

import (
	"github.com/OlegSadJktu/protoc-gen-errorbuilder/pkg"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {

	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		alreadyGenerated := make(map[string]bool)
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			if f.Proto.Package == nil {
				continue
			}
			if _, ok := alreadyGenerated[*f.Proto.Package]; ok {
				continue
			}
			alreadyGenerated[*f.Proto.Package] = true

			generateFile(gen, f)
		}
		return nil
	})

}

func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	filename := file.GeneratedFilenamePrefix + ".errorbuilder.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	pkg.GenerateHeader(g, file.GoPackageName)

	pkg.GenerateLocalFunctions(g)
	pkg.GenerateErrorBuilderMethods(g, file)

	return g
}
