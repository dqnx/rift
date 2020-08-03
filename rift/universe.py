"""universe

Defines game location classes."""

import math
import pandas as pd
import engine.coordinate as coord

# Constants

# Newtons method iterations
ECC_ANOM_ITERS = 30
# Gravitational constant
GM = 6.67430e-11 

def _keplers_eq(e: float, M: float, hist=False, iters=ECC_ANOM_ITERS) -> (float, list):
    """Calculate eccentric anomaly from eccentricity and mean anomaly using Newton's method."""
    E_hist = []
    # Newtons method of En+1 = E - F(E)/F'(E)
    def E_next(E: float) -> float:
        n = E - e*math.sin(E) - M
        d = 1.0 - e*math.cos(E)
        return E - n/d
    
    # initialize E_0
    if e > 0.8:
        E_n = math.pi
    else:
        E_n = M

    for _ in range(iters):
        if hist: E_hist.append(E_n) # debugging
        E_n = E_next(E_n)
        
    if hist: E_hist.append(E_n) # debugging
    
    return (E_n, E_hist)

class Body:
    """Body
    Stellar object base class. A stationary object in a relative position around another central object.
    All positions are in spherical coordinates."""
    def __init__(self, name: str, mass: float, position: coord.Cartesian, parent = None):
        self.name = name
        self.parent = parent
        self.mass = mass
        self._position = position

    @property
    def position(self):
        return self._position

class OrbitalBody(Body):
    """OrbitalBody
    Stellar object with orbital mechanics."""
    def __init__(self, name: str, mass: float, eccentricity: float, semi_major_axis: float, inclination: float, 
        la_node: float, arg_periapsis: float, parent: Body = None):
        super().__init__(name, mass, None, parent=parent)
        ## Kepler orbit elements
        # Elliptical orbit eccentricity
        self._eccentricity = eccentricity
        # Orbit semi-major axis length
        self._semi_major_axis = semi_major_axis
        # Longitude of the ascending node
        self._la_node = la_node
        # Inclination
        self._inclination = inclination
        # Argument of periapsis
        self._arg_periapsis = arg_periapsis

        ## Motion states
        self._true_anomaly = None
        self._radius = None
        self._sigma = None

        # Orbit position state: time since perphelion
        self._time = 0.0

        # Position state. If orbital params are changed, position must be recalculated.
        # This is expensive, so only calculate when params change.
        self._current_pos = True
        self._update_position()

    def update(self, dt: float):
        """Advances the body along its orbit by dt (seconds).
        Solves Kepler's Equation for finding radius and true anomaly."""
        # Calculate mean motion and mean anomaly
        self._time += dt
        n = 2*math.pi/self.period
        M = n*self._time

        # Calculate eccentric anomaly
        E = _keplers_eq(self.eccentricity, M)

        # Calculate true anomaly and radius
        self._true_anomaly = 2*math.atan(math.sqrt((1+self._eccentricity)/(1-self._eccentricity))*math.tan(E/2))
        self._radius = self._semi_major_axis*(1-math.pow(self.eccentricity,2))/(1+self._eccentricity*math.cos(self._true_anomaly))
        self._sigma = math.asin(math.sin(self._inclination)*math.sin(self._arg_periapsis+self._true_anomaly))

        # Reset time along the period.
        T = self.period()
        while self._time > 2*T:
            self._time -= T

    def _update_position(self):
        """Calculates location in 3D space relative to central body."""
        self._current_pos = True

        x = self._radius * math.cos(self._arg_periapsis) * math.cos(self._sigma)
        y = self._radius * math.sin(self._arg_periapsis) * math.cos(self._sigma)
        z = self._radius * math.sin(self._sigma)

        self._position = coord.Cartesian(x,y,z)

    @property
    def position(self):
        # Check if position is current
        if not self._current_pos:
            self._update_position()
        return self._position
    
    @property
    def true_anomaly(self):
        return self._true_anomaly
    
    @property
    def radius(self):
        return self._radius

    @property
    def eccentricity(self):
        return self._eccentricity
    
    @eccentricity.setter
    def eccentricity(self, e):
        self._eccentricity = e
        self._current_orbit = False

    @property
    def semi_major_axis(self):
        return self._semi_major_axis
    
    @semi_major_axis.setter
    def semi_major_axis(self, e):
        self._semi_major_axis = e
        self._current_orbit = False
    
    @property
    def la_node(self):
        return self._la_node
    
    @la_node.setter
    def la_node(self, e):
        self._la_node = e
        self._current_orbit = False
        
    @property
    def inclination(self):
        return self._inclination
    
    @inclination.setter
    def inclination(self, e):
        self._inclination = e
        self._current_orbit = False

    @property
    def arg_periapsis(self):
        return self._arg_periapsis
    
    @arg_periapsis.setter
    def arg_periapsis(self, e):
        self._arg_periapsis = e
        self._current_orbit = False
    
    @property
    def speed(self) -> float:
        v = math.sqrt(self.parent.grav_param*(2/self.position.r - 1/self.semi_major_axis))
        return v
    
    @property
    def period(self) -> float:
        T = 2*math.pi*math.sqrt(math.pow(self.semi_major_axis, 3)/self.parent.grav_param)
        return T
    
    @property
    def total_energy(self) -> float:
        E = -self.parent.grav_param*self.mass / 2*self.semi_major_axis
        return E
    
    @property
    def grav_param(self) -> float:
        return GM*self.mass


class System:
    """System
    System of bodies in an orbital configuration."""
    def __init__(self):
        self._bodies = []

class Galaxy:
    """Galaxy
    Arrangement of systems."""
    def __init__(self):
        self._systems = pd.DataFrame()
