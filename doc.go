/*
	Package scene provides an easy to use scene manager for ebiten.

		package main

		import (
			scene "github.com/fralonra/ebiten-scene"
			"github.com/hajimehoshi/ebiten"
		)

		// implement interface GameScene
		type SceneA struct{}

		func (s *SceneA) Draw(screen *ebiten.Image) {
			// Draw is the function where rendering takes place.
		}

		func (s *SceneA) Update() {
			// Update is the function where game logic updates. It is invoked before Draw.
		}

		func (s *SceneA) Start() {
			// Start is the function called every time you switch to the scene.
		}

		func (s *SceneA) Stop() {
			// Stop is the function called every time you exit the scene.
		}

		func main() {
			scene.Set(0, &SceneA{})
			scene.Switch(0)
			if err := ebiten.Run(scene.Update, 320, 240, 2, "Your game's title"); err != nil {
				log.Fatal(err)
			}
		}
*/
package scene
