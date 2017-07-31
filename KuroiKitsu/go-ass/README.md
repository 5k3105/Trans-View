# go-ass

## Usagi

You can look at `ass_test.go` for an example.

Here's the copy-paste.

```go
package main

import (
	"fmt"

	"github.com/KuroiKitsu/go-ass"
)

func main() {
	subs, err := ass.ParseFile("sample.ass")
	if err != nil {
		panic(err)
	}

	if info := subs.Section("Script Info"); info != nil {
		title := info.Get("Title")
		fmt.Printf("Title: %s\n", title)

		width := info.Get("PlayResX")
		height := info.Get("PlayResY")
		fmt.Printf("Resolution: %sx%s\n", width, height)
	}

	if events := subs.Section("Events"); events != nil {
		for _, pair := range events.Pairs {
			if pair.Key != "Dialogue" {
				continue
			}

			rawText := pair.Get("Text")
			text := ass.Text(rawText).Readable()
			fmt.Printf("> %s\n", text)
		}
	}

	// Output:
	// Title: Default Aegisub file
	// Resolution: 1280x720
	// > Onii-chan!
	// > Come on, how long are you staying in bed?
	// > It's morning.
	// > Your breakfast will get cold if you don't get up.
	// > I even made it just for you, too.
	// > Why is her hair pink?
}
```

## License

Apache 2.0
