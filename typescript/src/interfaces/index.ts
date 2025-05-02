import { E, Example } from './example.interface'

export const RUN = () => {
    console.log("=================== Interface Basics ===================")

    const EE = new Example('a', 3, true)

    console.log(E)

    console.log(EE)
    console.log(EE instanceof Example)

    console.log("=======================================================\n")
}