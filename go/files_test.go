package examples

import (
	"testing"
)

func TestFilesCreateText(t *testing.T) {
	_, err := FilesCreateText()
	if err != nil {
		t.Errorf("FilesCreateText returned an error: %v", err)
	}
}

func TestFilesCreateImage(t *testing.T) {
	_, err := FilesCreateImage()
	if err != nil {
		t.Errorf("FilesCreateImage returned an error: %v", err)
	}
}

func TestFilesCreateAudio(t *testing.T) {
	_, err := FilesCreateAudio()
	if err != nil {
		t.Errorf("FilesCreateAudio returned an error: %v", err)
	}
}

func TestFilesCreateVideo(t *testing.T) {
	_, err := FilesCreateVideo()
	if err != nil {
		t.Errorf("FilesCreateVideo returned an error: %v", err)
	}
}

func TestFilesCreatePdf(t *testing.T) {
	_, err := FilesCreatePdf()
	if err != nil {
		t.Errorf("FilesCreatePdf returned an error: %v", err)
	}
}

func TestFilesCreateFromIO(t *testing.T) {
	_, err := FilesCreateFromIO()
	if err != nil {
		t.Errorf("FilesCreateFromIO returned an error: %v", err)
	}
}

func TestFilesList(t *testing.T) {
	err := FilesList()
	if err != nil {
		t.Errorf("FilesList returned an error: %v", err)
	}
}

func TestFilesGet(t *testing.T) {
	_, err := FilesGet()
	if err != nil {
		t.Errorf("FilesGet returned an error: %v", err)
	}
}

func TestFilesDelete(t *testing.T) {
	err := FilesDelete()
	if err != nil {
		t.Errorf("FilesDelete returned an error: %v", err)
	}
}
