// https://www.typescriptlang.org/docs/handbook/2/mapped-types.html
// Exact number of keys and their names aren't known or are flexible.
// But nevertheless constrained by the following schema:
export type Q = {
    [key: string]: boolean | number;
}

export interface O {
    keyA?: string, // optionally present
    keyB: number
}