from map_objects.game_map import GameMap
from entity import Entity, Hero
from enum import Enum

class Game:
    def __init__(self):
        self.map_width = 80
        self.map_height = 45
        self.game_map = GameMap(self.map_width, self.map_height)

        self.player = Hero(int(self.map_width / 2), int(self.map_height / 2), '@', 'player')
        self.entities = [self.player]
        self.entities.append(Entity(int(self.map_width / 2 - 5), int(self.map_height / 2), '@', 'npc'))
        self.current_entity = 0

        self.energy_threshold = 5

    def process(self):
        if self.entities[self.current_entity].energy > self.energy_threshold:
            action = self.entities[self.current_entity].get_action()
            if action != None:
                action.perform(self.entities[self.current_entity])
            self.entities[self.current_entity].energy = 0
            
        self.current_entity = (self.current_entity + 1) % len(self.entities)

class Direction(Enum):
    NORTH = 0
    SOUTH = 1
    EAST = 2
    WEST = 3
