import { Component, Input, OnInit } from "@angular/core";
import { Planet } from "src/app/services/api/admin/planet";

@Component({
    selector: 'app-admin-planet',
    templateUrl: './admin-planet.component.html'
})
export class AdminPlanetComponent implements OnInit {

    @Input('planet')
    planet: Planet | undefined

    @Input('collapsed')
    collapsed: boolean

    atomSymbols: { [key: number]: { symbol: string, name: string } } = {}

    constructor() {
        this.collapsed = false
    }

    ngOnInit(): void {
        this.atomSymbols = {
            1: { symbol: "H", name: "Hydrogen" },
            2: { symbol: "He", name: "Helium" },
            7: { symbol: "N", name: "Nitrogen" },
            8: { symbol: "O", name: "Oxygen" },
            10: { symbol: "Ne", name: "Neon" },
            18: { symbol: "Ar", name: "Argon" },
            36: { symbol: "Kr", name: "Krypton" },
            54: { symbol: "Xe", name: "Xenon" },
            900: { symbol: "NH3", name: "Ammonia" },
            901: { symbol: "H2O", name: "Carbon Dioxide" },
            902: { symbol: "CO2", name: "Xenon" },
            903: { symbol: "O3", name: "Ozone" },
            904: { symbol: "CH4", name: "Methane" },
        }
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