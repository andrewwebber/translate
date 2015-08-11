package translate

import "testing"

func TestTranslate(t *testing.T) {
	data := make(map[string]string)
	data["foo"] = "bar"
	r := NewSimpleRepro(data)
	value, err := r.Translate("foo", LANG_ENGLISH)
	if err != nil {
		t.Fatal(err)
	}

	if value != "bar" {
		t.Fatal("unexpected translation")
	}

	value, err = r.Translate("foo2", LANG_ENGLISH)
	if err != ERR_NOTFOUND {
		t.Fatal("expected not found err")
	}

	value, err = r.Translate("foo2", LANG_GERMAN)
	if err != ERR_NOTSUPPORTED {
		t.Fatal("expected not supported error")
	}
}
