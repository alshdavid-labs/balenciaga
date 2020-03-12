export class EventEmitter {
  subscribers = {}
  resolveComplete
  completePromise = new Promise(res => this.resolveComplete = res)
  hasComplete = false

  toPromise() {
    return this.completePromise
  }

  subscribe(cb) {
    const key = (Math.random() * 1000000000000000).toFixed().toString()
    this.subscribers[key] = cb
    return {
      unsubscribe: () => delete this.subscribers[key]
    }
  }

  emit(value) {
    if (this.hasComplete) {
      throw new Error('Cannot next on complete subject')
    }
    for (const key of Object.keys(this.subscribers)) {
      this.subscribers[key](value)
    }
  }

  complete() {
    this.hasComplete = true
    this.resolveComplete()
  }
}