import { request } from './client'
import type { Game, GameConfig, Player } from './types'

export function createGame(config: GameConfig): Promise<Game> {
  return request<Game>('/games', { method: 'POST', body: { config } })
}

export function getGame(code: string): Promise<Game> {
  return request<Game>(`/games/${code}`)
}

export function joinGame(
  code: string,
  role: 'player' | 'spectator',
): Promise<Player> {
  return request<Player>(`/games/${code}/join`, {
    method: 'POST',
    body: { role },
  })
}

export function listPlayers(code: string): Promise<Player[]> {
  return request<Player[]>(`/games/${code}/players`)
}

export function expelPlayer(code: string, playerId: string): Promise<Player> {
  return request<Player>(`/games/${code}/players/${playerId}/expel`, {
    method: 'POST',
  })
}

export interface ResetRequest {
  confirmations_required: number
  confirmations_received: number
  executed: boolean
}

export function requestReset(code: string): Promise<ResetRequest> {
  return request<ResetRequest>(`/games/${code}/reset`, { method: 'POST' })
}

export function confirmReset(code: string): Promise<ResetRequest> {
  return request<ResetRequest>(`/games/${code}/reset/confirm`, {
    method: 'POST',
  })
}
