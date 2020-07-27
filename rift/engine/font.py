"""font

Specifies a font file (png/ttf) data structure for loading.
"""

class TrueType:
    def __init__(self, name: str, path: str):
        self.name = name
        self.path = path

class Bitmap:
    def __init__(self, name: str, path: str, tile_width: int, tile_height: int, charmap: str):
        self.name = name
        self.path = path
        self.tile_width = tile_width
        self.tile_height = tile_height
        self.charmap = charmap
