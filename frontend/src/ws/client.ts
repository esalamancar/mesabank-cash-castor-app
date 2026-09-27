// Cliente WebSocket con reconexión automática (backoff exponencial),
// soporta US-003 (recuperación de sesión) y US-080 (tiempo real).

const DEFAULT_BASE_URL = 'ws://localhost:8080'
const MAX_RECONNECT_DELAY_MS = 30_000

export class GameSocket {
  private readonly url: string
  private ws: WebSocket | null = null
  private reconnectAttempts = 0
  private shouldReconnect = false
  private readonly listeners = new Set<(event: MessageEvent) => void>()

  constructor(gameCode: string) {
    const base =
      (import.meta.env.VITE_WS_BASE_URL as string | undefined) ??
      DEFAULT_BASE_URL
    this.url = `${base}/games/${gameCode}/ws`
  }

  connect(): void {
    this.shouldReconnect = true
    this.open()
  }

  disconnect(): void {
    this.shouldReconnect = false
    this.ws?.close()
  }

  onMessage(listener: (event: MessageEvent) => void): () => void {
    this.listeners.add(listener)
    return () => this.listeners.delete(listener)
  }

  private open(): void {
    this.ws = new WebSocket(this.url)

    this.ws.addEventListener('open', () => {
      this.reconnectAttempts = 0
    })

    this.ws.addEventListener('message', (event) => {
      this.listeners.forEach((listener) => listener(event))
    })

    this.ws.addEventListener('close', () => {
      if (this.shouldReconnect) {
        this.scheduleReconnect()
      }
    })

    this.ws.addEventListener('error', () => {
      this.ws?.close()
    })
  }

  private scheduleReconnect(): void {
    const delay = Math.min(
      1000 * 2 ** this.reconnectAttempts,
      MAX_RECONNECT_DELAY_MS,
    )
    this.reconnectAttempts += 1
    setTimeout(() => {
      if (this.shouldReconnect) {
        this.open()
      }
    }, delay)
  }
}
