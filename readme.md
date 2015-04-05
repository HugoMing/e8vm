[![BuildStatus](https://travis-ci.org/h8liu/e8vm.png?branch=master)](https://travis-ci.org/h8liu/e8vm)

# E8VM

Goal: A book written in working code and markdown document that
describes how computer system works, including architecture,
assemblers, compilers, and operating systems.

Read it [here](http://8k.lonnie.io).

Planned Features:

- **Modularized.** File based modules. No circular dependency (not only on packages,
  but also on files). A reader can always read the project a file by
  a file, either from bottom to top, or from top to bottom.
- **Small files.** Each file is shorter than 200 lines of code.
- **Tested and Documented.**
  Each file (will) come with test cases, examples, and markdown description.
- **Real.** The simulation (will) work like a real computer.

## TODOs

- `FuncDecl` and `VarDecl`  // 4.5
- Return, break, continue   // 4.6
- Basic types, char         // 4.7
- Type conversion           // 4.8
- Basic built-in panic      // 4.10
- Pointer                   // 4.12
- Array and slice           // 4.13
- String                    // 4.14
- Struct                    // 4.15
- Fields and methods        // 4.17
- Interface					
- Big number constants
- Symbol usage track

Some code debt:

- `g8` building file split: the current build* functions all reside in
  two files, `build_stmt.go` and `build_expr.go`. Need split them
  into smaller files.

And more...

- Improve code reading website
- Complete consts in asm8
- Clean up the symbol linking in asm8 a little bit
- Package building system that tracks timestamps
- Online filesystem and online editing
- Code formatter
