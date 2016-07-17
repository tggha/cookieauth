# cookieauth

Cookie-based Basic-Authentication HTTP middleware for Go (golang). Prevents the need to keep entering basic-auth username and passwords over and over.

[![GoDoc](https://godoc.org/github.com/jpillora/cookieauth?status.svg)](https://godoc.org/github.com/jpillora/cookieauth)  [![CircleCI](https://circleci.com/gh/jpillora/cookieauth.svg?style=shield&circle-token=69ef9c6ac0d8cebcb354bb85c377eceff77bfb1b)](https://circleci.com/gh/jpillora/cookieauth)

### Usage

Get package:

``` sh
$ go get -v github.com/jpillora/cookieauth
```

Quick use:

``` go
handler := http.HandlerFunc(...)
protected := cookieauth.Wrap(handler, "foo", "bar")
http.ListenAndServe(":3000", protected)
```

Customized use:

``` go
handler := http.HandlerFunc(...)

ca := cookieauth.New()
ca.SetUserPass("foo", "bar")
ca.SetExpiry(2 * time.Hour)
ca.SetLogger(log.New(os.Stdout, "", log.LstdFlags))

protected := ca.Wrap(handler)

http.ListenAndServe(":3000", protected)
```

Successful logins are hashed using scrypt and stored in a cookie.

#### MIT License

Copyright © 2016 Jaime Pillora &lt;dev@jpillora.com&gt;

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
'Software'), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
