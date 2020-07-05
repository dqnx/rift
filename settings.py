"""settings

Defines core game settings singletons.
"""

from singleton import Singleton

class Settings(Singleton):
    def __init__(self):
        super().__init__()
        