import winston = require('winston');

const logger = winston.createLogger({
  format: winston.format.printf(info => {
    let logData;
    if (info.error instanceof Error) {
      logData = {
        error: {
          message: info.error.message,
          stack: info.error.stack
        },
        severity: info.level.toUpperCase(),
        level: info.level,
        message: info.message
      };
    } else {
      logData = {
        severity: info.level.toUpperCase(),
        level: info.level,
        message: info.message
      };
    }

    return JSON.stringify(logData, null, 0);
  }),
  transports: [new winston.transports.Console({ level: 'info' })]
});

export default logger;
