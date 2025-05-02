import {E} from './example.type'

// Type is an alias for interface
export class Example implements E {
    fieldA = ""
    fieldB = 0
    fieldC = false

    constructor (fieldA: string, fieldB: number, fieldC: boolean) {
      this.fieldA = fieldA
      this.fieldB = fieldB
      this.fieldC = fieldC
    }
}