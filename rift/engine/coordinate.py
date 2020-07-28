"""coordinate

Defines 2D and 3D coordinate system datatypes."""

class Spherical:
    def __init__(self, radius: float = 0.0, inclination: float = 0.0, azimuthal: float = 0.0):
        self.radius = radius
        self.inclination = inclination
        self.azimuthal = azimuthal

    #short hand accessors

    @property
    def r(self):
        return self.radius
    @property
    def theta(self):
        return self.inclination
    @property
    def phi(self):
        return self.azimuthal

class Cartesian:
    def __init__(self, x: float, y: float, z: float):
        self.x = x
        self.y = y
        self.z = z
    
    
    
    