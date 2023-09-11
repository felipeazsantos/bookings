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
	postedData := url.Values{}
	form := New(postedData)
	var has = form.Has("a")
	if has {
		t.Error("found the field when should not")
	}

	postedData = url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	has = form.Has("a")
	if !has {
		t.Error("The field must be found when was not")
	}
}

func TestForm_MinLength(t *testing.T) {

	postedData := url.Values{}
	postedData.Add("a", "abc")
	postedData.Add("b", "abcd")

	form := New(postedData)

	if form.MinLength("a", 4) {
		t.Error("The field doesn't have the min length required, but pass")
	}
	if !form.MinLength("b", 4) {
		t.Error("the field has the required min length")
	}
}