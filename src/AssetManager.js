// Asset Manager - Asset betöltés és kezelés
export class AssetManager {
    constructor() {
        this.assets = {
            interiors: {},
            characters: {},
            animated: {},
            ui: {},
            homeDesigns: {}
        };
        this.loadedAssets = 0;
        this.totalAssets = 0;
    }

    async loadAllAssets() {
        console.log('Asset betöltés kezdése...');
        
        try {
            // Interior assets betöltése
            await this.loadInteriorAssets();
            
            // Character assets betöltése
            await this.loadCharacterAssets();
            
            // Animated objects betöltése
            await this.loadAnimatedAssets();
            
            // UI elements betöltése
            await this.loadUIAssets();
            
            // Home designs betöltése
            await this.loadHomeDesignAssets();
            
            console.log('Minden asset sikeresen betöltve!');
            return true;
        } catch (error) {
            console.error('Hiba az asset betöltés során:', error);
            return false;
        }
    }

    async loadInteriorAssets() {
        const interiorTypes = ['16x16', '32x32', '48x48'];
        
        for (const size of interiorTypes) {
            try {
                // Interior spritesheet betöltése
                const interiorPath = `assets/1_Interiors/${size}/Interiors_${size}.png`;
                this.assets.interiors[`interiors_${size}`] = await this.loadImage(interiorPath);
                
                // Room builder spritesheet betöltése
                const roomBuilderPath = `assets/1_Interiors/${size}/Room_Builder_${size}.png`;
                this.assets.interiors[`room_builder_${size}`] = await this.loadImage(roomBuilderPath);
                
                this.loadedAssets += 2;
            } catch (error) {
                console.warn(`Nem sikerült betölteni a ${size} interior asseteket:`, error);
            }
        }
    }

    async loadCharacterAssets() {
        try {
            // Karakter generátor assets betöltése
            const characterSizes = ['16x16', '32x32', '48x48'];
            
            for (const size of characterSizes) {
                // Bodies
                const bodiesPath = `assets/2_Characters/Character_Generator/Bodies/Bodies_${size}.png`;
                this.assets.characters[`bodies_${size}`] = await this.loadImage(bodiesPath);
                
                // Hairstyles
                const hairPath = `assets/2_Characters/Character_Generator/Hairstyles/Hairstyles_${size}.png`;
                this.assets.characters[`hairstyles_${size}`] = await this.loadImage(hairPath);
                
                // Outfits
                const outfitPath = `assets/2_Characters/Character_Generator/Outfits/Outfits_${size}.png`;
                this.assets.characters[`outfits_${size}`] = await this.loadImage(outfitPath);
                
                // Eyes
                const eyesPath = `assets/2_Characters/Character_Generator/Eyes/Eyes_${size}.png`;
                this.assets.characters[`eyes_${size}`] = await this.loadImage(eyesPath);
                
                this.loadedAssets += 4;
            }
        } catch (error) {
            console.warn('Nem sikerült betölteni a karakter asseteket:', error);
        }
    }

    async loadAnimatedAssets() {
        try {
            const animatedSizes = ['16x16', '32x32', '48x48'];
            
            for (const size of animatedSizes) {
                // Ajtók
                const doorPath = `assets/3_Animated_objects/${size}/spritesheets/animated_door_1_${size}.png`;
                this.assets.animated[`door_${size}`] = await this.loadImage(doorPath);
                
                // Fürdőkád
                const bathtubPath = `assets/3_Animated_objects/${size}/spritesheets/animated_bathtub_${size}.png`;
                this.assets.animated[`bathtub_${size}`] = await this.loadImage(bathtubPath);
                
                // Csap
                const sinkPath = `assets/3_Animated_objects/${size}/spritesheets/animated_bathroom_sink_new_3-10 loop_${size}.png`;
                this.assets.animated[`sink_${size}`] = await this.loadImage(sinkPath);
                
                this.loadedAssets += 3;
            }
        } catch (error) {
            console.warn('Nem sikerült betölteni az animált asseteket:', error);
        }
    }

    async loadUIAssets() {
        try {
            const uiSizes = ['16x16', '32x32', '48x48'];
            
            for (const size of uiSizes) {
                // UI spritesheet
                const uiPath = `assets/4_User_Interface_Elements/UI_${size}.png`;
                this.assets.ui[`ui_${size}`] = await this.loadImage(uiPath);
                
                // Thinking emotes
                const emotesPath = `assets/4_User_Interface_Elements/UI_thinking_emotes_animation_${size}.png`;
                this.assets.ui[`emotes_${size}`] = await this.loadImage(emotesPath);
                
                this.loadedAssets += 2;
            }
        } catch (error) {
            console.warn('Nem sikerült betölteni a UI asseteket:', error);
        }
    }

    async loadHomeDesignAssets() {
        try {
            // Generic home designs
            const homePath = `assets/6_Home_Designs/Generic_Home_Designs/32x32/Generic_Home_1_Layer_1_32x32.png`;
            this.assets.homeDesigns['generic_home_1'] = await this.loadImage(homePath);
            
            this.loadedAssets += 1;
        } catch (error) {
            console.warn('Nem sikerült betölteni a home design asseteket:', error);
        }
    }

    async loadImage(path) {
        return new Promise((resolve, reject) => {
            const img = new Image();
            img.onload = () => resolve(img);
            img.onerror = () => reject(new Error(`Nem sikerült betölteni: ${path}`));
            img.src = path;
        });
    }

    getAsset(category, name) {
        return this.assets[category]?.[name] || null;
    }

    getInteriorAsset(name) {
        return this.getAsset('interiors', name);
    }

    getCharacterAsset(name) {
        return this.getAsset('characters', name);
    }

    getAnimatedAsset(name) {
        return this.getAsset('animated', name);
    }

    getUIAsset(name) {
        return this.getAsset('ui', name);
    }

    getHomeDesignAsset(name) {
        return this.getAsset('homeDesigns', name);
    }

    getLoadingProgress() {
        return this.totalAssets > 0 ? (this.loadedAssets / this.totalAssets) * 100 : 0;
    }

    // Sprite sheet kezelés
    getSpriteFromSheet(sheet, x, y, width, height) {
        if (!sheet) return null;
        
        const canvas = document.createElement('canvas');
        const ctx = canvas.getContext('2d');
        canvas.width = width;
        canvas.height = height;
        
        ctx.drawImage(sheet, x * width, y * height, width, height, 0, 0, width, height);
        return canvas;
    }

    // Karakter generálás
    generateCharacter(size = '32x32') {
        const bodies = this.getCharacterAsset(`bodies_${size}`);
        const hairstyles = this.getCharacterAsset(`hairstyles_${size}`);
        const outfits = this.getCharacterAsset(`outfits_${size}`);
        const eyes = this.getCharacterAsset(`eyes_${size}`);
        
        if (!bodies || !hairstyles || !outfits || !eyes) {
            console.warn('Hiányzó karakter assetek');
            return null;
        }
        
        // Véletlenszerű karakter összeállítása
        const bodyIndex = Math.floor(Math.random() * 4); // 0-3
        const hairIndex = Math.floor(Math.random() * 8); // 0-7
        const outfitIndex = Math.floor(Math.random() * 6); // 0-5
        const eyeIndex = Math.floor(Math.random() * 4); // 0-3
        
        const tileSize = parseInt(size);
        
        const canvas = document.createElement('canvas');
        const ctx = canvas.getContext('2d');
        canvas.width = tileSize;
        canvas.height = tileSize;
        
        // Rétegek rajzolása sorrendben
        ctx.drawImage(bodies, bodyIndex * tileSize, 0, tileSize, tileSize, 0, 0, tileSize, tileSize);
        ctx.drawImage(hairstyles, hairIndex * tileSize, 0, tileSize, tileSize, 0, 0, tileSize, tileSize);
        ctx.drawImage(outfits, outfitIndex * tileSize, 0, tileSize, tileSize, 0, 0, tileSize, tileSize);
        ctx.drawImage(eyes, eyeIndex * tileSize, 0, tileSize, tileSize, 0, 0, tileSize, tileSize);
        
        return canvas;
    }
} 