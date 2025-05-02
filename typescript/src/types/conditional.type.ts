// https://www.typescriptlang.org/docs/handbook/2/conditional-types.html

type V = {
    id: number
}

export interface VV extends V {
    msg: string
}

// Types can be set conditionally now
// Canditional/Dynamic types CANNOT be used in Class definitions!
export type VVV = VV extends V ? VV : V;

export class VVVV implements VV {
    msg: string;
    id: number;  
      
    constructor(msg: string, id: number) {
        this.msg = msg;
        this.id = id;
    }
}