package mathf

import (
	"fmt"
	"log"
	"math"
)

// http://quantumg.net/orbelems.cpp

type Orbit struct {
	epoch                    *Vec3
	mass                     float64
	semiMajorAxis            float64 // a
	eccentricity             float64 // e
	inclination              float64 // i
	longitudeOfAscendingNode float64 // l
	argumentOfPeriapsis      float64 // w
	trueAnomaly              float64 // t
}

func NewOrbitFromVectors(mass float64, radius *Vec3, velocity *Vec3) *Orbit {
	orbit := &Orbit{
		mass:  mass,
		epoch: NewZeroVec3(),
	}
	// calculate specific relative angular momement
	h := radius.Cross(velocity)

	// calculate vector to the ascending node
	n := NewVec3(-h.Y, h.X, 0.0)

	// standard gravity GM
	u := GravitationalConstant * mass

	// calculate eccentricity vector and scalar
	e := velocity.MultiplyVec(h).Multiply(1.0 / u).SubtractVec(radius.Multiply(1.0 / radius.Length()))
	orbit.eccentricity = e.Length()

	// calculate specific orbital energy and semi-major axis
	E := velocity.SqrtLength()*0.5 - u/radius.Length()
	orbit.semiMajorAxis = -u / (E * 2.0)

	// calculate inclination
	orbit.inclination = math.Acos(h.Z / h.Length())

	// calculate longitude of ascending node
	if orbit.inclination == 0.0 {
		orbit.longitudeOfAscendingNode = 0.0
	} else if n.Y >= 0.0 {
		orbit.longitudeOfAscendingNode = math.Acos(n.X / n.Length())
	} else {
		orbit.longitudeOfAscendingNode = math.Pi*2.0 - math.Acos(n.X/n.Length())
	}

	// calculate argument of periapsis
	if orbit.inclination == 0.0 {
		orbit.argumentOfPeriapsis = math.Acos(e.X / e.Length())
	} else if e.Z >= 0.0 {
		orbit.argumentOfPeriapsis = math.Acos(n.Dot(e) / (n.Length() * e.Length()))
	} else {
		orbit.argumentOfPeriapsis = 2.0*math.Pi - math.Acos(n.Dot(e)/(n.Length()*e.Length()))
	}

	// calculate true anomaly
	if radius.Dot(velocity) >= 0.0 {
		orbit.trueAnomaly = math.Acos(e.Dot(radius) / (e.Length() * radius.Length()))
	} else {
		orbit.trueAnomaly = 2.0*math.Pi - math.Acos(e.Dot(radius)/(e.Length()*radius.Length()))
	}

	// calculate epoch
	position := orbit.Position()
	orbit.epoch = position.SubtractVec(radius)

	return orbit
}

func (orbit *Orbit) semiparameter() float64 {
	return orbit.trueAnomaly * (1.0 - orbit.eccentricity*orbit.eccentricity)
}

func (orbit *Orbit) Position() *Vec3 {
	p := orbit.semiparameter()

	lSin := math.Sin(orbit.longitudeOfAscendingNode)
	lCos := math.Cos(orbit.longitudeOfAscendingNode)

	iSin := math.Sin(orbit.inclination)
	iCos := math.Cos(orbit.inclination)

	wtSin := math.Sin(orbit.argumentOfPeriapsis + orbit.trueAnomaly)
	wtCos := math.Cos(orbit.argumentOfPeriapsis + orbit.trueAnomaly)

	r := NewVec3(
		p*(lCos*wtCos-lSin*iCos*wtSin),
		p*(lSin*wtCos+lCos*iCos*wtSin),
		p*iSin*wtSin,
	)

	return r.SubtractVec(orbit.epoch)
}

func (orbit *Orbit) Velocity() *Vec3 {
	p := orbit.semiparameter()
	// standard gravity GM
	u := GravitationalConstant * orbit.mass

	g := -math.Sqrt(u / p)

	lSin := math.Sin(orbit.longitudeOfAscendingNode)
	lCos := math.Cos(orbit.longitudeOfAscendingNode)

	iSin := math.Sin(orbit.inclination)
	iCos := math.Cos(orbit.inclination)

	wSin := math.Sin(orbit.argumentOfPeriapsis)
	wCos := math.Cos(orbit.argumentOfPeriapsis)

	wtSin := math.Sin(orbit.argumentOfPeriapsis + orbit.trueAnomaly)
	wtCos := math.Cos(orbit.argumentOfPeriapsis + orbit.trueAnomaly)

	e := orbit.eccentricity

	/*
			v.x = g * (cos(l)          * (sin(w + t) + e * sin(w)) +
		            	sin(l) * cos(i) * (cos(w + t) + e * cos(w)));
		    v.y = g * (sin(l)          * (sin(w + t) + e * sin(w)) -
		                cos(l) * cos(i) * (cos(w + t) + e * cos(w)));
		    v.z = g * (sin(i) * (cos(w + t) + e * cos(w)));
	*/

	return NewVec3(
		g*(lCos*(wtSin+e*wSin)+lSin*iCos*(wtCos+e*wCos)),
		g*(lSin*(wtSin+e*wSin)-lCos*iCos*(wtCos+e*wCos)),
		g*(iSin*(wtCos+e*wCos)),
	)
}

func (orbit *Orbit) period() float64 {
	a := math.Abs(orbit.semiMajorAxis)
	u := GravitationalConstant * orbit.mass
	return 2 * math.Pi * math.Sqrt(a*a*a/u)
}

func (orbit *Orbit) eccentricAnomaly() float64 {
	e := orbit.eccentricity
	t := orbit.trueAnomaly

	//E := math.Acos((e + math.Cos(t)/(1.0+e*math.Cos(t))))

	//E := 2.0 * math.Atan(math.Sqrt(1-e)/(1+e)*math.Tan(t/2.0))

	E := math.Atan2(math.Sqrt((1-e)/(1+e))*math.Sin(t/2.0), math.Cos(t/2.0))

	if t > math.Pi && E < math.Pi {
		E = 2.0*math.Pi - E
	}
	return E
}

func (orbit *Orbit) meanAnomaly() float64 {
	E := orbit.eccentricAnomaly()
	M := E - orbit.eccentricity*math.Sin(E)
	if E > math.Pi && M < math.Pi {
		M = 2.0*math.Pi - M
	}
	return M
}

func (orbit *Orbit) meanMotion() float64 {
	// standard gravity GM
	u := GravitationalConstant * orbit.mass
	a := math.Abs(orbit.semiMajorAxis)
	return math.Sqrt(u / (a * a * a))
}

func (orbit *Orbit) estimateTrueAnomaly(meanAnomaly float64) float64 {
	M := meanAnomaly
	e := orbit.eccentricity
	return M + 2.0*e*math.Sin(M) + 1.25*e*e*math.Sin(2.0*M)
}

func (orbit *Orbit) calcEccentricAnomaly(meanAnomaly float64) float64 {
	t := orbit.estimateTrueAnomaly(meanAnomaly)
	E := math.Acos((orbit.eccentricity + math.Cos(t)) / (1.0 + orbit.eccentricity*math.Cos(t)))
	M := E - orbit.eccentricity*math.Sin(E)

	// iterate to get M closer to meanAnomaly
	rate := 0.01
	lastDec := false

	for {
		log.Printf("using approx %f to %f\n", M, meanAnomaly)
		if math.Abs(M-meanAnomaly) < 0.0000000000001 {
			break
		}

		if M > meanAnomaly {
			E -= rate
			lastDec = true
		} else {
			E += rate
			if lastDec {
				rate *= 0.1
			}
		}
		M = E - orbit.eccentricity*math.Sin(E)
	}

	if meanAnomaly > math.Pi && E < math.Pi {
		E = 2.0*math.Pi - E
	}

	return E
}

func (orbit *Orbit) calcTrueAnomaly(eccentricAnomaly float64) float64 {
	t := math.Acos((math.Cos(eccentricAnomaly) - orbit.eccentricity) / (1.0 - orbit.eccentricity*math.Cos(eccentricAnomaly)))
	if eccentricAnomaly > math.Pi && t < math.Pi {
		t = 2.0*math.Pi - t
	}
	return t
}

func (orbit *Orbit) Update(delta float64) {
	M := orbit.meanAnomaly()
	M += orbit.meanMotion() * delta

	M = math.Mod(M, 2.0*math.Pi)

	for M < (-2.0 * math.Pi) {
		log.Printf("M: %f < %f", M, -2.0*math.Pi)
		M = M + 2.0*math.Pi
	}

	if M < 0 {
		M = 2.0*math.Pi + M
	}

	for M > (2.0 * math.Pi) {
		log.Printf("M: %f > %f", M, 2.0*math.Pi)
		M = M - 2.0*math.Pi
	}

	E := orbit.calcEccentricAnomaly(M)
	orbit.trueAnomaly = orbit.calcTrueAnomaly(E)
}

func (orbit *Orbit) String() string {
	s := fmt.Sprintf("\n mass: %f\n", orbit.mass)
	s = fmt.Sprintf("\n GM: %f\n", GravitationalConstant*orbit.mass)
	s = fmt.Sprintf("%s semiMajorAxis: %f\n", s, orbit.semiMajorAxis)
	s = fmt.Sprintf("%s eccentricity: %f\n", s, orbit.eccentricity)
	s = fmt.Sprintf("%s longitudeOfAscendingNode: %f\n", s, orbit.longitudeOfAscendingNode)
	s = fmt.Sprintf("%s argumentOfPeriapsis: %f\n", s, orbit.argumentOfPeriapsis)
	s = fmt.Sprintf("%s inclination: %f\n", s, orbit.inclination)
	s = fmt.Sprintf("%s trueAnomaly: %f\n", s, orbit.trueAnomaly)
	s = fmt.Sprintf("%s eccentricAnomaly: %f\n", s, orbit.eccentricAnomaly())
	s = fmt.Sprintf("%s meanAnomaly: %f\n", s, orbit.meanAnomaly())
	s = fmt.Sprintf("%s period: %f\n", s, orbit.period())
	s = fmt.Sprintf("%s meanMotion: %f\n", s, orbit.meanMotion())
	s = fmt.Sprintf("%s epoch: %v\n", s, orbit.epoch)
	s = fmt.Sprintf("%s position: %v\n", s, orbit.Position())
	s = fmt.Sprintf("%s velocity: %v\n", s, orbit.Velocity())
	return s
}
