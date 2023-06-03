export class Handler<TSuccessFn, TFailFn> {
  constructor(public sfn: TSuccessFn, public ffn: TFailFn)
    {}
}

export function Handle<T, Sfn, Ffn, THandle = Handler<Sfn, Ffn>>(value: T, handler: THandle) {
  if (value === undefined || value === null) {
    handler.ffn
  } else {
    handler.ffn()
  }
}