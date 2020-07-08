from abc import ABC, abstractmethod
from engine.game import Direction

class Action(ABC):
    @abstractmethod
    def perform(self, o: object):
        pass

class WalkAction(Action):
    def __init__(self, direction):
        self.direction = direction

    def perform(self, actor):
        if self.direction == Direction.NORTH:
            actor.move(0,1)
        elif self.direction == Direction.SOUTH:
            actor.move(0,-1)
        elif self.direction == Direction.EAST:
            actor.move(1,0)
        elif self.direction == Direction.WEST:
            actor.move(-1,0)

