"""settings

Defines core game settings singletons.
"""

import tcod
from singleton import singleton

@singleton
class Settings:
    def initialize(self):
        """Load settings values. TODO: move to settings yaml file."""
        self._screen_width = 80
        self._screen_height = 50
        self._map_width = 80
        self._map_height = 45

        self.default_font_name = 'arial10'
        self.fonts = {
            'arial10': 'arial10x10.png'
        }

        self._colors = {
            'dark_wall': tcod.Color(0, 0, 100),
            'dark_ground': tcod.Color(50, 50, 150),
            'player': tcod.white,
            'npc': tcod.yellow,
            'default_fg': (255,255,255),
            'default_bg': (0,0,0)
        }

        self._title = "rift"

        # Run tcod init to apply settings.
        tcod.console_init_root(self._screen_width, self._screen_height, self._title, False)
        tcod.console_set_custom_font(self.fonts[self.default_font_name], tcod.FONT_TYPE_GREYSCALE | tcod.FONT_LAYOUT_TCOD)
    
    @property
    def screen_size(self) -> (int, int):
        """Gets game screen size as (width, height)."""
        return (self._screen_width, self._screen_height)
    
    @screen_size.setter
    def set_screen_size(self, w: int, h: int):
        """Sets screen size as width and height."""
        self._screen_width = w
        self._screen_height = h
        tcod.console_init_root(self._screen_width, self._screen_height, self._title, False)

    @property
    def map_size(self) -> (int, int):
        """Gets game map size as (width, height)."""
        return (self._map_width, self._map_height)
    
    @map_size.setter
    def set_map_size(self, w: int, h: int):
        """Sets map size as width and height."""
        self._map_width = w
        self._map_height = h
    
    def set_font(self, font_name: str):
        if self.fonts[font_name] is None:
            raise RuntimeWarning("set_font cannot set empty font name: " + font_name)

        tcod.console_set_custom_font(self.fonts[font_name], tcod.FONT_TYPE_GREYSCALE | tcod.FONT_LAYOUT_TCOD)

    @property
    def colors(self):
        return self._colors
    
    @property
    def title(self):
        return self._title
    


        
        