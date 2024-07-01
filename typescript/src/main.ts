import { B } from './subclass'
import { C, D, E, F } from './destructingkeys'

try {

    const obj = new B()
    
    const X = obj.testB1()
    console.log(X)

    const Y = obj.testB2()
    console.log(Y)

    console.log(X === Y)
    console.log("===================")
    
    const c: C = {
        field: 'c1',
        '%field': 'c2',
    }

    console.log(c.field)

    const d: D = {
        string_literal: 'd1'
    }

    console.log(d.string_literal)
    // console.log(d['*field']) // Doesn't exist - cannot set a constant as key without destructuring.

    // Destructered into the keys of the interface/type.
    const e: E = {
        '!field': 'e1'
    }

    console.log(e['!field'])

    // Most flexibility - any key that's a string allowed.
    const f: F = {
        'fe': 'f1'
    }
    console.log(f.fe)


} catch (ex) {
    console.log(ex)
}