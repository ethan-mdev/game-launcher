import { writable } from 'svelte/store';

export interface AuthState {
  isLoggedIn: boolean;
  userId: string;
  username: string;
  email: string;
  role: string;
  profileImage: string;
  accessToken: string;
  refreshToken: string;
}

export const auth = writable<AuthState>({
  isLoggedIn: false,
  userId: '',
  username: '',
  email: '',
  role: 'user',
  profileImage: '',
  accessToken: '',
  refreshToken: ''
});