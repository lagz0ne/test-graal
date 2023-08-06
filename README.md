## Graal ecosystem

Graal contains quite a few interesting concepts, especially interoptings with
Java ecocsystem.

Why Java is nice? Because it brings huge amount of libraries for PDF/printing to
earth, which is not available/solid as in other modern ecosystems

## Taskfile

Taskfile is used for build

- `npm install -g @go-task/cli`
- you can access taskfile using `task`

## First of all

- Make sure you have docker installed
- Run `task prepare`

## Experiments

### Static linked binaries

- `task native-exec` to build executable using graalvm
- Output then can be found and executed in `exec`

### Call Java from Graal NodeJS

- `task run:node`
- Details in `node-app/index.js``
