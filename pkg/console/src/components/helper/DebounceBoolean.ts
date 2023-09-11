export default class DebounceBoolean {
  _value = false;
  timerId = -1;

  get value() {
    return this._value;
  }

  set value(v: boolean) {
    if (this.timerId >= 0) {
      clearTimeout(this.timerId);
      this.timerId = -1;
    }
    if (v) {
      this._value = v;
      return;
    }
    this.timerId = window.setTimeout(() => {
      this._value = v;
    }, 100);
  }
}
