const test = require('ava');
const CONSTANTS = require('../constants.js');
const request = require('supertest');
const app = require('../app.js');

test("Request to install Chaincode to with valid parameters on " + CONSTANTS.ORG1_NAME, async t => {
  const response = await request(app)
  .post('/chaincodes')
  .set('Authorization', 'Bearer ' + CONSTANTS.ORG1_JWT)
  .send({
	peers: ['peer0.org1.ela.com','peer1.org1.ela.com','peer0.org2.ela.com','peer1.org2.ela.com'],
	chaincodeName: CONSTANTS.CHAINCODE_NAME,
	chaincodePath: CONSTANTS.CHAINCODE_PATH,
	chaincodeType: CONSTANTS.CHAINCODE_TYPE,
	chaincodeVersion: CONSTANTS.CHAINCODE_VERSION
});
  t.is(response.status, 200);
  t.deepEqual(response.body, {
    success: true,
    message: "Successfully installed chaincode"
  });
});

test("Request to instantiate Chaincode with valid parameters on Org1", async t => {
  const response = await request(app)
  .post('/channels/'+ CONSTANTS.CHANNEL_NAME +'/chaincodes')
  .set('Authorization', 'Bearer ' + CONSTANTS.ORG1_JWT)
  .send({
    peers: ["peer0.org1.ela.com", "peer1.org1.ela.com"],
  	chaincodeName: CONSTANTS.CHAINCODE_NAME,
    chaincodeVersion: CONSTANTS.CHAINCODE_VERSION,
  	chaincodeType: CONSTANTS.CHAINCODE_TYPE,
    fcn: CONSTANTS.FCN_INIT
  });
  t.is(response.status, 200);
  t.deepEqual(response.body, {
    success: true,
    message: "Successfully instantiated chaincode in organization " + CONSTANTS.ORG1_NAME + " to the channel " + CONSTANTS.CHANNEL_NAME
  });
});
