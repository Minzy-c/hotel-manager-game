// Game Data Manager - Játék adatok kezelése
export class GameData {
    constructor() {
        this.defaultData = {
            money: 10000,
            hotelLevel: 1,
            guestCount: 0,
            maxGuests: 5,
            satisfaction: 0,
            day: 1,
            hotelName: 'Új Hotel',
            rooms: [],
            guests: [],
            staff: [],
            services: {
                cafe: false,
                gym: false,
                wifi: true,
                cleaning: true
            },
            prices: {
                singleRoom: 50,
                doubleRoom: 80,
                luxuryRoom: 150
            },
            expenses: {
                utilities: 100,
                cleaning: 50,
                maintenance: 75
            },
            reputation: 50,
            totalGuests: 0,
            totalEarnings: 0,
            achievements: []
        };
        
        this.currentData = { ...this.defaultData };
    }

    // Játék adatok beállítása
    setData(data) {
        this.currentData = { ...this.currentData, ...data };
    }

    // Játék adatok lekérése
    getData() {
        return { ...this.currentData };
    }

    // Konkrét érték beállítása
    setValue(key, value) {
        this.currentData[key] = value;
    }

    // Konkrét érték lekérése
    getValue(key) {
        return this.currentData[key];
    }

    // Pénz kezelés
    addMoney(amount) {
        this.currentData.money += amount;
        this.currentData.totalEarnings += Math.max(0, amount);
    }

    spendMoney(amount) {
        if (this.currentData.money >= amount) {
            this.currentData.money -= amount;
            return true;
        }
        return false;
    }

    // Vendég kezelés
    addGuest(guest) {
        if (this.currentData.guestCount < this.currentData.maxGuests) {
            this.currentData.guests.push(guest);
            this.currentData.guestCount++;
            this.currentData.totalGuests++;
            return true;
        }
        return false;
    }

    removeGuest(guestId) {
        const index = this.currentData.guests.findIndex(g => g.id === guestId);
        if (index !== -1) {
            this.currentData.guests.splice(index, 1);
            this.currentData.guestCount--;
            return true;
        }
        return false;
    }

    // Szoba kezelés
    addRoom(room) {
        this.currentData.rooms.push(room);
        this.currentData.maxGuests += room.capacity;
    }

    removeRoom(roomId) {
        const index = this.currentData.rooms.findIndex(r => r.id === roomId);
        if (index !== -1) {
            const room = this.currentData.rooms[index];
            this.currentData.maxGuests -= room.capacity;
            this.currentData.rooms.splice(index, 1);
            return true;
        }
        return false;
    }

    // Elégedettség frissítése
    updateSatisfaction() {
        if (this.currentData.guestCount === 0) {
            this.currentData.satisfaction = 0;
            return;
        }

        let totalSatisfaction = 0;
        this.currentData.guests.forEach(guest => {
            totalSatisfaction += guest.satisfaction;
        });

        this.currentData.satisfaction = Math.round(totalSatisfaction / this.currentData.guestCount);
    }

    // Hírnév frissítése
    updateReputation() {
        const satisfactionBonus = (this.currentData.satisfaction - 50) * 0.1;
        const serviceBonus = this.getServiceBonus();
        
        this.currentData.reputation = Math.max(0, Math.min(100, 
            this.currentData.reputation + satisfactionBonus + serviceBonus
        ));
    }

    getServiceBonus() {
        let bonus = 0;
        if (this.currentData.services.cafe) bonus += 5;
        if (this.currentData.services.gym) bonus += 5;
        if (this.currentData.services.wifi) bonus += 2;
        if (this.currentData.services.cleaning) bonus += 3;
        return bonus;
    }

    // Nap frissítése
    nextDay() {
        this.currentData.day++;
        
        // Napi költségek
        const dailyExpenses = 
            this.currentData.expenses.utilities +
            this.currentData.expenses.cleaning +
            this.currentData.expenses.maintenance;
        
        this.spendMoney(dailyExpenses);
        
        // Vendégek frissítése
        this.updateGuests();
        
        // Elégedettség és hírnév frissítése
        this.updateSatisfaction();
        this.updateReputation();
    }

    updateGuests() {
        this.currentData.guests.forEach(guest => {
            guest.stayDuration--;
            
            // Vendég távozása
            if (guest.stayDuration <= 0) {
                this.removeGuest(guest.id);
            }
        });
    }

    // Szolgáltatás kezelés
    toggleService(serviceName) {
        if (this.currentData.services.hasOwnProperty(serviceName)) {
            this.currentData.services[serviceName] = !this.currentData.services[serviceName];
            return true;
        }
        return false;
    }

    // Ár beállítása
    setPrice(roomType, price) {
        if (this.currentData.prices.hasOwnProperty(roomType)) {
            this.currentData.prices[roomType] = Math.max(0, price);
            return true;
        }
        return false;
    }

    // Teljesítmény kezelés
    unlockAchievement(achievementId) {
        if (!this.currentData.achievements.includes(achievementId)) {
            this.currentData.achievements.push(achievementId);
            return true;
        }
        return false;
    }

    // Játék mentése
    saveGame(slotName = 'autosave') {
        try {
            const saveData = {
                timestamp: Date.now(),
                data: this.currentData
            };
            
            localStorage.setItem(`hotel_manager_save_${slotName}`, JSON.stringify(saveData));
            console.log(`Játék mentve: ${slotName}`);
            return true;
        } catch (error) {
            console.error('Hiba a mentés során:', error);
            return false;
        }
    }

    // Játék betöltése
    loadGame(slotName = 'autosave') {
        try {
            const saveData = localStorage.getItem(`hotel_manager_save_${slotName}`);
            if (saveData) {
                const parsed = JSON.parse(saveData);
                this.currentData = { ...this.defaultData, ...parsed.data };
                console.log(`Játék betöltve: ${slotName}`);
                return true;
            }
            return false;
        } catch (error) {
            console.error('Hiba a betöltés során:', error);
            return false;
        }
    }

    // Új játék
    newGame() {
        this.currentData = { ...this.defaultData };
        console.log('Új játék kezdve');
    }

    // Mentési slotok listázása
    getSaveSlots() {
        const slots = [];
        for (let i = 0; i < localStorage.length; i++) {
            const key = localStorage.key(i);
            if (key && key.startsWith('hotel_manager_save_')) {
                const slotName = key.replace('hotel_manager_save_', '');
                try {
                    const saveData = JSON.parse(localStorage.getItem(key));
                    slots.push({
                        name: slotName,
                        timestamp: saveData.timestamp,
                        hotelName: saveData.data.hotelName,
                        day: saveData.data.day,
                        money: saveData.data.money
                    });
                } catch (error) {
                    console.warn(`Hibás mentési slot: ${slotName}`);
                }
            }
        }
        return slots.sort((a, b) => b.timestamp - a.timestamp);
    }

    // Mentési slot törlése
    deleteSaveSlot(slotName) {
        try {
            localStorage.removeItem(`hotel_manager_save_${slotName}`);
            console.log(`Mentési slot törölve: ${slotName}`);
            return true;
        } catch (error) {
            console.error('Hiba a törlés során:', error);
            return false;
        }
    }

    // Játék statisztikák
    getStatistics() {
        return {
            totalGuests: this.currentData.totalGuests,
            totalEarnings: this.currentData.totalEarnings,
            averageSatisfaction: this.currentData.satisfaction,
            reputation: this.currentData.reputation,
            daysPlayed: this.currentData.day,
            roomsOwned: this.currentData.rooms.length,
            achievementsUnlocked: this.currentData.achievements.length
        };
    }

    // Játék állapot ellenőrzése
    isGameOver() {
        return this.currentData.money < -1000; // Ha túl sok adósság
    }

    // Játék nyertes állapot
    isGameWon() {
        return this.currentData.reputation >= 95 && this.currentData.money >= 100000;
    }
} 