package unicon_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/taybin/unicon"
)

var _ = Describe("JSONConfig", func() {
	var (
		err error
		cfg WritableConfig
	)

	BeforeEach(func() {
		cfg = NewJSONConfig("./config_valid.json")
		err = cfg.Load()
		Expect(err).ToNot(HaveOccurred())
	})

	Context("When the JSON config marshals properly", func() {
		It("Should have a string variable in config", func() {
			Expect(cfg.Get("test")).To(Equal("123"))
		})
		It("Should have an int variable in config", func() {
			Expect(cfg.GetInt("test_number")).To(Equal(1))
		})
		It("Should have a bool variable in config", func() {
			Expect(cfg.Get("test_bool")).To(Equal(true))
		})
		It("Should have a float variable in config", func() {
			Expect(cfg.Get("test_float")).To(Equal(12.34))
		})
		It("Should be case-insensitive", func() {
			Expect(cfg.GetBool("mixedcase")).To(BeTrue())
			Expect(cfg.GetBool("MIXEDCASE")).To(BeTrue())
		})
		It("Should not error", func() {
			Expect(err).NotTo(HaveOccurred())
		})
		It("Should have the string value from a nested map", func() {
			Expect(cfg.Get("test_object.nested_string")).To(Equal("abcd"))
		})
		It("Should have the string value from a deeply nested map", func() {
			Expect(cfg.Get("double_nested.nested_object.test_inner")).To(Equal("foo"))
		})
		It("Should have the int value from a nested map", func() {
			Expect(cfg.GetInt("test_object.nested_int")).To(Equal(987))
		})
		It("Should have case insensitive nesting", func() {
			Expect(cfg.Get("test_object.mixedcase")).To(BeTrue())
		})
		It("Should have the values from an input array", func() {
			Expect(cfg.GetInt("test_array[0].id")).To(Equal(1))
			Expect(cfg.GetString("test_array[1].id")).To(Equal("2"))
			Expect(cfg.Get("test_array[2].id")).To(Equal(3.0))
			Expect(cfg.Get("test_array.length")).To(Equal(3))
		})
	})

	Context("When config fails to marshal", func() {
		BeforeEach(func() {
			cfg = NewJSONConfig("./config_invalid.json")
			err = cfg.Load()
		})
		It("should return a functional config", func() {
			Expect(cfg).ToNot(BeZero())
			cfg.Set("QQ", "123")
			Expect(cfg.Get("QQ")).To(Equal("123"))
		})

		It("should error", func() {
			Expect(err).To(HaveOccurred())
		})
	})

	Context("When the JSON config does not exist", func() {
		BeforeEach(func() {
			cfg = NewJSONConfig("./config_nonexisting.json")
			err = cfg.Load()
		})
		It("should return a functional config", func() {
			Expect(cfg).ToNot(BeZero())
			cfg.Set("QQ", "123")
			Expect(cfg.Get("QQ")).To(Equal("123"))
		})
		It("should error", func() {
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("Config conversion", func() {
		It("Should be possible ro construct new JSON config from a gonfig hierarchy", func() {
			cfg := NewConfig(nil)
			cfg.Use("config_a", NewMemoryConfig())
			cfg.Use("config_b", NewMemoryConfig())
			cfg.Use("config_a").Set("config_a_var_a", "conf_a")
			cfg.Use("config_b").Set("config_b_var_a", "conf_b")
			jsonConf := &JSONConfig{
				Configurable: cfg,
				Path:         "./config.json",
			}
			err := jsonConf.Save()
			Expect(err).ToNot(HaveOccurred())
			jsonConf2 := NewJSONConfig("./config.json", cfg)
			err = jsonConf2.Save()
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
