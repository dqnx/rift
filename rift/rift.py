from bearlibterminal import terminal

from engine.gamestate import GameStateMachine
from engine.settings import Settings
from mainmenu import MenuState

def main():
    # Initialize global settings.
    sets = Settings()
    sets.initialize()

    terminal.open()

    # Load assets
    f = sets.font('dejavu_sans')
    if not terminal.set("font: %s, size=%d" % (f.path, sets.tile_size[0])):
        print("ERROR: Failed to set font.")
        raise SystemError
    if not terminal.set("window: title=%s, size=%dx%d" % (sets.title, *sets.screen_size)):
        print("ERROR: Failed to set screen size.")
        raise SystemError

    # Flag to prevent drawing before state update
    game_state_change = False

    # Create the game-scene state machine and initialize with the main menu.
    game_state = GameStateMachine(MenuState())
    game_state.render() 
    terminal.refresh()

    # Update/render loop
    while True: 
        # Wait for an event (animation, user input).
        while terminal.has_input():
            game_state_change = True
            # If the game state is over/exited, it will empty the state machine and return false.
            if not game_state.run(terminal.read()):
                terminal.close()
                raise SystemExit
            
        # Render a frame with a clear, draw, frame-swap operation.
        if (game_state_change):
            terminal.clear()
            game_state.render()
            terminal.refresh()

            # Reset game state change for renedering
            game_state_change = False

if __name__ == '__main__':
    main()