import { Injectable, EventEmitter } from "@angular/core";
import { EngineStats } from "./engine-stats";
import { ObjectsInfo } from "./objects-info";

@Injectable()
export class GameOverlayService {

    onInfo: EventEmitter<boolean>

    onObjectsInfo: EventEmitter<ObjectsInfo>

    onEngineStats: EventEmitter<EngineStats>

    constructor() {
        this.onInfo = new EventEmitter<boolean>(true)
        this.onEngineStats = new EventEmitter<EngineStats>(true)
        this.onObjectsInfo = new EventEmitter<ObjectsInfo>(true)
    }

}