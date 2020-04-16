package config

import (
	"io/ioutil"
	"testing"
)

// mock file reading
type iouUtilTest struct {
	readCount int
}

func (io *iouUtilTest) ReadFile(name string) ([]byte, error) {
	io.readCount++
	return ioutil.ReadFile("./testConfig.toml")
}

// Test_LoadsConfig : Tests if config file is read correctly.
func Test_LoadsConfig(t *testing.T) {
	// reset read from previous tests
	once.Reset()
	originalIouUtil := iouUtil
	iouUtil = &iouUtilTest{}
	defer func() {
		iouUtil = originalIouUtil
	}()

	// act
	config := Configuration()

	// assert
	if config == nil {
		t.Errorf("Configuration file not loaded.")
	}
	if config.MongoURL != "testMongoURLvalue" {
		t.Errorf("MongoURL not read correctly.")
	}
}

// Test_Singleton : Tests if config file is read only once after multiple calls.
func Test_Singleton(t *testing.T) {
	// reset read from previous tests
	once.Reset()
	originalIouUtil := iouUtil
	testIouUtil := &iouUtilTest{}
	iouUtil = testIouUtil
	defer func() {
		iouUtil = originalIouUtil
	}()

	// act - read multiple times
	config := Configuration()
	config = Configuration()

	// assert
	if config == nil {
		t.Errorf("Configuration file not loaded.")
	}
	if testIouUtil.readCount != 1 {
		t.Errorf("Configuration file read multiple times. %d", testIouUtil.readCount)
	}
}
