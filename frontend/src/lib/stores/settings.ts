import { writable } from 'svelte/store';

export const resolution = writable('1920x1080');
export const windowed = writable(false);

export const masterVolume = writable(100);
export const musicVolume = writable(100);
export const sfxVolume = writable(100);