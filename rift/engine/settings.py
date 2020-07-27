"""settings

Defines core game settings singletons.
"""

from bearlibterminal.terminal import color_from_argb

from engine.singleton import singleton
import engine.font

@singleton
class Settings:
    def initialize(self):
        """Load settings values. TODO: move to settings yaml file."""
        self._screen_width = 80
        self._screen_height = 50
        self._map_width = 80
        self._map_height = 45
        self._tile_width = 10
        self._tile_height = 10

        self._default_grid_font = 'dejavu_mono'
        self._default_text_font = 'dejavu_sans'

        self._asset_dir = "assets"
        self._fonts = {
            'arial_10x10': engine.font.Bitmap('arial_10x10', self._asset_dir+"/arial10x10.png", 10, 10, '437'),
            'dejavu_mono': engine.font.TrueType('dejavu_mono', self._asset_dir+"/DejaVuSansMono.ttf"),
            'dejavu_sans': engine.font.TrueType('dejavu_sans', self._asset_dir+"/DejaVuSans.ttf")
        }

        self._colors = {
            'default_fg': color_from_argb(255,255,255,255),
            'default_bg': color_from_argb(255,0,0,0),
            'highlight_fg': color_from_argb(255,102,153,255),
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
    


        
        