const test = require('ava');
const CONSTANTS = require('../constants.js');
const request = require('supertest');
const app = require('../app.js');

test("Request to create user records targeting peers from " + CONSTANTS.ORG1_NAME + " Org", async t => {
  const response = await request(app)
  .post('/channels/' + CONSTANTS.CHANNEL_NAME + '/chaincodes/' + CONSTANTS.CHAINCODE_NAME + '/' + CONSTANTS.FCN_CREATE_USER)
  .set('Authorization', 'Bearer ' + CONSTANTS.ORG1_JWT)
  .send({
    peers: ["peer6.patient.rxmed.com"],
    args: ["PL-USR001", {
      user_id : "PL-USR001",
      name: {
        first_name: "Adam",
        last_name: "Smith"
      },
      email : "user@propertylist.io",
      password : "qwerty123",
      country : "India",
      phone_number : "+91-9876543210",
      profile_image : "profile_image_url_ipfs",
      role : "buyer"
    }]
  });
  t.is(response.status, 200);
  t.is(response.body.success, true);
  t.is(response.body.message, 'Successfully invoked the chaincode for function '+ CONSTANTS.FCN_CREATE_USER +' to the channel ' + CONSTANTS.CHANNEL_NAME);
});
