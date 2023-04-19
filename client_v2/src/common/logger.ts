
// This code defines a Logger class with three static methods:
// Error, Info, and Warning. Each method takes two string parameters - group and
// body - and logs a message in the console with a specific formatting.
// The Error method logs an error message using the console.error
// function, with the message formatted as "[group] body".
//
// The Info method logs an informational message using
// the console.log function, with the message formatted as "[group] body".
//
// The Warning method logs a warning message using the console.warn
// function, with the message formatted as "[group] body".
//
// The use of static keyword before the method names means
// you don't need to create an instance of Logger to use these
// methods, you can access them directly on the class itself.
//
// For example: Logger.Error('API', 'Connection failed')
// will log "[API] Connection failed"
// in the console using console.error.
export class Logger {
  static Error<T>(group: string, body: T) {
    console.error(`[${group}]`, body)
  }

  static Info<T>(group: string, body: T) {
    console.log(`[${group}]`, body)
  }

  static Warring<T>(group: string, body: T) {
    console.warn(`[${group}]`, body)
  }
}