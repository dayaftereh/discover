import { PACKAGE_JSON } from './package-json';
import { CONFIG } from './config';

export const SERVER_CONFIG: any = {
    NAME: PACKAGE_JSON.name,
    VERSION: PACKAGE_JSON.version,

    MAIN_PKG: "github.com/dayaftereh/discover/server/main",

    OUTPUT: CONFIG.OUTPUT,
}