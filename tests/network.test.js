const test = require('ava');
const CONSTANTS = require('../constants.js');
const request = require('supertest');
const app = require('../app.js');

/* Checking REST Server Status - REST Route */
test("Checking REST server status.", async t => {
  const response = await request(app)
  .get('/status');
  t.is(response.status, 200);
  t.deepEqual(response.body, {
    status : true,
    message: 'REST server is up and running'
  });
});
