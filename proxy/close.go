package proxy

import (
	"io"
	"log"
)

func Close(closer io.Closer) {
	if err := closer.Close(); err != nil {
		log.Printf("Unable to close the closer. %v", err)
	}
}
