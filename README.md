ushios/configo
==============
[![Build Status](https://travis-ci.org/ushios/configo.svg?branch=master)](https://travis-ci.org/ushios/configo)
[![Coverage Status](https://coveralls.io/repos/ushios/configo/badge.svg?branch=master&service=github)](https://coveralls.io/github/ushios/configo?branch=master)

Get config value each environments

Install
========

```
go get github.com/ushios/config
```

Usage
=====

`sample.cfg`

```sample.cfg
[DEFAULT]
host: github.local

[dev]
host: dev.github.com

[prd]
host: github.com
```

`sample.go`
```sample.go
config, err := Instance("dev")
fmt.Println(config.String("host")) // dev.github.com

```


License
=======

The source files are distributed under the [Mozilla Public License, version 2.0](http://mozilla.org/MPL/2.0/),
unless otherwise noted.  
Please read the [FAQ](http://www.mozilla.org/MPL/2.0/FAQ.html)
if you have further questions regarding the license.
