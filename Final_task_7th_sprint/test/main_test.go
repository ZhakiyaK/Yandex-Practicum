package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCafeWhenOk(t *testing.T) {
	handler := http.HandlerFunc(mainHandle)

	requests := []string{
		"/cafe?count=2&city=moscow",
		"/cafe?city=tula",
		"/cafe?city=moscow&search=ложка",
	}

	for _, v := range requests {
		response := httptest.NewRecorder()
		req := httptest.NewRequest("GET", v, nil)
		handler.ServeHTTP(response, req)

		assert.Equal(t, http.StatusOK, response.Code)
	}
}

func TestCafeNegative(t *testing.T) {
	handler := http.HandlerFunc(mainHandle)

	requests := []struct {
		request string
		status  int
		message string
	}{
		{"/cafe", http.StatusBadRequest, "unknown city"},
		{"/cafe?city=omsk", http.StatusBadRequest, "unknown city"},
		{"/cafe?city=tula&count=na", http.StatusBadRequest, "incorrect count"},
	}
	for _, v := range requests {
		response := httptest.NewRecorder()
		req := httptest.NewRequest("GET", v.request, nil)
		handler.ServeHTTP(response, req)

		assert.Equal(t, v.status, response.Code)
		assert.Equal(t, v.message, strings.TrimSpace(response.Body.String()))
	}
}
func TestCafeCount(t *testing.T) {
	handler := http.HandlerFunc(mainHandle)

	for city := range cafeList {
		request := []struct {
			count int
			want  int
		}{
			{0, 0},
			{1, 1},
			{2, 2},
			{100, min(100, len(cafeList[city]))},
		}
		for _, v := range request {
			response := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/cafe?city="+city+"&count="+strconv.Itoa(v.count), nil)
			handler.ServeHTTP(response, req)
			require.Equal(t, http.StatusOK, response.Code)

			if response.Body.String() == "" {
				assert.Equal(t, v.want, 0)
			} else {
				assert.Equal(t, v.want, len(strings.Split(response.Body.String(), ",")))
			}

		}
	}
}
func TestCafeSearch(t *testing.T) {
	handler := http.HandlerFunc(mainHandle)

	city := "moscow"

	requests := []struct {
		search    string
		wantCount int
	}{
		{"фасоль", 0},
		{"кофе", 2},
		{"вилка", 1},
	}

	for _, v := range requests {
		url := fmt.Sprintf("/cafe?city=%s&search=%s", city, v.search)

		response := httptest.NewRecorder()
		req := httptest.NewRequest("GET", url, nil)
		handler.ServeHTTP(response, req)
		require.Equal(t, http.StatusOK, response.Code)

		for _, item := range strings.Split(response.Body.String(), ",") {
			if response.Body.String() != "" {
				assert.Contains(t, strings.ToUpper(item), strings.ToUpper(v.search))
			}
		}

	}
}
