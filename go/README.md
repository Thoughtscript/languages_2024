# Go Review 2024

[![](https://img.shields.io/badge/go-1.22.4-lightblue.svg)](https://go.dev/doc/)

Several newish features were adding **Go** `1.18+`.

```bash
interfaces > {1 2 3}
interfaces > 6
interfaces > {1 2 3}
interfaces > &{3 4 5}
interfaces > 0xc000102000
interfaces > {3 4 5}
interfaces > 6
interfaces > 6
empty interface > {}
empty interface >
generics > exampleOne {s} {t}
generics > exampleTwo I'm first I'm last
generics > I'm x
generics > I'm y
struct inheritance > {{fido}}
struct inheritance > {{lizzy}}
type inheritance > AA
type inheritance > BB
pointers > 42
pointers > 21
pointers > 73
pointers > 73
pointers > 0xc000102008
pointers > <variable> 'num' has <value>: 100
pointers > 'numAddress' obtains <pointer> via & <address operator> and <pointer type>: *int = &num  0xc000100060
pointers > <derefences> back to the <value> with <dereferencing operator> on 'numAddress': derefNum = *numAddress 100
pointers > call & <address operator> on 'derefNum' to obtain <address>: &derefNum 0xc000100068
pointers > can <dereference> back directly with & and * operators sequentially: *&derefNum 100
pointers > set <value> on <dereference> of 'numAddress', then get <address> from <pointer type>: *numAddress = 42 0xc000100060
pointers > <dereference> back to <value>: *numAddress  42
reflection > reflection.AA reflection.BB reflection.CC
reflection > reflection.S
reflection > true true true
reflection > exampleFive {a} {c}
reflection > exampleFive {a} string will implement subtype!
reflection > exampleSix {a} string will implement subtype!
reflection > exampleSix {a} {b}
reflection > exampleFive {a} [1 2 3 4 5]
reflection > exampleSix {a} [1 2 3 4 5]
reflection > exampleFive {a} {}
reflection > exampleSix {a} {}
reflection > exampleSeven {a} {b}
reflection > exampleSeven {a} {c}
reflection > exampleEight {a} {c}
Recovered from panic! reflection > t doesn't implement W
```

## Advanced Topics

1. Limits of Type Parameterization
2. Limits of Weak Composition as a Stand-In for Inheritance
3. Limts of Implementation
4. Reflection
5. Better use of `panic()`, `defer`, and `recover()`

## Resources and Links

1. https://jesseduffield.com/Gos-Shortcomings-3/
2. https://gobyexample.com/interfaces
3. https://medium.com/@peterbi_91340/implement-true-inheritance-in-go-ff6243bfd7a8
4. https://medium.com/@wambuirebeka/advanced-golang-concepts-channels-context-and-interfaces-dc3b71cd0ed8
5. https://www.toptal.com/go/golang-oop-tutorial
6. https://go.dev/blog/defer-panic-and-recover