package libtextcat

import (
    "testing"
    "io/ioutil"
)

const (
    TestConfigPath = "defaultcfg/conf.txt"
)

func testLanguage(t *testing.T, language string, filename string) {
    // Create textcat
    cat, err := NewTextCat(TestConfigPath)

    if nil != err {
        t.Fatalf("libtextcat wrapper init error: %s", err)
    }
    defer cat.Close()

    // Read test data file
    fbytes, err := ioutil.ReadFile(filename)

    if nil != err {
        t.Fatalf("Cannot read text file: %s", err)
    }

    res, err := cat.Classify(string(fbytes))

    if nil != err {
        t.Fatalf("Cannot classify: %s", err)
    }

    if len(res) == 0 ||
       res[0] != language {
        t.Errorf("Detected language mismatch for '%s': expected '%s', got '%s'", filename, language, res[0])
    }
}

func TestDefault(t *testing.T) {
    testLanguage(t, "danish",       "test/test_danish.txt")
    testLanguage(t, "english",      "test/test_english.txt")
    testLanguage(t, "finnish",      "test/test_finnish.txt")
    testLanguage(t, "french",       "test/test_french.txt")
    testLanguage(t, "german",       "test/test_german.txt")
    testLanguage(t, "hungarian",    "test/test_hungarian.txt")
    testLanguage(t, "italian",      "test/test_italian.txt")
    testLanguage(t, "norwegian",    "test/test_norwegian.txt")
    testLanguage(t, "portuguese",   "test/test_portuguese.txt")
    testLanguage(t, "romanian",     "test/test_romanian.txt")
    testLanguage(t, "russian",      "test/test_russian.txt")
    testLanguage(t, "spanish",      "test/test_spanish.txt")
    testLanguage(t, "swedish",      "test/test_swedish.txt")
    testLanguage(t, "turkish",      "test/test_turkish.txt")
}