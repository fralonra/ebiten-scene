package scene

import (
	"github.com/hajimehoshi/ebiten"
)

var (
	currentScene GameScene
	scenes       = make(map[GameSceneKey]*GameScene)
)

func Scene() GameScene {
	return currentScene
}

func Set(key GameSceneKey, scene GameScene) {
	scenes[key] = &scene
}

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

func Switch(key GameSceneKey) {
	SwitchWith(key, nil)
}

type GameScene interface {
	Draw(screen *ebiten.Image)
	Update()
	Start()
	Stop()
}

type GameSceneKey int
