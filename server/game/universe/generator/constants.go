package generator

const (
	K                    float64 = 50.0    //K = gas/dust ratio
	B                    float64 = 1.2E-5  //Used in Crit_mass calc
	N                    float64 = 3.0     // Used in density calcs
	Alpha                float64 = 5.0     // Used in density calcs
	ProtoPlanetMass      float64 = 1.0e-15 // Units of solar masses
	EccentricityCoeff    float64 = 0.077
	SunMassInEarthMasses float64 = 332775.64

	EarthAxialTilt float64 = 23.4    /* Units of degrees*/
	DaysInAYear    float64 = 365.256 /* Earth days per Earth year*/

	MaxSunAge float64 = 6e9
	MinSunAge float64 = 1e9
)
