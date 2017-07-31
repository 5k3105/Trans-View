package ass_test

import (
	"fmt"
	"testing"

	"github.com/KuroiKitsu/go-ass"
)

var expectedTextTokens = map[string][]string{
	`{\fad(485,1)}Episode 1  {\fs18}The Self-Defense Force Goes to Another World`: []string{
		`{\fad(485,1)}`,
		"Episode 1  ",
		`{\fs18}`,
		"The Self-Defense Force Goes to Another World",
	},
	"A Gothic lolita?!": []string{
		"A Gothic lolita?!",
	},
}

var expectedTextReadable = map[string]string{
	`{\fad(485,1)}Episode 1  {\fs18}The Self-Defense Force Goes to Another World`: "Episode 1  The Self-Defense Force Goes to Another World",
	`Episode 2  {\fs18}Two Military Forces`:                                       "Episode 2  Two Military Forces",
	`{\a6}Next time on GATE:`:                                                     "Next time on GATE:",
}

var expectedTextIsValid = map[string]bool{
	`Are you headed to the doujinshi \Nsale and exhibit, too?`: true,
	`Are you headed to the doujinshi \nsale and exhibit, too?`: true,
	"Are you headed to the doujinshi \nsale and exhibit, too?": false,

	`{\fad(1,250)}Watch Where\NYou Walk`:  true,
	`{{\fad(1,250)}Watch Where\NYou Walk`: false,
}

func TestText_Split(t *testing.T) {
	for rawText, expectedTokens := range expectedTextTokens {
		text := ass.Text(rawText)
		tokens := text.Split()

		if len(tokens) != len(expectedTokens) {
			t.Fatalf("expecting %d tokens (got %d)", len(expectedTokens), len(tokens))
		}

		for i, expectedToken := range expectedTokens {
			token := tokens[i]
			if token != expectedToken {
				t.Fatalf("expecting %#v (got %#v)", expectedToken, token)
			}
		}
	}
}

func TestText_Readable(t *testing.T) {
	for rawText, expectedReadable := range expectedTextReadable {
		text := ass.Text(rawText)
		readable := text.Readable()
		if readable != expectedReadable {
			t.Fatalf("invalid readable text %#v (expecting %#v)", readable, expectedReadable)
		}
	}
}

func ExampleText_Readable() {
	text := ass.Text(`{\i1}Onii-chan!{\i0}`)
	readable := text.Readable()

	fmt.Println(readable)
	// Output: Onii-chan!
}

func TestText_IsValid(t *testing.T) {
	for rawText, expectedIsValid := range expectedTextIsValid {
		text := ass.Text(rawText)
		isValid := text.IsValid()
		if isValid != expectedIsValid {
			if expectedIsValid {
				t.Fatalf("this text should be valid: %#v", rawText)
			} else {
				t.Fatalf("this text should be invalid: %#v", rawText)
			}
		}
	}
}
