const test = require('ava');
const CONSTANTS = require('../constants.js');
const request = require('supertest');
const app = require('../app.js');

test("Request to create channel with valid parameters and authorization token", async t => {
  const response = await request(app)
  .post('/channels')
  .set('Authorization', 'Bearer ' + CONSTANTS.ORG1_JWT)
  .send({channelName: CONSTANTS.CHANNEL_NAME, channelConfigPath: CONSTANTS.CHANNEL_CONFIG_PATH });
  t.is(response.status, 200);
  t.deepEqual(response.body, {
    success: true,
    message: "Channel " + CONSTANTS.CHANNEL_NAME + " created Successfully"
  });
});

test("Request to join with valid peers (peer 0 and peer1) from " + CONSTANTS.ORG1_NAME, async t => {
  const response = await request(app)
  .post('/channels/' + CONSTANTS.CHANNEL_NAME + '/peers/')
  .set('Authorization', 'Bearer ' + CONSTANTS.ORG1_JWT)
  .send({peers: ['peer0.org1.ela.com', 'peer1.org1.ela.com']});
  t.is(response.status, 200);
  t.deepEqual(response.body, {
    success: true,
    message: "Successfully joined peers in organization " + CONSTANTS.ORG1_NAME + " to the channel: " + CONSTANTS.CHANNEL_NAME
  });
});

test("Request to join with valid peers (peer 0 and peer1) from " + CONSTANTS.ORG2_NAME, async t => {
  const response = await request(app)
  .post('/channels/' + CONSTANTS.CHANNEL_NAME + '/peers/')
  .set('Authorization', 'Bearer ' + CONSTANTS.ORG2_JWT)
  .send({peers: ['peer0.org2.ela.com', 'peer1.org2.ela.com']});
  t.is(response.status, 200);
  t.deepEqual(response.body, {
    success: true,
    message: "Successfully joined peers in organization " + CONSTANTS.ORG2_NAME + " to the channel: " + CONSTANTS.CHANNEL_NAME
  });
});

test("Request to update anchor peer with valid anchor configuration - " + CONSTANTS.ORG1_NAME, async t => {
  const response = await request(app)
  .post('/channels/' + CONSTANTS.CHANNEL_NAME + '/anchorpeers')
  .set('Authorization', 'Bearer ' + CONSTANTS.ORG1_JWT)
  .send({
    configUpdatePath : CONSTANTS.ANCHOR_CONFIG_PATH_ORG1
  });
  t.is(response.status, 200);
  t.is(response.body.success, true);
  t.deepEqual(response.body, {
    "success": true,
    "message": "Successfully update anchor peers in organization " + CONSTANTS.ORG1_NAME + " to the channel " + CONSTANTS.CHANNEL_NAME
  });
});

test("Request to update anchor peer with valid anchor configuration - " + CONSTANTS.ORG2_NAME, async t => {
  const response = await request(app)
  .post('/channels/' + CONSTANTS.CHANNEL_NAME + '/anchorpeers')
  .set('Authorization', 'Bearer ' + CONSTANTS.ORG2_JWT)
  .send({
    configUpdatePath : CONSTANTS.ANCHOR_CONFIG_PATH_ORG2
  });
  t.is(response.status, 200);
  t.is(response.body.success, true);
  t.deepEqual(response.body, {
    "success": true,
    "message": "Successfully update anchor peers in organization " + CONSTANTS.ORG2_NAME + " to the channel " + CONSTANTS.CHANNEL_NAME
  });
});
