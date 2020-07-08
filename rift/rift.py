import tcod

from engine.gamestate import GameStateMachine
from settings import Settings
from mainmenu import MenuState

def main():
    # Initialize global settings.
    sets = Settings()
    sets.initialize()

    # Load the font and create a tileset.
    tileset = sets.font().tileset(*sets.tile_size)

    # Create the game-scene state machine and initialize with the main menu.
    game_state = GameStateMachine(MenuState())

    # Create a window based on this console and tileset.
    console = tcod.Console(*sets.screen_size)
    with tcod.context.new_terminal(
        console.width, console.height, tileset=tileset,
    ) as context:
        while True: 
            # If it exists, process user input(s).
            for event in tcod.event.get():
                # If the game state is over/exited, it will empty the state machine and return false.
                if not game_state.run(event):
                    raise SystemExit

            # Render a frame with a clear, draw, frame-swap operation.
            console.clear()
            game_state.render(console)
            context.present(console)
                
if __name__ == '__main__':
    main()