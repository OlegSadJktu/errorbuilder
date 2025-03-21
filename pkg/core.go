package pkg

import (
	"fmt"
	"slices"

	"google.golang.org/protobuf/compiler/protogen"
)

const buildErrorFormatString = `
func Build%sError(code %sErrorCode, msg ...string) (*%vResponse, error) {
	message := getMsg(msg)
	return &%vResponse{
		Error: &%vError{
			ErrorCode:    code,
			ErrorMessage: message,
		},
	}, nil
}
`

func buildErrorBuilder(methodName string) string {
	return fmt.Sprintf(
		buildErrorFormatString,
		methodName,
		methodName,
		methodName,
		methodName,
		methodName,
	)
}

func GenerateErrorBuilderMethods(g *protogen.GeneratedFile, src *protogen.File) {
	var enums []string
	for _, enum := range src.Enums {
		name := string(enum.Desc.Name())
		enums = append(enums, name)
	}

	for _, service := range src.Services {
		for _, method := range service.Methods {
			name := method.GoName
			errorName := name + "ErrorCode"
			if slices.Contains(enums, errorName) {
				formatedFunc := buildErrorBuilder(name)
				g.P(formatedFunc)
			}
		}
	}
}
