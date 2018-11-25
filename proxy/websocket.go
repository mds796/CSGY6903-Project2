package proxy

import (
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func hijack(destinationTarget *url.URL, w http.ResponseWriter, r *http.Request) {
	h, ok := w.(http.Hijacker)
	if !ok {
		log.Println("Failed to hijack HTTP connection. Response does not implement http.Hijacker.")
		return
	}

	bytes, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Printf("Failed to dump the HTTP request to a byte slice. %v", err)
		return
	}

	client, buffered, err := h.Hijack()
	if err != nil {
		log.Printf("Failed to hijack HTTP connection. %v", err)
	}

	defer Close(client)

	if buffered.Reader.Buffered() > 0 {
		log.Println("Failed to hijack HTTP connection. Client sent data before handshake is complete.")
		return
	}

	failed := make(chan bool, 2)

	destination, err := forward(destinationTarget, client, failed)
	if err != nil {
		log.Printf("Encountered an eror while forwarding Web Socket connection. %v", err)
		return
	}

	defer Close(destination)

	_, err = destination.Write(bytes)
	if err != nil {
		log.Printf("Failed to write the upgrade HTTP request to destination server. %v", err)
		return
	}

	<-failed
}

func forward(destinationTarget *url.URL, conn net.Conn, failed chan bool) (net.Conn, error) {
	destination, err := net.Dial("tcp", destinationTarget.Host)

	if err != nil {
		log.Printf("Failed to dial to destination server. %v", err)
		return nil, err
	}

	go func() {
		if _, err := io.Copy(destination, conn); err != nil {
			log.Printf("Encountered an eror copying client connection to destination. %v", err)
			failed <- true
		}
	}()

	go func() {
		if _, err := io.Copy(conn, destination); err != nil {
			log.Printf("Encountered an eror copying destination connection to client. %v", err)
			failed <- true
		}
	}()

	return destination, nil
}
