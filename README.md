# Hotel Manager - Üzleti Szimulációs Játék

## 📋 Leírás

A Hotel Manager egy üzleti szimulációs játék, ahol a játékos saját hotel vagy panzió üzemeltetését irányítja. A játék pixel art stílusban készül, használva a Modern Interiors asset pack-et.

## 🎮 Játék Funkciók

- **Hotel Építés és Berendezés**: Különböző típusú szobák építése és bútorok elhelyezése
- **Vendég Kezelés**: Vendégek fogadása, szoba kiosztás, elégedettség követése
- **Szolgáltatások**: Kávézó, konditerem, Wi-Fi, takarítási szolgáltatások
- **Gazdasági Menedzsment**: Árazás, költségek, bevételek, marketing
- **Karakter Rendszer**: Karakter generátor vendégekkel és alkalmazottakkal
- **Animált Objektumok**: Ajtók, fürdőkád, csapok és egyéb interaktív elemek

## 🛠️ Technikai Részletek

- **Játék Motor**: Ebitengine v2.6.0
- **Programozási Nyelv**: Go 1.21+
- **Grafika**: Pixel art (32x32, 48x48 méretek)
- **Nyelv**: Magyar
- **Platform**: Windows (elsődleges), macOS/Linux (később)
- **Steam Kiadás**: Igen, Steam platformra optimalizálva

## 📦 Telepítés

### Előfeltételek

- Go 1.21.0 vagy újabb
- Ebitengine v2.6.0
- Git

### Telepítési lépések

1. **Projekt klónozása**
   ```bash
   git clone https://github.com/Minzy-c/hotel-manager-game.git
   cd hotel-manager-game
   ```

2. **Függőségek telepítése**
   ```bash
   go mod tidy
   ```

3. **Fejlesztői mód indítása**
   ```bash
   go run main.go
   ```

4. **Játék buildelése**
   ```bash
   go build -o hotel-manager.exe main.go
   ```

5. **Windows GUI build**
   ```bash
   go build -ldflags="-H windowsgui" -o hotel-manager.exe main.go
   ```

## 🎯 Játék Mechanikák

### Alapvető Rendszer
- Hotel térkép és navigáció
- Karakter rendszer
- Asset betöltés és kezelés

### Vendég Kezelés
- Vendég fogadás és regisztrálás
- Szoba kiosztás
- Check-in/check-out folyamatok
- Vendég AI és viselkedés

### Gazdasági Rendszer
- Pénzügyi kezelés
- Árazási stratégia
- Költségek követése
- Befektetések tervezése

### Szolgáltatások
- Kávézó üzemeltetése
- Konditerem használat
- Extra szolgáltatások
- Marketing kampányok

### Jelenlegi Implementáció
- ✅ **Alapvető UI**: Főmenü, játék képernyő, szünet menü
- ✅ **Hotel Térkép**: 20x15 méretű térkép falakkal és padlóval
- ✅ **Szoba Rendszer**: Szobák létrehozása és megjelenítése
- ✅ **Vendég Rendszer**: Vendégek hozzáadása és kezelése
- ✅ **Gazdasági Rendszer**: Pénz kezelés, szoba árak
- ✅ **Input Kezelés**: Egér kattintás, billentyűzet (ESC)
- ✅ **Játék Állapotok**: Főmenü, játék, szünet

## 📁 Projekt Struktúra

```
hotel-manager/
├── assets/                 # Asset fájlok
│   ├── 1_Interiors/       # Bútorok és berendezés
│   ├── 2_Characters/      # Karakter assets
│   ├── 3_Animated_objects/ # Animált objektumok
│   ├── 4_User_Interface_Elements/ # UI elemek
│   └── 6_Home_Designs/    # Ház tervek
├── main.go                # Fő játék fájl (Go)
├── go.mod                 # Go modul fájl
├── package.json           # Projekt metaadatok
├── README.md              # Dokumentáció
├── LICENSE                # Licenc
└── .gitignore            # Git ignore fájl
```

## 🎮 Irányítás

- **Egér**: Navigáció és interakció
- **ESC**: Szünet menü
- **F11**: Teljes képernyő
- **Ctrl+S**: Gyors mentés

## 🏆 Teljesítmények

- Első vendég fogadása
- Pozitív értékelés elérése
- Hotel teljes kihasználása
- 5 csillagos értékelés
- Lánc hotelek létrehozása

## 🔧 Fejlesztés

### Fejlesztési Fázisok

1. **Fázis 1**: Alapvető rendszer
2. **Fázis 2**: Vendég kezelés
3. **Fázis 3**: Gazdasági rendszer
4. **Fázis 4**: Szolgáltatások
5. **Fázis 5**: Fejlesztések
6. **Fázis 6**: Finomhangolás

### Közreműködés

1. Fork a projektet
2. Hozz létre egy feature branch-et (`git checkout -b feature/uj-funkcio`)
3. Commit a változtatásokat (`git commit -am 'Új funkció hozzáadása'`)
4. Push a branch-et (`git push origin feature/uj-funkcio`)
5. Hozz létre egy Pull Request-et

## 📄 Licenc

Ez a projekt MIT licenc alatt van kiadva. Lásd a `LICENSE` fájlt részletekért.

## 🤝 Kapcsolat

- **Fejlesztő**: Hotel Manager Team
- **Email**: [email]
- **Steam**: [Steam link]

## 🙏 Köszönet

- **Limezu** - Modern Interiors asset pack
- **Ebitengine** - Játék motor
- **Közösség** - Visszajelzések és javaslatok

---

**Hotel Manager** - Építsd fel álmaid hotelét! 🏨✨ 