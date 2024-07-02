export interface C {
    field: string;
    '%field': string;
}

const string_literal = '*field'

export interface D {
    string_literal: string;
}

const string_literal_key = '!field'

export interface E {
    [string_literal_key]: string;
}

export interface F {
    [x: string]: string;
}

