package pkg

import "google.golang.org/protobuf/compiler/protogen"

const getMsgFunc = `func getMsg(msg []string) string {
	var message string
	if len(msg) > 0 {
		message = msg[0]
	}
	return message
}
`

func GenerateLocalFunctions(gen *protogen.GeneratedFile) {

	gen.P(getMsgFunc)
}
