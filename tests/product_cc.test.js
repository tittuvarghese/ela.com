const test = require('ava');
const CONSTANTS = require('../constants.js');
const request = require('supertest');
const app = require('../app.js');

test("Request to create product records targeting peers from " + CONSTANTS.ORG1_NAME + " Org", async t => {
  const response = await request(app)
  .post('/channels/' + CONSTANTS.CHANNEL_NAME + '/chaincodes/' + CONSTANTS.CHAINCODE_NAME + '/' + CONSTANTS.FCN_CREATE_PRODUCT)
  .set('Authorization', 'Bearer ' + CONSTANTS.ORG1_JWT)
  .send({
    peers: ["peer0.org1.ela.com"],
    args: ["ELP-PRODUCT001", {
      id : "ELP-PRODUCT001",
      name: "Tomato",
      quantity : 10,
      unit : "Nos",
      price : "10",
      date_of_harvest : "27-05-2019",
      farmer_id: "ELP-USR001"
    }]
  });
  t.is(response.status, 200);
  t.is(response.body.success, true);
  t.is(response.body.message, 'Successfully invoked the chaincode for function '+ CONSTANTS.FCN_CREATE_PRODUCT +' to the channel ' + CONSTANTS.CHANNEL_NAME);
});

test("Query product records from peer of " + CONSTANTS.ORG2_NAME + "Org", async t => {
  const response = await request(app)
  .get('/channels/' + CONSTANTS.CHANNEL_NAME + '/chaincodes/' + CONSTANTS.CHAINCODE_NAME + '/' + CONSTANTS.FCN_QUERY_PRODUCT + '/ELP-PRODUCT001?peer=peer0.org2.ela.com')
  .set('Authorization', 'Bearer ' + CONSTANTS.ORG2_JWT);
  t.is(response.status, 200);
  t.is(response.body.success, true);
  t.is(response.body.message, 'Found results for '+ CONSTANTS.FCN_QUERY_PRODUCT +' - ELP-PRODUCT001');
  t.deepEqual(response.body, {
    success: true,
    message: "Found results for queryProduct - ELP-PRODUCT001",
    data: {
      id: "ELP-PRODUCT001",
      name: "Tomato",
      quantity: 10,
      unit: "Nos",
      price: "10",
      date_of_harvest: "27-05-2019",
      weight_left: 10,
      farmer_id: "ELP-USR001"
    }
  });
});

test("Request to update product records targeting peers from " + CONSTANTS.ORG1_NAME + " Org", async t => {
  const response = await request(app)
  .post('/channels/' + CONSTANTS.CHANNEL_NAME + '/chaincodes/' + CONSTANTS.CHAINCODE_NAME + '/' + CONSTANTS.FCN_UPDATE_PRODUCT)
  .set('Authorization', 'Bearer ' + CONSTANTS.ORG1_JWT)
  .send({
    peers: ["peer0.org1.ela.com"],
    args: ["ELP-PRODUCT001", {
      quantity : 15
    }]
  });
  t.is(response.status, 200);
  t.is(response.body.success, true);
  t.is(response.body.message, 'Successfully invoked the chaincode for function '+ CONSTANTS.FCN_UPDATE_PRODUCT +' to the channel ' + CONSTANTS.CHANNEL_NAME);
});

test("Query product records from peer of " + CONSTANTS.ORG2_NAME + "Org to check update status", async t => {
  const response = await request(app)
  .get('/channels/' + CONSTANTS.CHANNEL_NAME + '/chaincodes/' + CONSTANTS.CHAINCODE_NAME + '/' + CONSTANTS.FCN_QUERY_PRODUCT + '/ELP-PRODUCT001?peer=peer0.org2.ela.com')
  .set('Authorization', 'Bearer ' + CONSTANTS.ORG2_JWT);
  t.is(response.status, 200);
  t.is(response.body.success, true);
  t.is(response.body.message, 'Found results for '+ CONSTANTS.FCN_QUERY_PRODUCT +' - ELP-PRODUCT001');
  t.deepEqual(response.body, {
    success: true,
    message: "Found results for queryProduct - ELP-PRODUCT001",
    data: {
      id: "ELP-PRODUCT001",
      name: "Tomato",
      quantity: 25,
      unit: "Nos",
      price: "10",
      date_of_harvest: "27-05-2019",
      weight_left: 25,
      farmer_id: "ELP-USR001"
    }
  });
});

test("Query history of product from peer of " + CONSTANTS.ORG2_NAME + "Org to check update status", async t => {
  const response = await request(app)
  .get('/channels/' + CONSTANTS.CHANNEL_NAME + '/chaincodes/' + CONSTANTS.CHAINCODE_NAME + '/' + CONSTANTS.FCN_PRODUCT_HISTORY + '/ELP-PRODUCT001?peer=peer0.org2.ela.com')
  .set('Authorization', 'Bearer ' + CONSTANTS.ORG2_JWT);
  t.is(response.status, 200);
  t.is(response.body.success, true);
  t.is(response.body.message, 'Found results for '+ CONSTANTS.FCN_PRODUCT_HISTORY +' - ELP-PRODUCT001');
  t.deepEqual(response.body, {
    success: true,
    message: "Found results for getHistory - ELP-PRODUCT001",
    data: {
      id: "ELP-PRODUCT001",
      name: "Tomato",
      quantity: 25,
      unit: "Nos",
      price: "10",
      date_of_harvest: "27-05-2019",
      weight_left: 25,
      farmer_id: "ELP-USR001"
    }
  });
});
