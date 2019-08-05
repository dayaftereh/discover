package types

type Atom struct {
	Num        int64
	Symbol     string
	Name       string
	Weight     float64
	Meth       float64
	Boil       float64
	Density    float64
	Abunde     float64
	Abunds     float64
	Reactivity float64
	MaxIpp     float64 // Max inspired partial pressure im millibars
}

const (
	EarthSurfPersInMMHG      float64 = 760.0 /* Dole p. 15 */
	EarthSurfPersInMilliBars float64 = 1013.25
	PPMPressure              float64 = EarthSurfPersInMilliBars / 1000000.0
	MMHG2MilliBars           float64 = EarthSurfPersInMilliBars / EarthSurfPersInMMHG
	// Gas IPP

	MinO2IPP  float64 = (72.0 * MMHG2MilliBars)  /* Dole p. 15      */
	MaxO2IPP  float64 = 400.0 * MMHG2MilliBars   /* Dole p. 15      */
	MaxHeIPP  float64 = 61000.0 * MMHG2MilliBars /* Dole p. 16      */
	MaxNeIPP  float64 = 3900.0 * MMHG2MilliBars  /* Dole p. 16      */
	MaxN2IPP  float64 = 2330.0 * MMHG2MilliBars  /* Dole p. 16      */
	MaxArIPP  float64 = 1220.0 * MMHG2MilliBars  /* Dole p. 16      */
	MaxKrIPP  float64 = 350.0 * MMHG2MilliBars   /* Dole p. 16      */
	MaxXeIPP  float64 = 160.0 * MMHG2MilliBars   /* Dole p. 16      */
	MaxCo2IPP float64 = 7.0 * MMHG2MilliBars     /* Dole p. 16      */
	// The next gases are listed as poisonous in parts per million by volume at 1 atm:
	MaxFIPP   float64 = 0.1 * PPMPressure    /* Dole p. 18      */
	MaxClIPP  float64 = 1.0 * PPMPressure    /* Dole p. 18      */
	MaxNh3IPP float64 = 100. * PPMPressure   /* Dole p. 18      */
	MaxO3IPP  float64 = 0.1 * PPMPressure    /* Dole p. 18      */
	MaxCh4IPP float64 = 50000. * PPMPressure /* Dole p. 18      */
)

var (
	//   An | Sym | name | Aw | melt | boil | dens | ABUNDe | ABUNDs | Rea | Max inspired pp
	Hydrogen = &Atom{1, "H", "Hydrogen", 1.0079, 14.06, 20.40, 8.99e-05, 0.00125893, 27925.4, 1, 0.0}
	Helium   = &Atom{2, "He", "Helium", 4.0026, 3.46, 4.20, 0.0001787, 7.94328e-09, 2722.7, 0, MaxHeIPP}
	Nitrogen = &Atom{7, "N", "Nitrogen", 14.0067, 63.34, 77.40, 0.0012506, 1.99526e-05, 3.13329, 0, MaxN2IPP}
	Oxygen   = &Atom{8, "O", "Oxygen", 15.9994, 54.80, 90.20, 0.001429, 0.501187, 23.8232, 10, MaxO2IPP}
	Neon     = &Atom{10, "Ne", "Neon", 20.1700, 24.53, 27.10, 0.0009, 5.01187e-09, 3.4435e-5, 0, MaxNeIPP}
	Argon    = &Atom{18, "Ar", "Argon", 39.9480, 84.00, 87.30, 0.0017824, 3.16228e-06, 0.100925, 0, MaxArIPP}
	Krypton  = &Atom{36, "Kr", "Krypton", 83.8000, 116.60, 119.70, 0.003708, 1e-10, 4.4978e-05, 0, MaxKrIPP}
	Xenon    = &Atom{54, "Xe", "Xenon", 131.3000, 161.30, 165.00, 0.00588, 3.16228e-11, 4.69894e-06, 0, MaxXeIPP}
	//
	Ammonia       = &Atom{900, "NH3", "Ammonia", 17.0000, 195.46, 239.66, 0.001, 0.002, 0.0001, 1, MaxNh3IPP}
	Water         = &Atom{901, "H2O", "Water", 18.0000, 273.16, 373.16, 1.000, 0.03, 0.001, 0, 0.0}
	CarbonDioxide = &Atom{902, "CO2", "CarbonDioxide", 44.0000, 194.66, 194.66, 0.001, 0.01, 0.0005, 0, MaxCo2IPP}
	Ozone         = &Atom{903, "O3", "Ozone", 48.0000, 80.16, 161.16, 0.001, 0.001, 0.000001, 2, MaxO3IPP}
	Methane       = &Atom{904, "CH4", "Methane", 16.0000, 90.16, 109.16, 0.010, 0.005, 0.0001, 1, MaxCh4IPP}
)

var ChemicalElements map[int64]*Atom = map[int64]*Atom{
	1:  Hydrogen,
	2:  Helium,
	7:  Nitrogen,
	8:  Oxygen,
	10: Neon,
	18: Argon,
	36: Krypton,
	54: Xenon,
	//
	900: Ammonia,
	901: Water,
	902: CarbonDioxide,
	903: Ozone,
	904: Methane,
}

var AtmosphereGases []*Atom = []*Atom{
	Hydrogen, Helium, Nitrogen, Oxygen, Neon, Argon, Krypton, Xenon /**/, Ammonia, Water, CarbonDioxide, Ozone, Methane,
}
