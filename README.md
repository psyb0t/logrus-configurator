# logrus-configurator 🤖

Welcome to `logrus-configurator`, the badass sidekick for your logging adventures with Go! You've stumbled upon the dark alley of loggers where things get configured so slick that even your professors can't help but nod in approval. 😎

## What's This Shit About? 💩

`logrus-configurator` is a Go package that whips your `logrus` logger into shape without you breaking a sweat. Think of it as that one plugin at a rave that just knows how to tune things up. Want to set the log level? Bam! 🎚️ Prefer JSON over plain text? Wham! 📄 Want to know who called the logger? Boom! 🔍 It's got you covered.

## Features

- No-nonsense log level setting (trace your bugs or go full-on panic mode, we don't judge).
- Formatting logs like a boss with JSON or text formats – keep it structured or keep it simple.
- Caller reporting for when you need to backtrack who messed up. It's like `CSI` for your code.
- Automated configuration using environment variables, because who has time for manual setup?

## Usage Example

Ready to rock with `logrus-configurator`? Check this out.

### main.go

```go
package main

import (
	_ "github.com/psyb0t/logrus-configurator"
	"github.com/sirupsen/logrus"
)

func main() {
	// Here's where the magic happens - just logging some stuff.
	logrus.Trace("this shit's a trace") // Ninja mode, won't show unless you want it to.
	logrus.Debug("this shit's a debug") // Debugging like a boss.
	logrus.Info("this shit's an info")  // Cool, calm, and collected info.
	logrus.Warn("this shit's a warn")   // Warning: badass logger at work.
	logrus.Error("this shit's an error") // Oh crap, something went sideways.
	logrus.Fatal("this shit's a fatal")  // Critical hit! It's super effective!
}
```

### Crank It Up

Get your environment dialed in like the soundboard at a goth concert:

```bash
export LOG_LEVEL="trace"   # Choose the verbosity level.
export LOG_FORMAT="text"   # Pick your poison: json or text.
export LOG_CALLER="true"   # Decide if you want to see who's calling the logs.
```

Unleash the beast with:

```bash
go run main.go
```

And let the good times roll with the output:

```plaintext
DEBU[0000]/github.com/psyb0t/logrus-configurator/log.go:28 github.com/psyb0t/logrus-configurator.config.log() logrus-configurator: level: trace, format: text, reportCaller: true
TRAC[0000]/github.com/psyb0t/logrus-configurator/.example/main.go:9 main.main() this shit's a trace
DEBU[0000]/github.com/psyb0t/logrus-configurator/.example/main.go:10 main.main() this shit's a debug
INFO[0000]/github.com/psyb0t/logrus-configurator/.example/main.go:11 main.main() this shit's an info
WARN[0000]/github.com/psyb0t/logrus-configurator/.example/main.go:12 main.main() this shit's a warn
ERRO[0000]/github.com/psyb0t/logrus-configurator/.example/main.go:13 main.main() this shit's an error
FATA[0000]/github.com/psyb0t/logrus-configurator/.example/main.go:14 main.main() this shit's a fatal
exit status 1
```

Wanna switch it up? Change the environment variables to mix the brew.

```bash
export LOG_LEVEL="warn"
export LOG_FORMAT="json"
export LOG_CALLER="false"
```

Then let it simmer with:

```bash
go run main.go
```

And enjoy the sweet sound of (almost) silence:

```plaintext
{"level":"warning","msg":"this shit's a warn","time":"2023-11-06T21:15:49+02:00"}
{"level":"error","msg":"this shit's an error","time":"2023-11-06T21:15:49+02:00"}
{"level":"fatal","msg":"this shit's a fatal","time":"2023-11-06T21:15:49+02:00"}
exit status 1
```

Whether you're in for a riot or a silent disco, `logrus-configurator` is your ticket. 🎟️ (check out all of the supported levels in [`level.go`](level.go))

And that's damn it. You've just pimped your logger!

## Contribute

Got an idea? Throw in a PR! Found a bug? Raise an issue! Let's make `logrus-configurator` as tight as your favorite jeans.

## License

It's MIT. Free as in 'do whatever the hell you want with it', just don't blame me if shit hits the fan.
