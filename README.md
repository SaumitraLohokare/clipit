# clipit
Copy output of the previous command to clipboard.

## Build and Install
- Clone the repo
- `cd` to the folder
- Run `go install`
- Make sure you have your `$GOPATH/bin` in your `PATH`

## Usage

### Use with pipe

```bash
$ echo "Hello World" | clipit
```
This will copy "Hello World" to your clipboard. You can use this with any command. eg:
```bash
$ ls -al ~ | clipit
```
This will copy the output of the `ls` to your clipboard.

### Without pipe

```bash
$ clipit
Hello World
^D # Use ^D to signal end of input
```
This will copy "Hello World" to your clipboard.
