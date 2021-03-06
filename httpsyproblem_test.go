package httpsyproblem

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testEmbeddedDetails struct {
	Details
	ID string `json:"id" xml:"id"`
}

func TestEmbed(t *testing.T) {
	detail := &testEmbeddedDetails{
		Details: *New(http.StatusBadRequest, nil),
		ID:      "myid",
	}

	detail.Detail = "invalid input"

	if detail.Error() != http.StatusText(http.StatusBadRequest) {
		t.Fatal()
	}

	if detail.StatusCode() != http.StatusBadRequest {
		t.Fatal()
	}

	t.Run("JSON", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", nil)
		r.Header.Set("Accept", "application/json")
		Serve(w, r, detail)
		b := w.Body.String()
		if b != `{"detail":"invalid input","status":400,"title":"Bad Request","type":"about:blank","id":"myid"}`+"\n" {
			t.Fatal(b)
		}
	})

	t.Run("XML", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", nil)
		r.Header.Set("Accept", "application/xml")
		Serve(w, r, detail)
		b := w.Body.String()
		if b != xml.Header+`<problem xmlns="urn:ietf:rfc:7807"><detail>invalid input</detail><status>400</status><title>Bad Request</title><type>about:blank</type><id>myid</id></problem>` {
			t.Fatal(b)
		}
	})
}

func TestError(t *testing.T) {
	t.Run("XML", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", nil)
		r.Header.Set("Accept", "application/xml")
		Serve(w, r, Wrap(http.StatusBadRequest, io.EOF))
		if w.Result().StatusCode != http.StatusBadRequest {
			t.Fatal()
		} else if w.Body.String() != xml.Header+`<problem xmlns="urn:ietf:rfc:7807"><detail>EOF</detail><status>400</status><title>Bad Request</title><type>about:blank</type></problem>` {
			t.Fatal(w.Body.String())
		}
	})

	t.Run("JSON", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", nil)
		r.Header.Set("Accept", "application/json")
		Serve(w, r, Wrap(http.StatusBadRequest, io.EOF))
		if w.Body.String() != `{"detail":"EOF","status":400,"title":"Bad Request","type":"about:blank"}`+"\n" {
			t.Fatal()
		}
	})

	t.Run("Text", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", nil)
		Serve(w, r, io.EOF)
		if w.Body.String() != "Internal Server Error\n" {
			t.Fatal()
		}
	})
}

func TestWrap(t *testing.T) {
	var err1 error = errors.New("permission denied")
	var err2 error = Wrap(http.StatusForbidden, err1)
	if errors.Unwrap(err2) != err1 {
		t.Fatal()
	}

	js, _ := json.Marshal(err2)
	if string(js) != `{"detail":"permission denied","status":403,"title":"Forbidden","type":"about:blank"}` {
		t.Fatal(string(js))
	}

	err3 := Wrapf(http.StatusBadRequest, "the %s could not be jiggered", "thing")
	if err3.(*Details).Detail != errors.Unwrap(err3).Error() {
		t.Fatal()
	}
}

type temporaryError struct{}

func (e temporaryError) Error() string   { return "" }
func (e temporaryError) Temporary() bool { return true }

func TestStatusCode(t *testing.T) {
	if StatusCode(nil) != 200 {
		t.Fatal()
	}

	if StatusCode(context.DeadlineExceeded) != 504 {
		t.Fatal()
	}

	if StatusCode(fmt.Errorf("error: %w", temporaryError{})) != 503 {
		t.Fatal()
	}

	if StatusCode(Wrapf(0, "error: %w", temporaryError{})) != 503 {
		t.Fatal()
	}

	if StatusCode(Wrap(400, errors.New("bla"))) != 400 {
		t.Fatal()
	}

	if StatusCode(errors.New("bla")) != 500 {
		t.Fatal()
	}
}
