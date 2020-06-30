import tcod
import tcod.event
from game import Game, Direction
from action import WalkAction

def handle_event(event, game):
    if event.type == "QUIT":
        raise SystemExit()
    elif event.type == "KEYDOWN":
        handle_hero(event, game)

def handle_hero(event, game):
    if event.sym == tcod.event.K_UP:
        game.player.set_action(WalkAction(Direction.NORTH))
    elif event.sym == tcod.event.K_DOWN:
        game.player.set_action(WalkAction(Direction.SOUTH))
    elif event.sym == tcod.event.K_RIGHT:
        game.player.set_action(WalkAction(Direction.EAST))
    elif event.sym == tcod.event.K_LEFT:
        game.player.set_action(WalkAction(Direction.WEST))
