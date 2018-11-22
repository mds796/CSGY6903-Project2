# DropPoint

A secure file server proxy that securely stores and retrieves files on an insecure file server.

DropPoint provides:

1. a secure HTTPS encryption for requests.
1. Secure storage of file contents using symmetric encryption.
1. File content verification using digital signatures.
1. Mitigations for replay attacks to prevent a malicious actor from sniffing your requests in order to obtain your file contents.  