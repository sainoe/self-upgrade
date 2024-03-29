package main

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestSignAndVerifiy(t *testing.T) {
	targetFile := "dist/self-upgrade.ver2"
	signatureFile := "dist/self-upgrade.ver2.RSAsignature"

	SignErr := SignRSA(targetFile, signatureFile)
	if SignErr != nil {
		t.Errorf("Signature failed: %s", SignErr)
	}
	verifyErr := VerifyRSASignature(targetFile, signatureFile)
	if verifyErr != nil {
		t.Errorf("Verification failed: %s", verifyErr)
	}
}

func TestDownloadFile(t *testing.T) {
	fmt.Println(checkNewVersion())
	DownloadErr := downloadAndVerifyFile(filepath.Join("dist", "self-upgrade.ver2"))
	if DownloadErr != nil {
		t.Errorf("Download failed failed: %s", DownloadErr)
	}
}
