## Ela Application
Ela blockchain application readme.

##

### Running Application for Test
```
ssh -i "bcoin.pem" ubuntu@107.20.69.123
cd project/directory
npm install
./runApp.sh
```

### Test
```
npm test
```

### Accessing APIs
Refer tests folder for sample API calls.

### Public IP
http://107.20.69.123:4000

### Running for Production
```
ssh -i "bcoin.pem" ubuntu@107.20.69.123
cd project/directory
npm install
./runApp.sh
npm production
```

### API Documentation
#### User Operations
```js
Function: createUser
METHOD: POST
URL : '/channels/mychannel/chaincodes/ela_cc/createUser'
```
```js
REQUEST
{
  "peers": ["peer0.org1.ela.com"],
  "args": ["ELP-USR001", {
    "user_id" : "ELP-USR001",
    "name": {
      "first_name": "Adam",
      "last_name": "Smith"
    },
    "email" : "user@org1.ela.com",
    "phone_number" : "+91-9876543210",
    "profile_image" : "profile_image_url_ipfs",
    "role" : "producer"
  }]
}

RESPONSE
{
    "success": true,
    "message": "Successfully invoked the chaincode for function createUser to the channel mychannel",
    "trxnID": "TxnID"
}
```
```js
Function: queryUser
METHOD: GET
URL : '/channels/mychannel/chaincodes/ela_cc/queryUser/ELP-USR001?peer=peer0.org2.ela.com'
```
```js
RESPONSE
{
  "success": true,
  "message": "Found results for queryUser - ELP-USR001",
  "data": {
    "user_id" : "ELP-USR001",
    "name": {
      "first_name": "Adam",
      "last_name": "Smith"
    },
    "email" : "user@org1.ela.com",
    "phone_number" : "+91-9876543210",
    "profile_image" : "profile_image_url_ipfs",
    "role" : "producer"
  }
}
```
```js
Function: updateUser
METHOD: POST
URL : '/channels/mychannel/chaincodes/ela_cc/updateUser'
```
```js
REQUEST
{
  "peers": ["peer0.org1.ela.com"],
  "args": ["ELP-USR001", {
    "name": {
      "first_name": "Adam",
      "last_name": "Smith"
    },
    "phone_number" : "+91-9876543210",
    "profile_image" : "profile_image_url_ipfs",
    "role" : "producer"
  }]
}

RESPONSE
{
    "success": true,
    "message": "Successfully invoked the chaincode for function updateUser to the channel mychannel",
    "trxnID": "TxnID"
}
```
