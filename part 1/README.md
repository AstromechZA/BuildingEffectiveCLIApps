# Part 1 - Where do we even start..

Start by being able to track your progress and your bugs. You can't receive accurate reports or reproduce issues if you have no idea what version anyone is running or where it came from.

This directly feeds into whatever ticketing system you use. Jira, Github Issues, and most other issue trackers have methods of tying version information to each ticket. Use the graphing and report facilities to track improvements and changelogs over time. Motivate your team by showing them how much has been improved between version X and Y.

## Types of version information

Use one or more of these, they all have different meaning and different uses:

- Version strings (what version is this meant to be)
- Build dates (when was this created)
- Commit SHA (what was the last change, and what is the history)
- Commit date (when was the last change authored)
- Publicly available MD5 or SHA checksums (is this an official release)

## Semantic versioning strings

Versioning things by MAJOR.MINOR.PATCH. Increment the right numbers at the right times. See http://semver.org for more. To ease development, start with just 0.1 (MAJOR.MINOR) and bump it to 1.0.0 when you have an actual public release that matters.

Or... don't, and use a corporate style YEAR.MONTH.RELEASE like 17.7.1 or whatever suites your fancy. Don't not include a version string just because there's politics around a versioning scheme.

## Implementation in Golang

Storing a version string in your code is easy enough.. thats what constants are for right? But build dates and commit information is a bit harder. You can't store that in the code. The Golang linker has a `-X` option that can be used to set variables at compile time.

```
$ go build -ldflags "-X \"importpath.variable=some value\""
```

*(Demo time)*

## Using it in issue reports

Make it compulsory (or cultural) that all bug reports include a copy-paste or screenshot of the version information on the users machine. You should never have any doubt about what version exhibited the bug - and you should be able to reproduce it using the same version.
