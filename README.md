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
        env:
            GITHUB_TOKEN: ${{ secrets.TOKEN }}

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
        env:
            GITHUB_TOKEN: ${{ secrets.TOKEN }}
            TYPES: feat|fix
            DESCRIPTIOM: \[([\w])-(\d{1,4})\] ([\w ]+) # Example for requiring Jira issue number before description.
```
