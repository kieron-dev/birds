package birdpedia_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBirdpedia(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Birdpedia Suite")
}
