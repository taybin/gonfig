package unicon_test

import (
	"fmt"
	"os"

	. "github.com/taybin/unicon"
)

func ExampleHierarchy() {
	conf := NewConfig(nil)             // root config
	conf.Use("second", NewConfig(nil)) // config in hierarchy as second
	conf.Use("second").Set("asd", "abc")
	fmt.Println(conf.Get("asd"))
	// Output: abc
}

func ExampleDefaults() {
	conf := NewConfig(nil) // root config
	conf.ResetDefaults(map[string]interface{}{
		"test_default":   "123",
		"test_default_b": "321",
	})
	conf.Use("second", NewConfig(nil)) // config in hierarchy as second
	conf.Use("second").Set("test_default", "333")
	fmt.Println(conf.Get("test_default"), conf.Get("test_default_b"))
	// Output: 333 321
}

func ExampleSaveToJSON() {
	conf := NewConfig(nil)
	conf.Set("some", "variable")
	jsonconf := NewJSONConfig("./config.json", conf)
	jsonconf.Save()
	// OR:
	// jsonconf := NewJSONConf("./config")
	// jsonconf.Reset(conf.All());
	// jsonconf.Save()

	if err := jsonconf.Save(); err != nil {
		fmt.Println("Error saving config", err)
	}
	jsonconf2 := NewJSONConfig("./config.json")
	if err := jsonconf2.Load(); err != nil {
		fmt.Println("Error loading config", err)
	}
	fmt.Println(jsonconf2.Get("some"))
	os.Remove("./config.json")
	// Output: variable
}

func ExampleConstruction() {
	var cfg MemoryConfig
	cfg.Set("example1", "123")
	fmt.Println(cfg.Get("example1"))
	cfg2 := NewConfig(&cfg)
	fmt.Println(cfg2.All())
	// Output: 123
	// map[example1:123]
}
