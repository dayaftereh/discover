package stargen

import (
	"bytes"
	"encoding/json"
)

type Gas struct {
	Num          int64
	SurfPressure float64
}

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

func (atom *Atom) String() string {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "  ")

	encoder.Encode(atom)

	return buffer.String()
}

var GasesTable map[int64]*Atom = map[int64]*Atom{
	//   An | Sym | name | Aw | melt | boil | dens | ABUNDe | ABUNDs | Rea | Max inspired pp
	1:  &Atom{1, "H", "Hydrogen", 1.0079, 14.06, 20.40, 8.99e-05, 0.00125893, 27925.4, 1, 0.0},
	2:  &Atom{2, "He", "Helium", 4.0026, 3.46, 4.20, 0.0001787, 7.94328e-09, 2722.7, 0, MaxHeIPP},
	7:  &Atom{7, "N", "Nitrogen", 14.0067, 63.34, 77.40, 0.0012506, 1.99526e-05, 3.13329, 0, MaxN2IPP},
	8:  &Atom{8, "O", "Oxygen", 15.9994, 54.80, 90.20, 0.001429, 0.501187, 23.8232, 10, MaxO2IPP},
	10: &Atom{10, "Ne", "Neon", 20.1700, 24.53, 27.10, 0.0009, 5.01187e-09, 3.4435e-5, 0, MaxNeIPP},
	18: &Atom{18, "Ar", "Argon", 39.9480, 84.00, 87.30, 0.0017824, 3.16228e-06, 0.100925, 0, MaxArIPP},
	36: &Atom{36, "Kr", "Krypton", 83.8000, 116.60, 119.70, 0.003708, 1e-10, 4.4978e-05, 0, MaxKrIPP},
	54: &Atom{54, "Xe", "Xenon", 131.3000, 161.30, 165.00, 0.00588, 3.16228e-11, 4.69894e-06, 0, MaxXeIPP},
	//
	900: &Atom{900, "NH3", "Ammonia", 17.0000, 195.46, 239.66, 0.001, 0.002, 0.0001, 1, MaxNh3IPP},
	901: &Atom{901, "H2O", "Water", 18.0000, 273.16, 373.16, 1.000, 0.03, 0.001, 0, 0.0},
	902: &Atom{902, "CO2", "CarbonDioxide", 44.0000, 194.66, 194.66, 0.001, 0.01, 0.0005, 0, MaxCo2IPP},
	903: &Atom{903, "O3", "Ozone", 48.0000, 80.16, 161.16, 0.001, 0.001, 0.000001, 2, MaxO3IPP},
	904: &Atom{904, "CH4", "Methane", 16.0000, 90.16, 109.16, 0.010, 0.005, 0.0001, 1, MaxCh4IPP},
}

type Oxygen string

const (
	None         Oxygen = "none"
	Toxic        Oxygen = "toxic"
	Breathable   Oxygen = "breathable"
	Unbreathable Oxygen = "unbreathable"
)
