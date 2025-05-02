# Typescript 2025 Notes

[![](https://img.shields.io/badge/typescript-5.8.3-royalblue.svg)](https://www.typescriptlang.org/)

## Gotcha's and Review Items

1. **Windows** doesn't like `'` in **Babel** build commands: `babel src --out-dir built --extensions .ts,.tsx` per [package.json](./package.json) and [this](https://stackoverflow.com/questions/62615555/babel-creates-output-directory-but-doesnt-transpile-any-files) helpful note.
2. One can `export` both a `const` and a `type` with the same Variable Name within the same file per [example.type.ts](./src/types/example.type.ts).
3. **Types** are aliases for **Interfaces** and can be used in `class implements` [definitions](./src/types/class.implements.type.ts).
4. `types` can't share the same Name within the same Namespace.
5. Typescript still doesn't support Dynamic Importing that well - for example `require('my_plain_js_file.js)` could be loaded and will auto-excute the entire script without defining callable methods/functions in Node.

## Resources and Links

1. https://stackoverflow.com/questions/62615555/babel-creates-output-directory-but-doesnt-transpile-any-files