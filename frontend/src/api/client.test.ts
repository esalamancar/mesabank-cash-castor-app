import { afterEach, describe, expect, it, vi } from 'vitest'
import { ApiError, request } from './client'
import { useSessionStore } from '../store/useSessionStore'

function mockFetchOnce(status: number, body: unknown) {
  vi.stubGlobal(
    'fetch',
    vi.fn().mockResolvedValue({
      ok: status >= 200 && status < 300,
      status,
      json: async () => body,
    }),
  )
}

describe('request', () => {
  afterEach(() => {
    vi.unstubAllGlobals()
    useSessionStore.getState().setToken(null)
  })

  it('devuelve el body parseado cuando la respuesta es exitosa', async () => {
    mockFetchOnce(200, { status: 'ok' })

    const result = await request<{ status: string }>('/healthz', {
      auth: false,
    })

    expect(result).toEqual({ status: 'ok' })
  })

  it('lanza ApiError con status y body cuando la respuesta falla', async () => {
    mockFetchOnce(401, { error: 'missing_or_invalid_token' })

    await expect(request('/users/me')).rejects.toMatchObject({
      status: 401,
      body: { error: 'missing_or_invalid_token' },
    } satisfies Partial<ApiError>)
  })

  it('agrega el header Authorization cuando hay token en el store', async () => {
    useSessionStore.getState().setToken('mi-token')
    mockFetchOnce(200, {})

    await request('/users/me')

    const fetchMock = fetch as unknown as ReturnType<typeof vi.fn>
    const [, options] = fetchMock.mock.calls[0]
    expect(options.headers.Authorization).toBe('Bearer mi-token')
  })
})
