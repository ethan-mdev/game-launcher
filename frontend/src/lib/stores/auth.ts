import { writable } from 'svelte/store';

export interface AuthState {
  isLoggedIn: boolean;
  userId: string;
  username: string;
  role: string;
  profileImage: string;
  accessToken: string;
  refreshToken: string;
  gameApiKey: string;
  gameLinked: boolean;
}

export const auth = writable<AuthState>({
  isLoggedIn: false,
  userId: '',
  username: '',
  role: 'user',
  profileImage: '',
  accessToken: '',
  refreshToken: '',
  gameApiKey: '',
  gameLinked: false
});