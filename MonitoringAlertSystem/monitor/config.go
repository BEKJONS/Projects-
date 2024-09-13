package monitor

import (
	"gopkg.in/yaml.v2" // YAML package to handle YAML files
	"io/ioutil"        // Provides functions to read files
)

// Config struct defines the structure of the configuration file
type Config struct {
	Thresholds struct {
		CPUUsage    int `yaml:"cpu_usage"`    // Threshold for CPU usage (in percentage)
		MemoryUsage int `yaml:"memory_usage"` // Threshold for memory usage (in percentage)
		DiskIO      int `yaml:"disk_io"`      // Threshold for Disk I/O (in MB/s)
	} `yaml:"thresholds"`
	Interval int `yaml:"interval"` // Monitoring interval in seconds
}

// LoadConfig function reads the configuration file and unmarshals the YAML content into the Config struct
func LoadConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path) // Read the config file from the given path
	if err != nil {
		return nil, err // Return an error if the file cannot be read
	}

	var config Config
	// Unmarshal the YAML data into the Config struct
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err // Return an error if the YAML is not correctly formatted
	}

	return &config, nil // Return the loaded configuration
}
