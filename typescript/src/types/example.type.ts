type A = {
    id: number
}

type B = {
    name: string
}

// Types can be combined (intersect, union, simple polymorphism, etc.)
type C = A & B

interface D {
    fieldA: string,
    fieldB: string,
    fieldC: number,
}

// Can pick specific fields (think subtype)
type E = Pick<D, "fieldA">;

const A: A = {
    id: 1
}

const C: C = {
    id: 1,
    name: "pronomen"
}

const E: E = {
    fieldA: 'message'
}

console.log(A)

console.log(C)

console.log(E)