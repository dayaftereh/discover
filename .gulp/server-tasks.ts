import * as gulp from 'gulp';
import { GO_CONFIG } from './go-config';
import * as shell from 'gulp-shell';
import * as path from "path"
import { promisify } from 'util';
import { exec, spawn, ChildProcess } from 'child_process';
import { SERVER_CONFIG } from './server-config';

const execAsync = promisify(exec);

export namespace Server {

    export async function build(os: string, arch: string, ldflags?: string[]) {
        if (!ldflags) {
            ldflags = []
        }

        const now: string = new Date().toString()
        ldflags.push('-X', `'main.VERSION=${SERVER_CONFIG.VERSION}'`, "-X", `'main.RELEASE=${now}'`)

        const command: string[] = [
            `${GO_CONFIG.GO_BUILD}`,
            "-buildmode=exe"
        ]

        const ldFlagsString: string = ldflags.join(" ")
        command.push("-ldflags", `"${ldFlagsString}"`)

        let name: string = `${SERVER_CONFIG.NAME}-${os}-${arch}`
        if (os === "windows") {
            name = `${name}.exe`
        }

        const output: string = path.join(SERVER_CONFIG.OUTPUT, name)

        command.push("-o", output, SERVER_CONFIG.MAIN_PKG)

        const cmd: string = command.join(" ")

        await execAsync(cmd, {
            env: {
                ...process.env,
                GOOS: os,
                GOARCH: arch
            }
        })
    }

    export async function run() {
        return new Promise((resolve, reject) => {
            const process: ChildProcess = spawn(
                GO_CONFIG.GO, [
                    "run",
                    SERVER_CONFIG.MAIN_PKG
                ], {
                    stdio: 'inherit'
                })

            process.on("error", (err: Error) => {
                reject(err)
            })

            process.on("exit", (code: number, signal: string) => {
                if (code !== 0) {
                    return reject(new Error(`go run exists with code [ ${code} ]`))
                }
                resolve()
            })
        })
    }

    export async function runStargen() {
        return new Promise((resolve, reject) => {
            const process: ChildProcess = spawn(
                GO_CONFIG.GO, [
                    "run",
                    SERVER_CONFIG.STARGEN_MAIN_PKG,
                    "-ldflags='-s -w'"
                ], {
                    stdio: 'inherit'
                })

            process.on("error", (err: Error) => {
                reject(err)
            })

            process.on("exit", (code: number, signal: string) => {
                if (code !== 0) {
                    return reject(new Error(`go run exists with code [ ${code} ]`))
                }
                resolve()
            })
        })
    }

    export function defaultTasks(): void {

        gulp.task("server:dependencies", gulp.series(
            ["go:install:dep"],
            shell.task(`${GO_CONFIG.GO_PKG} ensure`)
        ))

        gulp.task("server:build", gulp.series(["server:dependencies"], () => {
            return Server.build("windows", "amd64")
        }))

        gulp.task("server:run", gulp.series(["server:dependencies"], () => {
            return Server.run()
        }))

        gulp.task("server:stargen", gulp.series(["server:dependencies"], () => {
            return Server.runStargen()
        }))
    }

}
