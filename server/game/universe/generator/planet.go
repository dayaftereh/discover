package generator

import "fmt"

type PlanetType string

type OrbitZone int

type Gas struct {
	Num          int64
	SurfPressure float64
}

type Planet struct {
	PlanetNO             int64
	A                    float64   /* semi-major axis of solar orbit (in AU)*/
	E                    float64   /* eccentricity of solar orbit		 */
	AxialTilt            float64   /* units of degrees					 */
	Mass                 float64   /* mass (in solar masses)			 */
	GasGiant             bool      /* TRUE if the planet is a gas giant */
	DustMass             float64   /* mass, ignoring gas				 */
	GasMass              float64   /* mass, ignoring dust				 */
	MoonA                float64   /* semi-major axis of lunar orbit (in AU)*/
	MoonE                float64   /* eccentricity of lunar orbit		 */
	CoreRadius           float64   /* radius of the rocky core (in km)	 */
	Radius               float64   /* equatorial radius (in km)		 */
	OrbitZone            OrbitZone /* the 'zone' of the planet			 */
	Density              float64   /* density (in g/cc)				 */
	OrbPeriod            float64   /* length of the local year (days)	 */
	Day                  float64   /* length of the local day (hours)	 */
	ResonantPeriod       bool      /* TRUE if in resonant rotation		 */
	ESCVelocity          float64   /* units of cm/sec					 */
	SurfAccel            float64   /* units of cm/sec2					 */
	SurfGrav             float64   /* units of Earth gravities			 */
	RMSVelocity          float64   /* units of cm/sec					 */
	MolecWeight          float64   /* smallest molecular weight retained*/
	VolatileGasInventory float64
	SurfPressure         float64 /* units of millibars (mb)			 */
	GreenhouseEffect     bool    /* runaway greenhouse effect?		 */
	BoilPoint            float64 /* the boiling point of water (Kelvin)*/
	Albedo               float64 /* albedo of the planet				 */
	ExosphericTemp       float64 /* units of degrees Kelvin			 */
	EstimatedTemp        float64 /* quick non-iterative estimate (K)  */
	EstimatedTerrTemp    float64 /* for terrestrial moons and the like*/
	SurfTemp             float64 /* surface temperature in Kelvin	 */
	GreenhsRise          float64 /* Temperature rise due to greenhouse */
	HighTemp             float64 /* Day-time temperature              */
	LowTemp              float64 /* Night-time temperature			 */
	MaxTemp              float64 /* Summer/Day						 */
	MinTemp              float64 /* Winter/Night						 */
	Hydrosphere          float64 /* fraction of surface covered		 */
	CloudCover           float64 /* fraction of surface covered		 */
	IceCover             float64 /* fraction of surface covered		 */
	Gases                int     /* Count of gases in the atmosphere: */
	Atmosphere           *Gas
	Type                 PlanetType /* Type code						 */
	Moons                []*Planet
	/*   ZEROES end here               */
	NextPlanet *Planet
}

const (
	// Planet Type
	PlanetUnknown        PlanetType = "unknown"
	PlanetRock           PlanetType = "rock"
	PlanetVenusian       PlanetType = "vebusian"
	PlanetTerrestrial    PlanetType = "terrestrial"
	PlanetGasGiant       PlanetType = "gas-gaint"
	PlanetMartian        PlanetType = "martian"
	PlanetWater          PlanetType = "water"
	PlanetIce            PlanetType = "ice"
	PlanetSubGasGiant    PlanetType = "sub-gas-gaint"
	PlanetSubSubGasGiant PlanetType = "sub-sub-gas-gaint"
	PlanetAsteroids      PlanetType = "asteroids"
	Planet1Face          PlanetType = "face"
	// OrbitType
	Orbit1 OrbitZone = 1
	Orbit2 OrbitZone = 2
	Orbit3 OrbitZone = 3
)

func (planet *Planet) String() string {
	s := fmt.Sprintf("Planet: [\n")
	s = fmt.Sprintf("%s A: %f\n", s, planet.A)
	s = fmt.Sprintf("%s E: %f\n", s, planet.A)
	s = fmt.Sprintf("%s Mass: %.19f\n", s, planet.Mass)
	s = fmt.Sprintf("%s GasGiant: %v\n", s, planet.GasGiant)
	s = fmt.Sprintf("%s DustMass: %f\n", s, planet.DustMass)
	s = fmt.Sprintf("%s GasMass: %f\n", s, planet.GasMass)
	s = fmt.Sprintf("%s]\n", s)
	return s
}
