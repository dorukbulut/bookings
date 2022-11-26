package render

import (
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"github.com/tsawler/bookings-app/internal/config"
	"github.com/tsawler/bookings-app/internal/models"
	"net/http"
	"os"
	"testing"
	"time"
)

var session *scs.SessionManager
var testApp config.AppConfig

type myWriter struct{}

func (w *myWriter) Header() http.Header {
	return http.Header{}
}

func (w *myWriter) WriteHeader(statusCode int) {}

func (w *myWriter) Write(b []byte) (int, error) {
	length := len(b)

	return length, nil
}

func TestMain(m *testing.M) {
	// what I'm going to put in my session.
	gob.Register(models.Reservation{})
	// change this to true when in production
	testApp.InProduction = false
	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session
	app = &testApp
	os.Exit(m.Run())
}
