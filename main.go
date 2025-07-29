package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
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

// Game állapotok
const (
	GameStateMainMenu = iota
	GameStatePlaying
	GameStatePaused
)

// Game fő struktúra
type Game struct {
	state       int
	assets      *AssetManager
	gameData    *GameData
	hotelMap    [][]int
	rooms       []Room
	guests      []Guest
	ui          *UI
	lastUpdate  time.Time
}

// AssetManager asset kezelés
type AssetManager struct {
	images map[string]*ebiten.Image
}

// GameData játék adatok
type GameData struct {
	Money        int
	HotelLevel   int
	GuestCount   int
	MaxGuests    int
	Satisfaction int
	Day          int
	HotelName    string
}

// Room szoba struktúra
type Room struct {
	ID       int
	X, Y     int
	Width    int
	Height   int
	Type     string
	Occupied bool
	Price    int
}

// Guest vendég struktúra
type Guest struct {
	ID           int
	X, Y         float64
	Satisfaction int
	StayDuration int
}

// UI felhasználói felület
type UI struct {
	panels map[string]Panel
}

// Panel UI panel
type Panel struct {
	X, Y, Width, Height int
	Title               string
	Buttons             []Button
}

// Button gomb
type Button struct {
	X, Y, Width, Height int
	Text                string
	Action              func()
}

// NewGame új játék létrehozása
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
			HotelName:    "Új Hotel",
		},
		hotelMap: make([][]int, 15),
		rooms:    make([]Room, 0),
		guests:   make([]Guest, 0),
		ui:       &UI{panels: make(map[string]Panel)},
	}

	// Hotel térkép inicializálása
	game.initHotelMap()
	game.createInitialRooms()
	game.initUI()

	return game
}

// initHotelMap hotel térkép inicializálása
func (g *Game) initHotelMap() {
	for y := 0; y < 15; y++ {
		g.hotelMap[y] = make([]int, 20)
		for x := 0; x < 20; x++ {
			if y == 0 || y == 14 || x == 0 || x == 19 {
				g.hotelMap[y][x] = 1 // Fal
			} else {
				g.hotelMap[y][x] = 0 // Padló
			}
		}
	}
}

// createInitialRooms kezdeti szobák létrehozása
func (g *Game) createInitialRooms() {
	g.rooms = []Room{
		{ID: 1, X: 2, Y: 2, Width: 4, Height: 3, Type: "single", Occupied: false, Price: 50},
		{ID: 2, X: 8, Y: 2, Width: 4, Height: 3, Type: "single", Occupied: false, Price: 50},
		{ID: 3, X: 14, Y: 2, Width: 4, Height: 3, Type: "double", Occupied: false, Price: 80},
	}
}

// initUI felhasználói felület inicializálása
func (g *Game) initUI() {
	// Információs panel
	g.ui.panels["info"] = Panel{
		X: 1000, Y: 50, Width: 250, Height: 200,
		Title: "HOTEL INFORMÁCIÓK",
	}

	// Menü panel
	g.ui.panels["menu"] = Panel{
		X: 1000, Y: 270, Width: 250, Height: 400,
		Title: "MENÜ",
		Buttons: []Button{
			{X: 10, Y: 20, Width: 230, Height: 30, Text: "Új Szoba"},
			{X: 10, Y: 60, Width: 230, Height: 30, Text: "Vendég Fogadás"},
			{X: 10, Y: 100, Width: 230, Height: 30, Text: "Takarítás"},
			{X: 10, Y: 140, Width: 230, Height: 30, Text: "Szolgáltatások"},
			{X: 10, Y: 180, Width: 230, Height: 30, Text: "Pénzügyek"},
			{X: 10, Y: 220, Width: 230, Height: 30, Text: "Beállítások"},
			{X: 10, Y: 260, Width: 230, Height: 30, Text: "Mentés"},
			{X: 10, Y: 300, Width: 230, Height: 30, Text: "Főmenü"},
		},
	}
}

// Update játék frissítése
func (g *Game) Update() error {
	now := time.Now()
	if g.lastUpdate.IsZero() {
		g.lastUpdate = now
	}

	// Input kezelés
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		if g.state == GameStatePlaying {
			g.state = GameStatePaused
		} else if g.state == GameStatePaused {
			g.state = GameStatePlaying
		}
	}

	// Egér kattintás kezelése
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		g.handleMouseClick(x, y)
	}

	// Vendégek frissítése
	if g.state == GameStatePlaying {
		g.updateGuests()
	}

	g.lastUpdate = now
	return nil
}

// handleMouseClick egér kattintás kezelése
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

// handleMainMenuClick főmenü kattintás kezelése
func (g *Game) handleMainMenuClick(x, y int) {
	// Új játék gomb
	if x >= 540 && x <= 740 && y >= 280 && y <= 330 {
		g.state = GameStatePlaying
	}
	// Kilépés gomb
	if x >= 540 && x <= 740 && y >= 500 && y <= 550 {
		// TODO: Kilépés implementálása
	}
}

// handleGameClick játék kattintás kezelése
func (g *Game) handleGameClick(x, y int) {
	// UI panel kattintások
	if x >= 1000 && x <= 1250 {
		if y >= 270 && y <= 670 {
			// Menü panel kattintások
			buttonY := y - 270
			if buttonY >= 20 && buttonY <= 50 {
				// Új Szoba
				g.addNewRoom()
			} else if buttonY >= 60 && buttonY <= 90 {
				// Vendég Fogadás
				g.addNewGuest()
			}
		}
	}
}

// handlePauseClick szünet kattintás kezelése
func (g *Game) handlePauseClick(x, y int) {
	// ESC gomb kezelése már a Update-ben van
}

// addNewRoom új szoba hozzáadása
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

// addNewGuest új vendég hozzáadása
func (g *Game) addNewGuest() {
	if g.gameData.GuestCount < g.gameData.MaxGuests {
		guest := Guest{
			ID:           len(g.guests) + 1,
			X:            float64(rand.Intn(18)+1) * tileSize,
			Y:            float64(rand.Intn(13)+1) * tileSize,
			Satisfaction: 75 + rand.Intn(25),
			StayDuration: 3 + rand.Intn(5),
		}
		g.guests = append(g.guests, guest)
		g.gameData.GuestCount++
		g.gameData.Money += 50 // Check-in díj
	}
}

// updateGuests vendégek frissítése
func (g *Game) updateGuests() {
	for i := range g.guests {
		g.guests[i].StayDuration--
		if g.guests[i].StayDuration <= 0 {
			// Vendég távozik
			g.gameData.GuestCount--
			g.guests = append(g.guests[:i], g.guests[i+1:]...)
			break
		}
	}
}

// Draw rajzolás
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

// drawMainMenu főmenü rajzolása
func (g *Game) drawMainMenu(screen *ebiten.Image) {
	// Háttér
	screen.Fill(color.RGBA{44, 62, 80, 255})

	// Cím
	title := "HOTEL MANAGER"
	text.Draw(screen, title, basicfont.Face7x13, 640-len(title)*7/2, 150, color.RGBA{236, 240, 241, 255})

	// Gombok
	buttons := []string{"Új Játék", "Játék Betöltése", "Beállítások", "Kilépés"}
	for i, buttonText := range buttons {
		y := 300 + i*70
		// Gomb háttér
		ebitenutil.DrawRect(screen, 540, float64(y), 200, 50, color.RGBA{52, 152, 219, 255})
		// Gomb szöveg
		text.Draw(screen, buttonText, basicfont.Face7x13, 640-len(buttonText)*7/2, y+30, color.RGBA{255, 255, 255, 255})
	}
}

// drawGame játék rajzolása
func (g *Game) drawGame(screen *ebiten.Image) {
	// Hotel térkép rajzolása
	g.drawHotelMap(screen)

	// Szobák rajzolása
	g.drawRooms(screen)

	// Vendégek rajzolása
	g.drawGuests(screen)

	// UI rajzolása
	g.drawUI(screen)
}

// drawHotelMap hotel térkép rajzolása
func (g *Game) drawHotelMap(screen *ebiten.Image) {
	for y := 0; y < 15; y++ {
		for x := 0; x < 20; x++ {
			tileX := float64(x * tileSize)
			tileY := float64(y * tileSize)
			
			if g.hotelMap[y][x] == 1 {
				// Fal
				ebitenutil.DrawRect(screen, tileX, tileY, tileSize, tileSize, color.RGBA{52, 73, 94, 255})
			} else {
				// Padló
				ebitenutil.DrawRect(screen, tileX, tileY, tileSize, tileSize, color.RGBA{236, 240, 241, 255})
			}
		}
	}
}

// drawRooms szobák rajzolása
func (g *Game) drawRooms(screen *ebiten.Image) {
	for _, room := range g.rooms {
		roomX := float64(room.X * tileSize)
		roomY := float64(room.Y * tileSize)
		roomWidth := float64(room.Width * tileSize)
		roomHeight := float64(room.Height * tileSize)

		// Szoba háttér
		roomColor := color.RGBA{46, 204, 113, 255} // Zöld (szabad)
		if room.Occupied {
			roomColor = color.RGBA{231, 76, 60, 255} // Piros (foglalt)
		}
		ebitenutil.DrawRect(screen, roomX, roomY, roomWidth, roomHeight, roomColor)

		// Szoba keret
		ebitenutil.DrawRect(screen, roomX, roomY, roomWidth, 2, color.RGBA{44, 62, 80, 255})
		ebitenutil.DrawRect(screen, roomX, roomY, 2, roomHeight, color.RGBA{44, 62, 80, 255})
		ebitenutil.DrawRect(screen, roomX+roomWidth-2, roomY, 2, roomHeight, color.RGBA{44, 62, 80, 255})
		ebitenutil.DrawRect(screen, roomX, roomY+roomHeight-2, roomWidth, 2, color.RGBA{44, 62, 80, 255})

		// Szoba szám
		roomText := fmt.Sprintf("Szoba %d", room.ID)
		text.Draw(screen, roomText, basicfont.Face7x13, int(roomX+roomWidth/2)-len(roomText)*7/2, int(roomY+roomHeight/2)+5, color.RGBA{255, 255, 255, 255})
	}
}

// drawGuests vendégek rajzolása
func (g *Game) drawGuests(screen *ebiten.Image) {
	for _, guest := range g.guests {
		// Vendég (egyszerű kör)
		ebitenutil.DrawCircle(screen, guest.X+float64(tileSize)/2, guest.Y+float64(tileSize)/2, 8, color.RGBA{243, 156, 18, 255})
	}
}

// drawUI felhasználói felület rajzolása
func (g *Game) drawUI(screen *ebiten.Image) {
	// Információs panel
	infoPanel := g.ui.panels["info"]
	ebitenutil.DrawRect(screen, float64(infoPanel.X), float64(infoPanel.Y), float64(infoPanel.Width), float64(infoPanel.Height), color.RGBA{52, 73, 94, 255})
	
	// Információk
	text.Draw(screen, infoPanel.Title, basicfont.Face7x13, infoPanel.X+10, infoPanel.Y+20, color.RGBA{236, 240, 241, 255})
	text.Draw(screen, fmt.Sprintf("Pénz: $%d", g.gameData.Money), basicfont.Face7x13, infoPanel.X+10, infoPanel.Y+50, color.RGBA{46, 204, 113, 255})
	text.Draw(screen, fmt.Sprintf("Vendégek: %d/%d", g.gameData.GuestCount, g.gameData.MaxGuests), basicfont.Face7x13, infoPanel.X+10, infoPanel.Y+70, color.RGBA{52, 152, 219, 255})
	text.Draw(screen, fmt.Sprintf("Elégedettség: %d%%", g.gameData.Satisfaction), basicfont.Face7x13, infoPanel.X+10, infoPanel.Y+90, color.RGBA{243, 156, 18, 255})
	text.Draw(screen, fmt.Sprintf("Nap: %d", g.gameData.Day), basicfont.Face7x13, infoPanel.X+10, infoPanel.Y+110, color.RGBA{155, 89, 182, 255})

	// Menü panel
	menuPanel := g.ui.panels["menu"]
	ebitenutil.DrawRect(screen, float64(menuPanel.X), float64(menuPanel.Y), float64(menuPanel.Width), float64(menuPanel.Height), color.RGBA{52, 73, 94, 255})
	
	// Menü gombok
	for _, button := range menuPanel.Buttons {
		ebitenutil.DrawRect(screen, float64(menuPanel.X+button.X), float64(menuPanel.Y+button.Y), float64(button.Width), float64(button.Height), color.RGBA{52, 152, 219, 255})
		text.Draw(screen, button.Text, basicfont.Face7x13, menuPanel.X+button.X+10, menuPanel.Y+button.Y+20, color.RGBA{255, 255, 255, 255})
	}
}

// drawPause szünet menü rajzolása
func (g *Game) drawPause(screen *ebiten.Image) {
	// Átlátszó háttér
	ebitenutil.DrawRect(screen, 0, 0, screenWidth, screenHeight, color.RGBA{0, 0, 0, 128})
	
	// Szünet szöveg
	pauseText := "JÁTÉK SZÜNETELTETVE"
	text.Draw(screen, pauseText, basicfont.Face7x13, 640-len(pauseText)*7/2, 300, color.RGBA{255, 255, 255, 255})
	
	escText := "ESC - Folytatás"
	text.Draw(screen, escText, basicfont.Face7x13, 640-len(escText)*7/2, 350, color.RGBA{255, 255, 255, 255})
}

// Layout képernyő méretezése
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 720
}

// main fő függvény
func main() {
	// Random seed inicializálása
	rand.Seed(time.Now().UnixNano())

	// Ebitengine beállítások
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hotel Manager - Üzleti Szimulációs Játék")
	ebiten.SetWindowResizable(true)

	// Játék létrehozása és indítása
	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
} 