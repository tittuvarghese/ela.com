'use strict';

const CONSTANTS = require('../constants.js');

let util = require('util');
let helper = require('./helper.js');
let logger = helper.getLogger('install-chaincode');

/**
* installChaincode - function to install chaincode on peers
* @params peers {array}, chaincodeName, chaincodePath, chaincodeVersion
* @params chaincodeType, username, org_name
* @returns response {JSON_Array}
*/
let installChaincode = async function(peers, chaincodeName, chaincodePath,
	chaincodeVersion, chaincodeType, username, org_name) {
	logger.debug('\n\n============ Install chaincode on organizations ============\n');
	helper.setupChaincodeDeploy();
	let error_message = null;
	try {
		logger.info('Calling peers in organization "%s" to join the channel', org_name);

		// first setup the client for this org
		let client = await helper.getClientForOrg(org_name, username);
		logger.debug('Successfully got the fabric client for the organization "%s"', org_name);

		let request = {
			targets: peers,
			chaincodePath: chaincodePath,
			chaincodeId: chaincodeName,
			chaincodeVersion: chaincodeVersion,
			chaincodeType: chaincodeType
		};
		let results = await client.installChaincode(request);
		// the returned object has both the endorsement results
		// and the actual proposal, the proposal will be needed
		// later when we send a transaction to the orederer
		let proposalResponses = results[0];
		let proposal = results[1];

		// lets have a look at the responses to see if they are
		// all good, if good they will also include signatures
		// required to be committed
		let all_good = true;
		for (let i in proposalResponses) {
			let one_good = false;
			if (proposalResponses && proposalResponses[i].response &&
				proposalResponses[i].response.status === 200) {
				one_good = true;
				logger.info('install proposal was good');
			} else {
				logger.error('install proposal was bad');
			}
			all_good = all_good & one_good;
		}
		if (all_good) {
			logger.info('Successfully sent install Proposal and received ProposalResponse');
		} else {
			error_message = 'Failed to send install Proposal or receive valid response. Response null or status is not 200'
			logger.error(error_message);
		}
	} catch(error) {
		logger.error('Failed to install due to error: ' + error.stack ? error.stack : error);
		error_message = error.toString();
	}

	if (!error_message) {
		let message = util.format('Successfully installed chaincode');
		logger.info(message);
		let response = {
			success: true,
			message: message
		};
		return response;
	} else {
		let message = util.format('Failed to install due to: %s',error_message);
		logger.error(message);
		let response = {
			success: false,
			message: message
		};
		return response;
	}
};
exports.installChaincode = installChaincode;
