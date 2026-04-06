import { useState } from 'react'
import './App.css'

function App() {
  const [tableId, setTableId] = useState(null)
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState(null)

  const createGame = async () => {
    setLoading(true)
    setError(null)
    try {
      const res = await fetch('/api/create', { method: 'POST' })
      if (!res.ok) throw new Error('Failed to create game')
      const data = await res.json()
      setTableId(data.table_id)
    } catch (err) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="container">
      <h1>Degen Poker</h1>

      {!tableId ? (
        <div className="home">
          <p>Create a new poker table to get started</p>
          <button onClick={createGame} disabled={loading}>
            {loading ? 'Creating...' : 'Create Game'}
          </button>
          {error && <p className="error">{error}</p>}
        </div>
      ) : (
        <div className="created">
          <p>Table created!</p>
          <code className="table-id">{tableId}</code>
          <p className="hint">Share this code with friends to join</p>
          <button onClick={() => setTableId(null)}>Create Another</button>
        </div>
      )}
    </div>
  )
}

export default App
