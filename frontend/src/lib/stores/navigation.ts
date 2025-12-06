import { writable } from 'svelte/store';

export type Page = 'play' | 'news' | 'patchnotes' | 'settings';

export const currentPage = writable<Page>('play');