package common

type FilenameMessage struct {
	Filename string
	Dst string
}

func InitChannel() chan FilenameMessage {
	return make(chan FilenameMessage)
}
