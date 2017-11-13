package main

import sf "github.com/manyminds/gosfml"
import "fmt"

const (
	screenWidth  = 800
	screenHeight = 600
)

type PianoKey struct {
	graphic *sf.RectangleShape
	note *sf.Sound
	blackKey bool
	noteName *sf.Text
	keyToPress int
}

func NewPianoKey (n *sf.Sound, nname string, bK bool, key int) *PianoKey {
	p := new (PianoKey)

	p.note = n
	p.blackKey = bK
	p.graphic, _ = sf.NewRectangleShape ()
	p.keyToPress = key
	
	font, _ := sf.NewFontFromFile ("./Notes/trench100free.otf")
	p.noteName, _ = sf.NewText (font)
	p.noteName.SetString (nname)
	p.noteName.SetOrigin (sf.Vector2f {15, 15})

	var sizeX float32
	var sizeY float32

	if bK {
		sizeX = 18
		sizeY = 65
		p.graphic.SetFillColor (sf.ColorBlack())
		p.noteName.SetColor (sf.ColorWhite ())
		p.noteName.SetCharacterSize (25)
	} else {
		sizeX = 20
		sizeY = 120
		p.graphic.SetFillColor (sf.ColorWhite ())
		p.noteName.SetColor (sf.ColorBlack ())
	}

	p.graphic.SetSize (sf.Vector2f{sizeX, sizeY})
	p.graphic.SetOrigin (sf.Vector2f {sizeX, sizeY})

	return p
}

func PlayNote (key sf.KeyCode, wk []*PianoKey, blk []*PianoKey) {
	// TODO: Display keybindings.

	for i := 0; i < len (wk); i++ {
		if key == sf.KeyCode(wk[i].keyToPress) {
			wk[i].note.Play ()
		}
	}

	for i := 0; i < len (blk); i++ {
		if key == sf.KeyCode(blk[i].keyToPress) {
			blk[i].note.Play ()
		}
	}
}


func main() {
	window := sf.NewRenderWindow(sf.VideoMode{screenWidth, screenHeight, 32}, "Piano", sf.StyleDefault, sf.DefaultContextSettings())
	window.SetVSyncEnabled(true)
	window.SetFramerateLimit(60)

	whiteKeys := make ([]*PianoKey, 0)
	blackKeys := make ([]*PianoKey, 0)

	//Prepare yourself for a multitude of disgusting for loops

	// White Key Sounds
	for i := 0; i <= 6; i ++ {
		var snd *sf.Sound
		var sndBuffer *sf.SoundBuffer
		var key int
		var name string

		switch i {
		case 0:
			sndBuffer, _ = sf.NewSoundBufferFromFile ("./Notes/g.wav")
			name = "G"
			key = sf.KeyA
			break
		case 1:
			sndBuffer, _ = sf.NewSoundBufferFromFile ("./Notes/a.wav")
			name = "A"
			key = sf.KeyS
			break
		case 2:
			sndBuffer, _ = sf.NewSoundBufferFromFile ("./Notes/b.wav")
			name = "B"
			key = sf.KeyD
			break
		case 3:
			sndBuffer, _ = sf.NewSoundBufferFromFile ("./Notes/c.wav")
			name = "C"
			key = sf.KeyF
			break
		case 4:
			sndBuffer, _ = sf.NewSoundBufferFromFile ("./Notes/d.wav")
			name = "D"
			key = sf.KeyG
			break
		case 5:
			sndBuffer, _ = sf.NewSoundBufferFromFile ("./Notes/e.wav")
			name = "E"
			key = sf.KeyH
			break
		case 6:	
			sndBuffer, _ = sf.NewSoundBufferFromFile ("./Notes/f.wav")
			name = "F"
			key = sf.KeyJ
			break
		default:
			fmt.Println ("NO SOUND FILE")
			return
		}

		snd = sf.NewSound (sndBuffer)
		whiteKeys = append (whiteKeys, NewPianoKey (snd, name, false, key))
	}

	// Black Key Sounds
	for i := 0; i <= 4; i ++ {
		var snd *sf.Sound
		var sndBuffer *sf.SoundBuffer
		var key int
		var name string

		switch i {
		case 0:
			sndBuffer, _ = sf.NewSoundBufferFromFile ("./Notes/gsharp.wav")
			name = "G\n#"
			key = sf.KeyW
			break
		case 1:
			sndBuffer, _ = sf.NewSoundBufferFromFile ("./Notes/fsharp.wav")
			name = "F\n#"
			key = sf.KeyI
			break
		case 2:
			sndBuffer, _ = sf.NewSoundBufferFromFile ("./Notes/eflat.wav")
			name = "E\nb"
			key = sf.KeyY
			break
		case 3:
			sndBuffer, _ = sf.NewSoundBufferFromFile ("./Notes/csharp.wav")
			name = "C\n#"
			key = sf.KeyT
			break
		case 4:
			sndBuffer, _ = sf.NewSoundBufferFromFile ("./Notes/bflat.wav")
			name = "B\nb"
			key = sf.KeyE
			break
		default:
			fmt.Println ("NO SOUND FILE")
			return
		}

		snd = sf.NewSound (sndBuffer)
		blackKeys = append (blackKeys, NewPianoKey (snd, name, true, key))
	}

	fmt.Println (len(whiteKeys))

	fmt.Println (len(blackKeys))

	// White Keys Drawn
	for i := 0; i < len (whiteKeys); i ++ {
		placeX := float32 (screenWidth/2 - 60.5)

		if i == 0 {
			whiteKeys[i].graphic.SetPosition (sf.Vector2f {placeX, screenHeight/2})
		} else {
			whiteKeys[i].graphic.SetPosition (sf.Vector2f {whiteKeys [i-1].graphic.GetPosition ().X + whiteKeys[i-1].graphic.GetSize ().X + 1, screenHeight/2})
		} 
	}

	for i := 0; i < len (blackKeys); i ++ {
		switch i {
			case 0:
				blackKeys[i].graphic.SetPosition (sf.Vector2f {whiteKeys[0].graphic.GetPosition ().X + whiteKeys[0].graphic.GetSize ().X/2, whiteKeys[0].graphic.GetPosition ().Y - whiteKeys[0].graphic.GetSize ().Y/2})
				break
			case 1:
				blackKeys[i].graphic.SetPosition (sf.Vector2f {whiteKeys[6].graphic.GetPosition ().X + whiteKeys[6].graphic.GetSize ().X/2, whiteKeys[6].graphic.GetPosition ().Y - whiteKeys[6].graphic.GetSize ().Y/2})
				break
			case 2:
				blackKeys[i].graphic.SetPosition (sf.Vector2f {whiteKeys[4].graphic.GetPosition ().X + whiteKeys[4].graphic.GetSize ().X/2, whiteKeys[4].graphic.GetPosition ().Y - whiteKeys[4].graphic.GetSize ().Y/2})
				break
			case 3:
				blackKeys[i].graphic.SetPosition (sf.Vector2f {whiteKeys[3].graphic.GetPosition ().X + whiteKeys[3].graphic.GetSize ().X/2, whiteKeys[3].graphic.GetPosition ().Y - whiteKeys[3].graphic.GetSize ().Y/2})
				break
			case 4:
				blackKeys[i].graphic.SetPosition (sf.Vector2f {whiteKeys[1].graphic.GetPosition ().X + whiteKeys[1].graphic.GetSize ().X/2, whiteKeys[1].graphic.GetPosition ().Y - whiteKeys[0].graphic.GetSize ().Y/2})
				break
			default:
				fmt.Println ("UNRECOGNIZED BLACK KEY")
				return
		}
	}

	for window.IsOpen() {

		for event := window.PollEvent(); event != nil; event = window.PollEvent() {
			switch ev := event.(type) {
			case sf.EventKeyReleased:
				if ev.Code == sf.KeyEscape {
					window.Close()
				} else {
					PlayNote (ev.Code, whiteKeys, blackKeys)
				}
				break
			case sf.EventClosed:
				window.Close()
				break
			}
		}




		window.Clear(sf.Color {0, 170, 224, 255})
		
		// TODO: Fix the "Failed to activate window's context" error which stops the program
		// TODO: Find the source of aforementioned error

		for i := 0; i < len (whiteKeys); i ++ {
			whiteKeys[i].noteName.SetPosition (sf.Vector2f {whiteKeys[i].graphic.GetPosition().X, whiteKeys[i].graphic.GetPosition().Y - 20})	
			window.Draw (whiteKeys[i].graphic, sf.DefaultRenderStates ())
			window.Draw (whiteKeys[i].noteName, sf.DefaultRenderStates ())
		}

		for i := 0; i < len (blackKeys); i ++ {
			blackKeys[i].noteName.SetPosition (sf.Vector2f {blackKeys[i].graphic.GetPosition().X, blackKeys[i].graphic.GetPosition().Y - 40})			
			window.Draw (blackKeys[i].graphic, sf.DefaultRenderStates ())
			window.Draw (blackKeys[i].noteName, sf.DefaultRenderStates ())

		}

		window.Display()
	}
}