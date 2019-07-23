package mathf

import "math"

var EclipticNormal *Vec3 = NewVec3(0, 1, 0)

type Orbit2 struct {
	Position              *Vec3
	Velocity              *Vec3
	AttractorMass         float64
	GravitationalConstant float64
	Eccentricity          float64
	Period                float64
	MeanAnomaly           float64
	SemiMajorAxis         float64
	SemiMinorAxis         float64
	TrueAnomaly           float64
	EccentricAnomaly      float64
	FocalParameter        float64
	SemiMinorAxisBasis    *Vec3
	SemiMajorAxisBasis    *Vec3
	CenterPoint           *Vec3
}

func NewOrbitFromOrbitalVectors(attractorMass float64, gravitationalConstant float64, position *Vec3, velocity *Vec3) *Orbit2 {
	orbit := &Orbit2{
		Position:              position,
		Velocity:              velocity,
		AttractorMass:         attractorMass,
		GravitationalConstant: gravitationalConstant,
	}

	return orbit
}

func (orbit *Orbit2) CalculateOrbitStateFromOrbitalVectors() {
	MG := orbit.AttractorMass * orbit.GravitationalConstant
	AttractorDistance := NewVec3(0, 0, 0).DistanceTo(orbit.Position)
	angularMomentumVector := orbit.Position.Cross(orbit.Velocity)
	OrbitNormal := angularMomentumVector.Normalize()

	var eccVector *Vec3
	if OrbitNormal.SqrtLength() < 0.99 {
		OrbitNormal = orbit.Position.Cross(EclipticNormal)
		eccVector = NewZeroVec3()
	} else {
		eccVector = orbit.Velocity.Cross(angularMomentumVector).DivideVec(orbit.Position.Subtract(MG)).Divide(AttractorDistance)
	}

	//OrbitNormalDotEclipticNormal := OrbitNormal.Dot(EclipticNormal)
	orbit.FocalParameter = angularMomentumVector.SqrtLength() / MG
	orbit.Eccentricity = eccVector.Length()
	//EnergyTotal := orbit.Velocity.SqrtLength() - 2*MG/AttractorDistance
	orbit.SemiMinorAxisBasis = angularMomentumVector.Cross(eccVector.Negate()).Normalize()
	if orbit.SemiMinorAxisBasis.SqrtLength() < 0.99 {
		orbit.SemiMinorAxisBasis = OrbitNormal.Cross(orbit.Position)
	}
	orbit.SemiMajorAxisBasis = OrbitNormal.Cross(orbit.SemiMinorAxisBasis).Normalize()
	if orbit.Eccentricity < 1.0 {
		OrbitCompressionRatio := 1.0 - orbit.Eccentricity*orbit.Eccentricity
		orbit.SemiMajorAxis = orbit.FocalParameter / OrbitCompressionRatio
		orbit.SemiMinorAxis = orbit.SemiMajorAxis * math.Sqrt(OrbitCompressionRatio)
		orbit.CenterPoint = eccVector.Multiply(-orbit.SemiMajorAxis)
		orbit.Period = (math.Pi * 2.0) * math.Sqrt(math.Pow(orbit.SemiMajorAxis, 3)/MG)
		//Apoapsis := orbit.CenterPoint.SubtractVec(orbit.SemiMajorAxisBasis.Multiply(orbit.SemiMajorAxis))
		//Periapsis := orbit.CenterPoint.AddVec(orbit.SemiMajorAxisBasis.Multiply(orbit.SemiMajorAxis))
		//PeriapsisDistance := Periapsis.Length()
		//ApoapsisDistance := Apoapsis.Length()
		orbit.TrueAnomaly = orbit.Position.AngleTo(orbit.SemiMajorAxisBasis)
		if orbit.Position.Cross(orbit.SemiMajorAxisBasis.Negate()).Dot(OrbitNormal) < 0.0 {
			orbit.TrueAnomaly = (math.Pi * 2.0) - orbit.TrueAnomaly
		}
		orbit.EccentricAnomaly = ConvertTrueToEccentricAnomaly(orbit.TrueAnomaly, orbit.Eccentricity)
		orbit.MeanAnomaly = orbit.EccentricAnomaly - orbit.Eccentricity*math.Sin(orbit.EccentricAnomaly)
	} else {
		OrbitCompressionRatio := orbit.Eccentricity*orbit.Eccentricity - 1
		orbit.SemiMajorAxis = orbit.FocalParameter / OrbitCompressionRatio
		orbit.SemiMinorAxis = orbit.SemiMajorAxis * math.Sqrt(OrbitCompressionRatio)
		orbit.CenterPoint = eccVector.Multiply(orbit.SemiMinorAxis)
		orbit.Period = math.Inf(1)
		//Apoapsis := NewVec3(math.Inf(1), math.Inf(1), math.Inf(1))
		//Periapsis := orbit.CenterPoint.SubtractVec(orbit.SemiMajorAxisBasis.Multiply(orbit.SemiMajorAxis))
		//PeriapsisDistance := Periapsis.Length()
		//ApoapsisDistance := math.Inf(1)
		orbit.TrueAnomaly = orbit.Position.AngleTo(eccVector)
		if orbit.Position.Cross(orbit.SemiMajorAxisBasis.Negate()).Dot(OrbitNormal) < 0.0 {
			orbit.TrueAnomaly = -orbit.TrueAnomaly
		}
		orbit.EccentricAnomaly = ConvertTrueToEccentricAnomaly(orbit.TrueAnomaly, orbit.Eccentricity)
		orbit.MeanAnomaly = math.Sinh(orbit.EccentricAnomaly)*orbit.Eccentricity - orbit.EccentricAnomaly
	}
}

func (orbit *Orbit2) MeanMotion() float64 {
	if orbit.Eccentricity < 1.0 {
		return (math.Pi * 2.0) / orbit.Period
	}
	return math.Sqrt(orbit.AttractorMass * orbit.GravitationalConstant / math.Pow(orbit.SemiMajorAxis, 3.0))
}

func (orbit *Orbit2) UpdateOrbitDataByTime(delta float64) {
	orbit.UpdateOrbitAnomaliesByTime(delta)
	orbit.SetPositionByCurrentAnomaly()
	orbit.SetVelocityByCurrentAnomaly()
}

func (orbit *Orbit2) UpdateOrbitAnomaliesByTime(delta float64) {
	if orbit.Eccentricity < 1 {
		if orbit.Period > 1.401298e-45 {
			orbit.MeanAnomaly += orbit.MeanMotion() * delta
		}
		orbit.MeanAnomaly = math.Mod(orbit.MeanAnomaly, math.Pi*2.0)
		if orbit.MeanAnomaly < 0 {
			orbit.MeanAnomaly = (math.Pi * 2.0) - orbit.MeanAnomaly
		}
		EccentricAnomaly := KeplerSolver(orbit.MeanAnomaly, orbit.Eccentricity)
		cosE := math.Cos(EccentricAnomaly)
		orbit.TrueAnomaly = math.Acos((cosE - orbit.Eccentricity) / (1 - orbit.Eccentricity*cosE))
		if orbit.TrueAnomaly > math.Pi {
			orbit.TrueAnomaly = (math.Pi * 2.0) - orbit.TrueAnomaly
		}
	}

}

func (orbit *Orbit2) GetFocalPositionAtEccentricAnomaly(eccentricAnomaly float64) *Vec3 {
	return orbit.GetCentralPositionAtEccentricAnomaly(eccentricAnomaly).AddVec(orbit.CenterPoint)
}

func (orbit *Orbit2) GetCentralPositionAtEccentricAnomaly(eccentricAnomaly float64) *Vec3 {
	var result *Vec3
	if orbit.Eccentricity < 1 {
		result = NewVec3(
			math.Sin(eccentricAnomaly)*orbit.SemiMinorAxis,
			-math.Cos(eccentricAnomaly)*orbit.SemiMajorAxis,
			0,
		)
	} else {
		result = NewVec3(
			math.Sinh(eccentricAnomaly)*orbit.SemiMinorAxis,
			math.Cosh(eccentricAnomaly)*orbit.SemiMajorAxis,
			0,
		)
	}
	return orbit.SemiMinorAxisBasis.Multiply(-result.X).SubtractVec(orbit.SemiMajorAxisBasis.Multiply(result.Y))
}

func (orbit *Orbit2) SetPositionByCurrentAnomaly() {
	orbit.Position = orbit.GetFocalPositionAtEccentricAnomaly(orbit.EccentricAnomaly)
}

func (orbit *Orbit2) SetVelocityByCurrentAnomaly() {
	orbit.Velocity = orbit.GetVelocityAtEccentricAnomaly(orbit.EccentricAnomaly)
}

func (orbit *Orbit2) GetVelocityAtEccentricAnomaly(eccentricAnomaly float64) *Vec3 {
	return orbit.GetVelocityAtTrueAnomaly(ConvertEccentricToTrueAnomaly(eccentricAnomaly, orbit.Eccentricity))
}

func (orbit *Orbit2) GetVelocityAtTrueAnomaly(trueAnomaly float64) *Vec3 {
	if orbit.FocalParameter < 1.401298e-45 {
		return NewZeroVec3()
	}
	sqrtMGdivP := math.Sqrt(orbit.AttractorMass * orbit.GravitationalConstant / orbit.FocalParameter)
	vX := sqrtMGdivP * (orbit.Eccentricity + math.Cos(trueAnomaly))
	vY := sqrtMGdivP * math.Sin(trueAnomaly)
	return orbit.SemiMinorAxisBasis.Multiply(-vX).SubtractVec(orbit.SemiMajorAxisBasis.Multiply(vY))
}

func KeplerSolver(meanAnomaly float64, eccentricity float64) float64 {
	iterations := 3
	if eccentricity < 0.2 {
		iterations = 2
	}

	e := meanAnomaly
	for i := 0; i < iterations; i++ {
		esinE := eccentricity * math.Sin(e)
		ecosE := eccentricity * math.Cos(e)
		deltaE := e - esinE - meanAnomaly
		n := 1.0 - ecosE
		e += -5 * deltaE / (n + math.Sin(n)*math.Sqrt(math.Abs(16.0*n*n-20*deltaE*esinE)))
	}
	return e
}

func ConvertEccentricToTrueAnomaly(eccentricAnomaly float64, eccentricity float64) float64 {
	if eccentricity < 1.0 {
		cosE := math.Cos(eccentricAnomaly)
		tAnom := math.Acos((cosE - eccentricity) / (1.0 - eccentricity*cosE))
		if eccentricAnomaly > math.Pi {
			tAnom = (math.Pi * 2.0) - tAnom
		}
		return tAnom
	}
	tAnom := math.Atan2(
		math.Sqrt(eccentricity*eccentricity-1.0)*math.Sinh(eccentricAnomaly),
		eccentricity-math.Cosh(eccentricAnomaly),
	)
	return tAnom
}

func ConvertTrueToEccentricAnomaly(trueAnomaly float64, eccentricity float64) float64 {
	if math.IsNaN(eccentricity) || math.IsInf(eccentricity, 0) {
		return trueAnomaly
	}
	trueAnomaly = math.Mod(trueAnomaly, math.Pi*2.0)
	if eccentricity < 1.0 {
		if trueAnomaly < 0 {
			trueAnomaly = trueAnomaly + (math.Pi * 2.0)
		}
		cosT2 := math.Cos(trueAnomaly)
		eccAnom := math.Acos((eccentricity + cosT2) / (1 + eccentricity*cosT2))
		if trueAnomaly > math.Pi {
			eccAnom = (math.Pi * 2.0) - eccAnom
		}
		return eccAnom
	}

	cosT2 := math.Cos(trueAnomaly)
	eccAnom := math.Acosh((eccentricity+cosT2)/(1+eccentricity*cosT2)) * math.Sin(trueAnomaly)
	return eccAnom
}
