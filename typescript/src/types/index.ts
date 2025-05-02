import {A, C, E} from './example.type'
import {Example} from './class.implements.type'
import {O, Q} from './optional.keys.and.mapped.type'
import {VVV, VVVV} from './conditional.type'

export const RUN = () => {
    console.log("=================== Types ===================")

    console.log(A)
    console.log(C)
    console.log(E)

    const X = new Example('fieldA', 0 , true);
    console.log(X)

    const V: VVV = {
        id: 0,
        msg: 'msg'
    }
    console.log(V)
    console.log(`typeof will still only check for JS inbuilt Types: ${typeof V}`) // Will still only check for inbuilt Type

    const W = new VVVV('msg', 1)
    console.log(`instanceof can still only check for TS Types indirectly through Classes implementing Types: ${W instanceof VVVV}`) // Can check for types only indirectly...
    // E.g. - by checking instance of a Class that implements a Type.

    const QQ: Q = {
        "keyA": true,
        "keyB": 1
    }
    console.log(`${JSON.stringify(QQ)} is valid`)
    
    const QQQ: Q = {
        "keyA": true,
        "keyB": 1,
    
        // Without the coerced Type/alias through the Type Assertion as
        // this would be invalid and wouldn't compile.
        "keyC": {
            "keyD": 1
        } as any
    }
    console.log(`${JSON.stringify(QQQ)} is valid`)

    const OO: O = {
        "keyB": 1
    }
    console.log(`${JSON.stringify(OO)} is valid`)

    console.log("=============================================\n")
}