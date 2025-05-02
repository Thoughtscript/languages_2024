import * as SUBCLASSES_EXAMPLE from './classes'
import * as KEY_DESTRUCTURING_EXAMPLE from './keydestructuring'
import * as BASIC_INTERFACE_EXAMPLE from './interfaces'
import * as BASIC_TYPES from './types'

try {

    SUBCLASSES_EXAMPLE.RUN()
    KEY_DESTRUCTURING_EXAMPLE.RUN()
    BASIC_INTERFACE_EXAMPLE.RUN()
    BASIC_TYPES.RUN()

} catch (ex) {
    console.log(ex)
}