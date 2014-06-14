package pagination

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-martini/martini"
)

func Test_NewPagination_Defaults(t *testing.T) {
	pagi := NewPagination("", "")
	if pagi.Limit != DefaultLimit {
		t.Error("Limit should be DefaultLimit")
	}
	if pagi.Page != DefaultPage {
		t.Error("Page should be DefaultPage")
	}
}

func Test_NewPagination_Underlimit(t *testing.T) {
	pagi := NewPagination("0", "-1")
	if pagi.Limit != DefaultLimit {
		t.Error("Limit should be DefaultLimit")
	}
	if pagi.Page != DefaultPage {
		t.Error("Page should be DefaultPage")
	}
}

func Test_NewPagination_Overlimit(t *testing.T) {
	pagi := NewPagination("21", "18446744073709551616")
	if pagi.Limit != DefaultLimit {
		t.Error("Limit should be DefaultLimit")
	}
	if pagi.Page != DefaultPage {
		t.Error("Page should be DefaultPage")
	}
}

func Test_NewPagination_Min(t *testing.T) {
	pagi := NewPagination("1", "0")
	if pagi.Limit != uint64(1) {
		t.Error("Limit should be 1")
	}
	if pagi.Page != uint64(0) {
		t.Error("Page should be 0")
	}
}

func Test_NewPagination_Max(t *testing.T) {
	pagi := NewPagination("20", "18446744073709551615")
	if pagi.Limit != uint64(20) {
		t.Error("Limit should be 20")
	}
	if pagi.Page != uint64(18446744073709551615) {
		t.Error("Page should be 18446744073709551615")
	}
}

func Test_Pagination_Offset_0(t *testing.T) {
	pagi := &Pagination{Limit: 10, Page: 0}
	if pagi.Offset() != 0 {
		t.Error("Offset() should be 0")
	}
}

func Test_Pagination_Offset_45(t *testing.T) {
	pagi := &Pagination{Limit: 15, Page: 3}
	if pagi.Offset() != 45 {
		t.Error("Offset() should be 45")
	}
}

func Test_Handler_Defaults(t *testing.T) {
	m := martini.Classic()
	m.Get("/", Handler, func(pagi *Pagination) {
		if pagi.Limit != DefaultLimit {
			t.Error("Limit should be DefaultLimit")
		}
		if pagi.Page != DefaultPage {
			t.Error("Page should be DefaultPage")
		}
	})
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	m.ServeHTTP(rec, req)
}

func Test_Handler_Underlimit(t *testing.T) {
	m := martini.Classic()
	m.Get("/", Handler, func(pagi *Pagination) {
		if pagi.Limit != DefaultLimit {
			t.Error("Limit should be DefaultLimit")
		}
		if pagi.Page != DefaultPage {
			t.Error("Page should be DefaultPage")
		}
	})
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/?l=0&p=-1", nil)
	m.ServeHTTP(rec, req)
}

func Test_Handler_Overlimit(t *testing.T) {
	m := martini.Classic()
	m.Get("/", Handler, func(pagi *Pagination) {
		if pagi.Limit != DefaultLimit {
			t.Error("Limit should be DefaultLimit")
		}
		if pagi.Page != DefaultPage {
			t.Error("Page should be DefaultPage")
		}
	})
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/?l=21&p=18446744073709551616", nil)
	m.ServeHTTP(rec, req)
}

func Test_Handler_Min(t *testing.T) {
	m := martini.Classic()
	m.Get("/", Handler, func(pagi *Pagination) {
		if pagi.Limit != uint64(1) {
			t.Error("Limit should be 1")
		}
		if pagi.Page != uint64(0) {
			t.Error("Page should be 0")
		}
	})
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/?l=1&p=0", nil)
	m.ServeHTTP(rec, req)
}

func Test_Handler_Max(t *testing.T) {
	m := martini.Classic()
	m.Get("/", Handler, func(pagi *Pagination) {
		if pagi.Limit != uint64(20) {
			t.Error("Limit should be 20")
		}
		if pagi.Page != uint64(18446744073709551615) {
			t.Error("Page should be 18446744073709551615")
		}
	})
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/?l=20&p=18446744073709551615", nil)
	m.ServeHTTP(rec, req)
}
