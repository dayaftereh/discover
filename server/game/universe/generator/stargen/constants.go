package stargen

import "github.com/dayaftereh/discover/server/game/universe/generator/stargen/types"

const (
	K                       float64 = 50.0     //K = gas/dust ratio
	B                       float64 = 1.2E-5   //Used in Crit_mass calc
	N                       float64 = 3.0      // Used in density calcs
	Alpha                   float64 = 5.0      // Used in density calcs
	ProtoPlanetMass         float64 = 1.0e-15  // Units of solar masses
	EccentricityCoefficient float64 = 0.077    /* Dole's was 0.077			*/
	ChangeInErthAngVel      float64 = -1.3E-15 /* Units of radians/sec/year*/
	DustDensityCoefficient  float64 = 2.0e-3   /* A in Dole's paper		*/

	GasRetentionThreshold float64 = 6.0 /* ratio of esc vel to RMS vel */

	AsteroidMassLimit    float64 = 0.001 /* Units of Earth Masses	*/
	SunMassInEarthMasses float64 = 332775.64
	SolarMassInGrams     float64 = 1.989e33 /* Units of grams			*/

	EarthRadius           float64 = 6.378e8  /* Units of cm				*/
	EarthRadiusInKM       float64 = 6378.0   /* Units of km				*/
	EarthDensity          float64 = 5.52     /* Units of g/cc			*/
	EarthMassInGrams      float64 = 5.977e27 /* Units of grams			*/
	EarthExosphereTemp    float64 = 1273.0   /* Units of degrees Kelvin	*/
	EarthAcceleration     float64 = 980.7    /* Units of cm/sec2			*/
	EarthAxialTilt        float64 = 23.4     /* Units of degrees*/
	EarthAverageCelsius   float64 = 14.0     /* Average Earth Temperature */
	EarthAverageKelvin    float64 = EarthAverageCelsius + FreezingPointOfWater
	EarthEffectiveTemp    float64 = 250.0    /* Units of degrees Kelvin (was 255)	*/
	EarthConvectionFactor float64 = 0.43     /* from Hart, eq.20			*/
	EarthWaterMassPerArea float64 = 3.83e15  /* grams per square km		*/
	DaysInAYear           float64 = 365.256  /* Earth days per Earth year*/
	CloudCoverageFactor   float64 = 1.839e-8 /* Km^2/kg					*/

	MaxSunAge float64 = 6e9
	MinSunAge float64 = 1e9

	EarthAlbedo             float64 = 0.33 /* was .33 for a while */
	IceAlbedo               float64 = 0.7
	CloudAlbedo             float64 = 0.52
	AirlessIceAlbedo        float64 = 0.5
	GreenhouseTriggerAlbedo float64 = 0.2
	RockyAlbedo             float64 = 0.15
	RockyAirlessAlbedo      float64 = 0.07
	WaterAlbedo             float64 = 0.04
	GasGaintAlbedo          float64 = 0.5 /* albedo of a gas giant	*/

	FreezingPointOfWater float64 = 273.15 /* Units of degrees Kelvin	*/
	MilliBarsPerBar      float64 = 1000.0

	GravityConstant float64 = 6.672E-8 /* units of dyne cm2/gram2	*/
	MolarGasConst   float64 = 8314.41  /* units: g*m2/(sec2*K*mol) */
	J               float64 = 1.46E-19 /* Used in day-length calcs (cm2/sec2 g) */
	SecondsPerHour  float64 = 3600.0

	CMPerAU    float64 = 1.495978707e13 /* number of cm in an AU	*/
	CMPerKM    float64 = 1.0e5          /* number of cm in a km		*/
	KMPerAU    float64 = CMPerAU / CMPerKM
	CMPerMeter float64 = 100.0

	/*	Now for a few molecular weights (used for RMS velocity calcs):	   */
	/*	This table is from Dole's book "Habitable Planets for Man", p. 38  */
	AtomicHydrogen   float64 = 1.0   /* H   */
	MolHydrogen      float64 = 2.0   /* H2  */
	Helium           float64 = 4.0   /* He  */
	AtomicNitrogen   float64 = 14.0  /* N   */
	AtomicOxygen     float64 = 16.0  /* O   */
	Methane          float64 = 16.0  /* CH4 */
	Ammonia          float64 = 17.0  /* NH3 */
	WaterVapro       float64 = 18.0  /* H2O */
	Neon             float64 = 20.2  /* Ne  */
	MolNitrogen      float64 = 28.0  /* N2  */
	CarbonMonoxide   float64 = 28.0  /* CO  */
	NitricOxide      float64 = 30.0  /* NO  */
	MolOxygen        float64 = 32.0  /* O2  */
	HydrogenSulphide float64 = 34.1  /* H2S */
	Argon            float64 = 39.9  /* Ar  */
	CarbonDioxide    float64 = 44.0  /* CO2 */
	NitrousOxide     float64 = 44.0  /* N2O */
	NitrogenDioxide  float64 = 46.0  /* NO2 */
	Ozone            float64 = 48.0  /* O3  */
	SulphDioxide     float64 = 64.1  /* SO2 */
	SulphTrioxode    float64 = 80.1  /* SO3 */
	Krypton          float64 = 83.8  /* Kr  */
	Xenon            float64 = 131.3 /* Xe  */

	/*	The following defines are used in the kothari_radius function in	*/
	A1_20     float64 = 6.485e12  /* All units are in cgs system.	 */
	A2_20     float64 = 4.0032e-8 /*	 ie: cm, g, dynes, etc.		 */
	Beta20    float64 = 5.71e12
	JimsFudge float64 = 1.004

	/*	 The following defines are used in determining the fraction of a planet	 */
	/*	covered with clouds in function cloud_fraction in file enviro.c.		 */
	Q1_36 float64 = 1.258e19 /* grams	*/
	Q2_36 float64 = 0.0698   /* 1/Kelvin */

	H2OAssumedPressure float64 = 47.0 * MMHG2MilliBars /* Dole p. 15      */
	PPMPressure        float64 = types.EarthSurfPersInMilliBars / 1000000.0
	MMHG2MilliBars     float64 = types.EarthSurfPersInMilliBars / types.EarthSurfPersInMMHG
)
