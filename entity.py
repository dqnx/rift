class Entity:
    """
    A generic object to represent players, enemies, items, etc.
    """
    def __init__(self, x, y, char, color):
        self.x = x
        self.y = y
        self.char = char
        self.color = color
        self.energy = 0
        self.energy_gain = 1
        self.action = None

    def move(self, dx, dy):
        # Move the entity by a given amount
        self.x += dx
        self.y += dy
    
    def get_action(self):
        return self.action

class Hero(Entity):
    def __init__(self, x, y, char, color):
        super().__init__(x, y, char, color)

    def get_action(self):
        action = self.action
        self.action = None
        return action
    
    def set_action(self, action):
        self.action = action
    