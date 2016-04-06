// Package asset is a set of packages for calculating and evaluating financial asset figures.
//
// The broker package allows to get static broker data to calculate order prices and fees.
//
// The order package allows to calculate future price projections, broker fees and risk figures.
//
package asset

import (
	// broker package
	_ "github.com/MaximilianMeister/asset/broker"
	// order package
	_ "github.com/MaximilianMeister/asset/order"
)
