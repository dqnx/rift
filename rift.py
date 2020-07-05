import tcod
import tcod.event
from engine.input_handlers import InputHandler
from render_functions import clear_all, render_all
from map_objects.game_map import GameMap
from entity import Entity
from settings import Settings

def main():

    sets = Settings()
    sets.initialize()

    map_width, map_height = sets.map_size
    screen_width, screen_height = sets.screen_size

    player = Entity(int(screen_width / 2), int(screen_height / 2), '@', tcod.white)
    npc = Entity(int(screen_width / 2 - 5), int(screen_height / 2), '@', tcod.yellow)
    entities = [npc, player]

    con = tcod.console.Console(screen_width, screen_height)
    
    game_map = GameMap(map_width, map_height)

    """
    root_console = tcod.console_init_root(80, 60)
    state = State()
    while True:
    tcod.console_flush()
    for event in tcod.event.wait():
        state.dispatch(event)
    """
    state = InputHandler()

    console = tcod.Console(screen_width, screen_height)
    # Create a window based on this console and tileset.
    with tcod.context.new_terminal(
        console.width, console.height, tileset=tileset,
    ) as context:
        while True:  # Main loop, runs until SystemExit is raised.
            for event in tcod.event.get():
                action = state.dispatch(event)
                if action != None:
                    cmd, val = action
                    if cmd == 'move':
                        dx, dy = val
                        if not game_map.is_blocked(player.x + dx, player.y + dy):
                            player.move(dx, dy)

                    if cmd == 'exit':
                        raise SystemExit

            clear_all(console, entities)
            render_all(console, entities, game_map, screen_width, screen_height, sets.colors)
            tcod.console_flush()
        
if __name__ == '__main__':
    main()