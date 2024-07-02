import { A } from './superclass'

export class B extends A {
    constructor(){
        super()
    }

    funcB() {
        
    }

    testB1() {
        console.log("testB1")
        return this.funcA()
    }

    testB2() {
        console.log("testB2")
        return super.funcA()
    }
}