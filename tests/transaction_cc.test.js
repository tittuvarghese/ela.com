const test = require('ava');
const CONSTANTS = require('../constants.js');
const request = require('supertest');
const app = require('../app.js');

test("Request to create transaction records targeting peers from " + CONSTANTS.ORG1_NAME + " Org", async t => {
  const response = await request(app)
  .post('/channels/' + CONSTANTS.CHANNEL_NAME + '/chaincodes/' + CONSTANTS.CHAINCODE_NAME + '/' + CONSTANTS.FCN_CREATE_TXN)
  .set('Authorization', 'Bearer ' + CONSTANTS.ORG1_JWT)
  .send({
    peers: ["peer0.org1.ela.com"],
    args: ["ELP-TXN001", {
      id : "ELP-TXN001",
      buyer_id: "ELP-USR001",
      buyer_name : "Asha",
      seller_id : "ELP-USR002",
      seller_name : "Elton",
      product_id : "ELP-PRODUCT001",
      product_price: "10",
      quantity : 10,
      unit : "Nos",
      buyer_type : "Preserver Buy"
    }]
  });
  t.is(response.status, 200);
  t.is(response.body.success, true);
  t.is(response.body.message, 'Successfully invoked the chaincode for function '+ CONSTANTS.FCN_CREATE_TXN +' to the channel ' + CONSTANTS.CHANNEL_NAME);
});

test("Query transaction records from peer of " + CONSTANTS.ORG2_NAME + "Org", async t => {
  const response = await request(app)
  .get('/channels/' + CONSTANTS.CHANNEL_NAME + '/chaincodes/' + CONSTANTS.CHAINCODE_NAME + '/' + CONSTANTS.FCN_QUERY_TXN + '/ELP-TXN001?peer=peer0.org2.ela.com')
  .set('Authorization', 'Bearer ' + CONSTANTS.ORG2_JWT);
  t.is(response.status, 200);
  t.is(response.body.success, true);
  t.is(response.body.message, 'Found results for '+ CONSTANTS.FCN_QUERY_TXN +' - ELP-TXN001');
  t.deepEqual(response.body, {
    success: true,
    message: "Found results for queryTransaction - ELP-TXN001",
    data: {
      id: "ELP-TXN001",
      buyer_type: "Preserver Buy",
      buyer_id: "ELP-USR001",
      buyer_name: "Asha",
      seller_id: "ELP-USR002",
      seller_name: "Elton",
      product_id: "ELP-PRODUCT001",
      product_price: "10",
      quantity: 10,
      unit: "Nos",
      status: false,
      timestamp: "0001-01-01T00:00:00Z"
    }
  });
});

test("Request to update transaction records targeting peers from " + CONSTANTS.ORG1_NAME + " Org", async t => {
  const response = await request(app)
  .post('/channels/' + CONSTANTS.CHANNEL_NAME + '/chaincodes/' + CONSTANTS.CHAINCODE_NAME + '/' + CONSTANTS.FCN_UPDATE_TXN)
  .set('Authorization', 'Bearer ' + CONSTANTS.ORG1_JWT)
  .send({
    peers: ["peer0.org1.ela.com"],
    args: ["ELP-TXN001", {
      status: true
    }]
  });
  t.is(response.status, 200);
  t.is(response.body.success, true);
  t.is(response.body.message, 'Successfully invoked the chaincode for function '+ CONSTANTS.FCN_UPDATE_TXN +' to the channel ' + CONSTANTS.CHANNEL_NAME);
});

test("Query transaction records from peer of " + CONSTANTS.ORG2_NAME + "Org to check update status", async t => {
  const response = await request(app)
  .get('/channels/' + CONSTANTS.CHANNEL_NAME + '/chaincodes/' + CONSTANTS.CHAINCODE_NAME + '/' + CONSTANTS.FCN_QUERY_TXN + '/ELP-TXN001?peer=peer0.org2.ela.com')
  .set('Authorization', 'Bearer ' + CONSTANTS.ORG2_JWT);
  t.is(response.status, 200);
  t.is(response.body.success, true);
  t.is(response.body.message, 'Found results for '+ CONSTANTS.FCN_QUERY_TXN +' - ELP-TXN001');
  t.deepEqual(response.body, {
    success: true,
    message: "Found results for queryTransaction - ELP-TXN001",
    data: {
      id: "ELP-TXN001",
      buyer_type: "Preserver Buy",
      buyer_id: "ELP-USR001",
      buyer_name: "Asha",
      seller_id: "ELP-USR002",
      seller_name: "Elton",
      product_id: "ELP-PRODUCT001",
      product_price: "10",
      quantity: 10,
      unit: "Nos",
      status: true,
      timestamp: "0001-01-01T00:00:00Z"
    }
  });
});
