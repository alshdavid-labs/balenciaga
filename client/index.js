import { print } from './print.js'
import { WebSocketAdapter } from './websocket-adapter.js'

const ws = new WebSocketAdapter()

ws.onMessage$.subscribe(console.log)
ws.onMessage$.subscribe(print)

ws.onConnect$.subscribe(() => {
  // ws.send(JSON.stringify({ type: 'SUBSCRIBE', value: 'topic-a' }, null, 2))
  ws.send({ type: 'topic-a', value: 'show' })
  ws.send({ type: 'topic-b', value: 'hide' })
})

ws.connect('ws://localhost:8081')

window.ws = ws