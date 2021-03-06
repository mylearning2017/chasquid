package normalize

import "testing"

func TestUser(t *testing.T) {
	valid := []struct{ user, norm string }{
		{"ÑAndÚ", "ñandú"},
		{"Pingüino", "pingüino"},
	}
	for _, c := range valid {
		nu, err := User(c.user)
		if nu != c.norm {
			t.Errorf("%q normalized to %q, expected %q", c.user, nu, c.norm)
		}
		if err != nil {
			t.Errorf("%q error: %v", c.user, err)
		}

	}

	invalid := []string{
		"á é", "a\te", "x ", "x\xa0y", "x\x85y", "x\vy", "x\fy", "x\ry",
		"henry\u2163", "\u265a", "\u00b9",
	}
	for _, u := range invalid {
		nu, err := User(u)
		if err == nil {
			t.Errorf("expected User(%+q) to fail, but did not", u)
		}
		if nu != u {
			t.Errorf("%+q failed norm, but returned %+q", u, nu)
		}
	}
}

func TestDomain(t *testing.T) {
	valid := []struct{ user, norm string }{
		{"ÑAndÚ", "ñandú"},
		{"Pingüino", "pingüino"},
		{"xn--aca-6ma", "ñaca"},
		{"xn--lca", "ñ"}, // Punycode is for 'Ñ'.
		{"e\u0301", "é"}, // Transform to NFC form.
	}
	for _, c := range valid {
		nu, err := Domain(c.user)
		if nu != c.norm {
			t.Errorf("%q normalized to %q, expected %q", c.user, nu, c.norm)
		}
		if err != nil {
			t.Errorf("%q error: %v", c.user, err)
		}

	}

	invalid := []string{"xn---", "xn--xyz-ñ"}
	for _, u := range invalid {
		nu, err := Domain(u)
		if err == nil {
			t.Errorf("expected Domain(%+q) to fail, but did not", u)
		}
		if nu != u {
			t.Errorf("%+q failed norm, but returned %+q", u, nu)
		}
	}
}

func TestAddr(t *testing.T) {
	valid := []struct{ user, norm string }{
		{"ÑAndÚ@pampa", "ñandú@pampa"},
		{"Pingüino@patagonia", "pingüino@patagonia"},
		{"pe\u0301@le\u0301a", "pé@léa"}, // Transform to NFC form.
	}
	for _, c := range valid {
		nu, err := Addr(c.user)
		if nu != c.norm {
			t.Errorf("%q normalized to %q, expected %q", c.user, nu, c.norm)
		}
		if err != nil {
			t.Errorf("%q error: %v", c.user, err)
		}

	}

	invalid := []string{
		"á é@i", "henry\u2163@throne",
	}
	for _, u := range invalid {
		nu, err := Addr(u)
		if err == nil {
			t.Errorf("expected Addr(%+q) to fail, but did not", u)
		}
		if nu != u {
			t.Errorf("%+q failed norm, but returned %+q", u, nu)
		}
	}
}
