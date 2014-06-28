package words

import (
    "bufio"
    "io"
    "strings"
)

func Tokenize(r *bufio.Reader) []string {
    tokens := make([]string, 0)

    var err error
    var word string
    for err != io.EOF {
        word, err = readTo(r, isSeperator)
        if err != nil && err != io.EOF {
            break
        }
        word = stripFromEnd(stripFromStart(word))
        tokens = append(tokens, word)
    }

    return tokens
}

func readTo(r *bufio.Reader, end func(rune, string) bool) (word string, err error) {
    var buf rune
    for {
        buf, _, err = r.ReadRune()
        if err != nil {
            return
        }
        if end(buf, word) {
            if len(word) > 0 {
                return
            } else {
                continue
            }
        }
        word = word + string(buf)
    }
}

func stripFromEnd(word string) string {
    for len(word) > 0 {
        last := rune(word[len(word)-1])
        switch {
        case punctuation[last] || quotes[last] || whitespace[last] || brackets[last]:
            word = word[:len(word)-1]
        default:
            return word
        }
    }
    return word
}

func stripFromStart(word string) string {
    for len(word) > 0 {
        first := rune(word[0])
        switch {
        case quotes[first] || brackets[first]:
            word = word[1:]
        default:
            return word
        }
    }
    return word
}

func isSeperator(b rune, word string) bool {
    if b == '/' {
        w := strings.ToLower(word)
        if w == "http:" || w == "http:/" {
            return false
        } else if i := strings.LastIndex(word, "."); i > 0 && !isSeperator(rune(word[i-1]), "") {
            return false
        }
        return true
    }
    if punctuation[b] {
        if len(word) > 0 {
            return isSeperator(rune(word[len(word)-1]), "")
        }
        return true
    }
    return whitespace[b] || dashes[b] || b == '/'
}

var punctuation = map[rune]bool{
    '.': true,
    ',': true,
    ';': true,
    '!': true,
    '?': true,
}

var quotes = map[rune]bool{
    '"':  true,
    '\'': true,
}

var whitespace = map[rune]bool{
    ' ':  true,
    '\t': true,
    '\r': true,
    '\n': true,
}

var brackets = map[rune]bool{
    '(': true,
    ')': true,
    '[': true,
    ']': true,
}

var dashes = map[rune]bool{
    '-': true,
    '—': true, // em dash
    '–': true, // en dash
}
