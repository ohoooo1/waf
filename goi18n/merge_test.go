package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestMergeExecute(t *testing.T) {
	resetDir(t, "testdata/output")
	files := []string{
		"testdata/input/en-us.one.json",
		"testdata/input/en-us.two.json",
		"testdata/input/fr-fr.json",
		"testdata/input/ar-ar.one.json",
		"testdata/input/ar-ar.two.json",
	}

	mc := &mergeCommand{
		translationFiles:  files,
		sourceLanguageTag: "en-us",
		outdir:            "testdata/output",
		format:            "json",
	}
	if err := mc.execute(); err != nil {
		t.Fatal(err)
	}

	expectEqualFiles(t, "testdata/output/en-us.all.json", "testdata/expected/en-us.all.json")
	expectEqualFiles(t, "testdata/output/ar-ar.all.json", "testdata/expected/ar-ar.all.json")
	expectEqualFiles(t, "testdata/output/fr-fr.all.json", "testdata/expected/fr-fr.all.json")
	expectEqualFiles(t, "testdata/output/en-us.untranslated.json", "testdata/expected/en-us.untranslated.json")
	expectEqualFiles(t, "testdata/output/ar-ar.untranslated.json", "testdata/expected/ar-ar.untranslated.json")
	expectEqualFiles(t, "testdata/output/fr-fr.untranslated.json", "testdata/expected/fr-fr.untranslated.json")
}

func resetDir(t *testing.T, dir string) {
	if err := os.RemoveAll(dir); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(dir, 0777); err != nil {
		t.Fatal(err)
	}
}

func expectEqualFiles(t *testing.T, expectedName, actualName string) {
	actual, err := ioutil.ReadFile(actualName)
	if err != nil {
		t.Fatal(err)
	}
	expected, err := ioutil.ReadFile(expectedName)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(actual, expected) {
		t.Fatalf("contents of files did not match: %s, %s", expectedName, actualName)
	}
}
