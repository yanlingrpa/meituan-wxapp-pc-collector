package main

import (
	"os"
	"path/filepath"
)

func main() {
}

// getGoModCache 获取 GOMODCACHE 环境变量或使用默认路径
func getGoModCache() string {
	if goModCache := os.Getenv("GOMODCACHE"); goModCache != "" {
		return goModCache
	}

	// 优先使用 GOPATH
	if goPath := os.Getenv("GOPATH"); goPath != "" {
		return filepath.Join(goPath, "pkg", "mod")
	}

	// 如果 GOPATH 也不存在，使用 ${HOME}/go/pkg/mod
	home, err := os.UserHomeDir()
	if err != nil {
		// 最后的默认值（如果连 home 都获取不了）
		return filepath.Join("go", "pkg", "mod")
	}

	return filepath.Join(home, "go", "pkg", "mod")
}
