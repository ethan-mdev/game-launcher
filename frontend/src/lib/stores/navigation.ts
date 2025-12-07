import { writable } from 'svelte/store';

export type Page = 'play' | 'news' | 'settings';

export const currentPage = writable<Page>('play');