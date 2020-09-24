package unicon_test

import (
	"os"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/taybin/unicon"
)

var _ = Describe("EnvConfig", func() {
	var (
		err error
		cfg ReadableConfig
	)
	BeforeEach(func() {
		cfg = NewEnvConfig("")
		err = cfg.Load()
		Expect(err).ToNot(HaveOccurred())
	})
	It("Should load variables from environment", func() {
		Expect(len(cfg.All()) > 0).To(BeTrue())
		env := os.Environ()
		Expect(len(env) > 0).To(BeTrue())
		for _, kvpair := range env {
			pairs := strings.Split(kvpair, "=")
			Expect(len(pairs) >= 2).To(BeTrue())
			Expect(cfg.Get(strings.ToLower(pairs[0]))).To(Equal(pairs[1]))
		}
	})
	It("Should create namespaces if provided, split by _", func() {
		os.Setenv("POSTGRES_HOST", "localhost")
		cfg = NewEnvConfig("", "postgres")
		cfg.Load()
		Expect(cfg.Get("postgres.host")).To(Equal("localhost"))
	})
	It("Should create namespaces if provided in UPPERCASE", func() {
		os.Setenv("POSTGRES_HOST", "localhost")
		cfg = NewEnvConfig("", "POSTGRES")
		cfg.Load()
		Expect(cfg.Get("postgres.host")).To(Equal("localhost"))
	})
	It("Should not create namespaces if not provided", func() {
		os.Setenv("POSTGRES_HOST", "localhost")
		cfg = NewEnvConfig("")
		cfg.Load()
		Expect(cfg.Get("postgres_host")).To(Equal("localhost"))
	})
	It("Should create namespaces if provided, split by -", func() {
		os.Setenv("POSTGRES-HOST", "localhost")
		cfg = NewEnvConfig("", "postgres")
		cfg.Load()
		Expect(cfg.Get("postgres.host")).To(Equal("localhost"))
	})
	It("Should create namespaces if provided, split by :", func() {
		os.Setenv("POSTGRES:HOST", "localhost")
		cfg = NewEnvConfig("", "postgres")
		cfg.Load()
		Expect(cfg.Get("postgres.host")).To(Equal("localhost"))
	})
})
