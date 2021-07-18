package httpsyproblem_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/askeladdk/httpsyproblem"
)

func ExampleWrap() {
	// Errors are associated with 500 Internal Server Error by default.
	err := io.EOF
	fmt.Println(httpsyproblem.StatusCode(err), err)

	// Wrap any error to associate it with a status code.
	err = httpsyproblem.Wrap(http.StatusBadRequest, err)
	fmt.Println(httpsyproblem.StatusCode(err), err)

	// Wrapping an already wrapped error changes the status code but preserves the details.
	err = httpsyproblem.Wrap(http.StatusNotFound, err)
	fmt.Println(httpsyproblem.StatusCode(err), err)
	fmt.Println(err.(*httpsyproblem.Details).Detail)
	// Output:
	// 500 EOF
	// 400 Bad Request
	// 404 Not Found
	// EOF
}

func ExampleError_json() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	err := httpsyproblem.Wrap(http.StatusBadRequest, io.EOF)
	err.(*httpsyproblem.Details).Instance = "/products/42"
	httpsyproblem.Error(w, r, err)

	fmt.Println(w.Result().Status)
	_, _ = io.Copy(os.Stdout, w.Body)
	// Output:
	// 400 Bad Request
	// {"detail":"EOF","instance":"/products/42","status":400,"title":"Bad Request","type":"about:blank"}
}

func ExampleError_text() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	httpsyproblem.Error(w, r, io.EOF)

	fmt.Println(w.Result().Status)
	_, _ = io.Copy(os.Stdout, w.Body)
	// Output:
	// 500 Internal Server Error
	// EOF
}

func ExampleStatusCode() {
	fmt.Println(httpsyproblem.StatusCode(nil))
	fmt.Println(httpsyproblem.StatusCode(io.EOF))
	fmt.Println(httpsyproblem.StatusCode(httpsyproblem.Wrap(http.StatusBadRequest, io.EOF)))
	// Output:
	// 200
	// 500
	// 400
}
