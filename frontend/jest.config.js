module.exports = {
  collectCoverage: true,
  coverageReporters: ["clover", "text-summary"],
  preset: '@vue/cli-plugin-unit-jest/presets/typescript-and-babel',
  transform: {
    '^.+\\.vue$': 'vue-jest'
  },
  testResultsProcessor: "jest-sonar-reporter",
  "moduleNameMapper": {
    "\\.(css|jpg|png)$": "<rootDir>/mocks/empty-module.js"
  }
}
