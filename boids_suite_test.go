package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBoids(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Boids Suite")
}
