# `go-ircformat`

This package contains IRC text formatting and coloring helpers.

## Installation

```
go get -u github.com/n7st/go-ircformat
```

## Features

### Text coloring

There are specific helpers for foreground text colors:

```go
import ircColor "github.com/n7st/go-ircformat/color"

func main() {
    redText := ircColor.Red("some text")

    // ...
}
```

There is a helper for applying both foreground and background coloring:

```go
import ircColor "github.com/n7st/go-ircformat/color"

func main() {
    redTextWhiteYellowBackground := ircColor.Color("my text", &ircColor.Colors{
        Background: ircColor.New(ircColor.ColorYellow),
        Foreground: ircColor.New(ircColor.ColorRed),
    })

    // ...
}
```

### Emphasis

There are helpers for bold, italic, underline, strikethrough and monospace:

```go
import ircEmphasis "github.com/n7st/go-ircformat/emphasis"

func main() {
    boldText := ircEmphasis.Bold("my text")

    italicText := ircEmphasis.Italic("my text")

    underlineText := ircEmphasis.Underline("my text")

    strikethroughText := ircEmphasis.Strikethrough("my text")

    monospaceText := ircEmphasis.Monospace("my text")

    // ...
}
```

### Combining colors and emphasis

```go
import (
    ircColor    "github.com/n7st/go-ircformat/emphasis"
    ircEmphasis "github.com/n7st/go-ircformat/color"
)

func main() {
    boldRedText := ircEmphasis.Bold(ircColor.Red("my text"))

    boldRedYellowText := ircEmphasis.Bold(ircColor.Color("my text", &ircColor.Colors{
        Background: ircColor.New(ircColor.ColorYellow),
        Foreground: ircColor.New(ircColor.ColorRed),
    }))

    // ...
}
```
