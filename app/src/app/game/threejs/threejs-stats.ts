import { EventEmitter } from "@angular/core";

export class ThreeJSStats {

    onChanged: EventEmitter<{ fps: number, update: number }>

    private update: number
    private updateStarted: number | undefined

    private fps: number
    private frames: number
    private framesTime: number | undefined
    private framesTimer: number

    constructor() {
        this.fps = 0;
        this.frames = 0
        this.update = 0.0
        this.framesTimer = 0
        this.onChanged = new EventEmitter<{ fps: number, update: number }>(true)
    }

    beginUpdate(): void {
        // increment fps
        this.frames++
        // set the update
        this.updateStarted = Date.now()
    }

    endUpdate(): void {
        if (!this.updateStarted) {
            throw new Error("beginUpdate() not called")
        }
        const now: number = Date.now()

        // calculate the update
        const delta: number = now - this.updateStarted
        // reset the update stratred
        this.updateStarted = undefined
        this.update = (this.update * 0.5 + delta * 0.5) //more stable

        // caluclate the fps
        if (this.framesTime) {
            // add to frames timer
            this.framesTimer += now - this.framesTime
            // check if timer above 1 second
            if (this.framesTimer > 1000.0) {
                // subtract one second
                this.framesTimer -= 1000
                // update the fps
                this.fps = (this.fps * 0.5 + this.frames * 0.5)//more stable
                // reset the frames
                this.frames = 0

                this.onChanged.emit({
                    fps: this.fps,
                    update: this.update
                })
            }
        }
        // reset the frames time
        this.framesTime = now
    }

}

