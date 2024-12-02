![GitHub last commit (main)](https://img.shields.io/github/last-commit/yyewolf/goaoc/main)

![Golang logo](./template/readme/logo.png)

## 🎄 Golang Advent Of Code Template 🎅

To setup this template you only need `Go` !

To begin, you only need to type the following :

```
$ ./setup.sh
```

## 🎄🎄 CLI Usage

| Command       | Description                                                                      |
| ------------- | -------------------------------------------------------------------------------- |
| `bench (b)`   | Run the benchmark for the current day.                                           |
| `completion`  | Generate autocompletion scripts.                                                 |
| `get (g)`     | Get the challenge for the current day / flags.                                   |
| `prepare (p)` | Prepare for the current day / flags.                                             |
| `run (r)`     | Run the solution or tests for the current day (using i for input or t for test). |
| `solve (s)`   | Run the input for the current day and submit.                                    |
| `stars (S)`   | Get stars for the current day.                                                   |
| `user (u)`    | Get your user ID.                                                                |
| `workflow`    | Run the workflow.                                                                |

### 🎄🎄🎄 Example :

Here's a simple example of what you might do (beginning at the root) :

```
$ goaoc p
// Creates 2023/day17 
$ goaoc g 
// Prints out first part of 2023/day17
// User: solve part 1
$ goaoc r t 
// Output seems to match expected for test input 
$ goaoc r i
// Output seems legit
$ goaoc s
// You solved it!
$ goaoc g -p2
// Prints out second part of 2023/day17
// User: solve part 2
// Same commands as before to run and solve!
// How good is my code ?
$ goaoc b
```

## 🎄🎄 Setup

If you want to setup the workflow, don't forget to put your **secrets** in the Github Actions **Secrets**.

You will also need to allow Github Actions to push on the repo.


## 🎄🎄 Customization

You can customize your AoC as you wish !

Here's a list of interesting files to modify:

- `README.template.md`
- `template/markdown/*`
- `template/go/*`

## Special thanks

- The benches are displayed thanks to a modified version of the now archived [prettybench](https://github.com/cespare/prettybench) from @cespare.
- The compact stars idea comes from [this repo](https://github.com/tselmek/advent-of-code) from @tselmek.
