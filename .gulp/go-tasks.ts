import * as gulp from 'gulp';
import { promisify } from 'util';
import { exec } from 'child_process';
import { GO_CONFIG } from './go-config';
import * as fs from 'fs'

const execAsync = promisify(exec);

export namespace Go {

    export function defaultTasks(): void {

        gulp.task("go:install:dep", async () => {
            const files: string[] = [
                GO_CONFIG.GO_PKG,
                `${GO_CONFIG.GO_PKG}.exe`
            ]

            const exists: boolean = files.some((file:string)=>{
                return fs.existsSync(file)
            })
            
            if (exists) {
                return
            }

            await execAsync(`${GO_CONFIG.GO_GET} -u ${GO_CONFIG.GO_PKG_DEP}`)
        })

    }

}