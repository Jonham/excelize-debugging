# excelize debugging

help to reproduce excelize lib issue case

## about VSCode multiple workspace

- https://github.com/golang/tools/blob/master/gopls/doc/workspace.md
- https://code.visualstudio.com/docs/editor/multi-root-workspaces

## Before commit or create PR

run `go test` and coverage test yourself

```shell

# run golang coverage test
function runGoCoverageReport() {
  outFile="cover.out"
  if [ -n "$1" ]; then
    outFile="$1"
    echo "using param as outFile: $1"
  fi

  # should all pass
  go test

  # generate report file
  go test -coverprofile=$outFile
  # open report HTML
  go tool cover -html=$outFile
}
```
