package session

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/pkg/errors"
)

const (
	//SessionName to store session under
	SessionName = "mylopodsession"
)

var (
	Store sessions.Store
)

// determineEncryptionKey ensures the provided SESSION_ENCRYPTION_KEY is the
// correct size (16, 24 or 32 bytes). If it's too large it's truncated to the
// max. If it's otherwise incorrect size wise an error is returned. Otherwise
// the []byte version is returned.
func DetermineEncryptionKey() ([]byte, error) {
	sek := os.Getenv("SESSION_ENCRYPTION_KEY")
	lek := len(sek)
	switch {
	case lek >= 0 && lek < 16, lek > 16 && lek < 24, lek > 24 && lek < 32:
		return nil, errors.Errorf("SESSION_ENCRYPTION_KEY needs to be either 16, 24 or 32 characters long or longer, was: %d", lek)
	case lek == 16, lek == 24, lek == 32:
		return []byte(sek), nil
	case lek > 32:
		return []byte(sek[0:32]), nil
	default:
		return nil, errors.New("invalid SESSION_ENCRYPTION_KEY: " + sek)
	}

}

func HandleSessionError(w http.ResponseWriter, err error) {
	http.Error(w, "Application Error", http.StatusInternalServerError)
}
