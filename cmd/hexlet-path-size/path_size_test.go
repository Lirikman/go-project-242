package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSizeEmptyFolder(t *testing.T) {
	path := "/home/dmitriy/go-project-242/testdata/EmptyFolder"
	expected := int64(0)
	actual, err := GetSize(path, true)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestGetSizeAudioFile(t *testing.T) {
	path := "/home/dmitriy/go-project-242/testdata/audio_file.mp3"
	expected := int64(9327847)
	actual, err := GetSize(path, true)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestGetSizeOneFileInFolder(t *testing.T) {
	path := "/home/dmitriy/go-project-242/testdata/FilesFolder"
	expected := int64(27761)
	actual, err := GetSize(path, true)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestGetSizeTextFile(t *testing.T) {
	path := "/home/dmitriy/go-project-242/testdata/linux.txt"
	expected := int64(2095)
	actual, err := GetSize(path, true)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestGetSizeFilesAndFolders(t *testing.T) {
	path := "/home/dmitriy/go-project-242/testdata"
	expected := int64(9342567)
	actual, err := GetSize(path, true)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestGetSizeHiddenFile(t *testing.T) {
	path := "/home/dmitriy/go-project-242/testdata/HiddenFiles/.go1.2.txt"
	expected := int64(1943664)
	actual, err := GetSize(path, true)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestGetSizeFilesExceptHiddens(t *testing.T) {
	path := "/home/dmitriy/go-project-242/testdata/HiddenFiles"
	expected := int64(2687115)
	actual, err := GetSize(path, false)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestGetSizeAllFiles(t *testing.T) {
	path := "/home/dmitriy/go-project-242/testdata/HiddenFiles"
	expected := int64(4630779)
	actual, err := GetSize(path, true)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestFormatSizeZeroHumanTrue(t *testing.T) {
	size := float64(0)
	expected := "0.0MB"
	actual := FormatSize(size, true)
	assert.Equal(t, expected, actual)
}

func TestFormatSizeZeroHumanFalse(t *testing.T) {
	size := float64(0)
	expected := "0B"
	actual := FormatSize(size, false)
	assert.Equal(t, expected, actual)
}

func TestFormatSizeSmallHumanTrue(t *testing.T) {
	size := float64(1024)
	expected := "0.0MB"
	actual := FormatSize(size, true)
	assert.Equal(t, expected, actual)
}

func TestFormatSizeMediumHumanTrue(t *testing.T) {
	size := float64(1000050000)
	expected := "1000.0MB"
	actual := FormatSize(size, true)
	assert.Equal(t, expected, actual)
}

func TestFormatSizeMediumHumanFalse(t *testing.T) {
	size := float64(1000050000)
	expected := "1000050000B"
	actual := FormatSize(size, false)
	assert.Equal(t, expected, actual)
}
