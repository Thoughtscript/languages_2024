class A {
    // https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Using_classes#private_fields
    #privatefieldA = 2 // # affix indicates field
    #privatefieldB = -1 // # affix indicates field

    // https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Using_classes#static_properties
    static staticfield = "I'm only accessible through Class A not an Instance of"

    // https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Using_classes#accessor_fields
    get privatefieldA() {
        return this.#privatefieldA // Getter can expose privatefield like usual
    }
}

class B extends A {
    publicfieldB = "I'm publicfieldB"

    constructor(constructorfield = "I'm initialized!"){
        super() // Must be present before this
        this.publicfieldA = "I'm publicfieldA"
        this.constructorfield = constructorfield
    }
}

class C extends B {
    // https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Using_classes#constructor
    constructor(){
        super("I'm initialized from the subclass!") // Must precede this 
    }

    static staticmethod() {
        return "I'm a static method!"
    }
}

const a = new A()
console.log(`Class A: ${A}`)
console.log(`Object a: ${JSON.stringify(a)}`)
console.log(A.staticfield)
//console.log(a.#privatefieldA) // error
console.log(a.privatefieldA) // is accessible through getter
//console.log(a.#privatefieldB) // error
console.log(a.privatefieldB === undefined) // no getter

const b = new B()
console.log(`Class B: ${B}`)
console.log(`Object b: ${JSON.stringify(b)}`)
console.log(b.publicfieldA)
console.log(b.publicfieldB)
console.log(b.constructorfield)
//console.log(b.#privatefieldA) // error
console.log(b.privatefieldB === undefined) // no getter

const c = new C()
console.log(`Class C: ${C}`)
console.log(`Object c: ${JSON.stringify(c)}`)
console.log(c.constructorfield)
console.log(C.staticmethod())
//console.log(c.#privatefieldA) // error