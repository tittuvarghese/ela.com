/**
* Test for Enrolling User to Organizations
* @method POST
* @route /users
* @params username, orgName
*/

const test = require('ava');
const CONSTANTS = require('../constants.js');
const request = require('supertest');
const app = require('../app.js');

test("Enrolling user to " + CONSTANTS.ORG1_NAME + " Org to get authorization token", async t => {
  const response = await request(app)
  .post('/users')
  .send({username: 'user@'+CONSTANTS.ORG1_EMAIL, orgName: CONSTANTS.ORG1_NAME});
  t.is(response.status, 200);
  t.is(response.body.success, true);
});

test("Enrolling user to " + CONSTANTS.ORG2_NAME + " Org to get authorization token", async t => {
  const response = await request(app)
  .post('/users')
  .send({username: 'user@'+CONSTANTS.ORG2_EMAIL, orgName: CONSTANTS.ORG2_NAME});
  t.is(response.status, 200);
  t.is(response.body.success, true);
});
