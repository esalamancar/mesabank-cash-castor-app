// Tipos que reflejan los schemas de docs/08-openapi.yaml. Se actualizan a
// mano junto con la spec (T-07).

export interface User {
  id: string
  username: string
  is_guest: boolean
  created_at: string
}

export interface AuthResponse {
  token: string
  refresh_token: string
  user: User
}

export interface GameConfig {
  currency: string
  exchange_rate?: number
  /** RF-15: reservado, sin efecto funcional en el MVP. */
  inflation?: number | null
  fee?: number
  interest_type?: 'per_round' | 'per_time' | null
  interest_rate?: number | null
  interest_interval?: number | null
  max_debt_multiplier?: number
  total_digital: number
  total_paper: number
  initial_per_player: number
}

export interface Game {
  id: string
  code: string
  qr_url?: string
  link?: string
  status: 'lobby' | 'active' | 'finished'
  created_by: string
  config: GameConfig
}

export interface Player {
  id: string
  game_id: string
  user_id: string
  role: 'bank' | 'player' | 'spectator'
  digital_balance: number
  paper_balance: number
  debt: number
  is_bankrupt: boolean
}

export interface ApiErrorBody {
  error: string
  details?: string
}
