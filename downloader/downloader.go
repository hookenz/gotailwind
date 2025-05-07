package downloader

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func EnsureTailwindInstalled(version string) (string, error) {
	cacheDir, _ := os.UserCacheDir()
	toolDir := filepath.Join(cacheDir, "gotailwind", version)

	var binaryName string
	switch runtime.GOOS + "-" + runtime.GOARCH {
	case "darwin-arm64":
		binaryName = "tailwindcss-macos-arm64"
	case "darwin-amd64":
		binaryName = "tailwindcss-macos-x64"
	case "linux-arm64":
		binaryName = "tailwindcss-linux-arm64"
	case "linux-amd64":
		binaryName = "tailwindcss-linux-x64"
	case "windows-amd64":
		binaryName = "tailwindcss-windows-x64.exe"
	default:
		return "", fmt.Errorf("unsupported platform: %s-%s", runtime.GOOS, runtime.GOARCH)
	}

	binPath := filepath.Join(toolDir, binaryName)
	if _, err := os.Stat(binPath); err == nil {
		return binPath, nil
	}

	sumURL := "https://github.com/tailwindlabs/tailwindcss/releases/latest/download/sha256sums.txt"
	expectedSum, err := getExpectedSHA256(sumURL, binaryName)
	if err != nil {
		return "", fmt.Errorf("failed to fetch expected hash: %w", err)
	}

	if err := os.MkdirAll(toolDir, 0755); err != nil {
		return "", err
	}

	tmpFile := binPath + ".tmp"
	out, err := os.Create(tmpFile)
	if err != nil {
		return "", err
	}
	defer out.Close()

	binaryURL := fmt.Sprintf("https://github.com/tailwindlabs/tailwindcss/releases/download/%s/%s", version, binaryName)
	resp, err := http.Get(binaryURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("download failed: %s", resp.Status)
	}

	hasher := sha256.New()
	progress := &progressWriter{
		total:       resp.ContentLength,
		lastPercent: -1,
	}
	multi := io.MultiWriter(out, hasher, progress)

	if _, err := io.Copy(multi, resp.Body); err != nil {
		return "", err
	}
	fmt.Println() // newline after progress bar

	actualSum := hex.EncodeToString(hasher.Sum(nil))
	if actualSum != expectedSum {
		return "", fmt.Errorf("hash mismatch: expected %s, got %s", expectedSum, actualSum)
	}

	if err := os.Chmod(tmpFile, 0755); err != nil {
		return "", err
	}
	if err := os.Rename(tmpFile, binPath); err != nil {
		return "", err
	}

	return binPath, nil
}

func getExpectedSHA256(sumURL, binaryName string) (string, error) {
	resp, err := http.Get(sumURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("failed to fetch checksums: %s", resp.Status)
	}

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 2 && strings.TrimPrefix(parts[1], "./") == binaryName {
			return parts[0], nil
		}
	}

	return "", fmt.Errorf("no checksum found for %s", binaryName)
}

type progressWriter struct {
	written     int64
	total       int64
	lastPercent int
}

func (p *progressWriter) Write(b []byte) (int, error) {
	n := len(b)
	p.written += int64(n)

	if p.total > 0 {
		percent := int((p.written * 100) / p.total)
		if percent != p.lastPercent {
			p.lastPercent = percent
			bar := strings.Repeat("#", percent/2) // 50-char bar
			fmt.Printf("\rDownloading: [%-50s] %3d%%", bar, percent)
		}
	}

	return n, nil
}
