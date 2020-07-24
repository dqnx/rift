"""universe

Defines game location classes."""

import pandas as pd

class Body:
    """Body
    Stellar body base class."""
    def __init__(self):
        self.name = ""

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
