package infrastructures

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	config "github.com/spf13/viper"
)

// initConfig initializes the configuration
func initConfiguration() error {
	config.SetConfigName("App")
	config.AddConfigPath("configurations")
	if err := config.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

// SetConfiguration sets the configuration
func SetConfiguration(param string) error {
	if len(param) == 0 {
		// Get default configuration file
		if err := initConfiguration(); err != nil {
			return fmt.Errorf("%v", err)
		}
		return nil
	}

	// Get file extension
	ext := filepath.Ext(param)
	ext = strings.TrimPrefix(ext, ".")
	config.SetConfigType(ext)

	// Open configuration file
	file, err := os.Open(AbsolutePath(param))
	if err != nil {
		return err
	}
	defer file.Close()
	if err := config.ReadConfig(file); err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

// AbsolutePath get absolute path
func AbsolutePath(inPath string) string {
	if strings.HasPrefix(inPath, "$HOME") {
		inPath = userHomeDir() + inPath[5:]
	}

	if strings.HasPrefix(inPath, "$") {
		end := strings.Index(inPath, string(os.PathSeparator))
		inPath = os.Getenv(inPath[1:end]) + inPath[end:]
	}

	if filepath.IsAbs(inPath) {
		return filepath.Clean(inPath)
	}

	p, err := filepath.Abs(inPath)
	if err == nil {
		return filepath.Clean(p)
	}

	return ""
}
