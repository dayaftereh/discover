import { Gas } from "./gas";

export interface Planet {
    name: string
    type: string
    semiMajorAxis: number
    eccentricity: number
    axialTilt: number
    mass: number
    gasGiant: boolean
    dustMass: number
    gasMass: number
    coreRadius: number
    radius: number
    orbitZone: number
    density: number
    orbitPeriod: number
    day: number
    resonantPeriod: boolean
    escapeVelocity: number
    surfaceAcceleration: number
    surfaceGravity: number
    rootMeanSquareVelocity: number
    molecularWeight: number
    volatileGasInventory: number
    surfacePressure: number
    greenhouseEffect: boolean
    boilPoint: number
    albedo: number
    exosphericTemperature: number
    estimatedTemperature: number
    estimatedTerrestrialTemperature: number
    surfaceTemperature: number
    greenhouseRise: number
    highTemperature: number
    lowTemperature: number
    maxTemperature: number
    minTemperature: number
    hydrosphere: number
    cloudCover: number
    iceCover: number
    Atmosphere: Gas[]
    atmosphereType: string
    moons: Planet[]
}