package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields are missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r = httptest.NewRequest("POST", "/whatever", nil)
	r.PostForm = postedData
	form = New(r.PostForm)
	if !form.Valid() {
		t.Error("show does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	var has = form.Has("a", r)
	if has {
		t.Error("found the field when should not")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")

	r = httptest.NewRequest("POST", "/whatever", nil)
	r.PostForm = postedData
	form = New(r.PostForm)

	has = form.Has("a", r)
	if !has {
		t.Error("The field must be found when was not")
	}
}
