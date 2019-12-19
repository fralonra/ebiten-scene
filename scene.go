package scene

import (
	"github.com/hajimehoshi/ebiten"
)

var (
	currentScene GameScene
	scenes       = make(map[GameSceneKey]*GameScene)
)

// Scene returns the current scene.
func Scene() GameScene {
	return currentScene
}

// Set registers a GameScene with the given key.
func Set(key GameSceneKey, scene GameScene) {
	scenes[key] = &scene
}

// SwitchWith changes the current scene to the scene with the given key.
// It also receives a function which allows you to access the new scene.
func SwitchWith(key GameSceneKey, withFunc func(GameScene)) {
	scene := scenes[key]
	if scene == nil {
		return
	}
	if currentScene != nil {
		currentScene.Stop()
	}

	if withFunc != nil {
		withFunc(*scene)
	}

	currentScene = *scene
	currentScene.Start()
}

// Switch changes the current scene to the scene with the given key.
func Switch(key GameSceneKey) {
	SwitchWith(key, nil)
}

// Update is the function called every frame by ebiten.
func Update(screen *ebiten.Image) error {
	if currentScene == nil {
		return nil
	}
	currentScene.Update()

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	currentScene.Draw(screen)
	return nil
}

// GameScene is the interface that all the game scene should implement.
//
// Draw is the function where rendering takes place.
//
// Update is the function where game logic updates. It is invoked before Draw.
//
// Start is the function called every time you switch to the scene.
//
// Stop is the function called every time you exit the scene.
type GameScene interface {
	Draw(screen *ebiten.Image)
	Update()
	Start()
	Stop()
}

// GameSceneKey represents the identifier of each scene.
type GameSceneKey int
