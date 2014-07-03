package words

import (
    "testing"

    "bufio"
    "strings"
)

func assertStringSlicesEqual(t *testing.T, expected, actual []string) {
    for i, e := range expected {
        if e != actual[i] {
            t.Log(len(actual[i]))
            t.Errorf("'%s' should equal '%s'\n", actual[i], e)
        }
    }
}

func TestSimpleWords(t *testing.T) {
    input := bufio.NewReader(strings.NewReader("this is a test"))
    expected := []string{"this", "is", "a", "test"}
    actual := Tokenize(input)

    t.Log(actual)
    if len(expected) != len(actual) {
        t.Error("Slices are not the same length")
    }

    assertStringSlicesEqual(t, expected, actual)
}

func TestPeriod(t *testing.T) {
    input := bufio.NewReader(strings.NewReader("this is a test."))
    expected := []string{"this", "is", "a", "test"}
    actual := Tokenize(input)

    if len(expected) != len(actual) {
        t.Error("Slices are not the same length")
    }

    assertStringSlicesEqual(t, expected, actual)
}

func TestDoubleQuotes(t *testing.T) {
    input := bufio.NewReader(strings.NewReader("this \"is a test.\""))
    expected := []string{"this", "is", "a", "test"}
    actual := Tokenize(input)

    if len(expected) != len(actual) {
        t.Error("Slices are not the same length")
    }

    assertStringSlicesEqual(t, expected, actual)
}

func TestSingleQuotes(t *testing.T) {
    input := bufio.NewReader(strings.NewReader("this 'is a test.'"))
    expected := []string{"this", "is", "a", "test"}
    actual := Tokenize(input)

    if len(expected) != len(actual) {
        t.Error("Slices are not the same length")
    }

    assertStringSlicesEqual(t, expected, actual)
}

func TestPunctuation(t *testing.T) {
    input := bufio.NewReader(strings.NewReader("this 'is a test,' I said; I was thinking about bannanas! Really?"))
    expected := []string{"this", "is", "a", "test", "I", "said", "I", "was", "thinking", "about", "bannanas", "Really"}
    actual := Tokenize(input)

    if len(expected) != len(actual) {
        t.Error("Slices are not the same length")
    }

    assertStringSlicesEqual(t, expected, actual)
}

func TestWhitespace(t *testing.T) {
    input := bufio.NewReader(strings.NewReader("this is\ta\r\ntest"))
    expected := []string{"this", "is", "a", "test"}
    actual := Tokenize(input)

    if len(expected) != len(actual) {
        t.Error("Slices are not the same length")
    }

    assertStringSlicesEqual(t, expected, actual)
}

func TestParenthesis(t *testing.T) {
    input := bufio.NewReader(strings.NewReader("this is a test (really it is [really])."))
    expected := []string{"this", "is", "a", "test", "really", "it", "is", "really"}
    actual := Tokenize(input)

    if len(expected) != len(actual) {
        t.Error("Slices are not the same length")
    }

    assertStringSlicesEqual(t, expected, actual)
}

func TestDashes(t *testing.T) {
    input := bufio.NewReader(strings.NewReader("this is a test--really. I'm gonna have an em dash—told ya! Here comes an en dash–wow."))
    expected := []string{"this", "is", "a", "test", "really", "I'm", "gonna", "have", "an", "em", "dash", "told", "ya", "Here", "comes", "an", "en", "dash", "wow"}
    actual := Tokenize(input)

    if len(expected) != len(actual) {
        t.Error("Slices are not the same length")
    }

    assertStringSlicesEqual(t, expected, actual)
}

func TestURLs(t *testing.T) {
    input := bufio.NewReader(strings.NewReader("Go to http://www.google.com/ for info or bit.ly/123 but not bit .ly/123. Definitely https://www.pivotaltracker.com"))
    expected := []string{"Go", "to", "http://www.google.com/", "for", "info", "or", "bit.ly/123", "but", "not", "bit", "ly", "123", "Definitely", "https://www.pivotaltracker.com"}
    actual := Tokenize(input)

    if len(expected) != len(actual) {
        t.Error("Slices are not the same length")
    }

    assertStringSlicesEqual(t, expected, actual)
}

func TestLongURLs(t *testing.T) {
    input := bufio.NewReader(strings.NewReader("github.com/mceldeen/words google.com/#q=test"))
    expected := []string{"github.com/mceldeen/words", "google.com/#q=test"}
    actual := Tokenize(input)

    if len(expected) != len(actual) {
        t.Error("Slices are not the same length")
    }

    assertStringSlicesEqual(t, expected, actual)
}

func TestSlashes(t *testing.T) {
    input := bufio.NewReader(strings.NewReader("Talk to tom/jim for info"))
    expected := []string{"Talk", "to", "tom", "jim", "for", "info"}
    actual := Tokenize(input)

    if len(expected) != len(actual) {
        t.Error("Slices are not the same length")
    }

    assertStringSlicesEqual(t, expected, actual)
}
