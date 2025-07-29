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

- **Játék Motor**: Ebitengine
- **Grafika**: Pixel art (32x32, 48x48 méretek)
- **Nyelv**: Magyar
- **Platform**: Windows (elsődleges), macOS/Linux (később)
- **Steam Kiadás**: Igen, Steam platformra optimalizálva

## 📦 Telepítés

### Előfeltételek

- Node.js 18.0.0 vagy újabb
- Modern webböngésző (fejlesztéshez)

### Telepítési lépések

1. **Projekt klónozása**
   ```bash
   git clone [repository-url]
   cd hotel-manager
   ```

2. **Függőségek telepítése**
   ```bash
   npm install
   ```

3. **Fejlesztői szerver indítása**
   ```bash
   npm run dev
   ```

4. **Játék buildelése**
   ```bash
   npm run build
   ```

5. **Windows build**
   ```bash
   npm run build:windows
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

## 📁 Projekt Struktúra

```
hotel-manager/
├── assets/                 # Asset fájlok
│   ├── 1_Interiors/       # Bútorok és berendezés
│   ├── 2_Characters/      # Karakter assets
│   ├── 3_Animated_objects/ # Animált objektumok
│   ├── 4_User_Interface_Elements/ # UI elemek
│   └── 6_Home_Designs/    # Ház tervek
├── src/                   # Forráskód
│   ├── AssetManager.js    # Asset kezelő
│   ├── GameData.js        # Játék adatok
│   └── ...               # További modulok
├── main.js               # Fő játék fájl
├── package.json          # Projekt konfiguráció
└── README.md            # Dokumentáció
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