import { Component, OnInit } from "@angular/core";
import { ActivatedRoute, ParamMap } from "@angular/router";
import { first } from 'rxjs/operators';
import { AdminService } from "src/app/services/api/admin/admin.service";
import { StarSystem } from "src/app/services/api/admin/star-system";
import { Planet } from "src/app/services/api/admin/planet";
import { Gas } from "src/app/services/api/admin/gas";

@Component({
    templateUrl: "./admin-star-system.component.html"
})
export class AdminStarSystemComponent implements OnInit {

    starSystem: StarSystem | undefined

    constructor(private readonly adminService: AdminService,
        private readonly activatedRoute: ActivatedRoute) {

    }

    async ngOnInit(): Promise<void> {
        const paramMap: ParamMap = await this.activatedRoute.paramMap.pipe(first()).toPromise()
        const name: string | null = paramMap.get("name")

        if (name) {
            const starSystem: StarSystem = await this.adminService.starSystem(name)
            this.initStarSystem(starSystem)

            this.starSystem = starSystem
            console.log(this.starSystem)
        }
    }

    private initStarSystem(starSystem: StarSystem): void {
        starSystem.planets = starSystem.planets.sort((p1: Planet, p2: Planet) => {
            return p1.name.localeCompare(p2.name)
        })

        starSystem.planets.forEach((planet:Planet)=>{
            planet.atmosphere = planet.atmosphere.sort((g1:Gas,g2:Gas)=>{
                return g1.num-g2.num
            })
        })
    }

    color(x: number): string {
        if (!x) {
            return `#FFFFFF`
        }
        return `#${x.toString(16)}`
    }

    solarMass2KG(x: number): number {
        if (!x) {
            return 0.0
        }
        return x * 1.98847e30
    }

    au2KM(x: number): number {
        if (!x) {
            return 0.0
        }
        // 1 AU = 149597870700 m
        return x * 149597870.700
    }

    kelvin2Degrees(x: number): number {
        if (!x) {
            return 0.0
        }
        return x - 272.15
    }

}