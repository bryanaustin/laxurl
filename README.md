# laxurl

More lax URL parsing. Examples:

| Parse()               | Scheme | Host            | Path |
| --------------------- | ------ | --------------- | ------- |
| net://example/simple  | net    | example         | /sample |
| example.com:443/about |        | example.com:443 | /about |
| :443                  |        | :443            | |
| example.com           |        | example.com     | |
| [fd::1]:53            |        | [fd::1]:53      | |

Try these in the url.Parse function for comparison
