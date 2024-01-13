enum LogLevel {
  silent = 0,
  trace = 1,
  debug = 2,
  info = 3,
  warn = 4,
  error = 5,
}

export interface Logger {
  debug(msg: string, context?: object): void;
  info(msg: string, context?: object): void;
  warn(msg: string, context?: object): void;
  error(msg: string, context?: object): void;
  setLogLevel(level: LogLevel): void;
}

class AppLogger implements Logger {
  constructor(private level: LogLevel) {}

  setLogLevel(level: LogLevel): void {
    this.level = level;
  }

  debug(msg: string, context?: object): void {
    if (this.level < LogLevel.debug) return;
    console.log("debug: ", msg, context);
  }
  info(msg: string, context?: object): void {
    if (this.level < LogLevel.info) return;
    console.log("info: ", msg, context);
  }
  warn(msg: string, context?: object): void {
    if (this.level < LogLevel.warn) return;
    console.warn("warn: ", msg, context);
  }
  error(msg: string, context?: object): void {
    if (this.level < LogLevel.error) return;
    console.error("error: ", msg, context);
  }
}

export const logger: Logger = new AppLogger(
  import.meta.env.PROD ? LogLevel.silent : LogLevel.error
);
