# Humane Units

Just a few functions for helping humanize times and sizes.

`goinstall` it as `github.com/dustin/go-humanize`, import it as
`"github.com/dustin/go-humanize"`, use it as `humanize`

## Sizes

This lets you take numbers like `82854982` and convert them to useful
strings like, `83MB` or `79MiB` (whichever you prefer).

Example:

    fmt.Printf("That file is %s.", humanize.Bytes(82854982))

## Times

This lets you take a `time.Time` and spit it out in relative terms.
For example, `12 seconds ago` or `3 days from now`.

Example:

    fmt.Printf("This was touched %s", humanize.Time(someTimeInstance))

Thanks to Kyle Lemons for the time implementation from an IRC
conversation one day.  It's pretty neat.
