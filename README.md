goids
=====
[![Build Status](https://travis-ci.org/taybin/goids.svg?branch=master)](https://travis-ci.org/taybin/goids)
[![Coverage Status](https://coveralls.io/repos/github/taybin/goids/badge.svg?branch=master)](https://coveralls.io/github/taybin/goids?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/taybin/goids)](https://goreportcard.com/report/github.com/taybin/goids)

Boids in golang and javascript.

Install
-------
```
yarn install
yarn build
go get -d
go build
```

Run
---
This creates a 200x200 space with no Z dimension.  I've found a max-velocity
of 5 seems to work well.  The computation should be O(n log n) with n being
number of boids.  Feel free to play with other counts, velocities, and dimensions.
```
./goids --count=100 --max-velocity=5 -- -100:100 -100:100 0:0
```

Then open `http://localhost:3000`

Dev/Test
--------
`yarn buildw` will watch for changes to the javascript files and rebuild
`static/bundle.js` as appropriately.

`go test` or `ginkgo` will run the golang unittests

Thank You
---------
* [Craig Reynolds](http://www.red3d.com/cwr/boids/)
* [Jack Perkins](https://github.com/jackaperkins/boids)
* [Conrad Parker](http://www.kfish.org/boids/pseudocode.html)
