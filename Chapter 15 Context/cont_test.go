package main

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	/*
		t.Run("Happy Path Test", func(t *testing.T) {
			data := "hello, world"
			svr := Server(&SpyStore{response: data, t: t})

			request := httptest.NewRequest(http.MethodGet, "/", nil)
			response := httptest.NewRecorder()

			svr.ServeHTTP(response, request)

			if response.Body.String() != data {
				t.Errorf("got %s ,want %s ", response.Body.String(), data)
			}
		})
	*/
	t.Run("returns data from store", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}
		// *SpyResponseWriter implements http.ResponseWriter.because it has Header(),Write() and WriteHeader() methods.
		// Method names must be spelled correctly (or the same)

		svr.ServeHTTP(response, request)

		if response.written {
			t.Error("a response should NOT have been written")
		}

	})

}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}
