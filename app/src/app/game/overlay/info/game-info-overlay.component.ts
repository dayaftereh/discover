import { Component, OnInit, OnDestroy } from "@angular/core";
import { Subscription } from "rxjs";
import { GameOverlayService } from "../service/game-overlay.service";
import { EngineStats } from "../service/engine-stats";
import { ObjectsInfo } from "../service/objects-info";

@Component({
    selector: 'app-game-info-overlay',
    templateUrl: './game-info-overlay.component.html'
})
export class GameInfoOverlayComponent implements OnInit, OnDestroy {

    visible: boolean

    engineStats: EngineStats

    objectsInfo: ObjectsInfo

    private subscriptions: Subscription[]

    constructor(private readonly gameOverlayService: GameOverlayService) {
        this.visible = false
        this.subscriptions = []
        this.objectsInfo = {}
        this.engineStats = {} as EngineStats
    }

    ngOnInit(): void {
        const visibleSubscription: Subscription = this.gameOverlayService.onInfo.subscribe((state: boolean) => {
            if (state) {
                this.toggleVisible()
            }
        })

        const engineStatsSubscription: Subscription = this.gameOverlayService.onEngineStats.subscribe((stats: EngineStats) => {
            this.mergeEngineStats(stats)
        })

        const objectsInfoSubscription: Subscription = this.gameOverlayService.onObjectsInfo.subscribe((objectsInfo: ObjectsInfo) => {
            this.mergeObjectsInfo(objectsInfo)
        })

        this.subscriptions.push(visibleSubscription, engineStatsSubscription, objectsInfoSubscription)
    }

    private toggleVisible(): void {
        this.visible = !this.visible
    }

    private mergeEngineStats(stats: EngineStats): void {
        this.engineStats = Object.assign(this.engineStats, stats)
    }

    private mergeObjectsInfo(objectsInfo: ObjectsInfo): void {
        this.objectsInfo = Object.assign(this.objectsInfo, objectsInfo)
    }

    toDegrees(radians: number | undefined): number {
        if (radians === undefined || radians === null) {
            return 0.0
        }
        return radians * (180.0 / Math.PI)
    }

    ngOnDestroy(): void {
        if (this.subscriptions) {
            this.subscriptions.forEach((subscription: Subscription) => {
                subscription.unsubscribe()
            })
        }
    }

}