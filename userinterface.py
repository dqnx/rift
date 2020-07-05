"""userinterface

Provides a user interface base class and specific gameplay interfaces/menus.
"""

from abc import ABC, abstractmethod
from tcod import console

class UserInterface(ABC):
    @abstractmethod
    def render(self, con: console.Console):
        pass

class MainMenuUI(UserInterface):
    def __init__(self):
        super().__init__()

        self.selection = 0
        self.menu_items = ['Start', 'Options', 'Exit']

    def render(self, con, screen_width, screen_height, colors):
        pass

