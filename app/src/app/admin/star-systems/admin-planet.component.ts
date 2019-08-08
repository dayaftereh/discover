import { Component, Input, OnInit } from "@angular/core";
import { Planet } from "src/app/services/api/admin/planet";

@Component({
    selector: 'app-admin-planet',
    templateUrl: './admin-planet.component.html'
})
export class AdminPlanetComponent {

    @Input('planet')
    planet: Planet | undefined

    @Input('collapsed')
    collapsed: boolean

    constructor() {
        this.collapsed = false
    }

    infinity(x: number): number {
        if (!x) {
            return 0.0
        }
        if (x < (1E+300)) {
            return x
        }
        return Infinity
    }

    solarMass2KG(x: number): number {
        if (!x) {
            return 0.0
        }
        return this.infinity(x * 1.98847e30)
    }

    au2KM(x: number): number {
        if (!x) {
            return 0.0
        }
        // 1 AU = 149597870700 m
        return this.infinity(x * 149597870.700)
    }

    cmPerS2mPerS(x: number): number {
        if (!x) {
            return 0.0
        }
        return this.infinity(x * 0.01)
    }

    cmPerSSqr2mPerSSqr(x: number): number {
        if (!x) {
            return 0.0
        }
        return this.infinity(x * 0.01)
    }

    earthGravities2mPerSSqr(x: number): number {
        if (!x) {
            return 0.0
        }
        return this.infinity(x * 9.80665)
    }

    kelvin2Degrees(x: number): number {
        if (!x) {
            return 0.0
        }
        return this.infinity(x - 272.15)
    }
}