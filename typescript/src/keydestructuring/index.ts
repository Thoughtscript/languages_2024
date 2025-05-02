import { C, D, E, F } from './destructingkeys'

export const RUN = () => {
    console.log("=================== Key Destructuring ===================")
    
    const c: C = {
        field: 'c1',
        '%field': 'c2'
    }
    console.log(`A String with Special Chars can be set directly as a Key: ${JSON.stringify(c)}`)

    const d: D = {
        string_literal: 'd1'
    }
    console.log(`String Literal can't be set as a Key without Destructuring: d['*field'] will fail`)
    //console.log(d['*field']) // Doesn't exist - cannot set a constant as key without destructuring.
    console.log(`Can still use Key without String marks but won't be bound by Variable Name: ${d.string_literal}`)

    // Destructered into the keys of the interface/type.
    const e: E = {
        '!field': 'e1'
    }
    console.log(`Special Char passed through String Literal into Destructured Key as e['!field']: ${e['!field']}`)

    // Most flexibility - any key that's a string allowed.
    const f: F = {
        'fe': 'f1'
    }
    console.log(`f.fe dot notation: ${f.fe}`)

    console.log("=========================================================\n")
}