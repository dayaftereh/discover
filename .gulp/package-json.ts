export interface PackageJSON {
    version: string
    name: string
}

export const PACKAGE_JSON: PackageJSON = require("../package.json")