package https

import (
	"net/http"

	"github.com/mia0x75/pages/configuration"
	"github.com/mia0x75/pages/filenames"
)

// StartServer TODO
func StartServer(addr string, handler http.Handler) error {
	if configuration.Config.UseLetsEncrypt {
		server := buildLetsEncryptServer(addr, handler)
		return server.ListenAndServeTLS("", "")
	}
	checkCertificates()
	return http.ListenAndServeTLS(addr, filenames.HttpsCertFilename, filenames.HttpsKeyFilename, handler)
}
