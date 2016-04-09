[![GoDoc](https://godoc.org/github.com/MaximilianMeister/asset?status.svg)](https://godoc.org/github.com/MaximilianMeister/asset)
![Build Status](https://travis-ci.org/MaximilianMeister/asset.png?branch=master) ![Go Report Card](http://goreportcard.com/badge/MaximilianMeister/asset)

* [`broker`](https://gocover.io/github.com/MaximilianMeister/asset/order) ![Go Cover](https://gocover.io/_badge/github.com/MaximilianMeister/asset/broker)
* [`order`](https://gocover.io/github.com/MaximilianMeister/asset/order) ![Go Cover](https://gocover.io/_badge/github.com/MaximilianMeister/asset/order)

# asset

Package asset is a set of packages for calculating and evaluating financial asset figures.

The `broker` package allows to get static broker data to calculate order prices and fees.

The `order` package allows to calculate future price projections, broker fees and risk figures.

### features

* a risk reward ratio
* maximum gain
* maximum loss
* break even
* total commission rates of a specific broker
