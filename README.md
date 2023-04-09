# Conventional commit docker action

This action validates the PR title and checks if it is fomrmulated according to the spec.

## Inputs

No inputs are needed.

## Outputs

No ouputs are provided.

## Example usage

### Basic usage

```
name: Validate PR title
uses: lab42/conventional-commit
```

### Configuring types and description validation. 

The following example uses regular expressions to valudate the types and description. Adjust them according to your need as needed.

```
name: Validate PR title
uses: lab42/conventional-commit
env:
    TYPES: feat|fix # Default: build|chore|ci|docs|feat|fix|perf|refactor|revert|style|test
    DESCRIPTIOM: \[[\d{1,4}]\] ([\w ]+) # Defautl: ([\w ]+)
```
