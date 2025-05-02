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
export type E = Pick<D, "fieldA">;

export const A: A = {
    id: 1
}

export const C: C = {
    id: 1,
    name: "pronomen"
}

export const E: E = {
    fieldA: 'message'
}