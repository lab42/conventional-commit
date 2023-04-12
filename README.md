# Conventional commit docker action

This action validates the PR title and checks if it is formulated according to the conventional commit spec.

## Basic usage

```
name: Pull request

on:
  pull_request: {}

jobs:
  Validate:
    runs-on: ubuntu-latest
    steps:
      - name: Validate PR title
        uses: lab42/conventional-commit@main

```

## Configuring types and description validation. 

The following example uses regular expressions to valudate the types and description. Adjust them according to your needs.

* Default types      : build|chore|ci|docs|feat|fix|perf|refactor|revert|style|test
* Defautl description: ([\w ]+)

```
name: Pull request

on:
  pull_request: {}

jobs:
  Validate:
    runs-on: ubuntu-latest
    steps:
      - name: Validate PR title
        uses: lab42/conventional-commit@main
        with:
          # Allowed commit types
          allowed_types: |
              build
              chore
              ci
              docs
              feat
              fix
              perf
              refactor
              revert
              style
              test
          # Regular expression to validate description
          description_regexp: ([\w ]+)
          # Make scope requires (boolean)
          require_scope: "false"
          # Regular expression to validate scope. Only used when require_scope is true or when scope is provided.
          scope_regexp: (\w+)
```
