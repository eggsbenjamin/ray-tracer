package game

type Weapon interface {
	Fire()
	Reload()
	Update()
	Render()
}

type WeaponState struct {
	Sprite
}

type weapon struct {
	currentState         string
	states               map[string]*WeaponState
	clip, clipSize, ammo int
}

func NewWeapon(clip, clipSize, ammo int) Weapon {
	return &weapon{
		clip:     clip,
		clipSize: clipSize,
		ammo:     ammo,
	}
}

func (w *weapon) Update() {
	w.states[w.currentState].Sprite.Update()
}

func (w *weapon) Render() {

}

func (w *weapon) Fire() {
	// TODO: animate
	if w.clip > 0 {
		w.currentState = "firing"
		w.clip--
	}
}

func (w *weapon) Reload() {
	// TODO: animate
	if w.ammo > 0 && w.currentState != "reloading" {
		w.currentState = "reloading"

		if w.ammo > w.clipSize {
			d := w.clipSize - w.clip
			w.clip += d
			w.ammo -= d
		}

		d := w.ammo - w.clip
		w.clip += d
		w.ammo -= d
	}
}
