import * as path from 'path'

export const CWD: string = process.cwd()

export const CONFIG: any = {
    OUTPUT: path.join(CWD, "dist")
}