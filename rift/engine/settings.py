"""settings

Defines core game settings singletons.
"""

import tcod
from engine.singleton import singleton
from engine.font import PNGFont, TTFFont

@singleton
class Settings:
    def initialize(self):
        """Load settings values. TODO: move to settings yaml file."""
        self._screen_width = 80
        self._screen_height = 50
        self._map_width = 80
        self._map_height = 45
        self._tile_width = 22
        self._tile_height = 22

        self._default_grid_font = 'dejavu_mono'
        self._default_text_font = 'dejavu_sans'

        self._asset_dir = "assets"
        self._fonts = {
            'arial_10x10': PNGFont('arial_10x10', self._asset_dir+"/arial10x10.png", 32, 8, 10, 10, tcod.tileset.CHARMAP_TCOD),
            'dejavu_mono': TTFFont('dejavu_mono', self._asset_dir+"/DejaVuSansMono.ttf"),
            'dejavu_sans': TTFFont('dejavu_sans', self._asset_dir+"/DejaVuSans.ttf")
        }

        self._colors = {
            'dark_wall': tcod.Color(0, 0, 100),
            'dark_ground': tcod.Color(50, 50, 150),
            'player': tcod.white,
            'npc': tcod.yellow,
            'default_fg': (255,255,255),
            'default_bg': (0,0,0),
            'highlight_fg': (102, 153, 255),
        }

        self._title = "RIFT"

    
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
    
    @property
    def tile_size(self) -> (int, int):
        """Gets game tile size as (width, height)."""
        return (self._tile_width, self._tile_height)
    
    @tile_size.setter
    def set_tile_size(self, w: int, h: int):
        """Sets tile size as width and height."""
        self._tile_width = w
        self._tile_height = h
    

    def font(self, name=None):
        if name is None:
            return self._fonts[self._default_grid_font]
            
        return self._fonts[name]

    def set_font(self, font_name: str):
        raise RuntimeWarning("set_font not implemented")

    @property
    def colors(self):
        return self._colors
    
    @property
    def title(self):
        return self._title
    


        
        