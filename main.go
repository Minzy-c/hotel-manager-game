package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

const (
	screenWidth  = 1280
	screenHeight = 720
	tileSize     = 32
)

// Game states
const (
	GameStateMainMenu = iota
	GameStatePlaying
	GameStatePaused
)

// Game main structure
type Game struct {
	state       int
	assets      *AssetManager
	gameData    *GameData
	hotelMap    [][]int
	rooms       []Room
	guests      []Guest
	ui          *UI
	lastUpdate  time.Time
	mouseX      int
	mouseY      int
	clicked     bool
}

// AssetManager asset management
type AssetManager struct {
	images map[string]*ebiten.Image
	loaded bool
}

// GameData game data
type GameData struct {
	Money        int
	HotelLevel   int
	GuestCount   int
	MaxGuests    int
	Satisfaction int
	Day          int
	HotelName    string
}

// Room room structure
type Room struct {
	ID       int
	X, Y     int
	Width    int
	Height   int
	Type     string
	Occupied bool
	Price    int
	GuestID  int
}

// Guest guest structure
type Guest struct {
	ID           int
	X, Y         float64
	Satisfaction int
	StayDuration int
	RoomID       int
	Name         string
	CharacterID  int
	CheckInTime  time.Time
}

// UI user interface
type UI struct {
	panels map[string]Panel
}

// Panel UI panel
type Panel struct {
	X, Y, Width, Height int
	Title               string
	Buttons             []Button
}

// Button button
type Button struct {
	X, Y, Width, Height int
	Text                string
	Action              func()
	Enabled             bool
}

// NewGame create new game
func NewGame() *Game {
	game := &Game{
		state: GameStateMainMenu,
		assets: &AssetManager{
			images: make(map[string]*ebiten.Image),
		},
		gameData: &GameData{
			Money:        10000,
			HotelLevel:   1,
			GuestCount:   0,
			MaxGuests:    5,
			Satisfaction: 0,
			Day:          1,
			HotelName:    "New Hotel",
		},
		hotelMap: make([][]int, 15),
		rooms:    make([]Room, 0),
		guests:   make([]Guest, 0),
		ui:       &UI{panels: make(map[string]Panel)},
	}

	// Initialize hotel map
	game.initHotelMap()
	game.createInitialRooms()
	game.initUI()
	game.loadAssets()

	return game
}

// loadAssets load game assets
func (g *Game) loadAssets() {
	// Load interior assets
	interiorPath := "assets/1_Interiors/32x32/Interiors_32x32.png"
	if img, _, err := ebitenutil.NewImageFromFile(interiorPath); err == nil {
		g.assets.images["interiors"] = img
	} else {
		log.Printf("Failed to load interiors: %v", err)
	}

	// Load character assets
	for i := 1; i <= 20; i++ {
		characterPath := fmt.Sprintf("assets/2_Characters/Character_Generator/0_Premade_Characters/32x32/Premade_Character_32x32_%02d.png", i)
		if img, _, err := ebitenutil.NewImageFromFile(characterPath); err == nil {
			g.assets.images[fmt.Sprintf("character_%d", i)] = img
		} else {
			log.Printf("Failed to load character %d: %v", i, err)
		}
	}

	// Load UI assets
	uiPath := "assets/4_User_Interface_Elements/UI_32x32.png"
	if img, _, err := ebitenutil.NewImageFromFile(uiPath); err == nil {
		g.assets.images["ui"] = img
	} else {
		log.Printf("Failed to load UI: %v", err)
	}

	g.assets.loaded = true
}

// initHotelMap initialize hotel map
func (g *Game) initHotelMap() {
	for y := 0; y < 15; y++ {
		g.hotelMap[y] = make([]int, 20)
		for x := 0; x < 20; x++ {
			if y == 0 || y == 14 || x == 0 || x == 19 {
				g.hotelMap[y][x] = 1 // Wall
			} else {
				g.hotelMap[y][x] = 0 // Floor
			}
		}
	}
}

// createInitialRooms create initial rooms
func (g *Game) createInitialRooms() {
	g.rooms = []Room{
		{ID: 1, X: 2, Y: 2, Width: 4, Height: 3, Type: "single", Occupied: false, Price: 50},
		{ID: 2, X: 8, Y: 2, Width: 4, Height: 3, Type: "single", Occupied: false, Price: 50},
		{ID: 3, X: 14, Y: 2, Width: 4, Height: 3, Type: "double", Occupied: false, Price: 80},
	}
}

// initUI initialize user interface
func (g *Game) initUI() {
	// Information panel
	g.ui.panels["info"] = Panel{
		X: 1000, Y: 50, Width: 250, Height: 200,
		Title: "HOTEL INFORMATION",
	}

	// Menu panel - only New Room and Receive Guest
	g.ui.panels["menu"] = Panel{
		X: 1000, Y: 270, Width: 250, Height: 100,
		Title: "MENU",
		Buttons: []Button{
			{X: 10, Y: 20, Width: 230, Height: 30, Text: "New Room", Action: g.addNewRoom, Enabled: true},
			{X: 10, Y: 60, Width: 230, Height: 30, Text: "Receive Guest", Action: g.addNewGuest, Enabled: true},
		},
	}
}

// Update game update
func (g *Game) Update() error {
	now := time.Now()
	if g.lastUpdate.IsZero() {
		g.lastUpdate = now
	}

	// Get mouse position
	g.mouseX, g.mouseY = ebiten.CursorPosition()

	// Input handling
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		if g.state == GameStatePlaying {
			g.state = GameStatePaused
		} else if g.state == GameStatePaused {
			g.state = GameStatePlaying
		}
	}

	// Mouse click handling
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && !g.clicked {
		g.clicked = true
		g.handleMouseClick(g.mouseX, g.mouseY)
	} else if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.clicked = false
	}

	// Update guests
	if g.state == GameStatePlaying {
		g.updateGuests()
	}

	g.lastUpdate = now
	return nil
}

// handleMouseClick handle mouse click
func (g *Game) handleMouseClick(x, y int) {
	switch g.state {
	case GameStateMainMenu:
		g.handleMainMenuClick(x, y)
	case GameStatePlaying:
		g.handleGameClick(x, y)
	case GameStatePaused:
		g.handlePauseClick(x, y)
	}
}

// handleMainMenuClick handle main menu click
func (g *Game) handleMainMenuClick(x, y int) {
	// New Game button
	if x >= 540 && x <= 740 && y >= 280 && y <= 330 {
		g.state = GameStatePlaying
	}
	// Exit button
	if x >= 540 && x <= 740 && y >= 500 && y <= 550 {
		os.Exit(0)
	}
}

// handleGameClick handle game click
func (g *Game) handleGameClick(x, y int) {
	// UI panel clicks
	if x >= 1000 && x <= 1250 {
		if y >= 270 && y <= 370 {
			// Menu panel clicks
			buttonY := y - 270
			menuPanel := g.ui.panels["menu"]
			for _, button := range menuPanel.Buttons {
				if buttonY >= button.Y && buttonY <= button.Y+button.Height && button.Enabled {
					if button.Action != nil {
						button.Action()
					}
					break
				}
			}
		}
	}

	// Room clicks
	for _, room := range g.rooms {
		roomX := room.X * tileSize
		roomY := room.Y * tileSize
		roomWidth := room.Width * tileSize
		roomHeight := room.Height * tileSize

		if x >= roomX && x <= roomX+roomWidth && y >= roomY && y <= roomY+roomHeight {
			// Find room index
			for roomIndex, r := range g.rooms {
				if r.ID == room.ID {
					g.handleRoomClick(roomIndex)
					break
				}
			}
			break
		}
	}
}

// handleRoomClick handle room click
func (g *Game) handleRoomClick(roomIndex int) {
	room := &g.rooms[roomIndex]
	if room.Occupied {
		// Check out guest
		for i, guest := range g.guests {
			if guest.RoomID == room.ID {
				// Calculate money based on stay duration
				stayDuration := int(time.Since(guest.CheckInTime).Seconds()) / 10 // 10 seconds = 1 day in game
				if stayDuration < 1 {
					stayDuration = 1
				}
				g.gameData.Money += room.Price * stayDuration
				g.gameData.GuestCount--
				g.guests = append(g.guests[:i], g.guests[i+1:]...)
				room.Occupied = false
				room.GuestID = 0
				break
			}
		}
	} else {
		// Try to assign a guest
		if g.gameData.GuestCount < g.gameData.MaxGuests {
			g.addNewGuest()
		}
	}
}

// handlePauseClick handle pause click
func (g *Game) handlePauseClick(x, y int) {
	// ESC key handling is already in Update
}

// addNewRoom add new room
func (g *Game) addNewRoom() {
	if g.gameData.Money >= 1000 {
		g.gameData.Money -= 1000
		newRoom := Room{
			ID:       len(g.rooms) + 1,
			X:        2 + (len(g.rooms)%3)*6,
			Y:        6 + (len(g.rooms)/3)*4,
			Width:    4,
			Height:   3,
			Type:     "single",
			Occupied: false,
			Price:    50,
		}
		g.rooms = append(g.rooms, newRoom)
		g.gameData.MaxGuests++
	}
}

// addNewGuest add new guest
func (g *Game) addNewGuest() {
	if g.gameData.GuestCount < g.gameData.MaxGuests {
		// Find available room
		var availableRoom *Room
		for i := range g.rooms {
			if !g.rooms[i].Occupied {
				availableRoom = &g.rooms[i]
				break
			}
		}

		if availableRoom != nil {
			guestNames := []string{"John", "Mary", "David", "Sarah", "Michael", "Emma", "James", "Lisa", "Robert", "Anna"}
			guest := Guest{
				ID:           len(g.guests) + 1,
				X:            float64(availableRoom.X*tileSize + tileSize/2),
				Y:            float64(availableRoom.Y*tileSize + tileSize/2),
				Satisfaction: 75 + rand.Intn(25),
				StayDuration: 30 + rand.Intn(60), // 30-90 seconds
				RoomID:       availableRoom.ID,
				Name:         guestNames[rand.Intn(len(guestNames))],
				CharacterID:  rand.Intn(20) + 1, // 1-20 character sprites
				CheckInTime:  time.Now(),
			}
			g.guests = append(g.guests, guest)
			g.gameData.GuestCount++
			g.gameData.Money += 50 // Check-in fee
			availableRoom.Occupied = true
			availableRoom.GuestID = guest.ID
		}
	}
}

// updateGuests update guests
func (g *Game) updateGuests() {
	now := time.Now()
	for i := range g.guests {
		// Check if guest should leave based on time spent
		timeSpent := int(now.Sub(g.guests[i].CheckInTime).Seconds())
		if timeSpent >= g.guests[i].StayDuration {
			// Guest leaves
			guest := g.guests[i]
			for j := range g.rooms {
				if g.rooms[j].ID == guest.RoomID {
					g.rooms[j].Occupied = false
					g.rooms[j].GuestID = 0
					break
				}
			}
			g.gameData.GuestCount--
			g.guests = append(g.guests[:i], g.guests[i+1:]...)
			break
		}
	}
}

// Draw drawing
func (g *Game) Draw(screen *ebiten.Image) {
	switch g.state {
	case GameStateMainMenu:
		g.drawMainMenu(screen)
	case GameStatePlaying:
		g.drawGame(screen)
	case GameStatePaused:
		g.drawPause(screen)
	}
}

// drawMainMenu draw main menu
func (g *Game) drawMainMenu(screen *ebiten.Image) {
	// Background
	screen.Fill(color.RGBA{44, 62, 80, 255})

	// Title
	title := "HOTEL MANAGER"
	text.Draw(screen, title, basicfont.Face7x13, 640-len(title)*7/2, 150, color.RGBA{236, 240, 241, 255})

	// Buttons
	buttons := []string{"New Game", "Load Game", "Settings", "Exit"}
	for i, buttonText := range buttons {
		y := 300 + i*70
		// Button background
		ebitenutil.DrawRect(screen, 540, float64(y), 200, 50, color.RGBA{52, 152, 219, 255})
		// Button text
		text.Draw(screen, buttonText, basicfont.Face7x13, 640-len(buttonText)*7/2, y+30, color.RGBA{255, 255, 255, 255})
	}
}

// drawGame draw game
func (g *Game) drawGame(screen *ebiten.Image) {
	// Hotel map drawing
	g.drawHotelMap(screen)

	// Rooms drawing
	g.drawRooms(screen)

	// Guests drawing
	g.drawGuests(screen)

	// UI drawing
	g.drawUI(screen)
}

// drawHotelMap draw hotel map
func (g *Game) drawHotelMap(screen *ebiten.Image) {
	for y := 0; y < 15; y++ {
		for x := 0; x < 20; x++ {
			tileX := float64(x * tileSize)
			tileY := float64(y * tileSize)
			
			if g.hotelMap[y][x] == 1 {
				// Wall - use asset if available
				if wallImg := g.assets.images["interiors"]; wallImg != nil {
					// Draw wall tile from spritesheet (assuming wall is at position 0,0)
					op := &ebiten.DrawImageOptions{}
					op.GeoM.Translate(tileX, tileY)
					screen.DrawImage(wallImg.SubImage(image.Rect(0, 0, tileSize, tileSize)).(*ebiten.Image), op)
				} else {
					// Fallback to colored rectangle
					ebitenutil.DrawRect(screen, tileX, tileY, tileSize, tileSize, color.RGBA{52, 73, 94, 255})
				}
			} else {
				// Floor - use asset if available
				if floorImg := g.assets.images["interiors"]; floorImg != nil {
					// Draw floor tile from spritesheet (assuming floor is at position 32,0)
					op := &ebiten.DrawImageOptions{}
					op.GeoM.Translate(tileX, tileY)
					screen.DrawImage(floorImg.SubImage(image.Rect(32, 0, 64, 32)).(*ebiten.Image), op)
				} else {
					// Fallback to colored rectangle
					ebitenutil.DrawRect(screen, tileX, tileY, tileSize, tileSize, color.RGBA{236, 240, 241, 255})
				}
			}
		}
	}
}

// drawRooms draw rooms
func (g *Game) drawRooms(screen *ebiten.Image) {
	for _, room := range g.rooms {
		roomX := float64(room.X * tileSize)
		roomY := float64(room.Y * tileSize)
		roomWidth := float64(room.Width * tileSize)
		roomHeight := float64(room.Height * tileSize)

		// Room background - use asset if available
		if roomImg := g.assets.images["interiors"]; roomImg != nil {
			// Draw room tiles from spritesheet
			for y := 0; y < room.Height; y++ {
				for x := 0; x < room.Width; x++ {
					tileX := roomX + float64(x*tileSize)
					tileY := roomY + float64(y*tileSize)
					
					// Use different tile based on room state
					spriteX := 64 // Default room tile
					if room.Occupied {
						spriteX = 96 // Occupied room tile
					}
					
					op := &ebiten.DrawImageOptions{}
					op.GeoM.Translate(tileX, tileY)
					screen.DrawImage(roomImg.SubImage(image.Rect(spriteX, 0, spriteX+32, 32)).(*ebiten.Image), op)
				}
			}
		} else {
			// Fallback to colored rectangles
			roomColor := color.RGBA{46, 204, 113, 255} // Green (available)
			if room.Occupied {
				roomColor = color.RGBA{231, 76, 60, 255} // Red (occupied)
			}
			ebitenutil.DrawRect(screen, roomX, roomY, roomWidth, roomHeight, roomColor)
		}

		// Room border
		ebitenutil.DrawRect(screen, roomX, roomY, roomWidth, 2, color.RGBA{44, 62, 80, 255})
		ebitenutil.DrawRect(screen, roomX, roomY, 2, roomHeight, color.RGBA{44, 62, 80, 255})
		ebitenutil.DrawRect(screen, roomX+roomWidth-2, roomY, 2, roomHeight, color.RGBA{44, 62, 80, 255})
		ebitenutil.DrawRect(screen, roomX, roomY+roomHeight-2, roomWidth, 2, color.RGBA{44, 62, 80, 255})

		// Room number
		roomText := fmt.Sprintf("Room %d", room.ID)
		text.Draw(screen, roomText, basicfont.Face7x13, int(roomX+roomWidth/2)-len(roomText)*7/2, int(roomY+roomHeight/2)+5, color.RGBA{255, 255, 255, 255})
	}
}

// drawGuests draw guests
func (g *Game) drawGuests(screen *ebiten.Image) {
	for _, guest := range g.guests {
		// Draw character sprite if available
		if charImg := g.assets.images[fmt.Sprintf("character_%d", guest.CharacterID)]; charImg != nil {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(guest.X-16, guest.Y-16) // Center the sprite
			screen.DrawImage(charImg, op)
		} else {
			// Fallback to simple circle
			ebitenutil.DrawCircle(screen, guest.X, guest.Y, 8, color.RGBA{243, 156, 18, 255})
		}
		
		// Guest name
		text.Draw(screen, guest.Name, basicfont.Face7x13, int(guest.X)-len(guest.Name)*7/2, int(guest.Y)-25, color.RGBA{255, 255, 255, 255})
		
		// Show remaining time
		remainingTime := guest.StayDuration - int(time.Since(guest.CheckInTime).Seconds())
		if remainingTime > 0 {
			timeText := fmt.Sprintf("%ds", remainingTime)
			text.Draw(screen, timeText, basicfont.Face7x13, int(guest.X)-len(timeText)*7/2, int(guest.Y)+20, color.RGBA{255, 255, 0, 255})
		}
	}
}

// drawUI draw user interface
func (g *Game) drawUI(screen *ebiten.Image) {
	// Information panel
	infoPanel := g.ui.panels["info"]
	ebitenutil.DrawRect(screen, float64(infoPanel.X), float64(infoPanel.Y), float64(infoPanel.Width), float64(infoPanel.Height), color.RGBA{52, 73, 94, 255})
	
	// Information
	text.Draw(screen, infoPanel.Title, basicfont.Face7x13, infoPanel.X+10, infoPanel.Y+20, color.RGBA{236, 240, 241, 255})
	text.Draw(screen, fmt.Sprintf("Money: $%d", g.gameData.Money), basicfont.Face7x13, infoPanel.X+10, infoPanel.Y+50, color.RGBA{46, 204, 113, 255})
	text.Draw(screen, fmt.Sprintf("Guests: %d/%d", g.gameData.GuestCount, g.gameData.MaxGuests), basicfont.Face7x13, infoPanel.X+10, infoPanel.Y+70, color.RGBA{52, 152, 219, 255})
	text.Draw(screen, fmt.Sprintf("Satisfaction: %d%%", g.gameData.Satisfaction), basicfont.Face7x13, infoPanel.X+10, infoPanel.Y+90, color.RGBA{243, 156, 18, 255})
	text.Draw(screen, fmt.Sprintf("Day: %d", g.gameData.Day), basicfont.Face7x13, infoPanel.X+10, infoPanel.Y+110, color.RGBA{155, 89, 182, 255})

	// Menu panel
	menuPanel := g.ui.panels["menu"]
	ebitenutil.DrawRect(screen, float64(menuPanel.X), float64(menuPanel.Y), float64(menuPanel.Width), float64(menuPanel.Height), color.RGBA{52, 73, 94, 255})
	
	// Menu buttons
	for _, button := range menuPanel.Buttons {
		buttonColor := color.RGBA{52, 152, 219, 255}
		if !button.Enabled {
			buttonColor = color.RGBA{128, 128, 128, 255}
		}
		ebitenutil.DrawRect(screen, float64(menuPanel.X+button.X), float64(menuPanel.Y+button.Y), float64(button.Width), float64(button.Height), buttonColor)
		text.Draw(screen, button.Text, basicfont.Face7x13, menuPanel.X+button.X+10, menuPanel.Y+button.Y+20, color.RGBA{255, 255, 255, 255})
	}
}

// drawPause draw pause menu
func (g *Game) drawPause(screen *ebiten.Image) {
	// Transparent background
	ebitenutil.DrawRect(screen, 0, 0, screenWidth, screenHeight, color.RGBA{0, 0, 0, 128})
	
	// Pause text
	pauseText := "GAME PAUSED"
	text.Draw(screen, pauseText, basicfont.Face7x13, 640-len(pauseText)*7/2, 300, color.RGBA{255, 255, 255, 255})
	
	escText := "ESC - Continue"
	text.Draw(screen, escText, basicfont.Face7x13, 640-len(escText)*7/2, 350, color.RGBA{255, 255, 255, 255})
}

// Layout screen sizing
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 720
}

// main main function
func main() {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// Ebitengine settings
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hotel Manager - Business Simulation Game")
	ebiten.SetWindowResizable(true)

	// Create and start game
	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
} 