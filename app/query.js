const CONSTANTS = require('../constants.js');
let util = require('util');
let helper = require('./helper.js');
let logger = helper.getLogger('Query');

/**
* queryChaincode - To query chaincode
* @params peer, channelName, chaincodeName, args, fcn, username, org_name
* @return response {JSON_Array}
*/

let queryChaincode = async function(peer, channelName, chaincodeName, args, fcn, username, org_name) {
	try {

		let searchKey = args[0];

		// first setup the client for this org
		let client = await helper.getClientForOrg(org_name, username);
		logger.debug('Successfully got the fabric client for the organization "%s"', org_name);
		let channel = client.getChannel(channelName);
		if(!channel) {
			let message = util.format('Channel %s was not defined in the connection profile', channelName);
			logger.error(message);
			let response = {
				success: false,
				message: message
			}
			return response;
		}
		// send query
		let request = {
			targets : [peer], //queryByChaincode allows for multiple targets
			chaincodeId: chaincodeName,
			args: args,
			fcn: fcn,
		};

		let response_payloads = await channel.queryByChaincode(request);
		if (response_payloads) {
			for (let i = 0; i < response_payloads.length; i++) {
				logger.info(fcn + ' have ' + response_payloads[i].toString('utf8'));
			}
			let response = null;
			if(typeof(response_payloads[0]) == 'object') {
				try {
					let result = JSON.parse((response_payloads[0].toString('utf8')));
						response = {
							success: true,
							message: 'Found results for ' + fcn + ' - ' + searchKey,
							data: result
						}
				} catch (e) {
					response = {
						success: false,
						message: response_payloads[0].toString('utf8')
					}
				}
			} else {
				response = {
					success: false,
					message: "Unable to perform this operation."
				}
			}
			return response;
		} else {
			logger.error('response_payloads is null');
			response = {
				success: false,
				message: "response_payloads is null"
			}
			return response;
		}
	} catch(error) {
		logger.error('Failed to query due to error: ' + error.stack ? error.stack : error);
		response = {
			success: false,
			message: error.toString()
		}
		return response;
	}
};

/**
* getBlockByNumber - To query chaincode to get block by block number
* @params peer, channelName, blockNumber, username, org_name
* @return response {JSON_Array}
*/

let getBlockByNumber = async function(peer, channelName, blockNumber, username, org_name) {
	try {
		// first setup the client for this org
		let client = await helper.getClientForOrg(org_name, username);
		logger.debug('Successfully got the fabric client for the organization "%s"', org_name);
		let channel = client.getChannel(channelName);
		if(!channel) {
			let message = util.format('Channel %s was not defined in the connection profile', channelName);
			logger.error(message);
			let response = {
				success: false,
				message: message
			}
			return response;
		}

		let response_payload = await channel.queryBlock(parseInt(blockNumber, peer));
		if (response_payload) {
			logger.debug(response_payload);
			return response_payload;
		} else {
			logger.error('response_payload is null');
			let response = {
				success: false,
				message: "response_payload is null"
			}
			return response;
		}
	} catch(error) {
		logger.error('Failed to query due to error: ' + error.stack ? error.stack : error);
		let response = {
			success: false,
			message: error.toString()
		}
		return response;
	}
};

/**
* getTransactionByID - To query chaincode to get transaction by txn id
* @params peer, channelName, trxnID, username, org_name
* @return response {JSON_Array}
*/

let getTransactionByID = async function(peer, channelName, trxnID, username, org_name) {
	try {
		// first setup the client for this org
		let client = await helper.getClientForOrg(org_name, username);
		logger.debug('Successfully got the fabric client for the organization "%s"', org_name);
		let channel = client.getChannel(channelName);
		if(!channel) {
			let message = util.format('Channel %s was not defined in the connection profile', channelName);
			logger.error(message);
			let response = {
				success: false,
				message: message
			}
			return response;
		}

		let response_payload = await channel.queryTransaction(trxnID, peer);
		if (response_payload) {
			logger.debug(response_payload);
			return response_payload;
		} else {
			logger.error('response_payload is null');
			let response = {
				success: false,
				message: "response_payload is null"
			}
			return response;
		}
	} catch(error) {
		logger.error('Failed to query due to error: ' + error.stack ? error.stack : error);
		let response = {
			success: false,
			message: error.toString()
		}
		return response;
	}
};

/**
* getBlockByHash - To query chaincode to get block by block hash
* @params peer, channelName, hash, username, org_name
* @return response {JSON_Array}
*/

let getBlockByHash = async function(peer, channelName, hash, username, org_name) {
	try {
		// first setup the client for this org
		let client = await helper.getClientForOrg(org_name, username);
		logger.debug('Successfully got the fabric client for the organization "%s"', org_name);
		let channel = client.getChannel(channelName);
		if(!channel) {
			let message = util.format('Channel %s was not defined in the connection profile', channelName);
			logger.error(message);
			let response = {
				success: false,
				message: message
			}
			return response;
		}

		let response_payload = await channel.queryBlockByHash(Buffer.from(hash,'hex'), peer);
		if (response_payload) {
			logger.debug(response_payload);
			return response_payload;
		} else {
			logger.error('response_payload is null');
			let response = {
				success: false,
				message: "response_payload is null"
			}
			return response;
		}
	} catch(error) {
		logger.error('Failed to query due to error: ' + error.stack ? error.stack : error);
		let response = {
			success: false,
			message: error.toString()
		}
		return response;
	}
};

/**
* getChainInfo - To query chaincode to get chain info
* @params peer, channelName, username, org_name
* @return response {JSON_Array}
*/

let getChainInfo = async function(peer, channelName, username, org_name) {
	try {
		// first setup the client for this org
		let client = await helper.getClientForOrg(org_name, username);
		logger.debug('Successfully got the fabric client for the organization "%s"', org_name);
		let channel = client.getChannel(channelName);
		console.log('Channel ===> ' + channel);
		if(!channel) {
			let message = util.format('Channel %s was not defined in the connection profile', channelName);
			logger.error(message);
			//throw new Error(message);
			let response = {
				success: false,
				message: 'Channel %s was not defined in the connection profile'
			}
			return response;
		}

		let response_payload = await channel.queryInfo(peer);
		console.log('response_payload ===> ' + response_payload);
		if (response_payload) {
			logger.debug(response_payload);
			let response = {
				success: true,
				message: response_payload
			}
			return response;
		} else {
			logger.error('response_payload is null');
			let response = {
				success: false,
				message: 'response_payload is null'
			}
			return response;
		}
	} catch(error) {
		logger.error('Failed to query due to error: ' + error.stack ? error.stack : error);
		let response = {
			success: false,
			message: error.toString()
		}
		return response;
	}
};

/**
* getInstalledChaincodes - To query chaincode to get installed chaincodes
* @params peer, channelName, type, username, org_name
* @params type => installed, instantiated
* @return response {JSON_Array}
*/

let getInstalledChaincodes = async function(peer, channelName, type, username, org_name) {
	try {
		// first setup the client for this org
		let client = await helper.getClientForOrg(org_name, username);
		logger.debug('Successfully got the fabric client for the organization "%s"', org_name);

		let result = null
		if (type === 'installed') {
			result = await client.queryInstalledChaincodes(peer, true); //use the admin identity
		} else {
			let channel = client.getChannel(channelName);
			if(!channel) {
				let message = util.format('Channel %s was not defined in the connection profile', channelName);
				logger.error(message);
				// throw new Error(message);
				let response = {
					success: false,
					message: 'Channel %s was not defined in the connection profile'
				}
				return response;
			}
			result = await channel.queryInstantiatedChaincodes(peer, true); //use the admin identity
		}
		if (result) {
			if (type === 'installed') {
				logger.debug('<<< Installed Chaincodes >>>');
			} else {
				logger.debug('<<< Instantiated Chaincodes >>>');
			}
			let details = [];
			for (let i = 0; i < result.chaincodes.length; i++) {
				logger.debug('name: ' + result.chaincodes[i].name + ', version: ' +
					result.chaincodes[i].version + ', path: ' + result.chaincodes[i].path
				);
				let chaincodeDetails = {
					name: result.chaincodes[i].name,
					version: result.chaincodes[i].version,
					path: result.chaincodes[i].path,
				}
				details.push(chaincodeDetails);
			}
			let response = {
				success: true,
				message: details
			}
			return response;
		} else {
			logger.error('response is null, no chaincodes installed');
			let response = {
				success: false,
				message: 'Response is null, no chaincodes installed'
			}
			return response;
		}
	} catch(error) {
		logger.error('Failed to query due to error: ' + error.stack ? error.stack : error);
		let response = {
			success: false,
			message: error.toString()
		}
		return response;
	}
};

/**
* getInstalledChaincodes - To query chaincode to get channels in the network
* @params peer, username, org_name
* @return response {JSON_Array}
*/

let getChannels = async function(peer, username, org_name) {
	try {
		// first setup the client for this org
		let client = await helper.getClientForOrg(org_name, username);
		logger.debug('Successfully got the fabric client for the organization "%s"', org_name);

		let result = await client.queryChannels(peer);
		if (result) {
			logger.debug('<<< channels >>>');
			let channelNames = [];
			for (let i = 0; i < result.channels.length; i++) {
				channelNames.push('channel id: ' + result.channels[i].channel_id);
			}
			logger.debug(channelNames);
			let response = {
				success: true,
				message: result
			}
			return response;
		} else {
			logger.error('response_payloads is null');
			let response = {
				success: false,
				message: 'response_payloads is null'
			}
			return response;
		}
	} catch(error) {
		logger.error('Failed to query due to error: ' + error.stack ? error.stack : error);
		let response = {
			success: false,
			message: error.toString()
		}
		return response;
	}
};

exports.queryChaincode = queryChaincode;
exports.getBlockByNumber = getBlockByNumber;
exports.getTransactionByID = getTransactionByID;
exports.getBlockByHash = getBlockByHash;
exports.getChainInfo = getChainInfo;
exports.getInstalledChaincodes = getInstalledChaincodes;
exports.getChannels = getChannels;
