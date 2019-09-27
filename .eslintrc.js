module.exports = {
    parser:  '@typescript-eslint/parser',  // Specifies the ESLint parser
    plugins: ["@typescript-eslint", "prettier"],
    rules: {
      "prettier/prettier": ["error", { "singleQuote": true }]
    },
    extends:  [
      "eslint:recommended",
      "plugin:@typescript-eslint/recommended",
      "prettier/@typescript-eslint",
      "plugin:prettier/recommended"],
    parserOptions:  {
      ecmaVersion:  2018,  // Allows for the parsing of modern ECMAScript features
      sourceType:  'module',  // Allows for the use of imports
    },
    env: {
      node: true,
      mocha: true,
    }
  };