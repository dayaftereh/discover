import * as path from 'path'

const GO: string = "go"
const GO_PATH: string = process.env["GOPATH"]
const GO_BIN: string = path.join(GO_PATH, "bin")

export const GO_CONFIG: any = {
    GO,
    GO_BIN,
    GO_PATH,

    GO_GET: `${GO} get`,
    GO_RUN: `${GO} run`,
    GO_BUILD: `${GO} build`,
    GO_PKG: `${GO_BIN}/dep`,

    GO_PKG_DEP: "github.com/golang/dep/cmd/dep"
}
