package cookieauth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
)

func TestAll(t *testing.T) {
	//secret handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`hello world`))
	})
	//protect with cookieauth
	ca := New()
	ca.SetUserPass("foo", "bar")
	// ca.SetLogger(log.New(os.Stdout, "", log.LstdFlags))
	protected := ca.Wrap(handler)
	//start server
	server := httptest.NewServer(protected)
	defer server.Close()
	//begin
	e := httpexpect.New(t, server.URL)
	e.GET("/").Expect().Status(http.StatusUnauthorized)
	e.GET("/").WithBasicAuth("bazz", "bar").Expect().Status(http.StatusUnauthorized)
	c := e.GET("/").WithBasicAuth("foo", "bar").Expect().Status(http.StatusOK).Cookie("cookieauth")
	e.GET("/").WithCookie("cookieauth", "incorrect").Expect().Status(http.StatusUnauthorized)
	e.GET("/").WithCookie("cookieauth", c.Value().Raw()).Expect().Status(http.StatusOK)
	ca.SetUserPass("zip", "zop")
	e.GET("/").WithCookie("cookieauth", c.Value().Raw()).Expect().Status(http.StatusUnauthorized)
	c = e.GET("/").WithBasicAuth("zip", "zop").Expect().Status(http.StatusOK).Cookie("cookieauth")
	e.GET("/").WithCookie("cookieauth", c.Value().Raw()).Expect().Status(http.StatusOK)
}
