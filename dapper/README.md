NOTE: This was only used in the desktop version which is unsupported at the moment.

# dapper

Dapper is a funky little configuration management tool that exists to install you runtime dependencies of OpenOrch locally on your computer - mostly Docker and related dependencies like WSL on Windows.

There were some design decisions that shaped it:

- No dependencies (single binary)
- Be quick - never redo an already done job
- Stream things are they as happening for user feedback

You can see it in action when you click on the "Install Runtime" button on the Start screen of OpenOrch.

## How it works

Dapper runs [`Applications`](fixture/app.json) which are basically a collection of `Features`.

`Features` are a bunch of shell/powershell `Script`s for different operating systems.

These `Script`s have two main parts:
- a `Check` and
- an `Execute` script.

The flow goes like this:
- A `Check` script runs. If it succeeds, we move on to the next dependency in the list.
- If it fails, an `Execute` script is ran to fix the problem.
- If the `Execute` returns with an exit status other than `0` Dapper exits.
- If the `Execute` script returns with a success, we call the `Check` again to make sure.
- If the re`Check` fails, Dapper exists.
- If the re`Check` suceeds, we move on to the next dependency in the list.

The above describes the traversal of a Feature list, but Features also depend on each other recursively, forming a tree structure.
When a `Check` fails for a `Feature`, all dependencies get checked recursively.

## Run

### Example

```sh
go run main.go run --var-username=yourlogin fixture/app.json
```

```sh
PS C:\Users\dobro\mono\dapper> go run .\main.go run --var-username=dobro --var-assetfolder=$env:USERPROFILE .\fixture\app.json
```
