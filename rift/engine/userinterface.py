"""userinterface

Provides a user interface base class and specific gameplay interfaces/menus.
"""

from abc import ABC, abstractmethod
from enum import Enum
import tcod
from settings import Settings

class Direction(Enum):
    LEFT = 0
    RIGHT = 1
    UP = 2
    DOWN = 3

class UserInterface(ABC):
    @abstractmethod
    def render(self, console: tcod.console.Console):
        pass
