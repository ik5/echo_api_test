# Echo API Test

The following code is a test of mine to use the [echo](https://echo.labstack.com/) Golang web framework as a tool for writing
REST API.

## Reasons and goals

I wish to use Golang as a REST API system for high demand system.
After several frameworks that I have tested, I decided that to check both [Iris](https://www.iris-go.com/) and [echo](https://echo.labstack.com/).

While testing Iris, I encountered a few basic issues due to lack of proper documentation and examples that I could find.
Following the issues I have encountered, I have decided to work with Echo framework that is simpler for me to use without issues so far.

The test have few goals in mind (the order is irrelevant):

  * [-] API versioning.
    - [X] API version 1.
    - [ ] API version 2.
  * [X] Documentation of the API using [swagger](https://swagger.io/).
  * [ ] Documentation of the API using [openapi](https://www.openapis.org/)v3.
  * [X] Real-life like application (e.g. designing, structures etc).
  * [X] Reusable code as much as possible.
  * [X] Data validation using [validator](https://github.com/go-playground/validator).
  * [ ] Debugging of code using [DAP](https://microsoft.github.io/debug-adapter-protocol/)
  (I'm using [neovim](https://neovim.io/), but configuration is under `.vscode`).
  * [X] Using Golang's new [slog](https://pkg.go.dev/log/slog) logger on all libraries and tools.

# Build system

## General Information

The project was written and tested under Linux, and it is using [gnu-make](https://www.gnu.org/software/make/).

The `Makefile` has several entries, that some are called in chain for building the application.

At the moment the project testing the alignment of `struct` memory using [aligo](https://github.com/essentialkaos/aligo)
and other Linters arrive from [GolangCI linter](https://github.com/nametake/golangci-lint-langserver).

GolangCI has configuration file named `.golangci.yaml` with linters and rules what and how to check the written code.

While building the code, there are few steps that will be made:

  1. Removing old copy of the binary file (`api`) under the `bin/` directory.
  2. Installing dependencies configured under `go.mod`/`go.sum`.
  3. Executing Linters.
  4. Generating API documentation.
  5. Compiling and building the source code into a binary named `api` that located under `bin/` directory.

The Makefile system also add few details inside the source code when built (not touching the actual files):

  * Saving current `SHA1` commit version (of `HEAD`).
  * Saving built date and time.
  * Saving the current `git's` `brunch` that the built stood at.

The binary is statically built, so it will be easier to deploy, but it has some warnings, based on `libc` system
your system is using.

## How to use the `Makefile`

The `Makefile` has few entries that can be executed on their own:

  *  `linters` - Executing Linters one after the other.
  * `deps` - Executing `go mod tidy` on the project to install dependencies.
  * `clean-api` - Deleting the `bin/api` file (if exists, no error will be raised if it does not exists).
  * `generate-swagger` - Generating a new swagger documentation, and making sure that the schemas are in the proper
  name.
  * `install-deps` - Executing `go install` on the following dependencies:
    - `aligo` - `Struct` memory alignment `linter`.
    - `golangci-lint-langserver` - Language server of `golangci-lint`.
    - `golangci-lint` - The `linter` of `golangci-lint` to be executed by `CLI`.
    - `swag` - The Golang's implementation for static generating of swagger that the project is using.

## How to build the binary

There are two build entries for the code:

  * `build-api` - Building a development based binary named `api` under the `bin/` directory.
  * `build-released-api` - Building a release ready binary named `api` under the `bin/` directory.

### Example building

```bash
$ make build-api
```

### Example executing

```bash
$ bin/api
```

# Important

The following code was made only for me, and as such I do not take any responsibility on it.
I will fix my own bugs and I will add tests if/when suited for me.

If you are using this code and something happened to your machine, system or anything else,
you agreeing that it was your sole own responsibility to choose doing so, and you have no
blame, claims or anything else that takes or makes me responsible by any existed way.

You understand and accept that you are using the following code on your own risk.

There are no warranties of any kind.

**If you disagree with any or all that is written on this section -- _DO NOT USE OR EXECUTE THE CODE_**.

# License
Copyright 2023 ik_5

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  [http://www.apache.org/licenses/LICENSE-2.0](http://www.apache.org/licenses/LICENSE-2.0)

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
