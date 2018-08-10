<meta valign="center" halign="left">

### Developing Effective CLI Apps

#### In Golang, or Otherwise

10 August 2018

<meta footer="Ben Meier - Oracle Bristol Dojo Sessions - 10 August 2018">

---

### First a terrible (terrifying?) introduction..

Lets first establish what makes a *bad* command line interface.

You downloaded a "recommended" binary from some web page:

    $ ls -al demo
    -rwxr-xr-x  1 benmeier  staff   1.6M Jul 25 21:40 demo*

> "Solves the greatest programming problem known to man!" - Developer

---

### Run it!

    $ ./demo

    panic: runtime error: index out of range

    goroutine 1 [running]:
    main.main()
        /Users/benmeier/projects/go_workspace/src/github.com/AstromechZA/BuildingEffectiveCLIApps/demo/main.go:10 +0x23a

Not off to a good start..

At least it told you what the panic was!

--- 

### Help?

Nope.

    $ ./demo --help

    panic: runtime error: index out of range

    goroutine 1 [running]:
    main.main()
        /Users/benmeier/projects/go_workspace/src/github.com/AstromechZA/BuildingEffectiveCLIApps/demo/main.go:15 +0x1ef

---

### Man page?

No way.

    $ man ./demo
    No manual entry for ./demo

---

### So we try some arguments

    $ ./demo a
    panic: runtime error: index out of range

    goroutine 1 [running]:
    main.main()
        /Users/benmeier/projects/go_workspace/src/github.com/AstromechZA/BuildingEffectiveCLIApps/demo/main.go:15 +0x1ef

And finally get some output:

    $ ./demo a b
    Done.%
    $ echo $?
    0

Was that a good or bad thing?

Doesn't exactly _"Solve"_ anything!

---

### Need some help now..

- Ask the developer...

`.`

`.`

`.`

- Stackoverflow...

`.`

`.`

- Finally find an old comment on a cached version of a Google Groups page telling us to run it with an _integer and a string_.

---

### Aha!

    $ ./demo 10 world
    Hëllo  0	world
    Hëllo  1	world
    Hëllo  2	world
    Hëllo  3	world
    Hëllo  4	world
    Hëllo  5	world
    Hëllo  6	world
    Hëllo  7	world
    Hëllo  8	world
    Hëllo  9	world
    Done.%

---

<meta valign="center" halign="center" talign="center">

Now we can go and celebrate about how **awesome** free software is.

![Crying Gif](https://media.giphy.com/media/26ufcVAp3AiJJsrIs/giphy.gif)

If you have interfaces like this, you can only go up from here!

---

### Improvements

Lots of things that we can change here!

They generally fall in a few main areas:

1. Respect the users

2. Respect the environment

3. Transparency

We'll expand on these in the next slides and then implement some of them to improve the app.

---

## Respect the users

Have to balance out:

- How do you _expect_ average users, developers, noobs, and experts to use your app.

- How are users _actually going_ to use your app.

These are all difference categories of users and each have their own requirements and use cases that you need to anticipate.

---

### The average user

- Is probably only going to use 50-80% of the features and code paths that your CLI provides.

    - eg: How many of `kubectl`'s actions to do actually use in day to day?

- Will use it regularly enough to form semantic memories, but may still require prompting and help some days.

    - eg: How often do you look up the available subcommands for an object in the `docker` CLI?

---

### The noob

- HAS NO IDEA WHAT THEY ARE DOING

    - They *depend* on tutorials, manuals, or word of mouth

- May not even have a compatible system or environment ☹️

    - Does this tool work with Python 3?

- May be running it as root.. 

    - `rm -rf /etc`!

---

### The expert

- Knows what to expect whenever they run a command

- Is probably using even the most obscure flags and options

- Has their environment set up _just right_

---

### The developer

- Is probably running it from their development environment

- Is not scared off by bugs (probably just ignores them)

- > Works on my machine!

- Has probably never run it in production (in anger)

- Does not record their results

---

### The CI bot?

- Does not actually have a screen

- Can not respond to prompts

- Has unlimited patience or none at all

- Does not distinguish between failure cases

---

### Summary

- Try to use your CLI from the perspective of average users, noobs, experts, developers, and bots. 
  
- **Use these perspectives to drive usability improvements**.

---

## Respect the environment

It doesn't always run in the same context as it does on your machine.

Users have different

- operating systems
- shells
- platforms
- screen sizes
- hardware
    - clock speed
    - available memory
- character support
- ttys
- timezones

---

## Transparency

Don't hide information that would be better out in the open.

- Provide points of contact
    - Where is the documentation?
  
- Embed documentation
    - Explain concepts and common usecases
  
- Feedback/bug reports
    - Let real user experience drive your development
  
- Changelogs
    - What has changed?
  
- Hints  ("Did you mean X?")

---

<meta valign="center" halign="center" talign="center">

# Fixing things

Time for some opinionated changes 

\<rant>

---

### Preparing for development : Versioning

We can't track our improvements without any sense of versioning. 

This feeds directly into your **issue tracking** by tying bugs or feature requests to specific versions.

Types of versioning:

- Version strings (_what version is this meant to be_)
- Build dates (_when was this created_)
- Commit SHA (_what was the last change_)
- Commit date (_when was the last change_)
- Publicly available MD5/SHA checksums or GPG signatures (_is this official_)

---

### A version string

Use [SemVer](http://semver.org):

1. MAJOR version when you make incompatible API changes,
2. MINOR version when you add functionality in a backwards-compatible manner, and
3. PATCH version when you make backwards-compatible bug fixes.

You probably don't need MAJOR.MINOR.PATCH, start with just MAJOR.MINOR.

---

We could embed it in the CLI code base

```golang
const Version = "0.1"
```

Or get it from a Git branch or tag,

But lets prefer a version file for this demo

```
$ cat VERSION
0.1
```

With a build variable that we can inject

```golang
var Version = "unknown"
```

```
$ go build -ldflags "-X \"main.Version=$(cat VERSION)\""
```

---

Do the same with Git info

```golang
var CommitSHA = "unknown"
var CommitDate = "1970-01-01T00:00:00Z"
```

```
$ go build -ldflags "
    -X \"main.Version=$(cat VERSION)\"
    -X \"main.CommitSHA=$(git rev-parse --short HEAD)\"
    -X \"main.CommitDate=$(git show -s --format=%cI HEAD)\"
"
```

**Note:** Tools like [govvv](https://github.com/ahmetb/govvv) can do this for you, but it's easy enough to do in a `Makefile` or with bash.

---

Allow users to get at the version information

```golang
versionFlag := flag.Bool("--version", false, "Show the version information")

if versionFlag {
    fmt.Printf("Version: %s\n", Version)
    fmt.Printf("Git Commit: %s\n", CommitSHA)
    fmt.Printf("Git Date: %s\n", CommitDate)
    os.Exit(0)
}
```

```
$ ./demo --version
Version: 0.1
Git Commit: e1a06ec
Git Date: 2017-07-20T23:06:18+02:00
```

Now we can log bugs, reproduce issues, and generally be responsible developers.

---

## Use a proper args interface

- Don't parse `os.Args` yourself!

- *Always* check `len(os.Args)` if you are using positional arguments

Use an existing parser:

- Golang `flag` (often good enough)
- Python `argparse` (definitely good enough)
- [github.com/spf13/cobra](https://github.com/spf13/cobra)
- [github.com/urfave/cli](https://github.com/urfave/cli)
  
These give you more than just parsing, they also provide:

- `--help`
- "No such option --blah"
- Subcommands

---

### Developer gives us the new version (0.1)

```
$ ./demo
$ echo $?
0
```

At least it didn't stack trace this time..

```
$ ./demo --help
This binary solves the greatest programming problem known to man! It will print 'N' rows of hello text.

Usage:

$ demo (options...) [N] [text]

-version
        print the version information
```

```
$ ./demo --version
Version: 0.1
Git Commit: e1a06ec
Git Date: 2017-07-20T23:06:18+02:00
```

--- 

## Task 2 : Exit codes

**Very Important** when used in scripts, CI, recipes, automated environments

Some _traditional_ values, usually dictated by shells, but nothing official so just concentrate on `0` vs `1` - anything else probably won't be noticed unless its needed.

```golang
func mainInner() error {
    if rand.Float32() > 0.5 {
        return errors.New("Dice says no!")
    }
    return nil
}

func main() {
    if err := mainInner(); err != nil {
        fmt.Printf("Error: %s\n", err.Error())
        os.Exit(1)
    }
}
```

Exit codes are often not noticed in typical calls, but hit when using automation or shell pipelines (`pipefail`?).

---

## Task 3 : Error messages and language

If an error occurs *always* include a reason and *suggest* a fix if possible

    USELESS: "unexpected exception"
    MEH: "could not open file"
    USEFUL: "could not open file '/bob': file does not exist"

An error message that includes a suggestions means one less question from the user.

> "Could not write file '/something/else' due to a permission error. Make sure you have write permission on the file."

Be Kind.

---

## Task 4 : Exploit the different output streams

- The user (and shell) probably expects results to be written on `stdout`
  
- So you usually have `stderr` for your logs, errors, and messages

    - This is the default behaviour of standard library logging in Python and Golang.

- If you output machine readable content on `stdout`, don't muck it up by adding plain text log messages

---

## Task 5 : Allow for input streams

- It's great when a user can use the `stdin` stream

- Can add security in order to prevent sensitive data being written to disk

- `-` is often used as an argument to represent `stdin` (or `stdout`)

    - eg: `--input=- --output=-`

---

## A final version..

Versioning!

```
$ ./demo --version
Project: demo-cli (https://github.com/AstromechZA/BuildingEffectiveCLIApps)
Version: 0.2
Git Commit: e1a06ec
Git Date: 2017-07-20T23:06:18+02:00
```

Error messages and exit codes!

```
$ ./demo
Error: expected 2 arguments [N] [text]; see --help for info
$ echo $?
1
```

Argument type checking and parsing!

```
$ ./demo a world
Error: failed to parse argument 1 as a number: strconv.Atoi: parsing "a": invalid syntax
$ echo $?
1
```

---

<meta valign="center" halign="center" talign="center">

# Other CLI techniques, issues, and quirks

---

## stdin isn't always readable!

When running as a subprocess, `stdin` might not be attached to anything. Attempts to prompt the user may hang or fail.

```
import "golang.org/x/crypto/ssh/terminal"

...

if terminal.IsTerminal(os.Stdin.Fd()) {
    // prompt
    os.Stdin.Readline()
}
```

Another option is `Isatty()` from [github.com/andrew-d/go-termutil](https://github.com/andrew-d/go-termutil) 

---

## Environment variables

1. Name them in a way that relates to your app. You _may_ have to deal with conflicts otherwise..

    - `MY_APP_ENABLE_DEBUG=1` is more targetted than simply `DEBUG=1`

2. Always document these in `--help`

3. Tell the user when they are affecting the execution (if possible)

    > "Logging at debug level due to $MY_APP_ENABLE_DEBUG"

4. Often a `--debug` is a better choice as they are first class citizens in your CLI and its help menus

---

## Line rewriting and progress bars

Most terminals interpret control characters as more than just raw bytes.

- `\r` can reset the cursor to the line start
- `\b` can move the cursor one position to the left
- ANSI codes can move the cursor and page fairly arbitrarily.

Use these to continuously update a progress percentage, progress bar, or the entire screen.

**BUT** log files and piped output don't handle these well, so have options to disable or mute these.

---

## Detecting the terminal size

Editors like `vi` and `nano` expand to fill the full size of the terminal.

Most terminals will respond to various syscalls with information about their dimensions (although this is platform dependent).

Some shells will detect and export the dimensions as environment variables `$LINES`/`$COLS`. Some terminals will also forward a signal (`SIGWINCH`) to the process to indicate that the dimensions were changed.

Use this to adapt the width of tables, progress bars, and other things to the size of the users terminal.

_Remember to default to something sensible when not in a terminal!_

---

## Machine readable

In the spirit of the Linux philosophy, your app or tool should focus on doing _one specific thing_ really well.

This means it is likely to be used to compose more complicated workflows.

So design it to be machine readable and to operate with the standard line-by-line tools like `grep`/`sed`/`head`/`tail` where necessary.

```
Name        Age
---------------
Ted         39
Robin       37
Barney      41
```

vs `--json` 

```
{"name": "Ted", "age": 39}
{"name": "Robin", "age": 37}
{"name": "Barney", "age": 41}
```

---

## Subcommands

Subcommands can be convenient to keep things in one binary

```
$ my-server serve -p 80
$ my-server list-files
$ my-server prepare
```

But observe the problems that come from going too far in this direction

- `docker` cli has/had >50 subcommands each with different (and conflicting) options.

Sometimes separate binaries are better

```
$ my-server-serve -p 80
$ my-server-prepare
```

---

Noun-verb or verb-noun?

- `docker image ls blah` vs `kubectl get pod blah`?

Sometimes there are **too many verbs compared to nouns** and not enough commonality between the nouns. You can end up doing
crazy things like turning `cckube open nodeport` into `cckube create port-opening` just to use a common verb ("create").

Whatever you do, keep it consistent. Don't have `cli create balloon` and `cli balloon destroy`.

---

## Bash completion

Allow the shell to prompt the user with available subcommands or suitable files!

- This is _annoyingly_ platform and shell specific.

- Bash vs Zsh vs ?
- Linux vs MacOS

- If you're lucky, you can use a prebuilt library that hooks into your CLI processing and just handles it for you.

    - Python has `argcomplete` to complement `argparse`

---

## Using ANSI control codes

[ANSI](https://en.wikipedia.org/wiki/ANSI_escape_code) codes can be used to control:

- text color (16, 256, or RGB)
- text style (bold, underline)
- blinking
- window titles

Can work *very* well for user-centric tools and prompts.

But again, only work in tty's that will actual render them.

    ESC[38;5mmy important log lineESC[0m that you must look at

`--no-color` should disable colours..

And remember that terminals often have colour schemes that change the default colour maps.

Test for colour support by looking at `$TERM` or `$ tput colors`.

---

## `man` pages

- Very platform specific

- Don't attempt unless you know people will use it..

Better to implement a longer `help` or `manual` subcommand which delivers a large block of text straight into the 
`$PAGER` configured (pager is usually `less`).

---

## Config files

If you expect a config file, provide an example config in your `--help` or as an option:

    --generate-example-config   outputs an example json file on stdout and exits

This can help to keep your config example synchronised with the code that reads it.

    $ thing --generate-example-config > config.json
    $ thing --config ./config.json

It can also be handy to provide validation options as well. This would just validate the config and not actually do 
any work.

---

## Apply CI/CD practices

- You **can** unittest your CLI! (as long as you design it to be tested from the beginning)
    - Vital to test the way you _think_ your argument parsing works..
  
- Create functional and integration tests for your CLI
  
- Check your documentation examples (doctest?)
  
- Catch regressions before they are reported
  
- Automatically release new versions to github, repository
  
- Combined with versioning, this removes a lot of the manual effort involved in maintaining your tools

---

## A message about bash and/or Make

- People often ask what is wrong with writing and distributing long and complex bash scripts or make targets
    - ALL OF THE REASONS IN THESE SLIDES

- Yes you can implement some of the improvements, but really you're going to end up with an unmaintainable, and unsupported, mess.

- Start your bash scripts with `#!/usr/bin/env python` :trollface:

---

## So to recap..

### 1. Respect your users

- They _will_ need help
  
- Developers are also users
  
- So are machines
  
- Someone will try to run it via a Fish script through a Rails app and pass the output over HTTP in JSON to be manipulated in Javascript.
  
- Empower those who use it
  
- Provide remedy suggestions

---

### 2. Respect your environment

- Bash is not the only shell
  
- Either restrict compilation to one platform, or commit to being cross platform.
  
- Sometimes there isn't a TTY
  
- Multi-byte characters may not display correctly
  
- Control codes may not function correctly
  
- Colours may not display at all

### 3. Transparency

- Be contactable
  
- Make it easy to identify and reproduce bugs
  
- Explain your errors
