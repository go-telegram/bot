package bot

import "testing"

func TestEscapeMarkdown(t *testing.T) {
	res := EscapeMarkdown("foo \\! _*[]()~`>#+-=|{}.! bar")
	if res != "foo \\\\! \\_\\*\\[\\]\\(\\)\\~\\`\\>\\#\\+\\-\\=\\|\\{\\}\\.\\! bar" {
		t.Fatalf("unexpected result %q", res)
	}
}

func TestEscapeMarkdownUnescaped(t *testing.T) {
	res := EscapeMarkdownUnescaped("foo \\! _*[]()~`>#+-=|{}.! bar")
	if res != "foo \\! \\_\\*\\[\\]\\(\\)\\~\\`\\>\\#\\+\\-\\=\\|\\{\\}\\.\\! bar" {
		t.Fatalf("unexpected result %q", res)
	}
}

func TestRandomString(t *testing.T) {
	res := map[string]struct{}{}
	for i := 0; i < 100; i++ {
		r := RandomString(10)
		if _, ok := res[r]; ok {
			t.Fatalf("value already exists")
		}
		res[r] = struct{}{}
	}

	for i := 10; i < 50; i++ {
		r := RandomString(i)
		if len(r) != i {
			t.Fatalf("unexpected len")
		}
	}
}
