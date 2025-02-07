/*
More lax URL parsing. Examples:

Parse()               | Scheme | Host            | Path
---------------------------------------------------------
net://example/simple  | net    | example         | /sample
example.com:443/about |        | example.com:443 | /about
:443                  |        | :443            |
example.com           |        | example.com     |
[fd::1]:53            |        | [fd::1]:53      |
Try these in the url.Parse function for comparison
*/
package laxurl

import (
	"net/url"
	"strings"
)

var (
	// MagicString is a workaround to fight the URL parsing rules since I don't want to rewrite it from scratch.
	// It will vanish if you use it as the scheme or a host. That's why I made it changeable.
	// Change it if you invent a magicemphasisleader protocol... and then tell me what your new protocol does. Sounds interesting.
	MagicString = "magicemphasisleader"
)

// Parse is a more forgiving, more lax version of url.Parse generally for when you don't feel like passing a
// protocol or scheme. Or maybe not even a host. Or maybe no path. Or nothing at all.
func Parse(x string) (*url.URL, error) {
	if strings.HasPrefix(x, ":") {
		// For the case here you start with a port
		x = MagicString + x
	} else if strings.HasPrefix(x, "[") {
		// For the case where you have a ipv6 address with no scheme
		// It starts with an opening bracket [
		cb := strings.Index(x, "]")
		if cb > -1 {
			// has a closing bracket somewhere in the string ]
			sl := strings.Index(x, "/")
			if sl == -1 || sl > cb {
				/*
					Slash exists and it is after the closing bracket.
					Or the slash doesn't exist.
					Decision Table:
					cb | sl | expected
					 2 | -1 | true
					 2 |  3 | true
					 2 |  1 | false
				*/
				// This seems like the most solid check I could think of, apply the band-aid!
				x = MagicString + "://" + x
			}
		}
	}

	u, err := url.Parse(x)
	if err != nil {
		return nil, err
	}

	if len(u.Host) < 1 {
		// Host not parsed

		if len(u.Opaque) > 0 {
			// Uses opaque syntax
			ps := strings.Split(u.Opaque, "/")
			u.Host = u.Scheme + ":" + ps[0]
			u.Scheme = ""

			if len(ps) > 1 {
				u.Path = "/" + strings.Join(ps[1:], "/")
			}
		} else {
			// Probably means the first part of the path is the host
			ps := strings.Split(u.Path, "/")
			u.Host = ps[0]

			if len(ps) > 1 {
				u.Path = "/" + strings.Join(ps[1:], "/")
			} else {
				u.Path = ""
			}
		}
	}

	// The magic
	if u.Scheme == MagicString {
		u.Scheme = ""
	}
	if strings.HasPrefix(u.Host, MagicString) {
		u.Host = u.Host[len(MagicString):]
	}

	return u, nil
}

// Merge two URLs copy the base URL and override anything non-null
// with the extend URL.
func Merge(base, extend *url.URL) *url.URL {
	result := *base
	if len(extend.Scheme) > 0 {
		result.Scheme = extend.Scheme
	}
	if len(extend.Host) > 0 {
		result.Host = extend.Host
	}
	if len(extend.Path) > 0 {
		result.Path = extend.Path
	}
	if len(extend.RawQuery) > 0 {
		result.RawQuery = extend.RawQuery
	}
	if len(extend.Fragment) > 0 {
		result.Fragment = extend.Fragment
	}
	return &result
}
