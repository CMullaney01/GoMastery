package typedfunc

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type TransformFunc func(string) string

type Server struct {
	filenameTransformFunc TransformFunc
}

func (s *Server) handleRequest(filename string) error {
	newFilename := s.filenameTransformFunc(filename)

	fmt.Println("new filename: ", newFilename)

	return nil
}

// mediocre -- typed functions (note: we could also use an interface for this but if the thing we want to do does not take any state we can use typed function)
// all these filename shenanigans can be tested individually
func hashFilename(filename string) string {
	hash := sha256.Sum256([]byte(filename))
	newFilename := hex.EncodeToString(hash[:])

	return newFilename
}

func prefixFilename(prefix string) TransformFunc {
	return func(filename string) string {
		return prefix + filename
	}
}

func TypedFunc() {
	s := &Server{
		filenameTransformFunc: prefixFilename("BOB_"),
	}

	s.handleRequest("cool_picture.jpg")
}
