import { Sun } from "./sun";
import { Planet } from "./planet";

export interface StarSystem {
    name: string
    sun: Sun
    spawnLocation: { x: number, y: number, z: number },
    planets: Planet[]
}