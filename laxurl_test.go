package laxurl

import (
	"net/url"
	"testing"
)

func TestParsing(t *testing.T) {
	t.Parallel()
	if d, err := Parse("tcp://some.server:1234/coolthings/mine?t=fb#thing"); !checkerror(t, err) {
		compare(t, &url.URL{
			Scheme:   "tcp",
			Host:     "some.server:1234",
			Path:     "/coolthings/mine",
			RawQuery: "t=fb",
			Fragment: "thing",
		}, d)
	}
}

func TestEmpty(t *testing.T) {
	t.Parallel()
	if d, err := Parse(""); !checkerror(t, err) {
		compare(t, &url.URL{}, d)
	}
}

func TestNoPort(t *testing.T) {
	t.Parallel()
	if d, err := Parse("tcp://some.server/coolthings/mine?t=fb#thing"); !checkerror(t, err) {
		compare(t, &url.URL{
			Scheme:   "tcp",
			Host:     "some.server",
			Path:     "/coolthings/mine",
			RawQuery: "t=fb",
			Fragment: "thing",
		}, d)
	}
}

func TestNoScheme(t *testing.T) {
	t.Parallel()
	if d, err := Parse("some.server:98765/coolthings/mine?t=fb#thing"); !checkerror(t, err) {
		compare(t, &url.URL{
			Host:     "some.server:98765",
			Path:     "/coolthings/mine",
			RawQuery: "t=fb",
			Fragment: "thing",
		}, d)
	}
}

func TestNoSchemeOrPort(t *testing.T) {
	t.Parallel()
	if d, err := Parse("some.server/coolthings/mine?t=fb#thing"); !checkerror(t, err) {
		compare(t, &url.URL{
			Host:     "some.server",
			Path:     "/coolthings/mine",
			RawQuery: "t=fb",
			Fragment: "thing",
		}, d)
	}
}

func TestHostOnly(t *testing.T) {
	t.Parallel()
	if d, err := Parse("router.internal"); !checkerror(t, err) {
		compare(t, &url.URL{
			Host: "router.internal",
		}, d)
	}
}

func TestHostPortOnly(t *testing.T) {
	t.Parallel()
	if d, err := Parse("localhost:8080"); !checkerror(t, err) {
		compare(t, &url.URL{
			Host: "localhost:8080",
		}, d)
	}
}

func TestPortOnly(t *testing.T) {
	t.Parallel()
	if d, err := Parse(":53"); !checkerror(t, err) {
		compare(t, &url.URL{
			Host: ":53",
		}, d)
	}
}

func TestPortPath(t *testing.T) {
	t.Parallel()
	if d, err := Parse(":2233/final/count"); !checkerror(t, err) {
		compare(t, &url.URL{
			Host: ":2233",
			Path: "/final/count",
		}, d)
	}
}

func TestHostPath(t *testing.T) {
	t.Parallel()
	if d, err := Parse("example.com/about"); !checkerror(t, err) {
		compare(t, &url.URL{
			Host: "example.com",
			Path: "/about",
		}, d)
	}
}

func TestHostPortPath(t *testing.T) {
	t.Parallel()
	if d, err := Parse("example.com:443/about"); !checkerror(t, err) {
		compare(t, &url.URL{
			Host: "example.com:443",
			Path: "/about",
		}, d)
	}
}

func TestSchemeHost(t *testing.T) {
	t.Parallel()
	if d, err := Parse("noms://nom.nom"); !checkerror(t, err) {
		compare(t, &url.URL{
			Scheme: "noms",
			Host:   "nom.nom",
		}, d)
	}
}

func TestSchemeHostPort(t *testing.T) {
	t.Parallel()
	if d, err := Parse("noms://nom.nom:999"); !checkerror(t, err) {
		compare(t, &url.URL{
			Scheme: "noms",
			Host:   "nom.nom:999",
		}, d)
	}
}

func TestSchemeOnly(t *testing.T) {
	t.Parallel()
	if d, err := Parse("words://"); !checkerror(t, err) {
		compare(t, &url.URL{
			Scheme: "words",
		}, d)
	}
}

func TestIpv4(t *testing.T) {
	t.Parallel()
	if d, err := Parse("10.20.30.40/admin"); !checkerror(t, err) {
		compare(t, &url.URL{
			Host: "10.20.30.40",
			Path: "/admin",
		}, d)
	}
}

func TestIpv6NoScheme(t *testing.T) {
	t.Parallel()
	if d, err := Parse("[fd00:cafe::beef:face:1982]/admin"); !checkerror(t, err) {
		compare(t, &url.URL{
			Host: "[fd00:cafe::beef:face:1982]",
			Path: "/admin",
		}, d)
	}
}

func TestIpv6WithScheme(t *testing.T) {
	t.Parallel()
	if d, err := Parse("udp://[fd00:cafe::beef:face:1982]/admin"); !checkerror(t, err) {
		compare(t, &url.URL{
			Scheme: "udp",
			Host:   "[fd00:cafe::beef:face:1982]",
			Path:   "/admin",
		}, d)
	}
}

func TestIpv6NoSchemeWithPort(t *testing.T) {
	t.Parallel()
	if d, err := Parse("[fd00:cafe::beef:face:1982]:80/admin"); !checkerror(t, err) {
		compare(t, &url.URL{
			Host: "[fd00:cafe::beef:face:1982]:80",
			Path: "/admin",
		}, d)
	}
}

func TestIpv6WithSchemeWithPort(t *testing.T) {
	t.Parallel()
	if d, err := Parse("udp://[fd00:cafe::beef:face:1982]:80/admin"); !checkerror(t, err) {
		compare(t, &url.URL{
			Scheme: "udp",
			Host:   "[fd00:cafe::beef:face:1982]:80",
			Path:   "/admin",
		}, d)
	}
}

func checkerror(t *testing.T, err error) bool {
	t.Helper()
	if err != nil {
		t.Errorf("Expected to not get an error, got %q", err)
	}
	return t.Failed()
}

func compare(t *testing.T, x, y *url.URL) {
	t.Helper()
	if x.Scheme != y.Scheme {
		t.Errorf("Expected scheme to be %q, it was %q", x.Scheme, y.Scheme)
	}
	if x.Host != y.Host {
		t.Errorf("Expected host to be %q, it was %q", x.Host, y.Host)
	}
	if x.Path != y.Path {
		t.Errorf("Expected path to be %q, it was %q", x.Path, y.Path)
	}
	if x.RawQuery != y.RawQuery {
		t.Errorf("Expected query string to be %q, it was %q", x.RawQuery, y.RawQuery)
	}
	if x.Fragment != y.Fragment {
		t.Errorf("Expected query string to be %q, it was %q", x.Fragment, y.Fragment)
	}
}
