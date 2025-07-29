# 🏨 Hotel Manager - Projekt Összefoglaló

## 📋 Projekt Áttekintés

A **Hotel Manager** egy üzleti szimulációs játék, amelyet az **Ebitengine** játék motor segítségével fejlesztettünk. A játék célja, hogy a játékosok saját hotel vagy panzió üzemeltetését irányíthassák, miközben gazdasági döntéseket hoznak és vendégeket szolgálnak ki.

## 🎯 Főbb Jellemzők

### ✅ **Elkészült Funkciók (Fázis 1)**

1. **Alapvető Projekt Struktúra**
   - Ebitengine projekt beállítása
   - Moduláris kód struktúra
   - Asset kezelő rendszer
   - Játék adatkezelő rendszer

2. **Játék Motor Alapok**
   - Fő játék osztály (`HotelManager`)
   - Jelenet kezelés (`Scene` rendszer)
   - Asset betöltés és kezelés
   - Játék állapot kezelés

3. **Felhasználói Felület**
   - Fő menü jelenet
   - Hotel jelenet (alapvető)
   - Szünet menü
   - Magyar nyelvű felület

4. **Asset Integráció**
   - Modern Interiors asset pack használata
   - Karakter generátor rendszer
   - Animált objektumok támogatása
   - UI elemek integrálása

5. **Játék Adatkezelés**
   - Mentés/betöltés rendszer
   - Játék állapot követése
   - Statisztikák kezelése
   - Teljesítmény rendszer

## 🛠️ **Technikai Részletek**

### **Használt Technológiák**
- **Játék Motor**: Ebitengine 2.6.0
- **Programozási Nyelv**: JavaScript (ES6+)
- **Grafika**: Pixel art (32x32, 48x48 méretek)
- **Platform**: Windows (elsődleges)
- **Steam Integráció**: Tervezett

### **Projekt Struktúra**
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
│   └── GameData.js        # Játék adatok
├── main.js               # Fő játék fájl
├── index.html            # Web indító
├── package.json          # Projekt konfiguráció
├── build.js              # Build script
└── README.md            # Dokumentáció
```

## 🎮 **Játék Mechanikák (Tervezett)**

### **Fázis 2: Vendég Kezelés**
- Vendég fogadás és regisztrálás
- Szoba kiosztás rendszer
- Check-in/check-out folyamatok
- Vendég AI és viselkedés

### **Fázis 3: Gazdasági Rendszer**
- Pénzügyi kezelés
- Árazási stratégia
- Költségek követése
- Befektetések tervezése

### **Fázis 4: Szolgáltatások**
- Kávézó üzemeltetése
- Konditerem használat
- Extra szolgáltatások
- Marketing kampányok

### **Fázis 5: Fejlesztések**
- Hotel bővítés
- Lánc hotelek
- Verseny rendszer
- Steam integráció

## 📊 **Jelenlegi Állapot**

### **✅ Elkészült**
- [x] Projekt alapstruktúra
- [x] Ebitengine integráció
- [x] Asset kezelő rendszer
- [x] Játék adatkezelő
- [x] Alapvető UI rendszer
- [x] Fő menü és jelenetek
- [x] Mentés/betöltés rendszer
- [x] Magyar nyelvű felület
- [x] Build scriptek
- [x] Dokumentáció

### **🔄 Folyamatban**
- [ ] Asset betöltés optimalizálása
- [ ] Hotel térkép implementálása
- [ ] Vendég rendszer fejlesztése

### **📋 Tervezett**
- [ ] Gazdasági rendszer
- [ ] Szolgáltatások
- [ ] Steam integráció
- [ ] Teljesítmények
- [ ] Zene és hangok

## 🚀 **Használat**

### **Fejlesztői Környezet**
```bash
# Függőségek telepítése
npm install

# Fejlesztői szerver indítása
npm run dev

# Vagy használja a build scriptet
npm start
```

### **Build**
```bash
# Web build
npm run build

# Windows build
npm run build:windows

# Mindkettő
npm run build:all
```

## 🎯 **Következő Lépések**

### **Azonnali Feladatok**
1. **Asset Betöltés Javítása**
   - Sprite sheet kezelés optimalizálása
   - Karakter generátor tesztelése
   - Animált objektumok implementálása

2. **Hotel Térkép Fejlesztése**
   - Tile-based rendszer
   - Szoba elhelyezés
   - Navigáció és kamera

3. **Vendég Rendszer**
   - Vendég AI implementálása
   - Szoba kiosztás logika
   - Interakció rendszer

### **Középtávú Célok**
- Gazdasági rendszer implementálása
- Szolgáltatások hozzáadása
- Steam Workshop támogatás
- Teljesítmény rendszer

### **Hosszútávú Célok**
- Steam kiadás
- Mobil platform támogatás
- Multiplayer funkciók
- Modding támogatás

## 🏆 **Teljesítmények és Célok**

### **Játék Célok**
- Első vendég fogadása
- Pozitív értékelés elérése
- Hotel teljes kihasználása
- 5 csillagos értékelés
- Lánc hotelek létrehozása

### **Fejlesztési Célok**
- Stabil 60 FPS teljesítmény
- < 100MB játék méret
- < 30 másodperc betöltési idő
- 95%+ bug-mentes játék

## 🤝 **Közreműködés**

A projekt nyitott a közreműködésre! Ha szeretnél hozzájárulni:

1. Fork a projektet
2. Hozz létre egy feature branch-et
3. Implementáld a változtatásokat
4. Teszteld a funkciókat
5. Készíts Pull Request-et

## 📞 **Kapcsolat**

- **Fejlesztő**: Hotel Manager Team
- **Projekt**: Hotel Manager üzleti szimulációs játék
- **Platform**: Windows (Steam)
- **Státusz**: Fejlesztés alatt

---

**🏨 Hotel Manager** - Építsd fel álmaid hotelét! ✨

*Utolsó frissítés: 2024* 