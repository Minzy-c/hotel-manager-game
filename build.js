#!/usr/bin/env node

const { execSync } = require('child_process');
const fs = require('fs');
const path = require('path');

console.log('🏨 Hotel Manager - Build Script');
console.log('===============================\n');

// Ellenőrizzük, hogy a szükséges fájlok léteznek-e
const requiredFiles = [
    'package.json',
    'main.js',
    'index.html',
    'src/AssetManager.js',
    'src/GameData.js'
];

console.log('Fájlok ellenőrzése...');
for (const file of requiredFiles) {
    if (!fs.existsSync(file)) {
        console.error(`❌ Hiányzó fájl: ${file}`);
        process.exit(1);
    }
}
console.log('✅ Minden szükséges fájl megtalálható\n');

// NPM install
console.log('Függőségek telepítése...');
try {
    execSync('npm install', { stdio: 'inherit' });
    console.log('✅ Függőségek telepítve\n');
} catch (error) {
    console.error('❌ Hiba a függőségek telepítése során');
    process.exit(1);
}

// Build típus kiválasztása
const buildType = process.argv[2] || 'dev';

if (buildType === 'windows') {
    console.log('Windows build létrehozása...');
    try {
        execSync('npm run build:windows', { stdio: 'inherit' });
        console.log('✅ Windows build elkészült\n');
    } catch (error) {
        console.error('❌ Hiba a Windows build során');
        process.exit(1);
    }
} else if (buildType === 'web') {
    console.log('Web build létrehozása...');
    try {
        execSync('npm run build', { stdio: 'inherit' });
        console.log('✅ Web build elkészült\n');
    } catch (error) {
        console.error('❌ Hiba a web build során');
        process.exit(1);
    }
} else {
    console.log('Fejlesztői szerver indítása...');
    try {
        execSync('npm run dev', { stdio: 'inherit' });
    } catch (error) {
        console.error('❌ Hiba a fejlesztői szerver indítása során');
        process.exit(1);
    }
}

console.log('🎉 Build folyamat befejezve!'); 