import { createApp } from './app';

const { PORT = 8000 } = process.env;

createApp()
  .then(app => {
    app.listen(Number(PORT), err => {
      if (err) {
        throw err;
      }
      console.log(`server running on: ${PORT}`);
    });
  })
  .catch(console.error);
