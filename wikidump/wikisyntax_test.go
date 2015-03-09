package wikidump

import (
    "testing"
)

func assertStringEq(t *testing.T, a, b string) {
    if a != b {
        t.Errorf("%q != %q", a, b)
    }
}

func TestCleanup(t *testing.T) {
    in := "|}Hello,<ref group=\"note\">1</rf> world{{math|bla{{?}}}}!{{bla"
    assertStringEq(t, cleanup(in), "Hello, world!")
}

func TestExtractLinks(t *testing.T) {
    links := make(chan *link)
    go extractLinks("before[[_target _page_ #\nsection|inside]]after", links)
    l := <-links
    assertStringEq(t, l.anchor, "beforeinsideafter")
    assertStringEq(t, l.target, "Target page")
}