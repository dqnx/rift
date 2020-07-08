"""font

Specifies a font file (png/ttf) data structure for loading in tcod.
"""

from enum import Enum
from abc import ABC, abstractmethod
import tcod

class Font(ABC):
    def __init__(self, name: str, path: str):
        self._name = name
        self._path = path
    
    @property
    def name(self):
        return self._name

    @abstractmethod
    def tileset(self, tile_width: int, tile_height: int) -> tcod.tileset.Tileset:
        pass

class TTFFont(Font):
    def __init__(self, name: str, path: str):
        super().__init__(name, path)
    
    def tileset(self, tile_width: int, tile_height: int) -> tcod.tileset.Tileset:
        return tcod.tileset.load_truetype_font(self._path, tile_width, tile_height)
        
class PNGFont(Font):
    def __init__(self, name: str, path: str, columns: int, rows: int, tile_width: int, tile_height: int, charmap):
        super().__init__(name, path)
        self._columns = columns
        self._rows = rows
        self._charmap = charmap
        self._tile_width = tile_width
        self._tile_height = tile_height
    
    @property
    def tile_size(self):
        return (self._tile_width, self._tile_height)
    
    def tileset(self, tile_width: int, tile_height: int) -> tcod.tileset.Tileset:
        return tcod.tileset.load_tilesheet(self._path, self._columns, self._rows, self._charmap)
    
