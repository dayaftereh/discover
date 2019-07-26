package generator

/*

sun->mass		 = random_number(0.7,1.4);

outer_dust_limit	 = stellar_dust_limit(sun->mass);
sun->luminosity	 = luminosity(sun->mass);
sun->r_ecosphere	 = sqrt(sun->luminosity);
sun->life			 = 1.0E10 * (sun->mass / sun->luminosity); // age of the sun

long double min_age = 1.0E9;
long double max_age = 6.0E9;

if (sun->life < max_age)
	max_age = sun->life;

outer_planet_limit := 0.0 // wenn kein dopplestar

dust_density_coeff := 2.0e-3

innermost_planet = dist_planetary_masses(sun->mass,
										sun->luminosity,
										0.0,
										outer_dust_limit,
										outer_planet_limit,
										dust_density_coeff,
										seed_system,
										do_moons);
sun->age = random_number(min_age, max_age);

generate_planets()

*/

/*

dist_planetary_masses: long double stell_mass_ratio,
									 long double stell_luminosity_ratio, 
									 long double inner_dust, 
									 long double outer_dust,
									 long double outer_planet_limit,
									 long double dust_density_coeff,
									 planet_pointer seed_system,
									 int		 do_moons

set_initial_conditions(inner_dust,outer_dust);

planet_inner_bound = nearest_planet(stell_mass_ratio);
planet_outer_bound = farthest_planet(stell_mass_ratio);

for dust_left {

	a = random_number(planet_inner_bound,planet_outer_bound);
	e = random_eccentricity( );

	mass      = PROTOPLANET_MASS; // 1.0E-15 Units of solar masses
	dust_mass = 0;
	gas_mass  = 0;

	if (dust_available(inner_effect_limit(a, e, mass),  outer_effect_limit(a, e, mass))) {

		dust_density = dust_density_coeff * sqrt(stell_mass_ratio)  * exp(-ALPHA * pow(a,(1.0 / N)));
		crit_mass = critical_limit(a,e,stell_luminosity_ratio);

		accrete_dust(&mass, &dust_mass, &gas_mass,
						 a,e,crit_mass,
						 planet_inner_bound,
						 planet_outer_bound);

	}

}

*/
