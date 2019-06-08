exports.ORG1_NAME = "Org1"
exports.ORG2_NAME = "Org2"

exports.ORG1_EMAIL = "org1.ela.com"
exports.ORG2_EMAIL = "org2.ela.com"

/* Authentication Tokens */
exports.ORG1_JWT = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjA4MzgwODIsInVzZXJuYW1lIjoidXNlckBvcmcxLmVsYS5jb20iLCJvcmdOYW1lIjoiT3JnMSIsImlhdCI6MTU1OTk3NDA4Mn0.tF5YEE58D1XauGfWJF30rbQZ9zDunfjyvm60nk2Mi_U';
exports.ORG2_JWT = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjA4MzgxMDksInVzZXJuYW1lIjoidXNlckBvcmcyLmVsYS5jb20iLCJvcmdOYW1lIjoiT3JnMiIsImlhdCI6MTU1OTk3NDEwOX0.VRWpZKHitTz25zeEz6bwrNLVSwOnkeN2ms8DDWyu69Y';

exports.CHANNEL_NAME = 'mychannel';
exports.CHANNEL_CONFIG_PATH = '../artifacts/channel/mychannel.tx';
exports.CHAINCODE_NAME = 'ela_cc';
exports.CHAINCODE_PATH = 'ela_cc';
exports.CHAINCODE_TYPE = 'golang';
exports.CHAINCODE_VERSION = 'v1.0';
exports.ANCHOR_CONFIG_PATH_ORG1 = "../artifacts/channel/Org1MSPanchors.tx";
exports.ANCHOR_CONFIG_PATH_ORG2 = "../artifacts/channel/Org2MSPanchors.tx";

exports.FCN_INIT = "init";
exports.FCN_CREATE_USER = "createUser"
exports.FCN_QUERY_USER = "queryUser"
exports.FCN_UPDATE_USER = "updateUser"

/* Timeout Settings */
exports.JOIN_CHANNEL_TIMEOUT = 10000;
exports.INSTANTIATION_TIMEOUT = 600000;
exports.INVOKE_TIMEOUT = 40000;
