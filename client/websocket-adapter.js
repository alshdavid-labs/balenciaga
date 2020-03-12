import { EventEmitter } from "./event-emitter.js";

export class WebSocketAdapter {
  onConnect$ = new EventEmitter()
  onDisconnect$ = new EventEmitter()
  onMessage$ = new EventEmitter()
  onError$ = new EventEmitter()
  
  _connection
  
  connect(address) {
    const conn = new WebSocket(address)
    conn.onopen = () => this.onConnect$.emit()
    conn.onclose = () => this.onDisconnect$.emit()
    conn.onerror = (evt) => this.onError$.emit(evt)
    conn.onmessage = (evt) => this.onMessage$.emit(evt.data)
    this._connection = conn
  }

  disconnect() {
    if (!this._connection) {
      return
    }
    this._connection.close()
  }

  send(msg = '') {
    if (typeof msg === 'object') {
      msg = JSON.stringify(msg);
    }
    this._connection.send(msg)
  }
}