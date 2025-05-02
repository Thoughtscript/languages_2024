import { A } from './superclass'
import { B } from './subclass'

export const RUN = () => {
    console.log("=================== Subclassing ===================")

    const obj = new B()
    const X = obj.testB1()
    console.log(X)
    console.log(`obj is an instance of B: ${obj instanceof B}`)
    console.log(`obj is an instance of A: ${obj instanceof A}`)

    const Y = obj.testB2()
    console.log(Y)
    console.log(X === Y)

    console.log("==================================================\n")
}