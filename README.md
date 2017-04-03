goids
=====
[![Build Status](https://travis-ci.org/taybin/goids.svg?branch=master)](https://travis-ci.org/taybin/goids)
[![Coverage Status](https://coveralls.io/repos/github/taybin/goids/badge.svg?branch=master)](https://coveralls.io/github/taybin/goids?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/taybin/goids)](https://goreportcard.com/report/github.com/taybin/goids)

Boids in golang and javascript

Install
-------
```
yarn install
go get -d
go build
```

Run
---
```
./goids
```

Then open `http://localhost:4000`

Dev/Test
----
`yarn run buildw` will watch for changes to the javascript files and rebuild
`static/bundle.js` as appropriately.

`go test` or `ginkgo` will run the golang unittests
