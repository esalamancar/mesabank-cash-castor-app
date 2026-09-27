import { create } from 'zustand'

// Store de referencia (T-02): todavía sin lógica real, solo el patrón que
// van a seguir los stores de la app a partir de Sprint 1 (sesión, partida
// activa, etc.).
interface SessionState {
  token: string | null
  setToken: (token: string | null) => void
}

export const useSessionStore = create<SessionState>((set) => ({
  token: null,
  setToken: (token) => set({ token }),
}))
