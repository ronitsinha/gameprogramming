package main

import (
	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

func Wrap(a *sf.Sprite) {
	x := a.GetPosition ().X
	y := a.GetPosition ().Y

	width := a.GetGlobalBounds ().Width
	height := a.GetGlobalBounds ().Height
	
	if x < -width/2 {
		a.SetPosition (sf.Vector2f {screenWidth + width/2, y})
	} else if x > screenWidth + width/2 {
		a.SetPosition (sf.Vector2f {-width/2, y})
	}

	if y < -height/2 {
		a.SetPosition (sf.Vector2f {x,  screenHeight + height/2})
	} else if y > screenHeight + height/2 {
		a.SetPosition (sf.Vector2f {x, -height/2})
	}
}

func Intersects(s1, s2 *sf.Sprite) bool {
	returnBool, _ := s1.GetGlobalBounds ().Intersects (s2.GetGlobalBounds ())

	return returnBool
}