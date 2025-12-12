// Avatar utility functions
const AVATAR_COUNT = 20;
const DEFAULT_AVATAR = 'avatar-1.png';

/**
 * Get the avatar source URL from a filename
 * @param filename - Avatar filename like "avatar-1.png" or null/undefined
 * @returns The avatar image source URL
 */
export function getAvatarSrc(filename: string | null | undefined): string {
  const avatarName = filename || DEFAULT_AVATAR;
  
  try {
    return new URL(`../../assets/images/${avatarName}`, import.meta.url).href;
  } catch {
    return new URL(`../../assets/images/${DEFAULT_AVATAR}`, import.meta.url).href;
  }
}

/**
 * Get all available avatar options
 * @returns Array of avatar filenames
 */
export function getAvailableAvatars(): string[] {
  return Array.from({ length: AVATAR_COUNT }, (_, i) => `avatar-${i + 1}.png`);
}

