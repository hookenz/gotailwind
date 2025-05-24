package downloader

import (
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

func EnsureTailwindExtraInstalled(version string) (string, error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}

	toolDir := filepath.Join(cacheDir, "gotailwind-extra", version)

	var binaryName string
	switch runtime.GOOS + "-" + runtime.GOARCH {
	case "darwin-arm64":
		binaryName = "tailwindcss-extra-macos-arm64"
	case "darwin-amd64":
		binaryName = "tailwindcss-extra-macos-x64"
	case "linux-arm64":
		binaryName = "tailwindcss-extra-linux-arm64"
	case "linux-amd64":
		binaryName = "tailwindcss-extra-linux-x64"
	case "windows-amd64":
		binaryName = "tailwindcss-extra-windows-x64.exe"
	default:
		return "", fmt.Errorf("unsupported platform: %s-%s", runtime.GOOS, runtime.GOARCH)
	}

	binPath := filepath.Join(toolDir, binaryName)
	if _, err := os.Stat(binPath); err == nil {
		return binPath, nil
	}

	// todo: no checksums provided by tailwind-cli-extra
	// sumURL := fmt.Sprintf("https://github.com/tailwindlabs/tailwindcss/releases/download/%s/sha256sums.txt", version)
	// expectedSum, err := getExpectedSHA256(sumURL, binaryName)
	// if err != nil {
	// 	return "", fmt.Errorf("failed to fetch expected hash: %w", err)
	// }

	if err := os.MkdirAll(toolDir, 0755); err != nil {
		return "", err
	}

	tmpFile := binPath + ".tmp"
	out, err := os.Create(tmpFile)
	if err != nil {
		return "", err
	}
	defer out.Close()

	binaryURL := fmt.Sprintf("https://github.com/dobicinaitis/tailwind-cli-extra/releases/download/%s/%s", version, binaryName)
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

	// todo: no checksums provided by tailwind-cli-extra
	// actualSum := hex.EncodeToString(hasher.Sum(nil))
	// if actualSum != expectedSum {
	// 	return "", fmt.Errorf("hash mismatch: expected %s, got %s", expectedSum, actualSum)
	// }
	out.Close() // close before Chmod and Rename (Windows doesn't allow renaming opened file).

	if err := os.Chmod(tmpFile, 0755); err != nil {
		return "", err
	}

	if err := os.Rename(tmpFile, binPath); err != nil {
		return "", err
	}

	return binPath, nil
}
