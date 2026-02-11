package main

import (
	"code"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPathSizeEmptyFolder(t *testing.T) {
	path := "./testdata/EmptyFolder"
	expected := "0B"
	actual, err := code.GetPathSize(path, false, true, true)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestGetPathSizeAudioFileNoHuman(t *testing.T) {
	path := "./testdata/audio_file.mp3"
	expected := "9327847B"
	actual, err := code.GetPathSize(path, false, true, true)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestGetPathSizeAudioFileHuman(t *testing.T) {
	path := "./testdata/audio_file.mp3"
	expected := "9.3MB"
	actual, err := code.GetPathSize(path, true, true, true)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestGetPathSizeOneFileInFolderNoHuman(t *testing.T) {
	path := "./testdata/FilesFolder"
	expected := "27761B"
	actual, err := code.GetPathSize(path, false, true, true)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestGetPathSizeTextFileHuman(t *testing.T) {
	path := "./testdata/linux.txt"
	expected := "2.1KB"
	actual, err := code.GetPathSize(path, true, true, true)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestGetPathSizeFilesAndFoldersNoHumanNoRecursive(t *testing.T) {
	path := "./testdata"
	expected := "9342567B"
	actual, err := code.GetPathSize(path, false, true, false)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
func TestGetPathSizeFilesAndFoldersNoHumanNoHiddenRecursive(t *testing.T) {
	path := "./testdata"
	expected := "12057443B"
	actual, err := code.GetPathSize(path, false, false, true)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestGetPathSizeFilesAndFoldersHumanNoHiddenRecursive(t *testing.T) {
	path := "./testdata"
	expected := "12.1MB"
	actual, err := code.GetPathSize(path, true, false, true)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestGetPathSizeFilesAndFoldersNoHumanHiddenRecursive(t *testing.T) {
	path := "./testdata"
	expected := "14001107B"
	actual, err := code.GetPathSize(path, false, true, true)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestGetPathSizeHiddenFileNoHuman(t *testing.T) {
	path := "./testdata/HiddenFiles/.go1.2.txt"
	expected := "1943664B"
	actual, err := code.GetPathSize(path, false, true, true)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestGetPathSizeFilesExceptHiddensNoHumanNoRecursive(t *testing.T) {
	path := "./testdata/HiddenFiles"
	expected := "2687115B"
	actual, err := code.GetPathSize(path, false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestGetPathSizeAllFilesNoHumanNoRecursive(t *testing.T) {
	path := "./testdata/HiddenFiles"
	expected := "4630779B"
	actual, err := code.GetPathSize(path, false, true, false)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestGetPathSizeAllFilesHumanNoRecursive(t *testing.T) {
	path := "./testdata/HiddenFiles"
	expected := "4.6MB"
	actual, err := code.GetPathSize(path, true, true, false)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestFormatSizeZeroHumanTrue(t *testing.T) {
	size := int64(0)
	expected := "0B"
	actual := code.FormatSize(size, true)
	assert.Equal(t, expected, actual)
}

func TestFormatSizeZeroHumanFalse(t *testing.T) {
	size := int64(0)
	expected := "0B"
	actual := code.FormatSize(size, false)
	assert.Equal(t, expected, actual)
}

func TestFormatSizeSmallHumanTrue(t *testing.T) {
	size := int64(1024)
	expected := "1.0KB"
	actual := code.FormatSize(size, true)
	assert.Equal(t, expected, actual)
}

func TestFormatSizeMediumHumanTrue(t *testing.T) {
	size := int64(1000050000)
	expected := "1.0GB"
	actual := code.FormatSize(size, true)
	assert.Equal(t, expected, actual)
}

func TestFormatSizeMediumHumanFalse(t *testing.T) {
	size := int64(1000050000)
	expected := "1000050000B"
	actual := code.FormatSize(size, false)
	assert.Equal(t, expected, actual)
}
