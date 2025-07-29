import { Game, Scene, Vector2 } from 'ebiten';
import { AssetManager } from './src/AssetManager.js';
import { GameData } from './src/GameData.js';

// Játék állapotok
const GAME_STATES = {
    MAIN_MENU: 'main_menu',
    PLAYING: 'playing',
    PAUSED: 'paused',
    GAME_OVER: 'game_over'
};

// Fő játék osztály
class HotelManager extends Game {
    constructor() {
        super({
            title: 'Hotel Manager',
            width: 1280,
            height: 720,
            scale: 1,
            vsync: true,
            resizable: true
        });

        this.gameState = GAME_STATES.MAIN_MENU;
        this.currentScene = null;
        this.assetManager = new AssetManager();
        this.gameData = new GameData();

        this.init();
    }

    async init() {
        // Asset betöltés
        await this.loadAssets();
        
        // Jelenetek inicializálása
        this.scenes = {
            mainMenu: new MainMenuScene(this),
            hotel: new HotelScene(this),
            pause: new PauseScene(this)
        };

        this.currentScene = this.scenes.mainMenu;
    }

    async loadAssets() {
        console.log('Asset betöltés kezdése...');
        const success = await this.assetManager.loadAllAssets();
        if (!success) {
            console.error('Hiba az asset betöltés során!');
        }
    }

    update() {
        if (this.currentScene) {
            this.currentScene.update();
        }
    }

    draw(screen) {
        if (this.currentScene) {
            this.currentScene.draw(screen);
        }
    }

    changeScene(sceneName) {
        if (this.scenes[sceneName]) {
            this.currentScene = this.scenes[sceneName];
            this.currentScene.init();
        }
    }

    getGameData() {
        return this.gameData.getData();
    }

    setGameData(data) {
        this.gameData.setData(data);
    }
}

// Fő menü jelenet
class MainMenuScene extends Scene {
    constructor(game) {
        super();
        this.game = game;
        this.buttons = [];
        this.init();
    }

    init() {
        // Menü gombok létrehozása
        this.buttons = [
            {
                text: 'Új Játék',
                x: 640,
                y: 300,
                width: 200,
                height: 50,
                action: () => this.game.changeScene('hotel')
            },
            {
                text: 'Játék Betöltése',
                x: 640,
                y: 370,
                width: 200,
                height: 50,
                action: () => this.loadGame()
            },
            {
                text: 'Beállítások',
                x: 640,
                y: 440,
                width: 200,
                height: 50,
                action: () => this.openSettings()
            },
            {
                text: 'Kilépés',
                x: 640,
                y: 510,
                width: 200,
                height: 50,
                action: () => this.game.quit()
            }
        ];
    }

    update() {
        // Menü logika
    }

    draw(screen) {
        // Háttér rajzolása
        screen.fill('#2c3e50');
        
        // Cím rajzolása
        screen.drawText('HOTEL MANAGER', {
            x: 640,
            y: 150,
            fontSize: 48,
            color: '#ecf0f1',
            align: 'center'
        });

        // Gombok rajzolása
        this.buttons.forEach(button => {
            screen.drawRect(button.x - button.width/2, button.y - button.height/2, 
                          button.width, button.height, '#3498db');
            screen.drawText(button.text, {
                x: button.x,
                y: button.y + 8,
                fontSize: 18,
                color: '#ffffff',
                align: 'center'
            });
        });
    }

    loadGame() {
        // TODO: Játék betöltés implementálása
        console.log('Játék betöltése...');
    }

    openSettings() {
        // TODO: Beállítások menü implementálása
        console.log('Beállítások megnyitása...');
    }
}

// Hotel jelenet (fő játék)
class HotelScene extends Scene {
    constructor(game) {
        super();
        this.game = game;
        this.hotelMap = [];
        this.guests = [];
        this.rooms = [];
        this.ui = new HotelUI(this);
        this.init();
    }

    init() {
        // Hotel térkép inicializálása
        this.initHotelMap();
        
        // Kezdeti szobák létrehozása
        this.createInitialRooms();
    }

    initHotelMap() {
        // 20x15 méretű hotel térkép
        this.hotelMap = Array(15).fill().map(() => Array(20).fill(0));
        
        // Alapvető struktúra
        for (let y = 0; y < 15; y++) {
            for (let x = 0; x < 20; x++) {
                if (y === 0 || y === 14 || x === 0 || x === 19) {
                    this.hotelMap[y][x] = 1; // Fal
                }
            }
        }
    }

    createInitialRooms() {
        // Kezdeti szobák létrehozása
        this.rooms = [
            { id: 1, x: 2, y: 2, width: 4, height: 3, type: 'single', occupied: false, price: 50 },
            { id: 2, x: 8, y: 2, width: 4, height: 3, type: 'single', occupied: false, price: 50 },
            { id: 3, x: 14, y: 2, width: 4, height: 3, type: 'double', occupied: false, price: 80 }
        ];
    }

    update() {
        // Vendégek frissítése
        this.updateGuests();
        
        // UI frissítése
        this.ui.update();
    }

    updateGuests() {
        // Vendégek AI és mozgás
        this.guests.forEach(guest => {
            guest.update();
        });
    }

    draw(screen) {
        // Hotel térkép rajzolása
        this.drawHotelMap(screen);
        
        // Szobák rajzolása
        this.drawRooms(screen);
        
        // Vendégek rajzolása
        this.drawGuests(screen);
        
        // UI rajzolása
        this.ui.draw(screen);
    }

    drawHotelMap(screen) {
        const tileSize = 32;
        
        for (let y = 0; y < 15; y++) {
            for (let x = 0; x < 20; x++) {
                const tileX = x * tileSize;
                const tileY = y * tileSize;
                
                if (this.hotelMap[y][x] === 1) {
                    // Fal rajzolása
                    screen.drawRect(tileX, tileY, tileSize, tileSize, '#34495e');
                } else {
                    // Padló rajzolása
                    screen.drawRect(tileX, tileY, tileSize, tileSize, '#ecf0f1');
                }
            }
        }
    }

    drawRooms(screen) {
        const tileSize = 32;
        
        this.rooms.forEach(room => {
            const roomX = room.x * tileSize;
            const roomY = room.y * tileSize;
            const roomWidth = room.width * tileSize;
            const roomHeight = room.height * tileSize;
            
            // Szoba háttér
            const color = room.occupied ? '#e74c3c' : '#2ecc71';
            screen.drawRect(roomX, roomY, roomWidth, roomHeight, color);
            
            // Szoba keret
            screen.drawRect(roomX, roomY, roomWidth, roomHeight, '#2c3e50', 2);
            
            // Szoba információ
            screen.drawText(`Szoba ${room.id}`, {
                x: roomX + roomWidth/2,
                y: roomY + roomHeight/2,
                fontSize: 12,
                color: '#ffffff',
                align: 'center'
            });
        });
    }

    drawGuests(screen) {
        const tileSize = 32;
        
        this.guests.forEach(guest => {
            const guestX = guest.x * tileSize + tileSize/2;
            const guestY = guest.y * tileSize + tileSize/2;
            
            // Vendég rajzolása (egyszerű kör)
            screen.drawCircle(guestX, guestY, 8, '#f39c12');
        });
    }
}

// Hotel UI osztály
class HotelUI {
    constructor(scene) {
        this.scene = scene;
        this.game = scene.game;
        this.panels = {
            info: { x: 1000, y: 50, width: 250, height: 200 },
            menu: { x: 1000, y: 270, width: 250, height: 400 }
        };
    }

    update() {
        // UI frissítés
    }

    draw(screen) {
        // Információs panel
        this.drawInfoPanel(screen);
        
        // Menü panel
        this.drawMenuPanel(screen);
    }

    drawInfoPanel(screen) {
        const panel = this.panels.info;
        const gameData = this.game.getGameData();
        
        // Panel háttér
        screen.drawRect(panel.x, panel.y, panel.width, panel.height, '#34495e');
        screen.drawRect(panel.x, panel.y, panel.width, panel.height, '#2c3e50', 2);
        
        // Információk
        screen.drawText('HOTEL INFORMÁCIÓK', {
            x: panel.x + panel.width/2,
            y: panel.y + 20,
            fontSize: 16,
            color: '#ecf0f1',
            align: 'center'
        });
        
        screen.drawText(`Pénz: $${gameData.money}`, {
            x: panel.x + 10,
            y: panel.y + 50,
            fontSize: 14,
            color: '#2ecc71'
        });
        
        screen.drawText(`Vendégek: ${gameData.guestCount}/${gameData.maxGuests}`, {
            x: panel.x + 10,
            y: panel.y + 70,
            fontSize: 14,
            color: '#3498db'
        });
        
        screen.drawText(`Elégedettség: ${gameData.satisfaction}%`, {
            x: panel.x + 10,
            y: panel.y + 90,
            fontSize: 14,
            color: '#f39c12'
        });
        
        screen.drawText(`Nap: ${gameData.day}`, {
            x: panel.x + 10,
            y: panel.y + 110,
            fontSize: 14,
            color: '#9b59b6'
        });
    }

    drawMenuPanel(screen) {
        const panel = this.panels.menu;
        
        // Panel háttér
        screen.drawRect(panel.x, panel.y, panel.width, panel.height, '#34495e');
        screen.drawRect(panel.x, panel.y, panel.width, panel.height, '#2c3e50', 2);
        
        // Menü gombok
        const buttons = [
            { text: 'Új Szoba', y: 20 },
            { text: 'Vendég Fogadás', y: 60 },
            { text: 'Takarítás', y: 100 },
            { text: 'Szolgáltatások', y: 140 },
            { text: 'Pénzügyek', y: 180 },
            { text: 'Beállítások', y: 220 },
            { text: 'Mentés', y: 260 },
            { text: 'Főmenü', y: 300 }
        ];
        
        buttons.forEach(button => {
            screen.drawRect(panel.x + 10, panel.y + button.y, 230, 30, '#3498db');
            screen.drawText(button.text, {
                x: panel.x + 125,
                y: panel.y + button.y + 20,
                fontSize: 14,
                color: '#ffffff',
                align: 'center'
            });
        });
    }
}

// Szünet jelenet
class PauseScene extends Scene {
    constructor(game) {
        super();
        this.game = game;
    }

    init() {
        // Szünet menü inicializálása
    }

    update() {
        // Szünet logika
    }

    draw(screen) {
        // Átlátszó háttér
        screen.drawRect(0, 0, 1280, 720, 'rgba(0,0,0,0.5)');
        
        // Szünet menü
        screen.drawText('JÁTÉK SZÜNETELTETVE', {
            x: 640,
            y: 300,
            fontSize: 32,
            color: '#ffffff',
            align: 'center'
        });
        
        screen.drawText('ESC - Folytatás', {
            x: 640,
            y: 350,
            fontSize: 18,
            color: '#ffffff',
            align: 'center'
        });
    }
}

// Játék indítása
const game = new HotelManager();
game.run(); 