package stargen

import (
	"bytes"
	"encoding/json"
)

type PlanetType string

type OrbitZone int

type Planet struct {
	PlanetNO                        int64
	SemiMajorAxis                   float64   /* A - (Distance from primary star) semi-major axis of solar orbit (in AU)*/
	Eccentricity                    float64   /* E - eccentricity of solar orbit		 */
	AxialTilt                       float64   /* units of degrees					 */
	Mass                            float64   /* mass (in solar masses)			 */
	GasGiant                        bool      /* TRUE if the planet is a gas giant */
	DustMass                        float64   /* mass, ignoring gas				 */
	GasMass                         float64   /* mass, ignoring dust				 */
	MoonA                           float64   /* semi-major axis of lunar orbit (in AU)*/
	MoonE                           float64   /* eccentricity of lunar orbit		 */
	CoreRadius                      float64   /* radius of the rocky core (in km)	 */
	Radius                          float64   /* equatorial radius (in km)		 */
	OrbitZone                       OrbitZone /* the 'zone' of the planet			 */
	Density                         float64   /* density (in g/cc)				 */
	OrbitPeriod                     float64   /* length of the local year (days)	 */
	Day                             float64   /* length of the local day (hours)	 */
	ResonantPeriod                  bool      /* TRUE if in resonant rotation (Planet's rotation is in a resonant spin lock with the star)		 */
	EscapeVelocity                  float64   /* Escape Velocity (cm/sec)				 */
	SurfaceAcceleration             float64   /* Surface acceleration (cm/sec2	)				 */
	SurfaceGravity                  float64   /* Surface gravity (Earth gravities)			 */
	RootMeanSquareVelocity          float64   /* Root Mean Square Velocity units of cm/sec					 */
	MolecularWeight                 float64   /* Molecular weight smallest molecular weight retained*/
	VolatileGasInventory            float64
	SurfacePressure                 float64 /* Surface pressure (millibars [mb])			 */
	GreenhouseEffect                bool    /* runaway greenhouse effect?		 */
	BoilPoint                       float64 /* the boiling point of water (Kelvin)*/
	Albedo                          float64 /* albedo of the planet				 */
	ExosphericTemperature           float64 /* units of degrees Kelvin			 */
	EstimatedTemperature            float64 /* quick non-iterative estimate (K)  */
	EstimatedTerrestrialTemperature float64 /* for terrestrial moons and the like*/
	SurfaceTemperature              float64 /* surface temperature in Kelvin	 */
	GreenhouseRise                  float64 /* Temperature rise due to greenhouse */
	HighTemperature                 float64 /* Day-time temperature              */
	LowTemperature                  float64 /* Night-time temperature			 */
	MaxTemperature                  float64 /* Summer/Day						 */
	MinTemperature                  float64 /* Winter/Night						 */
	Hydrosphere                     float64 /* fraction of surface covered (%)		 */
	CloudCover                      float64 /* fraction of surface covered	(%)		 */
	IceCover                        float64 /* fraction of surface covered	(%)		 */
	Atmosphere                      []*Gas
	Type                            PlanetType /* Type code						 */
	Moons                           []*Planet
	Breathability                   Oxygen
}

const (
	// Planet Type
	PlanetUnknown        PlanetType = "unknown"
	PlanetRock           PlanetType = "rock"        // callisto
	PlanetVenusian       PlanetType = "vebusian"    // venuslike
	PlanetTerrestrial    PlanetType = "terrestrial" // Earthlike
	PlanetGasGiant       PlanetType = "gas-gaint"   // jupiterlike Jovian
	PlanetMartian        PlanetType = "martian"     // planet like mars
	PlanetWater          PlanetType = "water"
	PlanetIce            PlanetType = "ice"               // pluto
	PlanetSubGasGiant    PlanetType = "sub-gas-gaint"     // gasgiant Sub-Jovian
	PlanetSubSubGasGiant PlanetType = "sub-sub-gas-gaint" // GasDwarf
	PlanetAsteroids      PlanetType = "asteroids"
	Planet1Face          PlanetType = "face"
	// OrbitType
	Orbit1 OrbitZone = 1
	Orbit2 OrbitZone = 2
	Orbit3 OrbitZone = 3
)

func (planet *Planet) String() string {

	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "  ")

	encoder.Encode(planet)

	return buffer.String()
}
