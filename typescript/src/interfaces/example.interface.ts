// Interfaces are struct-like with typed constraints
export interface Example {
    fieldA: string
    fieldB: number
    fieldC: boolean
}

// variable types are specified via:
export const E: Example = {
    fieldA: "example",
    fieldB: 1,
    fieldC: true
}

// interfaces can be combined with class and constructor syntax
export class Example implements Example {
    fieldA = ""
    fieldB = 0
    fieldC = false

    constructor (fieldA: string, fieldB: number, fieldC: boolean) {
      this.fieldA = fieldA
      this.fieldB = fieldB
      this.fieldC = fieldC
    }
}