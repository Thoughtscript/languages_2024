# Go Review 2024

[![](https://img.shields.io/badge/go-1.22.4-lightblue.svg)](https://go.dev/doc/)

Several newish features were adding **Go** `1.18+`.

```bash
2024-07-01 21:28:10 {1 2 3}
2024-07-01 21:28:10 6
2024-07-01 21:28:10 {1 2 3}
2024-07-01 21:28:10 &{3 4 5}
2024-07-01 21:28:10 0xc000048028
2024-07-01 21:28:10 {3 4 5}
2024-07-01 21:28:10 6
2024-07-01 21:28:10 6
2024-07-01 21:28:10 {}
2024-07-01 21:28:10 
2024-07-01 21:28:10 {s}
2024-07-01 21:28:10 {t}
2024-07-01 21:28:10 I'm first
2024-07-01 21:28:10 I'm last
2024-07-01 21:28:10 I'm x
2024-07-01 21:28:10 I'm y
2024-07-01 21:28:10 {{fido}}
2024-07-01 21:28:10 {{lizzy}}
2024-07-01 21:28:10 AA
2024-07-01 21:28:10 BB
2024-07-01 21:28:10 42
2024-07-01 21:28:10 21
2024-07-01 21:28:10 73
2024-07-01 21:28:10 73
2024-07-01 21:28:10 0xc000048030
2024-07-01 21:28:10 <variable> 'num' has <value>: 100
2024-07-01 21:28:10 'numAddress' obtains <pointer> via & <address operator> and <pointer type>: *int = &num  0xc0000120e0
2024-07-01 21:28:10 <derefences> back to the <value> with <dereferencing operator> on 'numAddress': derefNum = *numAddress 100
2024-07-01 21:28:10 call & <address operator> on 'derefNum' to obtain <address>: &derefNum 0xc0000120e8
2024-07-01 21:28:10 can <dereference> back directly with & and * operators sequentially: *&derefNum 100
2024-07-01 21:28:10 set <value> on <dereference> of 'numAddress', then get <address> from <pointer type>: *numAddress = 42 0xc0000120e0
2024-07-01 21:28:10 <dereference> back to <value>: *numAddress  42
```

## Resources and Links

1. https://jesseduffield.com/Gos-Shortcomings-3/
2. https://gobyexample.com/interfaces
3. https://medium.com/@peterbi_91340/implement-true-inheritance-in-go-ff6243bfd7a8
4. https://medium.com/@wambuirebeka/advanced-golang-concepts-channels-context-and-interfaces-dc3b71cd0ed8
5. https://www.toptal.com/go/golang-oop-tutorial